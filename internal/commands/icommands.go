package commands

type ICommand interface {
	Decode() error
	Exec() error
	Command() string
}
