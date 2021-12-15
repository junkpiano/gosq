package gosq

type Projects struct {
	Components []struct {
		Key              string `json:"key"`
		LastAnalysisDate string `json:"lastAnalysisDate"`
		Name             string `json:"name"`
		Qualifier        string `json:"qualifier"`
		Revision         string `json:"revision"`
		Visibility       string `json:"visibility"`
	} `json:"components"`
	Paging struct {
		PageIndex int64 `json:"pageIndex"`
		PageSize  int64 `json:"pageSize"`
		Total     int64 `json:"total"`
	} `json:"paging"`
}
