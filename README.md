Backend for exercise tracking app

Terms:

Exercise - Name of the exercise (e.g push-up, pull-up)
Workout - Instance of an exercise contains an exercise repetitions and sets
Session - Plan for a gym session, contains a list of workouts and a name (e.g Leg day)


Useful commands:
Generate Swagger API 

`oapi-codegen --config=oapi-codegen.yaml pkg/api/v2/openapi.yaml`