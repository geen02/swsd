package swsdLib

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	rhttp "github.com/hashicorp/go-retryablehttp"
)

// return json data from WeatherSensorData struct
func GetJsonFromWeatherSensorData(wsd WeatherSensorData) ([]byte, error) {
	data, err := json.Marshal(wsd)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SendWeatherSensorData(method, url string, wsdJson []byte) (int, string, error) {
	bodyData := bytes.NewReader(wsdJson)

	client := rhttp.NewClient()
	client.HTTPClient = &http.Client{}
	client.RetryWaitMax = 10 * time.Second
	client.RetryMax = 5

	request, err := rhttp.NewRequest(method, url, bodyData)
	if err != nil {
		return 0, "", err
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	response, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	responseBodyData := buf.String()

	return response.StatusCode, responseBodyData, nil
}

func CallRequestToUploadData(url string, line string) {
	wsd, err := ParseDataByLine(line)
	if err != nil {
		log.Println(err)
	}

	data, err := GetJsonFromWeatherSensorData(wsd)
	if err != nil {
		log.Println(err)
	}

	_, _, err = SendWeatherSensorData("POST", url, data)
	if err != nil {
		log.Println(err)
	}
}
