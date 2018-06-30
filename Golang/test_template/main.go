package main

import (
	"fmt"
	"os"

	"text/template"
)

var templateTest string = `
func main() {
	Init{{.LogName}}
	Init{{.MonitorName}}
}
`

type Application struct {
	LogName string
	MonitorName string
}

func main() {
	var app Application

	app.LogName = "Seelog"
	app.MonitorName = "Statsd"

	appTemplate, err := template.New("application").Parse(templateTest)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	appTemplate.Execute(os.Stdout, app)
}
