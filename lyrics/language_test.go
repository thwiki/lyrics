package lyrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLanguageValidStandard(t *testing.T) {
	var language Language
	var err error

	texts := []string{"all", "zh", "ja"}
	languages := []string{"all", "zh", "ja"}

	// Standard
	for index, text := range texts {
		language, err = ParseLanguage(text)
		assert.Equal(t, languages[index], language.String())
		assert.Nil(t, err)
	}
}

func TestParseLanguageValidCases(t *testing.T) {
	var language Language
	var err error

	texts := []string{"All", "ZH", "jA"}
	languages := []string{"all", "zh", "ja"}

	// Standard
	for index, text := range texts {
		language, err = ParseLanguage(text)
		assert.Equal(t, languages[index], language.String())
		assert.Nil(t, err)
	}
}

func TestParseLanguageValidSpaces(t *testing.T) {
	var language Language
	var err error

	texts := []string{" All ", " zh", "ja "}
	languages := []string{"all", "zh", "ja"}

	// Standard
	for index, text := range texts {
		language, err = ParseLanguage(text)
		assert.Equal(t, languages[index], language.String())
		assert.Nil(t, err)
	}
}

func TestParseLanguageInvalid(t *testing.T) {
	var language Language
	var err error

	texts := []string{"123", "abc", "ja-"}
	languages := []string{"ja", "ja", "ja"}

	// Standard
	for index, text := range texts {
		language, err = ParseLanguage(text)
		assert.Equal(t, languages[index], language.String())
		assert.Error(t, &InvalidLanguageError{}, err)
	}
}
