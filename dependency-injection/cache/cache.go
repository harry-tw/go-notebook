package cache

type LRU struct {
	size int64
}

func NewLRU(size int64) (*LRU, error) {
	return &LRU{size: size}, nil
}
