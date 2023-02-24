#!/bin/sh

ng serve &
gin --port 1337 --path . --build ./back_end/src/ --i --all &

wait