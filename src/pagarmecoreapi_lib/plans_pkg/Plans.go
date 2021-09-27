/*
 * pagarmecoreapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package plans_pkg

import "time"
import "github.com/tradersclub/pagarme-core-api-go/src/pagarmecoreapi_lib/configuration_pkg"
import "github.com/tradersclub/pagarme-core-api-go/src/pagarmecoreapi_lib/models_pkg"

/*
 * Interface for the PLANS_IMPL
 */
type PLANS interface {
	GetPlan(string) (*models_pkg.GetPlanResponse, error)

	UpdatePlanMetadata(string, *models_pkg.UpdateMetadataRequest, *string) (*models_pkg.GetPlanResponse, error)

	UpdatePlanItem(string, string, *models_pkg.UpdatePlanItemRequest, *string) (*models_pkg.GetPlanItemResponse, error)

	CreatePlanItem(string, *models_pkg.CreatePlanItemRequest, *string) (*models_pkg.GetPlanItemResponse, error)

	GetPlanItem(string, string) (*models_pkg.GetPlanItemResponse, error)

	CreatePlan(*models_pkg.CreatePlanRequest, *string) (*models_pkg.GetPlanResponse, error)

	DeletePlanItem(string, string, *string) (*models_pkg.GetPlanItemResponse, error)

	GetPlans(*int64, *int64, *string, *string, *string, *time.Time, *time.Time) (*models_pkg.ListPlansResponse, error)

	UpdatePlan(string, *models_pkg.UpdatePlanRequest, *string) (*models_pkg.GetPlanResponse, error)

	DeletePlan(string, *string) (*models_pkg.GetPlanResponse, error)
}

/*
 * Factory for the PLANS interaface returning PLANS_IMPL
 */
func NewPLANS(config configuration_pkg.CONFIGURATION) *PLANS_IMPL {
	client := new(PLANS_IMPL)
	client.config = config
	return client
}
