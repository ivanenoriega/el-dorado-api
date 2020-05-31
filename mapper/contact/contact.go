package contact

import (
	"database/sql"

	m "github.com/ivanenoriega/el-dorado-api/model/contact"
)

func SelectByID(id int, db *sql.DB) (m.Contact, error) {
	selDB, err := db.Query("SELECT * FROM contact WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	c := m.Contact{}

	for selDB.Next() {
		var id int
		var name string
		var email string
		var phone string
		var subject string
		var message string
		err = selDB.Scan(&id, &name, &email, &phone, &subject, &message)
		if err != nil {
			panic(err.Error())
		}
		c.Id = id
		c.Name = name
		c.Email = email
		c.Phone = phone
		c.Subject = subject
		c.Message = message
	}

	return c, nil
}

func Insert(name string, email string, phone string, subject string, message string, db *sql.DB) error {
	_, err := db.Query("INSERT INTO contact (name, email, phone, subject, message) VALUES (?,?,?,?,?)", name, email, phone, subject, message)
	if err != nil {
		return err
	}
	return nil
}
