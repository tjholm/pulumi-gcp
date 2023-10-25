// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the current IAM policy data for accesspolicy
//
// ## example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := accesscontextmanager.LookupAccessPolicyIamPolicy(ctx, &accesscontextmanager.LookupAccessPolicyIamPolicyArgs{
//				Name: google_access_context_manager_access_policy.AccessPolicy.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAccessPolicyIamPolicy(ctx *pulumi.Context, args *LookupAccessPolicyIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessPolicyIamPolicyResult
	err := ctx.Invoke("gcp:accesscontextmanager/getAccessPolicyIamPolicy:getAccessPolicyIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessPolicyIamPolicy.
type LookupAccessPolicyIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name string `pulumi:"name"`
}

// A collection of values returned by getAccessPolicyIamPolicy.
type LookupAccessPolicyIamPolicyResult struct {
	// (Computed) The etag of the IAM policy.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// (Required only by `accesscontextmanager.AccessPolicyIamPolicy`) The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

func LookupAccessPolicyIamPolicyOutput(ctx *pulumi.Context, args LookupAccessPolicyIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPolicyIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPolicyIamPolicyResult, error) {
			args := v.(LookupAccessPolicyIamPolicyArgs)
			r, err := LookupAccessPolicyIamPolicy(ctx, &args, opts...)
			var s LookupAccessPolicyIamPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessPolicyIamPolicyResultOutput)
}

// A collection of arguments for invoking getAccessPolicyIamPolicy.
type LookupAccessPolicyIamPolicyOutputArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupAccessPolicyIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getAccessPolicyIamPolicy.
type LookupAccessPolicyIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPolicyIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyIamPolicyResult)(nil)).Elem()
}

func (o LookupAccessPolicyIamPolicyResultOutput) ToLookupAccessPolicyIamPolicyResultOutput() LookupAccessPolicyIamPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyIamPolicyResultOutput) ToLookupAccessPolicyIamPolicyResultOutputWithContext(ctx context.Context) LookupAccessPolicyIamPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyIamPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccessPolicyIamPolicyResult] {
	return pulumix.Output[LookupAccessPolicyIamPolicyResult]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o LookupAccessPolicyIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccessPolicyIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessPolicyIamPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyIamPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// (Required only by `accesscontextmanager.AccessPolicyIamPolicy`) The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o LookupAccessPolicyIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPolicyIamPolicyResultOutput{})
}
