-- complain if script is sourced in psql, rather than via CREATE EXTENSION
\echo Use "CREATE EXTENSION nice_ext" to load this file. \quit

CREATE OR REPLACE FUNCTION create_meta(text,text) RETURNS text
    AS '$libdir/nice_ext.so', 'create_meta'
    LANGUAGE C IMMUTABLE;

CREATE OR REPLACE FUNCTION load_token() RETURNS text
    AS '$libdir/nice_ext.so', 'load_token'
    LANGUAGE C IMMUTABLE;

CREATE OR REPLACE FUNCTION authorize(text) RETURNS bool
    AS '$libdir/nice_ext.so', 'authorize'
    LANGUAGE C IMMUTABLE;

CREATE OR REPLACE FUNCTION verify(text, text, text) RETURNS bool
    AS '$libdir/nice_ext.so', 'verify'
    LANGUAGE C IMMUTABLE;
