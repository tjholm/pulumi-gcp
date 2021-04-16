// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An isolated set of Cloud Spanner resources on which databases can be
// hosted.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/spanner/)
//
// ## Example Usage
// ### Spanner Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.NewInstance(ctx, "example", &spanner.InstanceArgs{
// 			Config:      pulumi.String("regional-us-central1"),
// 			DisplayName: pulumi.String("Test Spanner Instance"),
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			NumNodes: pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Spanner Instance Multi Regional
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.NewInstance(ctx, "example", &spanner.InstanceArgs{
// 			Config:      pulumi.String("nam-eur-asia1"),
// 			DisplayName: pulumi.String("Multi Regional Instance"),
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			NumNodes: pulumi.Int(2),
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
// Instance can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:spanner/instance:Instance default projects/{{project}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:spanner/instance:Instance default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:spanner/instance:Instance default {{name}}
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config pulumi.StringOutput `pulumi:"config"`
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes allocated to this instance.
	NumNodes pulumi.IntPtrOutput `pulumi:"numNodes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Instance status: 'CREATING' or 'READY'.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:spanner/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:spanner/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config *string `pulumi:"config"`
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName *string `pulumi:"displayName"`
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name *string `pulumi:"name"`
	// The number of nodes allocated to this instance.
	NumNodes *int `pulumi:"numNodes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Instance status: 'CREATING' or 'READY'.
	State *string `pulumi:"state"`
}

type InstanceState struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config pulumi.StringPtrInput
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName pulumi.StringPtrInput
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name pulumi.StringPtrInput
	// The number of nodes allocated to this instance.
	NumNodes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Instance status: 'CREATING' or 'READY'.
	State pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config string `pulumi:"config"`
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName string `pulumi:"displayName"`
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name *string `pulumi:"name"`
	// The number of nodes allocated to this instance.
	NumNodes *int `pulumi:"numNodes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config pulumi.StringInput
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName pulumi.StringInput
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name pulumi.StringPtrInput
	// The number of nodes allocated to this instance.
	NumNodes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
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
	return reflect.TypeOf(([]*Instance)(nil))
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
	return reflect.TypeOf((map[string]*Instance)(nil))
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

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
	return o.ApplyT(func(v Instance) *Instance {
		return &v
	}).(InstancePtrOutput)
}

type InstancePtrOutput struct {
	*pulumi.OutputState
}

func (InstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (o InstancePtrOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o
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
