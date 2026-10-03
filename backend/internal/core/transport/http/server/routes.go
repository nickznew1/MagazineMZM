package core_transport_http_server

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"
	core_middleware_auth "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/middleware/auth"
	core_middleware "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/middleware/logger"
	applications_repository_postgres "github.com/nickznew1/MagazineMZM/backend/internal/features/repository/postgres/applications"
	cart_repository_postgres "github.com/nickznew1/MagazineMZM/backend/internal/features/repository/postgres/cart"
	item_repository_postgres "github.com/nickznew1/MagazineMZM/backend/internal/features/repository/postgres/item"
	users_repository_postgres "github.com/nickznew1/MagazineMZM/backend/internal/features/repository/postgres/user"
	applications_service "github.com/nickznew1/MagazineMZM/backend/internal/features/service/applications"
	cart_service "github.com/nickznew1/MagazineMZM/backend/internal/features/service/cart"
	item_service "github.com/nickznew1/MagazineMZM/backend/internal/features/service/item"
	users_service "github.com/nickznew1/MagazineMZM/backend/internal/features/service/users"
	applications_transport_http "github.com/nickznew1/MagazineMZM/backend/internal/features/transport/http/applications"
	cart_transport_http "github.com/nickznew1/MagazineMZM/backend/internal/features/transport/http/cart"
	item_transport_http "github.com/nickznew1/MagazineMZM/backend/internal/features/transport/http/item"
	users_transport_http "github.com/nickznew1/MagazineMZM/backend/internal/features/transport/http/users"
)

func Routes(
	pool *core_postgres_pool.ConnectionPool,
	router chi.Router,
	log *slog.Logger) chi.Router {

	router.Use(middleware.RequestID)
	router.Use(core_middleware.LoggerMiddleware(log))
	router.Use(middleware.Recoverer)

	AuthManager, err := core_middleware_auth.NewManager(log)
	if err != nil {
		log.Error("error when initializing AuthManager: ", err)
		panic(err)
	}

	usersRepository := users_repository_postgres.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransport := users_transport_http.NewUsersHTTPHandler(usersService, AuthManager)

	cartRepository := cart_repository_postgres.NewCartRepository(pool)
	cartService := cart_service.NewCartService(cartRepository)
	cartTransport := cart_transport_http.NewCartHTTPHandler(cartService)

	itemRepository := item_repository_postgres.NewItemRepository(pool)
	itemService := item_service.NewItemService(itemRepository)
	itemTransport := item_transport_http.NewItemHTTPHandler(itemService)

	applicationRepository := applications_repository_postgres.NewApplicationRepository(pool)
	applicationService := applications_service.NewApplicationService(applicationRepository)
	applicationTransport := applications_transport_http.NewApplicationHTTPHandler(applicationService)

	router.Group(func(r chi.Router) {
		r.Use(AuthManager.AuthMiddleware)
		r.Get("/profile/", usersTransport.GetUserProfile)
		r.Get("/user/", usersTransport.GetUserById)
		r.Get("/cart/", cartTransport.GetCart)
		r.Get("/checkout", usersTransport.GetCheckoutInfo)
		r.Put("/applications", applicationTransport.CreateApplication)
		r.Get("/checkout/complete/{id}", applicationTransport.GetApplication)
		r.Get("/applications/all", applicationTransport.GetAllApplicationsForUser)
	})

	router.Route("/", func(r chi.Router) {
		r.Route("/cart", func(r chi.Router) {
			r.Post("/delete/", cartTransport.DeleteFromCart)
			r.Post("/add/", cartTransport.AddToCart)
			r.Post("/calc/", cartTransport.CalcItemFromCart)
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
			r.Post("/create", itemTransport.CreateItem)
			r.Get("/{id}", itemTransport.GetById)
			r.Get("/spec/{id}", itemTransport.GetSpecById)
			r.Get("/all", itemTransport.GetAll)
			r.Delete("/delete", itemTransport.DeleteItem)
		})

		r.Route("/admin", func(r chi.Router) {
			r.Get("/user", usersTransport.GetAllUsers)
			r.Get("/applications", applicationTransport.GetAllApplicationsForAdmin)
			r.Post("/status", applicationTransport.SetApplicationStatus)
			r.Get("/application/{id}", applicationTransport.GetApplicationForAdmin)
			r.Post("/visible/{id}", itemTransport.ChangeVisible)
			r.Get("/props", itemTransport.GetAllPropsName)
			r.Put("/newprops/{id}", itemTransport.SetProps)
		})
	})

	return router
}
