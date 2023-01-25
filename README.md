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

## References

Exercises will be populated from [API Ninjas Exercises API](https://www.api-ninjas.com/api/exercises)
