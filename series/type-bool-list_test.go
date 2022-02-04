package series

import "testing"

func Test_boolListElement_Set(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected []bool
	}{
		{
			[]string{"true", "t", "1", "false", "f", "0"},
			[]bool{true, true, true, false, false, false},
		},
		{
			[]int{0, 1},
			[]bool{false, true},
		},
		{
			[]float64{0.0, 1.0},
			[]bool{false, true},
		},
		{
			[]bool{false, true, false, true},
			[]bool{false, true, false, true},
		},
	}
	for testnum, test := range tests {
		e := &boolListElement{}
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
