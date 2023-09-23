- [ ] db views
- [ ] install from github
- [ ] cmd name => todo
- [x] db migration
- [x] status todo(o), wip(w), done(x)
- [x] todo add
- [x] todo list

## Usage

```sh
todo
# >> ? * Title: ______ [Enter]
# ? Description: ________ [Enter]
# Todo added: {
#   "id": "clmvixdt00000bmlqkrffaiyh",
#   "title": "Hello",
#   "desc": "World...",
#   "status": "",
#   "updatedAt": "2023-09-23T04:21:39.488Z",
#   "createdAt": "2023-09-23T04:21:39.488Z"
# }

npm i -g prisma  # needed by cmd `todo list`
todo list
# >> Prisma schema loaded from schema.prisma
# Prisma Studio is up on http://localhost:5555

todo wip  # or `todo w / doing`
# >> Loop: Marking a todo as 'wip'...
# Press CTRL-C to exit.
# ? * Id: ________ [Enter]
# Todo marked as 'done': {
#   "id": "clmw20nda0000bm7v7jh847ty",
#   "title": "Hi",
#   "desc": "",
#   "status": "wip",
#   "firstWipAt": "2023-09-23T13:16:16.815Z",
#   "lastDoneAt": null,
#   "updatedAt": "2023-09-23T13:16:16.816Z",
#   "createdAt": "2023-09-23T13:16:04.551Z"
# }

todo done  # or `todo x / check / finish`
todo undo  # or `todo o / uncheck / todo`
```

## Install from GitHub

Coming soon...

## Development & Run from Source

CLI:<br>
<https://cobra.dev/><br>
<https://github.com/go-survey/survey>

Prisma:<br>
<https://goprisma.org/docs/getting-started/quickstart><br>
<https://goprisma.org/docs/getting-started/advanced><br>
<https://prisma.io/>

```sh
go get prisma-go/db
go get github.com/steebchen/prisma-client-go

# sync the database with your schema
go run github.com/steebchen/prisma-client-go db push
# The Prisma Client Go client is automatically generated in your project.
# You can re-run this command any time to sync your schema with the database.

go mod tidy
go mod download
go run .
alias todo="go run ."
```

## Build from Source

```sh
go run github.com/steebchen/prisma-client-go generate
go run github.com/steebchen/prisma-client-go migrate dev --name xxx

go build
go install .
alias todo=~/go/bin/todo-hub
# or
go build -o ~/go/bin/todo
```

## Release & Deploy

Coming soon...

Developing with Prisma Migrate<br>
https://www.prisma.io/docs/guides/migrate/developing-with-prisma-migrate

Team development with Prisma Migrate<br>
https://www.prisma.io/docs/guides/migrate/developing-with-prisma-migrate/team-development

Prisma CLI with Go<br>
https://goprisma.org/docs/getting-started/cli

```sh
# just re-generate the Go client
go run github.com/steebchen/prisma-client-go generate

# sync the database with your schema for development
go run github.com/steebchen/prisma-client-go db push

# create a prisma schema from your existing database
go run github.com/steebchen/prisma-client-go db pull

# for production use, create a migration locally
go run github.com/steebchen/prisma-client-go migrate dev

# sync your production database with your migrations
go run github.com/steebchen/prisma-client-go migrate deploy
```
