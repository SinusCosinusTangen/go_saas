// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by geeks-accelerator/swag at
// 2019-06-24 13:07:10.15232 -0800 AKDT m=+0.076902855

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/geeks-accelerator/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server celler server.",
        "title": "SaaS Example API",
        "termsOfService": "http://geeksinthewoods.com/terms",
        "contact": {
            "name": "API Support",
            "url": "https://gitlab.com/geeks-accelerator/oss/saas-starter-kit",
            "email": "support@geeksinthewoods.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "example-api.saas.geeksinthewoods.com",
    "basePath": "/v1",
    "paths": {}
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
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
