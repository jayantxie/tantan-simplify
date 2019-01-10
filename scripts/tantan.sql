--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: relationship_id_seq; Type: SEQUENCE; Schema: public; Owner: xzy
--

CREATE SEQUENCE public.relationship_id_seq
    START WITH 1000000000
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.relationship_id_seq OWNER TO xzy;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: relationships; Type: TABLE; Schema: public; Owner: xzy
--

CREATE TABLE public.relationships (
    id integer DEFAULT nextval('public.relationship_id_seq'::regclass) NOT NULL,
    from_user_id integer NOT NULL,
    to_user_id integer NOT NULL,
    state integer
);


ALTER TABLE public.relationships OWNER TO xzy;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: xzy
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1076408511
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO xzy;

--
-- Name: users; Type: TABLE; Schema: public; Owner: xzy
--

CREATE TABLE public.users (
    id integer DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
    name character varying(30)
);


ALTER TABLE public.users OWNER TO xzy;


--
-- Name: relationship_id_seq; Type: SEQUENCE SET; Schema: public; Owner: xzy
--

SELECT pg_catalog.setval('public.relationship_id_seq', 1000000039, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: xzy
--

SELECT pg_catalog.setval('public.user_id_seq', 1076408556, true);


--
-- Name: relationships relationships_pkey; Type: CONSTRAINT; Schema: public; Owner: xzy
--

ALTER TABLE ONLY public.relationships
    ADD CONSTRAINT relationships_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: xzy
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: relationships_index; Type: INDEX; Schema: public; Owner: xzy
--

CREATE INDEX relationships_index ON public.relationships USING btree (from_user_id, to_user_id);


--
-- Name: relationships relationship_to_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: xzy
--

ALTER TABLE ONLY public.relationships
    ADD CONSTRAINT relationship_to_user_id_fk FOREIGN KEY (to_user_id) REFERENCES public.users(id);


--
-- Name: relationships relationship_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: xzy
--

ALTER TABLE ONLY public.relationships
    ADD CONSTRAINT relationship_user_id_fk FOREIGN KEY (from_user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

