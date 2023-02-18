package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
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

func scraping(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func extraction(data []byte, patthern string, groupId int) []string {
	re := regexp.MustCompile(patthern)
	matches := re.FindAllStringSubmatch(string(data), -1)
	result := make([]string, 0)
	for i := 0; i < len(matches); i++ {
		dataRow := strings.Join(matches[i], "")
		data := re.FindStringSubmatch(dataRow)
		result = append(result, data[groupId])
	}
	return result
}

func zip(data ...[]string) [][]string {
	if len(data) == 0 {
		return nil
	}
	minLen := len(data[0])
	for i := range data {
		if minLen > len(data[i]) {
			minLen = len(data[i])
		}
	}
	result := make([][]string, 0)
	for j := 0; j < minLen; j++ {
		tmp := make([]string, 0)
		for k := 0; k < len(data); k++ {
			tmp = append(tmp, data[k][j])
		}
		result = append(result, tmp)
	}
	return result
}

func transformation(raw []byte) ([][]string, error) {
	if len(raw) == 0 {
		return nil, ErrNoData
	}
	pfc := make([][]string, 0)
	productName := extraction(raw, nutritionFacts[0], ProductNameGroup)
	for i := 1; i < len(nutritionFacts); i++ {
		ProteinFatCarbohydrate := extraction(raw, nutritionFacts[i], ProteinFatCarbohydrateGroup)
		pfc = append(pfc, ProteinFatCarbohydrate)
	}
	result := zip(productName, pfc[0], pfc[1], pfc[2])
	return result, nil
}

func main() {
	rawData, err := scraping("https://calorizator.ru/product/mushroom")
	if err != nil {
		log.Fatal(err)
	}
	data, err := transformation(rawData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
