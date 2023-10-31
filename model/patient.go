package model

// Layer - Domain
// 患者
type Patient struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	Phone      string `json:"Phone"`
	Email      string `json:"Email"`
	Gender     string `json:"Gender"`
	CreateTime string `json:"CreateTime"`
	ModifyTime string `json:"ModifyTime"`
}
