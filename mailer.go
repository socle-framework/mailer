package mailer

type Mailer interface {
	InitServer() error
	Send(msg Message, isSandbox bool) error
}
