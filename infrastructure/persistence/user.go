package persistence

import (
    "LayeredArchitecture/domain"
    "LayeredArchitecture/domain/repository"
    "database/sql"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
    return &userPersistence{}
}

//ユーザ登録
func (up userPersistence) Insert(DB *sql.DB, userID, name, email string) error {
    stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email) VALUES(?, ?, ?)")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(userID, name, email)
    return err
}

//userIDによってユーザ情報を取得する
func (up userPersistence) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
    row := DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userID)
    //row型をgolangで利用できる形にキャストする。
    return convertToUser(row)
}

//row型をuser型に紐づける
func convertToUser(row *sql.Row) (*domain.User, error) {
    user := domain.User{}
    err := row.Scan(&user.UserID, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}