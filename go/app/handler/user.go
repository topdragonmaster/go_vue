package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sniper/door/app/middleware"
	"sniper/door/app/model"

	"gorm.io/gorm"
)

type BaseHandler struct {
	db *gorm.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func (handler *BaseHandler) Login(w http.ResponseWriter, r *http.Request) {

	type Props struct {
		Email    string
		Password string
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	var props Props
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&props)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	fmt.Printf("%+v \n", props)

	user := model.User{Email: props.Email, Password: props.Password}
	existUser := model.User{}
	// error := db.First(&existUser, user).Error

	if err := handler.db.First(&existUser, user).Error; err != nil {
		respondError(w, http.StatusNotFound, "no exist user")
		return
	}

	type Res struct {
		ExistUser model.User
		Token     string
	}

	tokenString, err := middleware.GenerateJWT(existUser.Email, existUser.UserName, existUser.UserRole)
	respondJSON(w, http.StatusOK, Res{ExistUser: existUser, Token: tokenString})

}

func (handler *BaseHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	type Props struct {
		UserName string
		Email    string
		Password string
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	var props Props
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&props)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user := model.User{
		Email:    props.Email,
		Password: props.Password,
		UserName: props.UserName,
		UserRole: "3",
	}
	if err := handler.db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func (handler *BaseHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	type Props struct {
		Email    string
		Password string
	}

	fmt.Printf("%v+", r.Header.Get("email"))
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	var props Props
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&props)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	fmt.Printf("%+v \n", props)

	user := model.User{Email: props.Email, Password: props.Password}
	existUser := model.User{}
	// error := db.First(&existUser, user).Error

	if err := handler.db.First(&existUser, user).Error; err != nil {
		respondError(w, http.StatusNotFound, "no exist user")
		return
	}

	type Res struct {
		ExistUser model.User
		Token     string
	}

	tokenString, err := middleware.GenerateJWT(existUser.Email, existUser.UserName, existUser.UserRole)
	respondJSON(w, http.StatusOK, Res{ExistUser: existUser, Token: tokenString})
}
