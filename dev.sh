#!/bin/bash

go get -u github.com/oxequa/realize

realize start &
P1=$!
npm run dev --prefix frontend &
P2=$!
wait $P1 $P2
