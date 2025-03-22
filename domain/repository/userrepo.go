package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/muhammad-reda/go-api-gin/domain/entity"
)

type UserRepository interface {
	FindAll(context.Context) ([]entity.User, error)
	FindById(ctx context.Context, id int64) (entity.User, error)
	Save(ctx context.Context, user entity.User) (*entity.User, error)
	Update(ctx context.Context, user entity.User, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserRepositoryImplementation struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserRepositoryImplementation{
		DB,
	}
}

func (ur *UserRepositoryImplementation) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	query := "SELECT id, email, username, password, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL"

	rows, errQuery := ur.DB.QueryContext(ctx, query)
	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User

		errScan := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if errScan != nil {
			return nil, errScan
		}

		users = append(users, user)

	}

	if rows.Err() != nil {
		return nil, errQuery
	}

	return users, nil
}

func (ur *UserRepositoryImplementation) FindById(ctx context.Context, id int64) (entity.User, error) {
	var user entity.User

	query := "SELECT id, email, username, password, created_at, updated_at, deleted_at FROM users WHERE id = ? AND deleted_at IS NULL"
	db := ur.DB

	errScan := db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return user, sql.ErrNoRows
		}
		return user, errScan
	}

	return user, nil
}

func (ur *UserRepositoryImplementation) Save(ctx context.Context, user entity.User) (*entity.User, error) {
	query := "INSERT INTO users (email, username, password) VALUES (?, ?, ?)"
	_, errExec := ur.DB.ExecContext(ctx, query, user.Email, user.Username, user.Password)
	if errExec != nil {

		switch errExec.(*mysql.MySQLError).Number {
		case 1062:
			return nil, errors.New("email or username already exist")
		}

		return nil, errExec
	}

	return &user, nil
}

func (ur *UserRepositoryImplementation) Update(ctx context.Context, user entity.User, id int64) (*entity.User, error) {
	query := "UPDATE users SET email = ?, username = ?, password = ? WHERE id = ? AND deleted_at IS NULL"

	_, errExec := ur.DB.ExecContext(ctx, query, user.Email, user.Username, user.Password, id)
	if errExec != nil {
		switch errExec.(*mysql.MySQLError).Number {
		case 1062:
			return nil, errors.New("email or username already exist")
		}

		return nil, errExec
	}

	return &user, nil
}

func (ur *UserRepositoryImplementation) Delete(ctx context.Context, id int64) error {
	query := "UPDATE users SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL"

	_, errExec := ur.DB.ExecContext(ctx, query, time.Now(), id)
	if errExec != nil {
		if errExec == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		return errExec
	}

	return nil
}
