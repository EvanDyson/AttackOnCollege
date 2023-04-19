# Welcome to Attack on College, our CEN3031 project

## Developed by Pierce Casey, Evan Dyson, Michael Hoctor, Sara Osmanovic

*Working on the front end will be Evan Dyson and Michael Hoctor.*

*Working on the back end will be Pierce Casey and Sara Osmanovic.*

## Documentation of the requirements for running and using our app.

### Front-End
Our front-end is based off of angular. 
In order to properly run angular applications it is recommended by the developers to have Node.js installed, as well as having the angular CLI downloaded locally.

* [Node.js](https://nodejs.org/en)
* After downloading Node.js, run the command ``` npm install -g @angular/cli ``` in yor console.

After downloading both required things and downloading our project AttackonCollege, enter the project folder and run the following commands in console.

* ``` npm install ```
* ``` ng install ```
* ``` ng serve ```

You can then navigate to [localhost:4200](http://localhost:4200)

**Please note that this will only launch the front end unconnected to the backend**


### Back-end

As of now, the front-end and back-end are linked up. Starting the server and Angular allows the two to communicate. GoLang needs to be installed and set up for execution on the machine to properly start the server. If both Angular and Go are installed, running ```npm start``` in the console of the AttackOnCollege directory will start both up. For testing purposes, running ```go run ./back_end/src/*.go``` will start up just the server. Although Codegangsta Gin is NOT used to start the server, it does give testing feedback in the console. The reason for not running with Gin is that Gin is very verbose and it is hard to read through all the ouput.

Inspiration from: https://lvluplife.com/
