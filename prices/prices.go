package prices

import (
	"fmt"

	"github.com/JoaoFerrareis02/price-calculator-go/conversion"
	"github.com/JoaoFerrareis02/price-calculator-go/filemanager"
)

type TaxIncludedPricesJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPricesJob) LoadData() {

	lines, err := job.IOManager.ReadLines()

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

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)

}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPricesJob {

	return &TaxIncludedPricesJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}

}
