# Sprint 3 Documentation

## Back-end Documentation

Current version of back-end source code is compiled into main.exe. Thus, without any updates made, running just main.exe will start the server. This makes running the app much faster for testing and examining purposes. 

### User model changed
The struct User is changed to contain a bool variable called IsAdmin, which is going to mark administrator accounts. Administrator accounts will be able to create achievements and add them into the database and potentially monitor/moderate communication (if we implement a feature like that). Administrator accounts will only be able to be added directly into the back-end by the developers.

### Login request
The function for login was implemented and developed a while ago, but I have updated it now to make it slightly easier for the front-end to read and store the JWT as a cookie. Nothing much has changed with respect to functionality other than that. 

### Administrator account
The administrator account has several privileges that the regular user account does not. The first privilege is the ability to create, modify, and delete achievements straight from the website. This makes populating and editing the achievements database more convenient and faster. An additional privilege the admin account has is deleting users from the database. This is not something that the administrator should be able to do in all cases regularly, but during development, this feature is very useful. The function for deleting a user by the administrator is already implemented, along with the routing. In addition, a function to edit certain aspects of users accounts is also included. Finally, a function which returns all users currently present in the database has been created. In the near future, a vision of an admin page would be to view all available users on a table that can be manipulated, much like the database we as developers can already see presently, but can not directly interact with. Again, these functions would likely change during normal operation of the website, as administrators rarely have this much unrestricted access to individual user accounts. 

### Unit testing

As usual, REST API extension in VS Code will be used for unit testing of the routing and functions. The new tests added in this stage will be for testing the login function which will take in a username and password, and will return a JWT. There are two types of tests for this functionality. First is the successful login, and the second one will be a wrong username/password attempt which should return Internal Server Error for wrong username (user not in the database) or Unauthorized Error for wrong password. Additionally, we will have tests for profile information access. As with login, this will have two tests. First one will be a successful attempt at getting profile information (user is logged in) and a failed one (JWT expired and user is now logged out). Additionally, we are going to test trying to access profile information with a token that no longer exists in the database.

### DOB Formatting 

When the request for the DOB of the registered user is received from the front-end, the day of the week and time are also received. These are values that are completely unnecessary for what we want to display, which is simply the birth date of the user that they put in, in MM/DD/YYYY format. Since every request received from the front-end is the same, we can do this through a simple string operation in Go. The first four characters are the day chosen, which is unnecessary, and every character after the last digit of the year, at position 15, is also unnecessary. So, we are able to just take the string from character positions 4-15, and it represents all of the information we need to obtain from the user, and truncates what we don't need to display. This value can now be easily calculated into the user's age, which will be implemented to be stored in their user model in the next sprint. 

## Front-End Documentation

### New Pages / Changes
Below will list the new pages that have been created this sprint and what each page's functionality is.
For the linking on this sprint, we were able to successfully link the following pages to the backend: Login, User Profile, Register.
For next sprint we will finish the linking on these pages: Add assignment, Add course, Achievement.

#### Header
The header has been edited to add two new buttons. 
 1. My Profile: This is the profile page button which reroutes the webpage to the user profile page, read the functionality below.
 2. New Assignment: This is the button to add a new assignment to the user when logged in, read the functionality below.
 
#### User profile page
The user profile page will display all of the user's information they inputted during registering for our website.
There is one new tag that displays the age of the user which is derived from their date of birth.
The page also contains 3 buttons. 
 1. Edit Profile: Once this is clicked the site is rerouted to the edit user page, read the functionality below.
 2. Add a course: Same as the edit profile button this reroutes the site to the add course page, read the functionality below.
 3. Add an assignment: This reroutes the webpage to add assignment page, read the functionality below.

#### Edit user profile
The edit user page will contain a form that has inputs for all of the user's data which will allow the user to change anything they would like. 
The page also contains 2 buttons.
 1. Cancel: This reroutes the webpage to the same page (edit user) which resets all fields.
 2. Save changes: This sends all of the data that is contained in the fields to the backend to be saved into the user's current data to overwrite all new changes that occurred.

#### Add course
The front-end part of the assignment page has been created and is ready for linking to the backend. 
We are only waiting for the functionality to be implemented in the backend.
The course page has 3 required fields to be filled in,
Course Code, Course Name, and Professor Name.
The page also contains 2 buttons.
 1. Clear: This reroutes the webpage to the same page (add assignment) which resets all fields.
 2. Add the course: This button is set to send a post function to the backend to send all of the course data with it.
For now when the Add the course button is clicked the user's inputted info is displayed in the console.

#### Add assignment
The front-end part of the assignment page has been created and is ready for linking to the backend. 
We are only waiting for the functionality to be implemented in the backend.
The assignment page has 4 required fields to be filled in,
Assignment name, Course Name, Assignment Type, and Due Date.
The assignment type is a drop-down field that contains options of: Homework, Quiz, Project, or Exam.
In the future, I would like to add this same feature for the course name and drop down all of the user's current courses.
The page also contains 2 buttons.
 1. Clear: This reroutes the webpage to the same page (add assignment) which resets all fields.
 2. Create Challenge: This button is set to send a post function to the backend to send all of the assignment data with it.
For now, when the Create Challenge button is clicked the user's inputted info is displayed in the console.

#### Achievement
The achievement page has not been implemented yet. The backend has been set up but is still waiting for the front-end development to catch up to their progress.
The component has been created and the linking is ready to be implemented into other pages.



### Unit Testing
As stated in previous sprints, using cypress to test is what I feel is ideal for unit testing our pages. I am not a fan of angular jasmine testing, it gives many errors for things that should work with little online documentation. Jasmine testing with angular causes many errors that by running our program we can see simply aren't true. This includes things like service creation failing even though we can most certainly use that service. For this reason, cypress is going to be our go to.
### Tests
1. Testing correct button visibility for landing page
2. Testing correct hidden buttons for landing page
3. Testing correct button visibility for login
4. Testing correct hidden buttons for login
5. Testing correct button visibility for profile
6. Testing correct hidden buttons for profile
7. Testing Login Link
8. Testing Register Link
9. Testing that login button will not work with invalid credentials (sends user to same page)
10. Testing home from anywhere links

### Login Implementation
Login now works once backend server is running. When a user logs in and they have valid credentials, a cookie is created on the front end that allows us to track that the user is indeed logged in. The backend has requests that can only be made by users that are logged in, and we let them know this by injecting the cookie into every HTTP request using the HTTPINJECTOR service. The JWTTOKENSERVICE is used to take in JWT tokens from backend, and the APPCOOKIE service is used ot store the cookies. Additionally, the profile apge has been configured to show the users information they gave us. This is retreived using a request from backend and it only works when a user is logged in.
