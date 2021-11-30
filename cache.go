package recdrive

type Cache interface {
	QueryPath(path string) (bool, string, error)
	SavePair(path string, id string) error
	DropPair(path string, id string) error
}

type NoCache struct{}

func (c *NoCache) QueryPath(path string) (bool, string, error) {
	return false, "", nil
}

func (c *NoCache) SavePair(path string, id string) error {
	return nil
}

func (c *NoCache) DropPair(path string, id string) error {
	return nil
}
