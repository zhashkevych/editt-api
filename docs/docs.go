// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-20 12:59:47.065073 +0300 EEST m=+0.044555375

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/feedback": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetFeedbacks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "GetFeedbacks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.getFeedbacksResponse"
                        }
                    },
                    "400": {},
                    "401": {},
                    "500": {}
                }
            }
        },
        "/admin/metrics": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetMetrics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "GetMetrics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MetricsAggregated"
                        }
                    },
                    "400": {},
                    "401": {},
                    "500": {}
                }
            }
        },
        "/admin/publications": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetPublications",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "GetPublications",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.getPublicationsResponse"
                        }
                    },
                    "400": {},
                    "500": {}
                }
            }
        },
        "/admin/publications/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "RemovePublication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "RemovePublication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Publication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {},
                    "401": {},
                    "500": {}
                }
            }
        },
        "/admin/sign-in": {
            "post": {
                "description": "Sign In",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Sign In",
                "parameters": [
                    {
                        "description": "Sign In Input",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.signInResponse"
                        }
                    },
                    "400": {},
                    "401": {},
                    "500": {}
                }
            }
        },
        "/api/feedback": {
            "post": {
                "description": "Create Feedback",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Create Feedback",
                "parameters": [
                    {
                        "description": "Feedback Input",
                        "name": "feedback",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Feedback"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {},
                    "500": {}
                }
            }
        },
        "/api/publications": {
            "get": {
                "description": "Gets all publications",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publications"
                ],
                "summary": "Gets all publications",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Publications type filter",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Publications count limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.getPublicationsResponse"
                        }
                    },
                    "400": {},
                    "500": {}
                }
            },
            "post": {
                "description": "Creates new publication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publications"
                ],
                "summary": "Creates new publication",
                "parameters": [
                    {
                        "description": "Create Publication",
                        "name": "publication",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.publishInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.getPublicationsResponse"
                        }
                    },
                    "400": {},
                    "500": {}
                }
            }
        },
        "/api/publications/{id}": {
            "get": {
                "description": "Gets publication by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publications"
                ],
                "summary": "Gets publication by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Publication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Publication"
                        }
                    },
                    "400": {},
                    "500": {}
                }
            }
        },
        "/api/publications/{id}/reaction": {
            "post": {
                "description": "Increments reactions count for publication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publications"
                ],
                "summary": "Increments reactions count for publication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Publication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {}
                }
            }
        },
        "/api/upload": {
            "post": {
                "description": "Upload file for publication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publications"
                ],
                "summary": "Upload file for publication",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File input",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.uploadResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.uploadResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "delivery.getFeedbacksResponse": {
            "type": "object",
            "properties": {
                "feedbacks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Feedback"
                    }
                }
            }
        },
        "delivery.getPublicationsResponse": {
            "type": "object",
            "properties": {
                "publications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Publication"
                    }
                }
            }
        },
        "delivery.signInInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "username": {
                    "type": "string",
                    "example": "editt"
                }
            }
        },
        "delivery.signInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "http.getPublicationsResponse": {
            "type": "object",
            "properties": {
                "publications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Publication"
                    }
                }
            }
        },
        "http.publishInput": {
            "type": "object",
            "required": [
                "author",
                "body",
                "imageLink",
                "tags",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "example": "Вася"
                },
                "body": {
                    "type": "string",
                    "example": "Очень крутая публикация"
                },
                "imageLink": {
                    "type": "string",
                    "example": "https://images.unsplash.com/photo-1571997804104-011c8c1d19b6?ixlib=rb-1.2.1\u0026ixid=eyJhcHBfaWQiOjEyMDd9\u0026auto=format\u0026fit=crop\u0026w=1650\u0026q=80"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "финансы",
                        "бюджет"
                    ]
                },
                "title": {
                    "type": "string",
                    "example": "Про личные финансы"
                }
            }
        },
        "http.uploadResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "ok"
                },
                "url": {
                    "type": "string",
                    "example": "https://editt-image-storage.fra1.digitaloceanspaces.com/image.png"
                }
            }
        },
        "models.Feedback": {
            "type": "object",
            "properties": {
                "features": {
                    "type": "array",
                    "items": {
                        "type": "integer",
                        "enum": [
                            1,
                            2
                        ]
                    }
                },
                "score": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "models.Metrics": {
            "type": "object",
            "properties": {
                "timestamp": {
                    "type": "string"
                },
                "unique_visitors_count": {
                    "type": "integer"
                }
            }
        },
        "models.MetricsAggregated": {
            "type": "object",
            "properties": {
                "last24": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Metrics"
                    }
                },
                "publications_count": {
                    "type": "integer"
                }
            }
        },
        "models.Publication": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string",
                    "example": "Вася"
                },
                "body": {
                    "type": "string",
                    "example": "Очень крутая публикация"
                },
                "id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "imageLink": {
                    "type": "string",
                    "example": "https://images.unsplash.com/photo-1571997804104-011c8c1d19b6?ixlib=rb-1.2.1\u0026ixid=eyJhcHBfaWQiOjEyMDd9\u0026auto=format\u0026fit=crop\u0026w=1650\u0026q=80"
                },
                "publishedAt": {
                    "type": "string"
                },
                "reactions": {
                    "type": "integer",
                    "example": 35
                },
                "readingTime": {
                    "type": "integer",
                    "example": 5
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "финансы",
                        "бюджет"
                    ]
                },
                "title": {
                    "type": "string",
                    "example": "Про личные финансы"
                },
                "views": {
                    "type": "integer",
                    "example": 586
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "editt API",
	Description: "editt back-end API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
