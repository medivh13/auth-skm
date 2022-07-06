package layanan_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type LayananDTOInterface interface {
	Validate() error
}

type GetLayananReqDTO struct {
	Token string `json:"token"`
	OpdID string `json:"opd_id"`
}

func (dto *GetLayananReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Token, validation.Required),
		validation.Field(&dto.OpdID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type LayananIntegrateDTO struct {
	Data []*LayananRespDTO `json:"data"`
}

type LayananRespDTO struct {
	ID             int64  `json:"id"`
	DisplayName    string `json:"display_name"`
	UnsurPelayanan string `json:"unsur_pelayanan"`
}
