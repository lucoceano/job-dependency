## Go lang job dependency

For this implementation was taking into account that a job can't be depended twice.
The general idea was to increment the implementation with small steps, following each use case.

For the ordering part, it searchs for each non depended job and follows it dependency tree recursively, removing each dependency from being checked twice.

### Test

```
go test
```

### Bechmarch

```
go test -bench=.
```


#### golang version

go version go1.12.5 darwin/amd64
