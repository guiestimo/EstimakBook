INSERT INTO usuarios (nome, nick, email, senha) 
VALUE
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$pmgCiKaIWkC9MkSKP27z9.83lbn4pUGgxdaBIgoZJ5JQGjll71ylK"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$pmgCiKaIWkC9MkSKP27z9.83lbn4pUGgxdaBIgoZJ5JQGjll71ylK"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$pmgCiKaIWkC9MkSKP27z9.83lbn4pUGgxdaBIgoZJ5JQGjll71ylK");


INSERT INTO seguidores (usuario_id, seguidor_id) 
VALUE
(1, 2),
(3, 1),
(1, 3);