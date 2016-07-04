package mbus

import (
    "fmt"

    "github.com/streadway/amqp"
)

const (
    ExchangeName = "daocloud"
    ExchangeType = "topic"
)

type Bus interface {
    Subscribe(topic string, handler Handler) error
    Publish(topic string, body []byte) error
    Unsubscribe(topic string) error
}

type Handler func([]byte)

type Client struct {
    conn *amqp.Connection
    channel *amqp.Channel

    // fields used by receiver
    queue string
    handlers map[string]Handler
}

func Connect(host string) (Bus, error) {
    url := fmt.Sprintf("amqp://guest:guest@%s/", host)
    c, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := c.Channel()
    if err != nil {
        return nil, err
    }

    err = ch.ExchangeDeclare(
        ExchangeName, // name
        ExchangeType, // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        ch.Close()
        return nil, err
    }

    return &Client{conn: c, channel: ch}, nil
}

func (c *Client) Subscribe(topic string, handler Handler) error {
    if c.queue == "" {
        q, err := c.channel.QueueDeclare("", false, false, true, false, nil)
        if err != nil {
            return err
        }

        msgs, err := c.channel.Consume(q.Name, "", true, false, false, false, nil)
        if err != nil {
            return err
        }

        go func() {
            for msg := range msgs {
                if fun, ok := c.handlers[msg.RoutingKey]; ok {
                    fun(msg.Body)
                }
            }
        } ()

        c.queue = q.Name
    }

    if err := c.channel.QueueBind(c.queue, topic, ExchangeName, false, nil); err != nil {
        return err
    }

    if c.handlers == nil {
        c.handlers = make(map[string]Handler)
    }

    c.handlers[topic] = handler
    return nil
}

func (c *Client) Unsubscribe(topic string) error {
    if err := c.channel.QueueUnbind(c.queue, topic, ExchangeName, nil); err != nil {
        return err
    }

    delete(c.handlers, topic)
    return nil

}

func (c *Client) Publish(topic string, body []byte) error {
    return c.channel.Publish(
        ExchangeName, topic, false, false,
        amqp.Publishing{ContentType: "text/plain", Body: []byte(body)},
    )
}

func (c *Client) Close() {
    c.channel.Close()
    c.conn.Close()
}
