## Pre-flight Checklist

1. Install Go
2. Prepare environment
3. Get my code

### Install Go

You'll need to install the Golang development environment in order to
run the tests.

    $ brew install go

### Prepare environment

Go requires an environment variable named `$GOPATH` which points to your
Go workspace. You can also export the Go `bin` directory in order to run 
your executables. Export `$GOPATH` from your `.zshrc` or analogous file like so:

    export GOPATH=/Users/sandropadin/Project/go 
    export PATH=$GOPATH/bin:$PATH # to access Go binaries

You will need to create that `go` directory with the following folder
structure.

    ~/Projects/go
    $ tree -L 1
    .
    ├── bin
    ├── pkg
    └── src

### Install dependencies

Golang can get dependencies from any repository including: Git, Mercurial and 
Bazaar repos. One of my Tic Tac Toe dependencies is in a Bazaar repo so you will 
need to install Bazaar DCVS:

    $ brew install bazaar

Now you can install the rest of the dependencies:

    $ go get launchpad.net/gocheck
    $ go get github.com/remogatto/prettytest 

### Get my code

Now that you have Go installed and the environment set up, you can grab
my code with the following command:

    $ go get github.com/spadin/tictactoe

Running this command downloads all of my files from Github to a local directory
at: `$GOPATH/src/github.com/spadin/tictactoe`

## Running specs

You'll have to run the specs for each package individually. You can do
so by switching into each sub-directory and running the following:

    $ go test -i # installs dependencies
    $ go test

## Running the game

    $ cd $GOPATH/src/github.com/spadin/tictactoe
    $ go run tictactoe.go

## Installing the game

    $ cd $GOPATH/src/github.com/spadin/tictactoe
    $ go install
    $ tictactoe # assuming $GOPATH/bin is in your $PATH

## Need more help?

The Go documentation is well written, organized and helpful.

In particular, I've found these links to be of great help:

1. [Go Programming Language Specification][2]
1. [How to Write Go Code][3]
1. [Effective Go][1]
1. [Go Package Documentation][4]

[1]: http://golang.org/doc/effective_go.html
[2]: http://golang.org/ref/spec
[3]: http://golang.org/doc/code.html
[4]: http://golang.org/pkg/
