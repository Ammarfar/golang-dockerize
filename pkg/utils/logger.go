package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Ammarfar/mezink-golang-assignment/internal/domain"
)

func WriteToErrorLog(payload *domain.LoggerPayload) string {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	filePath := fmt.Sprintf("./.log/%d/%d", year, month)
	fileName := fmt.Sprintf("%s/%d_error_log.log", filePath, day)

	payload.Time = now.Format("2006-01-02 15:04:05")
	jsonmarshal, _ := json.Marshal(payload)

	// mkdir if folder doesn't exist
	os.MkdirAll(filePath, os.ModePerm)

	// open file
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	// write to file
	fmt.Fprintf(file, "%s,\n", jsonmarshal)

	return string(jsonmarshal)
}
