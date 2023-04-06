package model

type Email struct {
	Subject string
	Content string
	To      []string
	Cc      []string
	Bcc     []string
	Files   []string
}
