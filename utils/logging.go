package utils

import (
	"log"
	"os"
	"io"
)

// logの関数設定
// LoggingSettingsはmainの前に設定する。(config.goのinitに設定する)
func LoggingSettings(logFile string) {
	/*
	os.O_RDWR = 読み書き
	os.O_CREATE = fileがなければ作成
	os.O_APPEND = fileに追記する
	パーミッション = 0666
	 */
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	// logの書き込み先設定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// logのフォーマットを指定
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	// 標準出力先設定
	log.SetOutput(multiLogFile)
}