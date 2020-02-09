package jwt

import (
    "smart-hr/library/logger"
	"golang.org/x/crypto/bcrypt"
)

type JWT struct{}

func(j *JWT) HashPassword(password string) (hashPassword string, err error){
	bytePwd := []byte(password)
	btyeHash, errHash :=bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if errHash != nil {
		logger.Log.Println(errHash.Error())
		return "", errHash
	}

	return string(btyeHash), nil
}

func(j *JWT) ComparePassword(password string, hash string) (err error) {
	pwd := []byte(password)
	hashPwd := []byte(hash)
	errMatch := bcrypt.CompareHashAndPassword(hashPwd, pwd)
	if errMatch != nil {
		logger.Log.Println(errMatch)
		return errMatch
	}
	return nil
}