// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the current IAM policy data for tunnel
//
// ## example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.LookupTunnelIamPolicy(ctx, &iap.LookupTunnelIamPolicyArgs{
//				Project: pulumi.StringRef(google_project_service.Project_service.Project),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTunnelIamPolicy(ctx *pulumi.Context, args *LookupTunnelIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupTunnelIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTunnelIamPolicyResult
	err := ctx.Invoke("gcp:iap/getTunnelIamPolicy:getTunnelIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTunnelIamPolicy.
type LookupTunnelIamPolicyArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getTunnelIamPolicy.
type LookupTunnelIamPolicyResult struct {
	// (Computed) The etag of the IAM policy.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Required only by `iap.TunnelIamPolicy`) The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupTunnelIamPolicyOutput(ctx *pulumi.Context, args LookupTunnelIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupTunnelIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTunnelIamPolicyResult, error) {
			args := v.(LookupTunnelIamPolicyArgs)
			r, err := LookupTunnelIamPolicy(ctx, &args, opts...)
			var s LookupTunnelIamPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTunnelIamPolicyResultOutput)
}

// A collection of arguments for invoking getTunnelIamPolicy.
type LookupTunnelIamPolicyOutputArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupTunnelIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTunnelIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getTunnelIamPolicy.
type LookupTunnelIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupTunnelIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTunnelIamPolicyResult)(nil)).Elem()
}

func (o LookupTunnelIamPolicyResultOutput) ToLookupTunnelIamPolicyResultOutput() LookupTunnelIamPolicyResultOutput {
	return o
}

func (o LookupTunnelIamPolicyResultOutput) ToLookupTunnelIamPolicyResultOutputWithContext(ctx context.Context) LookupTunnelIamPolicyResultOutput {
	return o
}

func (o LookupTunnelIamPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTunnelIamPolicyResult] {
	return pulumix.Output[LookupTunnelIamPolicyResult]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o LookupTunnelIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTunnelIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTunnelIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTunnelIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Required only by `iap.TunnelIamPolicy`) The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o LookupTunnelIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTunnelIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupTunnelIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTunnelIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTunnelIamPolicyResultOutput{})
}
