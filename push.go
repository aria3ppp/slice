package slice

func (slc *Slice[T]) Push(item T, index int) {
	slc.inner = append(slc.inner[:index+1], slc.inner[index:]...)
	slc.inner[index] = item
}
