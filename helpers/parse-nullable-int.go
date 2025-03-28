package helpers

import "strconv"

func ParseNullableInt(value string) (*int, error) {
	if value == "" {
		return nil, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}
