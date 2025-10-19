package mailer

type Distributor interface {
	Send(msg Message)
	ListenForMail(client Mailer, isSandbox bool)
}
