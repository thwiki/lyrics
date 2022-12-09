package lrc

type Document struct {
	Meta  Meta   `json:"meta"`
	Lines []Line `json:"lines"`
}

func NewDocument() Document {
	return Document{
		Meta:  Meta{},
		Lines: make([]Line, 0),
	}
}

func (lrc *Document) String() (result string) {
	result += lrc.Meta.String()

	for _, line := range lrc.Lines {
		result += line.String()
	}

	return
}
