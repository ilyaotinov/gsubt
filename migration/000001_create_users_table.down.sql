BEGIN;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;
DROP FUNCTION IF EXISTS update_updated_at_column;
COMMIT;