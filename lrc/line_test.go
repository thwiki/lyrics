package lrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineString(t *testing.T) {
	line := Line{
		Time: "12:34.56",
		Text: "line 1",
	}

	str := line.String()

	assert.Equal(t, `[12:34.56]line 1
`, str)
}

func TestLineStringWithoutText(t *testing.T) {
	line := Line{
		Time: "12:34.56",
	}

	str := line.String()

	assert.Equal(t, `[12:34.56]
`, str)
}
