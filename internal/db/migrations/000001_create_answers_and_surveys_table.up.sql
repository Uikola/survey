CREATE TABLE surveys (
    id serial NOT NULL unique,
    title text NOT NULL
);

CREATE TABLE answers (
    id serial   NOT NULL unique,
    text text NOT NULL,
    survey_id int,
    votes int DEFAULT 0,
    FOREIGN KEY (survey_id) REFERENCES surveys(id) ON DELETE CASCADE
)