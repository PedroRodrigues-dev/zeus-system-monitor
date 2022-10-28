# ZEUS SYSTEM MONITOR
## _Sistema gerenciador de infraestrutura_
![Badge Finalizado](http://img.shields.io/static/v1?label=STATUS&message=FINALIZADO&color=GREEN&style=for-the-badge)</br>
![Badge Licença](http://img.shields.io/static/v1?label=LICENÇA&message=MIT&color=GREEN&style=for-the-badge)</br>
**Free Software**
## Ferramentas para desenvolvimento
- Golang 1.19
- Sqlite3
## Funções
- Monitoramento de CPU, Memória e Disco com sensores reguláveis
- Envio de dados por websocket
- Interface web para monitoramento
- Api com swagger para alteração de variáveis de ambiente
## Desenvolvimento
```sh
go run main.go
```
## Compilação
```sh
go build main.go
```
# Interface [http://localhost:8080]
Por meio do Zeus Simple Interface você pode ver os dados em tempo real do monitoramento e receber notificações desktop de sobrecarga. 
# Manusear configurações de ambiente
## Swagger [http://localhost:8080/docs/index.html]
/configs -> Visualizar todas as configurações.</br>
/configs/{name} -> Visualizar uma configuração específica.</br>
/configs/{name} -> Editar uma configuração específica.</br>
PUT json body: { "value": "90" }
## Respostas
### Api [http://localhost:8080/configs]
Resposta JSON:
{"name":"nome_da_variavel", "value":"valor_da_variavel"}
### Websocket 
#### Resposta tempo real JSON: [ws://localhost:8080/monitor]
{"MonitorName":"nome_do_monitor", "Percent":"porcentagem", "IsOverload":"booleano_para_sobrecarga", "Alert":"mensagem_de_alerta"}
## Configurações
### Variáveis de ambiente
"*OverloadCounterLimit": Limite que vezes o contador pode atingir a porcentagem definida no período</br>
"*PercentLimit": Limite de porcentagem para ser classificado como sobrecarga</br>
"*OverloadMessage": Pode conter qualquer valor, o único padrão e deve conter %d onde a porcentagem limite deve aparecer</br>
### Parâmetros de execução
Porta da api e websocket: [Padrão: 8080]
```sh
go run main.go 8081
```
ou
```sh
./zeus 8081
```
**Interface Project: Julia Rodrigues Almeida**<br/>
**Por: Pedro Rodrigues**
![Badge Github](https://img.shields.io/github/followers/PedroRodrigues-dev?style=social)
