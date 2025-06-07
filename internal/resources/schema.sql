CREATE DATABASE IF NOT EXISTS sumup;

USE sumup;

CREATE TABLE IF NOT EXISTS user (
    id INT PRIMARY KEY,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS notification (
    id INT PRIMARY KEY,
    type VARCHAR(100) NOT NULL  
);

CREATE TABLE IF NOT EXISTS user_notification (
    id INT PRIMARY KEY,
    user_id INT NOT NULL,
    notification_id INT NOT NULL, 
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (notification_id) REFERENCES notification(id) ON DELETE CASCADE,
    UNIQUE (user_id, notification_id) 
);


INSERT INTO user (id, email)
SELECT 1, 'johndoe@email.com'
WHERE NOT EXISTS (SELECT 1 FROM user WHERE id = 1);

INSERT INTO user (id, email)
SELECT 2, 'janedoe@email.com'
WHERE NOT EXISTS (SELECT 1 FROM user WHERE id = 2);



INSERT INTO notification (id, type)
SELECT 1, 'email'
WHERE NOT EXISTS (SELECT 1 FROM notification WHERE id = 1);

INSERT INTO notification (id, type)
SELECT 2, 'sms'
WHERE NOT EXISTS (SELECT 1 FROM notification WHERE id = 2);



INSERT INTO user_notification (id, user_id, notification_id)
SELECT 1, 1, 1
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 1);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 2, 1, 2
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 2);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 3, 2, 1
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 3);