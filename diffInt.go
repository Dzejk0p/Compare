package compare

import (
	"errors"
	"fmt"
	"reflect"
)

func (d *Differ) diffInt(a, b reflect.Value, name string, parentName *string) error {
	if a.Kind() == reflect.Invalid {
		d.dodajPole(name, parentName, nil, int(b.Int()), true)
		return nil
	}

	if b.Kind() == reflect.Invalid {
		d.dodajPole(name, parentName, int(a.Int()), nil, true)
		return nil
	}

	if a.Kind() != b.Kind() {
		return errors.New(fmt.Sprintf("Rożne typy: %v - %v", a.Kind().String(), b.Kind().String()))
	}

	if a.Int() == b.Int() {
		d.dodajPole(name, parentName, nil, int(a.Int()), false)
		return nil
	}

	d.dodajPole(name, parentName, int(a.Int()), int(b.Int()), true)

	return nil
}
