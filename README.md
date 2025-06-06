# API de Pizzaria

Este projeto demonstra uma API REST escrita em Go utilizando o framework [Gin](https://github.com/gin-gonic/gin). Ela permite cadastrar e consultar pizzas armazenadas em um arquivo JSON.

## Requisitos

- Go 1.20 ou superior instalado.

## Como executar

1. Clone este repositório e acesse a pasta do projeto.
2. Execute `go run main.go` para iniciar o servidor.
3. A API ficará disponível em `http://localhost:8080`.

## Endpoints

### `GET /pizzas`
Retorna a lista de pizzas cadastradas.

Exemplo de resposta:
```json
{
  "pizzas": [
    {
      "id": 1,
      "nome": "Calabresa",
      "preco": 45.5
    }
  ]
}
```

### `GET /pizzas/:id`
Busca uma pizza pelo `id` informado na rota.

- **Sucesso:** código `200` e dados da pizza.
- **Falha:** código `404` quando o `id` não for encontrado.

### `POST /pizzas`
Adiciona uma nova pizza ao catálogo. O corpo da requisição deve conter o seguinte formato:
```json
{
  "nome": "Quatro Queijos",
  "preco": 50.0
}
```

Se bem-sucedido retorna código `201` e a pizza criada.

## Estrutura do projeto
```
API-REST--GO/
├── data/            # Arquivo JSON onde as pizzas são salvas
│   └── pizza.json
├── models/          # Definição dos modelos utilizados
│   └── pizzaria.go
├── main.go          # Código principal da aplicação
└── go.mod           # Dependências do projeto
```

## Autor
Pedro Andriola

## Licença
Distribuído sob a licença MIT. Consulte o arquivo `LICENSE` para mais informações.
