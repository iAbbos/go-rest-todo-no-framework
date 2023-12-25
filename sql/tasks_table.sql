CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    string VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    note TEXT,
    task_date DATE NOT NULL,
    status VARCHAR(16) NOT NULL
);
