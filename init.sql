CREATE TABLE IF NOT EXISTS member (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    is_active BOOLEAN NOT NULL,
    updated_at DATETIME NOT NULL
    );

-- mock data
INSERT INTO member (name, created_at, is_active, updated_at) VALUES
                ('Alice', NOW(), TRUE, NOW()),
                ('Bob', NOW(), FALSE, NOW()),
                ('Charlie', NOW(), TRUE, NOW()),
                ('Dave', NOW(), TRUE, NOW()),
                ('Eve', NOW(), FALSE, NOW());
