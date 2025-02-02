/*
 * pagarmecoreapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package transfers_pkg

import "github.com/tradersclub/pagarme-core-api-go/src/pagarmecoreapi_lib/configuration_pkg"
import "github.com/tradersclub/pagarme-core-api-go/src/pagarmecoreapi_lib/models_pkg"

/*
 * Interface for the TRANSFERS_IMPL
 */
type TRANSFERS interface {
	GetTransferById(string) (*models_pkg.GetTransfer, error)

	CreateTransfer(*models_pkg.CreateTransfer) (*models_pkg.GetTransfer, error)

	GetTransfers() (*models_pkg.ListTransfers, error)
}

/*
 * Factory for the TRANSFERS interaface returning TRANSFERS_IMPL
 */
func NewTRANSFERS(config configuration_pkg.CONFIGURATION) *TRANSFERS_IMPL {
	client := new(TRANSFERS_IMPL)
	client.config = config
	return client
}
