package file

type GrpcServer struct {
}

func New() (*GrpcServer, error) {
	return &GrpcServer{}, nil
}

func (s *GrpcServer) Send() {

}
