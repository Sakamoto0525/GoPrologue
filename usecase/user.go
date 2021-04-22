package usecase

import (
    "database/sql"
    "errors"
    "github.com/google/uuid"
    "LayeredArchitecture/domain"
    "LayeredArchitecture/domain/repository"
)

// User における UseCase のインターフェース
type UserUseCase interface {
    GetByUesrID(DB *sql.DB, userID string) (domain.User, error)
    Insert(DB *sql.DB, userID, name, email string) error
}

type userUseCase struct {
    userRepository repository.UserRepository
}

// Userデータに対するusecaseを生成
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
    return &userUseCase{
        userRepository: ur,
    }
}

func (uu UserUsecase) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
    user, err := uu.userRepository.GetByUserID(DB, userID)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (uu UserUsecase) Insert(DB *sql.DB, name, email string) error {
    //本来ならemailのバリデーションをする

    //一意でランダムな文字列を生成する
    userID, err := uuid.NewRandom()//返り値はuuid型
    if err != nil {
        return err
    }

    //domainを介してinfrastructureで実装した関数を呼び出す。
    // Persistence（Repository）を呼出
    err = uu.userRepository.Insert(DB, userID.String(), name, email)
    if err != nil {
        return err
    }
    return nil
}
