package engine

type Comparable interface {
	Less()
}

type SkipList struct {
	length int
}

func NewSkipList() *SkipList {
	return &SkipList{}
}

func (this *SkipList) Len() int {
	return this.length
}
