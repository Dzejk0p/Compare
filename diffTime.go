package diff

import (
	"fmt"
	"reflect"
	"time"
)

func (d *Differ) diffTime(a, b reflect.Value, name string, parentName *string) error {
	if a.Kind() == reflect.Invalid {
		d.dodajPole(name, parentName, nil, b.Interface(), true)
		return nil
	}

	if b.Kind() == reflect.Invalid {
		d.dodajPole(name, parentName, a.Interface(), nil, true)
		return nil
	}

	if a.Kind() != b.Kind() {
		return fmt.Errorf("Rożne typy: %v - %v", a.Kind().String(), b.Kind().String())
	}

	atime := a.Interface().(time.Time).UnixNano()
	btime := b.Interface().(time.Time).UnixNano()

	if atime == btime {
		d.dodajPole(name, parentName, nil, a.Interface(), false)
		return nil
	}

	d.dodajPole(name, parentName, a.Interface(), b.Interface(), true)

	return nil
}
