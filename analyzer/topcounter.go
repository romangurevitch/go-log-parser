package analyzer

import (
	"sort"
)

type CountValue struct {
	count int
	value string
}
type maxValues []CountValue

// Simple array to keep the 'keepAtMost' max count values, the array is sorted by the count value.
func (m *maxValues) addValue(newValue CountValue, keepAtMost int) {
	// if the new value is already one of the max values
	for k, v := range *m {
		if newValue.value == v.value {
			(*m)[k] = newValue
			sortByCount(*m)
			return
		}
	}

	if len(*m) < keepAtMost {
		*m = append(*m, newValue)
		sortByCount(*m)
		return
	}

	if (*m)[keepAtMost-1].count < newValue.count {
		(*m)[keepAtMost-1] = newValue
		sortByCount(*m)
	}
}

func sortByCount(maxValues []CountValue) {
	sort.SliceStable(maxValues, func(i, j int) bool {
		return maxValues[i].count > maxValues[j].count
	})
	return
}

func getValues(maxValues []CountValue) []string {
	var values []string
	for _, v := range maxValues {
		values = append(values, v.value)
	}
	return values
}
