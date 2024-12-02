// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/users": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "submit information to create users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "create users",
                "parameters": [
                    {
                        "description": "users information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUsersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CreateUsersReply"
                        }
                    }
                }
            }
        },
        "/api/v1/users/condition": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get users by condition",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get users by condition",
                "parameters": [
                    {
                        "description": "query condition",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Conditions"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetUsersByConditionReply"
                        }
                    }
                }
            }
        },
        "/api/v1/users/delete/ids": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "delete userss by batch id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "delete userss",
                "parameters": [
                    {
                        "description": "id array",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.DeleteUserssByIDsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.DeleteUserssByIDsReply"
                        }
                    }
                }
            }
        },
        "/api/v1/users/list": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "list of userss by last id and limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "list of userss by last id and limit",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "last id, default is MaxInt32",
                        "name": "lastID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "number per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "-id",
                        "description": "sort by column name of table, and the ",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ListUserssReply"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "list of userss by paging and conditions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "list of userss by query parameters",
                "parameters": [
                    {
                        "description": "query parameters",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Params"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ListUserssReply"
                        }
                    }
                }
            }
        },
        "/api/v1/users/list/ids": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "list of userss by batch id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "list of userss by batch id",
                "parameters": [
                    {
                        "description": "id array",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ListUserssByIDsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ListUserssByIDsReply"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get users detail by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get users detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetUsersByIDReply"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "update users information by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "update users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "users information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateUsersByIDRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.UpdateUsersByIDReply"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "delete users by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "delete users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.DeleteUsersByIDReply"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.Column": {
            "type": "object",
            "properties": {
                "exp": {
                    "description": "expressions, which default to = when the value is null, have =, !=, \u003e, \u003e=, \u003c, \u003c=, like",
                    "type": "string"
                },
                "logic": {
                    "description": "logical type, defaults to and when value is null, only \u0026(and), ||(or)",
                    "type": "string"
                },
                "name": {
                    "description": "column name",
                    "type": "string"
                },
                "value": {
                    "description": "column value"
                }
            }
        },
        "types.Conditions": {
            "type": "object",
            "properties": {
                "columns": {
                    "description": "columns info",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Column"
                    }
                }
            }
        },
        "types.CreateUsersReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "id": {
                            "description": "id",
                            "type": "integer"
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.CreateUsersRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "mobileNumber": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.DeleteUsersByIDReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data"
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.DeleteUserssByIDsReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data"
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.DeleteUserssByIDsRequest": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "id list",
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "types.GetUsersByConditionReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "users": {
                            "$ref": "#/definitions/types.UsersObjDetail"
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.GetUsersByIDReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "users": {
                            "$ref": "#/definitions/types.UsersObjDetail"
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.ListUserssByIDsReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "userss": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.UsersObjDetail"
                            }
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.ListUserssByIDsRequest": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "id list",
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "types.ListUserssReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "userss": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.UsersObjDetail"
                            }
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.Params": {
            "type": "object",
            "properties": {
                "columns": {
                    "description": "query conditions",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Column"
                    }
                },
                "limit": {
                    "description": "lines per page",
                    "type": "integer"
                },
                "page": {
                    "description": "page number, starting from page 0",
                    "type": "integer"
                },
                "sort": {
                    "description": "sorted fields, multi-column sorting separated by commas",
                    "type": "string"
                }
            }
        },
        "types.UpdateUsersByIDReply": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data"
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.UpdateUsersByIDRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "description": "uint64 id",
                    "type": "integer"
                },
                "mobileNumber": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.UsersObjDetail": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "description": "convert to uint64 id",
                    "type": "integer"
                },
                "mobileNumber": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type Bearer your-jwt-token to Value",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{"http", "https"},
	Title:            "user_service api docs",
	Description:      "http server api docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
