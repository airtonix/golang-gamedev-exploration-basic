# JUSTFILE
# https://github.com/casey/just
#

default:
    @just --choose

help:
    @just --list

# Project setup
setup:
    @echo '⛳ Begin project setup'
    go get
    

# Repo cleanup
fix:
    go mod tidy

build:
    goreleaser build --snapshot --rm-dist

release *ARGS="":
    goreleaser {{ARGS}}
