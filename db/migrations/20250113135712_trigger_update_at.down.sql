DROP TRIGGER IF EXISTS set_updated_at_folders ON folders;
DROP TRIGGER IF EXISTS set_updated_at_chats ON chats;

DROP FUNCTION IF EXISTS updating_updated_at ();