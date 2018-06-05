package server

import funk "github.com/thoas/go-funk"

func DiffSlice(old, new_ interface{}) (added []interface{}, deleted []interface{}) {
	funk.ForEach(old, func(e interface{}) {
		if !funk.Contains(new_, e) {
			deleted = append(deleted, e)
		}
	})
	funk.ForEach(new_, func(e interface{}) {
		if !funk.Contains(old, e) {
			added = append(added, e)
		}
	})
	return
}

func DiffSliceAsString(old, new_ interface{}) (added []string, deleted []string) {
	addedx, deletedx := DiffSlice(old, new_)
	for _, e := range addedx {
		added = append(added, e.(string))
	}
	for _, e := range deletedx {
		deleted = append(deleted, e.(string))
	}
	return
}
