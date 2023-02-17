-- add users
INSERT INTO user(username,password_hash) VALUES ('user1', '$2a$10$9IxVao19OqCVj9No1lySxupoM7Njl2jgxY6Artr4QSzvbdXel0feG');
INSERT INTO user(username,password_hash) VALUES ('user2', '$2a$10$fQxHSzY9sikIhBDRDRKBnePaOcMQMfKUC2r1BhsyxhfSV2PDCzKim');
-- add user1 workouts
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a5e0-0c0ebb3f9376', 'Ab Day', '3', '10', 'N/A', '1', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a5e0-0c0ebb3f9376', 'Ab Day', '3', '10', 'N/A', '1', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a5e0-0c0ebb3f9376', 'Ab Day', '3', '10', 'N/A', '1', '1', '1');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a770-0c0ebb3f9376', 'Abs', '3', '10', 'N/A', '1', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a770-0c0ebb3f9376', 'Abs', '3', '10', 'N/A', '1', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a770-0c0ebb3f9376', 'Abs', '3', '10', 'N/A', '1', '1', '1');
-- add user2 workouts
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-44e0-0c0ebb3f9376', 'Core', '2', '12', 'N/A', '2', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-44e0-0c0ebb3f9376', 'Core', '2', '12', 'N/A', '2', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-44e0-0c0ebb3f9376', 'Core', '2', '12', 'N/A', '2', '1', '1');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a599-0c0ebb3f9376', 'Core 2', '2', '12', 'N/A', '2', '2', '2');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a599-0c0ebb3f9376', 'Core 2', '2', '12', 'N/A', '2', '3', '3');
INSERT INTO workout(plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id) VALUES ('ec73413c-2156-46f2-a599-0c0ebb3f9376', 'Core 2', '2', '12', 'N/A', '2', '1', '1');
-- add user1 workout history
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('1-Feb-2023', '55:00', 'Gassed', 'ec73413c-2156-46f2-a5e0-0c0ebb3f9376', '1');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('2-Feb-2023', '51:00', 'Ok', 'ec73413c-2156-46f2-a5e0-0c0ebb3f9376', '1');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('3-Feb-2023', '52:00', 'Great', 'ec73413c-2156-46f2-a770-0c0ebb3f9376', '1');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('4-Feb-2023', '50:00', 'Great', 'ec73413c-2156-46f2-a770-0c0ebb3f9376', '1');
-- add user2 workout history
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('1-Feb-2023', '55:00', 'Sore', 'ec73413c-2156-46f2-44e0-0c0ebb3f9376', '2');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('2-Feb-2023', '51:00', 'Tired', 'ec73413c-2156-46f2-44e0-0c0ebb3f9376', '2');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('3-Feb-2023', '52:00', 'Not bad', 'ec73413c-2156-46f2-a599-0c0ebb3f9376', '2');
INSERT INTO history(date, duration, notes, plan_id, athlete_id) VALUES ('4-Feb-2023', '50:00', 'Ok', 'ec73413c-2156-46f2-a599-0c0ebb3f9376', '2');