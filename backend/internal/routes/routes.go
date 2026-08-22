package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/repository"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/service"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/usecase"
	"github.com/nickznew1/MagazineMZM/backend/internal/middleware/authMiddleware"
	"github.com/nickznew1/MagazineMZM/backend/internal/middleware/logger"
	"github.com/nickznew1/MagazineMZM/backend/pkg/auth"
	"log/slog"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/internal/config"

	"github.com/go-chi/cors"
)

func Routes(sql *pgxpool.Pool, cfg *config.Config, log *slog.Logger) {
	r := chi.NewRouter()

	frontendServerUrl := cfg.ClientConfig[0].Url
	serverPort := cfg.ServerConfig[0].Port
	r.Use(middleware.RequestID)
	r.Use(logger.HTTPLogger(log))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{frontendServerUrl},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "multipart/form-data"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	userRepo := repository.NewUserRepo(sql)
	itemRepo := repository.NewItemRepo(sql)
	itemUseCase := usecase.NewItemUseCase(itemRepo)
	itemService := service.NewItemService(itemUseCase)
	cartRepo := repository.NewCartRepo(sql)
	cartUseCase := usecase.NewCartUseCase(cartRepo)
	cartService := service.NewCartService(cartUseCase)
	applicationRepo := repository.NewApplicationRepo(sql)
	applicationUseCase := usecase.NewApplicationUseCase(applicationRepo)
	applicationService := service.NewApplicationService(applicationUseCase)

	auth, err := auth.NewManager()
	if err != nil {
		return
	}
	manager := authMiddleware.NewManager(auth)

	UserUseCase := usecase.NewUserUseCase(userRepo)
	UserService := service.NewUserService(UserUseCase, auth)
	ImageFs := http.FileServer(http.Dir("./public/images"))
	r.Handle("/images/*", http.StripPrefix("/images/", ImageFs))
	PdfFs := http.FileServer(http.Dir("./public/documents"))
	r.Handle("/documents/*", http.StripPrefix("/documents/", PdfFs))

	r.Group(func(r chi.Router) {
		r.Use(manager.AuthMiddleware)
		r.Get("/profile/", UserService.GetUserProfile)
		r.Get("/user/", UserService.GetUser)
		r.Get("/cart/", cartService.GetCart)
		r.Get("/checkout", UserService.GetCheckoutInfo)
		r.Put("/applications", applicationService.CreateApplication)
		r.Get("/checkout/complete/{id}", applicationService.GetApplication)
		r.Get("/applications/all", applicationService.GetAllApplicationsForUser)
	})

	r.Route("/", func(r chi.Router) {

		r.Route("/cart", func(r chi.Router) {

			r.Post("/delete/", cartService.DeleteUserItem)
			r.Post("/add/", cartService.CreateUserItem)
			r.Post("/calc/", cartService.CalcUserItem)
		})

		r.Route("/auth", func(r chi.Router) {

			r.Post("/", UserService.UserAuth)
			r.Post("/registry", UserService.CreateUser)
		})

		r.Route("/profile", func(r chi.Router) {

			r.Post("/personal", UserService.InsertPersonalInfo)
			r.Post("/delivery", UserService.InsertDeliveryInfo)
			r.Patch("/personal/up", UserService.UpdatePersonalInfo)
			r.Patch("/delivery/up", UserService.UpdateDeliveryInfo)
			r.Patch("/", UserService.UserEmailChange)
			r.Put("/changep", UserService.UserPasswordChange)
		})

		r.Route("/item", func(r chi.Router) {

			r.Post("/create", itemService.CreateItem)
			r.Get("/{id}", itemService.GetItemById)
			r.Get("/spec/{id}", itemService.GetItemSpecById)
			r.Get("/all", itemService.GetAllItems)
			r.Delete("/delete", itemService.DeleteItem)
		})

		r.Route("/admin", func(r chi.Router) {

			r.Get("/users", UserService.GetAllUsers)
			r.Get("/applications", applicationService.GetAllApplicationsForAdmin)
			r.Post("/status", applicationService.SetApplicationStatus)
			r.Get("/application/{id}", applicationService.GetApplicationForAdmin)
			r.Post("/visible/{id}", itemService.ChangeVisible)
			r.Get("/props", itemService.GetAllPropsName)
			r.Put("/newprops/{id}", itemService.SetNewProps)

		})
	})

	err = http.ListenAndServe(serverPort, r)
	if err != nil {
		log.Error("Error when creating backend server on env.port", slog.Any("error", err))
	}
}
