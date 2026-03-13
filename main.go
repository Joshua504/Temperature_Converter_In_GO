package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader :=  bufio.NewReader(os.Stdin)

	fmt.Println("input a NUMBER for conversion: ")
	userInputNum, _ := reader.ReadString('\n')
	trimmedI := strings.TrimSpace(userInputNum)
	intNum, _ := strconv.ParseFloat(trimmedI, 64)

	fmt.Println("convert from:\n 1: Celsius → Fahrenheit \n 2: Fahrenheit → Celsius \n 3: Celsius → Kelvin \n 4: Fahrenheit → Kelvin")
	userChoice, _ := reader.ReadString('\n')
	trimmedII := strings.TrimSpace(userChoice)
	intChoice, _ := strconv.Atoi(trimmedII)

	switch intChoice {
	case 1:
		message := fmt.Sprintf("\nANSWER: %.2f°F", celsiusToFahrenheit(intNum))
		fmt.Println(message)
	case 2:
		message := fmt.Sprintf("\nANSWER: %.2f°C", fahrenheitToCelsius(intNum))
		fmt.Println(message)
	case 3:
		message := fmt.Sprintf("\nANSWER: %.2f°K", celsiusToKelvin(intNum))
		fmt.Println(message)
	case 4 :
		message := fmt.Sprintf("\nANSWER: %.2f°K", fahrenheitToKelvin(intNum))
		fmt.Println(message)
	default:
		fmt.Println("invalid choice")
	}
}

func celsiusToFahrenheit(num float64) float64{
	conv := (float64(num) * 1.8) + 32
	return conv
}

func fahrenheitToCelsius(num float64) float64{
	conv := (float64(num) - 32) * 5 / 9
	return  conv
}

func celsiusToKelvin(num float64) float64{
	conv := (float64(num) + 273.15)
	return conv
}
func fahrenheitToKelvin(num float64) float64{
	conv := (float64(num) + 459.67) * 5 / 9
	return conv
}

