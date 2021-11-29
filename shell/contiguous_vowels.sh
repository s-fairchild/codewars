#!/bin/bash
# Steven Fairchild 2021-07-27
# https://www.codewars.com/kata/5d2d0d34bceae80027bffddb

# "a e ii ooo u" == "ooo ii a e u"
count_vowels() {
    s=( $(echo "$1" | tr ' ' '\n') )
    declare -a vc
    for i in "${!s[@]}"; do
        vc["$i"]=$(echo "${s[$i]}" | tr -c '[aeiouAEIOU]' ',')
        while [[ "${vc[$i]: -1}" == ',' ]] || [[ "${vc[$i]:0:1}" == ',' ]]; do
            vc["$i"]=$(echo "${vc[$i]}" | sed s/^,//g)
            vc["$i"]=$(echo "${vc[$i]}" | sed s/,$//g)
        done

        IFS=',' read -r -a tmparr <<< "${vc[$i]}"
        max=${#tmparr[0]}
        for n in "${tmparr[@]}" ; do
            if [ "${#n}" -gt "${#max}" ] ; then
                max="${#n}"
            fi
        done
        vc["$i"]="$max"
    done
    for (( i=0; i<"${#vc[@]}"; i++ )); do
        for (( j="${#vc[@]}"; j>"$i"; j-- )); do
            k=$(( j - 1 ))
            if [[ "${vc[$i]}" < "${vc[$j]}" ]]; then
                echo "before swap: ${s[$i]} ${s[$j]}"
                echo "before swap: ${vc[$i]} ${vc[$j]}"
                tmp="${s[$k]}"; vctmp="${vc[$k]}"
                s["$k"]="${s[$j]}"; vc["$k"]="${vc[$j]}"
                s["$j"]="$tmp"; vc["$j"]="$vctmp"
                #echo "${#vc[@]}"
                #echo "${s[@]}"
                echo "after swap: ${s[$i]} ${s[$j]}"
                echo "after swap: ${vc[$i]} ${vc[$j]}"
            fi
        done
    done
    echo "${s[@]}"
}

count_vowels "$1"