package templates

import (
	"github.com/matcornic/hermes/v2"
)

type template struct {
	hermes hermes.Hermes
}

func NewTemplate(link string) *template {
	var h hermes.Hermes
	h = hermes.Hermes{
		Product: hermes.Product{
			Name: "VGM",
			Link: link,
			Logo: "https://avatars.githubusercontent.com/u/99216816?s=400&u=632542b5c30ddecf0e29b584e9d55c7de8421d21&v=4",
		},
	}
	return &template{
		hermes: h,
	}

}

func (t *template) Greeting(username, link string) (string, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Name: "VGM",
			Intros: []string{
				"Welcome to VGM! We're very excited to have you on board.",
			},
			Dictionary: []hermes.Entry{
				{Key: "Username", Value: username},
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Text: "Go",
						Link: link,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	return t.hermes.GenerateHTML(email)
}
