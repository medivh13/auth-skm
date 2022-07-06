package layanan_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	dto "auth-skm/src/app/dtos/layanan"
	"auth-skm/src/domain/repositories"
	"context"
	"log"
)

type LayananUsecaseInterface interface {
	GetLayanan(ctx context.Context, data *dto.GetLayananReqDTO) ([]*dto.LayananRespDTO, error)
}

type layananUseCase struct {
	LayananRepo repositories.LayananRepository
}

func NewLayananUseCase(layananRepo repositories.LayananRepository) *layananUseCase {
	return &layananUseCase{
		LayananRepo: layananRepo,
	}
}

func (uc *layananUseCase) GetLayanan(ctx context.Context, data *dto.GetLayananReqDTO) ([]*dto.LayananRespDTO, error) {

	dataRegister, err := uc.LayananRepo.GetLayananByOpdId(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dataRegister, nil
}
