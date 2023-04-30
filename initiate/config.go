package initiate

import (
	"log"

	"github.com/spf13/viper" // 一個用於讀取配置檔案的套件
)

var Dsn string

func SetConfig() {
	viper.SetConfigName("config") // 設定配置檔案的名稱
	viper.SetConfigType("yaml")   // 設定配置檔案的類型
	viper.AddConfigPath(".")      // 設定配置檔案的路徑
	err := viper.ReadInConfig()   // 讀取配置檔案
	if err != nil {
		log.Fatal(err)
	}

	Dsn = viper.GetString("dsn") // 從配置檔案中獲取dsn
}
