package main

import (
	// "github.com/feifeijun97/GenshinBackEnd/repository"

	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"

	"github.com/feifeijun97/GenshinBackEnd/modules/character"
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
	// c.GetCharacterById(1)
	// fmt.Println(c)
	// character.GenerateCharactersFromJson("src/data/english/characters")
	// character.CreateCharacterPotraitImages("C:\\Users\\user\\Downloads\\api-mistress")
	// os.Exit(1)
	// host a HTTP server at 3000
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/images", func(r chi.Router) {
		//sub route under /images
		//to detemine the category of image user wan to view
		//Example : characters, weapons, items, materials and so on
		r.Route("/characters", func(r chi.Router) {
			r.Route("/{characterId}", func(r chi.Router) {
				r.Get("/", getCharacterImage)

			})
		})
	})

	r.Get("/images/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcomes"))
	})
	fmt.Println("Successfully establish http server on 0.0.0.0:3000")

	go http.ListenAndServe(":3000", r)

	//host the GRPC server at port 50051
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
	c := character.GetCharaacterList(in)

	return &characterpb.CharacterListResponse{Characters: c}, nil
}

// CHI ROUTE functions
func getCharacterImage(w http.ResponseWriter, r *http.Request) {
	// characterId := chi.URLParam(r, "characterId")
	// img, err := common.RenderImage(common.CharacterCategories, characterId)
	// if err != nil {
	// 	http.Error(w, http.StatusText(404), 404)
	// 	return
	// }

	// w.Write([]byte(fmt.Sprintf("title:%s", img)))
	imgFile, err := os.Open("assets/images/elements/normal/pyro.png") // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

	// png.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	// Embed into an html without PNG file
	img2html := "<html><body><img src=\"data:image/png;base64," + imgBase64Str + "\" /></body></html>"

	w.Write([]byte(fmt.Sprintf(img2html)))
}
