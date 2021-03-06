// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-09 03:00:17.047610189 +0800 CST m=+0.028286155

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "个人点对点支付服务",
        "title": "易支付",
        "contact": {
            "name": "halower",
            "url": "https://github.com/halower"
        },
        "license": {
            "name": "源码协议",
            "url": "https://github.com/halower/newbie-spring-boot-project/blob/master/license_996.txt"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/pay/list": {
            "get": {
                "tags": [
                    "支付接口"
                ],
                "summary": "获取所有支付信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页容",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.GetPaysListOutputDto"
                        }
                    }
                }
            }
        },
        "/api/pay/pending": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "支付接口"
                ],
                "summary": "发起支付",
                "parameters": [
                    {
                        "description": "交易信息",
                        "name": "payment_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.PaymentInfoInsertDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.PrePaidOutputDto"
                        }
                    }
                }
            }
        },
        "/api/pay/status/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "支付接口"
                ],
                "summary": "确认状态",
                "parameters": [
                    {
                        "description": "状态",
                        "name": "payment_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.TradingStatusInputDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/pay/stream": {
            "get": {
                "tags": [
                    "支付接口"
                ],
                "summary": "服务推送"
            }
        }
    },
    "definitions": {
        "models.GetPaysListOutputDto": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务状态代码",
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PaymentInfoOutputDto"
                    }
                },
                "error_detail": {
                    "description": "详细错误",
                    "type": "string",
                    "example": "详细错误"
                },
                "msg": {
                    "description": "信息",
                    "type": "string",
                    "example": "成功"
                }
            }
        },
        "models.PaymentInfoInsertDto": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "交易金额",
                    "type": "string"
                },
                "pay_type": {
                    "description": "支付方式",
                    "type": "string"
                },
                "payer_email": {
                    "description": "支付人邮箱",
                    "type": "string"
                },
                "payer_message": {
                    "description": "留言",
                    "type": "string"
                },
                "payer_mobile_number": {
                    "description": "支付人电话号码",
                    "type": "string"
                },
                "payer_name": {
                    "description": "支付人姓名",
                    "type": "string"
                }
            }
        },
        "models.PaymentInfoOutputDto": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "交易金额",
                    "type": "string"
                },
                "creation_date": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "流水号",
                    "type": "integer"
                },
                "pay_type": {
                    "description": "支付方式",
                    "type": "string"
                },
                "payee_email": {
                    "description": "收款人邮箱",
                    "type": "string"
                },
                "payee_mobile_number": {
                    "description": "收款人电话号码",
                    "type": "string"
                },
                "payee_name": {
                    "description": "收款人姓名",
                    "type": "string"
                },
                "payer_email": {
                    "description": "支付人邮箱",
                    "type": "string"
                },
                "payer_mobile_number": {
                    "description": "支付人电话号码",
                    "type": "string"
                },
                "payer_name": {
                    "description": "支付人姓名",
                    "type": "string"
                },
                "trading_status": {
                    "description": "交易状态",
                    "type": "string"
                }
            }
        },
        "models.PrePaidOutputDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "流水号",
                    "type": "integer"
                },
                "id_code": {
                    "description": "支付票据号",
                    "type": "string"
                }
            }
        },
        "models.TradingStatusInputDto": {
            "type": "object",
            "required": [
                "id",
                "trading_status"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "trading_status": {
                    "type": "string"
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
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
