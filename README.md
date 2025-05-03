# Testarossa CLI
![](./rose.jpg)
A CLI that generates tests for the specified file using AI.

Demo:

![demo of product](./demo.gif)


# Requirements

- Environment variable `GEMINI_KEY` set to a valid Gemini API key. Get one [here](https://aistudio.google.com/apikey).
- Go 1.24.2 and up

# Install
Run the following command to install:

    go install github.com/cheikh2shift/testarossa/cmd/test-gen@latest


# Flags

    extension string
        specifies which file extension type to extract soource code from (default ".go")
    file string
        specifies file to generate tests for (default "./file.go")
    output string
        specifies path to save generated file to (default "test.go")
    projectroot string
        specifies the root of the project (default "./")
