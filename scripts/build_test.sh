#/bin/bash

set -e

# Running any go tests
echo "Running go test.."
go test ./...

# Create empty sqlite db file. Will be overwritten everytime.
echo "Create empty sqlite db file"
touch db/sqlite/localfarm.db
chmod 775 db/sqlite/localfarm.db

echo "Building golang binaries..."
# Build and test Golang
make linux-arm linux-amd64 windows

echo "Configuring binary for running ..."
# Setting up configuration
cp conf.json.example conf.json
sed -i.bak "s|/Users/user/Code/golang/src/github.com/LocalFarm/localfarm-server|$TRAVIS_BUILD_DIR|g" conf.json
# Set DEMO_MODE to true to turn off the token validation
DEMO_MODE=false
echo "DEMO_MODE is set to ${DEMO_MODE}"

echo "Starting server for E2E testing ..."

# Run golang on linux
./localfarm.linux.amd64 > /dev/null 2>&1 &
LOCALFARM_PID=$!

echo "Server has running running in the background at pid ${LOCALFARM_PID}"

echo "Running Front-End Unit tests ..."
# build and run unit test
# npm install && npm run unit
npm install

# echo "Running end to end tests ..."
# build and test e2e
# npm run prod && npm run cypress:run
npm run prod
echo "Killing Server [$LOCALFARM_PID] ..."

# Move the screenshoot and recorded video from the test result into public folder
# mkdir public/assets
# cp -rf resources/tests/assets/videos public/assets/

kill -s TERM $LOCALFARM_PID
