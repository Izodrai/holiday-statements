CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login VARCHAR(255) NOT NULL,
    pwd   VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    rights INTEGER,
    active INTEGER,
    UNIQUE (login),
    UNIQUE (email)
);

CREATE TABLE friends (
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

CREATE TABLE spending_types (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL
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
    FOREIGN KEY(type_id) REFERENCES spending_types(id),
    FOREIGN KEY(payer_id) REFERENCES users(id)
);

CREATE TABLE spending_for (
    spending_id INTEGER,
    debtor_id INTEGER,
    debt REAL,
    refunded INTEGER,
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



INSERT INTO users (id, login, pwd, email, rights, active)
VALUES
	(1, "admin", "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918", "admin@mydomain.country", 1, 1),
	(2, "user1", "0a041b9462caa4a31bac3567e0b6e6fd9100787db2ab433d96f6d178cabfce90", "users1@mydomain.country", 0, 1),
	(3, "user2", "6025d18fe48abd45168528f18a82e265dd98d421a7084aa09f61b341703901a3", "users2@mydomain.country", 0, 1),
	(4, "user3", "5860faf02b6bc6222ba5aca523560f0e364ccd8b67bee486fe8bf7c01d492ccb", "users3@mydomain.country", 0, 1);

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
	(2,3),
	(3,1),
	(4,1);

INSERT INTO spending_types (id, name)
VALUES
	(1, "food"),
	(2, "accommodation");

INSERT INTO spending (id, event_id, type_id, description, amount, spending_at, created_at, payer_id)
VALUES
	(1, 1, 1, "beers", 12.35, 1440416501, 1440416502, 1),
	(2, 1, 2, "camping", 350.10, 1440436501, 1440436502, 1),
	(3, 1, 2, "food", 35, 1440436501, 1440436502, 3),
	(4, 1, 2, "resto", 43.53, 1440436501, 1440436502, 4),
	(5, 1, 2, "hotel", 954.27, 1440436501, 1440436502, 4);

INSERT INTO spending_for (spending_id, debtor_id, debt)
VALUES
	(1, 1, 6.175),
	(1, 2, 6.175),
	(2, 1, 87.525),
	(2, 2, 87.525),
	(2, 3, 87.525),
	(2, 4, 87.525),
	(3, 1, 11.6666666667),
	(3, 2, 11.6666666667),
	(3, 3, 11.6666666667),
	(4, 1, 21.765),
	(4, 2, 21.765),
	(5, 2, 318.09),
	(5, 3, 318.09),
	(5, 4, 318.09);

INSERT INTO friends (user_id_1, user_id_2)
VALUES
	(1, 2),
	(1, 3);
