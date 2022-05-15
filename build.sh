GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o release/linux/ddl-to-object
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o release/win/ddl-to-object.exe
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o release/mac/ddl-to-object
cp -r template/ release/linux/template
cp -r template/ release/win/template
cp -r template/ release/mac/template