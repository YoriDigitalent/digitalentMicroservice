package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/database"
	"github.com/YoriDigitalent/digitalentMicroservice/utils"
	"gorm.io/gorm"
)

type Menu struct {
	Db *gorm.DB
}

// AddMenu handle add menu
func (menu *Menu) AddMenu(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var datamenu database.Menu
	err = json.Unmarshal(body, &menu)

	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	err = datamenu.Insert(menu.Db)

	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
	return
}

// GetMenu handle get menu
func (handler *Menu) GetMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	menu := database.Menu{}

	menus, err := menu.GetAll(handler.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(w, r, menus, http.StatusOK, "success")
}
