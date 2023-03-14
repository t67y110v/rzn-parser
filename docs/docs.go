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
        "/admin/access": {
            "post": {
                "description": "checking that user have admin rights",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "AdminAccess",
                "parameters": [
                    {
                        "description": "check adming rights",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.EmailPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/department/condition/{email}": {
            "get": {
                "description": "getting current department condition",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Department"
                ],
                "summary": "Department condition",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.DepartmentRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/email/send": {
            "post": {
                "description": "send email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Email"
                ],
                "summary": "Email",
                "parameters": [
                    {
                        "description": "handler for sending message on email",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.EmailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/parser/parse": {
            "post": {
                "description": "pars site to get informaion about nr",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parser"
                ],
                "summary": "Parser",
                "parameters": [
                    {
                        "description": "create new user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ParserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ParserResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/user/change/password": {
            "put": {
                "description": "change users passwrod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Password Change",
                "parameters": [
                    {
                        "description": "change users password",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.EmailPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "creation of user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Create",
                "parameters": [
                    {
                        "description": "create new user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "description": "delete user from system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Delete",
                "parameters": [
                    {
                        "description": "delete user from system",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.Email"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/user/session": {
            "post": {
                "description": "creation new session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Session Create",
                "parameters": [
                    {
                        "description": "create new session",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.EmailPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "put": {
                "description": "update users info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Update",
                "parameters": [
                    {
                        "description": "update users information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserUpdate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.CreateUser": {
            "type": "object",
            "properties": {
                "departments": {
                    "$ref": "#/definitions/requests.Department"
                },
                "email": {
                    "type": "string"
                },
                "monitoring_responsible": {
                    "type": "integer"
                },
                "monitoring_specialist": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "seccond_name": {
                    "type": "string"
                },
                "user_role": {
                    "type": "string"
                }
            }
        },
        "requests.Department": {
            "type": "object",
            "properties": {
                "client_department": {
                    "type": "boolean"
                },
                "db_department": {
                    "type": "boolean"
                },
                "documentation_department": {
                    "type": "boolean"
                },
                "education_department": {
                    "type": "boolean"
                },
                "international_department": {
                    "type": "boolean"
                },
                "nr_department": {
                    "type": "boolean"
                },
                "periodic_reporting_department": {
                    "type": "boolean"
                },
                "source_tracking_department": {
                    "type": "boolean"
                }
            }
        },
        "requests.Email": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "requests.EmailPassword": {
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
        "requests.EmailReq": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "recipient_mail": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        },
        "requests.ParserLogin": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.UpdateUser": {
            "type": "object",
            "properties": {
                "departments": {
                    "$ref": "#/definitions/requests.Department"
                },
                "email": {
                    "type": "string"
                },
                "monitoring_responsible": {
                    "type": "integer"
                },
                "monitoring_specialist": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "seccond_name": {
                    "type": "string"
                },
                "user_role": {
                    "type": "string"
                }
            }
        },
        "responses.CreateUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "seccond_name": {
                    "type": "string"
                }
            }
        },
        "responses.Department": {
            "type": "object",
            "properties": {
                "client_department": {
                    "type": "boolean"
                },
                "db_department": {
                    "type": "boolean"
                },
                "documentation_department": {
                    "type": "boolean"
                },
                "education_department": {
                    "type": "boolean"
                },
                "international_department": {
                    "type": "boolean"
                },
                "nr_department": {
                    "type": "boolean"
                },
                "periodic_reporting_department": {
                    "type": "boolean"
                },
                "source_tracking_department": {
                    "type": "boolean"
                }
            }
        },
        "responses.DepartmentRes": {
            "type": "object",
            "properties": {
                "departments": {
                    "$ref": "#/definitions/responses.Department"
                },
                "monitoring_responsible": {
                    "type": "integer"
                },
                "monitoring_specialist": {
                    "type": "boolean"
                }
            }
        },
        "responses.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "responses.ParserResult": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "responses.Result": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "boolean"
                }
            }
        },
        "responses.UserUpdate": {
            "type": "object",
            "properties": {
                "departments": {
                    "$ref": "#/definitions/responses.Department"
                },
                "monitoring_responsible": {
                    "type": "integer"
                },
                "monitoring_specialist": {
                    "type": "boolean"
                },
                "user_role": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.12.0",
	Host:             "localhost:4000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "PVSystem24 API",
	Description:      "Swag documentaion for PVSystem24 API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}