// Code generated by swaggo/swag. DO NOT EDIT
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
        "/createUser": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/deleteUser": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "删除用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ConfigParam"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/index": {
            "get": {
                "tags": [
                    "创建用户"
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "service.ConfigParam": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "sdfdsf"
                },
                "userId": {
                    "type": "integer",
                    "example": 7
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
