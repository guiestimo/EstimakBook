INSERT INTO usuarios (nome, nick, email, senha) 
VALUES
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$pmgCiKaIWkC9MkSKP27z9.83lbn4pUGgxdaBIgoZJ5JQGjll71ylK"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$pmgCiKaIWkC9MkSKP27z9.83lbn4pUGgxdaBIgoZJ5JQGjll71ylK"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$pmgCiKaIWkC9MkSKP27z9.83lbn4pUGgxdaBIgoZJ5JQGjll71ylK");


INSERT INTO seguidores (usuario_id, seguidor_id) 
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO publicacoes (titulo, conteudo, autor_id)
VALUES
("Publicação do Usuário 1", "Essa é a publicação do usuario 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuario 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuario 3! Oba!", 3);