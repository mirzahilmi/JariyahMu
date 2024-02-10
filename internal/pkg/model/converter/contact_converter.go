package converter

import (
	"golang-clean-architecture/internal/pkg/model"
	"golang-clean-architecture/internal/pkg/model/entity"
)

func ContactToResponse(contact *entity.Contact) *model.ContactResponse {
	return &model.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}
}
