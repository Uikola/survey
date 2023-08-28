CREATE TABLE surveys (
    id serial NOT NULL unique,
    title text NOT NULL
);

CREATE TABLE questions (
    id serial NOT NULL unique,
    text text NOT NULL,
    survey_id int,
    FOREIGN KEY (survey_id) REFERENCES surveys(id)
)