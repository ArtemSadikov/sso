package grpc

import (
	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/utils"
	"github.com/ArtemSadikov/cinematic.back_protos/generated/go/sso"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapUserFromModel(user *model.User) *sso.User {
	contacts := make([]*sso.UserContact, len(user.Contacts))
	for _, contact := range user.Contacts {
		contacts = append(contacts, &sso.UserContact{
			Id:    contact.Id.String(),
			Value: contact.Value,
			Type:  utils.Enum(contact.Type, sso.UserContactTypeEnum_value, sso.UserContactTypeEnum_USER_CONTACT_TYPE_ENUM_UNSPECIFIED),
		})
	}

	return &sso.User{
		Id: user.Id.String(),
		Profile: &sso.UserProfile{
			Login: user.Login,
		},
		Contacts: contacts,
	}
}

func MapTokenFromModel(token *model.Token) *sso.Token {
	return &sso.Token{
		AvailableFor: timestamppb.New(*token.AvailableFor),
		Value:        token.Value,
	}
}
