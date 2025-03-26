package result

type PageResult struct {
	Total   int           `json:"total"`
	Records []interface{} `json:"records"`
}
