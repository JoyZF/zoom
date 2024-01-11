package app

type CliOptions interface {
	Validate() []error
}
