package commands

type ICommand interface {
	Decode() *Error
	Exec() (*Response, *Error)
	Command() string
}
