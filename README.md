# Sobre esse projeto

Esse projeto é um back-end escrito em golang que tem como objetivo, simular um sistema bancário

## Como rodar esse projeto?

Você pode iniciar esse projeto facilmante usando o makefile oferecido no repositório.  
basta rodar os seguintes códigos:  

primeiramente, suba o banco de dados para o funcionamento da aplicação:

```shell
make db-up
```

Agora, faça o seeding do nosso banco de dados rodando o comando:

```shell
make seed
```

Em seguida, rode esse comando para subir a sua aplicação:  

```shell
make run
``` 

Você também pode rodar ```make test``` para verificar se os testes do proejeto passam 

Para saber os endpoints desta api de forma mais fácil, você pode consultar a página [rotas](./ENDPOINTS.md)

Pronto! Agora usua aplicação está rodando na porta 8080!
