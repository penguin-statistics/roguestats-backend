package exprutils

import (
	"errors"
)

type ExprFunction struct {
}

// Public methods here will be available in the expression environment. Return type must be interface{} and error.

func (e ExprFunction) FlattenDropTickets(dropRecruitTickets interface{}) (interface{}, error) {
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

func (e ExprFunction) MapTotemArrayToValues(totemArray interface{}) (interface{}, error) {
	if totemArray == nil {
		return nil, nil
	}

	results := make([]interface{}, 0)
	totemValueMap := GetExprCommonData().GetTotemValueMap()
	for _, totemInterface := range totemArray.([]interface{}) {
		totem := totemInterface.(string)
		if _, ok := totemValueMap[totem]; !ok {
			return nil, errors.New("invalid totem " + totem)
		}
		results = append(results, totemValueMap[totem])
	}
	return results, nil
}

func (e ExprFunction) MapTotemArrayToColors(totemArray interface{}) (interface{}, error) {
	if totemArray == nil {
		return nil, nil
	}

	results := make([]interface{}, 0)
	for _, totemInterface := range totemArray.([]interface{}) {
		totem := totemInterface.(string)
		results = append(results, totem[len(totem)-4:len(totem)-3])
	}
	return results, nil
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
