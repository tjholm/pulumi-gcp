// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudidentity

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get list of the Cloud Identity Groups under a customer or namespace.
//
// https://cloud.google.com/identity/docs/concepts/overview#groups
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/cloudidentity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudidentity.GetGroups(ctx, &cloudidentity.GetGroupsArgs{
// 			Parent: "customers/A01b123xz",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	var rv GetGroupsResult
	err := ctx.Invoke("gcp:cloudidentity/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	// The parent resource under which to list all Groups. Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_id} for Google Groups.
	Parent string `pulumi:"parent"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	// The list of groups under the provided customer or namespace. Structure is documented below.
	Groups []GetGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Parent string `pulumi:"parent"`
}
