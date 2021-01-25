#!/bin/bash

hwi() {
    curl \
    --header "Content-Type: application/json" \
    --data '
    {
        "tag": "http-test",
        "port": 2080,
        "protocol": "http",
        "settings": {}
    }
    ' \
    http://127.0.0.1:8080/helper/inbound/add
}

hri() {
    curl \
    http://127.0.0.1:8080/helper/inbound/read/http-test
}

hai() {
    curl \
    http://127.0.0.1:8080/helper/inbound/apply/http-test
}

xao() {
    curl \
    --header "Content-Type: application/json" \
    --data '
    {
        "tag": "proxy-test",
        "protocol": "VLESS",
        "settings": {
            "vnext": [
                {
                    "address": "192.168.1.1",
                    "port": 443,
                    "users": [
                        {
                            "id": "00000000-0000-0000-0000-000000000000",
                            "alterId": 0,
                            "email": "t@t.tt",
                            "security": "none",
                            "encryption": "none",
                            "flow": "xtls-rprx-splice"
                        }
                    ]
                }
            ],
            "servers": null,
            "response": null
        },
        "streamSettings": {
            "network": "tcp",
            "security": "xtls",
            "xtlsSettings": {
                "allowInsecure": false,
                "serverName": "foo.bar"
            },
            "tcpSettings": null,
            "kcpSettings": null,
            "wsSettings": {
                "connectionReuse": true,
                "path": "/xray",
                "host": null,
                "headers": null
            },
            "httpSettings": null,
            "quicSettings": null
        },
        "mux": {
            "enabled": false
        }
    }
    ' \
    http://127.0.0.1:8080/xray/outbound/add
}

xai() {
    curl \
    --header "Content-Type: application/json" \
    --data '
    {
        "tag": "http-test",
        "port": 2080,
        "protocol": "http",
        "settings": {}
    }
    ' \
    http://127.0.0.1:8080/xray/inbound/add
}

xri() {
    curl \
    http://127.0.0.1:8080/xray/inbound/remove/http-test
}

xro() {
    curl \
    http://127.0.0.1:8080/xray/outbound/remove/proxy-test
}

case "$1" in
    hwi)
        hwi
        ;;
    hri)
        hri
        ;;
    hai)
        hai
        ;;
    xri)
        xri
        ;;
    xro)
        xro
        ;;
    xai)
        xai
        ;;
    xao)
        xao
        ;;
    *)
        echo "$0:  usage:  $0 {hwi|hri|hai|hwo|hro|hao|xri|xro|xai|xao}"
        ;;
esac