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

	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	var dataMenu database.Menu
	err = json.Unmarshal(body, &dataMenu)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = dataMenu.Insert(menu.Db)
	if err != nil {
		utils.WrapAPIError(w, r, "insert menu error : "+err.Error(), http.StatusInternalServerError)
	}

	utils.WrapAPISuccess(w, r, "success", 200)
}
