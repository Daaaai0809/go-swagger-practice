// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "go-swaggerで生成したAPIのサンプル",
    "title": "go-swgger-practice",
    "version": "1.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/hello": {
      "$ref": "./paths/hello.yaml"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "go-swaggerで生成したAPIのサンプル",
    "title": "go-swgger-practice",
    "version": "1.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/hello": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "getGreeting",
        "parameters": [
          {
            "type": "string",
            "description": "名前",
            "name": "name",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/helloResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/defaultErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "defaultErrorResponse": {
      "description": "Default error response",
      "properties": {
        "code": {
          "type": "integer",
          "example": 400
        },
        "message": {
          "type": "string",
          "example": "Bad Request"
        }
      }
    },
    "helloResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "example": "Hello, World!"
        }
      }
    }
  }
}`))
}
