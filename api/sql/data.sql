insert into users (name, nick, email, password)
values
("Usuário 1", "usuário_1", "usuario_1@example.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq"),
("Usuário 2", "usuário_2", "usuario_2@example.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq"),
("Usuário 3", "usuário_3", "usuario_3@example.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq"),
("Luis Barufi", "Barufi", "luisbarufi@gmail.com", "$2a$10$OMDrN1cJNg.QwYn8HLB5pu/WagBusESKe9lJWMR23O4m81znig7qq");

insert into followers (user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);
