### Tecnologias

- <img src="https://badges.aleen42.com/src/golang.svg" alt="badge"/> 
- <img src="https://badges.aleen42.com/src/docker.svg" alt="badge"/> 
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

### O que é ?
Implementação de uma API de fretes, para demonstrar na prática alguns conceitos sobre clean architecture.\
Para isso criei o projeto "fretes-api", usando a linguagem de programação Golang.\
Como base de dados uso o Mongo-db.

Estrutura das pastas do projeto [frete-api]:
- `adapters` Implementei a chamada HTTP;
- `drivers`  Implementei a conexão da base de dados MongoDB;
- `usecases` Implementei a regra de negócio do App;
- `entities` Implementei a regra de negócio da empresa;


### Arquitetura
![](https://github.com/lucianoortizsilva/clean-architecture/blob/main/static/clean-architecture.jpg?raw=true)

### Como rodar
- Execute na raiz do projeto o comando `docker-compose up`

- Para extrair, trasformar e carregar os dados:
Acesse a API GET: ` http://localhost:8080/v1/fretes`

```json
{
	"pedidoId": "12345",
	"regiao": "SUL",	
	"pesoTotalPedido": 100
}
```
