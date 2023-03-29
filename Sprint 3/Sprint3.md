# Sprint 3 Documentation

## Back-end Documentation

Current version of back-end source code is compiled into main.exe. Thus, without any updates made, running just main.exe will start the server. This makes running the app much faster for testing and examining purposes. 

### User model changed
The struct User is changed to contain a bool variable called IsAdmin, which is going to mark administrator accounts. Administrator accounts will be able to create achievements and add them into the database and potentially monitor/moderate communication (if we implement a feature like that). Administrator accounts will only be able to be added directly into the back-end by the developers.

### Login request
The function for login was implemented and developed a while ago, but I have updated it now to make it slightly easier for the front-end to read and store the JWT as a cookie. Nothing much has changed with respect to functionality other than that. 

### Administrator account
The administrator account has several privileges that the regular user account does not. The first privilege is the ability to create, modify, and delete achievements straight from the website. This makes populating and editing the achievements database more convenient and faster. An additional privilege the admin account has is deleting users from the database. This is not something that the administrator should be able to do in all cases regularily, but during development, this feature is very useful. The function for deleting a user by the administrator is already implemented, along with the routing.

## Front-End Documentation
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
