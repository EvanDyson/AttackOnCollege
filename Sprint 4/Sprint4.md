# Sprint 4 Documentation
## Front End
### Unit and e2e Testing
As Stated in previous sprints, testing in cypress has been what has worked best for accurate testing in our project. As all front-end functionality in this sprint requires access to the back-end, these tests work both as unit and end to end tests.

1. Test 1: Allows user to login and add a course, this is something they might actually do interactively
2. Test 2: Inverse of some of the sprint 3 tests, different buttons will be visible on landing page
3. Test 3: Inverse of some of the sprint 3 tests, different buttons will be hidden on landing page
4. Test 4: Navigates to profile page
5. Test 5: Logs user out
6. Test 6: Navigates to Achievement Page
7. Test 7: Navigates to Edit User
8. Test 8: Navigates to add assignment

Note: All of these require the user to be logged in, and as such every one of these tests also logs a profile in,
this is a back end supported function, hence why all of these are end to end.
