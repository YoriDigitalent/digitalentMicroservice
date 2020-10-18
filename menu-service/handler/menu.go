package handler

import (
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/utils"
)

// AddMenu handle add menu
func AddMenu(w http.ResponseWriter, r *http.Request) {

	utils.WrapAPISuccess(w, r, "success", 200)
}
