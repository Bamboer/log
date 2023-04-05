package bamboer
// GetAttr get structs fields
// judge the struct fields whether nil.
func GetAttr(instance interface{}, fields string) (result interface{}, err error) {
	if len(fields) == 0 {
		result = instance
		return
	}

	typ := reflect.TypeOf(instance)
	val := reflect.ValueOf(instance)

	return get(typ, val, fields)
}
func get(typ reflect.Type, val reflect.Value, fields string) (interface{}, error) {
	NotFound := fmt.Errorf("not found field %s", fields)
	PointerNil := fmt.Errorf("pointer is nil field %s", fields)
	UnExported := fmt.Errorf("unexported field %s", fields)
	tmps := strings.Split(fields, ".")

	switch typ.Kind() {
	case reflect.Struct:
		if newTyp, flag := typ.FieldByName(tmps[0]); flag {
			if newTyp.IsExported() {
				return get(newTyp.Type, val.FieldByName(tmps[0]), strings.Join(tmps[1:], "."))
			}
			return nil, UnExported
		}
		return nil, NotFound

	case reflect.Pointer:
		if val.IsNil() {
			return nil, PointerNil
		}
		return get(typ.Elem(), val.Elem(), fields)
	}
	if len(fields) == 0 {
		return val.Interface(), nil
	}
	return nil, NotFound
}
