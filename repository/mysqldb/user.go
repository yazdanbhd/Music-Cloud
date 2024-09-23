package mysqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yazdanbhd/Music-Cloud/entity"
)

func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	result, err := d.db.Exec(`insert into users (name, phone_number, password) values (?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("unexpected error: %w", err)
	}
	id, _ := result.LastInsertId()
	u.ID = uint(id)
	return u, nil
}

func (d *MySQLDB) IsAuthenticated(userName, password string) (bool, error) {
	var user entity.User
	var createdAt []uint8

	query := d.db.QueryRow("select * from users where name = ? and password = ?", userName, password)
	err := query.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		} else {
			return true, err
		}
	}

	return true, nil

}
