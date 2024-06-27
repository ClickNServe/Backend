package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	
	// authentication by google
	app.Get("/oauth/google", controllers.GoogleAuthEndpoint)
	app.Get("/oauth/redirect", controllers.GoogleRedirectEndpoint)

	// room route
	app.Get("/api/all_rooms", controllers.GetAllRoom)
	app.Get("/api/available_rooms", controllers.GetAllAvailableRoom)
	app.Get("/api/room_detail/:id", controllers.GetRoomDetail)

	// customer route
	app.Post("/api/reserve_room", controllers.ReserveRoom)
	app.Patch("/api/cancel_reservation/:id", controllers.CancelReservation)
	app.Post("/api/add_to_wishlist/:id", controllers.AddToWishlist)
	app.Delete("/api/delete_room_wishlist/:id", controllers.DropRoomWishlist)

	// admin route
	app.Post("/api/create_new_room", controllers.CreateNewRoom)
	app.Patch("/api/update_room/:id", controllers.UpdateRoom)
	app.Delete("/api/delete_room/:id", controllers.DeleteRoom)
	app.Post("/api/approve_reservation/:id", controllers.ApproveCustomerReservation)
	app.Patch("/api/reject_reservation/:id", controllers.RejectCustomerReservation)
	app.Patch("/api/handle_cancelation/:id", controllers.HandleCustomerCancelation)

}
