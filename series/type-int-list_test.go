package series

import "testing"

func Test_intListElement_Set(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected []int
	}{
		{
			[]string{"1", "2", "3"},
			[]int{1, 2, 3},
		},
		{
			[]int{4, 5, 6},
			[]int{4, 5, 6},
		},
		{
			[]float64{3.14, 6.28, 9.42, 18.84},
			[]int{3, 6, 9, 18},
		},
		{
			[]bool{false, true, false, true},
			[]int{0, 1, 0, 1},
		},
	}
	for testnum, test := range tests {
		e := &intListElement{}
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
