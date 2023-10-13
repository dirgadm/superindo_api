package domain

import (
	"context"
	"project-version3/superindo-task/service/domain/dto"
	"time"
)

// User is representing the User data struct
type User struct {
	Id          int `gorm:"primaryKey;autoIncrement:true"`
	Email       string
	Password    string
	Name        string
	Phone       string
	CreatedAt   time.Time
	LastLoginAt time.Time `gorm:"column:last_login_at;type:datetime"`
}

func (m *User) TableName() string {
	return "user"
}

// UserUsecase represent the article's usecases
type UserUsecase interface {
	Login(ctx context.Context, req dto.LoginRequest) (res dto.AuthResponse, err error)
	Register(ctx context.Context, req dto.RegisterRequest) (res dto.AuthResponse, err error)
}

// UserRepository represent the article's repository contract
type UserRepository interface {
	GetDetail(ctx context.Context, id int) (user User, err error)
	GetByEmail(ctx context.Context, email string) (user User, err error)
	Create(ctx context.Context, user *User) (err error)
}
