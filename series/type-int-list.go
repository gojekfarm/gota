package series

import (
	"fmt"
	"math"
	"strconv"
)

type intListElement struct {
	e   []int
	nan bool
}

// force intListElement struct to implement Element interface
var _ Element = (*intListElement)(nil)

func (e *intListElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case []string:
		l := len(val)
		e.e = make([]int, l)
		for i := 0; i < l; i++ {
			if val[i] == "NaN" {
				e.nan = true
				return
			}
			vi, err := strconv.Atoi(val[i])
			if err != nil {
				e.nan = true
				return
			}
			e.e[i] = vi
		}
	case []int:
		l := len(val)
		e.e = make([]int, l)
		for i := 0; i < l; i++ {
			e.e[i] = int(val[i])
		}
	case []float64:
		l := len(val)
		e.e = make([]int, l)
		for i := 0; i < l; i++ {
			f := val[i]
			if math.IsNaN(f) ||
				math.IsInf(f, 0) ||
				math.IsInf(f, 1) {
				e.nan = true
				return
			}
			e.e[i] = int(f)
		}
	case []bool:
		l := len(val)
		e.e = make([]int, l)
		for i := 0; i < l; i++ {
			b := val[i]
			if b {
				e.e[i] = 1
			} else {
				e.e[i] = 0
			}
		}
	case Element:
		v, err := val.IntList()
		if err != nil {
			e.nan = true
			return
		}
		e.e = v
	default:
		e.nan = true
		return
	}
}

func (e intListElement) Copy() Element {
	if e.IsNA() {
		return &intListElement{[]int{}, true}
	}
	return &intListElement{e.e, false}
}

func (e intListElement) IsNA() bool {
	return e.nan
}

func (e intListElement) Type() Type {
	return IntList
}

func (e intListElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.e
}

func (e intListElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e intListElement) Int() (int, error) {
	return 0, fmt.Errorf("can't convert []int to int")
}

func (e intListElement) Float() float64 {
	return 0
}

func (e intListElement) Bool() (bool, error) {
	return false, fmt.Errorf("can't convert []int to bool")
}

func (e intListElement) StringList() []string {
	if e.IsNA() {
		return []string{"NaN"}
	}

	l := make([]string, len(e.e))
	for i := 0; i < len(e.e); i++ {
		l[i] = fmt.Sprint(e.e[i])
	}
	return l
}

func (e intListElement) IntList() ([]int, error) {
	if e.IsNA() {
		return nil, fmt.Errorf("can't convert NaN to []int")
	}
	return e.e, nil
}

func (e intListElement) FloatList() []float64 {
	if e.IsNA() {
		return []float64{math.NaN()}
	}

	l := make([]float64, len(e.e))
	for i := 0; i < len(e.e); i++ {
		l[i] = float64(e.e[i])
	}
	return l
}

func (e intListElement) BoolList() ([]bool, error) {
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

func (e intListElement) Eq(elem Element) bool {
	list, err := elem.IntList()
	if err != nil {
		return false
	}

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

func (e intListElement) Neq(elem Element) bool {
	list, err := elem.IntList()
	if err != nil {
		return false
	}

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

func (e intListElement) Less(elem Element) bool {
	list, err := elem.IntList()
	if err != nil {
		return false
	}

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

func (e intListElement) LessEq(elem Element) bool {
	list, err := elem.IntList()
	if err != nil {
		return false
	}

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

func (e intListElement) Greater(elem Element) bool {
	list, err := elem.IntList()
	if err != nil {
		return false
	}

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

func (e intListElement) GreaterEq(elem Element) bool {
	list, err := elem.IntList()
	if err != nil {
		return false
	}

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
