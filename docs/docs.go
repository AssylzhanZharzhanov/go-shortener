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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/shorten": {
            "post": {
                "description": "create shorten url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Create shorten url",
                "operationId": "create-shorten-url",
                "parameters": [
                    {
                        "description": "news body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateShortenURLDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/api/shorten/{shorten_url}": {
            "get": {
                "description": "Get url by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Get url by id",
                "operationId": "get-url-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "shorten url",
                        "name": "shorten_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CreateShortenURLDTO": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string"
                }
            }
        },
        "helpers.Response": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {},
                        "status": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Swagger Mercury",
	Description:      "Documentation for Mercury APIs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
