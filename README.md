
# Pizzaria API

Esta é uma API simples para gerenciar um catálogo de pizzas usando o framework Gin em Go.

## Endpoints

### GET /pizzas

Retorna a lista de todas as pizzas.

**Resposta de Sucesso:**

- **Código:** 200
- **Conteúdo:**
  ```json
  {
    "pizzas": [
      {
        "ID": 1,
        "Name": "Margherita",
        "Ingredients": ["Tomato", "Mozzarella", "Basil"]
      },
      ...
    ]
  }
	GET /pizzas/:id
Retorna uma pizza específica pelo ID.

Parâmetros:

id: ID da pizza
Resposta de Sucesso:

Código: 200
Conteúdo:
{
  "ID": 1,
  "Name": "Margherita",
  "Ingredients": ["Tomato", "Mozzarella", "Basil"]
}
  Resposta de Erro:

Código: 404
Conteúdo:
{
  "message": "Pizza not found"
}
  POST /pizzas
Adiciona uma nova pizza ao catálogo.

Corpo da Requisição:

{
  "Name": "Pepperoni",
  "Ingredients": ["Tomato", "Mozzarella", "Pepperoni"]
}
  Resposta de Sucesso:

Código: 201
Conteúdo:
{
  "ID": 2,
  "Name": "Pepperoni",
  "Ingredients": ["Tomato", "Mozzarella", "Pepperoni"]
}
  Resposta de Erro:

Código: 400
Conteúdo:
{
  "error": "Invalid request payload"
}
  Estrutura do Projeto
  Pizzaria/
├── data/
│   └── pizza.json
├── main.go
└── models/
    └── pizza.go

	Como Executar
Certifique-se de ter o Go instalado em sua máquina.
Clone este repositório.
Navegue até o diretório do projeto.
Execute o comando go run main.go.
A API estará disponível em http://localhost:8080.
Dependências
Gin: Framework web para Go.
Autor
Pedro Andriola

Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para mais detalhes.

