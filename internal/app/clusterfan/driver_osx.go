package clusterfan

import (
	"os/exec"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type OSXDriver struct{}

func (OSXDriver) IsCapable() bool {
	_, err := exec.LookPath("osx-cpu-temp")
	return err == nil
}
func (OSXDriver) GetFunc() MeasureFunc {
	reg := regexp.MustCompile("[^0-9.]+")

	return func() int {
		b, err := exec.Command("osx-cpu-temp", "-c").Output()
		if err != nil {
			log.Fatal(err)
		}

		contentString := reg.ReplaceAllString(string(b), "")
		temp, err := strconv.ParseFloat(contentString, 64)

		return int(temp)
	}
}
