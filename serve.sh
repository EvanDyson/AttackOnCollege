#!/bin/sh

ng serve &
gin --port 1337 --path . --build ./src/back_end/src/ --i --all &

wait