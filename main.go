package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	yaml "gopkg.in/yaml.v2"
)

var envs []string

// Recursively extact the yaml data structure into flattened environment vars
func makeEnvs(m interface{}, prefix string) {
	mm := reflect.ValueOf(m)
	mk := mm.Type().Kind()
	if mk == reflect.Struct {
		mm = reflect.ValueOf(structs.Map(m))
	}
	for _, key := range mm.MapKeys() {
		var k string
		k = fmt.Sprintf("%v", key)
		v := m.(map[interface{}]interface{})[k]
		tv := reflect.ValueOf(v).Kind()
		newKey := prefix + "_" + k

		if tv == reflect.Map || tv == reflect.Struct {
			makeEnvs(v, newKey)
		} else if tv == reflect.Array || tv == reflect.Slice {
			for i, e := range v.([]interface{}) {
				ke := fmt.Sprintf("%s_%d", newKey, i)
				envs = append(envs, strings.Replace(fmt.Sprintf("%s=%v", ke, e), "\n", "\\n", -1))
			}
		} else {
			envs = append(envs, strings.Replace(fmt.Sprintf("%s=%v", newKey, v), "\n", "\\n", -1))
		}
	}
}

func main() {
	yamlfile := flag.String("yaml", "", "yaml file to load environment from")
	flag.Parse()

	envs = os.Environ()

	yamlfileExp := os.Getenv("YAMLSH_YAMLFILE")

	// --yaml option takes precedence
	if *yamlfile != "" {
		yamlfileExp = os.ExpandEnv(*yamlfile)
	}

	prefix := "YAMLSH"
	yamlshPrefix := os.Getenv("YAMLSH_PREFIX")
	if yamlshPrefix != "" {
		prefix = yamlshPrefix
	}

	shell := "/bin/bash"
	yamlshshell := os.Getenv("YAMLSH_SHELL")
	if yamlshshell != "" {
		shell = yamlshshell
	}

	if yamlfileExp != "" {
		y, err := ioutil.ReadFile(yamlfileExp)
		if err != nil {
			panic(err)
		}

		// parse yaml into t
		t := make(map[interface{}]interface{})
		err = yaml.Unmarshal(y, &t)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// Convert to flattened env vars for shell
		makeEnvs(t, prefix)
	}

	cmd := exec.Command(shell, flag.Args()...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = envs

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}
