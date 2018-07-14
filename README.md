# create-noypi-app

Download create-noypi-app
```sh

    > go get -u github.com/noypi/clitools/...

```

Install
```sh

    > go install -v github.com/noypi/clitools/create-noypi-app

    > go build -v -o $GOBIN/create-noypi-app github.com/noypi/clitools/create-noypi-app
```

Create new project
```sh

    > cd $GOPATH/src

    > create-noypi-app  -v  -basedir=./github.com/mygit/myproject -pkg=github.com/mygit/myproject

    > cd github.com/mygit/myproject

```
