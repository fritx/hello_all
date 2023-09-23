package db

func GetClient() (client *PrismaClient, disconnect func()) {
	client = NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	disconnect = func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}
	return
}
