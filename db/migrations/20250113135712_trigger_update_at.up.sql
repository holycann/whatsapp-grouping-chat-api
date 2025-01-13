
CREATE OR REPLACE FUNCTION updating_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_chats
BEFORE UPDATE ON chats
FOR EACH ROW
EXECUTE FUNCTION updating_updated_at();

CREATE TRIGGER set_updated_at_folders
BEFORE UPDATE ON folders
FOR EACH ROW
EXECUTE FUNCTION updating_updated_at();