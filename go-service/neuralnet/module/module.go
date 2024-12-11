package module

type Module interface {
	Forward(any) (any, error)
}
