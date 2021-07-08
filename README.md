# checkout-api

API HTTP que simula um e-commerce

## Quick Start (Docker Deploy)

```sh
$ docker-compose up --build
```

## Desenvolvimento (Non-Dockerized Deploy)

### 1.Clonar o código fonte

```shell
$ git clone https://github.com/murilosrg/checkout-api

$ cd checkout-api
```

### 2.Download das dependências

```shell
$ go mod download
```

### 3.Criar configuração da aplicação

```shell
$ export BLACK_FRIDAY_DATE="2021-07-07" #deve seguir exatamente esse layout de data
$ export DATABASE_FILE="products.json"
$ export DISCOUNT_URI="localhost:50051"
```

| Param     | Description                                           |
| --------- | ----------------------------------------------------- |
| BLACK_FRIDAY_DATE | Data quando ocorrerá a blackfriday (em dias de bf, um produto de presente é incluido no carrinho) |
| DATABASE_FILE | Arquivo json que contém produtos utilizados no serviço |
| DISCOUNT_URI | Caminho do serviço de desconto |

### 4.Rodando

```shell
$ go run ./cmd/server
```

**Nota**

Para utilização do serviço de desconto em desenvolvimento, é necessario a execução do comando para subir a aplicação de desconto

```shell
$ docker-compose up discount
```

### 5. Utilizando o serviço

### Request

`POST /checkout/`

    Host: localhost:8080
    Content-Type: application/json
    {
        "products": [
            {
                "id": 1,
                "quantity": 2
            },
            {
                "id": 2,
                "quantity": 1
            }
        ]
    }

### Response

    HTTP/1.1 200 Ok
    Status: 200 Ok
    Content-Type: application/json
    
    {
        "total_amount": 25000,
        "total_amount_with_discount": 23000,
        "total_discount": 2000,
        "products": [
            {
                "id": 1,
                "quantity": 2,
                "unit_amount": 10000,
                "total_amount": 20000,
                "discount": 2000,
                "is_gift": false
            },
            {
                "id": 2,
                "quantity": 1,
                "unit_amount": 5000,
                "total_amount": 5000,
                "discount": 0,
                "is_gift": false
            },
            {
                "id": 6,
                "quantity": 1,
                "unit_amount": 0,
                "total_amount": 0,
                "discount": 0,
                "is_gift": true
            }
        ]
    }

## Rodando os testes

```shell
$ go test ./...
```

## Licença

[MIT](https://opensource.org/licenses/MIT)