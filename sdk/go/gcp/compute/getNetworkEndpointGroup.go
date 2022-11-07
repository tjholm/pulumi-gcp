// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access a Network Endpoint Group's attributes.
//
// The NEG may be found by providing either a `selfLink`, or a `name` and a `zone`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = compute.LookupNetworkEndpointGroup(ctx, &compute.LookupNetworkEndpointGroupArgs{
//				Name: pulumi.StringRef("k8s1-abcdef01-myns-mysvc-8080-4b6bac43"),
//				Zone: pulumi.StringRef("us-central1-a"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.LookupNetworkEndpointGroup(ctx, &compute.LookupNetworkEndpointGroupArgs{
//				SelfLink: pulumi.StringRef("https://www.googleapis.com/compute/v1/projects/myproject/zones/us-central1-a/networkEndpointGroups/k8s1-abcdef01-myns-mysvc-8080-4b6bac43"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNetworkEndpointGroup(ctx *pulumi.Context, args *LookupNetworkEndpointGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkEndpointGroupResult, error) {
	var rv LookupNetworkEndpointGroupResult
	err := ctx.Invoke("gcp:compute/getNetworkEndpointGroup:getNetworkEndpointGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkEndpointGroup.
type LookupNetworkEndpointGroupArgs struct {
	// The Network Endpoint Group name.
	// Provide either this or a `selfLink`.
	Name *string `pulumi:"name"`
	// The ID of the project to list versions in.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Network Endpoint Group self\_link.
	SelfLink *string `pulumi:"selfLink"`
	// The Network Endpoint Group availability zone.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getNetworkEndpointGroup.
type LookupNetworkEndpointGroupResult struct {
	// The NEG default port.
	DefaultPort int `pulumi:"defaultPort"`
	// The NEG description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// The network to which all network endpoints in the NEG belong.
	Network string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group.
	NetworkEndpointType string  `pulumi:"networkEndpointType"`
	Project             *string `pulumi:"project"`
	SelfLink            *string `pulumi:"selfLink"`
	// Number of network endpoints in the network endpoint group.
	Size int `pulumi:"size"`
	// subnetwork to which all network endpoints in the NEG belong.
	Subnetwork string  `pulumi:"subnetwork"`
	Zone       *string `pulumi:"zone"`
}

func LookupNetworkEndpointGroupOutput(ctx *pulumi.Context, args LookupNetworkEndpointGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkEndpointGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkEndpointGroupResult, error) {
			args := v.(LookupNetworkEndpointGroupArgs)
			r, err := LookupNetworkEndpointGroup(ctx, &args, opts...)
			var s LookupNetworkEndpointGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkEndpointGroupResultOutput)
}

// A collection of arguments for invoking getNetworkEndpointGroup.
type LookupNetworkEndpointGroupOutputArgs struct {
	// The Network Endpoint Group name.
	// Provide either this or a `selfLink`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project to list versions in.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The Network Endpoint Group self\_link.
	SelfLink pulumi.StringPtrInput `pulumi:"selfLink"`
	// The Network Endpoint Group availability zone.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupNetworkEndpointGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkEndpointGroupArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkEndpointGroup.
type LookupNetworkEndpointGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkEndpointGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkEndpointGroupResult)(nil)).Elem()
}

func (o LookupNetworkEndpointGroupResultOutput) ToLookupNetworkEndpointGroupResultOutput() LookupNetworkEndpointGroupResultOutput {
	return o
}

func (o LookupNetworkEndpointGroupResultOutput) ToLookupNetworkEndpointGroupResultOutputWithContext(ctx context.Context) LookupNetworkEndpointGroupResultOutput {
	return o
}

// The NEG default port.
func (o LookupNetworkEndpointGroupResultOutput) DefaultPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) int { return v.DefaultPort }).(pulumi.IntOutput)
}

// The NEG description.
func (o LookupNetworkEndpointGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkEndpointGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkEndpointGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The network to which all network endpoints in the NEG belong.
func (o LookupNetworkEndpointGroupResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Network }).(pulumi.StringOutput)
}

// Type of network endpoints in this network endpoint group.
func (o LookupNetworkEndpointGroupResultOutput) NetworkEndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.NetworkEndpointType }).(pulumi.StringOutput)
}

func (o LookupNetworkEndpointGroupResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkEndpointGroupResultOutput) SelfLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) *string { return v.SelfLink }).(pulumi.StringPtrOutput)
}

// Number of network endpoints in the network endpoint group.
func (o LookupNetworkEndpointGroupResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) int { return v.Size }).(pulumi.IntOutput)
}

// subnetwork to which all network endpoints in the NEG belong.
func (o LookupNetworkEndpointGroupResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

func (o LookupNetworkEndpointGroupResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkEndpointGroupResultOutput{})
}
