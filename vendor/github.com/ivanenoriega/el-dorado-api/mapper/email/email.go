package email

import (
	"database/sql"

	m "github.com/ivanenoriega/el-dorado-api/model/email"
)

func SelectByID(id int, db *sql.DB) (m.Email, error) {
	selDB, err := db.Query("SELECT * FROM email WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	e := m.Email{}
	for selDB.Next() {
		var id int
		var email string
		err = selDB.Scan(&id, &email)
		if err != nil {
			panic(err.Error())
		}
		e.Id = id
		e.Email = email
	}
	return e, nil
}

func Insert(email string, db *sql.DB) error {
	_, err := db.Query("INSERT INTO email (email) VALUES (?)", email)
	if err != nil {
		return err
	}
	return nil
}
