package slice

func (slc *Slice[T]) Pop(index int) (item T) {
	item = slc.inner[index]
	slc.inner = append(slc.inner[:index], slc.inner[index+1:]...)
	return item
}
