# Rotas da api

## Create account

**Method: Post**

**Endpoint: /accounts**

## Get all accounts

**Method: Get**

**Endpoint: /accounts**

**Query params:**

É preciso passar os seguintes parametros para que o endpoint retorne os dados requeridos:

- search: vai procurar contas com o nome sobrenome igual
- sort: asc ou desc, vai dizer como procurar as contas
- page: paginação, por default vai ser 1 caso não seja passado um valor

## Get account by Id

**Method: Get**

**Endpoint: /accounts/{id:int}**

**Header:** 
- {x-jwt-token}: token retornado ao fazer o Login

## Delete account

**Method: Delete**

**Endpoint: /accounts/{id:int}**
