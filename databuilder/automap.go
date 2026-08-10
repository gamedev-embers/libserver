package databuilder

import (
	"fmt"
	"reflect"
)

var mapPools = map[reflect.Type]any{}

func AutoMap[T any](rows []T) map[uint32]T {

	makeKeyUInt32 := func(v reflect.Value) (uint32, error) {
		if !v.IsValid() {
			return 0, fmt.Errorf("invalid value")
		}
		switch v.Kind() {
		case reflect.Uint32:
			return uint32(v.Uint()), nil
		case reflect.Uint:
			return uint32(v.Uint()), nil
		case reflect.Int:
			return uint32(v.Int()), nil
		case reflect.Int32:
			return uint32(v.Int()), nil
		default:
			return 0, fmt.Errorf("unsupported kind: %s", v.Kind())
		}
	}

	// 使用反射，根据 T 的类型获取对应的 map
	typ := reflect.TypeOf((*T)(nil)).Elem()
	dataMap := mapPools[typ]
	if dataMap == nil {
		m := make(map[uint32]T)
		for _, row := range rows {
			// 使用反射获取 row 的 Id 字段或者 ID 字段
			v := reflect.ValueOf(row)
			var id uint32
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				for _, fieldName := range []string{"Id", "ID"} {
					idField := v.FieldByName(fieldName)
					_id, err := makeKeyUInt32(idField)
					if err != nil {
						continue
					}
					id = _id
					break
				}
			} else {
				panic(fmt.Sprintf("Type %s is not a struct or pointer to struct", typ.Name()))
			}
			m[id] = row
		}
		dataMap = m
		mapPools[typ] = dataMap
	}
	return dataMap.(map[uint32]T)
}
