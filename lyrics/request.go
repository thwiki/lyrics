package lyrics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thwiki/lyrics/lrc"
)

type Request struct {
	Title     string   `json:"title"`
	Index     int      `json:"index"`
	Language  Language `json:"language"`
	Extension string   `json:"extension"`
}

func (request *Request) FromName(name string) (err error) {
	parts := strings.Split(name, ".")
	lastIndex := len(parts) - 1

	request.Index = 1
	request.Language = LANGUAGE_ALL
	request.Extension = "lrc"

	if lastIndex > 0 {
		if parts[lastIndex] == "lrc" {
			request.Extension = parts[lastIndex]
			lastIndex--
		}
	}

	if lastIndex > 0 {
		if language, err := ParseLanguage(parts[lastIndex]); err == nil {
			request.Language = language
			lastIndex--
		}
	}

	if lastIndex > 0 && len(parts[lastIndex]) <= 2 {
		if index, err := strconv.Atoi(parts[lastIndex]); err == nil {
			request.Index = index
			lastIndex--
		}
	}

	request.Title = strings.Join(parts[0:lastIndex+1], ".")

	return
}

func (request *Request) String() string {
	return strings.ReplaceAll(fmt.Sprintf("%s.%d.%s.%s", request.Title, request.Index, request.Language, request.Extension), " ", "_")
}

func (request *Request) GetLrc() (string, error) {
	document := lrc.NewDocument()

	var tttResponse TTTResponse
	if err := tttResponse.FromRequest(request); err != nil {
		return "", err
	}
	if err := tttResponse.AddLines(request, &document); err != nil {
		return "", err
	}

	var askTrackResponse AskTrackResponse
	if err := askTrackResponse.FromRequest(request); err != nil {
		return "", err
	}
	if err := askTrackResponse.AddMeta(request, &document); err != nil {
		return "", err
	}

	return document.String(), nil
}
