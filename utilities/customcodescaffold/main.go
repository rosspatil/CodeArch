package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/rosspatil/codearch/runtime/services"
	"github.com/rosspatil/codearch/runtime/services/service"
)

func main() {
	ba, err := ioutil.ReadFile("./microservice.json")
	if err != nil {
		panic(err)
	}
	ms := &services.MicroService{}
	err = json.Unmarshal(ba, ms)
	if err != nil {
		panic(err)
	}
	tmplBa, _ := ioutil.ReadFile("./tmpl")
	t := template.Must(template.New("tmpl").Parse(string(tmplBa)))
	packages := map[string]string{}
	for _, s := range ms.Services {
		for _, step := range s.Steps {
			if step.Type != service.CustomCode {
				continue
			}
			packageName := strings.ToLower(step.CustomeCode.FunctionName)
			packages[packageName] = step.CustomeCode.FunctionName
			step.CustomeCode.Data = strings.Replace(step.CustomeCode.Data, "main", packageName, 1)
			os.Mkdir("./customcode/"+packageName, os.ModePerm)
			err := ioutil.WriteFile("./customcode/"+packageName+"/"+packageName+".go", []byte(step.CustomeCode.Data), os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, map[string]interface{}{"packages": packages})
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./customcode/init.go", buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
