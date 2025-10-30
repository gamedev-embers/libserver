package databuilder

import "fmt"

func Map[K comparable, V any](lst []V, key func(V) K) (map[K]V, error) {
	m := map[K]V{}
	for _, row := range lst {
		id := key(row)
		if _, exists := m[id]; exists {
			return nil, fmt.Errorf("Map: duplicate id: %v", id)
		}
		m[id] = row
	}
	return m, nil
}

func MapOrPanic[K comparable, V any](lst []V, key func(V) K) map[K]V {
	m, err := Map(lst, key)
	if err != nil {
		panic(err)
	}
	return m

}

func MapList[K comparable, V any](lst []V, key func(V) K) (map[K][]V, error) {
	m := map[K][]V{}
	for _, row := range lst {
		id := key(row)
		m[id] = append(m[id], row)
	}
	if len(m) <= 0 {
		return nil, fmt.Errorf("MapList: empty map")
	}
	return m, nil
}

func MapListOrPanic[K comparable, V any](lst []V, key func(V) K) map[K][]V {
	m, err := MapList(lst, key)
	if err != nil {
		panic(err)
	}
	return m
}
