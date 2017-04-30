package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func (app *Application) issueToken(
	writer http.ResponseWriter,
	request *http.Request,
) {
	login := request.FormValue("login")
	password := request.FormValue("password")

	user, err := (&User{}).Auth(login, password)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Wrong user credentials: ", err)
		return
	}

	token, err := (&UserToken{UserId: user.Id}).Issue(user.Login)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, err)
		return
	}

	writer.WriteHeader(http.StatusAccepted)
	fmt.Fprint(writer, token.Token)
}

func (app *Application) getLoginByToken(
	writer http.ResponseWriter,
	request *http.Request,
) {
	token := request.FormValue("token")

	user, err := (&User{}).FindByToken(token)
	if err != nil {
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprint(writer, "Not Found: ", err)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(writer, err)
		}
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, user.Login)
}
