package templates

import (
	"github.com/matcornic/hermes/v2"
)

type welcome struct {
	hermes hermes.Hermes
}

func (w *welcome) NewTemplate(link string) {
	w.hermes = hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			Name: "VGM",
			Link: link,
			// Optional product logo
			Logo: "https://avatars.githubusercontent.com/u/99216816?s=400&u=632542b5c30ddecf0e29b584e9d55c7de8421d21&v=4",
		},
	}
}

func Email(username, link string) hermes.Email {
	return hermes.Email{
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
}
