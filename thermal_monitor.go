package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const thermalFolder = "/sys/class/thermal/"

//GetTemperatures ...
func GetTemperatures() map[string]int64 {
	result := make(map[string]int64, 0)
	thermalZones := getThermalZones()
	for _, zone := range thermalZones {
		tempContent, err := ioutil.ReadFile(thermalFolder + zone + "/temp")
		if err != nil {
			panic(err)
		}

		typeContent, err := ioutil.ReadFile(thermalFolder + zone + "/type")
		if err != nil {
			panic(err)
		}

		tempString := strings.TrimSuffix(string(tempContent), "\n")
		typeString := strings.TrimSuffix(string(typeContent), "\n")

		converted, err := strconv.ParseInt(tempString, 10, 64)
		if err != nil {
			panic(err)
		}
		result[typeString] = converted / 1000
	}

	return result
}

func getThermalZones() []string {
	result := make([]string, 0)
	files, err := ioutil.ReadDir(thermalFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range files {
		if strings.HasPrefix(r.Name(), "thermal_zone") {
			result = append(result, r.Name())
		}
	}

	return result
}
