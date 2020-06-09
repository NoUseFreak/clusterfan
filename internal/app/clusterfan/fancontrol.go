package clusterfan

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const secretHeader = "X-ClusterFan"

var currentFanSpeed = 0
var speedMap map[int]int

func Run() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("speedMap", `{"32":100,"20":0}`)
	hostname, _ := os.Hostname()
	viper.SetDefault("nodeName", hostname)
	viper.SetEnvPrefix("clusterfan")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	log.SetFormatter(&log.JSONFormatter{})

	secret := viper.GetString("masterSecret")
	if viper.GetBool("ismaster") {

		if err := json.Unmarshal([]byte(viper.GetString("speedMap")), &speedMap); err != nil {
			log.Fatal(err)
		}

		master(secret)
	} else {
		masterUrl := viper.GetString("masterUrl")
		collector(masterUrl, secret)
	}
}
