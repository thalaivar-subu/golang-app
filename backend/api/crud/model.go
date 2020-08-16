package crud

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/thalaivar-subu/golang-app/backend/structs"
)

func GetUsersFromEmail(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	queryParams := r.URL.Query()
	email, ok := queryParams["email"]
	if !ok {
		glog.Info("Email Param is Missing")
	}
	records := structs.User{}
	err := db.Where("email = ?", email[0]).Find(&records).Error
	if err != nil {
		glog.Info(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJSON, err := json.Marshal(map[string]string{
		"message": "Successfully created",
		"name":    records.Name, "email": records.Email})
	if err != nil {
		glog.Info(err)
	}
	w.Write(responseJSON)

}

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	var data structs.InsertBody
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	err = db.Create(&structs.User{Name: data.Name, Email: data.Email}).Error
	if err != nil {
		glog.Info(err)
	}
	responseJSON, err := json.Marshal(map[string]string{
		"message": "Successfully created"})
	if err != nil {
		glog.Info(err)
	}
	w.Write(responseJSON)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	var data structs.InsertBody
	err := decoder.Decode(&data)
	if err != nil {
		glog.Info(err)
	}
	err = db.Model(&structs.User{}).Where("id = ?", data.ID).Updates(map[string]interface{}{"name": data.Name, "email": data.Email}).Error
	if err != nil {
		glog.Info(err)
	}
	responseJSON, err := json.Marshal(map[string]string{
		"message": "Successfully Updated"})
	if err != nil {
		glog.Info(err)
	}
	w.Write(responseJSON)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := r.URL.Query().Get("id")
	err := db.Where("id = ?", id).Delete(structs.User{}).Error
	if err != nil {
		glog.Info(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJSON, err := json.Marshal(map[string]string{
		"message": "Deleted Successfully",
		"id":      id})
	if err != nil {
		glog.Info(err)
	}
	w.Write(responseJSON)

}

// curl -XPOST http://127.0.0.1:3001/api/v1/crud -H 'Content-Type: application/json' -d '{"name":"subu","email": "vesubramanian1996@gmail.com"}'
// curl -XDELETE http://127.0.0.1:3001/api/v1/crud?id=1
// curl -XPUT http://127.0.0.1:3001/api/v1/crud  -H 'Content-Type: application/json' -d '{"id": 2,"name":"america","email": "vesubramanian1996@gmail.com"}'
