#!/bin/bash

DIR=$(pwd)
sub=${DIR//[^\/]}
count=${#sub}
proj_name=$(echo $DIR | cut -d '/' -f `expr 1 + $count`-)

go build . && ./$proj_name