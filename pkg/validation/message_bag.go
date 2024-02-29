package validation

type MessageBag struct {
	messages map[string][]string
}

func NewMessageBag() *MessageBag {
	return &MessageBag{
		messages: make(map[string][]string),
	}
}

func NewMessageBagWithMessages(messages map[string][]string) *MessageBag {
	return &MessageBag{
		messages: messages,
	}
}

func (m *MessageBag) Add(field string, message string) {
	m.messages[field] = append(m.messages[field], message)
}

func (m *MessageBag) Has(field string) bool {
	_, ok := m.messages[field]
	return ok
}

func (m *MessageBag) Get(field string) []string {
	return m.messages[field]
}

func (m *MessageBag) All() map[string][]string {
	return m.messages
}

func (m *MessageBag) Any() bool {
	return len(m.messages) > 0
}

func (m *MessageBag) Count() int {
	return len(m.messages)
}

func (m *MessageBag) First(field string) string {
	if m.Has(field) {
		return m.messages[field][0]
	}

	return ""
}

func (m *MessageBag) Clear() {
	m.messages = make(map[string][]string)
}
