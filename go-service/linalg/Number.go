package linalg

type Int interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Int | Float
}
