// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/v1/order": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for update an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update an order",
                "parameters": [
                    {
                        "description": "order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order_service.UpdateOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create an order",
                "parameters": [
                    {
                        "description": "order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order_service.CreateOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/order/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get an order",
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
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for delete an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete an order",
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
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orderNote": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for update an order notes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-notes"
                ],
                "summary": "Update an order notes",
                "parameters": [
                    {
                        "description": "order-notes",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order_notes.UpdateOrderNotes"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating an order notes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-notes"
                ],
                "summary": "Create an order notes",
                "parameters": [
                    {
                        "description": "order-notes",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order_notes.CreateOrderNotes"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orderNote/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting an order notes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-notes"
                ],
                "summary": "Get an order notes",
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
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for delete an order notes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-notes"
                ],
                "summary": "Delete an order notes",
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
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orderNotes": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for Get all order notes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-notes"
                ],
                "summary": "Get all order notes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orderProduct": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for update an order product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-products"
                ],
                "summary": "Update an order product",
                "parameters": [
                    {
                        "description": "order-products",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order_product_service.UpdateOrderProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating an order products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-products"
                ],
                "summary": "Create an orderProduct",
                "parameters": [
                    {
                        "description": "order-products",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order_product_service.CreateOrderProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orderProduct/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-products"
                ],
                "summary": "Get an order-product",
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
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for delete an order product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-products"
                ],
                "summary": "Delete an order product",
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
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orderProducts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for Get all order products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order-products"
                ],
                "summary": "Get all orders",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orders": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for Get all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get all orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "description": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "order_notes.CreateOrderNotes": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/order_notes.PaymentEnum"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "order_notes.PaymentEnum": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5
            ],
            "x-enum-varnames": [
                "PaymentEnum_waiting_for_payment",
                "PaymentEnum_collecting",
                "PaymentEnum_shipping",
                "PaymentEnum_waiting_on_branch",
                "PaymentEnum_finished",
                "PaymentEnum_cancelled"
            ]
        },
        "order_notes.UpdateOrderNotes": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/order_notes.PaymentEnum"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "order_product_service.CreateOrderProduct": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "discount_price": {
                    "type": "number"
                },
                "order_id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "order_product_service.UpdateOrderProduct": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "discount_price": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "order_service.CreateOrder": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "courier_id": {
                    "type": "string"
                },
                "courier_name": {
                    "type": "string"
                },
                "courier_phone": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "customer_name": {
                    "type": "string"
                },
                "customer_phone": {
                    "type": "string"
                },
                "delivery_price": {
                    "type": "number"
                },
                "discount_amount": {
                    "type": "number"
                },
                "external_id": {
                    "type": "string"
                },
                "paid": {
                    "type": "boolean"
                },
                "payment_type": {
                    "$ref": "#/definitions/order_service.PaymentType"
                },
                "status": {
                    "$ref": "#/definitions/order_service.PaymentEnum"
                },
                "to_address": {
                    "type": "string"
                },
                "to_location": {
                    "$ref": "#/definitions/order_service.Location"
                },
                "type": {
                    "$ref": "#/definitions/order_service.TypeEnum"
                }
            }
        },
        "order_service.Location": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                }
            }
        },
        "order_service.PaymentEnum": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5
            ],
            "x-enum-varnames": [
                "PaymentEnum_waiting_for_payment",
                "PaymentEnum_collecting",
                "PaymentEnum_shipping",
                "PaymentEnum_waiting_on_branch",
                "PaymentEnum_finished",
                "PaymentEnum_cancelled"
            ]
        },
        "order_service.PaymentType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "PaymentType_uzum",
                "PaymentType_cash",
                "PaymentType_terminal"
            ]
        },
        "order_service.TypeEnum": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "TypeEnum_self_pickup",
                "TypeEnum_delivery"
            ]
        },
        "order_service.UpdateOrder": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "courier_id": {
                    "type": "string"
                },
                "courier_name": {
                    "type": "string"
                },
                "courier_phone": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "customer_name": {
                    "type": "string"
                },
                "customer_phone": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "delivery_price": {
                    "type": "number"
                },
                "discount_amount": {
                    "type": "number"
                },
                "external_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "paid": {
                    "type": "boolean"
                },
                "payment_type": {
                    "$ref": "#/definitions/order_service.PaymentType"
                },
                "status": {
                    "$ref": "#/definitions/order_service.PaymentEnum"
                },
                "to_address": {
                    "type": "string"
                },
                "to_location": {
                    "$ref": "#/definitions/order_service.Location"
                },
                "type": {
                    "$ref": "#/definitions/order_service.TypeEnum"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Uzum Clone Api",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
