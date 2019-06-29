## Go lang job dependency

For this implementation was taking into account that a job can't be depended twice.
The general idea was to increment the implementation wiht small steps, following each use case.

For the ordering part, it searchs for the each non depended job and follows it dependency tree recursively, removing each dependency from being checked twice.

### Test

```
go test
```

### Bechmarch

```
go test -bench=.
```
