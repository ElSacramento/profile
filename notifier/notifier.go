package notifier

type Operation string

const (
	Get    Operation = "get"
	Delete Operation = "delete"
	Update Operation = "update"
	Create Operation = "create"
)

// Send notifications for all subscribers.
type Notifier interface {
	Push(id uint64, opName Operation)
	Stop()
}
