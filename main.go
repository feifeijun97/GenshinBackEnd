package main

import (
	// "github.com/feifeijun97/GenshinBackEnd/repository"

	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"

	"github.com/feifeijun97/GenshinBackEnd/modules/character/characterpb"
	"github.com/feifeijun97/GenshinBackEnd/repository"
)

//handle the routes for API request
// func apiRouter() *chi.Mux {
// 	router := chi.NewRouter()
// 	router.Use(middleware.Logger)
// 	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("welcome"))
// 	})

// 	return router
// }

type server struct {
	characterpb.UnimplementedCharacterListServiceServer
}

func main() {
	repository.ConnectToPostgreDb()
	// c := character.Character{}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to lsiten: %v", err)
	}
	fmt.Println("Successfully establish server on 0.0.0.0:50051")

	s := grpc.NewServer()
	characterpb.RegisterCharacterListServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CharacterList(ctx context.Context, in *characterpb.CharacterListRequest) (*characterpb.CharacterListResponse, error) {
	fmt.Println("Received a request: ", in.GetName())
	return &characterpb.CharacterListResponse{Name: "Amber"}, nil
}
