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
        "/segments": {
            "get": {
                "description": "fetch all segments data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "find segments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create segment entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "create segment",
                "parameters": [
                    {
                        "description": "Segment slug",
                        "name": "slug",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/segments/{id}": {
            "get": {
                "description": "find segment entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "find segment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Segment ID",
                        "name": "segment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "update segment entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "update segment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Segment ID",
                        "name": "segment_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Segment slug",
                        "name": "slug",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete segment entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "delete segment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Segment ID",
                        "name": "segment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user_segments": {
            "get": {
                "description": "fetch all user_segment data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_segments"
                ],
                "summary": "find segments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create segment entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_segments"
                ],
                "summary": "create segment",
                "parameters": [
                    {
                        "description": "User id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Segment id",
                        "name": "segment_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user_segments/{id}": {
            "get": {
                "description": "find user_segment entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_segments"
                ],
                "summary": "find user_segment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "UserSegment ID",
                        "name": "user_segment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "update user_segment entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_segments"
                ],
                "summary": "update user user_segment",
                "parameters": [
                    {
                        "description": "User id",
                        "name": "user_segment_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "User id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Segment id",
                        "name": "segment_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete segment entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_segments"
                ],
                "summary": "delete segment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "UserSegment ID",
                        "name": "user_segment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "fetch all users data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "find users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create user entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "create user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "find user entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "find user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user entry by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8001",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Rest api Golang Avito App API",
	Description:      "API Server for Segments Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}