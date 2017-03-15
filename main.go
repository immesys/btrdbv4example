package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"gopkg.in/btrdb.v4"
)

//This format is like
//"02 Jan 06 15:04 MST"
func dateToNs(date string) (ns int64) {
	time, err := time.Parse(time.RFC822, date)
	if err != nil {
		fmt.Printf("Could not parse time %q\n", date)
		os.Exit(2)
	}
	return time.UnixNano()
}
func main() {
	//If you want to put a timeout on your operations, the context is how you do it
	//background means no timeout
	ctx := context.Background()

	//You need to set the BTRDB_ENDPOINTS environment variable
	db, err := btrdb.Connect(ctx, btrdb.EndpointsFromEnv()...)
	if err != nil {
		fmt.Printf("Could not connect: %v\n", err)
		os.Exit(1)
	}

	// Lets get all streams from PSL and show how many points there are:
	psl, err := db.LookupStreams(ctx, "upmu/psl_alameda", true, nil, nil)
	for _, stream := range psl {
		col, err := stream.Collection(ctx)
		if err != nil {
			panic(err) // This won't normally happen
		}
		tags, err := stream.Tags(ctx)
		if err != nil {
			panic(err) // This won't normally happen
		}
		//The name of the stream (like L1MAG) comes from the name tag
		name := tags["name"]

		start := dateToNs("01 Jan 14 00:00 UTC")
		fouryears := 4 * 365 * 24 * time.Hour
		res, _, cerr := stream.Windows(ctx, //Timeouts or whatever
			start, //The time to start
			start+int64(fouryears), //The time to end
			uint64(fouryears),      //How long each window should be
			0,                      //The depth to go (0 is full resolution)
			btrdb.LatestVersion,    //The version of the stream to use
		)
		//We get back a stream of statistical windows, but we only need the one
		singleResult := <-res //read just one window

		//The cerr variable is a channel, just read the error from it and check
		//that it is nil. An empty stream does not support windows, so it will
		//return an error saying "stream is empty, cannot do windows"
		if err := <-cerr; err != nil {
			fmt.Printf("Collection %s stream %s is empty/error\n", col, name)
			continue
		}
		fmt.Printf("Collection %s stream %s has %d points\n", col, name, singleResult.Count)
	}
}
