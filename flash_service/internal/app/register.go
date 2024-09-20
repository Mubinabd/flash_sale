package app

import (
	"errors"

	"github.com/Mubinabd/flash_sale/internal/pkg/config"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

func Register(h *KafkaHandler, cfg *config.Config) error {

	brokers := []string{cfg.KafkaUrl}
	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "create", "create-id", h.Register()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update", "update-id", h.EditProfile()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "edit", "edit", h.EditSetting()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'edit' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	
	if err := kcm.RegisterConsumer(brokers, "update-flash", "update-flash-id", h.UpdateFlashSale()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-flash' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-flash-sale", "create-flash-sale-id", h.CreateFlashSaleProduct()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-flash-sale' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-flash-sale", "update-flash-sale-id", h.UpdateFlashSaleProduct()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-flash-sale' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "notif", "notif-id", h.CreateNotification()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'notif' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	
	if err := kcm.RegisterConsumer(brokers, "update-order", "update-order-id", h.UpdateOrder()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-order' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-product", "create-product-id", h.CreateProduct()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-product' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-product", "update-product-id", h.UpdateProduct()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-product' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	return nil
}
