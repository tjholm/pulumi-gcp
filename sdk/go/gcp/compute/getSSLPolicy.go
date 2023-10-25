// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an SSL Policy within GCE from its name, for use with Target HTTPS and Target SSL Proxies.
//
//	For more information see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies).
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
//			_, err := compute.LookupSSLPolicy(ctx, &compute.LookupSSLPolicyArgs{
//				Name: "production-ssl-policy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSSLPolicy(ctx *pulumi.Context, args *LookupSSLPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSSLPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSSLPolicyResult
	err := ctx.Invoke("gcp:compute/getSSLPolicy:getSSLPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSSLPolicy.
type LookupSSLPolicyArgs struct {
	// The name of the SSL Policy.
	//
	// ***
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getSSLPolicy.
type LookupSSLPolicyResult struct {
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// If the `profile` is `CUSTOM`, these are the custom encryption
	// ciphers supported by the profile. If the `profile` is *not* `CUSTOM`, this
	// attribute will be empty.
	CustomFeatures []string `pulumi:"customFeatures"`
	// Description of this SSL Policy.
	Description string `pulumi:"description"`
	// The set of enabled encryption ciphers as a result of the policy config
	EnabledFeatures []string `pulumi:"enabledFeatures"`
	// Fingerprint of this resource.
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The minimum supported TLS version of this policy.
	MinTlsVersion string `pulumi:"minTlsVersion"`
	Name          string `pulumi:"name"`
	// The Google-curated or custom profile used by this policy.
	Profile string  `pulumi:"profile"`
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink string `pulumi:"selfLink"`
}

func LookupSSLPolicyOutput(ctx *pulumi.Context, args LookupSSLPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSSLPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSSLPolicyResult, error) {
			args := v.(LookupSSLPolicyArgs)
			r, err := LookupSSLPolicy(ctx, &args, opts...)
			var s LookupSSLPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSSLPolicyResultOutput)
}

// A collection of arguments for invoking getSSLPolicy.
type LookupSSLPolicyOutputArgs struct {
	// The name of the SSL Policy.
	//
	// ***
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupSSLPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSSLPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getSSLPolicy.
type LookupSSLPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSSLPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSSLPolicyResult)(nil)).Elem()
}

func (o LookupSSLPolicyResultOutput) ToLookupSSLPolicyResultOutput() LookupSSLPolicyResultOutput {
	return o
}

func (o LookupSSLPolicyResultOutput) ToLookupSSLPolicyResultOutputWithContext(ctx context.Context) LookupSSLPolicyResultOutput {
	return o
}

func (o LookupSSLPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSSLPolicyResult] {
	return pulumix.Output[LookupSSLPolicyResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSSLPolicyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// If the `profile` is `CUSTOM`, these are the custom encryption
// ciphers supported by the profile. If the `profile` is *not* `CUSTOM`, this
// attribute will be empty.
func (o LookupSSLPolicyResultOutput) CustomFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) []string { return v.CustomFeatures }).(pulumi.StringArrayOutput)
}

// Description of this SSL Policy.
func (o LookupSSLPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The set of enabled encryption ciphers as a result of the policy config
func (o LookupSSLPolicyResultOutput) EnabledFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) []string { return v.EnabledFeatures }).(pulumi.StringArrayOutput)
}

// Fingerprint of this resource.
func (o LookupSSLPolicyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSSLPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The minimum supported TLS version of this policy.
func (o LookupSSLPolicyResultOutput) MinTlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.MinTlsVersion }).(pulumi.StringOutput)
}

func (o LookupSSLPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Google-curated or custom profile used by this policy.
func (o LookupSSLPolicyResultOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.Profile }).(pulumi.StringOutput)
}

func (o LookupSSLPolicyResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// The URI of the created resource.
func (o LookupSSLPolicyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSLPolicyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSSLPolicyResultOutput{})
}
