package lib

import (
	"strconv"
	"strings"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

// 기온,습도,풍속,풍향,기압,강우량
type WeatherSensorData struct {
	// "2016-01-02 15:04:05"
	Date        time.Time `json:"Date"`
	Temperature string    `json:"Temperature"`
	Humitidy    string    `json:"Humitidy"`
	DWind       string    `json:"DWind"`
	Pressure    string    `json:"Pressure"`
	Rainfall    string    `json:"Rainfall"`
}

func ParseStringToFloat(str string) (float64, error) {
	fData, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}

	return fData, nil
}

func ParseStringToTime(t string) (time.Time, error) {
	date, err := time.Parse(timeLayout, t)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

// return WeatherSensorData struct from content of file by line
func ParseDataByLine(line string) (WeatherSensorData, error) {
	result := strings.Split(line, ",")

	date, _ := ParseStringToTime(result[0])
	temperature := result[1]
	humitidy := result[2]
	dwind := result[3]
	pressure := result[4]
	rainfall := result[5]

	wsd := WeatherSensorData{
		Date:        date,
		Temperature: temperature,
		Humitidy:    humitidy,
		DWind:       dwind,
		Pressure:    pressure,
		Rainfall:    rainfall,
	}

	return wsd, nil
}
