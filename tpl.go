package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/Masterminds/sprig"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var (
	templateFiles = flag.StringArrayP("template", "t", []string{}, "template files")
	dataFile      = flag.StringP("data", "d", "-", "input file")
	outputFile    = flag.StringP("output", "o", "-", "output file")
)

// ExecuteTemplates Reads a YAML document from the valuesIn stream, uses it as values
// for the tplFiles templates and writes the executed templates to
// the out stream.
func ExecuteTemplates(valuesIn io.Reader, out io.Writer, tplFiles []string) error {
	name := path.Base(tplFiles[0])
	tpl, err := template.New(name).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFiles...)
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err)
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, valuesIn)
	if err != nil {
		return fmt.Errorf("Failed to parse input template: %v", err)
	}

	var values map[string]interface{}
	err = yaml.Unmarshal(buf.Bytes(), &values)
	if err != nil {
		return fmt.Errorf("Failed to parse input template: %v", err)
	}

	err = tpl.Execute(out, values)
	if err != nil {
		return fmt.Errorf("Failed to parse input template: %v", err)
	}
	return nil
}

func openFile(path string, minusDefault *os.File) (*os.File, error) {
	if path == "-" {
		return minusDefault, nil
	}
	switch minusDefault {
	case os.Stdin:
		return os.Open(path)
	case os.Stdout:
		return os.Create(path)
	}
	return &os.File{}, nil
}

func main() {
	flag.Parse()

	data, err := openFile(*dataFile, os.Stdin)
	if err != nil {
		log.Fatalf("Error opening dataFile '%s'", *dataFile)
		os.Exit(1)
	}

	output, err := openFile(*outputFile, os.Stdout)
	if err != nil {
		log.Fatalf("Error opening outputFile '%s'", *outputFile)
		os.Exit(1)
	}

	if len(*templateFiles) == 0 {
		log.Fatal("No template files specified")
	}

	err = ExecuteTemplates(data, output, *templateFiles)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
