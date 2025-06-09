CREATE DATABASE IF NOT EXISTS sumup;

USE sumup;

CREATE TABLE IF NOT EXISTS user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(20)
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



INSERT INTO user_notification (id, user_id, notification_id)
SELECT 1, 1, 1
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 1);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 2, 1, 2
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 2);

INSERT INTO user_notification (id, user_id, notification_id)
SELECT 3, 2, 1
WHERE NOT EXISTS (SELECT 1 FROM user_notification WHERE id = 3);