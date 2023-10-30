INSERT INTO teachers (login, name, password) VALUES
('art@art',	'aaa', '123');

INSERT INTO students (inviteHash, name) VALUES
('d0e8d5bc-6de7-11ee-b962-0242ac120002', 'bbb'),
('d0e8d5bc-6de7-11ee-b962-0242ac120003', 'bbb2');

INSERT INTO chats (teacherID, studentHash) VALUES
(1, 'd0e8d5bc-6de7-11ee-b962-0242ac120002'),
(1, 'd0e8d5bc-6de7-11ee-b962-0242ac120003');

INSERT INTO messages (chatID, text, isAuthorTeacher, time, isRead) VALUES
(1, 'mes1', true, '2017-03-17 19:43:35.178882-07', true),
(1, 'mes2', false, '2020-03-17 19:22:01.178882-07', false)
(2, 'mes21', true, '2017-03-17 19:43:35.178882-07', true),
(2, 'mes22', false, '2020-03-17 19:22:01.178882-07', false);