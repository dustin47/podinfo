// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-16 09:49:23.920068 +0300 EEST m=+0.052088436

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
        "contact": {
            "name": "Source Code",
            "url": "https://github.com/stefanprodan/podinfo"
        },
        "license": {
            "name": "MIT License",
            "url": "https://github.com/stefanprodan/podinfo/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "renders podinfo UI",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/echo": {
            "post": {
                "description": "forwards the call to the backend service and echos the posted content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Echo",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/api/info": {
            "get": {
                "description": "returns the runtime information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Runtime information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RuntimeResponse"
                        }
                    }
                }
            }
        },
        "/cache": {
            "post": {
                "description": "writes the posted content in cache and returns the SHA1 hash of the content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Save payload in cache",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/cache/{hash}": {
            "get": {
                "description": "returns the content from cache if key exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Get payload from cache",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/chunked/{seconds}": {
            "get": {
                "description": "uses transfer-encoding type chunked to give a partial response and then waits for the specified period",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Chunked transfer encoding",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/delay/{seconds}": {
            "get": {
                "description": "waits for the specified period",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Delay",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/env": {
            "get": {
                "description": "returns the environment variables as a JSON array",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Environment",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ArrayResponse"
                        }
                    }
                }
            }
        },
        "/headers": {
            "get": {
                "description": "returns a JSON array with the request HTTP headers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Headers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ArrayResponse"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "used by Kubernetes liveness probe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Liveness check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "description": "returns HTTP requests duration and Go runtime metrics",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Prometheus metrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/panic": {
            "get": {
                "description": "crashes the process with exit code 255",
                "tags": [
                    "HTTP API"
                ],
                "summary": "Panic"
            }
        },
        "/readyz": {
            "get": {
                "description": "used by Kubernetes readiness probe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Readiness check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/readyz/disable": {
            "post": {
                "description": "signals the Kubernetes LB to stop sending requests to this instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Disable ready state",
                "responses": {
                    "202": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/readyz/enable": {
            "post": {
                "description": "signals the Kubernetes LB that this instance is ready to receive traffic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Enable ready state",
                "responses": {
                    "202": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/status/{code}": {
            "get": {
                "description": "sets the response status code to the specified code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Status code",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/store": {
            "post": {
                "description": "writes the posted content to disk at /data/hash and returns the SHA1 hash of the content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Upload file",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/store/{hash}": {
            "get": {
                "description": "returns the content of the file /data/hash if exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Download file",
                "responses": {
                    "200": {
                        "description": "file",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "description": "issues a JWT token valid for one minute",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Generate JWT token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.TokenResponse"
                        }
                    }
                }
            }
        },
        "/token/validate": {
            "post": {
                "description": "validates the JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Validate JWT token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.TokenValidationResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "returns podinfo version and git commit hash",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/ws/echo": {
            "post": {
                "description": "echos content via websockets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Echo over websockets",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ArrayResponse": {
            "type": "array",
            "items": {
                "type": "string"
            }
        },
        "api.MapResponse": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "api.RuntimeResponse": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "goarch": {
                    "type": "string"
                },
                "goos": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "num_cpu": {
                    "type": "string"
                },
                "num_goroutine": {
                    "type": "string"
                },
                "revision": {
                    "type": "string"
                },
                "runtime": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "api.TokenResponse": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api.TokenValidationResponse": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "token_name": {
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
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "localhost:9898",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "Podinfo API",
	Description: "Go microservice template for Kubernetes.",
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
