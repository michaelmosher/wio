CREATE TABLE IF NOT EXISTS jira_users (
    id       serial PRIMARY KEY,
    username varchar(40) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS jira_issues (
    id          serial PRIMARY KEY,
    jira_id     varchar(10) NOT NULL,
    jira_key    varchar(10) NOT NULL UNIQUE,
    external_id varchar(20) NOT NULL
);
