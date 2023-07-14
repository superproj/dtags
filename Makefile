build-cli-plugin:
	go build -o ./dist/ -ldflags "-s -w -X github.com/superproj/dtags/cmd.cliPlugin=true"
