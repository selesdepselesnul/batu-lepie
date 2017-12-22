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

const batteryPath = "/sys/class/power_supply/BAT1/"

func trimNL(content []byte) string {
	return strings.TrimSuffix(string(content), "\n")
}

func readCapacity() (int, error) {
	capacity, err := ioutil.ReadFile(batteryPath + "capacity")

	if err != nil {
		return 0, err
	}

	parsedCapacity, errParsed := strconv.Atoi(trimNL(capacity))

	if errParsed != nil {
		return 0, errParsed
	}
	
	return parsedCapacity, nil
}

func readStatus() (*string, error) {
	status, err := ioutil.ReadFile(batteryPath + "status")

	if err != nil {
		return nil, err
	}

	statusStr := trimNL(status)
	return &statusStr, nil
}

func readBattery() (*Battery, error) {
	
	capacity, errCapacity := readCapacity()
	status, errStatus := readStatus()
	
	if errCapacity != nil || errStatus != nil {
		if errCapacity != nil {
			return nil, errCapacity
		} else {
			return nil, errStatus
		}
	}
	
	return &Battery{capacity: capacity, status: *status}, nil
}

func main() {
	fmt.Println(readBattery())
}
