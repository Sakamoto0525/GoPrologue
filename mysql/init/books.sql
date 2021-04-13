DROP TABLE IF EXISTS books;

CREATE TABLE books (
    id integer,
    title varchar(255),
    auther_id integer
);

INSERT INTO books VALUES
(1, 'タイトル１', 1),
(2, 'タイトル２', 2),
(3, 'タイトル３', 3)
