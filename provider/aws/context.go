package aws

import (
	"context"

	"github.com/convox/praxis/types"
)

func (p *Provider) WithContext(ctx context.Context) types.Provider {
	q := p
	q.Context = ctx
	return q
}
