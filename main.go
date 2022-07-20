package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

var (
	templatePath = flag.String("template_path", "example_template.txt", "Relative path to the template file")
	valuesPath   = flag.String("values_path", "values.yml", "Relative path to the values file")
	renderPath   = flag.String("render_path", "output.txt", "Relative path to the output file")
)

type Values map[string]interface{}

type Conf struct {
	Values Values `yaml:"values"`
}

func (v *Conf) getValues() *Conf {
	yamlFile, err := ioutil.ReadFile(*valuesPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return v
}

func (v Values) Add(x, y int) int {
	return x + y
}

func main() {
	flag.Parse()

	values := new(Conf)
	values.getValues()

	t, err := template.ParseFiles(*templatePath)
	if err != nil {
		log.Print(err)
		return
	}

	f, err := os.Create(*renderPath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, values.Values)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}
