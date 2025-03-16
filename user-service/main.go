package main

import (
	"context"
	"fmt"

	"github.com/Kodnavis/face2face-backend/user-service/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
