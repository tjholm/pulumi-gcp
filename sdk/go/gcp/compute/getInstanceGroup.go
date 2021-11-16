// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Compute Instance Group within GCE.
// For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "instance-group-name"
// 		opt1 := "us-central1-a"
// 		_, err := compute.LookupInstanceGroup(ctx, &compute.LookupInstanceGroupArgs{
// 			Name: &opt0,
// 			Zone: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupInstanceGroup(ctx *pulumi.Context, args *LookupInstanceGroupArgs, opts ...pulumi.InvokeOption) (*LookupInstanceGroupResult, error) {
	var rv LookupInstanceGroupResult
	err := ctx.Invoke("gcp:compute/getInstanceGroup:getInstanceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceGroup.
type LookupInstanceGroupArgs struct {
	// The name of the instance group. Either `name` or `selfLink` must be provided.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The self link of the instance group. Either `name` or `selfLink` must be provided.
	SelfLink *string `pulumi:"selfLink"`
	// The zone of the instance group. If referencing the instance group by name
	// and `zone` is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceGroup.
type LookupInstanceGroupResult struct {
	// Textual description of the instance group.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of instances in the group.
	Instances []string `pulumi:"instances"`
	Name      *string  `pulumi:"name"`
	// List of named ports in the group.
	NamedPorts []GetInstanceGroupNamedPortType `pulumi:"namedPorts"`
	// The URL of the network the instance group is in.
	Network string `pulumi:"network"`
	Project string `pulumi:"project"`
	// The URI of the resource.
	SelfLink string `pulumi:"selfLink"`
	// The number of instances in the group.
	Size int    `pulumi:"size"`
	Zone string `pulumi:"zone"`
}

func LookupInstanceGroupOutput(ctx *pulumi.Context, args LookupInstanceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceGroupResult, error) {
			args := v.(LookupInstanceGroupArgs)
			r, err := LookupInstanceGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceGroupResultOutput)
}

// A collection of arguments for invoking getInstanceGroup.
type LookupInstanceGroupOutputArgs struct {
	// The name of the instance group. Either `name` or `selfLink` must be provided.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The self link of the instance group. Either `name` or `selfLink` must be provided.
	SelfLink pulumi.StringPtrInput `pulumi:"selfLink"`
	// The zone of the instance group. If referencing the instance group by name
	// and `zone` is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupInstanceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceGroup.
type LookupInstanceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupResult)(nil)).Elem()
}

func (o LookupInstanceGroupResultOutput) ToLookupInstanceGroupResultOutput() LookupInstanceGroupResultOutput {
	return o
}

func (o LookupInstanceGroupResultOutput) ToLookupInstanceGroupResultOutputWithContext(ctx context.Context) LookupInstanceGroupResultOutput {
	return o
}

// Textual description of the instance group.
func (o LookupInstanceGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of instances in the group.
func (o LookupInstanceGroupResultOutput) Instances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) []string { return v.Instances }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// List of named ports in the group.
func (o LookupInstanceGroupResultOutput) NamedPorts() GetInstanceGroupNamedPortTypeArrayOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) []GetInstanceGroupNamedPortType { return v.NamedPorts }).(GetInstanceGroupNamedPortTypeArrayOutput)
}

// The URL of the network the instance group is in.
func (o LookupInstanceGroupResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupInstanceGroupResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Project }).(pulumi.StringOutput)
}

// The URI of the resource.
func (o LookupInstanceGroupResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The number of instances in the group.
func (o LookupInstanceGroupResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) int { return v.Size }).(pulumi.IntOutput)
}

func (o LookupInstanceGroupResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceGroupResultOutput{})
}
