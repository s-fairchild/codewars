#!/bin/bash
# Steven Fairchild 2021-08-03
# "2000 10003 1234000 44444444 9999 11 11 22 123" == "11 11 2000 10003 22 123 1234000 44444444 9999"
# "10003 1234000 44444444 9999 2000 123456789" == "2000 10003 1234000 44444444 9999 123456789"

orderweight() {
    IFS=' ' read -ra ws <<< "$1"
    declare -a dct; declare -i tmp
    for (( i=0; i<"${#ws[@]}"; i++ )); do
        for (( j=0; j<"${#ws[$i]}"; j++ )); do
            tmp+="${ws[$i]:$j:1}"
        done
        #echo "$tmp"
        rdct["$i"]="${ws[$i]}"
        wdct["$i"]="$tmp"
        tmp=0
    done
    echo "${rdct[@]}"
    echo "${wdct[@]}"
}

orderweight "$1"