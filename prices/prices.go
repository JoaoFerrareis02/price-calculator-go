package prices

import (
	"fmt"

	"github.com/JoaoFerrareis02/price-calculator-go/conversion"
	"github.com/JoaoFerrareis02/price-calculator-go/filemanager"
)

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string][]float64
}

func (job *TaxIncludedPricesJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPricesJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, priceValue := range job.InputPrices {
		taxIncludedPrice := priceValue * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", priceValue)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPricesJob {

	return &TaxIncludedPricesJob{
		TaxRate: taxRate,
	}

}
