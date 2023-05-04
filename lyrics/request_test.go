package lyrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestFromNameFull(t *testing.T) {
	var request Request
	var err error = request.FromName("中文 test.track_结尾.42.zh.lrc")
	assert.Equal(t, "中文 test.track_结尾", request.Title)
	assert.Equal(t, 42, request.Index)
	assert.Equal(t, LANGUAGE_ZH, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)
}

func TestRequestFromNameWithoutExtension(t *testing.T) {
	var request Request
	var err error = request.FromName("中文 test.track_结尾.42.zh")
	assert.Equal(t, "中文 test.track_结尾", request.Title)
	assert.Equal(t, 42, request.Index)
	assert.Equal(t, LANGUAGE_ZH, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)
}

func TestRequestFromNameWithoutLanguage(t *testing.T) {
	var request Request
	var err error = request.FromName("中文 test.track_结尾.42.lrc")
	assert.Equal(t, "中文 test.track_结尾", request.Title)
	assert.Equal(t, 42, request.Index)
	assert.Equal(t, LANGUAGE_ALL, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)
}

func TestRequestFromNameWithoutIndex(t *testing.T) {
	var request Request
	var err error = request.FromName("中文 test.track_结尾.ja.lrc")
	assert.Equal(t, "中文 test.track_结尾", request.Title)
	assert.Equal(t, 1, request.Index)
	assert.Equal(t, LANGUAGE_JA, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)
}

func TestRequestFromNameWithoutIndexAndLanguage(t *testing.T) {
	var request Request
	var err error = request.FromName("中文 test.track_结尾.lrc")
	assert.Equal(t, "中文 test.track_结尾", request.Title)
	assert.Equal(t, 1, request.Index)
	assert.Equal(t, LANGUAGE_ALL, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)
}

func TestRequestFromNameNumericTitle(t *testing.T) {
	var request Request
	var err error

	err = request.FromName("1234.zh.lrc")
	assert.Equal(t, "1234", request.Title)
	assert.Equal(t, 1, request.Index)
	assert.Equal(t, LANGUAGE_ZH, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)

	err = request.FromName("12.34.ja.lrc")
	assert.Equal(t, "12", request.Title)
	assert.Equal(t, 34, request.Index)
	assert.Equal(t, LANGUAGE_JA, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)

	err = request.FromName("12.345.ja.lrc")
	assert.Equal(t, "12.345", request.Title)
	assert.Equal(t, 1, request.Index)
	assert.Equal(t, LANGUAGE_JA, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)
}

func TestRequestFromNameNumericTitleWithoutLanguage(t *testing.T) {
	var request Request
	var err error = request.FromName("1234.lrc")
	assert.Equal(t, "1234", request.Title)
	assert.Equal(t, 1, request.Index)
	assert.Equal(t, LANGUAGE_ALL, request.Language)
	assert.Equal(t, "lrc", request.Extension)
	assert.Nil(t, err)
}

func TestRequestString(t *testing.T) {
	request := Request{
		Title:     "中文 test.track_结尾",
		Index:     12,
		Language:  LANGUAGE_JA,
		Extension: "lrc",
	}
	var str string = request.String()
	assert.Equal(t, "中文_test.track_结尾.12.ja.lrc", str)
}
