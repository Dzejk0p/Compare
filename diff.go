package compare

import (
	"errors"
	"fmt"
	"reflect"
)

type Differ struct {
	Tagname   string
	Zmiany    Zmiany
	CzyZmiana bool
}

type Zmiany map[string]interface{}

func NewDiffer() *Differ {
	return &Differ{
		Tagname:   "diff",
		Zmiany:    map[string]interface{}{},
		CzyZmiana: false,
	}
}

func Diff(a, b interface{}) (Zmiany, bool, error) {
	d := NewDiffer()
	return d.Zmiany, d.CzyZmiana, d.diff(reflect.ValueOf(a), reflect.ValueOf(b), "", nil)
}

func (d *Differ) diff(a, b reflect.Value, upperName string, parentName *string) error {
	if invalid(a, b) {
		return fmt.Errorf("Rożne typy: %v - %v", a.Kind().String(), b.Kind().String())
	}

	for i := 0; i < a.NumField(); i++ {

		f := a.Type().Field(i)

		name, ok := f.Tag.Lookup(d.Tagname)
		if !ok {
			continue
		}

		if name == "" {
			name = upperName
		}

		af := a.Field(i)

		bf := b.FieldByName(f.Name)

		var err error

		switch {
		case are(af, bf, reflect.Struct, reflect.Invalid):
			err = d.diffStruct(af, bf, name, parentName)
		case are(af, bf, reflect.String, reflect.Invalid):
			err = d.diffString(af, bf, name, parentName)
		case are(af, bf, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Invalid):
			err = d.diffInt(af, bf, name, parentName)
		case are(af, bf, reflect.Slice, reflect.Invalid):
			err = d.diffSlice(af, bf, name, parentName)
		default:
			return errors.New("Nieobsługiwany typ: " + a.Kind().String())
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Differ) dodajPole(nazwa string, parentName *string, bylo, jest interface{}, zmiana bool) {
	if zmiana {
		d.CzyZmiana = true
	}

	p := Pole{
		Bylo: bylo,
		Jest: jest,
	}

	if parentName != nil {
		if d.Zmiany[*parentName] == nil {
			d.Zmiany[*parentName] = []map[string]Pole{}
		}

		for _, m := range d.Zmiany[*parentName].([]map[string]Pole) {
			_, ok := m[nazwa]
			if !ok {
				m[nazwa] = p
				return
			}
		}
		m := map[string]Pole{
			nazwa: p,
		}
		d.Zmiany[*parentName] = append(d.Zmiany[*parentName].([]map[string]Pole), m)
		//d.Zmiany[*parentName] = append(d.Zmiany[*parentName].([]Pole), p)
		return
	}
	d.Zmiany[nazwa] = p
}

func invalid(a, b reflect.Value) bool {
	if a.Kind() == b.Kind() {
		return false
	}

	if a.Kind() == reflect.Invalid {
		return false
	}
	if b.Kind() == reflect.Invalid {
		return false
	}

	return true
}

func are(a, b reflect.Value, kinds ...reflect.Kind) bool {
	var amatch, bmatch bool

	for _, k := range kinds {
		if a.Kind() == k {
			amatch = true
		}
		if b.Kind() == k {
			bmatch = true
		}
	}

	return amatch && bmatch
}

func AreType(a, b reflect.Value, types ...reflect.Type) bool {
	var amatch, bmatch bool

	for _, t := range types {
		if a.Kind() != reflect.Invalid {
			if a.Type() == t {
				amatch = true
			}
		}
		if b.Kind() != reflect.Invalid {
			if b.Type() == t {
				bmatch = true
			}
		}
	}

	return amatch && bmatch
}
