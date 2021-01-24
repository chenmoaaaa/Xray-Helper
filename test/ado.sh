#!/bin/bash

ado() {
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
    http://127.0.0.1:8080/ado
}

adi() {
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
    http://127.0.0.1:8080/adi
}

rmi() {
    curl \
    http://127.0.0.1:8080/rmi/http-test
}

rmo() {
    curl \
    http://127.0.0.1:8080/rmo/proxy-test
}