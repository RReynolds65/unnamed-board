
CREATE TABLE posts (
post_id integer PRIMARY KEY,
image_file text,
user text,
date text,
thread_id integer NOT NULL,
post_text text
);

CREATE TABLE thread (
thread_id integer PRIMARY KEY,
title text NOT NULL,
thread_text NOT NULL,
image_file text NOT NULL,
board_id integer NOT NULL
);

CREATE TABLE boards (
board_id integer PRIMARY KEY,
board_name text NOT NULL,
board_desc text NOT NULL
);

