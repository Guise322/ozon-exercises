package sub_client

type subClient struct{}

func NewSubClient() (*subClient, error) {
	return &subClient{}, nil
}
