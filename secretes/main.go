package secretes

import (
	"errors"
	"log"
)

type sqlInfo struct {
	userName string
	password string
	token    string
}

func GetSqlInfo(q sqlInfo, token string) (map[string]string, error) {
	u, err := q.getUserName(token)
	p, err := q.getPassword(token)
	if err != nil {
		return nil, reject(err)
	}
	return map[string]string{
		"UserName": u,
		"Password": p,
	}, nil
}

func (q sqlInfo) getUserName(token string) (string, error) {
	if token != q.token {
		return "", errors.New("wrong token to get the username\n")
	}
	return q.userName, nil
}

func (q sqlInfo) getPassword(token string) (string, error) {
	if token != q.token {
		return "", errors.New("wrong token to get the password\n")
	}
	return q.password, nil
}

func reject(err error) error {
	log.Println("Access Denied.")
	return err
}
