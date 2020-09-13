CREATE TABLE test_db.accounts (
	id         int unsigned auto_increment PRIMARY KEY,
	username VARCHAR(255) NULL,
	email VARCHAR(255) NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	UNIQUE KEY unique_email (email)
) CHARACTER SET utf8mb4;