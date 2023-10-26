INSERT INTO teachers (login, name, password) VALUES
('art@art',	'aaa', '123');

INSERT INTO students (inviteHash, name) VALUES
('d0e8d5bc-6de7-11ee-b962-0242ac120002', 'bbb'),
('d0e8d5bc-6de7-11ee-b962-0242ac120003', 'bbb2');

INSERT INTO chats (teacherID, studentHash) VALUES
(1, 'd0e8d5bc-6de7-11ee-b962-0242ac120002'),
(1, 'd0e8d5bc-6de7-11ee-b962-0242ac120003');