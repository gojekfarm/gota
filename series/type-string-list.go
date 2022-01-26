package series

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type stringListElement struct {
	e   []string
	nan bool
}

// force stringListElement struct to implement Element interface
var _ Element = (*stringListElement)(nil)

func (e *stringListElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case []string:
		l := len(val)
		e.e = make([]string, l)
		for i := 0; i < l; i++ {
			e.e[i] = string(val[i])
			if e.e[i] == "NaN" {
				e.nan = true
				return
			}
		}
	case []int:
		l := len(val)
		e.e = make([]string, l)
		for i := 0; i < l; i++ {
			e.e[i] = strconv.Itoa(val[i])
		}
	case []float64:
		l := len(val)
		e.e = make([]string, l)
		for i := 0; i < l; i++ {
			e.e[i] = strconv.FormatFloat(val[i], 'f', 6, 64)
		}
	case []bool:
		l := len(val)
		e.e = make([]string, l)
		for i := 0; i < l; i++ {
			b := val[i]
			if b {
				e.e[i] = "true"
			} else {
				e.e[i] = "false"
			}
		}
	case Element:
		e.e = val.StringList()
	default:
		e.nan = true
		return
	}
}

func (e stringListElement) Copy() Element {
	if e.IsNA() {
		return &stringListElement{[]string{}, true}
	}
	return &stringListElement{e.e, false}
}

func (e stringListElement) IsNA() bool {
	return e.nan
}

func (e stringListElement) Type() Type {
	return IntList
}

func (e stringListElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.e
}

func (e stringListElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e stringListElement) Int() (int, error) {
	return 0, fmt.Errorf("can't convert []int to int")
}

func (e stringListElement) Float() float64 {
	return 0
}

func (e stringListElement) Bool() (bool, error) {
	return false, fmt.Errorf("can't convert []int to bool")
}

func (e stringListElement) StringList() []string {
	if e.IsNA() {
		return []string{"NaN"}
	}
	return e.e
}

func (e stringListElement) IntList() ([]int, error) {
	if e.IsNA() {
		return nil, fmt.Errorf("can't convert NaN to []int")
	}

	l := make([]int, len(e.e))
	for i := 0; i < len(e.e); i++ {
		val, err := strconv.Atoi(e.e[i])
		if err != nil {
			return nil, err
		}
		l[i] = val
	}
	return l, nil
}

func (e stringListElement) FloatList() []float64 {
	if e.IsNA() {
		return []float64{math.NaN()}
	}

	l := make([]float64, len(e.e))
	for i := 0; i < len(e.e); i++ {
		val, err := strconv.ParseFloat(e.e[i], 64)
		if err != nil {
			return []float64{math.NaN()}
		}
		l[i] = val
	}
	return l
}

func (e stringListElement) BoolList() ([]bool, error) {
	if e.IsNA() {
		return nil, fmt.Errorf("can't convert NaN to []bool")
	}

	l := make([]bool, len(e.e))
	for i := 0; i < len(e.e); i++ {
		switch strings.ToLower(e.e[i]) {
		case "true", "t", "1":
			l[i] = true
		case "false", "f", "0":
			l[i] = false
		}
	}
	return l, nil
}

func (e stringListElement) Eq(elem Element) bool {
	list := elem.StringList()

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

func (e stringListElement) Neq(elem Element) bool {
	list := elem.StringList()

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

func (e stringListElement) Less(elem Element) bool {
	list := elem.StringList()

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

func (e stringListElement) LessEq(elem Element) bool {
	list := elem.StringList()

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

func (e stringListElement) Greater(elem Element) bool {
	list := elem.StringList()

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

func (e stringListElement) GreaterEq(elem Element) bool {
	list := elem.StringList()

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
