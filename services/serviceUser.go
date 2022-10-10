package services

import (
	"dtskominfo-gin/apimodels"
	"dtskominfo-gin/auth"
	"dtskominfo-gin/database"
	"dtskominfo-gin/models"
	"dtskominfo-gin/utils"
	"errors"
)

func Registrasi(req apimodels.ReqRegistration) (res apimodels.ResRegistration) {
	res = apimodels.ResRegistration{
		Status:  -2,
		Message: "Failed on Process",
	}
	user := models.User{
		Name:     req.Name,
		Username: req.User,
		Email:    req.Email,
	}
	passhash, err := utils.HashPassword(req.Password)
	if err != nil {
		return res
	}
	user.Password = passhash
	err = database.GetDb().Create(&user).Error
	if err != nil {
		return res
	}
	res.Status = 200
	res.Message = "User Create Success"
	return res
}

func GetToken(req apimodels.ReqLogin) (token string, err error) {
	user := models.User{}
	db := database.GetDb()
	err = db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return "", err
	}
	if !utils.ValidatePassword(user.Password, req.Password) {
		return "", errors.New("Password Missmatch")
	}
	token, err = auth.GenerateJWT(req.Email, user.Username)
	if err != nil {
		return "", errors.New("Generate Token Failed")
	}
	return token, nil
}
