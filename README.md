# Project for learning go

## Tech Stack

- `fiber` https://docs.gofiber.io/
- `gorm` https://gorm.io/index.html
- `MySQL` Database

## Environment Variable

You can see all Environment Variable in `.env.example` file

###

To run project with `.env` file. added `local=development` in machine that run this project.

you can write this command in mac

```bash
export TYPE=development
```

Maybe in the future will implement tools like `air`(hot reload) and `docker` for easy to development as not need to setup environment variable in development machine. If have time.

## Run Project

You can run project buy

```bash
go run ./cmd/main.go
```

for now I always use nodemon to hot reload my project

```bash
nodemon --exec go run ./cmd/main.go
```

You can install in global to not install nodemon in this project!

```bash
npm i -g nodemon
```

`PS:I'm don't know is it best practice. but I hate when need to cmd+c or ctrl+c when work with go project to stop server and and run again `
