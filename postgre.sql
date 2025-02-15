
CREATE TABLE ActivityDesc (
    id BIGSERIAL PRIMARY KEY,
    category VARCHAR(50) NOT NULL,
    time VARCHAR(50) NOT NULL,
    frequency VARCHAR(20) NOT NULL,
    activity_name TEXT NOT NULL
);

INSERT INTO ActivityDesc (id, category, activity_name, time, frequency) VALUES
    (1, 'Athleticism', 'Advanced Mobility exercises', 'Max.', 'Maximize'),
    (2, 'Athleticism', 'Finger Skills', '3x/Week', '60 sec'),
    (3, 'Boosters', 'Knowledge Boosters (Follow daily plans)', '2x/Day', '30 sec'),
    (4, 'Music', 'Visual Solfage', '1x/Day', '30 sec'),
    (5, 'Memory', 'Auditory Memory (Song 2)', '1x/Day', '30 sec'),
    (6, 'Creativity', 'Auditory Magic (Set 2)', '2 sounds/Day', '60 sec'),
    (7, 'Creativity', 'Stimulus Explosion', '2x/Week', '60 sec'),
    (8, 'Languages', 'Talk, To Listen', '1x/Day', '60 sec'),
    (9, 'Logic', 'Foundation of Logic (Quantity)', '2x/Week', '60 sec');

CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO Users(name, email) VALUES
    ('Suraj Patel', 'patelsuraj301@gmail.com');


CREATE TABLE WeekPlan (
    id BIGSERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

-- as per current logic for a single day- there will be 9 activities
-- completed_at is Null means the activity is not yet completed.
CREATE TABLE DayPlan (
    id SERIAL PRIMARY KEY,
    userId INT REFERENCES Users(id) ON DELETE CASCADE,
    weekId INT REFERENCES WeekPlan(id) ON DELETE CASCADE,
    activityId INT REFERENCES ActivityDesc(id) ON DELETE CASCADE,
    completed_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);