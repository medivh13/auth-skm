package users_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	"net/http"

	dto "auth-skm/src/app/dtos/layanan"

	usecases "auth-skm/src/app/use_cases/layanan"

	common_error "auth-skm/src/infra/errors"
	"auth-skm/src/interface/rest/response"
	_ "net/http/pprof"
)

type LayananHandlerInterface interface {
	GetLayanan(w http.ResponseWriter, r *http.Request)
}

type layananHandler struct {
	response response.IResponseClient
	usecase  usecases.LayananUsecaseInterface
}

func NewLayananHandler(r response.IResponseClient, u usecases.LayananUsecaseInterface) LayananHandlerInterface {
	return &layananHandler{
		response: r,
		usecase:  u,
	}
}

func (h *layananHandler) GetLayanan(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetLayananReqDTO{}
	dataOpdId := r.URL.Query().Get("opd_id")

	getDTO.OpdID = dataOpdId

	getDTO.Token = r.Header.Get("token")

	err := getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.GetLayanan(r.Context(), &getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Register",
		data,
		nil,
	)
}
