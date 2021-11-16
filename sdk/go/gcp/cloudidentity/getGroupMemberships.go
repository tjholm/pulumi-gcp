// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudidentity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get list of the Cloud Identity Group Memberships within a given Group.
//
// https://cloud.google.com/identity/docs/concepts/overview#memberships
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudidentity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudidentity.GetGroupMemberships(ctx, &cloudidentity.GetGroupMembershipsArgs{
// 			Group: "groups/123eab45c6defghi",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetGroupMemberships(ctx *pulumi.Context, args *GetGroupMembershipsArgs, opts ...pulumi.InvokeOption) (*GetGroupMembershipsResult, error) {
	var rv GetGroupMembershipsResult
	err := ctx.Invoke("gcp:cloudidentity/getGroupMemberships:getGroupMemberships", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupMemberships.
type GetGroupMembershipsArgs struct {
	// The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}.
	Group string `pulumi:"group"`
}

// A collection of values returned by getGroupMemberships.
type GetGroupMembershipsResult struct {
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of memberships under the given group. Structure is documented below.
	Memberships []GetGroupMembershipsMembership `pulumi:"memberships"`
}

func GetGroupMembershipsOutput(ctx *pulumi.Context, args GetGroupMembershipsOutputArgs, opts ...pulumi.InvokeOption) GetGroupMembershipsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupMembershipsResult, error) {
			args := v.(GetGroupMembershipsArgs)
			r, err := GetGroupMemberships(ctx, &args, opts...)
			return *r, err
		}).(GetGroupMembershipsResultOutput)
}

// A collection of arguments for invoking getGroupMemberships.
type GetGroupMembershipsOutputArgs struct {
	// The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}.
	Group pulumi.StringInput `pulumi:"group"`
}

func (GetGroupMembershipsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupMembershipsArgs)(nil)).Elem()
}

// A collection of values returned by getGroupMemberships.
type GetGroupMembershipsResultOutput struct{ *pulumi.OutputState }

func (GetGroupMembershipsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupMembershipsResult)(nil)).Elem()
}

func (o GetGroupMembershipsResultOutput) ToGetGroupMembershipsResultOutput() GetGroupMembershipsResultOutput {
	return o
}

func (o GetGroupMembershipsResultOutput) ToGetGroupMembershipsResultOutputWithContext(ctx context.Context) GetGroupMembershipsResultOutput {
	return o
}

func (o GetGroupMembershipsResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupMembershipsResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupMembershipsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupMembershipsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of memberships under the given group. Structure is documented below.
func (o GetGroupMembershipsResultOutput) Memberships() GetGroupMembershipsMembershipArrayOutput {
	return o.ApplyT(func(v GetGroupMembershipsResult) []GetGroupMembershipsMembership { return v.Memberships }).(GetGroupMembershipsMembershipArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupMembershipsResultOutput{})
}
