#!/bin/bash

# algorithms=("md5" "sha256" "sha256old" "sha256new" "imohash")
algorithms=("sha256" "sha256old" "sha256new")
file=$1
loop_time=$2

function oriFunc() {
    sync;echo 3 > /proc/sys/vm/drop_caches
    used_before=$(free | grep Mem | awk '{print $3}')
    free_before=$(free | grep Mem | awk '{print $4}')
    time sha256sum "${file}"
    used_after=$(free | grep Mem | awk '{print $3}')
    free_after=$(free | grep Mem | awk '{print $4}')
    used=$((used_after-used_before))
    free=$((free_after-free_before))
    printf "algo: %-10s used: %-10d free: %-10d\n" sha256sum $used $free
}

function hashFunc() {
    for al in "${algorithms[@]}"; do
        sync;echo 3 > /proc/sys/vm/drop_caches
        used_before=$(free | grep Mem | awk '{print $3}')
        free_before=$(free | grep Mem | awk '{print $4}')
        ./hash -path "${file}" -algo "${al}" -op "${loop_time}" -gc
        used_after=$(free | grep Mem | awk '{print $3}')
        free_after=$(free | grep Mem | awk '{print $4}')
        used=$((used_after-used_before))
        free=$((free_after-free_before))
        printf "algo: %-10s used: %-10d free: %-10d\n" "${al}" $used $free
        echo "============================================"
    done
}

hashFunc
oriFunc
