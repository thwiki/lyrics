package lyrics

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thwiki/lyrics/lrc"
)

func getMockTTTResponse() TTTResponse {
	var tttResponse TTTResponse
	json.Unmarshal([]byte(`{
		"table": [
			{ "type": "lyrics-header", "index": 2 },
			{
				"0": { "lang": "time", "*": "00:01.23" },
				"1": { "lang": "ja", "*": "Section 1 Line 1" },
				"2": { "lang": "zh", "*": "段落 1 第 1 句" },
				"type": "main-ja",
				"index": 3
			},
			{ "0": { "lang": "sep", "*": "00:02.34" }, "type": "lyrics-sep", "index": 4 },
			{
				"0": { "lang": "time", "*": "00:03.45" },
				"1": { "lang": "en", "*": "Section 1 Line 2" },
				"2": { "lang": "zh", "*": "段落 1 第 2 句" },
				"type": "main-ja",
				"index": 5
			},
			{ "0": { "lang": "sep", "*": "01:23.45" }, "type": "lyrics-sep", "index": 6 },
			{
				"0": { "lang": "time", "*": "02:34.56" },
				"1": { "lang": "ja", "*": "Section 1 Line 3" },
				"2": { "lang": "zh", "*": "段落 1 第 3 句" },
				"type": "main-ja",
				"index": 7
			},
			{ "type": "lyrics-header", "index": 9 },
			{
				"0": { "lang": "time", "*": "00:11.23" },
				"1": { "lang": "ja", "*": "Section 2 Line 1" },
				"2": { "lang": "zh", "*": "段落 2 第 1 句" },
				"type": "main-ja",
				"index": 10
			},
			{ "0": { "lang": "sep", "*": "00:22.34" }, "type": "lyrics-sep", "index": 11 },
			{
				"0": { "lang": "time", "*": "00:33.45" },
				"1": { "lang": "en", "*": "Section 2 Line 2" },
				"2": { "lang": "zh", "*": "段落 2 第 2 句" },
				"type": "main-ja",
				"index": 12
			},
			{ "0": { "lang": "sep", "*": "02:23.45" }, "type": "lyrics-sep", "index": 13 },
			{
				"0": { "lang": "time", "*": "03:34.56" },
				"1": { "lang": "ja", "*": "Section 2 Line 3" },
				"2": { "lang": "zh", "*": "段落 2 第 3 句" },
				"type": "main-ja",
				"index": 14
			}
		]
	}`), &tttResponse)
	return tttResponse
}

func TestGetTTTSource(t *testing.T) {
	request := Request{
		Title:     "中文 test.track_结尾",
		Index:     12,
		Language:  LANGUAGE_JA,
		Extension: "lrc",
	}

	os.Setenv("TTT_API_SOURCE", "https://example.com/api.php")
	os.Setenv("SOURCE_NAMESPACE", "Lyrics:")

	str, err := getTTTSource(&request)

	assert.Equal(t, "https://example.com/api.php?action=ttt&format=json&lang=xx%2Clyrics%2Csep%2Ctime&parse=&title=Lyrics%3A%E4%B8%AD%E6%96%87+test.track_%E7%BB%93%E5%B0%BE", str)
	assert.Nil(t, err)
}

func TestTTTResponseAddLinesIndex1ALL(t *testing.T) {
	tttResponse := getMockTTTResponse()

	document := lrc.NewDocument()
	err := tttResponse.AddLines(&Request{Index: 1, Language: LANGUAGE_ALL}, &document)

	assert.Nil(t, err)
	assert.Len(t, document.Lines, 5)
	assert.Equal(t, []lrc.Line{
		{Time: "00:01.23", Text: "Section 1 Line 1 // 段落 1 第 1 句"},
		{Time: "00:02.34", Text: ""},
		{Time: "00:03.45", Text: "Section 1 Line 2 // 段落 1 第 2 句"},
		{Time: "01:23.45", Text: ""},
		{Time: "02:34.56", Text: "Section 1 Line 3 // 段落 1 第 3 句"},
	}, document.Lines)
}

func TestTTTResponseAddLinesIndex1JA(t *testing.T) {
	tttResponse := getMockTTTResponse()

	document := lrc.NewDocument()
	err := tttResponse.AddLines(&Request{Index: 1, Language: LANGUAGE_JA}, &document)

	assert.Nil(t, err)
	assert.Len(t, document.Lines, 5)
	assert.Equal(t, []lrc.Line{
		{Time: "00:01.23", Text: "Section 1 Line 1"},
		{Time: "00:02.34", Text: ""},
		{Time: "00:03.45", Text: "Section 1 Line 2"},
		{Time: "01:23.45", Text: ""},
		{Time: "02:34.56", Text: "Section 1 Line 3"},
	}, document.Lines)
}

func TestTTTResponseAddLinesIndex2ALL(t *testing.T) {
	tttResponse := getMockTTTResponse()

	document := lrc.NewDocument()
	err := tttResponse.AddLines(&Request{Index: 2, Language: LANGUAGE_ALL}, &document)

	assert.Nil(t, err)
	assert.Len(t, document.Lines, 5)
	assert.Equal(t, []lrc.Line{
		{Time: "00:11.23", Text: "Section 2 Line 1 // 段落 2 第 1 句"},
		{Time: "00:22.34", Text: ""},
		{Time: "00:33.45", Text: "Section 2 Line 2 // 段落 2 第 2 句"},
		{Time: "02:23.45", Text: ""},
		{Time: "03:34.56", Text: "Section 2 Line 3 // 段落 2 第 3 句"},
	}, document.Lines)
}

func TestTTTResponseAddLinesIndex2JA(t *testing.T) {
	tttResponse := getMockTTTResponse()

	document := lrc.NewDocument()
	err := tttResponse.AddLines(&Request{Index: 2, Language: LANGUAGE_JA}, &document)

	assert.Nil(t, err)
	assert.Len(t, document.Lines, 5)
	assert.Equal(t, []lrc.Line{
		{Time: "00:11.23", Text: "Section 2 Line 1"},
		{Time: "00:22.34", Text: ""},
		{Time: "00:33.45", Text: "Section 2 Line 2"},
		{Time: "02:23.45", Text: ""},
		{Time: "03:34.56", Text: "Section 2 Line 3"},
	}, document.Lines)
}

func TestSanitizeTabName(t *testing.T) {
	var text, result string

	text = "原版 = "
	result = sanitizeTabName(text)
	assert.Equal(t, "原版", result)

	text = "<tabber>  原版 = "
	result = sanitizeTabName(text)
	assert.Equal(t, "原版", result)

	text = "|-| extended mix = "
	result = sanitizeTabName(text)
	assert.Equal(t, "extended mix", result)

	text = " = "
	result = sanitizeTabName(text)
	assert.Equal(t, "默认", result)
}

func TestSanitizeTime(t *testing.T) {
	var text, result string

	text = "12:34.56"
	result = sanitizeTime(text)
	assert.Equal(t, "12:34.56", result)

	text = "12:34:56"
	result = sanitizeTime(text)
	assert.Equal(t, "12:34.56", result)

	text = " 12:34.56  "
	result = sanitizeTime(text)
	assert.Equal(t, "12:34.56", result)

	text = " 12:34:56  "
	result = sanitizeTime(text)
	assert.Equal(t, "12:34.56", result)

	text = "12:34"
	result = sanitizeTime(text)
	assert.Equal(t, "12:34.00", result)

	text = "34.56"
	result = sanitizeTime(text)
	assert.Equal(t, "00:34.56", result)

	text = "34:56"
	result = sanitizeTime(text)
	assert.Equal(t, "34:56.00", result)

	text = "2:34"
	result = sanitizeTime(text)
	assert.Equal(t, "02:34.00", result)

	text = "4.56"
	result = sanitizeTime(text)
	assert.Equal(t, "00:04.56", result)

	text = "4:56"
	result = sanitizeTime(text)
	assert.Equal(t, "04:56.00", result)

	text = "1:2.3"
	result = sanitizeTime(text)
	assert.Equal(t, "01:02.3", result)

	text = "123:456.789"
	result = sanitizeTime(text)
	assert.Equal(t, "123:456.789", result)

	text = " "
	result = sanitizeTime(text)
	assert.Equal(t, "00:00.00", result)

	text = "-"
	result = sanitizeTime(text)
	assert.Equal(t, "00:00.00", result)
}

func TestSanitizeHtml(t *testing.T) {
	var text, result string

	text = `abc <sup id="cite_ref-unsung_1-0" class="reference"><a href="#cite_note-unsung-1">1</a></sup> def</div>
<div class="mw-references-wrap"><ol class="references">
<li id="cite_note-unsung-1"><span class="mw-cite-backlink"><a href="#cite_ref-unsung_1-0">↑</a></span> <span class="reference-text">本句没被唱出</span>
</li>
</ol>`
	result = sanitizeHtml(text)
	assert.Equal(t, "abc def", result)

	text = `abc <b>def</b> ghi`
	result = sanitizeHtml(text)
	assert.Equal(t, "abc def ghi", result)

	text = `abc <span style="color:purple;">def</span> ghi`
	result = sanitizeHtml(text)
	assert.Equal(t, "abc def ghi", result)

	text = `abc <ruby lang="en"><rb>def</rb><rp> (</rp><rt>123</rt><rp>) </rp></ruby> ghi`
	result = sanitizeHtml(text)
	assert.Equal(t, "abc def ghi", result)

	text = `abc <unkown>def</unkown> ghi`
	result = sanitizeHtml(text)
	assert.Equal(t, "abc def ghi", result)
}
