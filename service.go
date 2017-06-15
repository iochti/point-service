package main

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/iochti/point-service/proto"
	"golang.org/x/net/context"
)

// PointSvc implements the gRPCs methods described in proto/point.proto
type PointSvc struct {
	Db DataLayerInterface
}

// CreatePoint creates a point in the database
func (p *PointSvc) CreatePoint(ctx context.Context, in *pb.Point) (*empty.Empty, error) {
	fields, err := anyToInterface(in.GetFields())
	if err != nil {
		return nil, err
	}
	fmt.Println(fields)
	return nil, nil
}

// GetPointsByThing gets all points for a given user & thing
func (p *PointSvc) GetPointsByThing(in *pb.ThingId, stream pb.PointSvc_GetPointsByThingServer) error {
	return nil
}

// GetPointsByGroup get all points for a given group
func (p *PointSvc) GetPointsByGroup(in *pb.GroupId, stream pb.PointSvc_GetPointsByGroupServer) error {
	return nil

}

func anyToInterface(value map[string]*any.Any) (map[string]interface{}, error) {
	var res map[string]interface{}
	for k, v := range value {
		var x ptypes.DynamicAny
		if err := ptypes.UnmarshalAny(v, &x); err != nil {
			return res, err
		}
		res[k] = x.Message
	}
	return res, nil
}
