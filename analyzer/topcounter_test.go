package analyzer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTopCounter(t *testing.T) {
	counter := maxValues{}

	pushValues := []CountValue{
		{1, "5"},
		{2, "5"},
		{3, "5"},
		{4, "5"},
		{5, "5"},
		{4, "4"},
		{1, "1"},
		{3, "3"},
		{2, "2"},
	}

	expectedOrder := []CountValue{
		{5, "5"},
		{4, "4"},
		{3, "3"},
		{2, "2"},
		{1, "1"},
	}

	for _, toAdd := range pushValues {
		counter.addValue(toAdd, 5)
	}

	for i, expected := range expectedOrder {
		actual := counter[i]
		if !reflect.DeepEqual(expected, actual) {
			t.Fatal("Expected: ", expected, "Actual: ", actual)
		}
	}
}

func TestKeepMaxItems(t *testing.T) {
	counter := maxValues{}

	pushValues := []CountValue{
		{1, "5"},
		{2, "5"},
		{3, "5"},
		{4, "5"},
		{5, "5"},
		{4, "4"},
		{1, "1"},
		{3, "3"},
		{2, "2"},
	}

	expectedOrder := []CountValue{
		{5, "5"},
		{4, "4"},
		{3, "3"},
	}

	for _, toAdd := range pushValues {
		counter.addValue(toAdd, 3)
	}

	fmt.Println(counter)
	for i, expected := range expectedOrder {
		actual := counter[i]
		if !reflect.DeepEqual(expected, actual) {
			t.Fatal("Expected: ", expected, "Actual: ", actual)
		}
	}
}
