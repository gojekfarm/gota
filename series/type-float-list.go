package series

import (
	"fmt"
	"math"
	"strconv"
)

type floatListElement struct {
	e   []float64
	nan bool
}

// force floatListElement struct to implement Element interface
var _ Element = (*floatListElement)(nil)

func (e *floatListElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case []string:
		l := len(val)
		e.e = make([]float64, l)
		for i := 0; i < l; i++ {
			if val[i] == "NaN" {
				e.nan = true
				return
			}
			f, err := strconv.ParseFloat(val[i], 64)
			if err != nil {
				e.nan = true
				return
			}
			e.e[i] = f
		}
	case []int:
		l := len(val)
		e.e = make([]float64, l)
		for i := 0; i < l; i++ {
			e.e[i] = float64(val[i])
		}
	case []int32:
		l := len(val)
		e.e = make([]float64, l)
		for i := 0; i < l; i++ {
			e.e[i] = float64(val[i])
		}
	case []int64:
		l := len(val)
		e.e = make([]float64, l)
		for i := 0; i < l; i++ {
			e.e[i] = float64(val[i])
		}
	case []float32:
		l := len(val)
		e.e = make([]float64, l)
		for i := 0; i < l; i++ {
			e.e[i] = float64(val[i])
		}
	case []float64:
		l := len(val)
		e.e = make([]float64, l)
		for i := 0; i < l; i++ {
			e.e[i] = float64(val[i])
		}
	case []bool:
		l := len(val)
		e.e = make([]float64, l)
		for i := 0; i < l; i++ {
			b := val[i]
			if b {
				e.e[i] = 1
			} else {
				e.e[i] = 0
			}
		}
	case Element:
		e.e = val.FloatList()
	default:
		e.nan = true
		return
	}
}

func (e floatListElement) Copy() Element {
	if e.IsNA() {
		return &floatListElement{[]float64{}, true}
	}
	return &floatListElement{e.e, false}
}

func (e floatListElement) IsNA() bool {
	return e.nan
}

func (e floatListElement) Type() Type {
	return IntList
}

func (e floatListElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.e
}

func (e floatListElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e floatListElement) Int() (int, error) {
	return 0, fmt.Errorf("can't convert []float64 to int")
}

func (e floatListElement) Float() float64 {
	return 0
}

func (e floatListElement) Bool() (bool, error) {
	return false, fmt.Errorf("can't convert []float64 to bool")
}

func (e floatListElement) StringList() []string {
	if e.IsNA() {
		return []string{"NaN"}
	}

	l := make([]string, len(e.e))
	for i := 0; i < len(e.e); i++ {
		l[i] = fmt.Sprintf("%f", e.e[i])
	}
	return l
}

func (e floatListElement) IntList() ([]int, error) {
	if e.IsNA() {
		return nil, fmt.Errorf("can't convert NaN to []int")
	}

	l := make([]int, len(e.e))
	for i := 0; i < len(e.e); i++ {
		f := e.e[i]
		if math.IsInf(f, 1) || math.IsInf(f, -1) {
			return nil, fmt.Errorf("can't convert Inf to int")
		}
		if math.IsNaN(f) {
			return nil, fmt.Errorf("can't convert NaN to int")
		}
		l[i] = int(f)
	}
	return l, nil
}

func (e floatListElement) FloatList() []float64 {
	if e.IsNA() {
		return []float64{math.NaN()}
	}
	return e.e
}

func (e floatListElement) BoolList() ([]bool, error) {
	if e.IsNA() {
		return nil, fmt.Errorf("can't convert NaN to []bool")
	}

	l := make([]bool, len(e.e))
	for i := 0; i < len(e.e); i++ {
		if e.e[i] == 1 {
			l[i] = true
		} else {
			l[i] = false
		}
	}
	return l, nil
}

func (e floatListElement) Eq(elem Element) bool {
	list := elem.FloatList()

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] != list[i] {
			return false
		}
	}

	return true
}

func (e floatListElement) Neq(elem Element) bool {
	list := elem.FloatList()

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] == list[i] {
			return false
		}
	}

	return true
}

func (e floatListElement) Less(elem Element) bool {
	list := elem.FloatList()

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] < list[i] {
			return false
		}
	}

	return true
}

func (e floatListElement) LessEq(elem Element) bool {
	list := elem.FloatList()

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] <= list[i] {
			return false
		}
	}

	return true
}

func (e floatListElement) Greater(elem Element) bool {
	list := elem.FloatList()

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] > list[i] {
			return false
		}
	}

	return true
}

func (e floatListElement) GreaterEq(elem Element) bool {
	list := elem.FloatList()

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] >= list[i] {
			return false
		}
	}

	return true
}
