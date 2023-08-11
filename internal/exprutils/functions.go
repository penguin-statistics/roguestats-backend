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
	for _, totemInterface := range totemArray.([]interface{}) {
		totem := totemInterface.(string)
		lastChar := totem[len(totem)-1:]
		var value int
		if lastChar == "1" || lastChar == "2" || lastChar == "3" {
			value = 8
		} else if lastChar == "4" || lastChar == "5" {
			value = 12
		} else if lastChar == "6" {
			value = 16
		}
		results = append(results, value)
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

func (e ExprFunction) MapRecruitTicketsToOperatorClasses(dropRecruitTickets interface{}) (interface{}, error) {
	if dropRecruitTickets == nil {
		return nil, nil
	}

	mapping := GetExprCommonData().GetRecruitTicketOperatorClassMap()
	classes := make([]interface{}, 0)
	slice, _ := dropRecruitTickets.([]interface{})
	for _, ticketsForOneDropBox := range slice {
		classSet := make(map[string]interface{})
		innerSlice, _ := ticketsForOneDropBox.([]interface{})
		for _, ticket := range innerSlice {
			classNames := mapping[ticket.(string)]
			for _, className := range classNames {
				classSet[className] = nil
			}
		}
		for class := range classSet {
			classes = append(classes, class)
		}
	}
	return classes, nil
}

func (e ExprFunction) MapIncidentTypeToName(incidentType interface{}) (interface{}, error) {
	if incidentType == nil {
		return nil, nil
	}
	mapping := GetExprCommonData().GetIncidentTypeNameMap()
	return mapping[incidentType.(string)], nil
}

func (e ExprFunction) MapRestChoicesToNames(restChoices interface{}) (interface{}, error) {
	if restChoices == nil {
		return nil, nil
	}
	mapping := GetExprCommonData().GetRestChoicesNameMap()
	results := make([]interface{}, 0)
	for _, choice := range restChoices.([]interface{}) {
		results = append(results, mapping[choice.(string)])
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
