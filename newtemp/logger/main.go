package main

import(
	"os"
	"log"
	"net/http"
	"go.uber.org/zap"

)
//var logger *zap.Logger
var logger *zap.SugaredLogger

func InitLogger() {
	//logger, _ = zap.NewProduction()

	baselogger, _ := zap.NewProduction()
	logger = baselogger.Sugar()
}

func simpleHttpGetByZapLogger(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("/Users/japhyzhang/workspace/go/go-sample-list/newtemp/logger/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}

func main() {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")

	InitLogger()
  	defer logger.Sync()
	simpleHttpGetByZapLogger("www.google.com")
	simpleHttpGetByZapLogger("http://www.google.com")


}