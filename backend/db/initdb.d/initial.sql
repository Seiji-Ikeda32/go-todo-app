USE go_db;

CREATE TABLE IF NOT EXISTS todos (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(30),
    discription VARCHAR(100),
    isCompleted BOOLEAN NOT NULL DEFAULT '0',
    dueTime DATETIME,
    createAt TIMESTAMP,
    updataAt TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO todos (title, discription, dueTime) VALUES ("宿題", "宿題をします", "2022-03-01 08:00:00");
INSERT INTO todos (title, discription, dueTime) VALUES ("洗濯", "乾燥まで予約します", "2022-02-22 08:00:00");
INSERT INTO todos (title, discription, dueTime) VALUES ("英語学習", "45ページから60ページまで", "2022-02-19 08:00:00");
INSERT INTO todos (title, discription, dueTime) VALUES ("CS学習", "コンピュータの仕組みを学習", "2022-03-02 08:00:00");
INSERT INTO todos (title, discription, dueTime) VALUES ("gopher", "gopher君可愛い", "2022-02-25 08:00:00");