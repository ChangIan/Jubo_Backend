package model

// Layer - Domain
// 醫囑
type Orders struct {
	Id         int    `json:"Id"`
	Patient_Id int    `json:"Patient_Id"`
	Message    string `json:"Message"`
	CreateTime string `json:"CreateTime"`
	ModifyTime string `json:"ModifyTime"`
}
