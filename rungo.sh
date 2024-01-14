#!/bin/bash
rungo () {
        if [ $# -eq 0 ]
                then nodemon --exec go run main.go --signal SIGTERM
        elif [ $# -eq 1 ]
                then nodemon --exec go run $1 --signal SIGTERM
        fi
}
rungo $1