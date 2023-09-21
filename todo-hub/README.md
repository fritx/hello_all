- [ ] install from github
- [ ] cmd name => todo
- [ ] status todo(o), wip(w), done(x)
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
go get github.com/steebchen/prisma-client-go
go get prisma-go/db

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
# >> ? What is the title? 
# ...

todo list
# >> Prisma schema loaded from schema.prisma
# Prisma Studio is up on http://localhost:5555
```
