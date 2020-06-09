package clusterfan

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func master(secret string) {
	app := fiber.New(&fiber.Settings{
		DisableStartupMessage: true,
	})

	store := newStore()
	pwm := newPWM()
	defer pwm.Close()

	app.Post("/", func(c *fiber.Ctx) {
		if c.Get(secretHeader) != secret {
			c.Status(403)
			c.Send()
			return
		}
		data := strings.Split(c.Body(), " ")
		temp, err := strconv.Atoi(data[1])
		if err != nil {
			c.Status(400)
			c.Send()
			return
		}
		c.Send()

		store.Add(data[0], temp)
		if speed, changed := fanCheck(store.Max()); changed {
			log.Infof("Fanspeed to %d%%", speed)
			pwm.SetSpeed(speed)
		}
	})

	app.Get("/metrics", func(c *fiber.Ctx) {
		c.Send(prometheus(store))
	})

	log.Infof("Listening on 0.0.0.0:%d", viper.GetInt("port"))
	log.Fatal(app.Listen(viper.GetInt("port")))
}

func fanCheck(temp int) (int, bool) {

	speed := 0
	for cTemp, fanSpeed := range speedMap {
		if temp >= cTemp && fanSpeed > speed {
			speed = fanSpeed
		}
	}

	changed := currentFanSpeed != speed
	if changed {
		currentFanSpeed = speed
	}

	return speed, changed
}
