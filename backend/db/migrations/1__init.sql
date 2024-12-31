CREATE TABLE IF NOT EXISTS profiles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT UNIQUE,
    is_default BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS expenses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    amount REAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    profile_id INTEGER,
    FOREIGN KEY (profile_id) REFERENCES profiles (id)
);

CREATE TABLE IF NOT EXISTS earnings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    amount REAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    profile_id INTEGER,
    FOREIGN KEY (profile_id) REFERENCES profiles (id)
);

CREATE TABLE IF NOT EXISTS expense_category (
    expense_id INTEGER,
    category_id INTEGER,
    PRIMARY KEY (expense_id, category_id),
    FOREIGN KEY (expense_id) REFERENCES expenses (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TRIGGER IF NOT EXISTS prevent_created_at_update_profiles
BEFORE UPDATE OF created_at ON profiles
BEGIN
    SELECT RAISE(FAIL, 'Cannot update created_at column in profiles table');
END;

CREATE TRIGGER IF NOT EXISTS prevent_created_at_update_expenses
BEFORE UPDATE OF created_at ON expenses
BEGIN
    SELECT RAISE(FAIL, 'Cannot update created_at column in expenses table');
END;

CREATE TRIGGER IF NOT EXISTS prevent_created_at_update_earnings
BEFORE UPDATE OF created_at ON earnings
BEGIN
    SELECT RAISE(FAIL, 'Cannot update created_at column in earnings table');
END;

CREATE TRIGGER IF NOT EXISTS prevent_created_at_update_categories
BEFORE UPDATE OF created_at ON categories
BEGIN
    SELECT RAISE(FAIL, 'Cannot update created_at column in categories table');
END;

INSERT INTO categories (description) SELECT 'Food' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Food');
INSERT INTO categories (description) SELECT 'Transport' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Transport');
INSERT INTO categories (description) SELECT 'Entertainment' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Entertainment');
INSERT INTO categories (description) SELECT 'Health' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Health');
INSERT INTO categories (description) SELECT 'Education' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Education');
INSERT INTO categories (description) SELECT 'Home' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Home');
INSERT INTO categories (description) SELECT 'Shopping' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Shopping');
INSERT INTO categories (description) SELECT 'Investments' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Investments');
INSERT INTO categories (description) SELECT 'Others' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Others');