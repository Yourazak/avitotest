package pvz

import (
	"avitotes/config"
	"avitotes/interal/dto/errorDto"
	"avitotes/interal/dto/payloadError"
	"avitotes/pkg/midware"
	"avitotes/pkg/req"
	"net/http"
)

type PvzHandlerDependency struct {
	*PvzService
	*config.Config
}

type PvzHandler struct {
	*PvzService
	*config.Config
}

func (pvzHandler *PvzHandler) CreatePVZ() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.PvzCreateRequest](&w, r)
		if err != nil {
			strError := "Ошибка возникла на этапе чтения из объекта Request. Тело ошибки"
			errorDto.ShowResponseError(&w, strError, err, http.StatusBadRequest)
			return
		}
		pvz, err := pvzHandler.PvzService.Register(body.Id, body.RegistrationDate, body.City)
		if err != nil {
			strError := "Ошибка возникла на этапе создания PVZ в базе данных.Тело ошибки"
			errorDto.ShowResponseError(&w, strError, err, http.StatusBadRequest)
			return
		}
		req.JsonResponse(&w, pvz)
	}
}

func NewPvzHandler(router *http.ServeMux, pvz *PvzHandlerDependency) *PvzHandler {
	pvzHandler := &PvzHandler{
		pvz.PvzService,
		pvz.Config,
	}

	router.Handle("POST /pvz", midware.CheckRoleByToken(pvzHandler.CreatePVZ(), "TOKEN_MODERATOR"))
	return pvzHandler
}
