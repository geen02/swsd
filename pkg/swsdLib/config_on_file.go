package swsdlib

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// SWSDConfig - store configuration info
// ConfigFile: path of sensor data file
// StoreServer: url of server to send to upload sensor data
type SWSDConfig struct {
	DataFile    string
	StoreServer string
}

func getPWD() string {
	pwd, _ := os.Getwd()
	return pwd
}

// CheckConfigFile - check whether config file exists
func CheckConfigFile(configFile string) {
	fmt.Println(configFile)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Panicf("Configuration File does not exist [%s]", configFile)
	}
}

// NewCreateConfig - return SWSDConfig reading from config.yaml
func NewCreateConfig(configFile string) (sc *SWSDConfig) {
	// check config file
	fmt.Println(configFile)

	flag.String("configfile", "", "Configuration file path\n Default : ./config.yaml")
	flag.String("datafile", "", "Data file of sensor data saved")
	flag.String("url", "", "Server URL to upload sensor data")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// Check flag for config file
	if viper.GetString("configfile") != "" {
		configFile = path.Join(getPWD(), viper.GetString("configfile"))
	}

	CheckConfigFile(configFile)

	if viper.GetString("datafile") == "" || viper.GetString("url") == "" {
		viper.SetConfigFile(configFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicf("Fatal error config file: %s", err)
		}
	}

	sc = &SWSDConfig{}

	if datafile := viper.GetString("datafile"); datafile != "" {
		sc.DataFile = viper.GetString("datafile")
	} else if viper.InConfig("datafile_path") {
		sc.DataFile = viper.Get("datafile_path").(string)
	}

	if url := viper.GetString("url"); url != "" {
		sc.StoreServer = viper.GetString("url")
	} else if viper.InConfig("store_server_url") {
		sc.StoreServer = viper.Get("store_server_url").(string)
	}

	return
}
