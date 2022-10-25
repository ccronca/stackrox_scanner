package types

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// Timestamp is an alias for the timestamppb.Timestamp type
type Timestamp = timestamppb.Timestamp

// TimestampProto converts a time.Time to a Timestamp proto. If the resulting Timestamp is not valid,
// an error is returned.
func TimestampProto(t time.Time) (*Timestamp, error) {
	ts := timestamppb.New(t)
	return ts, ts.CheckValid()
}
