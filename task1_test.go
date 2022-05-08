package homework

import "testing"

func TestTask1Avg1(t *testing.T) {
	sourceData := [15]float32{1, 2, 3, 4, 5, 6}
	res := average(sourceData)

	if res != 1.4 {
		t.Errorf("Result was incorect, got: %f, want: %f.", res, 1.4)
	}
}
