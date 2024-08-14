# Go-MongoDB-Aggregation
Projeto REST desenvolvido em Go com banco de dados MongoDB rodando em um conteiner Docker

### Atualização Go:
1) Rodar no terminal o seguinte comando, para fazer atualizar e baixar bibliotecas necessárias:<br>
**go mod tidy**

### Docker:
1) Inicialmente rodal o container docker no terminal:<br>
**docker-compose up -d**

2) Checar se o conteiner está rodando normalmente atraves do comando no terminal:<br>
**docker ps**

3) Se quiser parar (stop) o container, no terminal:<br>
**docker stop**

### Aplicação Go:
1) Para rodar a aplicação Go, digite no terminal:<br>
**go run main.go**

2) Parar (stop) a aplicação, no terminal:<br>
**ctrl + c**

### Postman - Para trabalhar com as rotas/CRUD:<br>
1) POST:<br>
**http://localhost:8080/sales**<br>
**Body - raw - JSON**<br>

{
    "product": "Ipad",<br>
    "category": "Electronics",
    "amount": 3000.00,
    "date": "2024-07-22"
}

2) PUT:<br>
**http://localhost:8080/sales/Ipad**<br>
**Body - raw - JSON**<br>
{
    "product": "Ipad",
    "category": "Electronics",
    "amount": 2700.00,
    "date": "2024-07-22"
}

**OBS: product é a chave na estrutura chave-valor, portanto não pode ser modificado. Caso precise alterar, delete e cadastre um novo produto**

3) DELETE:<br>
**http://localhost:8080/sales/Ipad**<br>
SEND

4) GET: Agrupamento por Category<br>
**http://localhost:8080/sales/aggregate?category=Electronics**<br>
**Params - Key: category - Value: Electronics**<br>
**OBS: Tras o valor total de uma categoria**

5) GET: Agrupamento por Dates (ini e fim)<br>
**http://localhost:8080/sales/aggregateByDate?startDate=2024-08-01&endDate=2024-08-31**<br>
**Params - key: startDate e endDate - Value: 2024-08-01 e 2024-08-31**<br>
**OBS: Traz o valor total agrupado vendido no periodo**

6) GET: Lista de todos os Products<br>
**http://localhost:8080/sales**<br>
SEND


