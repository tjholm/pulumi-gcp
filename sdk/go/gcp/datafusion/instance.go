// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafusion

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Data Fusion instance.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/data-fusion/docs/reference/rest/v1beta1/projects.locations.instances)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/data-fusion/docs/)
//
// ## Example Usage
// ### Data Fusion Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datafusion"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datafusion.NewInstance(ctx, "basicInstance", &datafusion.InstanceArgs{
// 			Region: pulumi.String("us-central1"),
// 			Type:   pulumi.String("BASIC"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Data Fusion Instance Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/appengine"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datafusion"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_default, err := appengine.GetDefaultServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafusion.NewInstance(ctx, "extendedInstance", &datafusion.InstanceArgs{
// 			Description:                 pulumi.String("My Data Fusion instance"),
// 			Region:                      pulumi.String("us-central1"),
// 			Type:                        pulumi.String("BASIC"),
// 			EnableStackdriverLogging:    pulumi.Bool(true),
// 			EnableStackdriverMonitoring: pulumi.Bool(true),
// 			Labels: pulumi.StringMap{
// 				"example_key": pulumi.String("example_value"),
// 			},
// 			PrivateInstance: pulumi.Bool(true),
// 			NetworkConfig: &datafusion.InstanceNetworkConfigArgs{
// 				Network:      pulumi.String("default"),
// 				IpAllocation: pulumi.String("10.89.48.0/22"),
// 			},
// 			Version:                pulumi.String("6.3.0"),
// 			DataprocServiceAccount: pulumi.String(_default.Email),
// 		}, pulumi.Provider(google_beta))
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
// Instance can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:datafusion/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:datafusion/instance:Instance default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:datafusion/instance:Instance default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:datafusion/instance:Instance default {{name}}
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	DataprocServiceAccount pulumi.StringPtrOutput `pulumi:"dataprocServiceAccount"`
	// An optional description of the instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolPtrOutput `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolPtrOutput `pulumi:"enableStackdriverMonitoring"`
	// The resource labels for instance to use to annotate any related underlying resources,
	// such as Compute Engine VMs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The ID of the instance or a fully qualified identifier for the instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	// Structure is documented below.
	NetworkConfig InstanceNetworkConfigPtrOutput `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapOutput `pulumi:"options"`
	// Specifies whether the Data Fusion instance should be private. If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	PrivateInstance pulumi.BoolPtrOutput `pulumi:"privateInstance"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the Data Fusion instance.
	Region pulumi.StringOutput `pulumi:"region"`
	// Service account which will be used to access resources in the customer project.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Endpoint on which the Data Fusion UI and REST APIs are accessible.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// The current state of this Data Fusion instance. - CREATING: Instance is being created - RUNNING: Instance is running and
	// ready for requests - FAILED: Instance creation failed - DELETING: Instance is being deleted - UPGRADING: Instance is
	// being upgraded - RESTARTING: Instance is being restarted
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// Represents the type of Data Fusion instance. Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	//   using point and click UI. However, there are certain limitations, such as fewer number
	//   of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	//   available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	//   with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	//   pipelines at low cost.
	//   Possible values are `BASIC`, `ENTERPRISE`, and `DEVELOPER`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Current version of the Data Fusion.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:datafusion/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:datafusion/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	CreateTime *string `pulumi:"createTime"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	DataprocServiceAccount *string `pulumi:"dataprocServiceAccount"`
	// An optional description of the instance.
	Description *string `pulumi:"description"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring *bool `pulumi:"enableStackdriverMonitoring"`
	// The resource labels for instance to use to annotate any related underlying resources,
	// such as Compute Engine VMs.
	Labels map[string]string `pulumi:"labels"`
	// The ID of the instance or a fully qualified identifier for the instance.
	Name *string `pulumi:"name"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	// Structure is documented below.
	NetworkConfig *InstanceNetworkConfig `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `pulumi:"options"`
	// Specifies whether the Data Fusion instance should be private. If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	PrivateInstance *bool `pulumi:"privateInstance"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the Data Fusion instance.
	Region *string `pulumi:"region"`
	// Service account which will be used to access resources in the customer project.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Endpoint on which the Data Fusion UI and REST APIs are accessible.
	ServiceEndpoint *string `pulumi:"serviceEndpoint"`
	// The current state of this Data Fusion instance. - CREATING: Instance is being created - RUNNING: Instance is running and
	// ready for requests - FAILED: Instance creation failed - DELETING: Instance is being deleted - UPGRADING: Instance is
	// being upgraded - RESTARTING: Instance is being restarted
	State *string `pulumi:"state"`
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage *string `pulumi:"stateMessage"`
	// Represents the type of Data Fusion instance. Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	//   using point and click UI. However, there are certain limitations, such as fewer number
	//   of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	//   available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	//   with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	//   pipelines at low cost.
	//   Possible values are `BASIC`, `ENTERPRISE`, and `DEVELOPER`.
	Type *string `pulumi:"type"`
	// The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	UpdateTime *string `pulumi:"updateTime"`
	// Current version of the Data Fusion.
	Version *string `pulumi:"version"`
}

type InstanceState struct {
	// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	CreateTime pulumi.StringPtrInput
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	DataprocServiceAccount pulumi.StringPtrInput
	// An optional description of the instance.
	Description pulumi.StringPtrInput
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolPtrInput
	// The resource labels for instance to use to annotate any related underlying resources,
	// such as Compute Engine VMs.
	Labels pulumi.StringMapInput
	// The ID of the instance or a fully qualified identifier for the instance.
	Name pulumi.StringPtrInput
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	// Structure is documented below.
	NetworkConfig InstanceNetworkConfigPtrInput
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapInput
	// Specifies whether the Data Fusion instance should be private. If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	PrivateInstance pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the Data Fusion instance.
	Region pulumi.StringPtrInput
	// Service account which will be used to access resources in the customer project.
	ServiceAccount pulumi.StringPtrInput
	// Endpoint on which the Data Fusion UI and REST APIs are accessible.
	ServiceEndpoint pulumi.StringPtrInput
	// The current state of this Data Fusion instance. - CREATING: Instance is being created - RUNNING: Instance is running and
	// ready for requests - FAILED: Instance creation failed - DELETING: Instance is being deleted - UPGRADING: Instance is
	// being upgraded - RESTARTING: Instance is being restarted
	State pulumi.StringPtrInput
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage pulumi.StringPtrInput
	// Represents the type of Data Fusion instance. Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	//   using point and click UI. However, there are certain limitations, such as fewer number
	//   of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	//   available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	//   with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	//   pipelines at low cost.
	//   Possible values are `BASIC`, `ENTERPRISE`, and `DEVELOPER`.
	Type pulumi.StringPtrInput
	// The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	UpdateTime pulumi.StringPtrInput
	// Current version of the Data Fusion.
	Version pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	DataprocServiceAccount *string `pulumi:"dataprocServiceAccount"`
	// An optional description of the instance.
	Description *string `pulumi:"description"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring *bool `pulumi:"enableStackdriverMonitoring"`
	// The resource labels for instance to use to annotate any related underlying resources,
	// such as Compute Engine VMs.
	Labels map[string]string `pulumi:"labels"`
	// The ID of the instance or a fully qualified identifier for the instance.
	Name *string `pulumi:"name"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	// Structure is documented below.
	NetworkConfig *InstanceNetworkConfig `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `pulumi:"options"`
	// Specifies whether the Data Fusion instance should be private. If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	PrivateInstance *bool `pulumi:"privateInstance"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the Data Fusion instance.
	Region *string `pulumi:"region"`
	// Represents the type of Data Fusion instance. Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	//   using point and click UI. However, there are certain limitations, such as fewer number
	//   of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	//   available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	//   with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	//   pipelines at low cost.
	//   Possible values are `BASIC`, `ENTERPRISE`, and `DEVELOPER`.
	Type string `pulumi:"type"`
	// Current version of the Data Fusion.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	DataprocServiceAccount pulumi.StringPtrInput
	// An optional description of the instance.
	Description pulumi.StringPtrInput
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolPtrInput
	// The resource labels for instance to use to annotate any related underlying resources,
	// such as Compute Engine VMs.
	Labels pulumi.StringMapInput
	// The ID of the instance or a fully qualified identifier for the instance.
	Name pulumi.StringPtrInput
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	// Structure is documented below.
	NetworkConfig InstanceNetworkConfigPtrInput
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapInput
	// Specifies whether the Data Fusion instance should be private. If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	PrivateInstance pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the Data Fusion instance.
	Region pulumi.StringPtrInput
	// Represents the type of Data Fusion instance. Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	//   using point and click UI. However, there are certain limitations, such as fewer number
	//   of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	//   available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	//   with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	//   pipelines at low cost.
	//   Possible values are `BASIC`, `ENTERPRISE`, and `DEVELOPER`.
	Type pulumi.StringInput
	// Current version of the Data Fusion.
	Version pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *Instance) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

type InstancePtrInput interface {
	pulumi.Input

	ToInstancePtrOutput() InstancePtrOutput
	ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput
}

type instancePtrType InstanceArgs

func (*instancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (i *instancePtrType) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *instancePtrType) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o.ToInstancePtrOutputWithContext(context.Background())
}

func (o InstanceOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Instance) *Instance {
		return &v
	}).(InstancePtrOutput)
}

type InstancePtrOutput struct{ *pulumi.OutputState }

func (InstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (o InstancePtrOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) Elem() InstanceOutput {
	return o.ApplyT(func(v *Instance) Instance {
		if v != nil {
			return *v
		}
		var ret Instance
		return ret
	}).(InstanceOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Instance)(nil))
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Instance {
		return vs[0].([]Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Instance)(nil))
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Instance {
		return vs[0].(map[string]Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstancePtrOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
