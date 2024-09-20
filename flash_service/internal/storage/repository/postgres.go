package repository

import (
	"database/sql"

	"github.com/Mubinabd/flash_sale/internal/pkg/config"
	"github.com/Mubinabd/flash_sale/internal/storage"
)

type Storage struct {
	OrderS           storage.OrderI
	ProductS         storage.ProductI
	AuthS            storage.AuthI
	UserS            storage.UserI
	NotificationS    storage.NotificationI
	FlashSaleS       storage.FlashSaleI
	FlashSaleProdctS storage.FlashSaleProductI
	ReviewS          storage.ReviewI
	SocialS          storage.SocialI
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		OrderS:           NewOrderRepo(db),
		ProductS:         NewProductRepo(db),
		AuthS:            NewAuthRepo(db),
		UserS:            NewUserRepo(db),
		NotificationS:    NewNotificationRepo(db, &config.Config{}),
		FlashSaleS:       NewFlashSaleRepo(db),
		FlashSaleProdctS: NewFlashSaleProductsRepo(db),
		ReviewS:          NewReviewRepo(db),
		SocialS:          NewSocialRepo(db),
	}
}

func (s *Storage) Order() storage.OrderI {
	return s.OrderS
}

func (s *Storage) Product() storage.ProductI {
	return s.ProductS
}

func (s *Storage) Auth() storage.AuthI {
	return s.AuthS
}

func (s *Storage) User() storage.UserI {
	return s.UserS
}

func (s *Storage) Notification() storage.NotificationI {
	return s.NotificationS
}

func (s *Storage) FlashSale() storage.FlashSaleI {
	return s.FlashSaleS
}

func (s *Storage) FlashSaleProduct() storage.FlashSaleProductI {
	return s.FlashSaleProdctS
}
func (s *Storage) Review() storage.ReviewI {
	return s.ReviewS
}

func (s *Storage) Social() storage.SocialI {
	return s.SocialS
}
