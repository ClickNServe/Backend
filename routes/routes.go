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
	app.Get("/api/all_rooms", controllers.GetAllRoom) // done postman
	app.Get("/api/available_rooms", controllers.GetAllAvailableRoom) // done postman
	app.Get("/api/room_detail/:id", controllers.GetRoomDetail) // done postman

	// bed route
	app.Get("/api/all_beds", controllers.GetAllBed) // done postman

	// facility route
	app.Get("/api/all_facilities", controllers.GetAllFacilities) // done postman

	// wishlist route
	app.Get("/api/all_wishlist", controllers.GetUserWishlist) // done postman

	// customer route
	app.Post("/api/reserve_room/:id", controllers.ReserveRoom) // done postman
	app.Post("/api/add_to_wishlist/:id", controllers.AddToWishlist) // done postman
	app.Delete("/api/delete_room_wishlist/:id", controllers.DropRoomWishlist) // done postman

	// admin route
	app.Post("/api/create_new_room", controllers.CreateNewRoom) // done postman
	app.Patch("/api/update_room/:id", controllers.UpdateRoom) // done postman
	app.Delete("/api/delete_room/:id", controllers.DeleteRoom) // done postman
	app.Post("/api/create_new_bed", controllers.CreateNewBed) // done postman
	app.Patch("/api/update_bed/:id", controllers.UpdateBed) // done postman
	app.Delete("/api/delete_bed/:id", controllers.DeleteBed) // done postman
	app.Post("/api/create_new_facility", controllers.CreateNewFacility) // done postman
	app.Patch("/api/update_facility/:id", controllers.UpdateFacility) // done postman
	app.Delete("/api/delete_facility/:id", controllers.DeleteFacility) // done postman
	app.Patch("/api/approve_reservation/:id", controllers.ApproveCustomerReservation) // done postman
	app.Patch("/api/reject_reservation/:id", controllers.RejectCustomerReservation) // done postman
	
}
