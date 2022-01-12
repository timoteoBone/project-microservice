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

	stmt, err := repo.DB.Prepare(utils.CreateUser)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return "", err
	}

	newId := utils.GenerateId()

	res, err := stmt.Exec(user.Name, newId, user.Age, user.Pass)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return "", err
	}

	repo.Logger.Log(repo.Logger, res, "rows affected")

	return newId, nil
}

func (repo *sqlRepo) GetUser(ctx context.Context, userId string) (entities.User, error) {

	stmt, err := repo.DB.Query(utils.GetUser, userId)
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{}
	for stmt.Next() {
		err := stmt.Scan(&user.Name, &user.Age)
		if err != nil {
			level.Error(repo.Logger).Log("error")
		}
	}
	return user, nil
}
