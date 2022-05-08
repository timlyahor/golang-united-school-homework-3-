package homework

import (
	"reflect"
	"testing"
)

func TestTask2Sort2(t *testing.T) {
	expectedData := []int64{8, 7, 6, 5, 4, 3, 2, 1}
	sourceData := []int64{1, 3, 2, 7, 8, 5, 4, 6}
	actualData := reverse(sourceData)

	if !reflect.DeepEqual(expectedData, actualData) {
		t.Errorf("Result was incorect, got: %d, want: %d.", actualData, expectedData)
	}
}
