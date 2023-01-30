# Attack on College - Sprint 1

## Plan for Sprint 1

### Front-end 

Index landing page, user profile page, registration and login pages to be built

### Back-end

Structs for User, Course, Assignment

REST api functions for user registration, login, profile

Gorm Databases need to be set up for Users, Assignments, etc. 

Currently, gin-gonic/gin package is used instead of gorilla/mux, rs/cors, net/http packages.

### User stories:

# Back-end progress

## 01/30

The AttackOnCollege_v0.0.1/src directory contains several parts of back-end functionality. So far, the following properties are implemented:

### User Struct and Database

The models package contains a struct called User which represents the profile that the website user will have. The struct has the following fields:
1. Username: website handle that will be displayed on the profile, and will be used for looking up other users and displaying the global status on the ranked list of users. Gorm will require all emails in the database to be unique
2. Email: a method of identification that will be used for logging in. Gorm will require all emails in the database to be unique
3. Password: another part of identification
4. First name and last name: Personalization of the user
5. Current course: the code of the course that the user is currently taking. This is used to look up the course the student is taking in the database and access the assignments they have to take
6. Achievements: this is a list of IDs of the achievements the user has made. The list will be incremented as the user completes assignments and courses, and will provide additional experience points that will go toward levels.
7. Level: this is the current level that the user has in our app. It is incremented based on a scale we are yet to determine, and it will be primarily used for additonal achievements and bonus experience points. It is supposed to present small step-by-step goals for users that would keep their interest and give them instant gratification that is neccessary when playing any game, and completing a class
8. Experience points: the total number of experience points will be used to build a leaderboard of users used for competition

So far, functions that edit First and last name, along with the username of the user are created, along with a function that can change the password. Both of these can only be done if the user is logged in. Getter functions are implemented so that logged in users can access a list of all users, their own profile information, and delete their account.

### Course Struct and Database

The course struct has the following fields:
1. Title: course title that is just there for an easier access by the user
2. Course code: unique course code that the User struct has access to. That is the key that is looked up in the database of courses and connects the two databases
3. Assignments: a list of unique IDs of assignments the user created. The list is a container of handles that can be then used in the database of assignments to access the specific user's assignments and modify them.
4. Final grade: the final grade earned in the course. It is used to calculate the experience points that will be earned after completing the course. 
5. Experience points: the number of experience points that will be added to user once they complete the course

So far, only a function that creates the course and adds it to the database, along with linking it to the user is created. Since the function for completing assignments it already created, it would be easy to use that function as a model.

### Assignment Struct and Database

The assignment struct has the following fields:
1. Title: the title of the assignment that is to be completed
2. Description: the description of the assignment. This can be the requirements for the assignment, details the professor gave, type of assignment, etc. It can also include the deadline since it should be displayed with the title
3. Number of Points: the total number of points that can be earned for the assignment
4. Weight: the grade weight. It is the percent of the final grade that this assignment will contribute to.
5. Experience points: Total experience points earned upon completion of the assignment
6. Points earned: the number of points that the user achieved upon completion

Two functions for assignment manipulation developed so far: one to create, and one to complete assignments. The function that create assignments links the new assignment created in the database to the course it is in, and the user. The function that completes the assignment also calculated the total number of experience points to be added to the user.

### Achievement Struct and Database

So far, only the basic fields of the achievement struct are made such as title, description and experience points. 

### Server setup

The server is set up with the gin-gonic/gin (https://github.com/gin-gonic/gin) package that acts as a router. I have been experimenting with the TLS encryption of servers, so the server.go file also routes all http requests to port 8080 to the https server on port 1337. In that spirit, all the getter functions do not return the password hash of the user profiles. All passwords are replaced with "Hidden" before being sent in the reply.

### Databases

Databases are set up using GORM (https://gorm.io/gorm) package for GoLang, and there are 4 so far. They are all local, and I'm still looking for ways to decentralize that, although it might be pointless for the low scale project such as this one.

### Controllers

Controllers is a package I split up from the models package and it includes files with just functions that describe behaviors and are mostly handlers for routing.

### Rest

The REST plugin for VS Code is used for testing and sending requests to the server. The src/rest directory has the list of .rest files that are used as tests for behaviors and debugging. 

### Total list of packages used in the project so far

    go get -u gorm.io/gorm

    go get -u gorm.io/driver/sqlite

    go get -u github.com/gin-gonic/gin

    go get -u golang.org/x/crypto

    go get -u github.com/golang-jwt/jwt