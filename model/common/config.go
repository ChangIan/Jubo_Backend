package modelCommon

// Layer - Domain
// 設定檔
type Config struct {
	WebSiteUrl  string `yaml:"WebSiteUrl"`
	WebSitePort string `yaml:"WebSitePort"`
	JwtKey      string `yaml:"JwtKey"`
	DbUrl       string `yaml:"DbUrl"`
	DbPort      string `yaml:"DbPort"`
	DbName      string `yaml:"DbName"`
	DbUserName  string `yaml:"DbUserName"`
	DbPassword  string `yaml:"DbPassword"`
}
