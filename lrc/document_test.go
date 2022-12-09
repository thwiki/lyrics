package lrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	document := NewDocument()

	assert.IsType(t, Document{}, document)
}

func TestDocumentString(t *testing.T) {
	document := Document{
		Meta: Meta{
			Title:  "Track Name",
			Album:  "Album Name",
			Artist: "Artist Name",
			Editor: "Editor Name",
		},
		Lines: []Line{
			{
				Time: "00:01.23",
				Text: "line 1",
			},
			{
				Time: "00:04.56",
				Text: "line 2",
			},
			{
				Time: "00:12.34",
				Text: "",
			},
		},
	}

	str := document.String()

	assert.Equal(t, `[ti:Track Name]
[al:Album Name]
[ar:Artist Name]
[re:Editor Name]
[00:01.23]line 1
[00:04.56]line 2
[00:12.34]
`, str)
}
