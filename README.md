# teste-golang


O teste consiste em um serviço em Golang que implemente uma API Restful para controle de pessoas. Esses registros contém o nome, sexo, peso, altura e o IMC (calculado automaticamente) da pessoa. Logo após o processo, essas pessoas são armazenadas em um Cluster do MongoDB, as informações são editáveis, contudo, não é possivel adicionar o registro de dois nomes iguais. Para criar a API, foi utilizado o framework Echo. Logo após isso, foi implementado uma fila no RabbitMQ para logar uma mensagem em cada registro. 

Essa aplicação foi construída em Go utilizando testes unitários e será criado um Dockerfile para que fique disponibilizada em um container Docker.

Para rodar a aplicação, basta rodar o comando
```
docker-compose up
```
Contudo, será preciso gerar uma URI no mongoDB Cluster e inserir num arquivo local.env com o código

```
MONGODB_URI = <URI>
```
