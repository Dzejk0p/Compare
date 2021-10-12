package compare

import (
	"fmt"
	"reflect"
)

func (d *Differ) diffString(a, b reflect.Value, name string, parentName *string) error {
	if a.Kind() == reflect.Invalid {
		d.dodajPole(name, parentName, nil, b.String(), true)
		return nil
	}

	if b.Kind() == reflect.Invalid {
		d.dodajPole(name, parentName, a.String(), nil, true)
		return nil
	}

	if a.Kind() != b.Kind() {
		return fmt.Errorf("Rożne typy: %v - %v", a.Kind().String(), b.Kind().String())
	}

	if a.String() == b.String() {
		d.dodajPole(name, parentName, nil, a.String(), false)
		return nil
	}

	d.dodajPole(name, parentName, a.String(), b.String(), true)

	return nil
}
