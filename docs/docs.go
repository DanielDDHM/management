// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transactions/{userId}": {
            "get": {
                "description": "get transactions by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get Transactions by user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortOrder",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortOrderName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "api secret",
                        "name": "x-api-secret",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.HTTPError"
                        }
                    }
                }
            }
        },
        "/wallets/{userId}": {
            "get": {
                "description": "Wallets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallets"
                ],
                "summary": "Get Wallets",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "api secret",
                        "name": "x-api-secret",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "handlers.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
