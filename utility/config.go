package utility

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	modelCommon "changian.com/jubo/model/common"
)

// 取得設定檔
func GetConfig() *modelCommon.Config {
	// 讀取 config/config.yaml 文件
	yamlFile, err := os.ReadFile("config/config.yaml")

	// 若出現錯誤，印出錯誤訊息
	if err != nil {
		fmt.Println(err.Error())
	}

	conf := new(modelCommon.Config)

	// 將讀取的字符串轉換成結構體
	err = yaml.Unmarshal(yamlFile, conf)

	if err != nil {
		fmt.Println(err.Error())
	}

	return conf
}
