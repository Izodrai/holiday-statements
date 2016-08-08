CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login VARCHAR(255),
    pwd   VARCHAR(255),
    email VARCHAR(255),
    UNIQUE (login),
    UNIQUE (email)
);

INSERT INTO users(login, pwd, email) values("admin", "$2y$10$lsliCJPCiCsPOrgTgidDrumYuTwg3MGW6CQIy4nn7ziu8OXNiHbpO", "admin@mydomain.country");

CREATE TABLE know (
    user_id_1 UNSIGNED INTEGER,
    user_id_2 UNSIGNED INTEGER,
    FOREIGN KEY(user_id_1) REFERENCES users(id),
    FOREIGN KEY(user_id_2) REFERENCES users(id)
);

