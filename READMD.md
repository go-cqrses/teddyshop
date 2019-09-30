# Teddy Shop

This is a dummy project to demonstrate https://github.com/go-cqrses libraries.

Welcome to all kinds of contributions :)

## Business logic

A customer wants to buy a teddy, the customer must register and then create an order and pay for the order. Once paid the build floor workers should make the order, once made they should report the order as dispatched.

## Components

### `apps/shopfloor`

todo

### `apps/buildfloor`

todo

### `apis/identities`

todo

### `apis/billing`

todo

### `apis/orders`

todo

## External services.

- DigitalOcean (required infra bootstrap scripts to start Kubernetes)
- CoackroachDB (event store backend)
- NATS (event publisher)
