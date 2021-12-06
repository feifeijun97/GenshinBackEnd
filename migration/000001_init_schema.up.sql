CREATE DATABASE genshin_wiki
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

COMMENT ON DATABASE genshin_wiki
    IS 'Database for Genshin Wiki App';

GRANT ALL ON DATABASE genshin_wiki TO postgres;

GRANT TEMPORARY, CONNECT ON DATABASE genshin_wiki TO PUBLIC;

GRANT ALL ON DATABASE genshin_wiki TO admin;