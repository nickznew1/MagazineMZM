package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/repository"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/service"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/usecase"
	"github.com/nickznew1/MagazineMZM/backend/internal/middleware"
	"github.com/nickznew1/MagazineMZM/backend/pkg/auth"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/config"

	"github.com/go-chi/cors"
)

func Routes(sql *pgxpool.Pool, cfg *config.Config) {
	r := chi.NewRouter()

	frontendServerUrl := cfg.ClientConfig[0].Url
	serverPort := cfg.ServerConfig[0].Port

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
	manager := middleware.NewManager(auth)

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
		r.Post("/cart/delete/", cartService.DeleteUserItem)
		r.Post("/cart/add/", cartService.CreateUserItem)
		r.Post("/cart/calc/", cartService.CalcUserItem)
		r.Get("/", UserService.GetAllUsers)
		r.Post("/auth", UserService.UserAuth)
		r.Post("/auth/registry", UserService.CreateUser)

		r.Route("/profile", func(r chi.Router) {
			r.Post("/personal", UserService.InsertPersonalInfo)
			r.Post("/delivery", UserService.InsertDeliveryInfo)
			r.Patch("/personal/up", UserService.UpdatePersonalInfo)
			r.Patch("/delivery/up", UserService.UpdateDeliveryInfo)
			r.Patch("/", UserService.UserEmailChange)
			r.Put("/changep", UserService.UserPasswordChange)
		})
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

	log.Fatal(http.ListenAndServe(serverPort, r))
}
