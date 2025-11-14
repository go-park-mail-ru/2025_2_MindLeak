package subscribers

type Usecase interface {
	Subscribe()
	Unsubscribe()
	GetSubscribers()
	GetSubscriptions()
}
