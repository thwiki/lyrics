package lrc

import (
	"fmt"
	"reflect"
)

type Meta struct {
	Title     string   `json:"title" lrc:"ti"`
	Album     string   `json:"album" lrc:"al"`
	Artist    string   `json:"artist" lrc:"ar"`
	Cover     string   `json:"cover" lrc:"cv"`
	Editor    []string `json:"editor" lrc:"re"`
	Author    string   `json:"author" lrc:"by"`
	Length    string   `json:"length" lrc:"length"`
	Offset    string   `json:"offset" lrc:"offset"`
	Relations []string `json:"relations" lrc:"rel"`
}

func (meta *Meta) String() (result string) {
	values := reflect.ValueOf(*meta)
	types := values.Type()

	for index := 0; index < values.NumField(); index++ {
		name := types.Field(index).Tag.Get("lrc")
		value := values.Field(index)

		if types.Field(index).Type.Kind() != reflect.Slice {
			text := value.String()
			if text != "" {
				result += fmt.Sprintf("[%s:%s]\n", name, text)
			}
			continue
		}
		if texts, ok := value.Interface().([]string); ok {
			for _, text := range texts {
				if text != "" {
					result += fmt.Sprintf("[%s:%s]\n", name, text)
				}
			}
			continue
		}
	}

	return
}
