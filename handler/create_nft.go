package handler

import (
	"fmt"
	"net/http"

	"github.com/shunsukenagashima/go_minting_api/gen/api"
)

type CreateNft struct {

}

var _ api.ServerInterface = (*CreateNft)(nil)

func (c *CreateNft) CreateNft(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a endpont for creating new nft")
}
