package entry

import (
	"errors"

	"github.com/MaximTretjakov/Go-tutorial/test-services/downloader/scraping"
	"github.com/MaximTretjakov/Go-tutorial/test-services/downloader/transformation"
)

func Run(url string) ([][]string, error) {
	if len(url) == 0 {
		return nil, errors.New("target url not passed")
	}
	rawData, err := scraping.Scraping(url)
	if err != nil {
		return nil, err
	}
	data, err := transformation.Transformation(rawData)
	if err != nil {
		return nil, err
	}
	return data, nil
}
