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

func readBattery() *Battery {
	
	capacityChan := make(chan int)
	statusChan := make(chan string)

	go func() {
		capacity, _ := readCapacity()
        capacityChan <- capacity
    }()
	
    go func() {
		status, _ := readStatus()
        statusChan <- *status
    }()

	return &Battery{capacity: <-capacityChan, status: <-statusChan}
}

func main() {
	fmt.Println(readBattery())
}


