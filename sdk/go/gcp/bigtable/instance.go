// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## +---
//
// subcategory: "Cloud Bigtable"
// layout: "google"
// page_title: "Google: bigtable.Instance"
// sidebar_current: "docs-google-bigtable-instance"
// description: |-
//
//	Creates a Google Bigtable instance.
//
// ---
//
// # bigtable.Instance
//
// Creates a Google Bigtable instance. For more information see:
//
// * [API documentation](https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/bigtable/docs)
//
// ## Example Usage
// ### Simple Instance
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigtable.NewInstance(ctx, "production-instance", &bigtable.InstanceArgs{
//				Clusters: bigtable.InstanceClusterArray{
//					&bigtable.InstanceClusterArgs{
//						ClusterId:   pulumi.String("tf-instance-cluster"),
//						NumNodes:    pulumi.Int(1),
//						StorageType: pulumi.String("HDD"),
//					},
//				},
//				Labels: pulumi.StringMap{
//					"my-label": pulumi.String("prod-label"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Replicated Instance
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigtable.NewInstance(ctx, "production-instance", &bigtable.InstanceArgs{
//				Clusters: bigtable.InstanceClusterArray{
//					&bigtable.InstanceClusterArgs{
//						ClusterId:   pulumi.String("tf-instance-cluster1"),
//						NumNodes:    pulumi.Int(1),
//						StorageType: pulumi.String("HDD"),
//						Zone:        pulumi.String("us-central1-c"),
//					},
//					&bigtable.InstanceClusterArgs{
//						AutoscalingConfig: &bigtable.InstanceClusterAutoscalingConfigArgs{
//							CpuTarget: pulumi.Int(50),
//							MaxNodes:  pulumi.Int(3),
//							MinNodes:  pulumi.Int(1),
//						},
//						ClusterId:   pulumi.String("tf-instance-cluster2"),
//						StorageType: pulumi.String("HDD"),
//						Zone:        pulumi.String("us-central1-b"),
//					},
//				},
//				Labels: pulumi.StringMap{
//					"my-label": pulumi.String("prod-label"),
//				},
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
// # Bigtable Instances can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:bigtable/instance:Instance default projects/{{project}}/instances/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/instance:Instance default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/instance:Instance default {{name}}
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters InstanceClusterArrayOutput `pulumi:"clusters"`
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	var resource Instance
	err := ctx.RegisterResource("gcp:bigtable/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:bigtable/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters []InstanceCluster `pulumi:"clusters"`
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName *string `pulumi:"displayName"`
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType *string `pulumi:"instanceType"`
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]string `pulumi:"labels"`
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type InstanceState struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters InstanceClusterArrayInput
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection pulumi.BoolPtrInput
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName pulumi.StringPtrInput
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.StringMapInput
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters []InstanceCluster `pulumi:"clusters"`
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName *string `pulumi:"displayName"`
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType *string `pulumi:"instanceType"`
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]string `pulumi:"labels"`
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters InstanceClusterArrayInput
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection pulumi.BoolPtrInput
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName pulumi.StringPtrInput
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.StringMapInput
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
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
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
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
//	InstanceMap{ "key": InstanceArgs{...} }
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
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
// See structure below.
func (o InstanceOutput) Clusters() InstanceClusterArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceClusterArrayOutput { return v.Clusters }).(InstanceClusterArrayOutput)
}

// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
func (o InstanceOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
func (o InstanceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
//
// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
func (o InstanceOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
func (o InstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o InstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
