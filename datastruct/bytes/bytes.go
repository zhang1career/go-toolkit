package bytes

type Bytes []byte


func Equal(a *[]byte, b *[]byte) (isEqual bool) {
	if len(*a) != len(*b) {
		return false
	}
	for i, aa := range *a {
		if aa != (*b)[i] {
			return false
		}
	}
	return true
}

func (a *Bytes) Equal(b *Bytes) (isEqual bool) {
	if len(*a) != len(*b) {
		return false
	}
	for i, aa := range *a {
		if aa != (*b)[i] {
			return false
		}
	}
	return true
}


func Contain(container *[]byte, target *[]byte) (isEqual bool) {
	if len(*container) < len(*target) {
		return false
	}
	for i, aa := range *target {
		if aa != (*container)[i] {
			return false
		}
	}
	return true
}

func (container *Bytes) Contain(target *Bytes) (isEqual bool) {
	if len(*container) < len(*target) {
		return false
	}
	for i, aa := range *target {
		if aa != (*container)[i] {
			return false
		}
	}
	return true
}

