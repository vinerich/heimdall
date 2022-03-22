package authenticators

import (
	"context"

	"github.com/dadrus/heimdall/pipeline"
)

var _ Authenticator = new(noopAuthenticator)

func newNoopAuthenticator(id string) (*noopAuthenticator, error) {
	return &noopAuthenticator{
		id: id,
	}, nil
}

type noopAuthenticator struct {
	id string
}

func (a *noopAuthenticator) Id() string {
	return a.id
}

func (*noopAuthenticator) Authenticate(ctx context.Context, as pipeline.AuthDataSource, sc *pipeline.SubjectContext) error {
	return nil
}
