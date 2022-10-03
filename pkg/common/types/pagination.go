package types

type Pagination struct {
	Limit  int         `json:"limit"`
	Filter interface{} `json:"filter"`
	Sort   string      `json:"sort"`
	Offset int         `json:"skip"`
	Select string      `json:"select"`
}
