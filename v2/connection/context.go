//
// DISCLAIMER
//
// Copyright 2017 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Jakub Wierzbowski
//

package connection

import (
	"context"
	"time"
)

type ContextKey string

const (
	keyUseQueueTimeout ContextKey = "arangodb-use-queue-timeout"
	keyMaxQueueTime    ContextKey = "arangodb-max-queue-time-seconds"
)

// contextOrBackground returns the given context if it is not nil.
// Returns context.Background() otherwise.
func contextOrBackground(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	return context.Background()
}

// WithArangoQueueTimeout is used to enable Queue timeout on the server side.
// If WithArangoQueueTime is used then its value takes precedence in other case value of ctx.Deadline will be taken
func WithArangoQueueTimeout(parent context.Context, useQueueTimeout bool) context.Context {
	return context.WithValue(contextOrBackground(parent), keyUseQueueTimeout, useQueueTimeout)
}

// WithArangoQueueTime defines max queue timeout on the server side.
func WithArangoQueueTime(parent context.Context, duration time.Duration) context.Context {
	return context.WithValue(contextOrBackground(parent), keyMaxQueueTime, duration)
}
