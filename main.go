package main

import (
	"net/http"

	"github.com/Shashwat5522/golan_mongodb/controllers"
	"github.com/Shashwat5522/golan_mongodb/initializers"
	"github.com/julienschmidt/httprouter"
)

func init() {
	initializers.ConnectDB()
}
func main() {
	r := httprouter.New()
	// uc := controllers.NewUserController(getSession())
	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id",controllers.GetUserById)
	r.PUT("/user/:id",controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.GET("/search",controllers.GlobalSearch)

	http.ListenAndServe("localhost:9000", r)

}

// func connectDB() {
// 	uri := "localhost:27017"
// 	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Connected!!!")
// 	Mgr = models.Manager{Connection: client, Ctx: ctx, Cancel: cancel}

// }
