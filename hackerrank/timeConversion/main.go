package main

import (
	"fmt"
	"strconv"
)

func main() {
	strTime := "12:01:00AM"
	resultado := timeConversion(strTime)

	fmt.Println(resultado)
}

func timeConversion(s string) string {
	amPm := s[len(s)-2:]
	horas := s[:2]
	horasCon, _ := strconv.Atoi(horas)

	if amPm == "PM" && horasCon != 12 {
		horasCon += 12
	} else if amPm == "AM" && horasCon == 12 {
		horasCon = 00
	}
	horaFormatada := fmt.Sprintf("%02d", horasCon) + s[2:8]

	return horaFormatada

}
