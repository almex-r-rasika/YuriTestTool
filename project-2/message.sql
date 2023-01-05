DROP TABLE IF EXISTS messages;

CREATE TABLE messages
(msg_id TEXT NOT NULL,
 send_time TEXT NOT NULL,
 send_user_id TEXT NOT NULL,
 address TEXT NOT NULL,
 subject TEXT NOT NULL,
 line1 TEXT,
 line2 TEXT,
 line3 TEXT,
 line4 TEXT,
 line5 TEXT,
 line6 TEXT,
 line7 TEXT,
 line8 TEXT,
 line9 TEXT,
 line10 TEXT
);


INSERT INTO messages
VALUES ('1','2022-12-27T19:02:10+09:00','install_01','0000001234,本人','case1','あ','い','う','え','お','か','き','く','け','こ');

INSERT INTO messages
VALUES ('2','2022-12-27T19:02:10+09:00','install_02','0000001234,本人','case2','複数同時','い','う','え','お','か','き','く','け','こ');

INSERT INTO messages
VALUES ('3','2022-12-27T19:02:10+09:00','install_01','0000001234,本人','case3','あ','い','う','え','お','か','き','く','け','こ');

