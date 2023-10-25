DROP TABLE IF EXISTS "messages" CASCADE;
DROP TABLE IF EXISTS "chats" CASCADE;
DROP TABLE IF EXISTS "students" CASCADE;
DROP TABLE IF EXISTS "teachers" CASCADE;

CREATE TABLE teachers (
    id  SERIAL PRIMARY KEY,
    login VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (30),
	name VARCHAR (100) NOT NULL,
    tgAccount VARCHAR (100),
    vkAccount VARCHAR (100),
    tgBotLink VARCHAR (100),
    vkBotLink VARCHAR (100)
);

CREATE TABLE students (
    inviteHash UUID PRIMARY KEY,
	name VARCHAR (100) NOT NULL,
    parentName VARCHAR (100),
    tgStudent VARCHAR (100),
    vkStudent VARCHAR (100),
    tgParent VARCHAR (100),
    vkParent VARCHAR (100)
);

CREATE TABLE chats
(
    id       SERIAL PRIMARY KEY,
    teacherID   INT REFERENCES teachers (id) ON DELETE CASCADE,
    studentHash   UUID REFERENCES students (inviteHash) ON DELETE CASCADE
);

CREATE TABLE messages
(
    id          SERIAL PRIMARY KEY,
    chatID      INT REFERENCES chats (id) ON DELETE CASCADE,
    text        TEXT NOT NULL,
    isAuthorTeacher    BOOLEAN     NOT NULL,
    attaches    VARCHAR[],
    time TIMESTAMP NOT NULL
);

INSERT INTO teachers (login, name, password) VALUES
('art@art',	'aaa', '123');

INSERT INTO students (inviteHash, name) VALUES
('d0e8d5bc-6de7-11ee-b962-0242ac120002', 'bbb');

INSERT INTO chats (id, teacherID, studentHash) VALUES
(1, 2, 'd0e8d5bc-6de7-11ee-b962-0242ac120002');
