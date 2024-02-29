package product

import (
	"fmt"
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *productService) Create(ctx *fiber.Ctx) (*model.Response, error) {
	// parse body
	payload := new(model.UpsertProductRequest)
	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// TODO: validate image url
	fmt.Println("len(payload.ImageUrl) : ", len(payload.ImageUrl))
	fmt.Println("len(payload.ImageUrl) : ", len(payload.ImageUrl))
	if len(payload.ImageUrl) > 0 {
		fmt.Println("payload.ImageUrl : ", payload.ImageUrl)
		fmt.Println("len(payload.ImageUrl) : ", len(payload.ImageUrl))
		for idx, url := range payload.ImageUrl {
			fmt.Println("idx : ", idx, " url : ", url)

			if !validatorhelper.IsValidUrl(strings.Trim(url, " ")) {
				return responsehelper.ResponseErrorValidation(fiber.Map{"imageUrl": fmt.Sprintf("Invalid image url at index %d", idx)}), nil
			}
		}
	}

	// TODO: validate video url
	// TODO: validate media file
	mediaFiles, _ := ctx.FormFile("mediaFile")
	fmt.Println("mediaFiles.Size : ", mediaFiles.Size)

	return responsehelper.Response200("Successfuly created product", nil, nil), nil
}
