package serv

import (
	"context"
	"errors"
	"net/http"

	openapi "github.com/iaPlotnikovv/New-breath/generated/go"
)

// BannerGet - Получение всех баннеров c фильтрацией по фиче и/или тегу
func (s *BannerService) BannerGet(ctx context.Context, token string, featureId int32, tagId int32, limit int32, offset int32) (openapi.ImplResponse, error) {
	// TODO - update BannerGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []BannerGet200ResponseInner{}) or use other options such as http.Ok ...
	// return Response(200, []BannerGet200ResponseInner{}), nil

	// TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	// return Response(401, nil),nil

	// TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	// return Response(403, nil),nil

	// TODO: Uncomment the next line to return response Response(500, UserBannerGet400Response{}) or use other options such as http.Ok ...
	// return Response(500, UserBannerGet400Response{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("BannerGet method not implemented")
}
