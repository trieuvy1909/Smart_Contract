package main

import (
    "fmt"
    "backend_go/database"
    "backend_go/routes"
    "github.com/joho/godotenv"
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    err:=godotenv.Load("menu.env")
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    gin.SetMode(os.Getenv("GIN_MODE"))
    db := database.InitDB()
    
    router:=routes.SetupRouter(db)
    fmt.Println("http://"+os.Getenv("IP_EC2")+":"+os.Getenv("PORT"))

    router.Run(":"+os.Getenv("PORT"))
}