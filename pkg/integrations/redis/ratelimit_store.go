package redis

import (
	"context"
	"fmt"
	"time"

	"aitigo/pkg/middleware/ratelimit"
	"github.com/redis/go-redis/v9"
)

type RateLimitStore struct {
	client *redis.Client
	prefix string
}

var tokenBucketScript = redis.NewScript(`
local key = KEYS[1]
local rate = tonumber(ARGV[1])
local burst = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local per = tonumber(ARGV[4])

local data = redis.call("HMGET", key, "tokens", "ts")
local tokens = tonumber(data[1]) or burst
local ts = tonumber(data[2]) or now

local delta = math.max(0, now - ts)
local refill = (delta / per) * rate
tokens = math.min(burst, tokens + refill)

local allowed = tokens >= 1
if allowed then
  tokens = tokens - 1
end

redis.call("HMSET", key, "tokens", tokens, "ts", now)
redis.call("EXPIRE", key, math.ceil(per * 2))
return {allowed, tokens}
`)

func NewRateLimitStore(client *redis.Client, prefix string) *RateLimitStore {
	if prefix == "" {
		prefix = "aitigo:ratelimit"
	}
	return &RateLimitStore{client: client, prefix: prefix}
}

func (s *RateLimitStore) Allow(ctx context.Context, key string, limit ratelimit.Limit) (bool, time.Duration, error) {
	if s.client == nil {
		return false, 0, fmt.Errorf("redis client is nil")
	}
	if limit.Per <= 0 {
		limit.Per = time.Second
	}
	fullKey := fmt.Sprintf("%s:%s", s.prefix, key)
	now := float64(time.Now().UnixNano()) / float64(time.Second)

	result, err := tokenBucketScript.Run(ctx, s.client, []string{fullKey}, limit.Rate, limit.Burst, now, limit.Per.Seconds()).Result()
	if err != nil {
		return false, 0, err
	}

	values, ok := result.([]interface{})
	if !ok || len(values) < 2 {
		return false, 0, fmt.Errorf("unexpected redis response")
	}

	allowed := toInt(values[0])
	tokens := toFloat(values[1])
	if allowed == 1 {
		return true, 0, nil
	}

	ratePerSecond := float64(limit.Rate) / limit.Per.Seconds()
	retrySeconds := (1 - tokens) / ratePerSecond
	if retrySeconds < 0 {
		retrySeconds = 0
	}
	retryAfter := time.Duration(retrySeconds * float64(time.Second))
	return false, retryAfter, nil
}

func toInt(value interface{}) int64 {
	switch v := value.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case float64:
		return int64(v)
	default:
		return 0
	}
}

func toFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case int64:
		return float64(v)
	case int:
		return float64(v)
	default:
		return 0
	}
}
