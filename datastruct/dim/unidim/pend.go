package unidim

func Prepend(slice []interface{}, elems ...interface{}) []interface{} {
	if len(elems) <= 0 {
		return slice
	}
	
	ret := slice
	for i := len(elems)-1; i >= 0; i-- {
		ret = append([]interface{}{elems[i]}, ret...)
	}
	return ret
}
