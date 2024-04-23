#!/bin/bash
set -o allexport; source .release.env; set +o allexport

echo "📦 create release: $TAG $MESSAGE"

# ------------------------------------
# Install GitHub CLI
# ------------------------------------
# sudo apt-get -y install gh
# set -o allexport; source .github.env; set +o allexport
# gh auth login

find . -name '.DS_Store' -type f -delete

git add .
git commit -m "📦 ${MESSAGE}"

git tag -a ${TAG} -m "${MESSAGE}"
git push origin ${TAG}

cp darwin/amd64/seven release/seven-darwin-amd64
cp darwin/arm64/seven release/seven-darwin-arm64
cp linux/amd64/seven release/seven-linux-amd64
cp linux/arm64/seven release/seven-linux-arm64

gh release create ${TAG} release/seven-darwin-amd64 release/seven-darwin-arm64 release/seven-linux-amd64 release/seven-linux-arm64 --title "${MESSAGE}" 
#--notes "${MESSAGE}"




