package reception

import (
	"avitotes/config"
	"avitotes/interal/dto/errorDto"
	payload "avitotes/interal/dto/payloadError"
	"avitotes/pkg/midware"
	"avitotes/pkg/req"
	"net/http"
)

type ReceptionHandlerDependency struct {
	*ReceptionService
	*config.Config
}

type ReceptionHandler struct {
	*ReceptionService
	*config.Config
}

func NewReceptionHandler(router *http.ServeMux, recep *ReceptionHandlerDependency) *ReceptionHandler {
	receptionHandler := &ReceptionHandler{
		recep.ReceptionService,
		recep.Config,
	}
	router.Handle("POST /receptions", midware.CheckRoleByToken(recepHandler.CreateReceptHandler(), "client"))
	return recepHandler
}

func (receptionHandler *ReceptionHandler) CreateReceptHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.ReceptionCreateRequest](&w, r)
		if err != nil {
			strError := "Ошибка возникла на этапе чтения из объекта Request. Тело ошибки"
			errorDto.ShowResponseError(&w, strError, err, http.StatusBadRequest)
			return
		}
		reception, err := receptionHandler.ReceptionService.CreateReception(body.PvzId)
		if err != nil {
			strError := "Ошибка возникла на этапе создания reception в базе данных. Тело ошибки"
			errorDto.ShowResponseError(&w, strError, err, http.StatusBadRequest)
			return
		}
		req.JsonResponse(&w, reception)
	}
}
