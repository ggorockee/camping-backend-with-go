// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "ggorockee",
            "email": "ggorockee@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login Schema",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.LoginInputSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/presenter.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Health Check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HealthCheck"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            }
        },
        "/spot": {
            "get": {
                "description": "GetSpots",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "GetSpots",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/presenter.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Spot"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "AddSpot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "AddSpot",
                "parameters": [
                    {
                        "description": "Create Spot",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.CreateSpotInputSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/presenter.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Spot"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            }
        },
        "/spot/{id}": {
            "get": {
                "description": "GetSpot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "GetSpot",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Spot id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/presenter.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Spot"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "UpdateSpot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "UpdateSpot",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Spot id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Spot",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UpdateSpotSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/presenter.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Spot"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "RemoveSpot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "RemoveSpot",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Spot id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "PartialUpdateSpot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "PartialUpdateSpot",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Spot id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Spot",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UpdateSpotSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/presenter.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Spot"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            }
        },
        "/user/changepw": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "ChangePassword",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "ChangePassword",
                "parameters": [
                    {
                        "description": "Change Password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.ChangePasswordInputSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Create User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Register user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.SignUpInputSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/presenter.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/presenter.JsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.ChangePasswordInputSchema": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "new_password_confirm": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "entities.CreateSpotInputSchema": {
            "type": "object",
            "properties": {
                "location": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entities.LoginInputSchema": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entities.SignUpInputSchema": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_confirm": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entities.Spot": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/entities.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "entities.UpdateSpotSchema": {
            "type": "object",
            "properties": {
                "location": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "presenter.JsonResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "ggocamping App",
	Description:      "This is an API for ggocamping Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
