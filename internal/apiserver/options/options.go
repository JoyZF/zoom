package options

// Options runs an iam api server.
type Options struct {
}

func NewOptions() *Options {
	return &Options{
		// TODO
	}
}

func (o *Options) Validate() []error {
	return nil
}
