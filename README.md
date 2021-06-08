# Social-Media-app
POC for api writen in go using goarm



## Problem statement
Let's create a mock social media app. We'll have users and profiles. Each user will be linked to a profile. Profile can't exist without a user
You'll need to support 4 APIs

* POST /user -> it will take user details in JSON (Name/Email/Address), on save in the database you'll generate a ID for the user
* POST /user/{generated_user_id}/profile -> creates a profile for a user containing his hobbies, languages he knows
* GET /user -> Get all users
* GET /profile -> Get all profile
* GET /user/{user_id}/profile -> profile of a particular user

Problem contributed by @rajat-godi
