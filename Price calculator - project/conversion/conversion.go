package conversion

import (
	"errors"
	"strconv"
)

func ConvertStringArrayToFloat(stringArray []string) ([]float64, error) {

	floatArray := make([]float64, len(stringArray))

	for index, line := range stringArray {
		floatValue, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}

		floatArray[index] = floatValue
	}

	return floatArray, nil
}
