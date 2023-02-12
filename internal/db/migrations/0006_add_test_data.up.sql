-- add users
INSERT INTO user(username,password_hash) VALUES ('user1', '$2a$10$9IxVao19OqCVj9No1lySxupoM7Njl2jgxY6Artr4QSzvbdXel0feG');
INSERT INTO user(username,password_hash) VALUES ('user2', '$2a$10$fQxHSzY9sikIhBDRDRKBnePaOcMQMfKUC2r1BhsyxhfSV2PDCzKim');
-- add user1 workouts
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('1', 'Ab Day', '3', '10', 'N/A', '1', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('1', 'Ab Day', '3', '10', 'N/A', '1', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('1', 'Ab Day', '3', '10', 'N/A', '1', '1', '1');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('3', 'Abs', '3', '10', 'N/A', '1', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('3', 'Abs', '3', '10', 'N/A', '1', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('3', 'Abs', '3', '10', 'N/A', '1', '1', '1');
-- add user2 workouts
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('2', 'Core', '2', '12', 'N/A', '2', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('2', 'Core', '2', '12', 'N/A', '2', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('2', 'Core', '2', '12', 'N/A', '2', '1', '1');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('4', 'Core 2', '2', '12', 'N/A', '2', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('4', 'Core 2', '2', '12', 'N/A', '2', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('4', 'Core 2', '2', '12', 'N/A', '2', '1', '1');
-- add user1 workout history
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('1-Feb-2023', '55:00', 'Gassed', '1', '1');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('2-Feb-2023', '51:00', 'Ok', '1', '1');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('3-Feb-2023', '52:00', 'Great', '2', '1');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('4-Feb-2023', '50:00', 'Great', '2', '1');
-- add user2 workout history
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('1-Feb-2023', '55:00', 'Sore', '2', '2');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('2-Feb-2023', '51:00', 'Tired', '2', '2');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('3-Feb-2023', '52:00', 'Not bad', '4', '2');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('4-Feb-2023', '50:00', 'Ok', '4', '2');