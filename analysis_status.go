package gosq

type ActivityStatus struct {
	Pending    int `json:"pending"`
	Failing    int `json:"failing"`
	InProgress int `json:"inProgress"`
}
