- [ ] db views
- [ ] install from github
- [ ] cmd name => todo
- [x] status todo(o), wip(w), done(x)
- [x] todo add
- [x] todo list

## Install from GitHub

```
todo...
```

## Development & Build from Source

CLI:
<https://cobra.dev/>
<https://github.com/go-survey/survey>

Prisma:
<https://goprisma.org/docs/getting-started/quickstart>
<https://goprisma.org/docs/getting-started/advanced>
<https://prisma.io/>

```sh
go get prisma-go/db
go get github.com/steebchen/prisma-client-go

# sync the database with your schema
go run github.com/steebchen/prisma-client-go db push
# The Prisma Client Go client is automatically generated in your project.
# You can re-run this command any time to sync your schema with the database.

# once if schema change & db migrate
go run github.com/steebchen/prisma-client-go migrate dev --name xxx

go run .

go build
go install .
alias todo=~/go/bin/todo-hub

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
```
