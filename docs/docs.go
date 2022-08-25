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
        "contact": {
            "name": "Pedro Rodrigues",
            "email": "pedroras004@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://mit-license.org"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/configs": {
            "get": {
                "description": "Responds with the list of all configs as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configs"
                ],
                "summary": "Get configs array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Config"
                            }
                        }
                    }
                }
            }
        },
        "/configs/{name}": {
            "get": {
                "description": "Responds with one config as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configs"
                ],
                "summary": "Get one config",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search config by name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Config"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Responds with http status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configs"
                ],
                "summary": "Update a config",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update config by name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Config JSON",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Config"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Config": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v2.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Zeus System Monitor API",
	Description:      "Real-time server resource monitor and overload notification",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
