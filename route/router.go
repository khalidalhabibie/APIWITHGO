package route

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"../handler"
	"../middleware"
)

func RunAPI(address string) error {

	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()
	transactionHandler := handler.NewTransactionHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users",middleware.AuthorizeMiddleware())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)
		userProtectedRoutes.GET("/:user", userHandler.GetUser)
		userProtectedRoutes.GET("/:user/products", userHandler.GetProductTransaction)
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)
	}

	productProtectedRoutes := apiRoutes.Group("/products", middleware.AuthorizeMiddleware())
	{
		productProtectedRoutes.GET("/", productHandler.GetAllProduct)
		productProtectedRoutes.GET("/:product", productHandler.GetProduct)
		productProtectedRoutes.POST("/", productHandler.AddProduct)
		productProtectedRoutes.PUT("/:product", productHandler.UpdateProduct)
		productProtectedRoutes.DELETE("/:product", productHandler.DeleteProduct)
	}

	transactionProtectedRoutes := apiRoutes.Group("/order", middleware.AuthorizeMiddleware())
	{
		transactionProtectedRoutes.POST("/product/:product/quantity/:quantity", transactionHandler.TransactionProduct)
	}


	return r.Run(address)

}