/*
Copyright The Ratify Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package context

import (
	"context"
	"fmt"
)

type contextKey string

const (
	ContextKeyNamespace  = contextKey("namespace")
	VerificationCounters = contextKey("verificationCounters")
)

// SetContextWithNamespace embeds namespace to the context.
func SetContextWithNamespace(ctx context.Context, namespace string) context.Context {
	return context.WithValue(ctx, ContextKeyNamespace, namespace)
}

// Initializes the verification counters map in the context, the map is heriarchical with the first key being the verifier name and the second key being image name, and last is the counter.
func InitContextWithVerificationCounters(ctx context.Context) context.Context {
	return context.WithValue(ctx, VerificationCounters, map[string]interface{}{})
}

// Retrieves the verification counters map from the context.
func GetVerificationCountersFromContext(ctx context.Context) map[string]interface{} {
	counters := ctx.Value(VerificationCounters)
	if counters == nil {
		return nil
	}
	return counters.(map[string]interface{})
}

// GetNamespace returns the embedded namespace from the context.
func GetNamespace(ctx context.Context) string {
	namespace := ctx.Value(ContextKeyNamespace)
	if namespace == nil {
		return ""
	}
	return namespace.(string)
}

// CreateCacheKey creates a new cache key prefixed with embedded namespace.
func CreateCacheKey(ctx context.Context, key string) string {
	namespace := ctx.Value(ContextKeyNamespace)
	if namespace == nil {
		return key
	}

	namespaceStr := namespace.(string)
	if namespaceStr == "" {
		return key
	}
	return fmt.Sprintf("%s:%s", namespaceStr, key)
}
