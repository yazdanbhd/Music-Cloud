package mysqldb

import (
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
