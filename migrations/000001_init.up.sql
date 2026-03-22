CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15),
    version BIGINT
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    title VARCHAR(100) NOT NULL,
    description VARCHAR(1000),
    completed BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    completed_at TIMESTAMPTZ,
    CONSTRAINT fk_tasks_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

