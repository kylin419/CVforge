package progress

type Reporter interface {
	Start(total int64, desc string)

	Update()

	Finish()
}
