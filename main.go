package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"text/template"
)


type Ctest struct {
	Name string `yaml:"name"`
	GlobalSetup []string 	`yaml:"globalSetup"`
	GlobalTearDown []string	`yaml:"globalTearDown"`
	Setup []string			`yaml:"setup"`
	TearDown[] string		`yaml:"tearDown"`
	ExitOnFail bool 	    `yaml:"exitOnFail"`
	Tests []struct {
		Name string `yaml:"name"`
		Steps []struct {
			Name string `yaml:"name"`
			Command string `yaml:"command"`
			RetCode *int `yaml:"retCode"`
			Output string `yaml:"output"`
			OutputExp string `yaml:"outputExp"`
			Echo bool `yaml:"echo"`
		} `yaml:"steps"`
	} `yaml:"tests"`
}


func checkTest(ctest Ctest) string {
	if ctest.Name == "" {
		return "Missing 'name' parameter"
	}
	return ""
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: ctest <testfile>\n")
		os.Exit(1)
	}

	fileName := os.Args[1]
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading '%s' as YAML file: %s\n", fileName, err)
		os.Exit(2)
	}

	var ctest Ctest
	err = yaml.UnmarshalStrict(yamlFile, &ctest)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		os.Exit(3)
	}

	mess := checkTest(ctest); if mess != "" {
		fmt.Printf("%s in '%s' test file\n", mess, fileName)
		os.Exit(5)
	}


	tmpl, err := template.New("template1").Parse(template1)
	if err != nil {
		fmt.Printf("Error loading template: %s\n", err)
		os.Exit(4)
	}
	err = tmpl.Execute(os.Stdout, ctest)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err)
		os.Exit(4)
	}


}
