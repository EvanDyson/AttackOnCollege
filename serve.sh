#!/bin/sh

<<<<<<< HEAD
ng serve &
gin --port 1337 --path . --build ./back_end/src/ --i --all &
=======
ng serve & go run ./back_end/src/*.go &
>>>>>>> ce53c2e2a020a8a5009001757e0c168bbe19f4e0

wait