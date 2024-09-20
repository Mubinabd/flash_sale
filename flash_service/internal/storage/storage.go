package storage

import (
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
)

type StorageI interface {
	Auth() AuthI
	User() UserI
	FlashSale() FlashSaleI
	FlashSaleProduct() FlashSaleProductI
	Notification() NotificationI
	Order() OrderI
	Product() ProductI
	Review() ReviewI
	Social() SocialI
}
type AuthI interface {
	Register(req *pb.RegisterReq) (*pb.Void, error)
	Login(req *pb.LoginReq) (*pb.User, error)
	ForgotPassword(req *pb.GetByEmail) (*pb.Void, error)
	ResetPassword(req *pb.ResetPassReq) (*pb.Void, error)
	SaveRefreshToken(req *pb.RefToken) (*pb.Void, error)
	GetAllUsers(req *pb.ListUserReq) (*pb.ListUserRes, error)
	GetUserById(req *pb.GetById) (*pb.UserRes, error)
}
type UserI interface {
	GetProfile(req *pb.GetByID) (*pb.UserRes, error)
	EditProfile(req *pb.UserRes) (*pb.UserRes, error)
	ChangePassword(req *pb.ChangePasswordReq) (*pb.Void, error)
	GetSetting(req *pb.GetByID) (*pb.Setting, error)
	EditSetting(req *pb.SettingReq) (*pb.Void, error)
	DeleteUser(req *pb.GetByID) (*pb.Void, error)
}

type FlashSaleI interface {
	CreateFlashSale(req *pb.CreateFlashSalesReq) (*pb.Void, error)
	UpdateFlashSale(req *pb.UpdateFlashSalesReq) (*pb.Void, error)
	ListAllFlashSales(req *pb.ListAllFlashSalesReq) (*pb.ListAllFlashSalesRes, error)
	GetFlashSale(req *pb.GetById) (*pb.FlashSale, error)
	DeleteFlashSale(req *pb.GetById) (*pb.Void, error)
	AddProductToFlashSale(req *pb.AddProductReq) (*pb.Void, error)
	RemoveProductFromFlashSale(req *pb.RemoveProductReq) (*pb.Void, error)
	CancelFlashSale(req *pb.GetById) (*pb.CancelFlashSaleRes, error)
	GetStoreLocation(req *pb.GetStoreLocationReq) (*pb.StoreLocation, error)
}
type FlashSaleProductI interface {
	CreateFlashSaleProduct(req *pb.CreateFlashSaleProductReq) (*pb.Void, error)
	UpdateFlashSaleProduct(req *pb.UpdateFlashSaleProductReq) (*pb.Void, error)
	ListAllFlashSaleProducts(req *pb.ListAllFlashSaleProductsReq) (*pb.ListAllFlashSaleProductsRes, error)
	GetFlashSaleProduct(req *pb.GetById) (*pb.FlashSaleProduct, error)
	DeleteFlashSaleProduct(req *pb.GetById) (*pb.Void, error)
}
type NotificationI interface {
	CreateNotification(req *pb.NotificationCreate) (*pb.Void, error)
	DeleteNotification(req *pb.GetById) (*pb.Void, error)
	UpdateNotification(req *pb.NotificationUpdate) (*pb.Void, error)
	GetNotifications(req *pb.NotifFilter) (*pb.NotificationList, error)
	GetNotification(req *pb.GetById) (*pb.NotificationGet, error)
}
type OrderI interface {
	CreateOrder(req *pb.CreateOrderReq) (*pb.Void, error)
	UpdateOrder(req *pb.UpdateOrderReq) (*pb.Void, error)
	ListAllOrders(req *pb.ListAllOrdersReq) (*pb.ListAllOrdersRes, error)
	GetOrder(req *pb.GetById) (*pb.Order, error)
	DeleteOrder(req *pb.GetById) (*pb.Void, error)
	GetOrderHistory(req *pb.OrderHistoryReq) (*pb.OrderHistoryRes, error)
	CancelOrder(req *pb.GetById) (*pb.CancelOrderRes, error)
}
type ProductI interface {
	CreateProduct(req *pb.CreateProductReq) (*pb.Void, error)
	UpdateProduct(req *pb.UpdateProductReq) (*pb.Void, error)
	ListAllProducts(req *pb.ListAllProductsReq) (*pb.ListAllProductsRes, error)
	GetProduct(req *pb.GetById) (*pb.Products, error)
	DeleteProduct(req *pb.GetById) (*pb.Void, error)
}

type SocialI interface {
	ShareDeal(req *pb.ShareDealReq) (*pb.Void, error)
	GetSharingStats(req *pb.GetSharingStatsReq) (*pb.SharingStatsRes, error)
}

type ReviewI interface {
	CreateReview(req *pb.CreateReviewReq) (*pb.Void, error)
	GetProductRating(req *pb.GetProductRatingReq) (*pb.ProductRatingRes, error)
}
