# Gestão Financeira - Backend

Stack: Go, Gin, PostgreSQL
Arquitetura: MVC Architecture
Público-alvo: Usuário comum

## Code Style

### Nomenclatura
- Usar **camelCase** para variáveis internas
- Usar **PascalCase** para struct, interfaces e funções exportadas
- Arquivos em **kebab-case** (ex: `user-controller.go`)
- Pacotes em **lowercase** (ex: `controllers`, `services`)
- Controllers devem terminar com `Controller` (ex: `UserController`)
- Services devem terminar com `Service` (ex: `UserService`)
- Models devem terminar com `Model` (ex: `UserModel`)

### Responsabilidades
- Controllers: lidar com requisições HTTP, validação básica, respostas
- Services: lógica de negócio, regras de validação, interação com o banco
- Models: definição de estruturas de dados, mapeamento para o banco

### API Design
1. Padrão REST
2. JSON como formato de comunicação
3. Status codes HTTP corretos

### Comunicação com Banco de Dados
- Usar `database/sql`
- Queries parametrizadas para evitar SQL injection

## Error Handling

- Tratar erros em todas as operações
- Nunca retornar erro cru do banco ao usuario
- Criar respostas padronizadas
- **Nunca retornar stacktrace**

## Comentários

Apenas em **lógica complexa** que não é auto-explicativa. Código deve ser legível por si.

## Restrições

❌ **NÃO usar**:
- SQL concatenado manualmente
- Lógica de negócio em controllers
- Controllers acessando o banco diretamente
- Variáveis globais desnecessárias

✅ **SEMPRE usar**:
- Arquitetura MVC
- Services para regras de negócio
- Repositories para acesso ao banco
- PostgreSQL como banco de dados
- Gin como framework HTTP
- JSON para comunicação

## Testing

Implementar testes apenas quando explicitamente solicitado.

## Build & Dev

```bash
go mod tidy      # instalar dependências
go run main.go   # rodar servidor
go build         # build produção
```
