package main

type User struct {
	Id       int64
	Login    string
	Password string
}

func (model *User) Auth(login string, password string) (*User, error) {
	err := App.Db.QueryRow(`
		SELECT id
		FROM users
		WHERE login = $1 AND password = md5($2)
	`, login, password).Scan(&model.Id)

	if err != nil {
		return nil, err
	}

	model.Login = login
	model.Password = password

	return model, nil
}

func (model *User) FindByToken(token string) (*User, error) {
	err := App.Db.QueryRow(`
		SELECT users.id, users.login
		FROM users
		JOIN user_tokens ON user_tokens.user_id = users.id
		WHERE token = $1
	`, token).Scan(&model.Id, &model.Login)

	if err != nil {
		return nil, err
	}

	return model, nil
}
