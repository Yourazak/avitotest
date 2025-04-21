package payload

type ReceptionCreateRequest struct {
	Pvz string `json:"pvzId" validate:"required"`
}
