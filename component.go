package gosq

type ComponentResponse struct {
	Component struct {
		Key      string `json:"key"`
		Measures []struct {
			Metric string `json:"metric"`
			Value  string `json:"value"`
		} `json:"measures"`
		Name      string `json:"name"`
		Qualifier string `json:"qualifier"`
	} `json:"component"`
}
