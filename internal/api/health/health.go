package health

import "net/http"

type HealthAPI struct{}

func NewHealthAPI() *HealthAPI {
	return &HealthAPI{}
}

func (h *HealthAPI) Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
