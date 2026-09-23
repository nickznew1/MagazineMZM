package core_transport_http_server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/repository"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/service"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/usecase"
	users_repository_postgres "github.com/nickznew1/MagazineMZM/backend/internal/features/repository/postgres"
	users_service "github.com/nickznew1/MagazineMZM/backend/internal/features/service/users"
	users_transport_http "github.com/nickznew1/MagazineMZM/backend/internal/features/transport/http/users"
	"github.com/nickznew1/MagazineMZM/backend/internal/middleware/authMiddleware"
	"github.com/nickznew1/MagazineMZM/backend/internal/middleware/logger"
	"github.com/nickznew1/MagazineMZM/backend/pkg/auth"
)

func Routes(pool *core_postgres_pool.ConnectionPool, router chi.Router, log *slog.Logger) {

	router.Use(middleware.RequestID)
	router.Use(logger.HTTPLogger(log))
	router.Use(middleware.Recoverer)
	auth, err := auth.NewManager()
	if err != nil {
		return
	}
	manager := authMiddleware.NewManager(auth)

	usersRepository := users_repository_postgres.NewUsersRepository(pool.Pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransport := users_transport_http.NewUsersHTTPHandler(usersService, auth)
	itemRepo := repository.NewItemRepo(sql, log)
	itemUseCase := usecase.NewItemUseCase(itemRepo)
	itemService := service.NewItemService(itemUseCase)
	cartRepo := repository.NewCartRepo(sql, log)
	cartUseCase := usecase.NewCartUseCase(cartRepo)
	cartService := service.NewCartService(cartUseCase)
	applicationRepo := repository.NewApplicationRepo(sql, log)
	applicationUseCase := usecase.NewApplicationUseCase(applicationRepo)
	applicationService := service.NewApplicationService(applicationUseCase)

	ImageFs := http.FileServer(http.Dir("./public/images"))
	router.Handle("/images/*", http.StripPrefix("/images/", ImageFs))
	PdfFs := http.FileServer(http.Dir("./public/documents"))
	router.Handle("/documents/*", http.StripPrefix("/documents/", PdfFs))

	router.Group(func(r chi.Router) {
		r.Use(manager.AuthMiddleware)
		r.Get("/profile/", usersTransport.GetUserProfile)
		r.Get("/user/", usersTransport.GetUserById)
		r.Get("/cart/", cartService.GetCart)
		r.Get("/checkout", usersTransport.GetCheckoutInfo)
		r.Put("/applications", applicationService.CreateApplication)
		r.Get("/checkout/complete/{id}", applicationService.GetApplication)
		r.Get("/applications/all", applicationService.GetAllApplicationsForUser)
	})

	router.Route("/", func(r chi.Router) {
		r.Route("/cart", func(r chi.Router) {
			r.Post("/delete/", cartService.DeleteUserItem)
			r.Post("/add/", cartService.CreateUserItem)
			r.Post("/calc/", cartService.CalcUserItem)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/", usersTransport.UserAuth)
			r.Post("/registry", usersTransport.CreateUser)
		})

		r.Route("/profile", func(r chi.Router) {
			r.Post("/personal", usersTransport.InsertPersonalInfo)
			r.Post("/delivery", usersTransport.InsertDeliveryInfo)
			r.Patch("/personal/up", usersTransport.UpdatePersonalInfo)
			r.Patch("/delivery/up", usersTransport.UpdateDeliveryInfo)
			r.Patch("/", usersTransport.UserEmailChange)
			r.Put("/changep", usersTransport.UserPasswordChange)
		})

		r.Route("/item", func(r chi.Router) {
			r.Post("/create", itemService.CreateItem)
			r.Get("/{id}", itemService.GetItemById)
			r.Get("/spec/{id}", itemService.GetItemSpecById)
			r.Get("/all", itemService.GetAllItems)
			r.Delete("/delete", itemService.DeleteItem)
		})

		r.Route("/admin", func(r chi.Router) {
			r.Get("/user", usersTransport.GetAllUsers)
			r.Get("/applications", applicationService.GetAllApplicationsForAdmin)
			r.Post("/status", applicationService.SetApplicationStatus)
			r.Get("/application/{id}", applicationService.GetApplicationForAdmin)
			r.Post("/visible/{id}", itemService.ChangeVisible)
			r.Get("/props", itemService.GetAllPropsName)
			r.Put("/newprops/{id}", itemService.SetNewProps)
		})
	})

}
