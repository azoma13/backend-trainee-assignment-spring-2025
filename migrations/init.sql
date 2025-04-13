DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'avitoPVZ') THEN
        CREATE DATABASE "avitoPVZ";
    END IF;
END $$;