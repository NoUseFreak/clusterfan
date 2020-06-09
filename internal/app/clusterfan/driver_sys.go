package clusterfan

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type SysDriver struct{}

func (SysDriver) IsCapable() bool {
	sysFile := "/sys/class/thermal/thermal_zone0/temp"
	return fileExists(sysFile)
}
func (SysDriver) GetFunc() MeasureFunc {
	sysFile := "/sys/class/thermal/thermal_zone0/temp"
	reg := regexp.MustCompile("[^0-9]+")

	return func() int {
		content, err := ioutil.ReadFile(sysFile)
		if err != nil {
			log.Fatal(err)
		}

		contentString := reg.ReplaceAllString(string(content), "")

		temp, err := strconv.ParseFloat(contentString, 64)
		if err != nil {
			log.Infof("Could not read: %s", err.Error())
			return 0
		}

		return int(math.Round(temp / 1000))
	}
}
