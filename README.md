# Go-portunities

**Go-portunities** é uma API desenvolvida em Go para cadastro de vagas de emprego. Ela utiliza o framework **Gin-Gonic** para construção de APIs rápidas e eficientes, **Swaggo** para documentação automática da API e **SQLite** como banco de dados para armazenar as vagas.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação backend.
- **Gin-Gonic**: Framework web de alto desempenho para Go.
- **Swaggo**: Biblioteca para geração automática de documentação Swagger.
- **SQLite**: Banco de dados leve para armazenar as informações de vagas.

## Funcionalidades

- Cadastro de novas vagas de emprego.
- Listagem de vagas cadastradas.
- Busca por vagas específicas.
- Exibição de detalhes de cada vaga.

## Instalação

### Pré-requisitos

Antes de rodar o projeto, certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina.

### 1. Clone o repositório

```bash
git clone https://github.com/costamarcus/go-portunities.git
cd go-portunities
```
### 2. Make
O projeto já inclui um Makefile que facilita a execução de tarefas como a instalação de dependências e execução do servidor. Use os seguintes comandos:

Usado para rodar o servidor localmente
```bash
make run
```
Usado para rodar o servidor e atualizar a documentação do swagger
```bash
run-with-docs
```
Usado para gerar o binario da aplicação
```bash
build
```
Usado para rodar os teste unitários
```bash
test
```
Usado para gerar a documentação do swagger
```bash
docs
```

### 3. Endpoints da API
![alt text](/assets/image.png)


