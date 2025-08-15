-- Increment User Beatdrops Trigger 
CREATE OR REPLACE FUNCTION increment_user_beatdrops()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE users
    SET beatdrops = beatdrops + 1
    WHERE id = NEW.userid;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_beat_insert
AFTER INSERT ON beats
FOR EACH ROW
EXECUTE FUNCTION increment_user_beatdrops();

-- Decrement User Beatdrops Trigger 
CREATE OR REPLACE FUNCTION decrement_user_beatdrops()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE users
    SET beatdrops = beatdrops - 1
    WHERE id = OLD.userid;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_beat_delete
AFTER DELETE ON beats
FOR EACH ROW
EXECUTE FUNCTION decrement_user_beatdrops();

-- Increment User Friends Trigger 
CREATE OR REPLACE FUNCTION increment_friends_on_accept()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 1 AND OLD.status <> 1 THEN
        UPDATE users
        SET friends = friends + 1
        WHERE id IN (NEW.alpha, NEW.beta);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_friend_accept
AFTER UPDATE ON friends
FOR EACH ROW
EXECUTE FUNCTION increment_friends_on_accept();

-- Decrement User Friends Trigger 
CREATE OR REPLACE FUNCTION decrement_friends_on_remove()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.status = 1 AND (TG_OP = 'DELETE' OR NEW.status = -1) THEN
        UPDATE users
        SET friends = friends - 1
        WHERE id IN (OLD.alpha, OLD.beta);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_friend_remove
AFTER UPDATE OR DELETE ON friends
FOR EACH ROW
EXECUTE FUNCTION decrement_friends_on_remove();

-- Increment Beat Comments Trigger
CREATE OR REPLACE FUNCTION increment_beat_comments()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE beats
    SET comments = comments + 1
    WHERE id = NEW.beatid;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_increment_comments
AFTER INSERT ON comments
FOR EACH ROW
EXECUTE FUNCTION increment_beat_comments();

-- Decrement Beat Comments Trigger
CREATE OR REPLACE FUNCTION decrement_beat_comments()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE beats
    SET comments = comments - 1
    WHERE id = OLD.beatid;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_decrement_comments
AFTER DELETE ON comments
FOR EACH ROW
EXECUTE FUNCTION decrement_beat_comments();
