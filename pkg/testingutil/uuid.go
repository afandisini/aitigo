package testingutil

type UUIDGenerator interface {
	New() string
}

type StaticUUID struct {
	Value string
}

func (s StaticUUID) New() string {
	return s.Value
}
