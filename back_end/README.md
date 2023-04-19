# Attack On College v0.0.1

This is a general documentation for the back-end of Attack on College version 0.0.1

# Functionality

The implemented functionality:

## Linking front-end and back-end

As of now, the front-end and back-end are linked up. Starting the server and Angular allows the two to communicate. GoLang needs to be installed and set up for execution on the machine to properly start the server. If both Angular and Go are installed, running 
    
    npm start 
    
in the console of the AttackOnCollege directory will start both up. For testing purposes, running 

    go run ./back_end/src/*.go 

will start up just the server. Although Codegangsta Gin is NOT used to start the server, it does give testing feedback in the console. The reason for not running with Gin is that Gin is very verbose and it is hard to read through all the ouput. For debugging purposes, Delve can be used which is very easy to install and set up. For instructions on setting up Delve with VS Code, check the Delve setup section at the bottom of the document.

## Register user

Along with the respective front-end functionality, the router and handler functions for registering a user into the database are fully implemented and functional. In Sprint 1, registering the user required JSON input from the POST requests, but that has been updated to take in any object with the matching fields (firstName, lastName, email, username, password, dob, major, college). As long as the object passed can be bound to the request object in the backend, and it passes GORM requirements, the new user will be created. 

When the user is created, the password string is hashed and the hash is placed into the database. This hash will directly be used to compare the password during login. 

Once the user is added to the database with no errors, a JSON object containing the email and username of the newly created user is passed back to the front-end as a confirmation method. This needs to be updated at this stage to make it more proctical for the front-end to conduct tests.

### Notes on Date of Birth

The date of birth has finally been fully updated and is stored in the format "MMM DD YYYY". This format can be used in Golang to convert the string into the time.Time variable which has a lot of important functionalities that are used to calculate the user age. The same function for date of birth formatting is used for formatting the due dates of assignments. This can also be used to assign priority to assignments based on the urgency of the due date and have them potentially displayed for the user to see assignments with the closest due dates. 

### Ideas for due dates

Change due date formatting to store time due. Add a function to obtain urgency (difference between time.Now() and due date including time of the day difference) and then store it in the Assignment model. Then run through the assignments and get the top five. Optimization of this part is neccessary.

## Login user

The logging in functionality is fully implemented and functional with full integration with the front-end. Users can now log in with their username and password, and will not be able to do so with a wrong password. The password the users input to log in is hashed and checked against the hash stored in the database. Once that is determined a success (if it's successful), the token is stored in the database to be used for navigation and modification. It is automatically validated periodically since the JWT expires after an hour. 

## Create course and assignment

These functions work in a similar fashion as the register user function. The only difference is that the request will have to contain an "Authorization" header with the current token. If this token is expired, an error message will be sent back. If the passed token cannot be found in the database, an error message will be sent back. If the token is not passed, an error message will be sent back. This is also now fully integrated with the front-end.

The assignments can be completed, too, and the achievement for completing the first assignment is created, so it can be given once the user completes their first assignment, although this has not yet been implemented on the front-end.

The getter functions and routers have been set up. The user can now get basic information about all of their assignments (name, due date, course) and pick an assignment to get more information for (name, description, number of points, due date, course, final grade weight, etc.). This has not been implemented on the front-end yet.

## Achievements

Some checkers are added to the code to give user's achievements (probably should be set as a separate function that is called every time a user does something relevant on the profile such as add and complete assignments and courses, edits an assignment, etc.). A function and router to return all achievements a user has obtained has been implemented and fully integrated with the front-end. 

The array of achievements that is sent back has the first element which is the number of achievements a user has. This is stored in the "ExperiencePoints" field of the Achievement model and allows the front-end to read the size of the array sent from the server to iterate through and display for the user. This could potentially be explored for a better/more efficient implementation instead of sending an empty object as the first element of the array. 

## Admin and Testing accounts

For the purposes of testing the achievement and assignment functionality, the back-end calls a helper function which creates two important accounts. One is an administrator account (AOCAdmin) which will have very high privileges on our site. This will allow the admin to add achievements from the front-end, edit users, delete users (site moderator functionality), and increase the convenience of adding information into databases that don't have to be hardcoded. 

The other account is the testing account (AOCTest) which has several achievements hardcoded into the profile. This account is used to perform testing on more advanced functionalities such as displaying achievements without having to do all the steps to get the achievement which can be tedious. 

# Structs and packages

## User Struct and Database

The models package contains a struct called User which represents the profile that the website user will have. The struct has the following fields:

1. Username: website handle that will be displayed on the profile, and will be used for looking up other users and displaying the global status on the ranked list of users. Gorm will require all emails in the database to be unique
2. Email: a method of identification that will be used for logging in. Gorm will require all emails in the database to be unique
3. Password: another part of identification
4. First name and last name: Personalization of the user
5. Date of Birth: user's date of birth. Could potentially be used to sent Birthday notifications
6. Age: more personalized information for the user
7. College and Major: these two will be used to potentially organize the leaderboards based on the user's environment. This way they will be able to compete against peers in their major and/or college.
8. Current course: the code of the course that the user is currently taking. This is used to look up the course the student is taking in the database and access the assignments they have to take
9. Course ID: the ID number of the course the user is taking. This is going to be used to implement user being able to track more than one course at a time
10. Achievements: this is a list of IDs of the achievements the user has made. The list will be incremented as the user completes assignments and courses, and will provide additional experience points that will go toward levels.
11. Assignments: this is a list of IDs of assignments the user had to do. The list is used to navigate the Assignments database and obtain details about assignments they have due. Since there are no unique restraints other than the ID, several users and the same user can have different assignments share the same information, but IDs will be different and each user will only be able to access their own assignments. Since the ID is stored into the array the moment the Assignment is added, each user has access to the IDs of only the assignments they added. Same applies to courses they add, so only they can access their final grade for the course.
12. Level: this is the current level that the user has in our app. It is incremented based on a scale we are yet to determine, and it will be primarily used for additonal achievements and bonus experience points. It is supposed to present small step-by-step goals for users that would keep their interest and give them instant gratification that is neccessary when playing any game, and completing a class
13. Experience points: the total number of experience points will be used to build a leaderboard of users used for competition

## Course Struct and Database

The course struct has the following fields:

1. Title: course title that is just there for an easier access by the user
2. Course code: the code that user's university uses for a course
4. Final grade: the final grade earned in the course. It is used to calculate the experience points that will be earned after completing the course. 
5. Experience points: the number of experience points that will be added to user once they complete the course 
6. Gorm.Model has the field ID which is used to link users to the courses they are taking

## Assignment Struct and Database

The assignment struct has the following fields:

1. Title: the title of the assignment that is to be completed
2. Description: the description of the assignment. This can be the requirements for the assignment, details the professor gave, type of assignment, etc.
3. Due date: the due date of the assignment that can be used to display the most urgent assignments for the user
4. Assignment type: homework, quiz, exam, paper, etc.
5. Number of Points: the total number of points that can be earned for the assignment
6. Weight: the grade weight. It is the percent of the final grade that this assignment will contribute to.
7. Experience points: Total experience points earned upon completion of the assignment
8. Points earned: the number of points that the user achieved upon completion

## Achievement Struct and Database

So far, only the basic fields of the achievement struct are made such as title, description and experience points.

1. Title: name of the achievement
2. Description: a message for the user that can contain information about how to obtain the achievement
3. Experience points: the number of experience points a user will gain when they obtain the achevement

## Server setup

The server is set up with the [gin-gonic/gin package](https://github.com/gin-gonic/gin) that acts as a router. The TLS encryption was providing difficulty in linking with the front-end, so that has been taken out and replaced with just HTTP communication. Nevertheless, all the getter functions do not return the password hash of the user profiles.

## Databases

Databases are set up using [GORM package](https://gorm.io/gorm) for GoLang, and there are 4 so far. They are all local, and we are still looking for ways to decentralize that, although it might be pointless for the low scale project such as this one. PostgreSQL didn't work the way we expected it, and mySQL doesn't allow non-primitive datatypes which we need for the arrays of achievements, assignments, courses, etc.

## Controllers

Controllers is a package we split up from the models package, and it includes files with just functions that describe behaviors and are mostly handlers for routing.

### User controllers

These are functions mainly used for registering the user. This includes creating the account and formatting the date of birth, as well as calculating the age.

### Token controllers

These are functions used for logging in. They work with the JWT to ensure security and logged in functionalities

### Secured controllers

Most of these are functions only a logged in user can use, such as getting and editing profile information (including the password).

### Course and assignment controllers

These are functions necessary for course and assignment creation and editing. They also require the user to be logged in.

### Achievement controllers

Functions for creating, editing, adding, obtaining and retreiving achievement information. They will be used in tandem with admin controllers to allow the administrator to manipulate achievements in the database.

### Administrator controllers

These functions perform important administrator functionalities.

## Helper

The helper package has been developed to aid in testing. It adds some hardcoded achievements and accounts (admin and test) which are new functionalities that need to be tested but can be tedious to set up.

## Packages

go get -u gorm.io/gorm

    go get -u gorm.io/driver/sqlite

    go get -u github.com/gin-gonic/gin

    go get -u golang.org/x/crypto

    go get -u github.com/golang-jwt/jwt

    go get -u github.com/gin-contrib/cors

    go get -u github.com/lib/pq

    go get -u github.com/codegangsta/gin

## Delve setup

### Installing Delve

Run

    go get -u github.com/go-delve/delve/cmd/dlv 

from the console in the directory that contains go.mod file. Once the command has executed, check if it was installed properly by running 

    dlv

in the console. If the message is dlv command not found, you need to set up the Go-related ENV. Ensure that GOROOT, GOPATH and PATH environment variables are set up properly in the Environment variables. 

### Configuration

The .vscode/launch.json file updated with this document already has the configuration necessary for running the VS Code debugger. To ensure that your version of the launch.json file is correct, check if it has the following configuration:

    {
      "name": "Go Debugger",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}\\back_end\\src",
      "env": {"GITHUB_TOKEN": "xxxxx"}
    }

Once the file is updated, running the debugger from VS Code works as regular. The only thing that needs to be done after this is comment line 10 out in back_end/src/main.go, and uncomment line 13.
