package exprutils

import (
	"errors"
)

func FlattenDropTickets(dropRecruitTickets interface{}) (interface{}, error) {
	if dropRecruitTickets == nil {
		return nil, nil
	}

	ticketsArray, err := convertToSliceOfSliceString(dropRecruitTickets)
	if err != nil {
		return nil, err
	}

	elements := make([]interface{}, 0)
	for _, el := range ticketsArray {
		for _, s := range el {
			elements = append(elements, s)
		}
	}
	return elements, nil
}

func convertToSliceOfSliceString(input interface{}) ([][]string, error) {
	result := [][]string{}
	if slice, ok := input.([]interface{}); ok {
		result = make([][]string, len(slice))
		for i, v := range slice {
			if innerSlice, ok := v.([]interface{}); ok {
				result[i] = make([]string, len(innerSlice))
				for j, s := range innerSlice {
					if str, ok := s.(string); ok {
						result[i][j] = str
					} else {
						return nil, errors.New("element " + s.(string) + " is not a string")
					}
				}
			} else {
				return nil, errors.New("element " + v.(string) + " is not a slice")
			}
		}
	} else {
		return nil, errors.New("input is not a slice")
	}
	return result, nil
}
