package helper

func Zip(data ...[]string) [][]string {
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
