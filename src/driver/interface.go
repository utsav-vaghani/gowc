package driver

type Driver interface {
	Run(args []string) error
}
