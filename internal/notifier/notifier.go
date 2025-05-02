package notifier

type Notifier interface {
	Notify(title string, fields map[string]string) error
}
