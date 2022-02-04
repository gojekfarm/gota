package series

import "testing"

func Test_stringListElement_Set(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected []string
	}{
		{
			[]string{"1", "2", "3"},
			[]string{"1", "2", "3"},
		},
		{
			[]int{4, 5, 6},
			[]string{"4", "5", "6"},
		},
		{
			[]float64{3.14, 6.28, 9.42, 18.84},
			[]string{"3.140000", "6.280000", "9.420000", "18.840000"},
		},
		{
			[]bool{false, true, false, true},
			[]string{"false", "true", "false", "true"},
		},
	}
	for testnum, test := range tests {
		e := &stringListElement{}
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
