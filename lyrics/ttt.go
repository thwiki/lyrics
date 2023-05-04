package lyrics

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/thwiki/lyrics/lrc"
)

type TTTResponse struct {
	Table []*TTTResponseRow `json:"table,omitempty"`
	Error *TTTResponseError `json:"error,omitempty"`
}

type TTTResponseError struct {
	Code    string `json:"code"`
	Info    string `json:"info"`
	Content string `json:"*"`
}

type TTTResponseRow struct {
	Type  string           `json:"type"`
	Index int64            `json:"index"`
	Cell0 *TTTResponseCell `json:"0,omitempty"`
	Cell1 *TTTResponseCell `json:"1,omitempty"`
	Cell2 *TTTResponseCell `json:"2,omitempty"`
}

type TTTResponseCell struct {
	Lang    string `json:"lang"`
	Content string `json:"*"`
}

var (
	spaceRegex  = regexp.MustCompile(`\s+`)
	authorRegex = regexp.MustCompile(`(?:歌词翻译：)(.+\s+（.+）)`)
	tabRegex    = regexp.MustCompile(`(?:\|-\||<[^>]*>)?(.+)=\s*`)
)

func sanitizeAuthor(text string) (name string) {
	match := authorRegex.FindStringSubmatch(text)
	if match == nil {
		name = ""
	} else {
		name = sanitizeHtml(match[1])
	}

	return
}

func sanitizeTabName(text string) (name string) {
	match := tabRegex.FindStringSubmatch(text)
	if match == nil {
		name = "默认"
		return
	}
	name = strings.TrimSpace(match[1])
	if name == "" {
		name = "默认"
	}
	return
}

func sanitizeHtml(text string) string {
	p := &bluemonday.Policy{}
	p.SkipElementsContent(
		"abbr",
		"acronym",
		"address",
		"article",
		"aside",
		"bdi",
		"blockquote",
		"br",
		"button",
		"canvas",
		"caption",
		"center",
		"cite",
		"code",
		"col",
		"colgroup",
		"datalist",
		"dd",
		"details",
		"dfn",
		"div",
		"dl",
		"dt",
		"fieldset",
		"figcaption",
		"figure",
		"footer",
		"h1",
		"h2",
		"h3",
		"h4",
		"h5",
		"h6",
		"header",
		"hgroup",
		"hr",
		"html",
		"kbd",
		"mark",
		"marquee",
		"nav",
		"ol",
		"optgroup",
		"option",
		"p",
		"pre",
		"q",
		"rp",
		"rt",
		"samp",
		"script",
		"section",
		"select",
		"strike",
		"style",
		"sub",
		"summary",
		"sup",
		"table",
		"tbody",
		"td",
		"tfoot",
		"th",
		"thead",
		"title",
		"time",
		"tr",
		"tt",
		"var",
		"wbr",
	)
	p.AllowElementsContent("span", "a", "small", "s", "del", "em", "ins", "b", "strong", "i", "u", "font", "ul", "li", "rb", "ruby")

	text = p.Sanitize(text)

	text = html.UnescapeString(text)
	text = spaceRegex.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)
	return text
}

func getTTTSource(request *Request) (string, error) {
	source, err := url.Parse(os.Getenv("TTT_API_SOURCE"))

	if err != nil {
		return "", err
	}

	query := source.Query()
	query.Set("action", "ttt")
	query.Set("format", "json")
	query.Set("title", "Lyrics:"+request.Title)
	query.Set("lang", "xx,lyrics,sep,time")
	query.Set("parse", "")
	source.RawQuery = query.Encode()

	return source.String(), nil
}

func (r *TTTResponse) FromRequest(request *Request) (err error) {
	source, err := getTTTSource(request)
	if err != nil {
		return
	}

	resp, err := http.Get(source)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonBytes, &r)

	return
}

func (r *TTTResponse) AddLines(request *Request, document *lrc.Document) (err error) {
	if r.Error != nil && r.Error.Info != "" {
		err = fmt.Errorf("Error: %s", r.Error.Info)
		return
	}

	table := r.Table
	index := 0
	document.Meta.Relations = make([]string, 0)
	isInfo := true

	for _, row := range table {
		if row.Type == "xx" {
			header := row.FindLang("xx")
			if isInfo {
				isInfo = false
				document.Meta.Author = sanitizeAuthor(header)
			} else {
				nextRequest := Request{
					Title:     request.Title,
					Index:     index + 1,
					Language:  request.Language,
					Extension: request.Extension,
				}
				document.Meta.Relations = append(document.Meta.Relations, fmt.Sprintf("%s: https://%s/%s", sanitizeTabName(header), os.Getenv("SERVICE_HOST"), nextRequest.String()))
			}
		} else if row.Type == "lyrics-header" {
			index++
		}
		if index < request.Index {
			continue
		} else if index > request.Index {
			break
		}

		typeParts := append(strings.SplitN(row.Type, "-", 2), "", "")
		firstPart := typeParts[0]
		secondPart := typeParts[1]

		if secondPart == "sep" {
			document.Lines = append(document.Lines, lrc.Line{Time: row.FindLang("sep")})
		} else if firstPart == "main" {
			if secondPart == "zh" {
				document.Lines = append(document.Lines, lrc.Line{Time: row.FindLang("time"), Text: row.FindLang("zh")})
			} else {
				if request.Language == LANGUAGE_ALL {
					lineText := row.FindLang(secondPart)
					zhText := row.FindLang("zh")
					if zhText != "" {
						lineText += " // " + zhText
					}
					document.Lines = append(document.Lines, lrc.Line{Time: row.FindLang("time"), Text: lineText})
				} else if request.Language == "zh" {
					zhText := row.FindLang("zh")
					if zhText == "" {
						zhText = row.FindLang("ja")
					}
					document.Lines = append(document.Lines, lrc.Line{Time: row.FindLang("time"), Text: zhText})
				} else {
					document.Lines = append(document.Lines, lrc.Line{Time: row.FindLang("time"), Text: row.FindLang(secondPart)})
				}
			}
		}
	}

	return
}

func (row *TTTResponseRow) FindLang(lang string) string {
	cells := []*TTTResponseCell{row.Cell0, row.Cell1, row.Cell2}

	for _, cell := range cells {
		if cell != nil {
			isMatch := cell.Lang == lang
			if !isMatch && lang == "ja" {
				if cell.Lang != "ja" && cell.Lang != "zh" && cell.Lang != "time" && cell.Lang != "sep" {
					isMatch = true
				}
			}
			if isMatch {
				return sanitizeHtml(cell.Content)
			}
		}
	}
	return ""
}
