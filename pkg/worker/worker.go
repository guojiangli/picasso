package worker

// Worker could scheduled by picasso or customized scheduler
type Worker interface {
	Run() error
	Stop() error
}
