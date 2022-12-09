package lrc

import "fmt"

type Line struct {
	Time string `json:"time"`
	Text string `json:"text"`
}

func (meta *Line) String() (result string) {
	return fmt.Sprintf("[%s]%s\n", meta.Time, meta.Text)
}
