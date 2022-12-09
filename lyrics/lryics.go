package lyrics

func GetLyrics(name string) (text string, err error) {
	var request Request

	err = request.FromName(name)

	if err != nil {
		return
	}

	text, err = request.GetLrc()

	return
}
