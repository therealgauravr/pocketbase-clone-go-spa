package endpoints

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/therealgauravr/pocketbase-clone-spa/runtime"
	types "github.com/therealgauravr/pocketbase-clone-spa/types"
	"gorm.io/gorm"
)

func GenerateAuthEndpoints(router *fiber.Router) {

	authRouter := (*router).Group("/auth")
	authRouter.Post("/login", login)
	authRouter.Post("/signup", signup)

}

func login(ctx *fiber.Ctx) error {
	log.Println("login request received")
	req_body := new(types.LoginPayload)
	if err := ctx.BodyParser(req_body); err != nil {
		return ctx.JSON(map[string]interface{}{
			"status":  -1,
			"message": err.Error(),
		})
	}
	log.Println(req_body)
	return ctx.JSON(map[string]interface{}{
		"status": 0,
	})
}

func signup(ctx *fiber.Ctx) error {
	log.Println("signup request received")
	signup_payload := new(types.SignupPayload)
	if err := ctx.BodyParser(signup_payload); err != nil {
		return ctx.JSON(map[string]interface{}{
			"status":  -1,
			"message": err.Error(),
		})
	}
	log.Println(signup_payload)
	gs := runtime.RetrieveGlobalStore()

	db := gs.Get("db_client").(*gorm.DB)
	newUser := types.InitUserForProvider("basic", signup_payload.User, signup_payload.Email, signup_payload.Password)
	err := db.Create(&newUser).Error
	if err != nil {
		log.Println("Error occured while creating new user... Error is " + err.Error())
	}
	return ctx.JSON(map[string]interface{}{
		"status": 0,
	})
}
