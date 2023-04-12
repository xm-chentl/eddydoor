package views

type DataGrid struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}
