package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cheikh2shift/testarossa"
)

func main() {

	root := flag.String("projectroot", "./", "specifies the root of the project")
	suffix := flag.String("extension", ".go", "specifies which file extension type to extract soource code from")
	file := flag.String("file", "./file.go", "specifies file to generate tests for")
	outputfile := flag.String("output", "test.go", "specifies path to save generated file to")

	flag.Parse()

	testFile, err := os.ReadFile(*file)

	if err != nil {
		log.Fatalf("error reading target file")
	}

	files, err := testarossa.GetFileContents(*root, *suffix)

	if err != nil {
		log.Fatalf("error reading contents: %v", err)
	}

	context := "Here are the current contents of the project\n"

	for _, file := range files {
		context += fmt.Sprintf(`
path: %s
Content: %s`, file.Path, file.Content)
	}

	prompt := fmt.Sprintf(
		"Write a tests, with 100%% coverage for the following code: \n%s",
		testFile,
	)

	// log.Println(prompt, context)
	apikey := os.Getenv("GEMINI_KEY")

	data, err := testarossa.GenerateTests(apikey, context+"\n"+prompt)

	if err != nil {
		log.Fatalf("error generating tests: %v", err)
	}

	// clean up output
	data = strings.Split(data, "```go")[1]
	data = strings.Split(data, "```")[0]

	if data == "" {
		log.Fatal("unrecognized output received")
	}
	// write file out
	err = os.WriteFile(*outputfile, []byte(data), os.ModePerm)

	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}

}
