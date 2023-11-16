package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"server/database"
	"server/models"
)

type LoginData struct {
	Email    string
	Password string
}

type Message struct {
	Msg     string
	User_id string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data LoginData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		msgFail := map[string]string{
			"msg": "err when body parse",
		}
		json.NewEncoder(w).Encode(msgFail)
		return
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Email == "" {
		msgFail := map[string]string{
			"msg": "email not found!",
		}
		json.NewEncoder(w).Encode(msgFail)
		return
	}

	if user.Password != data.Password {
		msgFail := map[string]string{
			"msg": "password incorrect!",
		}
		json.NewEncoder(w).Encode(msgFail)
		return
	}

	id := strconv.FormatUint(uint64(user.Id), 10)

	msgSuccess := map[string]string{
		"msg":     "success",
		"user_id": id,
	}

	cookie := http.Cookie{
		Name:    "user_id",
		Value:   "1",
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 2), // 2hour
	}

	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(msgSuccess)
}
