package server

import (
	"fmt"
	"go-swagger-practice/server/gen/models"
	"go-swagger-practice/server/gen/restapi/operations"
	"math/rand"

	"github.com/go-openapi/runtime/middleware"
)

func GetGreeting(params operations.GetGreetingParams) middleware.Responder {
	greeting := fmt.Sprintf("Hello, %s!", params.Name)

	payload := &models.HelloResponse{
		Message: greeting,
	}

	err := testFunc()
	if err != nil {
		msg := err.Error()
		return operations.NewGetGreetingInternalServerError().WithPayload(&models.DefaultErrorResponse{
			Code:    INTERNAL_SERVER_ERROR,
			Message: msg,
		})
	}

	return operations.NewGetGreetingOK().WithPayload(payload)
}

func testFunc() error {
	random := rand.Intn(10)

	if random%2 == 0 {
		return fmt.Errorf("random number is even")
	}

	return nil
}
