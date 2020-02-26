package secretes

import (
	"errors"
	"log"
)

type SqlInfo struct {
	userName string
	password string
	token    string
}

type G interface {
	Key() SqlInfo
}

func GetSqlInfo(q SqlInfo, token string) (map[string]string, error) {
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

func (g SqlInfo) getUserName(token string) (string, error) {
	if token != g.token {
		return "", errors.New("wrong token to get the username\n")
	}
	return g.userName, nil
}

func (g SqlInfo) getPassword(token string) (string, error) {
	if token != g.token {
		return "", errors.New("wrong token to get the password\n")
	}
	return g.password, nil
}

func reject(err error) error {
	log.Println("Access Denied.")
	return err
}
