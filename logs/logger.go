package insuredlog // 改用不同於標準庫的package名稱

import (
	"insured/config"
	"log"
	"os"
)

var logger *log.Logger // 將logger設定為全域變數

func Init() {
	file, err := os.OpenFile("logs/insured.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 創建一個文件來寫入日誌
	if err != nil {
		log.Fatal(err) // 如果創建文件失敗，終止程式
	}
	logger = log.New(file, "", log.LstdFlags) // 使用log.New來創建一個logger
	logger.Println("Hello, world!")           // 使用logger.Println來印出日誌

	_, err = config.Load()
	if err != nil {
		logger.Fatalf("Error loading config: %s\n", err) // 如果載入config失敗，終止程式
	}
}
