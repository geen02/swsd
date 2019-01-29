package swsdLib

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"
)

func TestSensorDataSender(t *testing.T) {
	var timeOut = time.Duration(10)
	var requestWG = sync.WaitGroup{}
	var fileName = "../../test/Minute_1.log"

	t.Run("run to test Function ReadSensorDataByLineFromFile", func(t *testing.T) {
		server := makeEchoServer()
		cString := make(chan string)

		go ReadSensorDataByLineFromFile(fileName, &cString)

		go func() {
			for rawData := range cString {
				requestWG.Add(1)
				defer requestWG.Done()

				structedData, _ := ParseDataByLine(rawData)
				jsonData, _ := GetJsonFromWeatherSensorData(structedData)
				statusCode, returnData, _ := SendWeatherSensorData("POST", server.URL, jsonData)
				fmt.Printf("Return: %v - %v\n", statusCode, returnData)
			}
		}()

		addLinesToFile(t, fileName, 10)

		fmt.Println("readFileWG.Wait()")
		requestWG.Wait()

		<-time.After(timeOut * time.Second)
		fmt.Printf("Timeout to run test %v", timeOut)
	})
}

func addLinesToFile(t *testing.T, filename string, lineNumber int) {
	t.Helper()

	go func() {
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			t.Errorf("Could not write content to file(%v)", err)
		}

		defer f.Close()

		for i := 0; i < lineNumber; i++ {
			lineContent := fmt.Sprintf("2018-01-04 16:26:00,25.7,17.8,0.2,212,10%d.%d,0.0\n", i, i)
			f.WriteString(lineContent)

			time.Sleep(1 * time.Second)
		}
	}()
}

func makeEchoServer() *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, buf)
	}))

	return srv
}
