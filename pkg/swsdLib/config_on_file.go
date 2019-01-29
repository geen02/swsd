package swsdLib

import (
	"fmt"

	"github.com/spf13/viper"
)

// SWSDConfig - store configuration info
// ConfigFile: path of sensor data file
// StoreServer: url of server to send to upload sensor data
type SWSDConfig struct {
	DataFile    string
	StoreServer string
}

// NewCreateConfig - return SWSDConfig reading from config.yaml
func NewCreateConfig(configFile string) (sc *SWSDConfig) {
	fmt.Println(configFile)

	// var runtime_viper = viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	sc.DataFile = viper.Get("datafile_path").(string)
	sc.StoreServer = viper.Get("store_server_url").(string)

	return
}
