package utils

import (
	"fmt"
	"strconv"
)

func TransformStringToUInt64(s string) (uint64, error) {
	id, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("parse unit error: %v", err)
	}
	return id, nil
}

func TransformUInt64ToString(number uint64) string {
	return strconv.FormatUint(number, 10)
}
