package http

import (
	"context"
)

type ContextualizeRequest func(currentCtx context.Context) context.Context
