package cutils_default

func defaultString(origin *string, d ...any) string {
	if origin != nil {
		return *origin
	}
	if len(d) > 0 {
		if s, ok := d[0].(string); ok {
			return s
		}
	}
	return ""
}

func defaultInt(rest ...any) int {
	if len(rest) == 0 {
		return 0
	}

	var value int
	isNil := false

	switch v := rest[0].(type) {
	case *int:
		if v != nil {
			value = *v
		} else {
			isNil = true
		}
	case *int8:
		if v != nil {
			value = int(*v)
		} else {
			isNil = true
		}
	case *int16:
		if v != nil {
			value = int(*v)
		} else {
			isNil = true
		}
	case *int32:
		if v != nil {
			value = int(*v)
		} else {
			isNil = true
		}
	case *int64:
		if v != nil {
			value = int(*v)
		} else {
			isNil = true
		}
	default:
		isNil = true
	}

	if len(rest) == 1 {
		if isNil {
			return 0
		}
		return value
	}

	if len(rest) >= 2 {
		if isNil {
			if defaultValue, ok := rest[1].(int); ok {
				return defaultValue
			}
		} else {
			return value
		}
	}

	return 0
}
