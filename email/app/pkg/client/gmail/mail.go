package gmail

type gmailConfig struct {
	smtpAuthAddress   string
	smtpServerAddress string
	name              string
	fromAddress       string
	fromPassword      string
}

type gmailClient struct {
	name         string
	fromAddress  string
	fromPassword string
}

func NewPgConfig(
	smtpAuthAddress string,
	smtpServerAddress string,
	name string,
	fromAddress string,
	fromPassword string) *gmailConfig {
	return &gmailConfig{
		smtpAuthAddress:   smtpAuthAddress,
		smtpServerAddress: smtpServerAddress,
		name:              name,
		fromAddress:       fromAddress,
		fromPassword:      fromPassword,
	}
}

func NewSMTPClient(name string, address string, password string) *gmailClient {
	return &gmailClient{
		name:         name,
		fromAddress:  address,
		fromPassword: password,
	}
}
