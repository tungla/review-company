package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"

  "golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func HashPasswordBcrypt(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func ConnectDataBase() {
  database, err := gorm.Open("sqlite3", "test.db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Company{})
  database.AutoMigrate(&User{})
    database.AutoMigrate(&CompanyReview{})

  DB = database
}