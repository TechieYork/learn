package main

import (
	"fmt"
	"os"

	"html/template"
	"strings"
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

var templateHtmlMapTest string = `
<!doctype html>
<html lang="en">
<head>
    <title>Hello, world!</title>
</head>
<body>

<table border="1">
    <tr>
        <th bgcolor="yellow">Name</th>
        <th bgcolor="yellow">Price</th>
        <th bgcolor="yellow">Status</th>
    </tr>
	{{ range $key, $value := . }}
    <tr>
		{{ if $value.IsChanged }}
        	<td bgcolor="red" width="200">{{$value.Name}}</td>
		{{ else }}
        	<td bgcolor="white" width="200">{{$value.Name}}</td>
		{{ end }}

		{{ if eq $value.Tag "test" }}
        	<td bgcolor="yellow" width="100">[Test]{{$value.Price}}</td>
		{{ else }}
        	<td bgcolor="yellow" width="100">[Online]{{$value.Price}}</td>
		{{ end }}
        <td bgcolor="#7cfc00"width="100">{{$value.Status}}</td>
    </tr>
	{{ end }}
</table>

</body>
</html>
`

var templateHtmlSliceTest string = `
<!doctype html>
<html lang="en">
<head>
    <title>Hello, world!</title>
</head>
<body>

<table border="1">
    <tr>
        <th bgcolor="yellow">Name</th>
        <th bgcolor="yellow">Price</th>
        <th bgcolor="yellow">Status</th>
    </tr>
	{{ range . }}
    <tr>
		{{ if .IsChanged }}
        	<td bgcolor="red" width="200">{{.Name}}</td>
		{{ else }}
        	<td bgcolor="white" width="200">{{.Name}}</td>
		{{ end }}

		{{ if eq .Tag "test" }}
        	<td bgcolor="yellow" width="100">[Test]{{.Price}}</td>
		{{ else }}
        	<td bgcolor="yellow" width="100">[Online]{{.Price}}</td>
		{{ end }}
        <td bgcolor="#7cfc00"width="100">{{.Status}}</td>
    </tr>
	{{ end }}
</table>

</body>
</html>
`

type Goods struct {
	Name string
	Price string
	Status string
	Tag string

	IsChanged bool
}

func main() {
	//Test template
	//testTemplate()

	//Test template html
	testTemplateHtml()
}

func testTemplate() {
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

func testTemplateHtml() {
	good := Goods{
		Name: "Good Name 1",
		Price: "9.99 -> 19.99",
		Status: "Grab-away -> In-stock",
		Tag: "test",
		IsChanged: false,
	}

	good1 := Goods{
		Name: "Good Name 2",
		Price: "9.99 -> 29.99",
		Status: "In-stock -> Grab-away",
		Tag: "online",
		IsChanged: true,
	}

	//Map test
	mapGood := make(map[string]Goods)

	mapGood[good.Name] = good
	mapGood[good1.Name] = good1

	htmlMapTemplate, err := template.New("htmlTest").Parse(templateHtmlMapTest)

	if err != nil {
		fmt.Printf("template.New failed! error:%v", err.Error())
		return
	}

	err = htmlMapTemplate.Execute(os.Stdout, mapGood)

	if err != nil {
		fmt.Printf("htmlTemplate.Execute failed! error:%v", err.Error())
		return
	}

	//Slice test
	sliceGood := make([]Goods, 0)
	sliceGood = append(sliceGood, good)
	sliceGood = append(sliceGood, good1)

	htmlSliceTemplate, err := template.New("htmlTest").Parse(templateHtmlMapTest)

	if err != nil {
		fmt.Printf("template.New failed! error:%v", err.Error())
		return
	}

	sliceStringBuilder := strings.Builder{}

	err = htmlSliceTemplate.Execute(&sliceStringBuilder, sliceGood)

	if err != nil {
		fmt.Printf("htmlTemplate.Execute failed! error:%v", err.Error())
		return
	}

	fmt.Printf(sliceStringBuilder.String())
}


