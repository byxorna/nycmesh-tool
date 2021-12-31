package app

import (
	"fmt"
	"log"
)

func SetSectorFrequency(frequency string, devices ...[]string) error {
	log.Printf("setSectorFrequency called: freq:%s devs:%v\n", frequency, devices)
	return fmt.Errorf("oof... implement me to lookup devices and get frequencies and stuff")
}
