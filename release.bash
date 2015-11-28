rm -rf bin &&
mkdir bin &&
GOARCH=386 GOOS=linux go build -o bin/icollege.linux &&
cp -r server bin &&
cp -r public bin &&
scp -r bin odchazel@odchazel.com:/home/odchazel/icollege &&
rm -rf bin