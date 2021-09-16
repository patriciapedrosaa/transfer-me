# Transfer-me

It is an API designed to perform internal transfers between digital bank accounts.
This application allows you to: create an account, list accounts, consult the account balance, be authenticated and make transfers to other accounts.

## Based on Clean architecture
    transfer-me/
    ┣ app/
    ┃ ┣ domain/
    ┃ ┃  ┣ account/
    ┃ ┃  ┣ authentication/
    ┃ ┃  ┣ entities/
    ┃ ┃  ┣ transfer/
    ┃ ┃  ┗ vos/
    ┃ ┗ gateways/
    ┃    ┣ db/
    ┃    ┗ http/
    ┣ cmd/
    ┃  ┗ main
    ┗ Dockerfile

## Project Overview

### Endpoints

#### Create Accounts
Create an account
###### Request
- Path: `/accounts`
- Method: `POST`
- Request body
```json
  {
    "name": "gopher",
    "cpf": "12345678910",
    "secret": "mySecret"
  }
```

#### List Accounts
Get the list of accounts
###### Request
- Path: `/accounts`
- Method: `GET`

#### Get Balance
Get account balance
###### Request
- Path: `/accounts/{account_id}/balance`
- Method: `GET`

#### Login
Authenticate the user with jwt token
###### Request
- Path: `/login`
- Method: `POST`
- Request body:
```json
  {
    "cpf": "12345678910",
    "secret": "mySecret"
  }
```

###### Error message
Otherwise, you will receive the following error message:

```json
 {
   "error": "incorrect username or password"
 }
```


#### Create Transfers
Transfers from one Account to another.
###### Request
- Path: `/transfers`
- Method: `POST`
- Request body:
```json
  {
    "account_destination_id" : "7f3412f2-97cd-46de-afa5-35f72f34f3d3",
    "amount": 30
  }
```
- Header Params:
```json
  {
    "Authorization" : "Bearer Token"
  }
```

###### Error message
Otherwise, you will receive the following error message:

```json
 {
  "error": "invalid token"
 }
```

#### List Transfers
Get the list of transfers from the authenticated user.
###### Request
- Path: `http://localhost:8000/transfers`
- Method: `GET`

## Environment variables

The existing environment variables in the application are listed below followed by their descriptions:

| Name                 |  Description              |
| -------------------- | ------------------------- |
|  ACCESS_SECRET       | Sign key to generate jwt  |
|  API_PORT            | port to start application |
|  DB_USER             | database user             |
|  DB_PASSWORD         | database secret           |
|  DB_NAME             | database name             |  
|  DB_HOST             | database host             |
|  DB_PORT             | database port             |
|  DB_SSLMODE          | database SSL mode         |


## Stack

- [golang](https://golang.org/) (*1.16.5*)
- [docker](https://docs.docker.com/get-docker/)
- [testify](github.com/stretchr/testify)
- [validator](https://github.com/go-playground/validator)
- [jwt](https://github.com/dgrijalva/jwt-go)
- [migrate](https://github.com/golang-migrate/migrate)
- [uuid](https://github.com/google/uuid)
- [mux](https://github.com/gorilla/mux)
- [pgx](https://github.com/jackc/pgx)
- [zerolog](https://github.com/rs/zerolog)
- [crypto](https://pkg.go.dev/golang.org/x/crypto)

To download it with Go previously downloaded:
```bash
$ go mod download
```

## Starting (Docker required)
You can use the makefile to make it easier to build and run the application.
A typical day to day would be:
```bash
$  make postgres
```

```bash
$  make createdb
```

```bash
$  make build
```

```bash
$  make run
```

```bash
$  make test
```

## Hands On with Moq

We're using [moq](https://github.com/matryer/moq) to mock use cases.

How to use:
1) Indicate the interface you want to mock inserting a "//go:generate" instruction inside the code, like this:
```go
package account

//go:generate moq -stub -out use_case_mock.go . UseCase

type UseCase interface {
	Method1() error
	Method2(i int)
}
```

2) Run `go generate` indicating the location of the interface you want to mock, like this:
```bash
go generate ./app/domain/account/usecase.go
```

## Unit Tests
Run the command:
```bash
$  make test
```

To see the coverage run the command at the root of the project: 
```bash
$ go test -cover ./...
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Acknowledgments


Special thanks to all the authors of the libraries we use and the contents, videos and articles, on which I based this application.
Thanks also to my mentors and code reviewers.