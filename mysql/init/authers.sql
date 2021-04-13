DROP TABLE IF EXISTS authers;

CREATE TABLE authers (
    id integer,
    first_name varchar(255),
    last_name varchar(255)
);

INSERT INTO authers VALUES
(1, '田中', '太郎'),
(2, '山田', '花子'),
(3, 'テスト', 'テスト')
