package repositories

import (
	dto "auth-skm/src/app/dtos/layanan"
)

type LayananRepository interface {
	GetLayananByOpdId(reqDTO *dto.GetLayananReqDTO) ([]*dto.LayananRespDTO, error)
}
