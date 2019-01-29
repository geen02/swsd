package swsd

import (
	"fmt"
	"sync"
	"time"

	lib "github.com/geen02/swsd/tree/master/pkg/lib"
)

var timeOut = time.Duration(10)
var requestWG = sync.WaitGroup{}
var fileName = "./Minute_1.log"

// CreateNewSendAgent : return agent struct
func CreateNewSendAgent() {
	lib.NewCreateConfig("./config.yaml")
	cString := make(chan string)

	go ReadSensorDataByLineFromFile(fileName, &cString)

	go func() {
		for rawData := range cString {
			requestWG.Add(1)
			defer requestWG.Done()

			structedData, _ := ParseDataByLine(rawData)
			jsonData, _ := GetJsonFromWeatherSensorData(structedData)
			fmt.Printf("jsonData: %v\n", string(jsonData))
			// statusCode, returnData, _ := SendWeatherSensorData("POST", server.URL, jsonData)
			// fmt.Printf("Return: %v - %v\n", statusCode, returnData)

		}
	}()

	fmt.Println("readFileWG.Wait()")
	requestWG.Wait()

	<-time.After(timeOut * time.Second)
	fmt.Printf("Timeout to run test %v", timeOut)
}
