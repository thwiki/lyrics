package lyrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/thwiki/lyrics/lrc"
)

type AskTrackQuery struct {
	Name       *[]string `json:"name"`
	AlName     *[]string `json:"alname"`
	CircleName *[]string `json:"circlename"`
	Cover      *[]string `json:"cover2x"`
	Lyrics     *[]string `json:"lyrics"`
}

type AskTrackResponse struct {
	Results    []AskTrackResponseResult `json:"results"`
	Error      string                   `json:"error,omitempty"`
	HTTPCode   int                      `json:"httpCode,omitempty"`
	HTTPReason string                   `json:"httpReason,omitempty"`
}

type AskTrackResponseResult struct {
	ID         string                        `json:"id"`
	Name       []string                      `json:"name"`
	AlName     []string                      `json:"alname"`
	CircleName []string                      `json:"circlename"`
	Cover      []AskTrackResponseResultCover `json:"cover2x"`
}

type AskTrackResponseResultCover struct {
	FullText     string `json:"fulltext"`
	FullUrl      string `json:"fullurl"`
	Namespace    int    `json:"namespace"`
	Exists       bool   `json:"exists"`
	Displaytitle string `json:"displaytitle"`
}

func (r *AskTrackResponse) FromRequest(request *Request) (err error) {
	source := os.Getenv("ASKTRACK_API_SOURCE")
	if err != nil {
		return
	}

	query := AskTrackQuery{
		Lyrics: &[]string{os.Getenv("SOURCE_NAMESPACE") + request.Title},
	}
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return
	}

	resp, err := http.Post(source, "application/json", bytes.NewBuffer(queryBytes))
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

func (r *AskTrackResponse) AddMeta(request *Request, document *lrc.Document) (err error) {
	if r.Error != "" {
		err = fmt.Errorf("Error: %s", r.HTTPReason)
		return
	}

	if len(r.Results) == 0 {
		return
	}

	track := r.Results[0]

	document.Meta.Title = strings.Join(track.Name, "/")
	document.Meta.Album = strings.Join(track.AlName, "/")
	document.Meta.Artist = strings.Join(track.CircleName, "/")
	if len(track.Cover) >= 1 && track.Cover[0].Exists {
		document.Meta.Cover = (track.Cover[0].FullUrl)
	}

	document.Meta.Editor = []string{
		fmt.Sprintf("%s: https://%s/%s", os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_HOST"), request.String()),
		fmt.Sprintf("%s: https://%s/%s", os.Getenv("SOURCE_NAME"), os.Getenv("SOURCE_HOST"), os.Getenv("SOURCE_NAMESPACE")+request.Title),
	}

	return
}
