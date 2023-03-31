package email

import (
	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
)

func (s *server) Create(req *emailPb.CreateEmailRequest) {
	model := model.ModelFromPB(req)
	s.emailService.Send(model)
}
