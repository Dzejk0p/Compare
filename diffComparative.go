package compare

import "reflect"

func (d *Differ) diffComparative(c *ComparativeList, name string) error {
	nv := reflect.ValueOf(nil)

	for _, k := range c.keys {
		if c.m[k].A == nil {
			c.m[k].A = &nv
		}

		if c.m[k].B == nil {
			c.m[k].B = &nv
		}

		err := d.diff(*c.m[k].A, *c.m[k].B, name, &name)
		if err != nil {
			return err
		}
	}

	return nil
}
