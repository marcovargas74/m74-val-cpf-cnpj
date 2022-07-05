<h2 align="center">APP Go Valid CPF and CNPJ:</h2>
<p>
  <img alt="In Development" align="center" src="atWork.png" />

  <img alt="Version" src="https://img.shields.io/badge/version-0.00.1-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>

</p>

- Go Valid CPF and CNPJ is a simple CPF and CNPJ number validator. It also has a CRUD.

## Requirements/dependencies
- GO
- Docker
- Docker-compose
- MongoDB 

## Getting Started

- After installing Go and setting up your GOPATH, 
- [How To install Go](https://github.com/larien/aprenda-go-com-testes/blob/master/primeiros-passos-com-go/instalacao-do-go.md) 


- [Clone project](https://github.com/marcovargas74/m74-val-cpf-cnpj.git)
```sh
git clone https://github.com/marcovargas74/m74-val-cpf-cnpj.git
```

- HOW TO RUN   
```sh
 cd m74-val-cpf-cnpj

 ## start dockers
 make all

 ## stop dockers
 make stop
```

> :warning: **Mongo DB can take up to 3 minutes to start**: Be very careful here!


- Enter in project

```sh
cd m74-val-cpf-cnpj/src/bank
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

| Endpoint        | HTTP Method           | Description           |
| --------------- | :-------------------: | :-------------------: |
| `/cpfs`         | `POST`                | `Create CPF`          |
| `/cpfs`         | `GET`                 | `Find CPF`            |
| `/cpfs`         | `DELETE`              | `Delete CPF`          |
| `/cpfs/all`     | `GET`                 | `List CPF`            |
| `/cnpjs`        | `POST`                | `Create CPF`          |
| `/cnpjs`        | `GET`                 | `Find CPF`            |
| `/cnpjs`        | `DELETE`              | `Delete CPF`          |
| `/cnpjs/all`    | `GET`                 | `List CPF`            |
| `/status`       | `GET`                 | `Get status`          |


## Test endpoints API using curl

- #### Creating new CPF

`Request`
```bash
curl -i --request POST 'http://localhost:5000/cpfs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cpf": "111.111.111-11"
}'
```

`Response`
```json
{
    "id": "28f0d8fa-f76f-47bd-bd65-58a3c4ee9c12",
    "cpf": "111.111.111-11",
    "is_valid": true,
    "is_cpf": false,
    "is_cnpj": false,
    "created_at": "2022-07-05T12:44:29.770512516Z"
}
```
- #### Listing CPFs

`Request`
```bash
curl -i --request GET 'http://localhost:5000/cpfs/all'
```

`Response`
```json
[
    {
    "id":"5cf59c6c-0047-4b13-a118-65878313e329",
    "cpf":"111.111.111-11",
    "status":"isValid",
    "created_at":"2022-01-24T10:10:02Z"
    }
]
```

- #### Fetching CPF number is Valid

`Request`
```bash
curl -i --request GET 'http://localhost:5000/cpfs/{{cpf_number}}'
```

`Response`
```json
{
    "status":"isValid",
}
```
- #### Delete CPF Number

`Request`
```bash
curl -i --request DELETE 'http://localhost:5000/cpfs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cpf": "111.111.111-11",
}'
```

- #### Creating new CNPJ

`Request`
```bash
curl -i --request POST 'http://localhost:5000/cnpj' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cpf": "73.212.132/0001-50",
}'
```

`Response`
```json
{
    "id":"7cf59c6c-0047-4b13-a118-65878313e329",
    "cnpj":"73.212.132/0001-50",
    "status":"isValid",
    "created_at":"2022-01-24T10:10:02Z"
}
```
- #### Listing CNPJs

`Request`
```bash
curl -i --request GET 'http://localhost:5000/cnpj/all'
```

`Response`
```json
[
    {
    "id":"7cf59c6c-0047-4b13-a118-65878313e329",
    "cnpj":"73.212.132/0001-50",
    "status":"isValid",
    "created_at":"2022-01-24T10:10:02Z"
    }
]
```

- #### Fetching CNPJ number is Valid

`Request`
```bash
curl -i --request GET 'http://localhost:5000/cnpj/{{cnpj_number}}'
```

`Response`
```json
{
    "status":"isValid",
}
```
- #### Delete CNPJ Number

`Request`
```bash
curl -i --request DELETE 'http://localhost:5000/cnpj' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cnpj":"73.212.132/0001-50",
}'
```

- #### Check status

`Request`
```bash
curl -i --request GET 'http://localhost:5000/status' \
```

`Response`
```json
{
    "number_of_queries_made": 1,
    "up-time_at": "2022-01-24T10:12:05Z"
}
```


## Code status
- Development

## Next Steps
- Make a refactory
- Fix some bugs
- Add more tests

## Author
- Marco Antonio Vargas - [marcovargas74](https://github.com/marcovargas74)

## License
Copyright © 2022 [marcovargas74](https://github.com/marcovargas74).
This project is [MIT](LICENSE) licensed.
