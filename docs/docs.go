// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "go-core",
            "url": "https://im.pltrue.top",
            "email": "pltrueover@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ByFriendRequest": {
            "post": {
                "description": "同意好友请求接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "同意好友请求接口"
                ],
                "summary": "同意好友请求接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求描述",
                        "name": "information",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求记录id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "1.同意 0 拒绝",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/CreateGroup": {
            "post": {
                "description": "创建群聊",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "创建群聊"
                ],
                "summary": "创建群聊",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "群聊名称",
                        "name": "group_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "description": "群聊用户",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/FriendList": {
            "get": {
                "description": "获取好友列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取好友列表"
                ],
                "summary": "获取好友列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/FriendPlacedTop": {
            "post": {
                "description": "好友置顶功能接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "好友置顶功能接口"
                ],
                "summary": "好友置顶功能接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "好友id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态 1.置顶 0.取消置顶",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/GetFriendForRecord": {
            "get": {
                "description": "获取好友申请记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取好友申请记录"
                ],
                "summary": "获取好友申请记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/GetGroupList": {
            "get": {
                "description": "获取群聊列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取群聊列表"
                ],
                "summary": "获取群聊列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/GetGroupMessageList": {
            "get": {
                "description": "获取群聊历史消息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取群聊历史消息"
                ],
                "summary": "获取群聊历史消息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "群聊id",
                        "name": "group_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/InformationHistory": {
            "get": {
                "description": "获取用户历史消息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取用户历史消息"
                ],
                "summary": "获取用户历史消息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "to_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "分页条数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "第几页",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/JoinGroup": {
            "post": {
                "description": "添加用户到指定群聊",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "添加用户到指定群聊"
                ],
                "summary": "添加用户到指定群聊",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "群聊id",
                        "name": "group_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/ReadMessage": {
            "get": {
                "description": "历史消息读取",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史消息读取"
                ],
                "summary": "历史消息读取[废弃]",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "图片上传",
                        "name": "voice",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/RemoveFriend": {
            "post": {
                "description": "删除好友接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "删除好友接口"
                ],
                "summary": "删除好友接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "好友id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/RemoveGroup": {
            "post": {
                "description": "删除群聊",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "删除群聊"
                ],
                "summary": "删除群聊",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "群聊id",
                        "name": "group_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/RemovedUserFromGroup": {
            "post": {
                "description": "移除群聊用户",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "移除群聊用户"
                ],
                "summary": "移除群聊用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "群聊id",
                        "name": "group_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/SendFriendRequest": {
            "post": {
                "description": "发送好友请求接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "发送好友请求接口"
                ],
                "summary": "发送好友请求",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求描述",
                        "name": "information",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "f_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "客户端类型 0.网页端登录 1.设备端登录 2.app端",
                        "name": "client_type",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/Update": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/UpdateFriendNote": {
            "post": {
                "description": "更新好友备注接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "更新好友备注接口"
                ],
                "summary": "更新好友备注接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "好友id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "note",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/UpdatePwd": {
            "post": {
                "description": "更新密码",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "更新密码"
                ],
                "summary": "更新密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "旧密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "new_password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "重复密码",
                        "name": "password_confirm",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/UploadImg": {
            "post": {
                "description": "图片上传接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片上传接口"
                ],
                "summary": "图片上传接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "图片上传",
                        "name": "Smfile",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/UploadVoiceFile": {
            "post": {
                "description": "音频文件上传接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "音频文件上传接口"
                ],
                "summary": "音频文件上传接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "图片上传",
                        "name": "voice",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/UsersList": {
            "get": {
                "description": "根据昵称查询非好友用户列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "根据昵称查询非好友用户列表"
                ],
                "summary": "根据昵称查询非好友用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/bindingEmail": {
            "post": {
                "description": "绑定用户邮箱接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "绑定用户邮箱"
                ],
                "summary": "绑定用户邮箱",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "登录接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录接口"
                ],
                "summary": "这是一个登录接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "客户端类型 0.网页端登录 1.设备端登录",
                        "name": "client_type",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/me": {
            "post": {
                "description": "获取用户信息接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取用户信息接口"
                ],
                "summary": "获取用户信息接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/registered": {
            "post": {
                "description": "注册用户接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "注册用户"
                ],
                "summary": "注册用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "password_confirm",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/seedRegisteredEmail": {
            "get": {
                "description": "发送注册邮箱验证码接口",
                "tags": [
                    "发送注册邮箱验证码接口"
                ],
                "summary": "发送注册邮箱验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
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
	Version:     "2.0",
	Host:        "114.132.40.112:9502",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "go-im  接口文档",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
