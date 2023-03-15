// УНИВЕРСАЛЬНЫЕ, АБСТРАКТНЫЕ ИНТЕРФЕЙСЫ, ДЛЯ ПОЛУЧЕНИЯ ОБНОВЛЕНИЙ

package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

const (
	Unknown Type = iota
	Message
)

type Type int

type Event struct { // ЕДИНСТВЕННОЕ СОБЫТИЕ - СООБЩЕНИЕ
	Type Type
	Text string
	Meta interface{}
}
