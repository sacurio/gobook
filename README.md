# gobook
Go test project uses:

GORM => ORM framework.
MUX => URL router and dispatcher.
JWT-GO => Sing and verifies JSON Web Tokens.
GODOTENV => Load and read .env files.
GO MODULES => Manage Go dependencies.

The main goal of this project is connect to a Postgresql BDD, create and exposes some REST Services.

1. Getting started with Modules

    `go mod init`

2. Install external dependencies

    `go get github.com/jinzhu/gorm`
    `go get github.com/gorilla/mux`
    `go get github.com/dgrijalva/jwt-go`
    `go get github.com/joho/godotenv`

