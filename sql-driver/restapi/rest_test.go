package restapi_test

import (
	"database/sql"
	_ "github.com/zhang1career/lib/sql-driver/restapi"
	"testing"
)

func Test_Open(t *testing.T) {
	_, err := sql.Open("rest", "user:password@tcp(127.0.0.1:8080)/v1")
	if err != nil {
		t.Error(err.Error())
	}
}

func Test_Exec(t *testing.T) {
	db, err := sql.Open("rest", "user:password@tcp(127.0.0.1:8080)/v1")
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := db.Exec("SELECT * FROM rules where Id=2")
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(res)
}