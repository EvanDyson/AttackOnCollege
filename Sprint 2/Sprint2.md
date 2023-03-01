# Sprint 2 Documentation

# Front-End
## Testing
### Unit test

### Cypress test

## Updates

Since Sprint 1 we finished up the page linking so that we can have properly functioning link buttons to bring the user to all of our created pages (components). 

# Back-end

## Functionality

The implemented functionality:

### Linking front-end and back-end

As of now, the front-end and back-end are linked up. Starting the server and Angular allows the two to communicate. GoLang needs to be installed and set up for execution on the machine to properly start the server. If both Angular and Go are installed, running 
    
    npm start 
    
in the console of the AttackOnCollege directory will start both up. For testing purposes, running 

    go run ./back_end/src/*.go 

will start up just the server. Although Codegansta Gin is NOT used to start the server, it does give testing feedback in the console. The reason for not running with Gin is that Gin is very verbose and it is hard to read through all the ouput.

### Register user

Along with the respective front-end functionality, the router and handler functions for registering a user into the database are fully implemented and functional. In Sprint 1, registering the user required JSON input from the POST requests, but that has been updated to take in any object with the matching fields (firstName, lastName, email, username, password, dob, major, college). As long as the object passed can be bound to the request object in the backend, and it passes GORM requirements, the new user will be created. 

When the user is created, the password string is hashed and the hash is placed into the database. This hash will directly be used to compare the password during login. 

Once the user is added to the database with no errors, a JSON object containing the email and username of the newly created user is passed back to the front-end as a confirmation method. This needs to be updated at this stage to make it more proctical for the front-end to conduct tests.

### Notes on Date of Birth

For now, the entire string passed from the front-end is used as the DOB. A new function needs to be developed to extract only the day, month, and year from the input. 

### Login user

Although the back-end function for logging a user in (generateToken()) is implemented, some changes will probably be introduced once the front-end implements the login functionality. For now, a JWT is generated and stored in the database for retrieval, and the token is sent back to the front-end. 

### Create course and assignment

These functions work in a similar fashion as the register user function. The only difference is that the request will have to contain an "Authorization" header with the current token. If this token is expired, an error message will be sent back. If the passed token cannot be found in the database, an error message will be sent back. If the token is not passed, an error message will be sent back.

The assignments can be completed, too, and the achievement for completing the first assignment is created, so it can be given once the user completes their first assignment.

## Structs and packages

### User Struct and Database

The models package contains a struct called User which represents the profile that the website user will have. The struct has the following fields:
1. Username: website handle that will be displayed on the profile, and will be used for looking up other users and displaying the global status on the ranked list of users. Gorm will require all emails in the database to be unique
2. Email: a method of identification that will be used for logging in. Gorm will require all emails in the database to be unique
3. Password: another part of identification
4. First name and last name: Personalization of the user
5. Date of Birth: user's date of birth. Could potentially be used to sent Birthday notifications
6. College and Major: these two will be used to potentially organize the leaderboards based on the user's environment. This way they will be able to compete against peers in their major and/or college.
7. Current course: the code of the course that the user is currently taking. This is used to look up the course the student is taking in the database and access the assignments they have to take
8. Achievements: this is a list of IDs of the achievements the user has made. The list will be incremented as the user completes assignments and courses, and will provide additional experience points that will go toward levels.
9. Level: this is the current level that the user has in our app. It is incremented based on a scale we are yet to determine, and it will be primarily used for additonal achievements and bonus experience points. It is supposed to present small step-by-step goals for users that would keep their interest and give them instant gratification that is neccessary when playing any game, and completing a class
10. Experience points: the total number of experience points will be used to build a leaderboard of users used for competition

### Course Struct and Database

The course struct has the following fields:
1. Title: course title that is just there for an easier access by the user
2. Course code: unique course code that the User struct has access to. That is the key that is looked up in the database of courses and connects the two databases
3. Assignments: a list of unique IDs of assignments the user created. The list is a container of handles that can be then used in the database of assignments to access the specific user's assignments and modify them.
4. Final grade: the final grade earned in the course. It is used to calculate the experience points that will be earned after completing the course. 
5. Experience points: the number of experience points that will be added to user once they complete the course

### Assignment Struct and Database

The assignment struct has the following fields:
1. Title: the title of the assignment that is to be completed
2. Description: the description of the assignment. This can be the requirements for the assignment, details the professor gave, type of assignment, etc. It can also include the deadline since it should be displayed with the title
3. Number of Points: the total number of points that can be earned for the assignment
4. Weight: the grade weight. It is the percent of the final grade that this assignment will contribute to.
5. Experience points: Total experience points earned upon completion of the assignment
6. Points earned: the number of points that the user achieved upon completion

### Achievement Struct and Database

So far, only the basic fields of the achievement struct are made such as title, description and experience points. 

### Server setup

The server is set up with the [gin-gonic/gin package](https://github.com/gin-gonic/gin) that acts as a router. The TLS encryption was providing difficulty in linking with the front-end, so that has been taken out and replaced with just HTTP communication. Nevertheless, all the getter functions do not return the password hash of the user profiles. All passwords are replaced with "Hidden" before being sent in the reply.

### Databases

Databases are set up using [GORM package](https://gorm.io/gorm) for GoLang, and there are 4 so far. They are all local, and we are still looking for ways to decentralize that, although it might be pointless for the low scale project such as this one. PostgreSQL didn't work the way we expected it, and mySQL doesn't allow non-primitive datatypes which we need for the arrays of achievements, assignments, courses, etc.

### Controllers

Controllers is a package we split up from the models package, and it includes files with just functions that describe behaviors and are mostly handlers for routing. 

### Achievement constructor and accessor functions

Functions used for creating and accessing achievements were created. These will potentially be used later as controllers employed by an Administrator user instead of only being accessible from the back end. These will NOT be accessible to regular users and visitors of the website, but for now, no authentication and restriction is implemented. 

Two more functions (modifier and destructor) need to be implemented for full control over achievements.

Created an achievement controller that will manage the delegation of achievements to users who earn them. When the conditions of the achievement are met, the controller finds the achievement in the given database with the string passed into the function, along with a reference to the user who earned it. The GetAchievement function appends the achievement ID code to the achievement slice that each user struct contains. 

The first achievement implemented will be given when the user fully completes their first course, and it is titled “First Blood.” This can be changed to fit the flavor or style of the project that we eventually solidify.  

### Packages

go get -u gorm.io/gorm

    go get -u gorm.io/driver/sqlite

    go get -u github.com/gin-gonic/gin

    go get -u golang.org/x/crypto

    go get -u github.com/golang-jwt/jwt

    go get -u github.com/gin-contrib/cors

    go get -u github.com/lib/pq

    go install github.com/codegangsta/gin@latest
