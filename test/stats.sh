#!/bin/bash

statssys() {
    curl \
    http://127.0.0.1:8080/api/xray/statssys
}

query() {
    curl \
    http://127.0.0.1:8080/api/xray/statsquery/
}

stats() {
    curl \
    http://127.0.0.1:8080/api/xray/stats/
}

case "$1" in
    sys)
        statssys
        ;;
    query)
        statsquery
        ;;
    stats)
        stats
        ;;
    *)
        "$0:  usage:  $0 {sys|query|stats}"
        ;;
esac