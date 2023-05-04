package lrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetaStringMeta(t *testing.T) {
	meta := Meta{
		Title:  "Track Name",
		Album:  "Album Name",
		Artist: "Artist Name",
		Editor: []string{"Editor Name"},
		Relations: []string{
			"Relation 1",
			"",
			"Relation 2",
		},
	}

	str := meta.String()

	assert.Equal(t, `[ti:Track Name]
[al:Album Name]
[ar:Artist Name]
[re:Editor Name]
[rel:Relation 1]
[rel:Relation 2]
`, str)
}

func TestMetaStringPartialMeta(t *testing.T) {
	meta := Meta{
		Title:  "ABCDEF",
		Editor: []string{"efghij"},
	}

	str := meta.String()

	assert.Equal(t, `[ti:ABCDEF]
[re:efghij]
`, str)
}
