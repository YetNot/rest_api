CREATE TABLE author (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    author_id UUID NOT NULL,
    CONSTRAINT author_fr FOREIGN KEY(author_id) REFERENCES author(id)
);