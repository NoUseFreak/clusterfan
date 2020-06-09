package clusterfan

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func collector(masterURL, secret string) {
	f := FindmeasureFunc()
	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})

	temps := make([]int, 5)
	for {
		select {
		case <-ticker.C:
			temp := f()
			temps = append(temps[1:], temp)
			max := maxInSlice(temps)
			log.Infof("Temp %v with max %d", temps, max)
			publishResult(masterURL, max, secret)

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func publishResult(url string, temp int, secret string) {
	msg := fmt.Sprintf("%s %d", viper.GetString("nodeName"), strconv.Itoa(temp))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(msg)))
	req.Header.Set(secretHeader, secret)
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Errorf("publish failed: %s", err.Error())
	}
}
