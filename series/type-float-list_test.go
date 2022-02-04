package series

import "testing"

func Test_floatListElement_Set(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected []float64
	}{
		{
			[]string{"3.14", "6.28", "9.42", "18.84"},
			[]float64{3.14, 6.28, 9.42, 18.84},
		},
		{
			[]int{4, 5, 6},
			[]float64{4.0, 5.0, 6.0},
		},
		{
			[]float64{3.14, 6.28, 9.42, 18.84},
			[]float64{3.14, 6.28, 9.42, 18.84},
		},
		{
			[]bool{false, true, false, true},
			[]float64{0, 1, 0, 1},
		},
	}
	for testnum, test := range tests {
		e := &floatListElement{}
		e.Set(test.value)

		received := e.e
		expected := test.expected
		for i := 0; i < len(expected); i++ {
			if expected[i] != received[i] {
				t.Errorf(
					"Test:%v\nExpected:\n%v\nReceived:\n%v",
					testnum, expected, received,
				)
			}
		}
	}
}
