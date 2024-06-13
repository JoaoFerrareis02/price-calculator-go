package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {

	floats := make([]float64, len(strings))

	for stringIndex, stringValue := range strings {

		floatValue, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats[stringIndex] = floatValue

	}

	return floats, nil

}
