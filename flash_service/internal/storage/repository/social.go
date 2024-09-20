package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
)

type SocialRepo struct {
	DB *sql.DB
}

func NewSocialRepo(db *sql.DB) *SocialRepo {
	return &SocialRepo{
		DB: db,
	}
}

func (s *SocialRepo) ShareDeal(req *pb.ShareDealReq) (*pb.Void, error) {
	_, err := s.DB.Exec(`INSERT INTO shared_deals (user_id, flash_sale_id, platform, message, shared_at)
        VALUES ($1, $2, $3, $4, $5)`, req.UserId, req.FlashSaleId, req.Platform, req.Message, req.SharedAt)

	if err != nil {
		return nil, err
	}

	var sharesByPlatform map[string]int64
	row := s.DB.QueryRow(`
        SELECT shares_by_platform FROM sharing_stats WHERE flash_sale_id = $1
    `, req.FlashSaleId)

	var platformData string
	err = row.Scan(&platformData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if platformData != "" {
		err = json.Unmarshal([]byte(platformData), &sharesByPlatform)
		if err != nil {
			return nil, fmt.Errorf("failed to parse platform data: %v", err)
		}
	} else {
		sharesByPlatform = make(map[string]int64)
	}

	sharesByPlatform[req.Platform] += 1

	platformDataBytes, _ := json.Marshal(sharesByPlatform)

	_, err = s.DB.Exec(`
        INSERT INTO sharing_stats (flash_sale_id, total_shares, shares_by_platform)
        VALUES ($1, $2, $3)
        ON CONFLICT (flash_sale_id) DO UPDATE
        SET total_shares = sharing_stats.total_shares + 1,
            shares_by_platform = $3
    `, req.FlashSaleId, 1, platformDataBytes)

	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (s *SocialRepo) GetSharingStats(req *pb.GetSharingStatsReq) (*pb.SharingStatsRes, error) {
	var totalShares int64
	var platformData string
	sharesByPlatform := make(map[string]int64)

	row := s.DB.QueryRow(`
        SELECT total_shares, shares_by_platform
        FROM sharing_stats
        WHERE flash_sale_id = $1
    `, req.FlashSaleId)

	err := row.Scan(&totalShares, &platformData)
	if err != nil {
		return nil, err
	}

	if platformData != "" {
		err = json.Unmarshal([]byte(platformData), &sharesByPlatform)
		if err != nil {
			return nil, fmt.Errorf("failed to parse platform data: %v", err)
		}
	}

	return &pb.SharingStatsRes{
		FlashSaleId:       req.FlashSaleId,
		TotalShares:       totalShares,
		SharesByPlatform: sharesByPlatform,
	}, nil
}
