package types

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type HandlerWithDb func(w http.ResponseWriter, r *http.Request, db *gorm.DB)
