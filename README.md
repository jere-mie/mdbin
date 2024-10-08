# mdbin
FOSS platform for quickly sharing markdown

## Getting Set Up

- Ensure you have Go v1.23.1+ installed with CGO enabled (CGO is required for sqlite)
    - If you're on Windows, check out [this tutorial](https://code.visualstudio.com/docs/languages/cpp#_example-install-mingwx64-on-windows) by Microsoft to install MinGW-x64 and enable you to use the gcc suite
    - If you're on MacOS or Linux, you probably already have a C compiler installed
- Install dependencies with the command `go mod tidy`
- Copy the contents of `example.env` to the file `.env` and change the default values to whatever you desire
- Run the application with `go run .`
- Visit the app at [localhost:3000](http://localhost:3000) (or whatever port you specified in `.env`)

### Air

You can use [Air](https://github.com/air-verse/air) for live reloading during development. Simply install Air with the following command:

```sh
go install github.com/air-verse/air@latest
```

and then you can type `air` in your terminal to run the application.