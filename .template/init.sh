#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

PROJECTNAME=$(basename $(git rev-parse --show-toplevel))
echo "Initialising the repo - ${PROJECTNAME}"

echo "Setting up the README.md"
sed "s/PROJECTNAME/${PROJECTNAME}/g" .template/README.md >README.md

echo "Setting up base main.go"
sed "s/PROJECTNAME/${PROJECTNAME}/g" .template/main.go >main.go

echo "Setting up the base pkg"
mv .template/pkg pkg

echo "Setting up the gomod"
go mod init github.com/ghostsecurity/${PROJECTNAME}
go mod tidy
go mod vendor

echo "Setting up the golangci config"
mv .template/.golangci.yml .golangci.yml

echo "Setting up a base .gitignore"
mv .template/.gitignore .gitignore

echo "Setting up base Dockerfile"
sed "s/PROJECTNAME/${PROJECTNAME}/g" .template/Dockerfile >Dockerfile

echo "Setting up basic cloudbuild"
sed "s/PROJECTNAME/${PROJECTNAME}/g" .template/cloudbuild.yaml >cloudbuild.yaml

echo "Setting up .github"
mv .template/.github .github

echo "Setting up Makefile"
mv .template/Makefile Makefile

echo "Cleaning up"
rm -rf .template

echo "Resetting .git"
git add .
git commit --amend --no-edit >/dev/null
