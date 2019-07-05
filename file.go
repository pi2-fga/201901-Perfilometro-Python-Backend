package main

import (
    "fmt"
	"strings"
)

func formatSensorData(content string)  {
	// sensorsData := [10][100000]int
	numberOfRows := strings.Count(content, "\n")
	fmt.Println("numberOfRows")
	fmt.Printf("%v", numberOfRows)

	// // fmt.Printf("[%v] string has %d of characters of [a] ", strings., numberOfRows)


}
