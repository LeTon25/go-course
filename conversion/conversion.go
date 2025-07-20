package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloat64s(strings []string) ([]float64, error) {
	var floatValues []float64
	for _, str := range strings {
		floatValue, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing string '%s': %v", str, err)
		}
		floatValues = append(floatValues, floatValue)
	}
	return floatValues, nil
}
