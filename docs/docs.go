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
        "/": {
            "get": {
                "description": "Responds with \"Hello world\" message.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Test root",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.RootResponse"
                        }
                    }
                }
            }
        },
        "/patients": {
            "get": {
                "description": "Responds with the list of all patients as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patients"
                ],
                "summary": "Get patients array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.PatientsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Responds with the list of all patients as JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patients"
                ],
                "summary": "Add patient",
                "parameters": [
                    {
                        "description": "Add patient",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.PatientsResponse"
                        }
                    }
                }
            }
        },
        "/patients/{id}": {
            "get": {
                "description": "Responds with a patient as JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patients"
                ],
                "summary": "Get patient by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.PatientResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Responds with a patient as JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patients"
                ],
                "summary": "Toggles active status of patient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.PatientResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.HTTPError"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Responds with the list of all users as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UsersResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Responds with the list of all users as JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add User",
                "parameters": [
                    {
                        "description": "Add user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UsersResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Responds with a user as JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Responds with a user as JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Toggles active status of user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "geo": {
                    "type": "object",
                    "properties": {
                        "lat": {
                            "type": "string"
                        },
                        "lng": {
                            "type": "string"
                        }
                    }
                },
                "street": {
                    "type": "string"
                },
                "suite": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        },
        "models.Company": {
            "type": "object",
            "properties": {
                "bs": {
                    "type": "string"
                },
                "catchPhrase": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Patient": {
            "type": "object",
            "properties": {
                "BSN": {
                    "type": "string",
                    "example": "999999999"
                },
                "active": {
                    "type": "boolean",
                    "example": true
                },
                "birthdate": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "example": "a@b.io"
                },
                "familyName": {
                    "type": "string",
                    "example": "Doe"
                },
                "gender": {
                    "type": "string",
                    "example": "Female"
                },
                "givenName": {
                    "type": "string",
                    "example": "Jane"
                },
                "id": {
                    "type": "string",
                    "example": "46cb51b9-68ae-4560-943a-8ea0ae3ddc14"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "address": {
                    "$ref": "#/definitions/models.Address"
                },
                "company": {
                    "$ref": "#/definitions/models.Company"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "responses.HTTPError": {
            "type": "object",
            "properties": {
                "messasge": {
                    "type": "string",
                    "example": "bad request"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "responses.PatientResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.Patient"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "responses.PatientsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Patient"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "responses.RootResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "responses.UserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.User"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "responses.UsersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Give Me One More Shot API",
	Description:      "Give Me One More Shot service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
