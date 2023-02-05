-- add test user
INSERT INTO user(username,password_hash) VALUES ('test', '$2a$10$oemb0b.FDaWWlsvbGGrHwuD1K/k6NBSbbuLNvPvLzKNTe5U969Jai');
-- add test workout plan
INSERT INTO workout_plan(workout_id, name, sets, reps, creator_id, exercise_id) VALUES ('1', 'Test Workout', '1', '1', '1', '1');
INSERT INTO workout_plan(workout_id, name, sets, reps, creator_id, exercise_id) VALUES ('1', 'Test Workout', '2', '2', '1', '2');
INSERT INTO workout_plan(workout_id, name, sets, reps, creator_id, exercise_id) VALUES ('1', 'Test Workout', '3', '3', '1', '3');