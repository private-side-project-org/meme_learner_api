DROP table users;
DROP table memes;
Drop SEQUENCE public.users_id_seq;
--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL PRIMARY KEY,
    name text NOT NULL,
    password text NOT NULL, 
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: memes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.memes (
    id integer NOT NULL PRIMARY KEY,
    title text,
    description text,
    rating integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    user_id integer NOT NULL
);

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO users VALUES (1, 'David', 'testTest', '2021-12-24 00:00:00', '2021-12-24 00:00:00');
INSERT INTO users VALUES (2, 'Sato', 'testTest', '2021-12-24 00:00:00', '2021-12-24 00:00:00');
INSERT INTO users VALUES (3, 'Yamada', 'testTest', '2021-12-24 00:00:00', '2021-12-24 00:00:00');
INSERT INTO users VALUES (4, 'Tokyu', 'testTest', '2021-12-24 00:00:00', '2021-12-24 00:00:00');
INSERT INTO users VALUES (5, 'Shitty', 'testTest', '2021-12-24 00:00:00', '2021-12-24 00:00:00');

--
-- Data for Name: memes; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO memes VALUES (1, 'coffin song', 'coffin song is something funny', 5, '2021-05-17 12:12:12', '2021-05-17 12:12:12', 1);
INSERT INTO memes VALUES (2, 'test1', 'test test test', 3, '2021-05-17 12:12:12', '2021-05-17 12:12:12', 2);
INSERT INTO memes VALUES (3, 'test2', 'test2 is something funny', 3, '2021-05-17 12:12:12', '2021-05-17 12:12:12', 2);
INSERT INTO memes VALUES (4, 'test3', 'test3 is something funny', 2, '2021-05-17 12:12:12', '2021-05-17 12:12:12', 1);
INSERT INTO memes VALUES (5, 'test4', 'test4 is something funny', 1, '2021-05-17 12:12:12', '2021-05-17 12:12:12', 1);

-- ---
-- ---Joint users table on memes based on user_id
-- ---

-- --- get all info from memes
SELECT memes.id, memes.title, memes.description, memes.rating, memes.created_at, memes.updated_at from memes
-- --- create joint table that join memes on users if user_id is match users
INNER JOIN users ON memes.user_id = users.id WHERE user_id = 2