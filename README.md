# chorus
![Heroku](https://heroku-badge.herokuapp.com/?app=chorus-rr&style=flat&svg=1) [![Maintainability](https://api.codeclimate.com/v1/badges/4562119d0db2daec068b/maintainability)](https://codeclimate.com/github/bottleneckco/chorus/maintainability) [![Known Vulnerabilities](https://snyk.io/test/github/bottleneckco/chorus/badge.svg)](https://snyk.io/test/github/bottleneckco/chorus)

chorus is the place for sharing music with your friends in real time. 

## Get started
To start sharing, head over to http://chorus.2017.reactriot.com. Create your channel and share the URL with others to invite them to your channel. Add some music to get the party started 🎉

## Development
This repository is using the new `dep` package manager for Golang.
1. Clone this repo to `$GOPATH/src/github.com/bottleneckco/chorus`. _*NO OTHER DIRECTORY WORKS*_
2. Run `dep ensure` (This is like `yarn install`)
3. Run `foreman -f Procfile.dev` (or alternatively use https://github.com/ddollar/forego, a Go equivalent).
4. Visit http://localhost:8080
