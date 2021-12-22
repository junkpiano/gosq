package gosq

type AlmSettings struct {
	AlmSettings []struct {
		Alm string `json:"alm"`
		Key string `json:"key"`
		URL string `json:"url"`
	} `json:"almSettings"`
}
