package mbus

import (
)

type MockClient struct {
    handlers map[string]Handler
}


func MockConnect() (Bus, error) {
    return &MockClient{}, nil
}

func (c *MockClient) Subscribe(topic string, handler Handler) error {
    if c.handlers == nil {
        c.handlers = make(map[string]Handler)
    }

    c.handlers[topic] = handler
    return nil
}

func (c *MockClient) Unsubscribe(topic string) error {
    delete(c.handlers, topic)
    return nil
}

func (c *MockClient) Publish(topic string, body []byte) error {
    if c.handlers == nil {
        return nil
    }

    if fun, ok := c.handlers[topic]; ok {
        fun(body)
    }

    return nil
}
