package abstractions

type UnitOfWork interface {
	Commit() error
}
