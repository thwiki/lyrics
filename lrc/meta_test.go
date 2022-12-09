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
		Editor: "Editor Name",
	}

	str := meta.String()

	assert.Equal(t, `[ti:Track Name]
[al:Album Name]
[ar:Artist Name]
[re:Editor Name]
`, str)
}

func TestMetaStringPartialMeta(t *testing.T) {
	meta := Meta{
		Title:  "ABCDEF",
		Editor: "efghij",
	}

	str := meta.String()

	assert.Equal(t, `[ti:ABCDEF]
[re:efghij]
`, str)
}
