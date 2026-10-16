#📘 DevBook
Uma mini rede social desenvolvida em Go (Golang), com API REST, frontend HTML + JavaScript, e integração com MySQL.
Projeto criado para fins de estudo durante o curso Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA! ministrado por Otávio Augusto Gallego na Udemy.

#📌 Sobre o projeto
O DevBook é uma aplicação completa que combina:
- Backend em Go com API REST
- Frontend em HTML, CSS, Bootstrap e JavaScript (com jQuery)
- Banco de dados MySQL
- Autenticação via cookies
- CRUD completo de usuários e publicações
- Funcionalidades típicas de uma rede social
- O objetivo do projeto é proporcionar uma experiência prática no desenvolvimento de aplicações web utilizando Go, desde a criação da API até o consumo dos endpoints no frontend.

#🧩 Funcionalidades
##👤 Usuários
Cadastro de usuários

Login e logout

Validação de e-mail

Confirmação de senha

Atualização de dados do perfil

Exclusão de conta

Autenticação via cookies

##📝 Publicações
Criar publicações

Editar publicações

Excluir publicações

Curtir e descurtir publicações

Visualizar publicações de usuários específicos

Feed com todas as publicações

##🔗 Interações sociais
Seguir usuários

Parar de seguir usuários

Visualizar seguidores

Visualizar quem o usuário está seguindo

##🔄 Atualizações dinâmicas
Requisições AJAX para ações como curtir, seguir e atualizar publicações

#🛠 Tecnologias utilizadas
Backend
Go 1.26.3

MySQL Server 5.7

Gorilla Mux v1.8.1

Gorilla SecureCookie v1.1.2

Go SQL Driver (mysql) v1.10.1

Checkmail v1.2.4

JWT (v3.2.2 e v5.3.1)

Godotenv v1.5.1

x/crypto v0.57.0

filippo.io/edwards25519 v1.2.0

Frontend
HTML5

CSS / Bootstrap

JavaScript

jQuery

Outros
Arquivo .env para configuração

Script SQL para criação do banco

Sem uso de Docker

#📁 Estrutura do projeto
```
DevBook/
│
├── api/
│   ├── sql/
│   ├── src/
│   │   ├── autenticação/
│   │   ├── banco/
│   │   ├── config/
│   │   ├── controllers/
│   │   ├── middlewares/
│   │   ├── modelos/
│   │   ├── repositórios/
│   │   ├── respostas/
│   │   ├── router/
│   │   │   └── rotas/
│   │   └── segurança/
│   └── main.go
│
└── webapp/
    ├── assets/
    │   ├── css/
    │   └── js/
    ├── src/
    │   ├── config/
    │   ├── controllers/
    │   ├── cookies/
    │   ├── middlewares/
    │   ├── modelos/
    │   ├── requisições/
    │   ├── respostas/
    │   ├── router/
    │   │   └── rotas/
    │   └── utils/
    └── views/
        └── templates/
```

#🚀 Como rodar o projeto
1. Clonar o repositório
```
git clone https://github.com/seu-usuario/DevBook.git
```

3. Criar o banco de dados
Código
```
DevBook/api/sql/sql.sql
```
Execute o script no MySQL Server 5.7.

5. Configurar o arquivo .env
Crie um arquivo .env na pasta api/src/config contendo:

Código
```
DB_USUARIO=seu_usuario
DB_SENHA=sua_senha
DB_NOME=devbook
API_PORT=5000
SECRET_KEY=SGEhdclYeSGaNF5kyWclemAWWkRjZUJU3pjeKFZKKtghP0EBwQ2zxWealu/g5dRBFtxrNN8gaWRjOvKd7b4Fhw==
```

4. Rodar o servidor da API
bash
```
cd api
go run main.go
```
ou

bash
```
go build
go run ./api
```

5. Rodar o frontend
O frontend é servido pelo próprio servidor Go da pasta webapp.
bash
```
cd webapp
go run main.go
```

#📡 Endpoints da API
##🔐 Autenticação
| Método | Endpoint | Descrição |
| --- | --- | --- |
| POST | ``/login`` | Realiza login |

##📝 Publicações
| Método | Endpoint | Descrição |
| --- | --- | --- |
| POST | ``/publicacoes`` | Criar publicação |
| GET | ``/publicacoes`` | Listar publicações |
| GET | ``/publicacoes/{publicacaoId}`` | Buscar publicação |
| PUT | ``/publicacoes/{publicacaoId}`` | Atualizar publicação |
| DELETE | ``/publicacoes/{publicacaoId}`` | Excluir publicação |
| GET | ``/usuarios/{usuarioId}/publicacoes`` | Publicações de um usuário |
| POST | ``/publicacoes/{publicacaoId}/curtir`` | Curtir |
| POST | ``/publicacoes/{publicacaoId}/descurtir`` | Descurtir |

##👤 Usuários
| Método | Endpoint | Descrição |
| --- | --- | --- |
| POST | ``/usuarios`` | Criar usuário |
| GET | ``/usuarios`` | Listar usuários |
| GET | ``/usuarios/{usuarioId}`` | Buscar usuário |
| PUT | ``/usuarios/{usuarioId}`` | Atualizar usuário |
| DELETE | ``/usuarios/{usuarioId}`` | Excluir usuário |
| POST | ``/usuarios/{usuarioId}/seguir`` | Seguir |
| POST | ``/usuarios/{usuarioId}/parar-de-seguir`` | Parar de seguir |
| GET | ``/usuarios/{usuarioId}/seguidores`` | Listar seguidores |
| GET | ``/usuarios/{usuarioId}/seguindo`` | Listar seguindo |
| POST | ``/usuarios/{usuarioId}/atualizar-senha`` | Atualizar senha |

##🎓 Créditos
Projeto desenvolvido durante o curso:

Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!  
Ministrado por Otávio Augusto Gallego na Udemy.

Este projeto foi criado exclusivamente para fins de estudo e desenvolvimento pessoal.
