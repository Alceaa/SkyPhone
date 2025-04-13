package tests

import (
	"testing"

	"github.com/Alceaa/SkyPhone/handlers"
	"github.com/Alceaa/SkyPhone/models"
	"github.com/Alceaa/SkyPhone/utils"
	"github.com/Alceaa/SkyPhone/validators"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("Ошибка хеширования: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		t.Errorf("Хешированный пароль не совпадает с оригинальным")
	}
}

func TestValidateUser(t *testing.T) {
	username := "Ivan"
	password := "123"

	user := models.User{Username: username, Password: password}

	err := validators.ValidateUser(user)
	if err != nil {
		t.Errorf("Ошибка валидации пользователя: %v", err)
	} else {
		return
	}
}

func TestGetUser(t *testing.T) {
	username := "Ivan"

	var dbUser models.User
	err := handlers.DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1", username)

	if err != nil {
		t.Errorf("Ошибка получения пользователя: %v", err)
	}

	if dbUser.Username != username {
		t.Errorf("Ожидаемый логин %s, получен %s", username, dbUser.Username)
	}
}
