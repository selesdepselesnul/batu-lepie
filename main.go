package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Battery struct {
	capacity int
	status string
}

func readBattery() (*Battery, error) {
	const batteryPath = "/sys/class/power_supply/BAT1/"
	
	capacity, errCapacity := ioutil.ReadFile(batteryPath + "capacity")
	status, errStatus := ioutil.ReadFile(batteryPath + "status")
	
	if errCapacity != nil || errStatus != nil {
		if errCapacity != nil {
			return nil, errCapacity
		} else {
			return nil, errStatus
		}
	}

	capacityInt, errCapacityParsed := strconv.Atoi(strings.TrimSuffix(string(capacity), "\n"))

	if errCapacityParsed != nil {
		return nil, errCapacityParsed
	}
	
	return &Battery{capacity: capacityInt, status: string(status)}, nil
}

func main() {
	fmt.Println(readBattery())
}





