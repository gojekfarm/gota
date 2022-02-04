package series

import (
	"fmt"
	"math"
	"strings"
)

type boolListElement struct {
	e   []bool
	nan bool
}

// force boolListElement struct to implement Element interface
var _ Element = (*boolListElement)(nil)

func (e *boolListElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case []string:
		l := len(val)
		e.e = make([]bool, l)
		for i := 0; i < l; i++ {
			if val[i] == "NaN" {
				e.nan = true
				return
			}
			switch strings.ToLower(val[i]) {
			case "true", "t", "1":
				e.e[i] = true
			case "false", "f", "0":
				e.e[i] = false
			default:
				e.nan = true
				return
			}
		}
	case []int:
		l := len(val)
		e.e = make([]bool, l)
		for i := 0; i < l; i++ {
			switch val[i] {
			case 1:
				e.e[i] = true
			case 0:
				e.e[i] = false
			default:
				e.nan = true
				return
			}
		}
	case []int32:
		l := len(val)
		e.e = make([]bool, l)
		for i := 0; i < l; i++ {
			switch val[i] {
			case 1:
				e.e[i] = true
			case 0:
				e.e[i] = false
			default:
				e.nan = true
				return
			}
		}
	case []int64:
		l := len(val)
		e.e = make([]bool, l)
		for i := 0; i < l; i++ {
			switch val[i] {
			case 1:
				e.e[i] = true
			case 0:
				e.e[i] = false
			default:
				e.nan = true
				return
			}
		}
	case []float32:
		l := len(val)
		e.e = make([]bool, l)
		for i := 0; i < l; i++ {
			switch val[i] {
			case 1:
				e.e[i] = true
			case 0:
				e.e[i] = false
			default:
				e.nan = true
				return
			}
		}
	case []float64:
		l := len(val)
		e.e = make([]bool, l)
		for i := 0; i < l; i++ {
			switch val[i] {
			case 1:
				e.e[i] = true
			case 0:
				e.e[i] = false
			default:
				e.nan = true
				return
			}
		}
	case []bool:
		l := len(val)
		e.e = make([]bool, l)
		for i := 0; i < l; i++ {
			e.e[i] = val[i]
		}
	case Element:
		v, err := val.BoolList()
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

func (e boolListElement) Copy() Element {
	if e.IsNA() {
		return &boolListElement{[]bool{}, true}
	}
	return &boolListElement{e.e, false}
}

func (e boolListElement) IsNA() bool {
	return e.nan
}

func (e boolListElement) Type() Type {
	return IntList
}

func (e boolListElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.e
}

func (e boolListElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e boolListElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return 0, fmt.Errorf("can't convert []bool to int")
}

func (e boolListElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return 0
}

func (e boolListElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	return false, fmt.Errorf("can't convert []int to bool")
}

func (e boolListElement) StringList() []string {
	if e.IsNA() {
		return []string{"NaN"}
	}
	return []string{fmt.Sprint(e.e)}
}

func (e boolListElement) IntList() ([]int, error) {
	if e.IsNA() {
		return nil, fmt.Errorf("can't convert NaN to []int")
	}

	l := make([]int, len(e.e))
	for i := 0; i < len(e.e); i++ {
		if e.e[i] {
			l[i] = 1
		} else {
			l[i] = 0
		}
	}
	return l, nil
}

func (e boolListElement) FloatList() []float64 {
	if e.IsNA() {
		return []float64{math.NaN()}
	}
	l := make([]float64, len(e.e))
	for i := 0; i < len(e.e); i++ {
		if e.e[i] {
			l[i] = 1.0
		} else {
			l[i] = 0.0
		}
	}
	return l
}

func (e boolListElement) BoolList() ([]bool, error) {
	if e.IsNA() {
		return nil, fmt.Errorf("can't convert NaN to []bool")
	}
	return e.e, nil
}

func (e boolListElement) Eq(elem Element) bool {
	list, err := elem.BoolList()
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

func (e boolListElement) Neq(elem Element) bool {
	list, err := elem.BoolList()
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

func (e boolListElement) Less(elem Element) bool {
	list, err := elem.BoolList()
	if err != nil {
		return false
	}

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if !e.e[i] && list[i] {
			return false
		}
	}

	return true
}

func (e boolListElement) LessEq(elem Element) bool {
	list, err := elem.BoolList()
	if err != nil {
		return false
	}

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if !e.e[i] || list[i] {
			return false
		}
	}

	return true
}

func (e boolListElement) Greater(elem Element) bool {
	list, err := elem.BoolList()
	if err != nil {
		return false
	}

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] && !list[i] {
			return false
		}
	}

	return true
}

func (e boolListElement) GreaterEq(elem Element) bool {
	list, err := elem.BoolList()
	if err != nil {
		return false
	}

	if len(e.e) != len(list) {
		return false
	}

	for i := 0; i < len(e.e); i++ {
		if e.e[i] || !list[i] {
			return false
		}
	}

	return true
}
