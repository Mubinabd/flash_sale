package repository

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
)

func TestShareDeal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewSocialRepo(db)

	mock.ExpectExec(`INSERT INTO shared_deals`).
		WithArgs("user1", "flashsale1", "Facebook", "Great deal!", time.Now().Format(time.RFC3339)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"shares_by_platform"}).
		AddRow(`{"Facebook": 1}`)
	mock.ExpectQuery(`SELECT shares_by_platform FROM sharing_stats`).
		WithArgs("flashsale1").
		WillReturnRows(rows)

	mock.ExpectExec(`INSERT INTO sharing_stats`).
		WithArgs("flashsale1", 1, `{"Facebook":2}`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.ShareDealReq{
		UserId:      "user1",
		FlashSaleId: "flashsale1",
		Platform:    "Facebook",
		Message:     "Great deal!",
		SharedAt:    time.Now().Format(time.RFC3339),
	}

	_, err = repo.ShareDeal(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetSharingStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewSocialRepo(db)

	sharesByPlatform := map[string]int64{"Facebook": 5, "Twitter": 2}
	sharesByPlatformJSON, _ := json.Marshal(sharesByPlatform)

	mock.ExpectQuery(`SELECT total_shares, shares_by_platform`).
		WithArgs("flashsale1").
		WillReturnRows(sqlmock.NewRows([]string{"total_shares", "shares_by_platform"}).
			AddRow(7, sharesByPlatformJSON))

	req := &pb.GetSharingStatsReq{
		FlashSaleId: "flashsale1",
	}

	res, err := repo.GetSharingStats(req)
	assert.NoError(t, err)
	assert.Equal(t, "flashsale1", res.FlashSaleId)
	assert.Equal(t, int64(7), res.TotalShares)
	assert.Equal(t, sharesByPlatform, res.SharesByPlatform)
	assert.NoError(t, mock.ExpectationsWereMet())
}
