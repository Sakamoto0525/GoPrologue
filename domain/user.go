package domain

import (
    "LayeredArchitecture/infrastructure"
    "database/sql"
)

type User struct {
    UserID   string
    Name     string
    Email    string
}

func GetUserByID(DB *sql.DB, userID string) (*User, error) {
    //インフラストラクチャレイヤの実装を利⽤する。
    userDTO, err := infrastructure.GetUserByID(DB, userID)
    if err != nil {
        return nil, err
    }
    user := &User{
        UserID: userDTO.UserID,
        Name  : userDTO.Name,
        Email : userDTO.Email 
    }
    return user, nil
}
