package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/entity"
	"github.com/gorilla/context"
	"github.com/wskurniawan/intro-microservice/service-product/config"
	"github.com/wskurniawan/intro-microservice/utils"
)

type AuthHandler struct {
	AuthService config.AuthService
}

func (handler *AuthHandler) ValidateAuth(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest(http.MethodPost, handler.AuthService.Host+"/validate-admin", nil)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header

		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody, err := ioutil.ReadAll(authResponse.Body)

		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var authResult entity.AuthResponse
		err = json.Unmarshal(responseBody, &authResult)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		if authResponse.StatusCode != http.StatusOK {
			utils.WrapAPIError(w, r, authResult.ErrorDetails, authResponse.StatusCode)
			return

		}

		context.Set(r, "user", authResult.Data.Username)

		nextHandler(w, r)

	}

}
