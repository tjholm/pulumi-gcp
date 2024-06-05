// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workstations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkstationConfigIamPolicy(ctx *pulumi.Context, args *LookupWorkstationConfigIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupWorkstationConfigIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkstationConfigIamPolicyResult
	err := ctx.Invoke("gcp:workstations/getWorkstationConfigIamPolicy:getWorkstationConfigIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkstationConfigIamPolicy.
type LookupWorkstationConfigIamPolicyArgs struct {
	// The location where the workstation cluster config should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project              *string `pulumi:"project"`
	WorkstationClusterId string  `pulumi:"workstationClusterId"`
	WorkstationConfigId  string  `pulumi:"workstationConfigId"`
}

// A collection of values returned by getWorkstationConfigIamPolicy.
type LookupWorkstationConfigIamPolicyResult struct {
	// (Computed) The etag of the IAM policy.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
	// (Required only by `workstations.WorkstationConfigIamPolicy`) The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData           string `pulumi:"policyData"`
	Project              string `pulumi:"project"`
	WorkstationClusterId string `pulumi:"workstationClusterId"`
	WorkstationConfigId  string `pulumi:"workstationConfigId"`
}

func LookupWorkstationConfigIamPolicyOutput(ctx *pulumi.Context, args LookupWorkstationConfigIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupWorkstationConfigIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkstationConfigIamPolicyResult, error) {
			args := v.(LookupWorkstationConfigIamPolicyArgs)
			r, err := LookupWorkstationConfigIamPolicy(ctx, &args, opts...)
			var s LookupWorkstationConfigIamPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkstationConfigIamPolicyResultOutput)
}

// A collection of arguments for invoking getWorkstationConfigIamPolicy.
type LookupWorkstationConfigIamPolicyOutputArgs struct {
	// The location where the workstation cluster config should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project              pulumi.StringPtrInput `pulumi:"project"`
	WorkstationClusterId pulumi.StringInput    `pulumi:"workstationClusterId"`
	WorkstationConfigId  pulumi.StringInput    `pulumi:"workstationConfigId"`
}

func (LookupWorkstationConfigIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkstationConfigIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getWorkstationConfigIamPolicy.
type LookupWorkstationConfigIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupWorkstationConfigIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkstationConfigIamPolicyResult)(nil)).Elem()
}

func (o LookupWorkstationConfigIamPolicyResultOutput) ToLookupWorkstationConfigIamPolicyResultOutput() LookupWorkstationConfigIamPolicyResultOutput {
	return o
}

func (o LookupWorkstationConfigIamPolicyResultOutput) ToLookupWorkstationConfigIamPolicyResultOutputWithContext(ctx context.Context) LookupWorkstationConfigIamPolicyResultOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o LookupWorkstationConfigIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationConfigIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupWorkstationConfigIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationConfigIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkstationConfigIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationConfigIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// (Required only by `workstations.WorkstationConfigIamPolicy`) The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o LookupWorkstationConfigIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationConfigIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupWorkstationConfigIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationConfigIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupWorkstationConfigIamPolicyResultOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationConfigIamPolicyResult) string { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o LookupWorkstationConfigIamPolicyResultOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationConfigIamPolicyResult) string { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkstationConfigIamPolicyResultOutput{})
}
