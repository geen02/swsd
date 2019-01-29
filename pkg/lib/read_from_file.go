package lib

import (
	"log"

	"github.com/hpcloud/tail"
)

// Read data from file line by line and return as []string
func ReadSensorDataByLineFromFile(filename string, cString *chan string) {
	t, err := tail.TailFile(filename, tail.Config{Follow: true})
	if err != nil {
		log.Fatalf("Could not open file %s :%v", filename, err)
	}

	for {
		for line := range t.Lines {
			*cString <- line.Text

			// outputData, err := ParseDataByLine(data)
			if err != nil {
				log.Printf("Error during reading file :%v", err)
			}
		}
	}
}
