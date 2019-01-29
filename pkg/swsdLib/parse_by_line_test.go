package swsdLib

import (
	"testing"
	"time"
)

func TestParseLineData(t *testing.T) {
	LineData := "2018-01-04 16:26:00,25.7,17.8,0.2,212,1018.5,0.0"

	got, _ := ParseDataByLine(LineData)

	date, _ := time.Parse(timeLayout, "2018-01-04 16:26:00")

	want := WeatherSensorData{
		Date:        date,
		Temperature: "25.7",
		Humitidy:    "17.8",
		DWind:       "0.2",
		Pressure:    "212",
		Rainfall:    "1018.5",
	}

	t.Logf("\ngot '%v',\n want '%v'", got, want)

	if got != want {
		t.Errorf("\ngot '%v',\n want '%v'", got, want)
	}
}
