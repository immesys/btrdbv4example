# btrdbv4example

To use this, set the server you are querying, and run it. e.g

```
go get github.com/immesys/btrdbv4example
export BTRDB_ENDPOINTS=my.server.com:4410
btrdbv4example
```


You should get an output something like:

```
Collection upmu/psl_alameda stream C1ANG has 302400 points
Collection upmu/psl_alameda stream C1MAG has 302400 points
Collection upmu/psl_alameda stream C2ANG has 302400 points
Collection upmu/psl_alameda stream C2MAG has 302400 points
Collection upmu/psl_alameda stream C3ANG has 302400 points
Collection upmu/psl_alameda stream C3MAG has 302400 points
Collection upmu/psl_alameda stream L1ANG has 302400 points
Collection upmu/psl_alameda stream L1MAG has 302400 points
Collection upmu/psl_alameda stream L2ANG has 302400 points
Collection upmu/psl_alameda stream L2MAG has 302400 points
Collection upmu/psl_alameda stream L3ANG has 302400 points
Collection upmu/psl_alameda stream L3MAG has 302400 points
Collection upmu/psl_alameda stream LSTATE has 302400 points
```

For more information on the btrdb v4 API look at https://godoc.org/gopkg.in/btrdb.v4
