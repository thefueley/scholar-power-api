# scholar-power-api

UMGC CMSC 495 API for Scholar Power Workout App
Backend API for [Scholar Power](https://github.com/MoistCode/scholar-power) workout tracking app.

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

To ease testing with curl, login with a valid user, then set an env variable with the output of the login command. For example:

`export TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjllMjYwYWYwLTkxNzctNDJlNS05ZDkwLTFiYzI4YjMwYjEzOSIsInVpZCI6IjEiLCJpc3N1ZXIiOiJTY2hvbGFyLVBvd2VyIiwidXNlcm5hbWUiOiJ0ZXN0MSIsImlzc3VlZF9hdCI6IjIwMjMtMDItMThUMDA6NTg6NTkuNDk5OTI1LTA1OjAwIiwiZXhwaXJlZF9hdCI6IjIwMjMtMDItMTlUMDA6NTg6NTkuNDk5OTI1LTA1OjAwIn0.Jh6Fzzn1aXMJ8bH0TwkG4ETNwG88cNetoKKQtG2RG5o`

Create a test user
`curl -iL -X POST https://test.seismos.io/api/v1/register -H "Content-Type: application/json" -d '{"username": "user1", "password": "user1"}'`

Get user info by userid
`curl -iL https://test.seismos.io/api/v1/user/{uid} -H "Content-Type: application/json"`

Get user info by username
`curl -iL https://test.seismos.io/api/v1/user/{username} -H "Content-Type: application/json"`

Update password (JWT required)
`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/user/{uid} -H "Content-Type: application/json" -d '{"password": "newpassword"}'`

Delete User (JWT required)
`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/user/{uid} -H "Content-Type: application/json"`

Login
`curl -iL -X POST https://test.seismos.io/api/v1/auth -H "Content-Type: application/json" -d '{"username": "user1", "password": "user1"}'`

Get exercise by id
`curl -iL https://test.seismos.io/api/v1/exercise/1 -H "Content-Type: application/json"`

Get exercise by name
`curl -iL https://test.seismos.io/api/v1/exercise/name -H "Content-Type: application/json" -d '{"name": "Chest dip"}'`

Get exercises by muscle group
`curl -iL https://test.seismos.io/api/v1/exercise/lower_back -H "Content-Type: application/json"`

Get exercises by equipment
`curl -iL https://test.seismos.io/api/v1/exercise/equipment -H "Content-Type: application/json" -d '{"equipment": "barbell"}'`

Create workout plan (JWT required)
`curl -iL -H "Authorization: Bearer ${TOKEN}" -X POST https://test.seismos.io/api/v1/workout -H "Content-Type: application/json" -d '{"uid": "1", "name": "Lower #3", "exercises": [ {"sets": "9", "reps": "9", "load": "135", "exercise_id": "89"}, {"sets": "9", "reps": "9", "load": "135", "exercise_id": "77"}]}'`

Get workout plans by user (JWT required) (username is from "Get info by id")
`curl -iL -H "Authorization: Bearer ${TOKEN}" https://test.seismos.io/api/v1/workout/{username} -H "Content-Type: application/json"`

Get exercises in workout plan (JWT required) (plan_id is from above command)
`curl -iL -H "Authorization: Bearer ${TOKEN}" https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json"`

Update workout plan (JWT required) (Only id is required. All other params optional. id's in -d must come from above command)
`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json" -d [{"id": "15","reps": "15","load": "999"},{"id": "16","reps": "14","load": "888"}]`

Delete workout plan (JWT required)
`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/workout/{plan_id} -H "Content-Type: application/json"`

Complete workout (JWT required) (plan_id is from "Get workouts by user")
`curl -iL -X POST https://test.seismos.io/api/v1/history -H "Content-Type: application/json" -d {"date": "3-Feb-2023", "duration": "55:00", "notes": "Gassed", "plan_id": "4f719e89-0d7d-45a6-9fdf-fec7a5351bfd", "athlete_id": "1"}`

Get workout history (JWT required) (uid is from "Get user info by username")
`curl -iL -H "Authorization: Bearer ${TOKEN}" https://test.seismos.io/api/v1/history/{uid} -H "Content-Type: application/json"`

Update workout history (JWT required) (only updating workout notes implemented)
`curl -iL -H "Authorization: Bearer ${TOKEN}" -X PUT https://test.seismos.io/api/v1/history -H "Content-Type: application/json" -d {"notes": "so-so", "plan_id": "4f719e89-0d7d-45a6-9fdf-fec7a5351bfd"}`

Delete workout history (JWT required)
`curl -iL -H "Authorization: Bearer ${TOKEN}" -X DELETE https://test.seismos.io/api/v1/history -H "Content-Type: application/json" -d {"plan_id": "4f719e89-0d7d-45a6-9fdf-fec7a5351bfd"}`

## API Endpoints

### User Routes

Action | Method-Endpoint | Params | Required | Response 
--- | -------| --------------| ---------- | ------------
Create User | POST /api/v1/user |  | username: string, password: string  |  Return: token or error "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImRjMWU4OWRiLTIxZGQtNDM1ZS04Yzk0LTA3YWIyNzMwOWUxMiIsImlzc3VlciI6IlNjaG9sYXItUG93ZXIiLCJ1c2VybmFtZSI6InVzZXIzIiwiaXNzdWVkX2F0IjoiMjAyMy0wMi0xMFQxNTo0OToyOS45NDEwMDUtMDU6MDAiLCJleHBpcmVkX2F0IjoiMjAyMy0wMi0xMFQxNjowNDoyOS45NDEwMDUtMDU6MDAifQ.bz5PUFR2a_JfXgOC0vCFkGspDHhu4-eMCoRqQHeASJA"
Get User By ID | GET /api/v1/user/{id:[0-9]+} | none | | Return: User or error <br /> {<br />"ID": "1", <br /> "UserName": "user1", <br /> "PasswordHash": "$2a$10$9IxVao19OqCVj9No1lySxupoM7Njl2jgxY6Artr4QSzvbdXel0feG"<br />}
Get User By Username | GET /api/v1/user/{username} | none | | Return: User or error <br /> {<br />"ID": "1", <br /> "UserName": "user1", <br /> "PasswordHash": "$2a$10$9IxVao19OqCVj9No1lySxupoM7Njl2jgxY6Artr4QSzvbdXel0feG"<br />}
Update User Password | PUT /api/v1/user/{id:[0-9]+} || (auth required), password: string | Return: error <br /> {<br /> "username": "", <br /> "password": "user33"<br /> }
Delete User | DELETE /api/v1/user/{id:[0-9]+} | none (auth required) || Return: error <br />{ <br /> "Message": "Poof, it's gone!" <br /> }
Login | POST /api/v1/auth || username: string, password: string | Return: token or error <br /> "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjcwM2ZjYzI4LTg5MjQtNDNmMi05YzRiLTBlMDNmY2E3N2NlOCIsImlzc3VlciI6IlNjaG9sYXItUG93ZXIiLCJ1c2VybmFtZSI6InVzZXIzIiwiaXNzdWVkX2F0IjoiMjAyMy0wMi0xMFQxNTo1MjoxMC4wNDIwNTItMDU6MDAiLCJleHBpcmVkX2F0IjoiMjAyMy0wMi0xMFQxNjowNzoxMC4wNDIwNTItMDU6MDAifQ.U22MN2UGO2K6-y9Xw6nocoq7QYoJTPhLlx6V2MldloM"


### Exercise Routes

Action | Method-Endpoint | Params | Required| Response 
--- | -------| --------------| ------------| ------------
Get Exercise By ID | GET /api/v1/exercise/{id:[0-9]+} | none || Return : exercise or error <br /> { <br /> "ID": "1", <br /> "Name": "Landmine twist",<br /> "Muscle": "abdominals", <br /> "Equipment": "other", <br /> "Instructions": "Position a bar into a landmine or securely anchor it in a corner. Load the bar to an appropriate weight. Raise the bar from the floor, taking it to shoulder height with both hands with your arms extended in front of you. Adopt a wide stance. This will be your starting position. Perform the movement by rotating the trunk and hips as you swing the weight all the way down to one side. Keep your arms extended throughout the exercise. Reverse the motion to swing the weight all the way to the opposite side. Continue alternating the movement until the set is complete." <br /> }
Get Exercise By Name | GET /api/v1/exercise/name | none | name: string | Return: array of exercises or error <br /> [ <br /> {<br /> "ID": "58", <br /> "Name": "Chest dip", <br />"Muscle": "chest", <br /> "Equipment": "other", <br /> "Instructions": "For this exercise you will need access to parallel bars. To get yourself into the starting position, hold your body at arms length (arms locked) above the bars. While breathing in, lower yourself slowly with your torso leaning forward around 30 degrees or so and your elbows flared out slightly until you feel a slight stretch in the chest. Once you feel the stretch, use your chest to bring your body back to the starting position as you breathe out. Tip: Remember to squeeze the chest at the top of the movement for a second. Repeat the movement for the prescribed amount of repetitions.  Variations: If you are new at this exercise and do not have the strength to perform it, use a dip assist machine if available. These machines use weight to help you push your bodyweight. Otherwise, a spotter holding your legs can help. More advanced lifters can add weight to the exercise by using a weight belt that allows the addition of weighted plates." <br /> } <br /> ]
Get Exercise By Muscle Group | GET /api/v1/exercise/muscle | none | muscle: string | Return: array of exercises or error <br /> [ <br /> { <br /> "ID": "31", <br /> "Name": "Incline Hammer Curls", <br /> "Muscle": "biceps", <br /> "Equipment": "dumbbell", <br /> "Instructions": "Seat yourself on an incline bench with a dumbbell in each hand. You should pressed firmly against he back with your feet together. Allow the dumbbells to hang straight down at your side, holding them with a neutral grip. This will be your starting position. Initiate the movement by flexing at the elbow, attempting to keep the upper arm stationary. Continue to the top of the movement and pause, then slowly return to the start position." <br /> }, ... <br /> ]
Get Exercise By Equipment | GET /api/v1/exercise/equipment | none | equipment: string | Return: array of exercises or error <br /> [ <br /> { <br /> "ID": "32", <br /> "Name": "Wide-grip barbell curl", <br /> "Muscle": "biceps", <br /> "Equipment": "barbell", <br /> "Instructions": "Stand up with your torso upright while holding a barbell at the wide outer handle. The palm of your hands should be facing forward. The elbows should be close to the torso. This will be your starting position. While holding the upper arms stationary, curl the weights forward while contracting the biceps as you breathe out. Tip: Only the forearms should move. Continue the movement until your biceps are fully contracted and the bar is at shoulder level. Hold the contracted position for a second and squeeze the biceps hard. Slowly begin to bring the bar back to starting position as your breathe in. Repeat for the recommended amount of repetitions.  Variations:  You can also perform this movement using an E-Z bar or E-Z attachment hooked to a low pulley. This variation seems to really provide a good contraction at the top of the movement. You may also use the closer grip for variety purposes." <br />}, ... <br /> ]

### Workout Routes

Action | Method-Endpoint | Params | Required| Response 
--- | -------| --------------| ------------| ------------
Create Workout | POST /api/v1/workout | (auth required) | creator_id: string <br /> Optional: name: string, sets: string, reps: string, load: string, plan_id: string, exercise_id: string, instructions_id: string | Return: message <br /> { <br /> "Message": "workout created" <br /> }
Get Workouts By User | GET /api/v1/workout/user/{username} | (auth required) | none |Return: list of workouts or error <br /> [ <br /> { <br />   "PlanID": "36", <br /> "Name": "Upper", <br /> "CreatedAt": "2023-02-10 21:01:09.000", <br /> "EditedAt": "2023-02-10 21:01:09.000", <br /> "CreatorID": "4" <br /> } <br /> ]
Get Workout Exercises | GET /api/v1/workout/{plan_id} | (auth required) | none | Return: list of exercises in a workout or error <br /> [ <br /> { <br /> "ID": "7", <br /> "PlanID": "36", <br /> "Name": "Upper", <br /> "Sets": "3", <br /> "Reps": "8", <br /> "Load": "135", <br /> "ExerciseName": "Low-cable cross-over", <br /> "ExerciseMuscle": "chest", <br /> "ExerciseEquipment": "cable", <br /> "ExerciseInstructions": "To move into the starting position, place the pulleys at the low position, select the resistance to be used and grasp a handle in each hand. Step forward, gaining tension in the pulleys. Your palms should be facing forward, hands below the waist, and your arms straight. This will be your starting position. With a slight bend in your arms, draw your hands upward and toward the midline of your body. Your hands should come together in front of your chest, palms facing up. Return your arms back to the starting position after a brief pause." <br /> } <br /> ]
Update Workout | PUT /api/v1/workout/{plan_id} | (auth required) | id: string | Return: message <br /> { <br /> "Message": "workout updated" <br /> }
Delete Workout | DELETE /api/v1/workout/{plan_id} | (auth required) | none | Return: message <br /> { <br /> "Message": "workout deleted" <br /> }

### Workout History Routes

Action | Method-Endpoint | Params | Required| Response 
--- | -------| --------------| ------------| ------------
Create Workout History | POST /api/v1/history | (auth required) | date: string, duration: string, plan_id: string, athlete_id: string <br /> Optional: notes: string | Return: message or error <br /> {<br /> "Message":"History created"<br />}
Get Workout History | GET /api/v1/history/{id} | none | (auth required) |Return: message or error <br />[{<br />"ID":"1","Date":"1-Feb-2023","Duration":"55:00","Notes":"Gassed","PlanID":"1","AthleteID":"1"},<br />{"ID":"2","Date":"2-Feb-2023","Duration":"51:00","Notes":"Ok","PlanID":"1","AthleteID":"1"},<br />{"ID":"3","Date":"3-Feb-2023","Duration":"52:00","Notes":"Great","PlanID":"2","AthleteID":"1"},<br />{"ID":"4","Date":"4-Feb-2023","Duration":"50:00","Notes":"Great","PlanID":"2","AthleteID":"1"<br />}]
Update Workout History | PUT /api/v1/history/{id} | (auth required) | notes: string <br /> Optional: date: string (not implemented yet), duration: string (not implemented yet) | Return: message or error <br /> {<br />"Message":"workout history updated"<br />}
Delete Workout History | DELETE /api/v1/history/{id} | (auth required) | none | Return: message or error <br /> {<br />"Message":"Poof! It's gone."<br />}

## References

Exercises will be populated from [API Ninjas Exercises API](https://www.api-ninjas.com/api/exercises)

[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=thefueley_scholar-power-api&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=thefueley_scholar-power-api)
