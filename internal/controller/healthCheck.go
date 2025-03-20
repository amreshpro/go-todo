package controller

import (
	"net/http"
	"encoding/json"
)


type HealthCheckResponse struct{
	Status string `json:"status"`
	Message string `json:"message"`
}

func HealthCheck( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	res := HealthCheckResponse{
		Status: http.StatusText(200),
		Message: "Yes it works baby",

	}

	json.NewEncoder(w).Encode(res)

}