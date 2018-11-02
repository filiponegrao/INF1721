env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64_linux_linear_selection
env GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64_windows_linear_selection.exe
env GOOS=windows GOARCH=386 go build -o bin/windows/386_windows_linear_select.exe
env GOOS=darwin GOARCH=386 go build -o bin/macos/386_macos_linear_select
env GOOS=darwin GOARCH=amd64 go build -o bin/macos/amd64_macos_linear_select
env GOOS=darwin GOARCH=arm go build -o bin/macos/arm_macos_linear_select
env GOOS=darwin GOARCH=arm64 go build -o bin/macos/arm64_macos_linear_select

go build -o linear_selection

