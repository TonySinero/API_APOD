// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/album/": {
            "get": {
                "description": "gets all albums from API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "getAllAlbums",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListNasa"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create new albums",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "createPerm",
                "parameters": [
                    {
                        "description": "NASA",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Nasa"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/album/filter": {
            "get": {
                "description": "gets all images with filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "getWithFilter",
                "parameters": [
                    {
                        "description": "Filter",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ApodQueryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListNasa"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/album/images": {
            "get": {
                "description": "gets all images from DB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "getImagesFromDB",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListNasa"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ApodQueryInput": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "thumbs": {
                    "type": "boolean"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.ListNasa": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Nasa"
                    }
                }
            }
        },
        "model.Nasa": {
            "type": "object",
            "properties": {
                "copyright": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "explanation": {
                    "type": "string"
                },
                "hdurl": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "media_type": {
                    "type": "string"
                },
                "service_version": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "APOD Service",
	Description:      "NASA daily imaging service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
