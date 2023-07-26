package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	pg "github.com/fergusstrange/embedded-postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/keyauth"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	orm "github.com/therealgauravr/pocketbase-clone-spa/db"
	endpoints "github.com/therealgauravr/pocketbase-clone-spa/endpoints"
	runtime "github.com/therealgauravr/pocketbase-clone-spa/runtime"
	"github.com/therealgauravr/pocketbase-clone-spa/types"

	"github.com/therealgauravr/pocketbase-clone-spa/frontend"
)

var (
	gs = runtime.RetrieveGlobalStore()
)

func apiAuthFilter(c *fiber.Ctx) bool {
	protectedURLs := gs.Get("apiProtectedURLs").([]*regexp.Regexp)
	requestURL := c.OriginalURL()
	for _, pattern := range protectedURLs {
		if pattern.MatchString(requestURL) {
			return false
		}
	}
	return true

}

func basicAuthFilter(c *fiber.Ctx) bool {
	protectedURLs := gs.Get("basicProtectedURLs").([]*regexp.Regexp)
	requestURL := c.OriginalURL()
	for _, pattern := range protectedURLs {
		if pattern.MatchString(requestURL) {
			return false
		}
	}
	return true

}

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	apiKey := gs.Get("apiKey").(string)
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func main() {

	// Init the Global Store with defaults
	gs.Init()

	// Init embedded Postgres
	path, _ := filepath.Abs("./data")
	db := pg.NewDatabase(pg.DefaultConfig().DataPath(path).Logger(os.Stdout).Username("default").Password("default").Database("postgres").RuntimePath("C:\\temp").Port(8000))
	err := db.Start()

	if err != nil {
		log.Fatal("[POSTGRES INIT] DB Init failed.. Reason " + err.Error())
	}

	dbClient, err := orm.InitORM()

	// Save dbClient to Global Store

	gs.Set("db_client", dbClient)

	userResult := orm.GetUsers(dbClient)
	users := gs.Get("users").(map[string]types.User)
	for _, v := range userResult {
		users[v.User] = v
	}

	log.Printf("%+v", users)
	// orm.ShowTables(dbClient)
	// Qdrant API Key - 3IcFiUoLYn8SGWI3-elVKlU6NsIMIFJV7bCffynFcUhPqpWBeE3UEA

	// * Check dev mode. Branching paths for retrieving necessary credentials
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "DEVELOPMENT"
		err := godotenv.Load()
		if err != nil {
			log.Fatal("ENVIRONMENT COULD NOT BE LOADED")
		}
	}

	port := os.Getenv("PORT")
	// apiKey = os.Getenv("API_KEY")

	// Start the server
	app := fiber.New()

	endpointsRouter := app.Group("/")
	endpoints.GenerateAuthEndpoints(&endpointsRouter)

	// app.Use(basicauth.New(basicauth.Config{
	// 	Authorizer: func(user, pass string) bool {
	// 		if _, ok := users[user]; ok {
	// 			return users[user].IsValidPassword(pass)
	// 		}
	// 		return false
	// 	},
	// 	Next: basicAuthFilter,
	// }))
	app.Use(keyauth.New(keyauth.Config{
		KeyLookup: "cookie:access_token",
		Validator: validateAPIKey,
		Next:      apiAuthFilter,
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         frontend.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))

	app.Listen(fmt.Sprintf(":%s", port))

	defer db.Stop()
}
