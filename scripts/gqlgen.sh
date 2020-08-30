#!/bin/bash

printf "\nRegenerating gql=gen files\n"
rm -rf internal/gql/generated.go \
	internal/gql/models/generated.go \
	interal/gql/resolvers/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"
