CREATE TABLE surveys (
    id serial NOT NULL unique,
    title text NOT NULL
);

CREATE TABLE questions (
    id serial NOT NULL unique,
    text text NOT NULL,
    survey_id int,
    FOREIGN KEY (survey_id) REFERENCES surveys(id) ON DELETE CASCADE
);

CREATE TABLE answers (
    id serial   NOT NULL unique,
    text text NOT NULL,
    survey_id int,
    question_id int,
    votes int DEFAULT 0,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE,
    FOREIGN KEY (survey_id) REFERENCES surveys(id) ON DELETE CASCADE
)