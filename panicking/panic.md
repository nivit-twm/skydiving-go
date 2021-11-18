# Defer, Panic, Recover
ref: https://go.dev/blog/defer-panic-and-recover

## Pitfall
log.Fatal or os.Exit does not respect deferred calls whereas t.Fatal does, where t is a testing.T.
