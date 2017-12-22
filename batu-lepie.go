package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
	"os"
)

type Battery struct {
	capacity int
	status string
}

func (b *Battery) String() string {
	return fmt.Sprintf(
		"battery : %d%%, status : %s",
		b.capacity,
		b.status)
}

const powerSupplyPath = "/sys/class/power_supply/"

func readBatteryVendor() string {
	files, _ := ioutil.ReadDir(powerSupplyPath)
	for _, f := range files {
		name := f.Name()
		match, _ := regexp.MatchString("BAT.*", name)
		if match {
			return name
		}
    }

	return "";
}

var batteryPath string = "/sys/class/power_supply/" + readBatteryVendor() + "/"

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

	statusStr := strings.ToLower(trimNL(status))
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

	args := os.Args

	if len(args) == 2 {
		switch arg := args[1]; arg {
		case "--all":
			fmt.Println(readBattery())
		case "--capacity":
			capacity, _ := readCapacity()
			fmt.Println(capacity)
		case "--status":
			status, _ := readStatus()
			fmt.Println(*status)
		default:
			fmt.Println("argument doesnt valid !")
		} 
	} else {
		fmt.Println("please fill the argument !")
	}

}





