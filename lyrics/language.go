package lyrics

import (
	"fmt"
	"strings"
)

type Language string

const (
	LANGUAGE_ALL Language = "all"
	LANGUAGE_ZH  Language = "zh"
	LANGUAGE_JA  Language = "ja"
)

type InvalidLanguageError struct {
	Text string
}

func (e *InvalidLanguageError) Error() string {
	return fmt.Sprintf("invalid language: %s", e.Text)
}

func ParseLanguage(text string) (Language, error) {
	languages := map[Language]struct{}{
		LANGUAGE_ALL: {},
		LANGUAGE_ZH:  {},
		LANGUAGE_JA:  {},
	}

	text = strings.TrimSpace(strings.ToLower(text))
	language := Language(text)
	_, ok := languages[language]
	if !ok {
		return LANGUAGE_ALL, &InvalidLanguageError{Text: text}
	}
	return language, nil
}

func (language Language) String() string {
	return string(language)
}
