package serv

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	openapi "github.com/iaPlotnikovv/New-breath/generated/go"
)

func (b *BannerService) BannerPost(ctx context.Context, bannerPostRequest openapi.BannerPostRequest, token string) (openapi.ImplResponse, error) {

	fmt.Println("TODO POST")
	// TODO - update BannerPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, BannerPost201Response{}) or use other options such as http.Ok ...
	// return Response(201, BannerPost201Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, UserBannerGet400Response{}) or use other options such as http.Ok ...
	// return Response(400, UserBannerGet400Response{}), nil

	// TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	// return Response(401, nil),nil

	// TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	// return Response(403, nil),nil

	// TODO: Uncomment the next line to return response Response(500, UserBannerGet400Response{}) or use other options such as http.Ok ...
	// return Response(500, UserBannerGet400Response{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("BannerPost method not implemented")

}
