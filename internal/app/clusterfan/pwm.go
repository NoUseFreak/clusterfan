package clusterfan

import (
	log "github.com/sirupsen/logrus"

	"github.com/stianeikeland/go-rpio"
)

type pwm struct {
	pin  rpio.Pin
	freq int
}

func newPWM() pwm {
	err := rpio.Open()
	if err != nil {
		log.Error(err)
	}

	p := pwm{
		freq: 64000,
	}
	if err == nil {
		log.Info("Setting pin 19 for WPM fan control")
		p.pin = rpio.Pin(19)
		p.pin.Mode(rpio.Pwm)
		p.pin.Freq(p.freq)
		p.pin.DutyCycle(99, 100)
	}

	return p
}

func (p *pwm) Close() {
	if p.pin == 0 {
		return
	}
	rpio.Close()
}

func (p *pwm) SetSpeed(speed int) {
	if speed > 99 {
		speed = 99
	} else if speed < 0 {
		speed = 0
	}

	if p.pin == 0 {
		log.Warn("Did not change wpm, unsupported platform")
		return
	}
	p.pin.DutyCycle(uint32(speed), 100)
}
