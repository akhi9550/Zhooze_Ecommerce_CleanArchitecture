// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"Zhooze/pkg/api"
	"Zhooze/pkg/api/handlers"
	"Zhooze/pkg/config"
	"Zhooze/pkg/db"
	"Zhooze/pkg/repository"
	"Zhooze/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	orderRepository := repository.NewOrderRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, orderRepository)
	userHandler := handlers.NewUserHandler(userUseCase)
	productRepository := repository.NewProductRepository(gormDB)
	offerRepository := repository.NewOfferRepository(gormDB)
	productUseCase := usecase.NewProductUseCase(productRepository, offerRepository)
	productHandler := handlers.NewProductHandler(productUseCase)
	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(otpRepository)
	otpHandler := handlers.NewOtpHandler(otpUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, orderRepository, paymentRepository)
	adminHandler := handlers.NewAdminHandler(adminUseCase)
	cartRepository := repository.NewCartRepository(gormDB)
	couponRepository := repository.NewCouponRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, couponRepository, productRepository, offerRepository, orderRepository)
	cartHandler := handlers.NewCartHandler(cartUseCase, userUseCase)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository)
	orderHandler := handlers.NewOrderHandler(orderUseCase)
	couponUseCase := usecase.NewCouponUseCase(couponRepository, cartRepository, orderRepository)
	couponHandler := handlers.NewCouponHandler(couponUseCase)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository, orderRepository)
	paymentHandler := handlers.NeWPaymentHandler(paymentUseCase, orderUseCase)
	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryUseCase)
	offerUseCase := usecase.NewOfferUseCase(offerRepository)
	offerHandler := handlers.NewOfferHandler(offerUseCase)
	wishlistRepository := repository.NewWishlistRepository(gormDB)
	wishListUseCase := usecase.NewWishListUseCase(wishlistRepository, productRepository)
	wishListHandler := handlers.NewWishListHandler(wishListUseCase)
	walletRepository := repository.NewWalletRepository(gormDB)
	walletUseCase := usecase.NewWalletUseCase(walletRepository)
	walletHandler := handlers.NewWalletHandler(walletUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, productHandler, otpHandler, adminHandler, cartHandler, orderHandler, couponHandler, paymentHandler, categoryHandler, offerHandler, wishListHandler, walletHandler)
	return serverHTTP, nil
}
