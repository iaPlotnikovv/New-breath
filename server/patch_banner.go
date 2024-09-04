package serv

import (
	"context"
	"errors"
	"net/http"

	openapi "github.com/iaPlotnikovv/New-breath/generated/go"
)

func (s *BannerService) BannerIdPatch(ctx context.Context, id int32, bannerIdPatchRequest openapi.BannerIdPatchRequest, token string) (openapi.ImplResponse, error) {
	// TODO - update BannerIdPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(400, UserBannerGet400Response{}) or use other options such as http.Ok ...
	// return Response(400, UserBannerGet400Response{}), nil

	// TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	// return Response(401, nil),nil

	// TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	// return Response(403, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(500, UserBannerGet400Response{}) or use other options such as http.Ok ...
	// return Response(500, UserBannerGet400Response{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("BannerIdPatch method not implemented")
}
