// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authdto.LoginReq"
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
        "/category": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetCategoryList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "GetCategoryList",
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
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "CreateCategory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "CreateCategory",
                "parameters": [
                    {
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/categorydto.CreateCategoryReq"
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
        "/category/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetCategory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "GetCategory",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category id",
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
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "UpdateCategory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "UpdateCategory",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/categorydto.UpdateCategoryReq"
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
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "DeleteCategory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "DeleteCategory",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category id",
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
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetAllSpots",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "GetAllSpots",
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
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/spotdto.CreateSpotReq"
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
        "/spot/amenity": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetAmenities",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Amenity"
                ],
                "summary": "GetAmenities",
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
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "CreateAmenity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Amenity"
                ],
                "summary": "CreateAmenity",
                "parameters": [
                    {
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/amenitydto.CreateAmenityReq"
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
        "/spot/amenity/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetAmenity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Amenity"
                ],
                "summary": "GetAmenity",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "amenity id",
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
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "UpdateAmenity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Amenity"
                ],
                "summary": "UpdateAmenity",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "amenity id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/amenitydto.UpdateAmenityReq"
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
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "DeleteAmenity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Amenity"
                ],
                "summary": "DeleteAmenity",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "amenity id",
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
            }
        },
        "/spot/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
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
                        "description": "spot id",
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
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
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
                        "description": "spot id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/spotdto.UpdateSpotReq"
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
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
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
                        "description": "spot id",
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
            }
        },
        "/spot/{id}/review": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "AddSpotReview",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "AddSpotReview",
                "parameters": [
                    {
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reviewdto.CreateSpotReviewReq"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "spot id",
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
            }
        },
        "/spot/{id}/reviews": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "SpotReviews",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spot"
                ],
                "summary": "SpotReviews",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "spot id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
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
        "/user/change-password": {
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
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userdto.ChangePasswordReq"
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
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authdto.SignUpReq"
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
        }
    },
    "definitions": {
        "amenitydto.CreateAmenityReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "amenitydto.UpdateAmenityReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "authdto.LoginReq": {
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
        "authdto.SignUpReq": {
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
        "categorydto.CreateCategoryReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "categorydto.UpdateCategoryReq": {
            "type": "object",
            "properties": {
                "name": {
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
        },
        "reviewdto.CreateSpotReviewReq": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "spotdto.CreateSpotReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "amenities": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "category": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pet_friendly": {
                    "type": "boolean"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "spotdto.UpdateSpotReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "amenities": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "category": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pet_friendly": {
                    "type": "boolean"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "userdto.ChangePasswordReq": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
