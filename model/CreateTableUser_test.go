package model_test

import (
	"testing"
	"todo/model"
)

func TestCreateTableUser(t *testing.T) {
	InitMySQL()
	err := model.CreateTableUser()
	if err != nil {
		t.Error(err.Error())
	}
}
