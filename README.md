

<h1 align="center">WAIT... MANS AT WORK... </h1>
<p>
  <div style="text-align:center"><img src="atWork.png" /></div>
</p>

<h2 align="center">API Go Bank Transfer :bank:</h2>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0.00.1-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>

</p>

- Go Bank Transfer is a simple API for some banking routines, such as creating accounts, listing accounts, listing balance for a specific account, transfers between accounts and listing transfers.

## Requirements/dependencies
- Docker
- Docker-compose

## Getting Started

- After installing Go and setting up your GOPATH, 
- [How To install Go](https://github.com/larien/aprenda-go-com-testes/blob/master/primeiros-passos-com-go/instalacao-do-go.md) 


- [Clone project](https://github.com/marcovargas74/m74-bank-api)
```sh
git clone https://github.com/marcovargas74/m74-bank-api
```

- Star Docker Myql DB 
```sh
 ## Run compiled project
 docker run --name bank-mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:8.0.26
```


- Enter in project

```sh
cd m74-bank-api/src/bank
```

- Build e RUN golang project
```sh
 ## Run compiled project
	go run main.go
```

- Build golang project

```sh
go build -o main.go
```
- Run api(port 5000)
```sh
 ## Run compiled project
	go run main.go
```



## API Request

| Endpoint        | HTTP Method           | Description       |
| --------------- | :---------------------: | :-----------------: |
| `/accounts` | `POST`                | `Create accounts` |
| `/accounts` | `GET`                 | `List accounts`   |
| `/accounts/{{account_id}}/balance`   | `GET`                |    `Find balance account` |
| `/transfers`| `POST`                | `Create transfer` |
| `/transfers`| `GET`                 | `List transfers`  |


## Test endpoints API using curl

- #### Creating new account

`Request`
```bash
curl -i --request POST 'http://localhost:5000/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Test",
    "cpf": "111.111.111-11",
    "balance": 10.00
}'
```

`Response`
```json
{
    "id":"5cf59c6c-0047-4b13-a118-65878313e329",
    "name":"Test",
    "cpf":"111.111.111-11",
    "balance":10.00,
    "created_at":"2022-01-24T10:10:02Z"
}
```
- #### Listing accounts

`Request`
```bash
curl -i --request GET 'http://localhost:5000/accounts'
```

`Response`
```json
[
    {
    "id":"5cf59c6c-0047-4b13-a118-65878313e329",
    "name":"Test",
    "cpf":"111.111.111-11",
    "balance":10.00,
    "created_at":"2022-01-24T10:10:02Z"
    }
]
```

- #### Fetching account balance

`Request`
```bash
curl -i --request GET 'http://localhost:5000/accounts/{{account_id}}/balance'
```

`Response`
```json
{
    "balance": 10.00
}
```

- #### Creating new transfer

`Request`
```bash
curl -i --request POST 'http://localhost:5000/transfers' \
--header 'Content-Type: application/json' \
--data-raw '{
	"account_origin_id": "{{account_id}}",
	"account_destination_id": "{{account_id}}",
	"amount": 100
}'
```

`Response`
```json
{
    "id": "b51cd6c7-a55c-491e-9140-91903fe66fa9",
    "account_origin_id": "{{account_id}}",
    "account_destination_id": "{{account_id}}",
    "amount": 1,
    "created_at": "2022-01-24T10:12:05Z"
}
```

- #### Listing transfers

`Request`
```bash
curl -i --request GET 'http://localhost:5000/transfers'
```

`Response`
```json
[
    {
        "id": "b51cd6c7-a55c-491e-9140-91903fe66fa9",
        "account_origin_id": "{{account_id}}",
        "account_destination_id": "{{account_id}}",
        "amount": 1,
        "created_at": "2020-11-02T14:57:35Z"
    }
]
```

## Code status
- Development

## Author
- Marco Antonio Vargas - [marcovargas74](https://github.com/marcovargas74)

## License
Copyright © 2022 [marcovargas74](https://github.com/marcovargas74).
This project is [MIT](LICENSE) licensed.
