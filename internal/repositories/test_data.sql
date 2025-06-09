CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT,
    phone TEXT
);

CREATE TABLE IF NOT EXISTS notification (
    id INTEGER PRIMARY KEY, 
    type TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_notification (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    notification_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (notification_id) REFERENCES notification(id) ON DELETE CASCADE,
    UNIQUE (user_id, notification_id)
);

INSERT INTO user (id, name, email, phone)
SELECT 1, 'John Doe', 'johndoe@email.com', '123456789'
WHERE NOT EXISTS (SELECT 1 FROM user WHERE id = 1);

INSERT INTO user (id, name, email, phone)
SELECT 2, 'Jane Doe', 'janedoe@email.com', '987654321'
WHERE NOT EXISTS (SELECT 1 FROM user WHERE id = 2);

INSERT INTO notification (id, type)
SELECT 1, 'email'
WHERE NOT EXISTS (SELECT 1 FROM notification WHERE id = 1);

INSERT INTO notification (id, type)
SELECT 2, 'sms'
WHERE NOT EXISTS (SELECT 1 FROM notification WHERE id = 2);

INSERT INTO notification (id, type)
SELECT 3, 'invalid'
WHERE NOT EXISTS (SELECT 1 FROM notification WHERE id = 3);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 1, 1, 1
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 1);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 2, 1, 2
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 2);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 3, 2, 1
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 3);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 4, 1, 3
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 4);