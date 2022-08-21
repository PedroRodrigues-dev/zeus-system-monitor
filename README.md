# ZEUS SYSTEM MONITOR
## _Sistema gerenciador de infraestrutura_
![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)</br>
![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=LICENÇA&message=MIT&color=GREEN&style=for-the-badge)</br>
**Free Software**
## Ferramentas para desenvolvimento
- Golang 1.19
- Sqlite3

## Funções
- Monitoramento de CPU, Memória e Disco com sensores reguláveis
- Envio de dados por amqp e armazenamento em log
- Api para alteração de variáveis de ambiente em tempo real

## Desenvolvimento
```sh
go run main.go
```
# Compilação
```sh
go build main.go
```
# Manusear configurações de ambiente
Utilizando insomnia ou qualquer ferramenta de sua preferência utilize:
[GET] http://localhost:8080/configs -> para visualizar todas as configurações.
[GET] http://localhost:8080/configs/:name -> para visualizar uma configuração específica.
[PUT] http://localhost:8080/configs/:name -> para editar uma configuração específica.
PUT json body:
{
    "value": "90"
}

**Por: Pedro Rodrigues**
![Badge Github](https://img.shields.io/github/followers/PedroRodrigues-dev?style=social)
