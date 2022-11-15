package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ong-gtp/go-chat/pkg/errors"

	"github.com/ong-gtp/go-chat/pkg/domain/requests"
	"github.com/ong-gtp/go-chat/pkg/services"
	"github.com/ong-gtp/go-chat/pkg/utils"
)

var authServ = services.NewAuthService()

func Login(w http.ResponseWriter, r *http.Request) {
	lP := requests.LoginPayload{}
	err := utils.ParseBody(r, &lP)
	errors.ErrorCheck(err)

	res, err := authServ.Login(lP.Uname, lP.Password)
	errors.ErrorCheck(err)

	res.User.Password = nil
	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	lP := requests.SignUpPayload{}
	err := utils.ParseBody(r, &lP)
	errors.ErrorCheck(err)

	res, err := authServ.SignUp(lP.Uname, lP.Password)
	errors.ErrorCheck(err)

	res.User.Password = nil
	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}
