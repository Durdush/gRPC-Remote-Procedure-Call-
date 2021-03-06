package informatica

import (
	"context"
	"fmt"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/mongodb"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/proto"
	"log"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Informatica struct {
	ID       objectid.ObjectID `bson:"_id,omitempty"`
	Sequence int32             `bson:"sequence"`
	Title    string            `bson:"title"`
	Info     string            `bson:"info"`
	HostName string            `bson:"host_name"`
}

func InsertDataInInformatica(req *informaticapb.Informatica) error {
	data := Informatica{
		Sequence: req.GetSequence(),
		Title:    req.GetTitle(),
		Info:     req.GetInfo(),
		HostName: req.GetHostName(),
	}

	res, err := mongodb.CreateCollection("informaticaData").InsertOne(context.Background(), data)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln("Internal MongoDb Error:", err))
	}

	oid, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln("Internal oid Error:", ok))
	}

	log.Println("oid: ", oid)

	return nil
}
