package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const thermalFolder = "/sys/class/thermal/"

//GetTemperatures ...
func GetTemperatures() map[string]int64 {
	result := make(map[string]int64, 0)
	thermalZones := getThermalZones()
	for _, zone := range thermalZones {

		tempString := readLine(thermalFolder + zone + "/temp")
		typeString := readLine(thermalFolder + zone + "/type")

		converted, err := strconv.ParseInt(tempString, 10, 32)
		if err != nil {
			panic(err)
		}
		result[typeString] = converted / 1000
	}

	return result
}

func getThermalZones() []string {
	result := make([]string, 0)

	f, err := os.Open(thermalFolder)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fileInfo, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	for _, file := range fileInfo {
		filename := file.Name()
		if strings.HasPrefix(filename, "thermal_zone") {
			result = append(result, filename)
		}
	}

	return result
}

func readLine(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReaderSize(file, 4*1024)
	line, _, err := reader.ReadLine()

	stringResult := string(line)
	return stringResult
}
