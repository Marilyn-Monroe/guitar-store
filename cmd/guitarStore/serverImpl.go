//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=../../api/cfg.yaml ../../api/openapi.yaml

package main

import (
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"guitarStore/api"
	"guitarStore/internal/configuration"
	"guitarStore/internal/configuration/database"
	"guitarStore/internal/controller"
	repositoryImpl "guitarStore/internal/repository/impl"
	serviceImpl "guitarStore/internal/service/impl"
	"net/http"
)

type Server struct {
	cartItemController  *controller.CartItemController
	guitarController    *controller.GuitarController
	promocodeController *controller.PromocodeController
	reviewController    *controller.ReviewController
	userController      *controller.UserController
}

var _ api.ServerInterface = (*Server)(nil)

func NewServer(postgresqlMaster string, postgresqlSlaves string, dbMaxOpenConnections int, dbMaxIdleConnections int, dbConnectionMaxLifetime int, redisAddrs string) *Server {
	postgresqlCluster := database.NewPostgreSQLCluster(postgresqlMaster, postgresqlSlaves, dbMaxOpenConnections, dbMaxIdleConnections, dbConnectionMaxLifetime)
	redisCluster := configuration.NewRedisCluster(redisAddrs)

	cartItemRepository := repositoryImpl.NewCartItemRepositoryImpl(postgresqlCluster)
	guitarRepository := repositoryImpl.NewGuitarRepositoryImpl(postgresqlCluster, redisCluster)
	promocodeRepository := repositoryImpl.NewPromocodeRepositoryImpl(postgresqlCluster, redisCluster)
	reviewRepository := repositoryImpl.NewReviewRepositoryImpl(postgresqlCluster, redisCluster)
	userRepository := repositoryImpl.NewUserRepositoryImpl(postgresqlCluster, redisCluster)

	cartItemService := serviceImpl.NewCartItemServiceImpl(&cartItemRepository)
	guitarService := serviceImpl.NewGuitarServiceImpl(&guitarRepository, &reviewRepository)
	promocodeService := serviceImpl.NewPromocodeServiceImpl(&promocodeRepository)
	reviewService := serviceImpl.NewReviewServiceImpl(&reviewRepository)
	userService := serviceImpl.NewUserServiceImpl(&userRepository)

	cartItemController := controller.NewCartItemController(&cartItemService)
	guitarController := controller.NewGuitarController(&guitarService)
	promocodeController := controller.NewPromocodeController(&promocodeService)
	reviewController := controller.NewReviewController(&reviewService)
	userController := controller.NewUserController(&userService)

	return &Server{
		cartItemController:  cartItemController,
		guitarController:    guitarController,
		promocodeController: promocodeController,
		reviewController:    reviewController,
		userController:      userController,
	}
}

func (s Server) GetCartItemsByUserId(w http.ResponseWriter, r *http.Request, params api.GetCartItemsByUserIdParams) {
	s.cartItemController.GetCartItemsByUserId(w, r, params)
}

func (s Server) EditCartItem(w http.ResponseWriter, r *http.Request) {
	s.cartItemController.EditCartItem(w, r)
}

func (s Server) FindAllGuitars(w http.ResponseWriter, r *http.Request, params api.FindAllGuitarsParams) {
	s.guitarController.FindAllGuitars(w, r, params)
}

func (s Server) GetGuitarById(w http.ResponseWriter, r *http.Request, guitarId openapiTypes.UUID) {
	s.guitarController.GetGuitarById(w, r, guitarId)
}

func (s Server) GetPromocodeByCode(w http.ResponseWriter, r *http.Request, params api.GetPromocodeByCodeParams) {
	s.promocodeController.GetPromocodeByCode(w, r, params)
}

func (s Server) FindReviewsByGuitarId(w http.ResponseWriter, r *http.Request, params api.FindReviewsByGuitarIdParams) {
	s.reviewController.FindReviewsByGuitarId(w, r, params)
}

func (s Server) CreateReview(w http.ResponseWriter, r *http.Request) {
	s.reviewController.CreateReview(w, r)
}
