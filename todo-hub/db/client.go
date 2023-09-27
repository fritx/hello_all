package db

func GetClient() (client *PrismaClient, disconnect func(), err error) {
	client = NewClient()
	if err = client.Prisma.Connect(); err != nil {
		return
	}
	disconnect = func() {
		err := client.Prisma.Disconnect()
		if err != nil {
			panic(err)
		}
	}
	return
}
