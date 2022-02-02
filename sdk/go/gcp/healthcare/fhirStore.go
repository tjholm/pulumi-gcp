// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/fhir/STU3/)
// standard for Healthcare information exchange
//
// To get more information about FhirStore, see:
//
// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.fhirStores)
// * How-to Guides
//     * [Creating a FHIR store](https://cloud.google.com/healthcare/docs/how-tos/fhir)
//
// ## Example Usage
// ### Healthcare Fhir Store Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		topic, err := pubsub.NewTopic(ctx, "topic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		dataset, err := healthcare.NewDataset(ctx, "dataset", &healthcare.DatasetArgs{
// 			Location: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewFhirStore(ctx, "default", &healthcare.FhirStoreArgs{
// 			Dataset:                     dataset.ID(),
// 			Version:                     pulumi.String("R4"),
// 			EnableUpdateCreate:          pulumi.Bool(false),
// 			DisableReferentialIntegrity: pulumi.Bool(false),
// 			DisableResourceVersioning:   pulumi.Bool(false),
// 			EnableHistoryImport:         pulumi.Bool(false),
// 			NotificationConfig: &healthcare.FhirStoreNotificationConfigArgs{
// 				PubsubTopic: topic.ID(),
// 			},
// 			Labels: pulumi.StringMap{
// 				"label1": pulumi.String("labelvalue1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Healthcare Fhir Store Streaming Config
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dataset, err := healthcare.NewDataset(ctx, "dataset", &healthcare.DatasetArgs{
// 			Location: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bqDataset, err := bigquery.NewDataset(ctx, "bqDataset", &bigquery.DatasetArgs{
// 			DatasetId:               pulumi.String("bq_example_dataset"),
// 			FriendlyName:            pulumi.String("test"),
// 			Description:             pulumi.String("This is a test description"),
// 			Location:                pulumi.String("US"),
// 			DeleteContentsOnDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewFhirStore(ctx, "default", &healthcare.FhirStoreArgs{
// 			Dataset:                     dataset.ID(),
// 			Version:                     pulumi.String("R4"),
// 			EnableUpdateCreate:          pulumi.Bool(false),
// 			DisableReferentialIntegrity: pulumi.Bool(false),
// 			DisableResourceVersioning:   pulumi.Bool(false),
// 			EnableHistoryImport:         pulumi.Bool(false),
// 			Labels: pulumi.StringMap{
// 				"label1": pulumi.String("labelvalue1"),
// 			},
// 			StreamConfigs: healthcare.FhirStoreStreamConfigArray{
// 				&healthcare.FhirStoreStreamConfigArgs{
// 					ResourceTypes: pulumi.StringArray{
// 						pulumi.String("Observation"),
// 					},
// 					BigqueryDestination: &healthcare.FhirStoreStreamConfigBigqueryDestinationArgs{
// 						DatasetUri: pulumi.All(bqDataset.Project, bqDataset.DatasetId).ApplyT(func(_args []interface{}) (string, error) {
// 							project := _args[0].(string)
// 							datasetId := _args[1].(string)
// 							return fmt.Sprintf("%v%v%v%v", "bq://", project, ".", datasetId), nil
// 						}).(pulumi.StringOutput),
// 						SchemaConfig: &healthcare.FhirStoreStreamConfigBigqueryDestinationSchemaConfigArgs{
// 							RecursiveStructureDepth: pulumi.Int(3),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewTopic(ctx, "topic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// FhirStore can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/fhirStores/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/{{name}}
// ```
type FhirStore struct {
	pulumi.CustomResourceState

	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
	// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
	// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
	// will skip referential integrity check. Consequently, operations that rely on references, such as
	// Patient.get$everything, will not return all the results if broken references exist.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableReferentialIntegrity pulumi.BoolPtrOutput `pulumi:"disableReferentialIntegrity"`
	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
	// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
	// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
	// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
	// attempts to read the historical versions.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableResourceVersioning pulumi.BoolPtrOutput `pulumi:"disableResourceVersioning"`
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
	// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
	// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
	// will fail with an error.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
	EnableHistoryImport pulumi.BoolPtrOutput `pulumi:"enableHistoryImport"`
	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
	// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
	// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
	// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
	// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
	// notifications.
	EnableUpdateCreate pulumi.BoolPtrOutput `pulumi:"enableUpdateCreate"`
	// User-supplied key-value pairs used to organize FHIR stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for the FhirStore.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	Name pulumi.StringOutput `pulumi:"name"`
	// A nested object resource
	// Structure is documented below.
	NotificationConfig FhirStoreNotificationConfigPtrOutput `pulumi:"notificationConfig"`
	// The fully qualified name of this dataset
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
	// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
	// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
	// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
	// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
	// the order of dozens of seconds) is expected before the results show up in the streaming destination.
	// Structure is documented below.
	StreamConfigs FhirStoreStreamConfigArrayOutput `pulumi:"streamConfigs"`
	// The FHIR specification version.
	// Default value is `STU3`.
	// Possible values are `DSTU2`, `STU3`, and `R4`.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewFhirStore registers a new resource with the given unique name, arguments, and options.
func NewFhirStore(ctx *pulumi.Context,
	name string, args *FhirStoreArgs, opts ...pulumi.ResourceOption) (*FhirStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	var resource FhirStore
	err := ctx.RegisterResource("gcp:healthcare/fhirStore:FhirStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirStore gets an existing FhirStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirStoreState, opts ...pulumi.ResourceOption) (*FhirStore, error) {
	var resource FhirStore
	err := ctx.ReadResource("gcp:healthcare/fhirStore:FhirStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirStore resources.
type fhirStoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset *string `pulumi:"dataset"`
	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
	// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
	// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
	// will skip referential integrity check. Consequently, operations that rely on references, such as
	// Patient.get$everything, will not return all the results if broken references exist.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableReferentialIntegrity *bool `pulumi:"disableReferentialIntegrity"`
	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
	// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
	// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
	// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
	// attempts to read the historical versions.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableResourceVersioning *bool `pulumi:"disableResourceVersioning"`
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
	// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
	// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
	// will fail with an error.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
	EnableHistoryImport *bool `pulumi:"enableHistoryImport"`
	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
	// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
	// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
	// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
	// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
	// notifications.
	EnableUpdateCreate *bool `pulumi:"enableUpdateCreate"`
	// User-supplied key-value pairs used to organize FHIR stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the FhirStore.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	Name *string `pulumi:"name"`
	// A nested object resource
	// Structure is documented below.
	NotificationConfig *FhirStoreNotificationConfig `pulumi:"notificationConfig"`
	// The fully qualified name of this dataset
	SelfLink *string `pulumi:"selfLink"`
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
	// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
	// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
	// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
	// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
	// the order of dozens of seconds) is expected before the results show up in the streaming destination.
	// Structure is documented below.
	StreamConfigs []FhirStoreStreamConfig `pulumi:"streamConfigs"`
	// The FHIR specification version.
	// Default value is `STU3`.
	// Possible values are `DSTU2`, `STU3`, and `R4`.
	Version *string `pulumi:"version"`
}

type FhirStoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringPtrInput
	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
	// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
	// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
	// will skip referential integrity check. Consequently, operations that rely on references, such as
	// Patient.get$everything, will not return all the results if broken references exist.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableReferentialIntegrity pulumi.BoolPtrInput
	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
	// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
	// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
	// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
	// attempts to read the historical versions.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableResourceVersioning pulumi.BoolPtrInput
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
	// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
	// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
	// will fail with an error.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
	EnableHistoryImport pulumi.BoolPtrInput
	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
	// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
	// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
	// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
	// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
	// notifications.
	EnableUpdateCreate pulumi.BoolPtrInput
	// User-supplied key-value pairs used to organize FHIR stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the FhirStore.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	Name pulumi.StringPtrInput
	// A nested object resource
	// Structure is documented below.
	NotificationConfig FhirStoreNotificationConfigPtrInput
	// The fully qualified name of this dataset
	SelfLink pulumi.StringPtrInput
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
	// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
	// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
	// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
	// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
	// the order of dozens of seconds) is expected before the results show up in the streaming destination.
	// Structure is documented below.
	StreamConfigs FhirStoreStreamConfigArrayInput
	// The FHIR specification version.
	// Default value is `STU3`.
	// Possible values are `DSTU2`, `STU3`, and `R4`.
	Version pulumi.StringPtrInput
}

func (FhirStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreState)(nil)).Elem()
}

type fhirStoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset string `pulumi:"dataset"`
	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
	// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
	// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
	// will skip referential integrity check. Consequently, operations that rely on references, such as
	// Patient.get$everything, will not return all the results if broken references exist.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableReferentialIntegrity *bool `pulumi:"disableReferentialIntegrity"`
	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
	// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
	// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
	// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
	// attempts to read the historical versions.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableResourceVersioning *bool `pulumi:"disableResourceVersioning"`
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
	// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
	// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
	// will fail with an error.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
	EnableHistoryImport *bool `pulumi:"enableHistoryImport"`
	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
	// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
	// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
	// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
	// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
	// notifications.
	EnableUpdateCreate *bool `pulumi:"enableUpdateCreate"`
	// User-supplied key-value pairs used to organize FHIR stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the FhirStore.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	Name *string `pulumi:"name"`
	// A nested object resource
	// Structure is documented below.
	NotificationConfig *FhirStoreNotificationConfig `pulumi:"notificationConfig"`
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
	// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
	// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
	// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
	// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
	// the order of dozens of seconds) is expected before the results show up in the streaming destination.
	// Structure is documented below.
	StreamConfigs []FhirStoreStreamConfig `pulumi:"streamConfigs"`
	// The FHIR specification version.
	// Default value is `STU3`.
	// Possible values are `DSTU2`, `STU3`, and `R4`.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a FhirStore resource.
type FhirStoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringInput
	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
	// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
	// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
	// will skip referential integrity check. Consequently, operations that rely on references, such as
	// Patient.get$everything, will not return all the results if broken references exist.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableReferentialIntegrity pulumi.BoolPtrInput
	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
	// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
	// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
	// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
	// attempts to read the historical versions.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	DisableResourceVersioning pulumi.BoolPtrInput
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
	// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
	// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
	// will fail with an error.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
	EnableHistoryImport pulumi.BoolPtrInput
	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
	// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
	// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
	// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
	// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
	// notifications.
	EnableUpdateCreate pulumi.BoolPtrInput
	// User-supplied key-value pairs used to organize FHIR stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the FhirStore.
	// ** Changing this property may recreate the FHIR store (removing all data) **
	Name pulumi.StringPtrInput
	// A nested object resource
	// Structure is documented below.
	NotificationConfig FhirStoreNotificationConfigPtrInput
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
	// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
	// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
	// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
	// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
	// the order of dozens of seconds) is expected before the results show up in the streaming destination.
	// Structure is documented below.
	StreamConfigs FhirStoreStreamConfigArrayInput
	// The FHIR specification version.
	// Default value is `STU3`.
	// Possible values are `DSTU2`, `STU3`, and `R4`.
	Version pulumi.StringPtrInput
}

func (FhirStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreArgs)(nil)).Elem()
}

type FhirStoreInput interface {
	pulumi.Input

	ToFhirStoreOutput() FhirStoreOutput
	ToFhirStoreOutputWithContext(ctx context.Context) FhirStoreOutput
}

func (*FhirStore) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStore)(nil)).Elem()
}

func (i *FhirStore) ToFhirStoreOutput() FhirStoreOutput {
	return i.ToFhirStoreOutputWithContext(context.Background())
}

func (i *FhirStore) ToFhirStoreOutputWithContext(ctx context.Context) FhirStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreOutput)
}

// FhirStoreArrayInput is an input type that accepts FhirStoreArray and FhirStoreArrayOutput values.
// You can construct a concrete instance of `FhirStoreArrayInput` via:
//
//          FhirStoreArray{ FhirStoreArgs{...} }
type FhirStoreArrayInput interface {
	pulumi.Input

	ToFhirStoreArrayOutput() FhirStoreArrayOutput
	ToFhirStoreArrayOutputWithContext(context.Context) FhirStoreArrayOutput
}

type FhirStoreArray []FhirStoreInput

func (FhirStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FhirStore)(nil)).Elem()
}

func (i FhirStoreArray) ToFhirStoreArrayOutput() FhirStoreArrayOutput {
	return i.ToFhirStoreArrayOutputWithContext(context.Background())
}

func (i FhirStoreArray) ToFhirStoreArrayOutputWithContext(ctx context.Context) FhirStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreArrayOutput)
}

// FhirStoreMapInput is an input type that accepts FhirStoreMap and FhirStoreMapOutput values.
// You can construct a concrete instance of `FhirStoreMapInput` via:
//
//          FhirStoreMap{ "key": FhirStoreArgs{...} }
type FhirStoreMapInput interface {
	pulumi.Input

	ToFhirStoreMapOutput() FhirStoreMapOutput
	ToFhirStoreMapOutputWithContext(context.Context) FhirStoreMapOutput
}

type FhirStoreMap map[string]FhirStoreInput

func (FhirStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FhirStore)(nil)).Elem()
}

func (i FhirStoreMap) ToFhirStoreMapOutput() FhirStoreMapOutput {
	return i.ToFhirStoreMapOutputWithContext(context.Background())
}

func (i FhirStoreMap) ToFhirStoreMapOutputWithContext(ctx context.Context) FhirStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreMapOutput)
}

type FhirStoreOutput struct{ *pulumi.OutputState }

func (FhirStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStore)(nil)).Elem()
}

func (o FhirStoreOutput) ToFhirStoreOutput() FhirStoreOutput {
	return o
}

func (o FhirStoreOutput) ToFhirStoreOutputWithContext(ctx context.Context) FhirStoreOutput {
	return o
}

type FhirStoreArrayOutput struct{ *pulumi.OutputState }

func (FhirStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FhirStore)(nil)).Elem()
}

func (o FhirStoreArrayOutput) ToFhirStoreArrayOutput() FhirStoreArrayOutput {
	return o
}

func (o FhirStoreArrayOutput) ToFhirStoreArrayOutputWithContext(ctx context.Context) FhirStoreArrayOutput {
	return o
}

func (o FhirStoreArrayOutput) Index(i pulumi.IntInput) FhirStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FhirStore {
		return vs[0].([]*FhirStore)[vs[1].(int)]
	}).(FhirStoreOutput)
}

type FhirStoreMapOutput struct{ *pulumi.OutputState }

func (FhirStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FhirStore)(nil)).Elem()
}

func (o FhirStoreMapOutput) ToFhirStoreMapOutput() FhirStoreMapOutput {
	return o
}

func (o FhirStoreMapOutput) ToFhirStoreMapOutputWithContext(ctx context.Context) FhirStoreMapOutput {
	return o
}

func (o FhirStoreMapOutput) MapIndex(k pulumi.StringInput) FhirStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FhirStore {
		return vs[0].(map[string]*FhirStore)[vs[1].(string)]
	}).(FhirStoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreInput)(nil)).Elem(), &FhirStore{})
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreArrayInput)(nil)).Elem(), FhirStoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreMapInput)(nil)).Elem(), FhirStoreMap{})
	pulumi.RegisterOutputType(FhirStoreOutput{})
	pulumi.RegisterOutputType(FhirStoreArrayOutput{})
	pulumi.RegisterOutputType(FhirStoreMapOutput{})
}
