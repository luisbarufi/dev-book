insert into users (name, nick, email, password)
values
("Usuário 1", "usuário_1", "usuario_1@example.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq"),
("Usuário 2", "usuário_2", "usuario_2@example.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq"),
("Usuário 3", "usuário_3", "usuario_3@example.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq"),
("Luis Barufi", "Barufi", "luisbarufi@gmail.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq");

insert into followers (user_id, follower_id)
values
(4, 1),
(4, 2),
(4, 3),
(1, 4),
(2, 4),
(1, 2),
(3, 1),
(1, 3);

insert into posts(title, content, author_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Outra publicação do Usuário 1", "Essa é outra publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Outra publicação do Usuário 2", "Essa é a outra publicação do usuário 2! Oba!", 2),
("Meu Post", "Este é meu Post", 4),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3),
("Outra publicação do Usuário 3", "Essa é a outra publicação do usuário 3! Oba!", 3);
