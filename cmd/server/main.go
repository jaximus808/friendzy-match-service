package main

import (
	"log"
	"net"

	"github.com/jaximus808/friendzy-match-service/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Supabase client
	supabase, err := db.NewSupabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialize Supabase: %v", err)
	}

	// Create service instance
	svc := service.NewService(supabase)

	// Initialize gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, svc)

	log.Printf("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
