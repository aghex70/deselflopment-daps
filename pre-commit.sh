#!/bin/bash

bold=$(tput bold)
green=$(tput setaf 2)
yellow=$(tput setaf 3)
orange=$(tput setaf 202)
reset=$(tput sgr0)

echo "${bold}${yellow}***** Executing Go pre-commit (7 steps) *****${reset}"
docker-compose run backend sh -c \
  "echo '${bold}${green}1- Installing dependencies...${reset}' &&
  go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest  > /dev/null 2>&1 && \
  go install -v github.com/fzipp/gocyclo/cmd/gocyclo@latest  > /dev/null 2>&1 && \
  go install -v github.com/go-critic/go-critic/cmd/gocritic@latest  > /dev/null 2>&1 && \
  go install -v golang.org/x/tools/cmd/goimports@latest  > /dev/null 2>&1 && \
  go install -v github.com/securego/gosec/v2/cmd/gosec@latest  > /dev/null 2>&1 && \
  echo '${bold}${green}2- Running gofmt...${reset}' && \
  gofmt -w . && \
  echo '${bold}${green}3- Running goimports...${reset}' && \
  goimports -w . && \
  echo '${bold}${green}4- Running gocyclo...${reset}' && \
  gocyclo -over 20 . && \
  echo '${bold}${green}5- Running gocritic...${reset}' && \
  gocritic check ./... && \
  echo '${bold}${green}6- Running gosec...${reset}' && \
  gosec ./... && \
  echo '${bold}${green}7- Running golangci-lint...${reset}' && \
  golangci-lint run" || { echo "${bold}${orange}One or more commands failed.${reset}"; exit 1; }

exit 0;
