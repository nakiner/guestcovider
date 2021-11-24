package userRepository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/nakiner/guestcovider/internal/database"
)

var (
	ConnError = errors.New("get connection error")
)

type Repository interface {
	FindBySurname(ctx context.Context, surname string) ([]*User, error)
	UpdateUser(ctx context.Context, data *User) error
}

type userDBRepository struct {
	dbConn *database.Connection
}

func NewUserDBRepository(pool *database.Connection) Repository {
	return &userDBRepository{dbConn: pool}
}

func (r *userDBRepository) FindBySurname(ctx context.Context, surname string) ([]*User, error) {
	conn, err := database.GetMasterConn(ctx, r.dbConn)

	if err != nil {
		return nil, errors.Wrap(ConnError, err.Error())
	}

	var records []*User

	if err := conn.Debug().Where("surname ilike ?", "%" + surname + "%").Find(&records).Error; err != nil {
		return nil, errors.Wrap(err, err.Error())
	}


	return records, nil
}

func (r *userDBRepository) UpdateUser(ctx context.Context, data *User) error {
	conn, err := database.GetMasterConn(ctx, r.dbConn)

	if err != nil {
		return errors.Wrap(ConnError, err.Error())
	}

	var record User

	if err := conn.First(&record, data.ID).Error; err != nil {
		return err
	}

	// small fix
	record.CovidPass = data.CovidPass
	record.Checkin = data.Checkin

	if err := conn.Model(&record).Select("*").Updates(&record).Error; err != nil {
		return err
	}

	return nil
}