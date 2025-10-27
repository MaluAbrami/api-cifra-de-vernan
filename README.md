# Integrantes
- Maria Luiza Abrami ; Carlos Eduardo Silva

## Descrição
API que implementa a **cifra de Vernam**, permitindo criptografia e descriptografia de mensagens via endpoints REST.

---

## Pré-requisitos

Antes de rodar o projeto, é necessário ter instalado no seu computador:

1. **Go** (versão 1.20 ou superior)
   - Baixe em: https://go.dev/dl/
   - Verifique a instalação:
     ```bash
     go version
     ```

2. **Git** (opcional, se for clonar o repositório)
   - Verifique a instalação:
     ```bash
     git --version
     ```

3. Um **editor de código** (recomendado VS Code, GoLand, ou outro)

---

## Passos para rodar a API

1. Clone o repositório ou baixe o ZIP:

```bash
git clone https://github.com/SeuUsuario/api-cifra-de-vernam.git
cd api-cifra-de-vernam
```

2. Inicialize os módulos Go
```bash
go mod tidy
```

3. Execute a API:
```bash
go run main.go
```

> Obs: As dependências do Go (Gin, Swagger) serão baixadas automaticamente ao rodar `go run main.go`.  
> Se os arquivos Swagger (`docs`) já estiverem no projeto, não é necessário instalar a CLI do Swag.
- caso não tenha os 'docs' instale o swag com esse comando e depois de init, para depois sim executar a api igual é apresentado no passo 3 mais acima:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```