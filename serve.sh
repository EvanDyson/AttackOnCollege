#!/bin/sh

ng serve &
gin --port 4040 --path . --build ./back_end/src/ --i --all &

wait