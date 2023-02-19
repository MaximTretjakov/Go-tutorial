package transformation

import (
	"errors"

	"github.com/MaximTretjakov/Go-tutorial/test-services/downloader/extraction"
	"github.com/MaximTretjakov/Go-tutorial/test-services/downloader/helper"
)

var (
	ErrNoData = errors.New("scraping: no data found")
)

const ProductNameGroup = 3
const ProteinFatCarbohydrateGroup = 2

var nutritionFacts = []string{
	`(td.*active">)(.*\n.*a.*product.*".)(.*)(</a>)`,
	`(protein-value.*\D)(.*\d)(.*\D)`,
	`(fat-value.*\D)(.*\d)(.*\D)`,
	`(carbohydrate-value.*\D)(.*\d)(.*\D)`,
}

func Transformation(raw []byte) ([][]string, error) {
	if len(raw) == 0 {
		return nil, ErrNoData
	}
	pfc := make([][]string, 0)
	productName := extraction.Extraction(raw, nutritionFacts[0], ProductNameGroup)
	for i := 1; i < len(nutritionFacts); i++ {
		ProteinFatCarbohydrate := extraction.Extraction(raw, nutritionFacts[i], ProteinFatCarbohydrateGroup)
		pfc = append(pfc, ProteinFatCarbohydrate)
	}
	result := helper.Zip(productName, pfc[0], pfc[1], pfc[2])
	return result, nil
}
