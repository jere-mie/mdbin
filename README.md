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

## Downloading Pre-Built Binaries

You can find pre-built mdbin binaries for Windows, Linux, and MacOS on the mdbin repo's releases page From there, you can download the appropriate binary and either add it to your system's PATH variable, or run it directly from whatever directory you place it in.

If you prefer downloading via the cli, use one of the following commands below:

```sh
# Windows amd64
irm -Uri https://github.com/jere-mie/mdbin/releases/latest/download/mdbin_windows_amd64.exe -O mdbin.exe

# Linux amd64
curl -L https://github.com/jere-mie/mdbin/releases/latest/download/mdbin_linux_amd64 -o mdbin && chmod +x mdbin

# MacOS arm64 (Apple Silicon)
curl -L https://github.com/jere-mie/mdbin/releases/latest/download/mdbin_darwin_arm64 -o mdbin && chmod +x mdbin

# MacOS amd64 (Intel)
curl -L https://github.com/jere-mie/mdbin/releases/latest/download/mdbin_darwin_amd64 -o mdbin && chmod +x mdbin
```
