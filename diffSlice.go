package diff

import (
	"fmt"
	"reflect"
)

type Comparative struct {
	A, B *reflect.Value
}

type ComparativeList struct {
	m    map[interface{}]*Comparative
	keys []interface{}
}

func (d *Differ) diffSlice(a, b reflect.Value, name string, parentName *string) error {
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

	c := createComparativeList(a, b)

	return d.diffComparative(c, name)
}

func createComparativeList(a, b reflect.Value) *ComparativeList {
	c := &ComparativeList{
		m:    map[interface{}]*Comparative{},
		keys: []interface{}{},
	}

	for i := 0; i < a.Len(); i++ {
		av := a.Index(i)
		c.addA(i, &av)
	}

	for i := 0; i < b.Len(); i++ {
		bv := b.Index(i)
		c.addB(i, &bv)
	}

	return c
}

func (c *ComparativeList) addA(key interface{}, val *reflect.Value) {
	if _, ok := (*c).m[key]; !ok {
		(*c).m[key] = &Comparative{}
		(*c).keys = append((*c).keys, key)
	}
	(*c).m[key].A = val
}

func (c *ComparativeList) addB(key interface{}, val *reflect.Value) {
	if _, ok := (*c).m[key]; !ok {
		(*c).m[key] = &Comparative{}
		(*c).keys = append((*c).keys, key)
	}
	(*c).m[key].B = val
}
