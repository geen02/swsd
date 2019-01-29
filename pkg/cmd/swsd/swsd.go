package swsd

import (
	"fmt"
	"sync"
	"time"

	"github.com/geen02/swsd/pkg/swsdlib"
)

var timeOut = time.Duration(10)
var requestWG = sync.WaitGroup{}

// Run : start agent service
func Run() {
	cfg := swsdlib.NewCreateConfig("./config.yaml")

	cString := make(chan string)

	go swsdlib.ReadSensorDataByLineFromFile(cfg.DataFile, &cString)

	for rawData := range cString {
		go func(rawData string) {
			requestWG.Add(1)
			defer requestWG.Done()

			structedData, _ := swsdlib.ParseDataByLine(rawData)
			jsonData, _ := swsdlib.GetJsonFromWeatherSensorData(structedData)
			fmt.Printf("jsonData: %v\n", string(jsonData))
			// statusCode, returnData, _ := SendWeatherSensorData("POST", cfg.StoreServer, jsonData)
			// fmt.Printf("Return: %v - %v\n", statusCode, returnData)
		}(rawData)
	}

	fmt.Println("readFileWG.Wait()")
	requestWG.Wait()
}
