package modelBase

// Layer - Domain
// 回傳結構
type Response struct {
	ErrorCode    int    `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         any    `json:"Data"` // type any interface{}
}
