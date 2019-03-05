package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"unique_index"`
	Email     string     `gorm:"unique_index"`
	Avatar    string     `json:"avatar"`
	Pwd       string     `json:"-"`
	Role      int        `gorm:"default:0" json:"role"` // 0 管理员 1正常用户
	Editor    string     `json:"editor"`
}

func (db *DB) QueryUserByEmailAndPassword(email, password string) (user User, err error) {
	return user, db.db.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error
}

func (db *DB) QueryUserByName(name string) (user User, err error) {

	return user, db.db.Where("name = ?", name).Take(&user).Error
}

func (db *DB) QueryUserByEmail(email string) (user User, err error) {
	return user, db.db.Where("email = ?", email).Take(&user).Error
}

func (db *DB) UpdateUserEditor(editor string) (err error) {
	return db.db.Model(&User{}).Update("editor", editor).Error
}

func SaveUser(user *User) error {
	return db.Create(user).Error
}
