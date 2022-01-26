package interfaces

type Entity interface {
	GetId() string
	IsValid() error
}
