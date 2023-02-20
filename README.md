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

### Get exercise by id

`curl -iL https://test.seismos.io/api/v1/exercise/1 -H "Content-Type: application/json"`

### Get exercise by name

`curl -iL https://test.seismos.io/api/v1/exercise/name -H "Content-Type: application/json" -d '{"name": "Chest dip"}'`

### Get exercises by muscle group

`curl -iL https://test.seismos.io/api/v1/exercise/lower_back -H "Content-Type: application/json"`

### Get exercises by equipment

`curl -iL https://test.seismos.io/api/v1/exercise/equipment -H "Content-Type: application/json" -d '{"equipment": "barbell"}'`

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

### Below request assumes exercise id's: 15, 16, 17 exist. Request updates exercise 15, adds a new exercise with a new id, and drop exercises 16 and 17

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json" -d [{"id": "15","reps": "15","load": "999"},{"reps": "14","load": "888"}]`

### Delete workout plan (JWT required)

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json"`

### Complete workout (JWT required)  

#### plan_id is from "Get workouts by user"

`curl -iL -X POST https://test.seismos.io/api/v1/history -H "Content-Type: application/json" -d {"date": "3-Feb-2023", "duration": "55:00", "notes": "Gassed", "plan_id": "4f719e89-0d7d-45a6-9fdf-fec7a5351bfd", "athlete_id": "1"}`

### Get workout history (JWT required)  

#### uid is from "Get user info by username"

`curl -iL -H "Authorization: Bearer ${TOKEN}" https://test.seismos.io/api/v1/history/{uid} -H "Content-Type: application/json"`

### Update workout history (JWT required)  

#### only updating workout **notes** implemented

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/history -H "Content-Type: application/json" -d {"notes": "so-so", "plan_id": "4f719e89-0d7d-45a6-9fdf-fec7a5351bfd"}`

### Delete workout history (JWT required)

`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/history -H "Content-Type: application/json" -d {"plan_id": "4f719e89-0d7d-45a6-9fdf-fec7a5351bfd"}`

## API Endpoints

### User Routes

Action|Method-Endpoint|Required|Optional|Response|Notes  
-|--|-|-|-|-
Create User|**POST** /api/v1/user|username: string, password: string|none|`"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImRjMWU4OWRiLTIxZGQtNDM1ZS04Yzk0LTA3YWIyNzMwOWUxMiIsImlzc3VlciI6IlNjaG9sYXItUG93ZXIiLCJ1c2VybmFtZSI6InVzZXIzIiwiaXNzdWVkX2F0IjoiMjAyMy0wMi0xMFQxNTo0OToyOS45NDEwMDUtMDU6MDAiLCJleHBpcmVkX2F0IjoiMjAyMy0wMi0xMFQxNjowNDoyOS45NDEwMDUtMDU6MDAifQ.bz5PUFR2a_JfXgOC0vCFkGspDHhu4-eMCoRqQHeASJA"`|none
Get User By ID|**GET** /api/v1/user/{id}|none|none|`{"ID": "1", "UserName": "user1", "PasswordHash": "$2a$10$9IxVao19OqCVj9No1lySxupoM7Njl2jgxY6Artr4QSzvbdXel0feG"}`|Needs Updating. Remove hash response.
Get User By Username|**GET** /api/v1/user/{username}|none|none|`{"ID": "1", "UserName": "user1", "PasswordHash": "$2a$10$9IxVao19OqCVj9No1lySxupoM7Njl2jgxY6Artr4QSzvbdXel0feG"}`|Needs Updating. Remove hash response.
Update User Password|**PUT** /api/v1/user/{id}|password: string|none|`{"username": "","password": "user33"}`|Needs Updating. Send message instead.
Delete User|**DELETE** /api/v1/user/{id}|none|none|`{"Message": "Poof! It's gone."}`|none
Login|**POST** /api/v1/auth|username: string, password: string|none|`"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjcwM2ZjYzI4LTg5MjQtNDNmMi05YzRiLTBlMDNmY2E3N2NlOCIsImlzc3VlciI6IlNjaG9sYXItUG93ZXIiLCJ1c2VybmFtZSI6InVzZXIzIiwiaXNzdWVkX2F0IjoiMjAyMy0wMi0xMFQxNTo1MjoxMC4wNDIwNTItMDU6MDAiLCJleHBpcmVkX2F0IjoiMjAyMy0wMi0xMFQxNjowNzoxMC4wNDIwNTItMDU6MDAifQ.U22MN2UGO2K6-y9Xw6nocoq7QYoJTPhLlx6V2MldloM"`|none

### Exercise Routes

Action|Method-Endpoint|Required|Optional|Response|Notes  
-|-|-|-|-|-
Get Exercise By ID|**GET** /api/v1/exercise/{id}|none|none|`{"ID": "1","Name": "Landmine twist","Muscle": "abdominals","Equipment": "other","Instructions": "Position a bar into a landmine or securely anchor it in a corner. Load the bar to an appropriate weight. Raise the bar from the floor, taking it to shoulder height with both hands with your arms extended in front of you. Adopt a wide stance. This will be your starting position. Perform the movement by rotating the trunk and hips as you swing the weight all the way down to one side. Keep your arms extended throughout the exercise. Reverse the motion to swing the weight all the way to the opposite side. Continue alternating the movement until the set is complete."}`|none
Get Exercise By Name|**GET** /api/v1/exercise/name|name: string|none|`[{"ID": "58", "Name": "Chest dip", "Muscle": "chest", "Equipment": "other", "Instructions": "For this exercise you will need access to parallel bars. To get yourself into the starting position, hold your body at arms length (arms locked) above the bars. While breathing in, lower yourself slowly with your torso leaning forward around 30 degrees or so and your elbows flared out slightly until you feel a slight stretch in the chest. Once you feel the stretch, use your chest to bring your body back to the starting position as you breathe out. Tip: Remember to squeeze the chest at the top of the movement for a second. Repeat the movement for the prescribed amount of repetitions.  Variations: If you are new at this exercise and do not have the strength to perform it, use a dip assist machine if available. These machines use weight to help you push your bodyweight. Otherwise, a spotter holding your legs can help. More advanced lifters can add weight to the exercise by using a weight belt that allows the addition of weighted plates."}]`|Needs Updating. Remove body param.
Get Exercise By Muscle Group|**GET** /api/v1/exercise/{muscle}|none|none|`[{"ID": "31","Name": "Incline Hammer Curls", "Muscle": "biceps", "Equipment": "dumbbell", "Instructions": "Seat yourself on an incline bench with a dumbbell in each hand. You should pressed firmly against he back with your feet together. Allow the dumbbells to hang straight down at your side, holding them with a neutral grip. This will be your starting position. Initiate the movement by flexing at the elbow, attempting to keep the upper arm stationary. Continue to the top of the movement and pause, then slowly return to the start position."},...]`|none
Get Exercise By Equipment|**GET** /api/v1/exercise/equipment|equipment: string|none|`[{"ID": "32", "Name": "Wide-grip barbell curl", "Muscle": "biceps", "Equipment": "barbell", "Instructions": "Stand up with your torso upright while holding a barbell at the wide outer handle. The palm of your hands should be facing forward. The elbows should be close to the torso. This will be your starting position. While holding the upper arms stationary, curl the weights forward while contracting the biceps as you breathe out. Tip: Only the forearms should move. Continue the movement until your biceps are fully contracted and the bar is at shoulder level. Hold the contracted position for a second and squeeze the biceps hard. Slowly begin to bring the bar back to starting position as your breathe in. Repeat for the recommended amount of repetitions.  Variations:  You can also perform this movement using an E-Z bar or E-Z attachment hooked to a low pulley. This variation seems to really provide a good contraction at the top of the movement. You may also use the closer grip for variety purposes."},...]`|Needs Updating. Remove body param.

### Workout Routes

Action|Method-Endpoint|Required|Optional|Response|Notes  
-|-|-|-|-|-
Create Workout|**POST** /api/v1/workout|creator_id: string|name: string, sets: string, reps: string, load: string, exercise_id: string|`{"Message": "workout created"}`|JWT Required
Get Workouts By User|**GET** /api/v1/workout/user/{username}|none|none|`["PlanID": "36","Name": "Upper","CreatedAt": "2023-02-10 21:01:09.000","EditedAt": "2023-02-10 21:01:09.000","CreatorID": "4"}]`|JWT Required
Get Workout Exercises|**GET** /api/v1/workout/{plan_id}|none|none|`[{"ID": "7","PlanID": "36","Name": "Upper","Sets": "3", "Reps": "8", "Load": "135","ExerciseName": "Low-cable cross-over","ExerciseMuscle": "chest", "ExerciseEquipment": "cable", "ExerciseInstructions": "To move into the starting position, place the pulleys at the low position, select the resistance to be used and grasp a handle in each hand. Step forward, gaining tension in the pulleys. Your palms should be facing forward, hands below the waist, and your arms straight. This will be your starting position. With a slight bend in your arms, draw your hands upward and toward the midline of your body. Your hands should come together in front of your chest, palms facing up. Return your arms back to the starting position after a brief pause."}]`|JWT Required
Update Workout|**PUT** /api/v1/workout/{plan_id}|id: string|name: string, sets: string, reps: string, load: string, exercise_name: string|`{"Message": "workout updated"}`|JWT Required. Accepts an array of exercises. Updates workout to match what the request sends.
Delete Workout|**DELETE** /api/v1/workout/{plan_id}|none|none|`{"Message": "Poof! It's gone."}`|JWT Required

### Workout History Routes

Action|Method-Endpoint|Required|Optionl|Response|Notes
-|-|-|-|-|-
Create Workout History|**POST** /api/v1/history|date: string, duration: string, plan_id: string, athlete_id: string|notes: string|`{"Message":"History created"}`|JWT Required
Get Workout History|**GET** /api/v1/history/{id} |none|none|`[{"ID":"1","Date":"1-Feb-2023","Duration":"55:00","Notes":"Gassed","PlanID":"1","AthleteID":"1"},{"ID":"2","Date":"2-Feb-2023","Duration":"51:00","Notes":"Ok","PlanID":"1","AthleteID":"1"},{"ID":"3","Date":"3-Feb-2023","Duration":"52:00","Notes":"Great","PlanID":"2","AthleteID":"1"},{"ID":"4","Date":"4-Feb-2023","Duration":"50:00","Notes":"Great","PlanID":"2","AthleteID":"1"}]`|JWT Required
Update Workout History|**PUT** /api/v1/history/{id}|notes:string|date: string (not implemented), duration: string (not implemented) |`{"Message":"workout history updated"}`|JWT Required
Delete Workout History|**DELETE** /api/v1/history/{id}|none|none|`{"Message":"Poof! It's gone."}`|JWT Required

## References

Exercises will be populated from [API Ninjas Exercises API](https://www.api-ninjas.com/api/exercises)
