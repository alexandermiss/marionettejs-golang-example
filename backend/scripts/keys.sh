#!/bin/sh

configureKeys(){
    mkdir -p ~/.ssh
    chmod 700 ~/.ssh
    touch ~/.ssh/known_hosts
    ssh-keyscan "$DO_IP" >> ~/.ssh/known_hosts
    chmod 644 ~/.ssh/known_hosts

    eval $(ssh-agent -s)

    echo "$SSH_PRIVATE_KEY" | ssh-add -
    echo "$SSH_KNOWN_HOSTS" >> ~/.ssh/known_hosts
}

configureKeys
