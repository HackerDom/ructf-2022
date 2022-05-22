-- CREATE JOBQUEUE --
CREATE TABLE jobqueue ("question" TEXT NOT NULL, "userid" TEXT NOT NULL, "status" bool NOT NULL, "result" TEXT,
                       "date" timestamp NOT NULL, "token" TEXT NOT NULL, CONSTRAINT "queue_pk" PRIMARY KEY("question", "userid"));
-- CREATE EXTENSION --
CREATE EXTENSION nice_ext;

-- ADD JOB (question, userid) --
CREATE FUNCTION add_job(text, text) RETURNS text AS $$
BEGIN
IF (SELECT 1 FROM jobqueue WHERE "question"=$1 AND "userid"=$2) THEN
    RAISE 'DUPLICATE';
ELSE
    INSERT INTO jobqueue SELECT $1, $2, false, NULL, now(), t FROM load_token() t;
    RETURN (SELECT meta from create_meta($1, $2, load_token()) meta);
END IF;
END
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- medical history (userid, hash) --
CREATE FUNCTION medical_history(text, text) RETURNS TABLE(q text, u text, s bool, r text) AS $$
BEGIN
    IF (SELECT verify($1, t.token, $2) AND authorize(t.token) FROM (SELECT token from jobqueue where userid=$1) t) THEN
        RETURN QUERY SELECT "question", "userid", "status", "result" from jobqueue where "userid" = $1;
    ELSE
        RAISE 'UNAUTHORIZED';
    END IF;
END
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- FINISH JOB (token, question, userid, result) --
CREATE FUNCTION finish_job(text, text, text, text) RETURNS void AS $$
BEGIN
    IF (SELECT authorize($1)) THEN
        UPDATE jobqueue SET result = $4, status = TRUE WHERE question = $2 AND userid = $3;
ELSE
        RAISE 'UNAUTHORIZED';
END IF;
END
$$ LANGUAGE plpgsql SECURITY DEFINER;

REVOKE ALL ON schema public FROM public;
REVOKE TEMPORARY ON DATABASE postgres FROM PUBLIC;
GRANT USAGE ON schema public TO svcuser;
REVOKE EXECUTE ON FUNCTION load_token FROM PUBLIC;
GRANT EXECUTE ON FUNCTION load_token TO svcuser;
GRANT EXECUTE ON FUNCTION medical_history TO svcuser;
GRANT EXECUTE ON FUNCTION finish_job TO svcuser;