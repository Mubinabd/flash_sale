package app

import (
	"context"
	"log"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/usecase/service"
	"google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	auth *service.AuthService
	user *service.UserService
	product *service.ProductService
	order *service.OrderService
	flashSale *service.FlashSaleService
	flashSaleProduct *service.FlashSaleProductService
	notification *service.NotificationService
}

func (h *KafkaHandler) Register() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.RegisterReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.auth.Register(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Register User: %+v", res)
	}
}
func (h *KafkaHandler) EditProfile() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UserRes
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.user.EditProfile(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Edit profile: %+v", res)
	}
}

func (h *KafkaHandler) EditSetting() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.SettingReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.user.EditSetting(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Edit Setting: %+v", res)
	}
}
func (h *KafkaHandler) CreateFlashSale() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateFlashSalesReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.flashSale.CreateFlashSale(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Create Flash Sale: %+v", res)
	}
}
func (h *KafkaHandler) UpdateFlashSale() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateFlashSalesReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.flashSale.UpdateFlashSale(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Update Flash Sale: %+v", res)
	}
}
func (h *KafkaHandler) CreateFlashSaleProduct() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateFlashSaleProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.flashSaleProduct.CreateFlashSaleProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create flash sale product: %+v", res)
	}
}
func (h *KafkaHandler) UpdateFlashSaleProduct() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateFlashSaleProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.flashSaleProduct.UpdateFlashSaleProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("flash sale product: %+v", res)
	}
}
func (h *KafkaHandler) CreateNotification() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.NotificationCreate
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.notification.CreateNotification(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Create Notification: %+v", res)
	}

}
func (h *KafkaHandler) CreateOrder() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateOrderReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.order.CreateOrder(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create order: %+v", res)
	}
}
func (h *KafkaHandler) UpdateOrder() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateOrderReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.order.UpdateOrder(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update order: %+v", res)
	}
}
func (h *KafkaHandler) CreateProduct() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.product.CreateProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create product: %+v", res)
	}
}
func (h *KafkaHandler) UpdateProduct() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.product.UpdateProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update product: %+v", res)
	}
}