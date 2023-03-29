# Sprint 3 Documentation

## Back-end Documentation

Current version of back-end source code is compiled into main.exe. Thus, without any updates made, running just main.exe will start the server. This makes running the app much faster for testing and examining purposes. 

### User model changed
The struct User is changed to contain a bool variable called IsAdmin, which is going to mark administrator accounts. Administrator accounts will be able to create achievements and add them into the database and potentially monitor/moderate communication (if we implement a feature like that). Administrator accounts will only be able to be added directly into the back-end by the developers.

### Login request
The function for login was implemented and developed a while ago, but I have updated it now to make it slightly easier for the front-end to read and store the JWT as a cookie. Nothing much has changed with respect to functionality other than that. 

### Administrator account
The administrator account has several privileges that the regular user account does not. The first privilege is the ability to create, modify, and delete achievements straight from the website. This makes populating and editing the achievements database more convenient and faster. An additional privilege the admin account has is deleting users from the database. This is not something that the administrator should be able to do in all cases regularily, but during development, this feature is very useful. The function for deleting a user by the administrator is already implemented, along with the routing.

### Front-End Documentation
##Unit Testing
As stated in previous sprints, using cypress to test is what I feel is ideal for unit testing our pages. I am not a fan of angular jasmine testing, it gives many errors for things that should work with little online documentation
##Tests
Testing correct button visibility for various pages
Testing that buttons link to correct pages
