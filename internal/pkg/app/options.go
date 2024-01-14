package app

type Optioners interface {
	Validate() []error
}
