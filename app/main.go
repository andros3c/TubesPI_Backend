package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"APIRestaurant/app/routes"
	userUseCase "APIRestaurant/bussiness/users"
	userController "APIRestaurant/controllers/users"
	userRepo "APIRestaurant/drivers/databases/users"

	_middleware "APIRestaurant/app/middleware"

	"APIRestaurant/drivers/mysql"
)

func init() {
	
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB){
	db.AutoMigrate(&userRepo.User{})


}

func main(){
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host: viper.GetString(`database.host`),
		DB_Port: viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDb.InitialDb()
	dbMigrate(db)
	jwt := _middleware.ConfigJWT{
		SecretJWT : viper.GetString(`jwt.secret`),
		ExpiresDuration : viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	e.Use(middleware.CORS())

	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUseCase.NewUserUseCase(userRepoInterface,timeoutContext,&jwt)
	usercontrollerInterface := userController.NewUserController(userUseCaseInterface)

	

	routesInit := routes.RouteControllerList{
	UserController: *usercontrollerInterface,
	JWTConfig	: jwt.Init(),
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}