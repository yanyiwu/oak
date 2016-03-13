package mem

const (
	BLOCK_CAP = 512
)

type Block [BLOCK_CAP]interface{}

type Deque struct {
	buffer []*Block //TODO
	size   int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (this *Deque) Size() int {
	return this.size
}

func (this *Deque) Get(idx int) interface{} {
	blockid := idx / BLOCK_CAP
	offset := idx % BLOCK_CAP
	return (*this.buffer[blockid])[offset]
}

func (this *Deque) PushBack(v interface{}) {
	id := this.size / BLOCK_CAP
	// id == len
	if id >= len(this.buffer) {
		this.buffer = append(this.buffer, &Block{})
	}
	off := this.size % BLOCK_CAP
	(*this.buffer[id])[off] = v
	this.size += 1
}

//TODO
//func (this *Deque) PushFront(v interface{}) {
//	//TODO
//}
