package main

import (
    "fmt"
    "backend_go/database"
    "backend_go/routes"
    "github.com/joho/godotenv"
    "os"
    "github.com/gin-gonic/gin"
    _ "backend_go/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "runtime"
    "os/exec"
)


func openBrowser(url string) {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "linux":
        cmd = "xdg-open"
    case "windows":
        cmd = "rundll32"
        args = append(args, "url.dll,FileProtocolHandler")
    case "darwin":
        cmd = "open"
    default:
        return
    }

    args = append(args, url)
    exec.Command(cmd, args...).Start()
}
// @title Smart Contract

// @contact.url http://www.swagger.io/support
// @host localhost:3000
// @BasePath /api/
func main() {
    err:=godotenv.Load("menu.env")
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    gin.SetMode(os.Getenv("GIN_MODE"))
    db := database.InitDB()
    
    router:=routes.SetupRouter(db)
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    fmt.Println("http://"+os.Getenv("IP_EC2")+":"+os.Getenv("PORT"))

    // go func() {
    //     time.Sleep(2 * time.Second) // Đợi server khởi động
    //     openBrowser("http://"+os.Getenv("IP_EC2")+":"+os.Getenv("PORT") + "/swagger/index.html")
    // }()
    router.Run(":"+os.Getenv("PORT"))
}