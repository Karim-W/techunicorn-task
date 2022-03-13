package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/karim-w/techunicorn-task/ent"
	"github.com/karim-w/techunicorn-task/ent/user"
	"github.com/karim-w/techunicorn-task/models"
	"github.com/karim-w/techunicorn-task/utils"
	"go.uber.org/fx"
)

type AuthService struct {
	db  *ent.Client
	ctx context.Context
}

func (a *AuthService) Register(user *models.User) (*models.User, string, error) {
	if hashedText, err := utils.HashPassword(user.Password); err != nil {
		return nil, "", err
	} else {
		user.Password = hashedText
	}
	if dbUser, err := a.db.User.Create().SetUserId(uuid.New()).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetPhone(user.Phone).
		SetType(user.Type).
		Save(a.ctx); err != nil {
		return nil, "", err
	} else {
		userModel := models.User{
			ID:        dbUser.UserId.String(),
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
			Email:     dbUser.Email,
			Phone:     dbUser.Phone,
			Type:      dbUser.Type,
		}
		if token, err := utils.CreateToken(&userModel); err != nil {
			return nil, "", err
		} else {
			return &userModel, token, nil
		}

	}
}
func (a *AuthService) Login(use *models.User) (*string, error) {
	if u, err := a.db.User.Query().Where(user.EmailEQ(use.Email)).Only(a.ctx); err != nil {
		return nil, err
	} else {
		if ok := utils.VerifyPasswordHash(use.Password, u.Password); !ok {
			return nil, fmt.Errorf("password is incorrect")
		} else {
			userModel := models.User{
				ID:   u.UserId.String(),
				Type: u.Type,
			}
			if token, err := utils.CreateToken(&userModel); err != nil {
				return nil, err
			} else {
				return &token, nil
			}
		}
	}
}
func NewAuthService(db *ent.Client) *AuthService {
	return &AuthService{
		db:  db,
		ctx: context.Background(),
	}
}

var AuthServiceModule = fx.Option(fx.Provide(NewAuthService))
