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
Utilizando insomnia ou qualquer ferramenta de sua preferência utilize:</br>
[GET] http://localhost:8080/configs -> para visualizar todas as configurações.</br>
[GET] http://localhost:8080/configs/:name -> para visualizar uma configuração específica.</br>
[PUT] http://localhost:8080/configs/:name -> para editar uma configuração específica.</br>
PUT json body: { "value": "90" }

# Respostas
## Rabbit
A queue recebe as mensagens prontas
## Api [http://localhost:8080/configs]
Resposta JSON:
{"name":"nome_da_variavel", "value":"valor_da_variavel"}
## Websocket [ws://localhost:8080/zeus]
Resposta JSON:
{"MonitorName":"nome_do_monitor", "Percent":"porcentagem"}

# Configurações
## Variáveis de ambiente
"TimeNotificationLimit": Tempo entre as notificações</br>
"*OverloadCounterLimit": Limite que vezes o contador pode atingir a porcentagem definida no período</br>
"*PercentLimit": Limite de porcentagem para ser classificado como sobrecarga</br>
"*OverloadMessage": Pode conter qualquer valor, o único padrão e deve conter %d onde a porcentagem limite deve aparecer</br>
"RabbitHost": Host do Rabbit</br>
"RabbitPort": Porta do Rabbit</br>
"RabbitUser": Usuário do Rabbit</br>
"RabbitPassword": Senha do Rabbit</br>
"RabbitQueue": Nome da fila do Rabbit</br>
## Parâmetros de execução
Porta da api e websocket: [Padrão: 8080]
```sh
go run main.go 8081
```
ou
```sh
./zeus 8081
```


**Por: Pedro Rodrigues**
![Badge Github](https://img.shields.io/github/followers/PedroRodrigues-dev?style=social)
