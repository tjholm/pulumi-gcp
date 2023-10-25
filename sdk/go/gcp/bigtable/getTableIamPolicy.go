// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the current IAM policy data for a Bigtable Table.
//
// ## example
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
//			_, err := bigtable.LookupTableIamPolicy(ctx, &bigtable.LookupTableIamPolicyArgs{
//				Instance: google_bigtable_instance.Instance.Name,
//				Table:    google_bigtable_table.Table.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTableIamPolicy(ctx *pulumi.Context, args *LookupTableIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupTableIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTableIamPolicyResult
	err := ctx.Invoke("gcp:bigtable/getTableIamPolicy:getTableIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTableIamPolicy.
type LookupTableIamPolicyArgs struct {
	// The name or relative resource id of the instance that owns the table.
	Instance string  `pulumi:"instance"`
	Project  *string `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table string `pulumi:"table"`
}

// A collection of values returned by getTableIamPolicy.
type LookupTableIamPolicyResult struct {
	// (Computed) The etag of the IAM policy.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Instance string `pulumi:"instance"`
	// (Computed) The policy data
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Table      string `pulumi:"table"`
}

func LookupTableIamPolicyOutput(ctx *pulumi.Context, args LookupTableIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupTableIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableIamPolicyResult, error) {
			args := v.(LookupTableIamPolicyArgs)
			r, err := LookupTableIamPolicy(ctx, &args, opts...)
			var s LookupTableIamPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableIamPolicyResultOutput)
}

// A collection of arguments for invoking getTableIamPolicy.
type LookupTableIamPolicyOutputArgs struct {
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringInput    `pulumi:"instance"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringInput `pulumi:"table"`
}

func (LookupTableIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getTableIamPolicy.
type LookupTableIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupTableIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableIamPolicyResult)(nil)).Elem()
}

func (o LookupTableIamPolicyResultOutput) ToLookupTableIamPolicyResultOutput() LookupTableIamPolicyResultOutput {
	return o
}

func (o LookupTableIamPolicyResultOutput) ToLookupTableIamPolicyResultOutputWithContext(ctx context.Context) LookupTableIamPolicyResultOutput {
	return o
}

func (o LookupTableIamPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTableIamPolicyResult] {
	return pulumix.Output[LookupTableIamPolicyResult]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o LookupTableIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTableIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTableIamPolicyResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableIamPolicyResult) string { return v.Instance }).(pulumi.StringOutput)
}

// (Computed) The policy data
func (o LookupTableIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupTableIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupTableIamPolicyResultOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableIamPolicyResult) string { return v.Table }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableIamPolicyResultOutput{})
}
