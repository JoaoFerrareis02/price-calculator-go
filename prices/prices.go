package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string][]float64
}

func (job TaxTaxIncludedPricesJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

}

func (job TaxIncludedPricesJob) Process() {

	result := make(map[string]float64)

	for _, priceValue := range job.InputPrices {
		result[fmt.Sprintf("%.2f", priceValue)] = priceValue * (1 + job.TaxRate)
	}

	fmt.Println(result)

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPricesJob {

	return &TaxIncludedPricesJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}

}
