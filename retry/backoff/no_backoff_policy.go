package backoff

import "github.com/liuhailove/gretry/retry"

// NoBackOffPolicy 没有任何操作的策略，所有的重试操作处理都不会暂停
type NoBackOffPolicy struct {
}

func (n NoBackOffPolicy) Start(_ retry.RtyContext) BackoffContext {
	return nil
}

func (n NoBackOffPolicy) BackOff(_ BackoffContext) {
}

func (n NoBackOffPolicy) String() string {
	return "NoBackOffPolicy []"
}
