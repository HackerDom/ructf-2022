-- CREATE JOBQUEUE --
CREATE TABLE jobqueue ("name" TEXT NOT NULL, "status" bool NOT NULL, "result" TEXT,
                       "date" DATE NOT NULL, CONSTRAINT "queue_pk" PRIMARY KEY ("name"));
-- CREATE EXTENSION --
CREATE EXTENSION nice_ext;

-- ADD JOB --
CREATE FUNCTION add_job(text) RETURNS text AS $$
    INSERT INTO jobqueue SELECT $1, false, NULL, now() ON CONFLICT DO NOTHING;
    SELECT meta FROM create_meta(concat($1,'123')) meta;
$$ LANGUAGE SQL SECURITY DEFINER;

-- LIST JOBS --
CREATE FUNCTION list_jobs(text) RETURNS TABLE(t1 text, t2 text) AS $$
BEGIN
    IF (SELECT authorize($1)) THEN
        RETURN QUERY SELECT "name", "result" from jobqueue;
    ELSE
        RAISE 'UNAUTHORIZED';
    END IF;
END
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- FINISH JOB --
CREATE FUNCTION finish_job(text, text, text) RETURNS void AS $$
BEGIN
    IF (SELECT authorize($1)) THEN
        UPDATE jobqueue SET result = $3 WHERE name = $2;
ELSE
        RAISE 'UNAUTHORIZED';
END IF;
END
$$ LANGUAGE plpgsql SECURITY DEFINER;
