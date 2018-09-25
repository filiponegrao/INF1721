env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64_linux_linear_selection
env GOOS=linux GOARCH=amd64 go build -o bin/windows/amd64_windows_linear_selection.exe
env GOOS=linux GOARCH=386 go build -o bin/windows/386_windows_linear_select.exe
go build -o linear_selection

