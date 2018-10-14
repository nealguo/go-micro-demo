package main

import (
	pb "consignment-service/proto/consignment"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
	vessel "vessel-service/proto/vessel"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment) {
	return repo.consignments
}

type service struct {
	repo         Repository
	vesselClient vessel.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	response, err := s.vesselClient.FindAvailable(context.Background(), &vessel.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel:%s \n", response.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = response.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment

	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &ConsignmentRepository{}

	// Create a new service. Optionally include some options here
	srv := micro.NewService(
		// This name must match the package name give in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	client := vessel.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, client})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
