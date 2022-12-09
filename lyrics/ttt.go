package lyrics

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
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

func getTTTSource(request *Request) (string, error) {
	source, err := url.Parse(os.Getenv("TTT_API_SOURCE"))

	if err != nil {
		return "", err
	}

	query := source.Query()
	query.Set("action", "ttt")
	query.Set("format", "json")
	query.Set("title", "Lyrics:"+request.Title)
	query.Set("lang", "")
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

	for _, row := range table {
		if row.Type == "lyrics-header" {
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
				return strings.TrimSpace(strings.ReplaceAll(html.UnescapeString(strip.StripTags(cell.Content)), "\n", " "))
			}
		}
	}
	return ""
}
