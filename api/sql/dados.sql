insert into usuarios(nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@email.com", "$2a$10$neOohv5HIXyRuOVbHzsPjux3gXfnAlfg6COZeDeMt6SHy50QREQGi"),
("Usuario 2", "usuario_2", "usuario2@email.com", "$2a$10$neOohv5HIXyRuOVbHzsPjux3gXfnAlfg6COZeDeMt6SHy50QREQGi"),
("Usuario 3", "usuario_3", "usuario3@email.com", "$2a$10$neOohv5HIXyRuOVbHzsPjux3gXfnAlfg6COZeDeMt6SHy50QREQGi");

insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Usuário 1", "Publi do usuário 1", 1),
("Usuário 2", "Publi do usuário 2", 2),
("Usuário 3", "Publi do usuário 3", 3);
