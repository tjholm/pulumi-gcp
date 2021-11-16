// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Cloud Function. For more information see:
//
// * [API documentation](https://cloud.google.com/functions/docs/reference/rest/v1/projects.locations.functions)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/functions/docs)
//
// > **Warning:** As of November 1, 2019, newly created Functions are
// private-by-default and will require [appropriate IAM permissions](https://cloud.google.com/functions/docs/reference/iam/roles)
// to be invoked. See below examples for how to set up the appropriate permissions,
// or view the [Cloud Functions IAM resources](https://www.terraform.io/docs/providers/google/r/cloudfunctions_cloud_function_iam.html)
// for Cloud Functions.
//
// ## Example Usage
// ### Public Function
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
// 			Location: pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		archive, err := storage.NewBucketObject(ctx, "archive", &storage.BucketObjectArgs{
// 			Bucket: bucket.Name,
// 			Source: pulumi.NewFileAsset("./path/to/zip/file/which/contains/code"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		function, err := cloudfunctions.NewFunction(ctx, "function", &cloudfunctions.FunctionArgs{
// 			Description:         pulumi.String("My function"),
// 			Runtime:             pulumi.String("nodejs14"),
// 			AvailableMemoryMb:   pulumi.Int(128),
// 			SourceArchiveBucket: bucket.Name,
// 			SourceArchiveObject: archive.Name,
// 			TriggerHttp:         pulumi.Bool(true),
// 			EntryPoint:          pulumi.String("helloGET"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudfunctions.NewFunctionIamMember(ctx, "invoker", &cloudfunctions.FunctionIamMemberArgs{
// 			Project:       function.Project,
// 			Region:        function.Region,
// 			CloudFunction: function.Name,
// 			Role:          pulumi.String("roles/cloudfunctions.invoker"),
// 			Member:        pulumi.String("allUsers"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Single User
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
// 			Location: pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		archive, err := storage.NewBucketObject(ctx, "archive", &storage.BucketObjectArgs{
// 			Bucket: bucket.Name,
// 			Source: pulumi.NewFileAsset("./path/to/zip/file/which/contains/code"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		function, err := cloudfunctions.NewFunction(ctx, "function", &cloudfunctions.FunctionArgs{
// 			Description:         pulumi.String("My function"),
// 			Runtime:             pulumi.String("nodejs14"),
// 			AvailableMemoryMb:   pulumi.Int(128),
// 			SourceArchiveBucket: bucket.Name,
// 			SourceArchiveObject: archive.Name,
// 			TriggerHttp:         pulumi.Bool(true),
// 			Timeout:             pulumi.Int(60),
// 			EntryPoint:          pulumi.String("helloGET"),
// 			Labels: pulumi.AnyMap{
// 				"my-label": pulumi.Any("my-label-value"),
// 			},
// 			EnvironmentVariables: pulumi.AnyMap{
// 				"MY_ENV_VAR": pulumi.Any("my-env-var-value"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudfunctions.NewFunctionIamMember(ctx, "invoker", &cloudfunctions.FunctionIamMemberArgs{
// 			Project:       function.Project,
// 			Region:        function.Region,
// 			CloudFunction: function.Name,
// 			Role:          pulumi.String("roles/cloudfunctions.invoker"),
// 			Member:        pulumi.String("user:myFunctionInvoker@example.com"),
// 		})
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
// Functions can be imported using the `name` or `{{project}}/{{region}}/name`, e.g.
//
// ```sh
//  $ pulumi import gcp:cloudfunctions/function:Function default function-test
// ```
//
// ```sh
//  $ pulumi import gcp:cloudfunctions/function:Function default {{project}}/{{region}}/function-test
// ```
type Function struct {
	pulumi.CustomResourceState

	// Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
	AvailableMemoryMb pulumi.IntPtrOutput `pulumi:"availableMemoryMb"`
	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables pulumi.MapOutput `pulumi:"buildEnvironmentVariables"`
	// Description of the function.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint pulumi.StringPtrOutput `pulumi:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables pulumi.MapOutput `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger FunctionEventTriggerOutput `pulumi:"eventTrigger"`
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl pulumi.StringOutput `pulumi:"httpsTriggerUrl"`
	// String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings pulumi.StringPtrOutput `pulumi:"ingressSettings"`
	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances pulumi.IntPtrOutput `pulumi:"maxInstances"`
	// A user-defined name of the function. Function names must be unique globally.
	Name pulumi.StringOutput `pulumi:"name"`
	// Project of the function. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region of function. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The runtime in which the function is going to run.
	// Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket pulumi.StringPtrOutput `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject pulumi.StringPtrOutput `pulumi:"sourceArchiveObject"`
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository FunctionSourceRepositoryPtrOutput `pulumi:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
	TriggerHttp pulumi.BoolPtrOutput `pulumi:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector pulumi.StringPtrOutput `pulumi:"vpcConnector"`
	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
	VpcConnectorEgressSettings pulumi.StringOutput `pulumi:"vpcConnectorEgressSettings"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Runtime == nil {
		return nil, errors.New("invalid value for required argument 'Runtime'")
	}
	var resource Function
	err := ctx.RegisterResource("gcp:cloudfunctions/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("gcp:cloudfunctions/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
	AvailableMemoryMb *int `pulumi:"availableMemoryMb"`
	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables map[string]interface{} `pulumi:"buildEnvironmentVariables"`
	// Description of the function.
	Description *string `pulumi:"description"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint *string `pulumi:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger *FunctionEventTrigger `pulumi:"eventTrigger"`
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl *string `pulumi:"httpsTriggerUrl"`
	// String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings *string `pulumi:"ingressSettings"`
	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]interface{} `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances *int `pulumi:"maxInstances"`
	// A user-defined name of the function. Function names must be unique globally.
	Name *string `pulumi:"name"`
	// Project of the function. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region of function. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The runtime in which the function is going to run.
	// Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
	Runtime *string `pulumi:"runtime"`
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket *string `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject *string `pulumi:"sourceArchiveObject"`
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository *FunctionSourceRepository `pulumi:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout *int `pulumi:"timeout"`
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
	TriggerHttp *bool `pulumi:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector *string `pulumi:"vpcConnector"`
	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
	VpcConnectorEgressSettings *string `pulumi:"vpcConnectorEgressSettings"`
}

type FunctionState struct {
	// Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
	AvailableMemoryMb pulumi.IntPtrInput
	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables pulumi.MapInput
	// Description of the function.
	Description pulumi.StringPtrInput
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint pulumi.StringPtrInput
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables pulumi.MapInput
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger FunctionEventTriggerPtrInput
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl pulumi.StringPtrInput
	// String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.MapInput
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances pulumi.IntPtrInput
	// A user-defined name of the function. Function names must be unique globally.
	Name pulumi.StringPtrInput
	// Project of the function. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region of function. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The runtime in which the function is going to run.
	// Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
	Runtime pulumi.StringPtrInput
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail pulumi.StringPtrInput
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket pulumi.StringPtrInput
	// The source archive object (file) in archive bucket.
	SourceArchiveObject pulumi.StringPtrInput
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository FunctionSourceRepositoryPtrInput
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout pulumi.IntPtrInput
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
	TriggerHttp pulumi.BoolPtrInput
	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector pulumi.StringPtrInput
	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
	VpcConnectorEgressSettings pulumi.StringPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
	AvailableMemoryMb *int `pulumi:"availableMemoryMb"`
	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables map[string]interface{} `pulumi:"buildEnvironmentVariables"`
	// Description of the function.
	Description *string `pulumi:"description"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint *string `pulumi:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger *FunctionEventTrigger `pulumi:"eventTrigger"`
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl *string `pulumi:"httpsTriggerUrl"`
	// String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings *string `pulumi:"ingressSettings"`
	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]interface{} `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances *int `pulumi:"maxInstances"`
	// A user-defined name of the function. Function names must be unique globally.
	Name *string `pulumi:"name"`
	// Project of the function. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region of function. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The runtime in which the function is going to run.
	// Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
	Runtime string `pulumi:"runtime"`
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket *string `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject *string `pulumi:"sourceArchiveObject"`
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository *FunctionSourceRepository `pulumi:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout *int `pulumi:"timeout"`
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
	TriggerHttp *bool `pulumi:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector *string `pulumi:"vpcConnector"`
	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
	VpcConnectorEgressSettings *string `pulumi:"vpcConnectorEgressSettings"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
	AvailableMemoryMb pulumi.IntPtrInput
	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables pulumi.MapInput
	// Description of the function.
	Description pulumi.StringPtrInput
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint pulumi.StringPtrInput
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables pulumi.MapInput
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger FunctionEventTriggerPtrInput
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl pulumi.StringPtrInput
	// String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.MapInput
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances pulumi.IntPtrInput
	// A user-defined name of the function. Function names must be unique globally.
	Name pulumi.StringPtrInput
	// Project of the function. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region of function. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The runtime in which the function is going to run.
	// Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
	Runtime pulumi.StringInput
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail pulumi.StringPtrInput
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket pulumi.StringPtrInput
	// The source archive object (file) in archive bucket.
	SourceArchiveObject pulumi.StringPtrInput
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository FunctionSourceRepositoryPtrInput
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout pulumi.IntPtrInput
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
	TriggerHttp pulumi.BoolPtrInput
	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector pulumi.StringPtrInput
	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
	VpcConnectorEgressSettings pulumi.StringPtrInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((*Function)(nil))
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

func (i *Function) ToFunctionPtrOutput() FunctionPtrOutput {
	return i.ToFunctionPtrOutputWithContext(context.Background())
}

func (i *Function) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionPtrOutput)
}

type FunctionPtrInput interface {
	pulumi.Input

	ToFunctionPtrOutput() FunctionPtrOutput
	ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput
}

type functionPtrType FunctionArgs

func (*functionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil))
}

func (i *functionPtrType) ToFunctionPtrOutput() FunctionPtrOutput {
	return i.ToFunctionPtrOutputWithContext(context.Background())
}

func (i *functionPtrType) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionPtrOutput)
}

// FunctionArrayInput is an input type that accepts FunctionArray and FunctionArrayOutput values.
// You can construct a concrete instance of `FunctionArrayInput` via:
//
//          FunctionArray{ FunctionArgs{...} }
type FunctionArrayInput interface {
	pulumi.Input

	ToFunctionArrayOutput() FunctionArrayOutput
	ToFunctionArrayOutputWithContext(context.Context) FunctionArrayOutput
}

type FunctionArray []FunctionInput

func (FunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (i FunctionArray) ToFunctionArrayOutput() FunctionArrayOutput {
	return i.ToFunctionArrayOutputWithContext(context.Background())
}

func (i FunctionArray) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionArrayOutput)
}

// FunctionMapInput is an input type that accepts FunctionMap and FunctionMapOutput values.
// You can construct a concrete instance of `FunctionMapInput` via:
//
//          FunctionMap{ "key": FunctionArgs{...} }
type FunctionMapInput interface {
	pulumi.Input

	ToFunctionMapOutput() FunctionMapOutput
	ToFunctionMapOutputWithContext(context.Context) FunctionMapOutput
}

type FunctionMap map[string]FunctionInput

func (FunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (i FunctionMap) ToFunctionMapOutput() FunctionMapOutput {
	return i.ToFunctionMapOutputWithContext(context.Background())
}

func (i FunctionMap) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionMapOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Function)(nil))
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionPtrOutput() FunctionPtrOutput {
	return o.ToFunctionPtrOutputWithContext(context.Background())
}

func (o FunctionOutput) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Function) *Function {
		return &v
	}).(FunctionPtrOutput)
}

type FunctionPtrOutput struct{ *pulumi.OutputState }

func (FunctionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil))
}

func (o FunctionPtrOutput) ToFunctionPtrOutput() FunctionPtrOutput {
	return o
}

func (o FunctionPtrOutput) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return o
}

func (o FunctionPtrOutput) Elem() FunctionOutput {
	return o.ApplyT(func(v *Function) Function {
		if v != nil {
			return *v
		}
		var ret Function
		return ret
	}).(FunctionOutput)
}

type FunctionArrayOutput struct{ *pulumi.OutputState }

func (FunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Function)(nil))
}

func (o FunctionArrayOutput) ToFunctionArrayOutput() FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) Index(i pulumi.IntInput) FunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Function {
		return vs[0].([]Function)[vs[1].(int)]
	}).(FunctionOutput)
}

type FunctionMapOutput struct{ *pulumi.OutputState }

func (FunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Function)(nil))
}

func (o FunctionMapOutput) ToFunctionMapOutput() FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) MapIndex(k pulumi.StringInput) FunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Function {
		return vs[0].(map[string]Function)[vs[1].(string)]
	}).(FunctionOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionOutput{})
	pulumi.RegisterOutputType(FunctionPtrOutput{})
	pulumi.RegisterOutputType(FunctionArrayOutput{})
	pulumi.RegisterOutputType(FunctionMapOutput{})
}
