package main

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/influxdata/influxdb/client/v2"
	pb "github.com/iochti/point-service/proto"
	"golang.org/x/net/context"
)

// PointSvc implements the gRPCs methods described in proto/point.proto
type PointSvc struct {
	Db DataLayerInterface
}

// CreatePoint creates a point in the database
func (p *PointSvc) CreatePoint(ctx context.Context, in *pb.Point) (*empty.Empty, error) {
	fields := make(map[string]interface{}, len(in.GetFields()))
	for k, v := range in.GetFields() {
		msg, err := AnyToMessageType(v)
		if err != nil {
			return nil, err
		}
		val := ExtractDatasFromMessage(v.GetTypeUrl(), msg)
		fields[k] = val
	}
	pt, err := client.NewPoint(in.GetUser(), in.GetTags(), fields, time.Now())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	err = p.Db.CreatePoint(pt)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

// GetPointsByThing gets all points for a given user & thing
func (p *PointSvc) GetPointsByThing(ctx context.Context, in *pb.ThingId) (*pb.InfluxResult, error) {
	cmd := fmt.Sprintf("SELECT * FROM \"%s\" WHERE \"thing_id\" = '%s'",
		in.GetUser(),
		in.GetThingId(),
	)
	if in.GetEnd().GetSeconds() != 0 {
		cmd = cmd + fmt.Sprintf(" AND time < '%s'", time.Unix(in.GetEnd().GetSeconds(), 0).Format(time.RFC3339Nano))
	}
	if in.GetStart().GetSeconds() != 0 {
		cmd = cmd + fmt.Sprintf(" AND time > '%s'", time.Unix(in.GetStart().GetSeconds(), 0).Format(time.RFC3339Nano))
	}
	fmt.Println(cmd)
	res, err := p.Db.QueryDB(cmd)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return &pb.InfluxResult{Item: bytes}, nil
}

// GetPointsByGroup get all points for a given group
func (p *PointSvc) GetPointsByGroup(ctx context.Context, in *pb.GroupId) (*pb.InfluxResult, error) {
	cmd := fmt.Sprintf("SELECT * FROM \"%s\" WHERE group_id = '%s' AND time < '%s' AND time > '%s'",
		in.GetUser(),
		in.GetGroupId(),
		time.Unix(in.GetEnd().GetSeconds(), 0).Format(time.RFC3339Nano),
		time.Unix(in.GetStart().GetSeconds(), 0).Format(time.RFC3339Nano),
	)
	res, err := p.Db.QueryDB(cmd)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return &pb.InfluxResult{Item: bytes}, nil

}

func AnyToMessageType(elmt *any.Any) (proto.Message, error) {
	var value proto.Message
	switch elmt.GetTypeUrl() {
	case "iochti.com/string":
		value = &pb.StringPoint{}
	case "iochti.com/float":
		value = &pb.FloatPoint{}
	case "iochti.com/int":
		value = &pb.IntegerPoint{}
	case "iochti.com/bool":
		value = &pb.BoolPoint{}
	case "iochti.com/duration":
		value = &pb.DurationPoint{}
	case "iochti.com/date-time":
		value = &pb.DateTimePoint{}
	}
	err := proto.Unmarshal(elmt.GetValue(), value)
	return value, err
}

func ExtractDatasFromMessage(typeString string, msg proto.Message) interface{} {
	switch typeString {
	case "iochti.com/string":
		return msg.(*pb.StringPoint).GetValue()
	case "iochti.com/float":
		return msg.(*pb.FloatPoint).GetValue()
	case "iochti.com/int":
		return msg.(*pb.IntegerPoint).GetValue()
	case "iochti.com/bool":
		return msg.(*pb.BoolPoint).GetValue()
	case "iochti.com/duration":
		return msg.(*pb.DurationPoint).GetValue()
	case "iochti.com/date-time":
		return msg.(*pb.DateTimePoint).GetValue()
	}
	return nil
}
