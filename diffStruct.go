package diff

import (
	"reflect"
	"time"
)

func (d *Differ) diffStruct(a, b reflect.Value, name string, parentName *string) error {
	if AreType(a, b, reflect.TypeOf(time.Time{})) {
		return d.diffTime(a, b, name, parentName)
	}

	return d.diff(a, b, name, parentName)
}
