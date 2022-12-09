package lrc

import (
	"fmt"
	"reflect"
)

type Meta struct {
	Title  string `json:"ti" lrc:"ti"`
	Album  string `json:"al" lrc:"al"`
	Artist string `json:"ar" lrc:"ar"`
	Editor string `json:"re" lrc:"re"`
	Length string `json:"length" lrc:"length"`
	Offset string `json:"offset" lrc:"offset"`
}

func (meta *Meta) String() (result string) {
	values := reflect.ValueOf(*meta)
	types := values.Type()

	for index := 0; index < values.NumField(); index++ {
		name := types.Field(index).Tag.Get("lrc")
		value := values.Field(index).String()
		if value != "" {
			result += fmt.Sprintf("[%s:%s]\n", name, value)
		}
	}

	return
}
