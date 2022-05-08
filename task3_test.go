package homework

import (
	"reflect"
	"testing"
)

func TestTask3(t *testing.T) {
	expectedData := []string{"b", "a", "c"}
	sourceData := map[int]string{
		2: "a",
		1: "b",
		3: "c",
	}
	actualData := sortMapValues(sourceData)

	if !reflect.DeepEqual(expectedData, actualData) {
		t.Errorf("Result was incorect, got: %s, want: %s.", actualData, expectedData)
	}
}
