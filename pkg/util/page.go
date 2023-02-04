package util

type Page[T interface{}] struct {
	Num  int
	Size int
	Data []T
}

func NewPage[T interface{}](num, size int) *Page[T] {
	if num < 0 || size <= 0 {
		num = 0
		size = 5
	}

	return &Page[T]{
		Num:  num,
		Size: size,
	}
}
