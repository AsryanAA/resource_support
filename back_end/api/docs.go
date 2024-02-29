// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/accessor": {
            "get": {
                "description": "Возвращает массив всех accessors",
                "tags": [
                    "Фурнитура (accessors)"
                ],
                "summary": "Возвращает список Фурнитура",
                "parameters": [
                    {
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "PasswordWeb",
                        "name": "password_web",
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
                            "$ref": "#/definitions/models.Accessor"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Фурнитура (accessors)"
                ],
                "summary": "Создание новой записи Фурнитура",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PasswordWeb",
                        "name": "password_web",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Accessor"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "tags": [
                    "Фурнитура (accessors)"
                ],
                "summary": "Обновление записи Фурнитура",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PasswordWeb",
                        "name": "password_web",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Accessor"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            }
        },
        "/accessor/:id": {
            "delete": {
                "tags": [
                    "Фурнитура (accessors)"
                ],
                "summary": "Удаление записи Фурнитура",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PasswordWeb",
                        "name": "password_web",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Accessor"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Возвращает данные пользователя",
                "tags": [
                    "User (auth)"
                ],
                "summary": "Возвращает пользователя",
                "parameters": [
                    {
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "PasswordWeb",
                        "name": "password_web",
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
                            "$ref": "#/definitions/mw.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            }
        },
        "/calculation": {
            "get": {
                "description": "Возвращает рассчет потребности по лицевой карточке",
                "tags": [
                    "Calculations"
                ],
                "summary": "Возвращает рассчет потребности",
                "responses": {}
            }
        },
        "/munition": {
            "get": {
                "description": "Возвращает массив всех munitions",
                "tags": [
                    "Имущество (munitions)"
                ],
                "summary": "Возвращает список Имущество",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PasswordWeb",
                        "name": "password_web",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Munition"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Имущество (munitions)"
                ],
                "summary": "Создание новой записи Имущество",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PasswordWeb",
                        "name": "password_web",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Munition"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "tags": [
                    "Имущество (munitions)"
                ],
                "summary": "Обновление записи Имущество",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PasswordWeb",
                        "name": "password_web",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Munition"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            }
        },
        "/munition/:id": {
            "delete": {
                "tags": [
                    "Имущество (munitions)"
                ],
                "summary": "Удаление записи Имущество",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthId",
                        "name": "auth_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PasswordWeb",
                        "name": "password_web",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Munition"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/mw.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Accessor": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "crn": {
                    "type": "string"
                },
                "dicnomns": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nommodif": {
                    "type": "string"
                },
                "rn": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.Munition": {
            "type": "object",
            "properties": {
                "begstepheihth": {
                    "type": "string"
                },
                "cloth": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "crn": {
                    "type": "string"
                },
                "dicnomns": {
                    "type": "string"
                },
                "ignoresupply": {
                    "type": "string"
                },
                "kompsum": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nommodif": {
                    "type": "string"
                },
                "packmodif": {
                    "type": "string"
                },
                "rn": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "usefaktheihth": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "mw.HTTPError": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                }
            }
        },
        "mw.User": {
            "type": "object",
            "properties": {
                "license": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "user_data": {
                    "$ref": "#/definitions/mw.UserData"
                }
            }
        },
        "mw.UserData": {
            "type": "object",
            "properties": {
                "auth_id": {
                    "type": "string"
                },
                "password_web": {
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
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
