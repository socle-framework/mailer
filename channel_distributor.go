package mailer

type ChannelDistributor struct {
	Jobs    chan Message
	Results chan Result
}

// Result contains information regarding the status of the sent email message
type Result struct {
	Success bool
	Error   error
	//StatusCode int
}

func (c *ChannelDistributor) Send(msg Message) {
	//c.Jobs <- msg
}

// ListenForMail listens to the mail channel and sends mail
// when it receives a payload. It runs continually in the background,
// and sends error/success messages back on the Results channel.
// Note that if api and api key are set, it will prefer using
// an api to send mail
func (c *ChannelDistributor) ListenForMail(client Mailer, isSandbox bool) {
	for {
		msg := <-c.Jobs
		err := client.Send(msg, isSandbox)
		if err != nil {
			c.Results <- Result{false, err}
		} else {
			c.Results <- Result{true, nil}
		}
	}
}
