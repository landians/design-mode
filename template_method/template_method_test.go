package template_method

import (
	"strings"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func Test_TemplateMethod(t *testing.T) {
	// setup sqlmock ///////////////////////////////////////////
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock error: '%s'", err)
	}
	defer func() {
		_ = db.Close()
	}()
	// end setup sqlmock ///////////////////////////////////////////

	// test UserDAO.GetUserByID ////////////////////////////////////
	mock.ExpectQuery("select").WillReturnRows(
		mock.
			NewRows(strings.Split("id,name,pwd,org_id,role_id", ",")).
			AddRow(1, "John", "abcdefg", 11, "guest"))

	ud := newUserDAO()
	e, u := ud.GetUserByID(db, 1)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("user = %v", u)
	}
	// end test UserDAO.GetUserByID ////////////////////////////////

	// test UserDAO.GetUsersByOrgID ///////////////////////////////
	mock.ExpectQuery("select").WillReturnRows(
		mock.
			NewRows(strings.Split("id,name,pwd,org_id,role_id", ",")).
			AddRow(1, "John", "abcdefg", 11, "guest").
			AddRow(2, "Mike", "aaaaaa", 11, "admin"))

	e, ul := ud.GetUsersByOrgID(db, 11)
	if e != nil {
		t.Error(e)
	} else {
		for i, it := range ul {
			t.Logf("users[%d] = %v", i, it)
		}
	}
	// end test UserDAO.GetUsersByOrgID ///////////////////////////
}
