/*
 * pagarmecoreapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package transactions_pkg

import "github.com/tradersclub/pagarme-core-api-go/src/pagarmecoreapi_lib/configuration_pkg"
import "github.com/tradersclub/pagarme-core-api-go/src/pagarmecoreapi_lib/models_pkg"

/*
 * Interface for the TRANSACTIONS_IMPL
 */
type TRANSACTIONS interface {
	GetTransaction(string) (*models_pkg.GetTransactionResponse, error)
}

/*
 * Factory for the TRANSACTIONS interaface returning TRANSACTIONS_IMPL
 */
func NewTRANSACTIONS(config configuration_pkg.CONFIGURATION) *TRANSACTIONS_IMPL {
	client := new(TRANSACTIONS_IMPL)
	client.config = config
	return client
}
