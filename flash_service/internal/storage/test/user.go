package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)

	req := &pb.GetByID{Id: "1"}
	expectedUser := &pb.UserRes{
		Id:          "1",
		Username:    "testuser",
		Email:       "test@example.com",
		FullName:    "Test User",
		DateOfBirth: "2000-01-01",
		Role:        "user",
	}

	mock.ExpectQuery("SELECT id, username, email, full_name, date_of_birth, role").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "full_name", "date_of_birth", "role"}).
			AddRow(expectedUser.Id, expectedUser.Username, expectedUser.Email, expectedUser.FullName, "2000-01-01", expectedUser.Role))

	res, err := repo.GetProfile(req)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, res)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEditProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)

	req := &pb.UserRes{
		Id:          "1",
		Username:    "updateduser",
		Email:       "updated@example.com",
		FullName:    "Updated User",
		DateOfBirth: "1990-05-15",
	}

	mock.ExpectExec("UPDATE users SET updated_at = NOW()").
		WithArgs(req.Username, req.Email, req.FullName, req.DateOfBirth, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.EditProfile(req)
	assert.NoError(t, err)
	assert.Empty(t, res)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestChangePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)

	req := &pb.ChangePasswordReq{
		Id:              "1",
		CurrentPassword: "currentpass",
		NewPassword:     "newpass",
	}

	mock.ExpectQuery("SELECT password FROM users WHERE id = $1").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("currentpass"))

	mock.ExpectExec("UPDATE users SET updated_at = NOW(), password = $1 WHERE id = $2").
		WithArgs(req.NewPassword, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.ChangePassword(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetSetting(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)

	req := &pb.GetByID{Id: "1"}
	expectedSetting := &pb.Setting{
		PrivacyLevel: "high",
		Notification: "enabled",
		Language:     "en",
		Theme:        "dark",
	}

	mock.ExpectQuery("SELECT privacy_level, notification, language, theme").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"privacy_level", "notification", "language", "theme"}).
			AddRow(expectedSetting.PrivacyLevel, expectedSetting.Notification, expectedSetting.Language, expectedSetting.Theme))

	res, err := repo.GetSetting(req)
	assert.NoError(t, err)
	assert.Equal(t, expectedSetting, res)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEditSetting(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)

	req := &pb.SettingReq{
		Id:           "1",
		PrivacyLevel: "medium",
		Notification: "disabled",
		Language:     "es",
		Theme:        "light",
	}

	mock.ExpectExec("UPDATE settings SET updated_at = NOW()").
		WithArgs(req.PrivacyLevel, req.Notification, req.Language, req.Theme, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.EditSetting(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)

	req := &pb.GetByID{Id: "1"}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE users SET deleted_at = EXTRACT").WithArgs(req.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM settings WHERE user_id = $1").WithArgs(req.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = repo.DeleteUser(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
