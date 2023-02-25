# scholar-power-api

UMGC CMSC 495 API for Scholar Power Workout App
Backend API for [Scholar Power](https://github.com/MoistCode/scholar-power) workout tracking app.

[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=thefueley_scholar-power-api&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=thefueley_scholar-power-api) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=thefueley_scholar-power-api&metric=bugs)](https://sonarcloud.io/summary/new_code?id=thefueley_scholar-power-api) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=thefueley_scholar-power-api&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=thefueley_scholar-power-api) [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=thefueley_scholar-power-api&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=thefueley_scholar-power-api)

## Purpose

Workout Tracker is an application to help people keep track of their workouts.

## Features

Scholar Power is an application to help people keep track of their workouts. Users will login to their account or sign up for a new account. Users will be able to create an account by choosing a unique username and password. Users can then login to workout, create a new workout, or view a history of previous workouts completed.

## Startup Notes

Using the makefile is the easiest way to start the app.

`make run`

Of course, you can always run `go run cmd/server/main.go`

All endpoints that modify the resources will require a valid bearer token.

## Testing

Unit tests have been created.  

`make test`

Testing the API manually is accomplished using Postman or with simple `curl` commands.  

To ease testing with curl, create a user, instructions below. Login with a valid user, instructions below. Then set an env variable with the output of the login command. For example:  

`export TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJpZCI6IjllMjYwYWYwLTkxNzctNDJlNS05ZDkwLTFiYzI4YjMwYjEzOSIsInVpZCI6IjEiLCJpc3N1ZXIiOiJTY2hvbGFyLVBvd2VyIiwidXNlcm5hbWUiOiJ0ZXN0MSIsImlzc3VlZF9hdCI6IjIwMjMtMDItMThUMDA6NTg6NTkuNDk5OTI1LTA1OjAwIiwiZXhwaXJlZF9hdCI6IjIwMjMtMDItMTlUMDA6NTg6NTkuNDk5OTI1LTA1OjAwIn0.Jh6Fzzn1aXMJ8bH0TwkG4ETNwG88cNetoKKQtG2RG5o`

***

### Create a test user

`curl -iL -X POST https://test.seismos.io/api/v1/register -H "Content-Type: application/json" -d '{"username": "user1", "password": "user1"}'`

### Get user info by userid

`curl -iL https://test.seismos.io/api/v1/user/{uid} -H "Content-Type: application/json"`

### Get user info by username

`curl -iL https://test.seismos.io/api/v1/user/{username} -H "Content-Type: application/json"`

### Update password (JWT required)

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/user/{uid} -H "Content-Type: application/json" -d '{"password": "newpassword"}'`

### Delete User (JWT required)

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/user/{uid} -H "Content-Type: application/json"`

### Login

`curl -iL -X POST https://test.seismos.io/api/v1/auth -H "Content-Type: application/json" -d '{"username": "user1", "password": "user1"}'`

***

### Get exercise by id

`curl -iL https://test.seismos.io/api/v1/exercise/1 -H "Content-Type: application/json"`

### Get exercise by name

`curl -iL https://test.seismos.io/api/v1/exercise/name -H "Content-Type: application/json" -d '{"name": "Chest dip"}'`

### Get exercises by muscle group

`curl -iL https://test.seismos.io/api/v1/exercise/lower_back -H "Content-Type: application/json"`

### Get exercises by equipment

`curl -iL https://test.seismos.io/api/v1/exercise/equipment -H "Content-Type: application/json" -d '{"equipment": "barbell"}'`

***

### Create workout plan (JWT required)

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X POST https://test.seismos.io/api/v1/workout -H "Content-Type: application/json" -d '{"uid": "1", "name": "Lower #3", "exercises": [ {"sets": "9", "reps": "9", "load": "135", "exercise_id": "89"}, {"sets": "9", "reps": "9", "load": "135", "exercise_id": "77"}]}'`

### Get workout plans by user (JWT required)  

#### username is from "Get info by id"

`curl -iL -H "Authorization: Bearer ${TOKEN}" https://test.seismos.io/api/v1/workout/{username} -H "Content-Type: application/json"`

### Get exercises in workout plan (JWT required)  

#### plan_id is from above command

`curl -iL -H "Authorization: Bearer ${TOKEN}" https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json"`

### Update workout plan (JWT required)  

#### Send full details of a workout: uid: string, name: string, exercises: {id: string, sets: string, reps: string, load: string, exercise_id: string}

#### If a new exercise is added, no id is required in the exercises element. The exercise_id is still required

#### If request omits an existing exercise, that exercise is dropped from current plan

#### Below request assumes multiple exercises exist in workout plan. Updates exercise 15, adds a new exercise with a new id, and drops all other exercises

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json" -d [{"id": "15","reps": "15","load": "999"},{"reps": "14","load": "888"}]`

### Delete workout plan (JWT required)

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json"`

***

### Complete workout (JWT required)  

#### plan_id is from "Get workouts by user"

`curl -iL -X POST https://test.seismos.io/api/v1/history -H "Content-Type: application/json" -d {"date": "3-Feb-2023", "duration": "55:00", "notes": "Gassed", "plan_id": "4f719e89-0d7d-45a6-9fdf-fec7a5351bfd", "athlete_id": "1"}`

### Get workout history (JWT required)  

#### uid is from "Get user info by username"

`curl -iL -H "Authorization: Bearer ${TOKEN}" https://test.seismos.io/api/v1/history/{uid} -H "Content-Type: application/json"`

### Update workout history (JWT required)  

#### only updating workout **notes** implemented

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/history/{id} -H "Content-Type: application/json" -d {"notes": "so-so", "athlete_id": "1}`

### Delete workout history (JWT required)

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/history/{id} -H "Content-Type: application/json" -d {"athlete_id": "1"}`

***

## API Endpoints

### User Routes

Action|Method-Endpoint|Required|Optional|Response|Notes  
-|--|-|-|-|-
Create User|**POST** /api/v1/register|username: string, password: string|none|token|none
Get User By ID|**GET** /api/v1/user/{id}|none|none|username|none
Get User By Username|**GET** /api/v1/user/{username}|none|none|uid|none
Update User Password|**PUT** /api/v1/user/{id}|password: string|none|message|none
Delete User|**DELETE** /api/v1/user/{id}|none|none|message|none
Login|**POST** /api/v1/auth|username: string, password: string|none|token|none

### Exercise Routes

Action|Method-Endpoint|Required|Optional|Response|Notes  
-|-|-|-|-|-
Get Exercise By ID|**GET** /api/v1/exercise/{id}|none|none|`{id, name, muscle, equipment, instructions}`|none
Get Exercise By Name|**GET** /api/v1/exercise/name|name: string|none|`{id, name, muscle, equipment, instructions}`|Needs Updating. Remove body param.
Get Exercise By Muscle Group|**GET** /api/v1/exercise/{muscle}|none|none|`[{id, name, muscle, equipment, instructions}...{}]`|none
Get Exercise By Equipment|**GET** /api/v1/exercise/equipment|equipment: string|none|`[{id, name, muscle, equipment, instructions}...{}]`|Needs Updating. Remove body param.

### Workout Routes

Action|Method-Endpoint|Required|Optional|Response|Notes  
-|-|-|-|-|-
Create Workout|**POST** /api/v1/workout|creator_id: string|name: string, sets: string, reps: string, load: string, exercise_id: string|message|JWT Required
Get Workouts By User|**GET** /api/v1/workout/user/{username}|none|none|`[{plan_id, name, created_at, edited_at, creator_id},...{}]`|JWT Required
Get Workout Exercises|**GET** /api/v1/workout/{plan_id}|none|none|`[{id, plan_id, name, sets, reps, load, exercise_name, exercise_muscle, exercise_equipment, exercise_instructions, exercise_id},...{}]`|JWT Required
Update Workout|**PUT** /api/v1/workout/{plan_id}|id: string|name: string, sets: string, reps: string, load: string, exercise_name: string|message|JWT Required. **Accepts an array of exercises. Updates workout to match what the request sends.**
Delete Workout|**DELETE** /api/v1/workout/{plan_id}|none|none|message|JWT Required

### Workout History Routes

Action|Method-Endpoint|Required|Optionl|Response|Notes
-|-|-|-|-|-
Create Workout History|**POST** /api/v1/history|date: string, duration: string, plan_id: string, athlete_id: string|notes: string|message|JWT Required
Get Workout History|**GET** /api/v1/history/{id} |none|none|`[{id, date, duration, notes, plan_id, athlete_id},...{}]`|JWT Required
Update Workout History|**PUT** /api/v1/history/{id}|notes:string, athlete_id: string|date: string (not implemented), duration: string (not implemented) |message|JWT Required
Delete Workout History|**DELETE** /api/v1/history/{id}|athlete_id: string|none|message|JWT Required

## References

Exercises will be populated from [API Ninjas Exercises API](https://www.api-ninjas.com/api/exercises)
