// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firestore

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Cloud Firestore Database.
//
// If you wish to use Firestore with App Engine, use the
// `appengine.Application`
// resource instead. If you were previously using the `appengine.Application` resource exclusively for managing a Firestore database
// and would like to use the `firestore.Database` resource instead, please follow the instructions
// [here](https://cloud.google.com/firestore/docs/app-engine-requirement).
//
// To get more information about Database, see:
//
// * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/firestore/docs/)
//
// ## Example Usage
// ### Firestore Default Database
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firestore.NewDatabase(ctx, "database", &firestore.DatabaseArgs{
//				LocationId: pulumi.String("nam5"),
//				Project:    pulumi.String("my-project-name"),
//				Type:       pulumi.String("FIRESTORE_NATIVE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firestore Database
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firestore.NewDatabase(ctx, "database", &firestore.DatabaseArgs{
//				AppEngineIntegrationMode:      pulumi.String("DISABLED"),
//				ConcurrencyMode:               pulumi.String("OPTIMISTIC"),
//				DeleteProtectionState:         pulumi.String("DELETE_PROTECTION_ENABLED"),
//				DeletionPolicy:                pulumi.String("DELETE"),
//				LocationId:                    pulumi.String("nam5"),
//				PointInTimeRecoveryEnablement: pulumi.String("POINT_IN_TIME_RECOVERY_ENABLED"),
//				Project:                       pulumi.String("my-project-name"),
//				Type:                          pulumi.String("FIRESTORE_NATIVE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firestore Default Database In Datastore Mode
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firestore.NewDatabase(ctx, "datastoreModeDatabase", &firestore.DatabaseArgs{
//				LocationId: pulumi.String("nam5"),
//				Project:    pulumi.String("my-project-name"),
//				Type:       pulumi.String("DATASTORE_MODE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firestore Database In Datastore Mode
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firestore.NewDatabase(ctx, "datastoreModeDatabase", &firestore.DatabaseArgs{
//				AppEngineIntegrationMode:      pulumi.String("DISABLED"),
//				ConcurrencyMode:               pulumi.String("OPTIMISTIC"),
//				DeleteProtectionState:         pulumi.String("DELETE_PROTECTION_ENABLED"),
//				DeletionPolicy:                pulumi.String("DELETE"),
//				LocationId:                    pulumi.String("nam5"),
//				PointInTimeRecoveryEnablement: pulumi.String("POINT_IN_TIME_RECOVERY_ENABLED"),
//				Project:                       pulumi.String("my-project-name"),
//				Type:                          pulumi.String("DATASTORE_MODE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Database can be imported using any of these accepted formats* `projects/{{project}}/databases/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Database can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:firestore/database:Database default projects/{{project}}/databases/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firestore/database:Database default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firestore/database:Database default {{name}}
//
// ```
type Database struct {
	pulumi.CustomResourceState

	// The App Engine integration mode to use for this database.
	// Possible values are: `ENABLED`, `DISABLED`.
	AppEngineIntegrationMode pulumi.StringOutput `pulumi:"appEngineIntegrationMode"`
	// The concurrency control mode to use for this database.
	// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
	ConcurrencyMode pulumi.StringOutput `pulumi:"concurrencyMode"`
	// Output only. The timestamp at which this database was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
	// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
	// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
	// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
	DeleteProtectionState pulumi.StringOutput `pulumi:"deleteProtectionState"`
	// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
	// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
	// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
	// 'delete_protection'.
	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.
	// This value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	EarliestVersionTime pulumi.StringOutput `pulumi:"earliestVersionTime"`
	// Output only. This checksum is computed by the server based on the value of other fields,
	// and may be sent on update and delete requests to ensure the client has an
	// up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Output only. The keyPrefix for this database.
	// This keyPrefix is used, in combination with the project id ("~") to construct the application id
	// that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
	// This value may be empty in which case the appid to use for URL-encoded keys is the projectId (eg: foo instead of v~foo).
	KeyPrefix pulumi.StringOutput `pulumi:"keyPrefix"`
	// The location of the database. Available locations are listed at
	// https://cloud.google.com/firestore/docs/locations.
	LocationId pulumi.StringOutput `pulumi:"locationId"`
	// The ID to use for the database, which will become the final
	// component of the database's resource name. This value should be 4-63
	// characters. Valid characters are /[a-z][0-9]-/ with first character
	// a letter and the last a letter or a number. Must not be
	// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	// "(default)" database id is also valid.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether to enable the PITR feature on this database.
	// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
	// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
	// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
	// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
	// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
	// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
	PointInTimeRecoveryEnablement pulumi.StringPtrOutput `pulumi:"pointInTimeRecoveryEnablement"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The type of the database.
	// See https://cloud.google.com/datastore/docs/firestore-or-datastore
	// for information about how to choose.
	// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
	//
	// ***
	Type pulumi.StringOutput `pulumi:"type"`
	// Output only. The system-generated UUID4 for this Database.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The timestamp at which this database was most recently updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Output only. The period during which past versions of data are retained in the database.
	// Any read or query can specify a readTime within this window, and will read the state of the database at that time.
	// If the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	VersionRetentionPeriod pulumi.StringOutput `pulumi:"versionRetentionPeriod"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocationId == nil {
		return nil, errors.New("invalid value for required argument 'LocationId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("gcp:firestore/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("gcp:firestore/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// The App Engine integration mode to use for this database.
	// Possible values are: `ENABLED`, `DISABLED`.
	AppEngineIntegrationMode *string `pulumi:"appEngineIntegrationMode"`
	// The concurrency control mode to use for this database.
	// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
	ConcurrencyMode *string `pulumi:"concurrencyMode"`
	// Output only. The timestamp at which this database was created.
	CreateTime *string `pulumi:"createTime"`
	// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
	// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
	// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
	// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
	DeleteProtectionState *string `pulumi:"deleteProtectionState"`
	// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
	// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
	// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
	// 'delete_protection'.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.
	// This value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	EarliestVersionTime *string `pulumi:"earliestVersionTime"`
	// Output only. This checksum is computed by the server based on the value of other fields,
	// and may be sent on update and delete requests to ensure the client has an
	// up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Output only. The keyPrefix for this database.
	// This keyPrefix is used, in combination with the project id ("~") to construct the application id
	// that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
	// This value may be empty in which case the appid to use for URL-encoded keys is the projectId (eg: foo instead of v~foo).
	KeyPrefix *string `pulumi:"keyPrefix"`
	// The location of the database. Available locations are listed at
	// https://cloud.google.com/firestore/docs/locations.
	LocationId *string `pulumi:"locationId"`
	// The ID to use for the database, which will become the final
	// component of the database's resource name. This value should be 4-63
	// characters. Valid characters are /[a-z][0-9]-/ with first character
	// a letter and the last a letter or a number. Must not be
	// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	// "(default)" database id is also valid.
	Name *string `pulumi:"name"`
	// Whether to enable the PITR feature on this database.
	// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
	// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
	// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
	// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
	// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
	// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
	PointInTimeRecoveryEnablement *string `pulumi:"pointInTimeRecoveryEnablement"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The type of the database.
	// See https://cloud.google.com/datastore/docs/firestore-or-datastore
	// for information about how to choose.
	// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
	//
	// ***
	Type *string `pulumi:"type"`
	// Output only. The system-generated UUID4 for this Database.
	Uid *string `pulumi:"uid"`
	// Output only. The timestamp at which this database was most recently updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Output only. The period during which past versions of data are retained in the database.
	// Any read or query can specify a readTime within this window, and will read the state of the database at that time.
	// If the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	VersionRetentionPeriod *string `pulumi:"versionRetentionPeriod"`
}

type DatabaseState struct {
	// The App Engine integration mode to use for this database.
	// Possible values are: `ENABLED`, `DISABLED`.
	AppEngineIntegrationMode pulumi.StringPtrInput
	// The concurrency control mode to use for this database.
	// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
	ConcurrencyMode pulumi.StringPtrInput
	// Output only. The timestamp at which this database was created.
	CreateTime pulumi.StringPtrInput
	// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
	// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
	// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
	// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
	DeleteProtectionState pulumi.StringPtrInput
	// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
	// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
	// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
	// 'delete_protection'.
	DeletionPolicy pulumi.StringPtrInput
	// Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.
	// This value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	EarliestVersionTime pulumi.StringPtrInput
	// Output only. This checksum is computed by the server based on the value of other fields,
	// and may be sent on update and delete requests to ensure the client has an
	// up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Output only. The keyPrefix for this database.
	// This keyPrefix is used, in combination with the project id ("~") to construct the application id
	// that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
	// This value may be empty in which case the appid to use for URL-encoded keys is the projectId (eg: foo instead of v~foo).
	KeyPrefix pulumi.StringPtrInput
	// The location of the database. Available locations are listed at
	// https://cloud.google.com/firestore/docs/locations.
	LocationId pulumi.StringPtrInput
	// The ID to use for the database, which will become the final
	// component of the database's resource name. This value should be 4-63
	// characters. Valid characters are /[a-z][0-9]-/ with first character
	// a letter and the last a letter or a number. Must not be
	// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	// "(default)" database id is also valid.
	Name pulumi.StringPtrInput
	// Whether to enable the PITR feature on this database.
	// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
	// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
	// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
	// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
	// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
	// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
	PointInTimeRecoveryEnablement pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The type of the database.
	// See https://cloud.google.com/datastore/docs/firestore-or-datastore
	// for information about how to choose.
	// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
	//
	// ***
	Type pulumi.StringPtrInput
	// Output only. The system-generated UUID4 for this Database.
	Uid pulumi.StringPtrInput
	// Output only. The timestamp at which this database was most recently updated.
	UpdateTime pulumi.StringPtrInput
	// Output only. The period during which past versions of data are retained in the database.
	// Any read or query can specify a readTime within this window, and will read the state of the database at that time.
	// If the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	VersionRetentionPeriod pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The App Engine integration mode to use for this database.
	// Possible values are: `ENABLED`, `DISABLED`.
	AppEngineIntegrationMode *string `pulumi:"appEngineIntegrationMode"`
	// The concurrency control mode to use for this database.
	// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
	ConcurrencyMode *string `pulumi:"concurrencyMode"`
	// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
	// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
	// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
	// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
	DeleteProtectionState *string `pulumi:"deleteProtectionState"`
	// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
	// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
	// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
	// 'delete_protection'.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The location of the database. Available locations are listed at
	// https://cloud.google.com/firestore/docs/locations.
	LocationId string `pulumi:"locationId"`
	// The ID to use for the database, which will become the final
	// component of the database's resource name. This value should be 4-63
	// characters. Valid characters are /[a-z][0-9]-/ with first character
	// a letter and the last a letter or a number. Must not be
	// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	// "(default)" database id is also valid.
	Name *string `pulumi:"name"`
	// Whether to enable the PITR feature on this database.
	// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
	// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
	// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
	// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
	// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
	// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
	PointInTimeRecoveryEnablement *string `pulumi:"pointInTimeRecoveryEnablement"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The type of the database.
	// See https://cloud.google.com/datastore/docs/firestore-or-datastore
	// for information about how to choose.
	// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
	//
	// ***
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The App Engine integration mode to use for this database.
	// Possible values are: `ENABLED`, `DISABLED`.
	AppEngineIntegrationMode pulumi.StringPtrInput
	// The concurrency control mode to use for this database.
	// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
	ConcurrencyMode pulumi.StringPtrInput
	// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
	// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
	// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
	// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
	DeleteProtectionState pulumi.StringPtrInput
	// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
	// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
	// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
	// 'delete_protection'.
	DeletionPolicy pulumi.StringPtrInput
	// The location of the database. Available locations are listed at
	// https://cloud.google.com/firestore/docs/locations.
	LocationId pulumi.StringInput
	// The ID to use for the database, which will become the final
	// component of the database's resource name. This value should be 4-63
	// characters. Valid characters are /[a-z][0-9]-/ with first character
	// a letter and the last a letter or a number. Must not be
	// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	// "(default)" database id is also valid.
	Name pulumi.StringPtrInput
	// Whether to enable the PITR feature on this database.
	// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
	// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
	// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
	// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
	// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
	// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
	PointInTimeRecoveryEnablement pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The type of the database.
	// See https://cloud.google.com/datastore/docs/firestore-or-datastore
	// for information about how to choose.
	// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
	//
	// ***
	Type pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// The App Engine integration mode to use for this database.
// Possible values are: `ENABLED`, `DISABLED`.
func (o DatabaseOutput) AppEngineIntegrationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.AppEngineIntegrationMode }).(pulumi.StringOutput)
}

// The concurrency control mode to use for this database.
// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
func (o DatabaseOutput) ConcurrencyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ConcurrencyMode }).(pulumi.StringOutput)
}

// Output only. The timestamp at which this database was created.
func (o DatabaseOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
func (o DatabaseOutput) DeleteProtectionState() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DeleteProtectionState }).(pulumi.StringOutput)
}

// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
// 'delete_protection'.
func (o DatabaseOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

// Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.
// This value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o DatabaseOutput) EarliestVersionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.EarliestVersionTime }).(pulumi.StringOutput)
}

// Output only. This checksum is computed by the server based on the value of other fields,
// and may be sent on update and delete requests to ensure the client has an
// up-to-date value before proceeding.
func (o DatabaseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Output only. The keyPrefix for this database.
// This keyPrefix is used, in combination with the project id ("~") to construct the application id
// that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
// This value may be empty in which case the appid to use for URL-encoded keys is the projectId (eg: foo instead of v~foo).
func (o DatabaseOutput) KeyPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.KeyPrefix }).(pulumi.StringOutput)
}

// The location of the database. Available locations are listed at
// https://cloud.google.com/firestore/docs/locations.
func (o DatabaseOutput) LocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.LocationId }).(pulumi.StringOutput)
}

// The ID to use for the database, which will become the final
// component of the database's resource name. This value should be 4-63
// characters. Valid characters are /[a-z][0-9]-/ with first character
// a letter and the last a letter or a number. Must not be
// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
// "(default)" database id is also valid.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether to enable the PITR feature on this database.
// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
func (o DatabaseOutput) PointInTimeRecoveryEnablement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.PointInTimeRecoveryEnablement }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DatabaseOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The type of the database.
// See https://cloud.google.com/datastore/docs/firestore-or-datastore
// for information about how to choose.
// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
//
// ***
func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Output only. The system-generated UUID4 for this Database.
func (o DatabaseOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The timestamp at which this database was most recently updated.
func (o DatabaseOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Output only. The period during which past versions of data are retained in the database.
// Any read or query can specify a readTime within this window, and will read the state of the database at that time.
// If the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.
// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
func (o DatabaseOutput) VersionRetentionPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.VersionRetentionPeriod }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
