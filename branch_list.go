package gosq

type BranchList struct {
	Branches []struct {
		AnalysisDate      string `json:"analysisDate"`
		ExcludedFromPurge bool   `json:"excludedFromPurge"`
		IsMain            bool   `json:"isMain"`
		Name              string `json:"name"`
		Status            struct {
			QualityGateStatus string `json:"qualityGateStatus"`
		} `json:"status"`
		Type string `json:"type"`
	} `json:"branches"`
}
