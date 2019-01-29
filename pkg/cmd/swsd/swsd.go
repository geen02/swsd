package swsd

import (
	"fmt"
	"sync"
	"time"

	"github.com/geen02/swsd/pkg/swsdLib"
)

var timeOut = time.Duration(10)
var requestWG = sync.WaitGroup{}

// RUN : start agent service
func RUN() {
	cfg := swsdLib.NewCreateConfig("./config.yaml")

	cString := make(chan string)

	go swsdLib.ReadSensorDataByLineFromFile(cfg.DataFile, &cString)

	go func() {
		for rawData := range cString {
			requestWG.Add(1)
			defer requestWG.Done()

			structedData, _ := swsdLib.ParseDataByLine(rawData)
			jsonData, _ := swsdLib.GetJsonFromWeatherSensorData(structedData)
			fmt.Printf("jsonData: %v\n", string(jsonData))
			// statusCode, returnData, _ := SendWeatherSensorData("POST", cfg.StoreServer, jsonData)
			// fmt.Printf("Return: %v - %v\n", statusCode, returnData)
		}
	}()

	fmt.Println("readFileWG.Wait()")
	requestWG.Wait()

	<-time.After(timeOut * time.Second)
	fmt.Printf("Timeout to run test %v", timeOut)
}
