# scholar-power-api

UMGC CMSC 495 API for Scholar Power Workout App
Backend API for [Scholar Power](https://github.com/MoistCode/scholar-power) workout tracking app.

## Purpose

Workout Tracker is an application to help people keep track of their workouts.

## Features

 Users may create an account or choose “demo” to access pre generated content. The demo feature will show the user workout entry screen along with a sample history. Users will be able to create an account by choosing a unique username and password. Users can then login in to create their “workout” or view history to see previous workouts completed. When a user selects the option to create a workout a timer will start. The user will select an exercise from a drop-down menu and then insert Reps, Sets and weight values.

## Startup Notes

Using the makefile is the easiest way to start the app.

`make run`

Of course, you can always run `go run cmd/server/main.go`

All endpoints that modify the resources will require a valid bearer token.

## API Endpoints

### User Routes

Create User
POST /api/v1/user
Params: username: string, password: string
Return: token or an error

Get User By ID
GET /api/v1/user/{id:[0-9]+}
Params: None
Return: User or an error

Get User By Username
GET /api/v1/user/{username}
Params: None
Return: User or an error

Update User Password
PUT /api/v1/user/{id:[0-9]+}
Params: password: string
Return: error

Delete User
DELETE /api/v1/user/{id:[0-9]+}
Params: None
Return: error

Login
POST /api/v1/auth
Params: username: string, password: string
Return: token or an error

### Exercise routes

Get Exercise By ID
GET /api/v1/exercise/{id:[0-9]+}
Params: None
Return: Exercise or an error

Get Exercise By Name
GET /api/v1/exercise/name
Params: name: string
Return: []Exercise or an error

Get Exercise By Muscle Group
GET /api/v1/exercise/muscle
Params: muscle: string
Return: []Exercise or an error

Get Exercise By Equipment
GET /api/v1/exercise/equipment
Params: equipment: string
Return: []Exercise or an error

### Workout routes

Create Workout
POST /api/v1/workout
Params: workout_id: string, name: string, sets: string, reps: string, creator_id: string, exercise_id: string
Return: Workout

Get Workout By ID
GET /api/v1/workout/{workout_id:[0-9]+}
Params: None
Return: []Workout or an error

Get Workout By User
GET /api/v1/workout/user
Params: name: string
Return: []Workout or an error

Update Workout
PUT /api/v1/workout/{workout_id:[0-9]+}
Params: workout_id: string, name: string, sets: string, reps: string, creator_id: string, exercise_id: string
Return: error

Delete Workout
DELETE /api/v1/workout/{workout_id:[0-9]+}
Params: None
Return: error

## References

Exercises will be populated from [API Ninjas Exercises API](https://www.api-ninjas.com/api/exercises)

[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=thefueley_scholar-power-api&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=thefueley_scholar-power-api)