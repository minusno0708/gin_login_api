package models

import (
    "jwt-gin/utils/token"
    "strings"

    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    gorm.Model
    UserID string `gorm:"size:255;not null;unique" json:"user_id"`
    Username string `gorm:"size:255;not null;" json:"username"`
    Password string `gorm:"size:255;not null;" json:"password"`
}

func (u User) Save() (User, error) {
    err := DB.Create(&u).Error
    if err != nil {
        return User{}, err
    }
    return u, nil
}

func (u *User) BeforeSave() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    u.Password = string(hashedPassword)

    u.Username = strings.ToLower(u.Username)

    return nil
}

func (u User) PrepareOutput() User {
    u.Password = ""
    return u
}

func GenerateToken(user_id string, password string) (string, error) {
    var user User
    
    err := DB.Where("user_id = ?", user_id).First(&user).Error

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

    if err != nil {
        return "", err
    }

    token, err := token.GenerateToken(user.ID)

    if err != nil {
        return "", err
    }

    return token, nil
}