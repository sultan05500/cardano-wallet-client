// Package wallet provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package wallet

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// ScriptTemplateValue defines model for ScriptTemplateValue.
type ScriptTemplateValue interface{}

// ScriptValue defines model for ScriptValue.
type ScriptValue interface{}

// TransactionMetadataValue defines model for TransactionMetadataValue.
type TransactionMetadataValue interface{}

// PostAnyAddressJSONBody defines parameters for PostAnyAddress.
type PostAnyAddressJSONBody struct {
	Payment *interface{} `json:"payment,omitempty"`
	Stake   *interface{} `json:"stake,omitempty"`

	// Script validation level. Required validation sifts off scripts that would not
	// be accepted by the ledger. Recommended level filters out scripts that do not pass
	// required validation and additionally when:
	//   * 'all' is non-empty
	//   * there are redundant timelocks in a given level
	//   * there are no duplicated verification keys in a given level
	//   * 'at_least' coeffcient is positive
	//   * 'all', 'any' are non-empty and `'at_least' has no less elements in the list
	//      than the coeffcient after timelocks are filtered out.
	Validation *string `json:"validation,omitempty"`
}

// PostByronWalletJSONBody defines parameters for PostByronWallet.
type PostByronWalletJSONBody interface{}

// PutByronWalletJSONBody defines parameters for PutByronWallet.
type PutByronWalletJSONBody struct {
	Name *string `json:"name,omitempty"`
}

// ListByronAddressesParams defines parameters for ListByronAddresses.
type ListByronAddressesParams struct {

	// An optional filter on the address state.
	State *string `json:"state,omitempty"`
}

// CreateAddressJSONBody defines parameters for CreateAddress.
type CreateAddressJSONBody struct {

	// An address derivation index.
	AddressIndex *float32 `json:"address_index,omitempty"`

	// A master passphrase to lock and protect the wallet for sensitive operation (e.g. sending funds)
	Passphrase string `json:"passphrase"`
}

// ImportAddressesJSONBody defines parameters for ImportAddresses.
type ImportAddressesJSONBody struct {

	// The imported addresses.
	Addresses []string `json:"addresses"`
}

// ByronSelectCoinsJSONBody defines parameters for ByronSelectCoins.
type ByronSelectCoinsJSONBody struct {

	// A list of target outputs
	Payments []struct {
		Address string `json:"address"`

		// Coins, in Lovelace. Only relates to 'Ada'. Refer to `assets` for multi-assets wallets instead.
		Amount struct {
			Quantity int    `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"amount"`

		// A flat list of assets.
		Assets *[]struct {

			// The asset on-chain type which acts as a sub-identifier within a
			// policy. Although we call it "asset name", the value needn't be
			// text, and it could even be empty.
			//
			// For policies with a single fungible asset item, asset name is
			// typically an empty string.
			//
			// This value can be up to 32 bytes of arbitrary data (which is 64
			// hexadecimal digits).
			AssetName string `json:"asset_name"`

			// A unique identifier of the asset's monetary policy. The policy
			// controls how assets of this kind are created and destroyed.
			//
			// The contents are the blake2b-224 hash of the monetary policy
			// script, encoded in hexadecimal.
			PolicyId string `json:"policy_id"`

			// Number of assets for the given `policy_id` and `asset_name`.
			Quantity int `json:"quantity"`
		} `json:"assets,omitempty"`
	} `json:"payments"`
}

// MigrateByronWalletJSONBody defines parameters for MigrateByronWallet.
type MigrateByronWalletJSONBody struct {

	// The recipient addresses.
	Addresses []string `json:"addresses"`

	// The wallet's master passphrase.
	Passphrase string `json:"passphrase"`
}

// PutByronWalletPassphraseJSONBody defines parameters for PutByronWalletPassphrase.
type PutByronWalletPassphraseJSONBody struct {

	// A master passphrase to lock and protect the wallet for sensitive operation (e.g. sending funds).
	NewPassphrase string `json:"new_passphrase"`

	// The current passphrase if present.
	OldPassphrase *string `json:"old_passphrase,omitempty"`
}

// PostByronTransactionFeeJSONBody defines parameters for PostByronTransactionFee.
type PostByronTransactionFeeJSONBody struct {

	// A list of target outputs
	Payments []struct {
		Address string `json:"address"`

		// Coins, in Lovelace. Only relates to 'Ada'. Refer to `assets` for multi-assets wallets instead.
		Amount struct {
			Quantity int    `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"amount"`

		// A flat list of assets.
		Assets *[]struct {

			// The asset on-chain type which acts as a sub-identifier within a
			// policy. Although we call it "asset name", the value needn't be
			// text, and it could even be empty.
			//
			// For policies with a single fungible asset item, asset name is
			// typically an empty string.
			//
			// This value can be up to 32 bytes of arbitrary data (which is 64
			// hexadecimal digits).
			AssetName string `json:"asset_name"`

			// A unique identifier of the asset's monetary policy. The policy
			// controls how assets of this kind are created and destroyed.
			//
			// The contents are the blake2b-224 hash of the monetary policy
			// script, encoded in hexadecimal.
			PolicyId string `json:"policy_id"`

			// Number of assets for the given `policy_id` and `asset_name`.
			Quantity int `json:"quantity"`
		} `json:"assets,omitempty"`
	} `json:"payments"`
}

// ListByronTransactionsParams defines parameters for ListByronTransactions.
type ListByronTransactionsParams struct {

	// An optional start time in ISO 8601 date-and-time format. Basic and
	// extended formats are both accepted. Times can be local (with a
	// timezone offset) or UTC.
	//
	// If both a start time and an end time are specified, then the start
	// time must not be later than the end time.
	//
	// Example: `2008-08-08T08:08:08Z`
	Start *string `json:"start,omitempty"`

	// An optional end time in ISO 8601 date-and-time format. Basic and
	// extended formats are both accepted. Times can be local (with a
	// timezone offset) or UTC.
	//
	// If both a start time and an end time are specified, then the start
	// time must not be later than the end time.
	//
	// Example: `2008-08-08T08:08:08Z`
	End *string `json:"end,omitempty"`

	// An optional sort order.
	Order *string `json:"order,omitempty"`
}

// PostByronTransactionJSONBody defines parameters for PostByronTransaction.
type PostByronTransactionJSONBody struct {

	// The wallet's master passphrase.
	Passphrase string `json:"passphrase"`

	// A list of target outputs
	Payments []struct {
		Address string `json:"address"`

		// Coins, in Lovelace. Only relates to 'Ada'. Refer to `assets` for multi-assets wallets instead.
		Amount struct {
			Quantity int    `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"amount"`

		// A flat list of assets.
		Assets *[]struct {

			// The asset on-chain type which acts as a sub-identifier within a
			// policy. Although we call it "asset name", the value needn't be
			// text, and it could even be empty.
			//
			// For policies with a single fungible asset item, asset name is
			// typically an empty string.
			//
			// This value can be up to 32 bytes of arbitrary data (which is 64
			// hexadecimal digits).
			AssetName string `json:"asset_name"`

			// A unique identifier of the asset's monetary policy. The policy
			// controls how assets of this kind are created and destroyed.
			//
			// The contents are the blake2b-224 hash of the monetary policy
			// script, encoded in hexadecimal.
			PolicyId string `json:"policy_id"`

			// Number of assets for the given `policy_id` and `asset_name`.
			Quantity int `json:"quantity"`
		} `json:"assets,omitempty"`
	} `json:"payments"`
}

// GetNetworkClockParams defines parameters for GetNetworkClock.
type GetNetworkClockParams struct {

	// NTP checks are cached for short duration to avoid sending too many queries to the central NTP servers. In some cases however, a client may want to force a new check.
	//
	// When this flag is set, the request **will block** until NTP server responds or will timeout after a while without any answer from the NTP server.
	ForceNtpCheck *bool `json:"forceNtpCheck,omitempty"`
}

// PutSettingsJSONBody defines parameters for PutSettings.
type PutSettingsJSONBody struct {

	// Settings
	Settings *struct {

		// Select stake pool metadata fetching strategy:
		//   - `none` - metadata is not fetched at all,
		//   - `direct` - metadata is fetched directly URLs registered on chain,
		//   - `uri` - metadata is fetched from an external Stake-Pool Metadata Aggregation Server (SMASH)
		//
		// After update existing metadata will be dropped forcing it to re-sync automatically with the new setting.
		PoolMetadataSource string `json:"pool_metadata_source"`
	} `json:"settings,omitempty"`
}

// PostSharedWalletJSONBody defines parameters for PostSharedWallet.
type PostSharedWalletJSONBody interface{}

// PatchSharedWalletInDelegationJSONBody defines parameters for PatchSharedWalletInDelegation.
type PatchSharedWalletInDelegationJSONBody struct {
	AdditionalProperties map[string]string `json:"-"`
}

// PatchSharedWalletInPaymentJSONBody defines parameters for PatchSharedWalletInPayment.
type PatchSharedWalletInPaymentJSONBody struct {
	AdditionalProperties map[string]string `json:"-"`
}

// GetCurrentSmashHealthParams defines parameters for GetCurrentSmashHealth.
type GetCurrentSmashHealthParams struct {

	// check this url for health instead of the currently configured one
	Url *string `json:"url,omitempty"`
}

// ListStakePoolsParams defines parameters for ListStakePools.
type ListStakePoolsParams struct {

	// The stake the user intends to delegate in Lovelace. Required.
	Stake int `json:"stake"`
}

// QuitStakePoolJSONBody defines parameters for QuitStakePool.
type QuitStakePoolJSONBody struct {

	// The source Byron wallet's master passphrase.
	Passphrase string `json:"passphrase"`
}

// PostMaintenanceActionJSONBody defines parameters for PostMaintenanceAction.
type PostMaintenanceActionJSONBody struct {
	MaintenanceAction string `json:"maintenance_action"`
}

// JoinStakePoolJSONBody defines parameters for JoinStakePool.
type JoinStakePoolJSONBody struct {

	// The source Byron wallet's master passphrase.
	Passphrase string `json:"passphrase"`
}

// PostWalletJSONBody defines parameters for PostWallet.
type PostWalletJSONBody interface{}

// PutWalletJSONBody defines parameters for PutWallet.
type PutWalletJSONBody struct {
	Name *string `json:"name,omitempty"`
}

// ListAddressesParams defines parameters for ListAddresses.
type ListAddressesParams struct {

	// An optional filter on the address state.
	State *string `json:"state,omitempty"`
}

// SelectCoinsJSONBody defines parameters for SelectCoins.
type SelectCoinsJSONBody interface{}

// PostAccountKeyJSONBody defines parameters for PostAccountKey.
type PostAccountKeyJSONBody struct {

	// Determines whether extended (with chain code) or normal (without chain code) key is requested
	Extended bool `json:"extended"`

	// A master passphrase to lock and protect the wallet for sensitive operation (e.g. sending funds)
	Passphrase string `json:"passphrase"`
}

// MigrateShelleyWalletJSONBody defines parameters for MigrateShelleyWallet.
type MigrateShelleyWalletJSONBody struct {

	// The recipient addresses.
	Addresses []string `json:"addresses"`

	// The wallet's master passphrase.
	Passphrase string `json:"passphrase"`
}

// PutWalletPassphraseJSONBody defines parameters for PutWalletPassphrase.
type PutWalletPassphraseJSONBody struct {

	// A master passphrase to lock and protect the wallet for sensitive operation (e.g. sending funds).
	NewPassphrase string `json:"new_passphrase"`

	// The current passphrase.
	OldPassphrase string `json:"old_passphrase"`
}

// PostTransactionFeeJSONBody defines parameters for PostTransactionFee.
type PostTransactionFeeJSONBody interface{}

// SignMetadataJSONBody defines parameters for SignMetadata.
type SignMetadataJSONBody struct {

	// **⚠️ WARNING ⚠️**
	//
	// _Please note that metadata provided in a transaction will be
	// stored on the blockchain forever. Make sure not to include any sensitive data,
	// in particular personally identifiable information (PII)._
	//
	// Extra application data attached to the transaction.
	//
	// Cardano allows users and developers to embed their own
	// authenticated metadata when submitting transactions. Metadata can
	// be expressed as a JSON object with some restrictions:
	//
	// 1. All top-level keys must be integers between `0` and `2^64 - 1`.
	//
	// 2. Each metadata value is tagged with its type.
	//
	// 3. Strings must be at most 64 bytes when UTF-8 encoded.
	//
	// 4. Bytestrings are hex-encoded, with a maximum length of 64 bytes.
	//
	// Metadata aren't stored as JSON on the Cardano blockchain but are
	// instead stored using a compact binary encoding (CBOR).
	//
	// The binary encoding of metadata values supports three simple types:
	//
	// * Integers in the range `-(2^64 - 1)` to `2^64 - 1`
	// * Strings (UTF-8 encoded)
	// * Bytestrings
	//
	// And two compound types:
	//
	// * Lists of metadata values
	// * Mappings from metadata values to metadata values
	//
	// It is possible to transform any JSON object into this schema.
	//
	// However, if your application uses floating point values, they will
	// need to be converted somehow, according to your
	// requirements. Likewise for `null` or `bool` values. When reading
	// metadata from chain, be aware that integers may exceed the
	// javascript numeric range, and may need special "bigint" parsing.
	Metadata *SignMetadataJSONBody_Metadata `json:"metadata"`

	// A master passphrase to lock and protect the wallet for sensitive operation (e.g. sending funds)
	Passphrase string `json:"passphrase"`
}

// SignMetadataJSONBody_Metadata defines parameters for SignMetadata.
type SignMetadataJSONBody_Metadata struct {
	AdditionalProperties map[string]TransactionMetadataValue `json:"-"`
}

// ListTransactionsParams defines parameters for ListTransactions.
type ListTransactionsParams struct {

	// An optional start time in ISO 8601 date-and-time format. Basic and
	// extended formats are both accepted. Times can be local (with a
	// timezone offset) or UTC.
	//
	// If both a start time and an end time are specified, then the start
	// time must not be later than the end time.
	//
	// Example: `2008-08-08T08:08:08Z`
	Start *string `json:"start,omitempty"`

	// An optional end time in ISO 8601 date-and-time format. Basic and
	// extended formats are both accepted. Times can be local (with a
	// timezone offset) or UTC.
	//
	// If both a start time and an end time are specified, then the start
	// time must not be later than the end time.
	//
	// Example: `2008-08-08T08:08:08Z`
	End *string `json:"end,omitempty"`

	// An optional sort order.
	Order *string `json:"order,omitempty"`

	// Returns only transactions that have at least one withdrawal above the given amount.
	// This is particularly useful when set to `1` in order to list the withdrawal history of a wallet.
	MinWithdrawal *int `json:"minWithdrawal,omitempty"`
}

// PostTransactionJSONBody defines parameters for PostTransaction.
type PostTransactionJSONBody interface{}

// PostAnyAddressJSONRequestBody defines body for PostAnyAddress for application/json ContentType.
type PostAnyAddressJSONRequestBody PostAnyAddressJSONBody

// PostByronWalletJSONRequestBody defines body for PostByronWallet for application/json ContentType.
type PostByronWalletJSONRequestBody PostByronWalletJSONBody

// PutByronWalletJSONRequestBody defines body for PutByronWallet for application/json ContentType.
type PutByronWalletJSONRequestBody PutByronWalletJSONBody

// CreateAddressJSONRequestBody defines body for CreateAddress for application/json ContentType.
type CreateAddressJSONRequestBody CreateAddressJSONBody

// ImportAddressesJSONRequestBody defines body for ImportAddresses for application/json ContentType.
type ImportAddressesJSONRequestBody ImportAddressesJSONBody

// ByronSelectCoinsJSONRequestBody defines body for ByronSelectCoins for application/json ContentType.
type ByronSelectCoinsJSONRequestBody ByronSelectCoinsJSONBody

// MigrateByronWalletJSONRequestBody defines body for MigrateByronWallet for application/json ContentType.
type MigrateByronWalletJSONRequestBody MigrateByronWalletJSONBody

// PutByronWalletPassphraseJSONRequestBody defines body for PutByronWalletPassphrase for application/json ContentType.
type PutByronWalletPassphraseJSONRequestBody PutByronWalletPassphraseJSONBody

// PostByronTransactionFeeJSONRequestBody defines body for PostByronTransactionFee for application/json ContentType.
type PostByronTransactionFeeJSONRequestBody PostByronTransactionFeeJSONBody

// PostByronTransactionJSONRequestBody defines body for PostByronTransaction for application/json ContentType.
type PostByronTransactionJSONRequestBody PostByronTransactionJSONBody

// PutSettingsJSONRequestBody defines body for PutSettings for application/json ContentType.
type PutSettingsJSONRequestBody PutSettingsJSONBody

// PostSharedWalletJSONRequestBody defines body for PostSharedWallet for application/json ContentType.
type PostSharedWalletJSONRequestBody PostSharedWalletJSONBody

// PatchSharedWalletInDelegationJSONRequestBody defines body for PatchSharedWalletInDelegation for application/json ContentType.
type PatchSharedWalletInDelegationJSONRequestBody PatchSharedWalletInDelegationJSONBody

// PatchSharedWalletInPaymentJSONRequestBody defines body for PatchSharedWalletInPayment for application/json ContentType.
type PatchSharedWalletInPaymentJSONRequestBody PatchSharedWalletInPaymentJSONBody

// QuitStakePoolJSONRequestBody defines body for QuitStakePool for application/json ContentType.
type QuitStakePoolJSONRequestBody QuitStakePoolJSONBody

// PostMaintenanceActionJSONRequestBody defines body for PostMaintenanceAction for application/json ContentType.
type PostMaintenanceActionJSONRequestBody PostMaintenanceActionJSONBody

// JoinStakePoolJSONRequestBody defines body for JoinStakePool for application/json ContentType.
type JoinStakePoolJSONRequestBody JoinStakePoolJSONBody

// PostWalletJSONRequestBody defines body for PostWallet for application/json ContentType.
type PostWalletJSONRequestBody PostWalletJSONBody

// PutWalletJSONRequestBody defines body for PutWallet for application/json ContentType.
type PutWalletJSONRequestBody PutWalletJSONBody

// SelectCoinsJSONRequestBody defines body for SelectCoins for application/json ContentType.
type SelectCoinsJSONRequestBody SelectCoinsJSONBody

// PostAccountKeyJSONRequestBody defines body for PostAccountKey for application/json ContentType.
type PostAccountKeyJSONRequestBody PostAccountKeyJSONBody

// MigrateShelleyWalletJSONRequestBody defines body for MigrateShelleyWallet for application/json ContentType.
type MigrateShelleyWalletJSONRequestBody MigrateShelleyWalletJSONBody

// PutWalletPassphraseJSONRequestBody defines body for PutWalletPassphrase for application/json ContentType.
type PutWalletPassphraseJSONRequestBody PutWalletPassphraseJSONBody

// PostTransactionFeeJSONRequestBody defines body for PostTransactionFee for application/json ContentType.
type PostTransactionFeeJSONRequestBody PostTransactionFeeJSONBody

// SignMetadataJSONRequestBody defines body for SignMetadata for application/json ContentType.
type SignMetadataJSONRequestBody SignMetadataJSONBody

// PostTransactionJSONRequestBody defines body for PostTransaction for application/json ContentType.
type PostTransactionJSONRequestBody PostTransactionJSONBody

// Getter for additional properties for PatchSharedWalletInDelegationJSONBody. Returns the specified
// element and whether it was found
func (a PatchSharedWalletInDelegationJSONBody) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for PatchSharedWalletInDelegationJSONBody
func (a *PatchSharedWalletInDelegationJSONBody) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for PatchSharedWalletInDelegationJSONBody to handle AdditionalProperties
func (a *PatchSharedWalletInDelegationJSONBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for PatchSharedWalletInDelegationJSONBody to handle AdditionalProperties
func (a PatchSharedWalletInDelegationJSONBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for PatchSharedWalletInPaymentJSONBody. Returns the specified
// element and whether it was found
func (a PatchSharedWalletInPaymentJSONBody) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for PatchSharedWalletInPaymentJSONBody
func (a *PatchSharedWalletInPaymentJSONBody) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for PatchSharedWalletInPaymentJSONBody to handle AdditionalProperties
func (a *PatchSharedWalletInPaymentJSONBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for PatchSharedWalletInPaymentJSONBody to handle AdditionalProperties
func (a PatchSharedWalletInPaymentJSONBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for SignMetadataJSONBody_Metadata. Returns the specified
// element and whether it was found
func (a SignMetadataJSONBody_Metadata) Get(fieldName string) (value TransactionMetadataValue, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for SignMetadataJSONBody_Metadata
func (a *SignMetadataJSONBody_Metadata) Set(fieldName string, value TransactionMetadataValue) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]TransactionMetadataValue)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for SignMetadataJSONBody_Metadata to handle AdditionalProperties
func (a *SignMetadataJSONBody_Metadata) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]TransactionMetadataValue)
		for fieldName, fieldBuf := range object {
			var fieldVal TransactionMetadataValue
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for SignMetadataJSONBody_Metadata to handle AdditionalProperties
func (a SignMetadataJSONBody_Metadata) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
