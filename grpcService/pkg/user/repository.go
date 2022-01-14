package user

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	entities "github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/utils"
)

type sqlRepo struct {
	DB     *sql.DB
	Logger log.Logger
}

func NewSQL(db *sql.DB, log log.Logger) *sqlRepo {
	return &sqlRepo{db, log}
}

func (repo *sqlRepo) CreateUser(ctx context.Context, user entities.User) (string, error) {

	repo.Logger.Log(repo.Logger, "Repository method", "Create user")

	stmt, err := repo.DB.PrepareContext(ctx, utils.CreateUser)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return "", err
	}

	newId := utils.GenerateId()

	res, err := stmt.ExecContext(ctx, user.Name, newId, user.Pass, user.Age, user.Email)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return "", err
	}

	repo.Logger.Log(repo.Logger, res, "rows affected")

	return newId, nil
}

func (repo *sqlRepo) GetUser(ctx context.Context, userId string) (entities.User, error) {

	user := entities.User{}
	stmt, err := repo.DB.PrepareContext(ctx, utils.GetUser)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return entities.User{}, err
	}
	err = stmt.QueryRowContext(ctx, utils.GetUser).Scan(&user.Name, &user.Age, &user.Email)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (repo *sqlRepo) AuthenticateUser(ctx context.Context, email string) (string, error) {

	stmt, err := repo.DB.Prepare(utils.GetPassword)
	if err != nil {
		return "", err
	}

	var password string
	err = stmt.QueryRow(email).Scan(password)
	if err != nil {
		return "", err
	}

	return password, nil
}
