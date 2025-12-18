package tests

import (
	"amass-test/models"
	"amass-test/repositories"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAuthRepository_CheckUsernameExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewAuthRepository(gormDB)

	t.Run("Username Exists", func(t *testing.T) {
		username := "duplicate_user"

		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users`")).
			WithArgs(username).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		exists, err := repo.CheckUsernameExists(username)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Username Not Exists", func(t *testing.T) {
		username := "new_user"

		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users`")).
			WithArgs(username).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		exists, err := repo.CheckUsernameExists(username)
		assert.NoError(t, err)
		assert.False(t, exists)
	})
}

func TestAuthRepository_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewAuthRepository(gormDB)

	t.Run("User Found", func(t *testing.T) {
		username := "test_user"

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`") + ".*").
			WithArgs(username, 1).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "created_at", "updated_at", "deleted_at"}).
				AddRow("uuid-123", username, time.Now(), time.Now(), nil))

		user, err := repo.GetByUsername(username)
		assert.NoError(t, err)
		assert.Equal(t, username, user.Username)
	})
}

func TestAuthRepository_Register(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewAuthRepository(gormDB)

	t.Run("Register Success", func(t *testing.T) {
		user := models.User{
			UserID:   "uuid-123",
			Username: "newuser",
			Password: "hashedpassword",
		}

		mock.ExpectBegin()

		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
			WithArgs(
				user.UserID,
				user.Username,
				user.Password,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		_, err := repo.Register(user)
		assert.NoError(t, err)
	})
}
