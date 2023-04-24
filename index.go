package slice

func (slc Slice[T]) Index(item T) (index int) {
	var v T
	for index, v = range slc.inner {
		if item == v {
			return index
		}
	}
	return -1
}
