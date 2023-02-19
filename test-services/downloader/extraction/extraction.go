package extraction

import (
	"regexp"
	"strings"
)

func Extraction(data []byte, patthern string, groupId int) []string {
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
