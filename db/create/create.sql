CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login VARCHAR(255) NOT NULL,
    pwd   VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    rights INTEGER,
    UNIQUE (login),
    UNIQUE (email)
);

CREATE TABLE know (
    user_id_1 INTEGER,
    user_id_2 INTEGER,
    FOREIGN KEY(user_id_1) REFERENCES users(id),
    FOREIGN KEY(user_id_2) REFERENCES users(id)
);

CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    reference VARCHAR(255) NOT NULL,
    created_at INTEGER,
    promoter_id INTEGER,
    FOREIGN KEY(promoter_id) REFERENCES users(id)
);

CREATE TABLE participants (
    user_id INTEGER,
    event_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(event_id) REFERENCES events(id)
);

CREATE TABLE spending_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    reference VARCHAR(255) NOT NULL
);

CREATE TABLE spending (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER,
    type_id INTEGER,
    payer_id INTEGER,
    description VARCHAR(255) NOT NULL,
    amount REAL,
    spending_at INTEGER,
    created_at INTEGER,
    FOREIGN KEY(event_id) REFERENCES events(id),
    FOREIGN KEY(type_id) REFERENCES spending_type(id),
    FOREIGN KEY(payer_id) REFERENCES users(id)
);

CREATE TABLE spending_for (
    spending_id INTEGER,
    debtor_id INTEGER,
    FOREIGN KEY(spending_id) REFERENCES spending(id),
    FOREIGN KEY(debtor_id) REFERENCES users(id)
);


INSERT INTO users(login, pwd, email, rights) 
values
	("admin", "$2y$10$lsliCJPCiCsPOrgTgidDrumYuTwg3MGW6CQIy4nn7ziu8OXNiHbpO", "admin@mydomain.country", 1),
	("user1", "$2y$10$UgGYS7.GcVpauYMdmwGLEuf7bcOFii0g/OPQtjnCBCRdmHGmCSm.K", "users1@mydomain.country", 0);