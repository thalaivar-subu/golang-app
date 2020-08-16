package crud

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

// Handler -> get Url has input -> CRUD operation
func Handler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	switch r.Method {
	case "GET":
		GetUsersFromEmail(w, r, db)
		break
	case "POST":
		CreateUser(w, r, db)
		break
	case "PUT":
		UpdateUser(w, r, db)
		break
	case "DELETE":
		DeleteUser(w, r, db)
		break
	}
}
