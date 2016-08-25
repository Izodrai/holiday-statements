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
    debt REAL,
    FOREIGN KEY(spending_id) REFERENCES spending(id),
    FOREIGN KEY(debtor_id) REFERENCES users(id)
);

CREATE TABLE debts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    debtor_id INTEGER,
    creditor_id INTEGER,
    event_id INTEGER,
    amount REAL,
    paid INTEGER,
    FOREIGN KEY(debtor_id) REFERENCES users(id),
    FOREIGN KEY(creditor_id) REFERENCES users(id),
    FOREIGN KEY(event_id) REFERENCES events(id)
);



INSERT INTO users (id, login, pwd, email, rights) 
VALUES
	(1, "admin", "$2y$10$lsliCJPCiCsPOrgTgidDrumYuTwg3MGW6CQIy4nn7ziu8OXNiHbpO", "admin@mydomain.country", 1),
	(2, "user1", "$2y$10$UgGYS7.GcVpauYMdmwGLEuf7bcOFii0g/OPQtjnCBCRdmHGmCSm.K", "users1@mydomain.country", 0);
	
INSERT INTO events (id, reference, created_at, promoter_id)
VALUES 
	(1,"holidays in Barcelona !", 1440414793, 1),
	(2,"picnic at user2", 1472062820, 1),
	(3,"camping paradise", 1472070142, 2);
	
INSERT INTO participants (user_id, event_id)
VALUES
	(1,1),
	(2,1),
	(1,2),
	(2,3);
	
INSERT INTO spending_type (id, reference)
VALUES 
	(1, "food"),
	(2, "accommodation");
	
INSERT INTO spending (id, event_id, type_id, description, amount, spending_at, created_at, payer_id)
VALUES 
	(1, 1, 1, "beers", 12.35, 1440416501, 1440416502, 1),
	(2, 1, 2, "camping", 350.10, 1440436501, 1440436502, 1);
	
INSERT INTO spending_for (spending_id, debtor_id, debt)
VALUES 
	(1, 1, 6.175),
	(1, 2, 6.175),
	(2, 1, 175.05),
	(2, 2, 175.05);