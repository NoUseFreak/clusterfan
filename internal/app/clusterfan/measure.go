package clusterfan

import (
	log "github.com/sirupsen/logrus"
)

type MeasureFunc func() int

func FindmeasureFunc() MeasureFunc {
	drivers := []Driver{
		SysDriver{},
		OSXDriver{},
	}

	for _, driver := range drivers {
		if driver.IsCapable() {
			return driver.GetFunc()
		}
	}

	log.Fatal("Platform not implemented")
	return nil
}

type Driver interface {
	IsCapable() bool
	GetFunc() MeasureFunc
}
