package protoclient

import (
	"google.golang.org/grpc"
	pb "github.com/rajverve/protobuf"
)

type ProtoClient struct {
	client pb.SegmentationClient
}

func (c *ProtoClient) Connect(conn *grpc.ClientConn)  {
	return &ProtoClient{
		client: conn,
	}
}
