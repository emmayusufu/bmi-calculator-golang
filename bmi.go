package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/emmayusufu/bmi-calculator-golang/info"
)

func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.ReplaceAll(weightInput, "\n", "")
	heightInput = strings.ReplaceAll(heightInput, "\n", "")

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI: %.2f", bmi)
}
