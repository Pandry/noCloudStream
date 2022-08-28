package sqlite

import (
	"auther/models"
	"time"
)

func (d *Database) CreateUser(user models.User) error {
	user.CreatedAt = time.Now()
	user.HashUserPassword()
	result := d.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *Database) GetUserByMail(mail string) (models.User, error) {
	user := models.User{}
	result := d.db.First(&user, "email = ?", mail)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
