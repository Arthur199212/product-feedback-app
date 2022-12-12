package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseFeedbackIdsFromQuery(str string) ([]int, error) {
	if str == "" {
		return []int{}, nil
	}

	ids := strings.Split(str, ",")
	parsedIds := make([]int, len(ids))
	for i := range ids {
		parsedId, err := strconv.Atoi(ids[i])
		if err != nil {
			return parsedIds, fmt.Errorf("invalid feedbackId param")
		}
		parsedIds[i] = parsedId
	}
	return parsedIds, nil
}
