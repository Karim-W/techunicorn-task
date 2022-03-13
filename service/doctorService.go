package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/karim-w/techunicorn-task/ent"
	"github.com/karim-w/techunicorn-task/ent/user"
	"github.com/karim-w/techunicorn-task/models"
	"go.uber.org/fx"
)

type DocService struct {
	db  *ent.Client
	ctx context.Context
}

func (d *DocService) GetAll(limit int, offset int, search string) ([]models.User, error) {
	var users []models.User
	if search != "" {
		if userList, err := d.db.User.Query().Where(user.TypeEQ("doctor")).Where(user.EmailContains(search)).Limit(limit).Offset(offset).All(d.ctx); err != nil {
			return nil, err
		} else {
			for _, u := range userList {
				users = append(users, models.User{
					ID:        u.UserId.String(),
					FirstName: u.FirstName,
					LastName:  u.LastName,
					Email:     u.Email,
					Phone:     u.Phone,
					Type:      u.Type,
				})
			}
		}
	} else {
		if userList, err := d.db.User.Query().Where(user.TypeEQ("doctor")).Limit(limit).Offset(offset).All(d.ctx); err != nil {
			return nil, err
		} else {
			for _, u := range userList {
				users = append(users, models.User{
					ID:        u.UserId.String(),
					FirstName: u.FirstName,
					LastName:  u.LastName,
					Email:     u.Email,
					Phone:     u.Phone,
					Type:      u.Type,
				})
			}
		}
	}
	return users, nil

}

func (d *DocService) GetOne(id string) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	if u, err := d.db.User.Query().Where(user.UserIdEQ(uid)).First(d.ctx); err != nil {
		return nil, err
	} else {
		us := models.User{
			ID:        u.UserId.String(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
			Type:      u.Type,
		}
		return &us, nil
	}
}

func NewDocService(db *ent.Client) *DocService {
	return &DocService{
		db:  db,
		ctx: context.Background(),
	}
}

var DocServiceModule = fx.Option(fx.Provide(NewDocService))
