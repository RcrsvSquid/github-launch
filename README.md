# Github Launch

Opens the github webpage for a git repository.

Handles ssh remote urls as well. For example the remote "git@github.com:Rcrsvsquid/github-launch.git" will launch
"https://github.com/RcrsvSquid/github-launch.git" in the browser

## Installation
```bash
$ go get github.com/RcrsvSquid/github-launch.git
$ cd $GOPATH/src/github.com/RcrsvSquid/github-launch
$ go install gh-launch.go
```

## Usage
The command defaults to launching the origin url if no git remote name is passed.
Instructions assume you are inside a git repository with a valid remote
```
$ gh-launch          # opens origin url in the default browser
$ gh-launch upstream # opens upstream url in the default browser
$ gh-launch --help   # display help text
```
