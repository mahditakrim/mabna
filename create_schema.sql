--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Homebrew)
-- Dumped by pg_dump version 14.5 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: instruments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.instruments (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.instruments OWNER TO postgres;

--
-- Name: instruments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.instruments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.instruments_id_seq OWNER TO postgres;

--
-- Name: instruments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.instruments_id_seq OWNED BY public.instruments.id;


--
-- Name: trades; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trades (
    id integer NOT NULL,
    instrument_id integer NOT NULL,
    dateen timestamp with time zone NOT NULL,
    open numeric,
    high numeric,
    low numeric,
    close numeric
);


ALTER TABLE public.trades OWNER TO postgres;

--
-- Name: trades_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.trades_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.trades_id_seq OWNER TO postgres;

--
-- Name: trades_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.trades_id_seq OWNED BY public.trades.id;


--
-- Name: instruments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.instruments ALTER COLUMN id SET DEFAULT nextval('public.instruments_id_seq'::regclass);


--
-- Name: trades id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trades ALTER COLUMN id SET DEFAULT nextval('public.trades_id_seq'::regclass);


--
-- Data for Name: instruments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.instruments (id, name) FROM stdin;
5	AAPL
6	GOOGL
8	BTC
\.


--
-- Data for Name: trades; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trades (id, instrument_id, dateen, open, high, low, close) FROM stdin;
15	5	2020-01-01 00:00:00+03:30	1001	1001	301	401
16	5	2020-01-02 00:00:00+03:30	1002	1002	302	402
17	5	2020-01-03 00:00:00+03:30	1003	1003	303	403
18	6	2020-01-01 00:00:00+03:30	1004	1004	304	404
19	6	2020-01-03 00:00:00+03:30	1005	1005	305	405
20	5	2021-01-01 00:00:00+03:30	1007	1007	307	407
21	8	2021-02-02 00:00:00+03:30	1007	1007	307	407
22	8	2021-02-01 00:00:00+03:30	1007	1007	307	407
\.


--
-- Name: instruments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.instruments_id_seq', 9, true);


--
-- Name: trades_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trades_id_seq', 22, true);


--
-- Name: instruments instruments_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.instruments
    ADD CONSTRAINT instruments_name_key UNIQUE (name);


--
-- Name: instruments instruments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.instruments
    ADD CONSTRAINT instruments_pkey PRIMARY KEY (id);


--
-- Name: trades trades_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trades
    ADD CONSTRAINT trades_pkey PRIMARY KEY (id);


--
-- Name: trades trades_instrument_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trades
    ADD CONSTRAINT trades_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES public.instruments(id);


--
-- PostgreSQL database dump complete
--

