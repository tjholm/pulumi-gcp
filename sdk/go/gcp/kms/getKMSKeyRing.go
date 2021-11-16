// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides access to Google Cloud Platform KMS KeyRing. For more information see
// [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_ring)
// and
// [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings).
//
// A KeyRing is a grouping of CryptoKeys for organizational purposes. A KeyRing belongs to a Google Cloud Platform Project
// and resides in a specific location.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.GetKMSKeyRing(ctx, &kms.GetKMSKeyRingArgs{
// 			Location: "us-central1",
// 			Name:     "my-key-ring",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetKMSKeyRing(ctx *pulumi.Context, args *GetKMSKeyRingArgs, opts ...pulumi.InvokeOption) (*GetKMSKeyRingResult, error) {
	var rv GetKMSKeyRingResult
	err := ctx.Invoke("gcp:kms/getKMSKeyRing:getKMSKeyRing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKMSKeyRing.
type GetKMSKeyRingArgs struct {
	// The Google Cloud Platform location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location string `pulumi:"location"`
	// The KeyRing's name.
	// A KeyRing name must exist within the provided location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getKMSKeyRing.
type GetKMSKeyRingResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Project  *string `pulumi:"project"`
}

func GetKMSKeyRingOutput(ctx *pulumi.Context, args GetKMSKeyRingOutputArgs, opts ...pulumi.InvokeOption) GetKMSKeyRingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKMSKeyRingResult, error) {
			args := v.(GetKMSKeyRingArgs)
			r, err := GetKMSKeyRing(ctx, &args, opts...)
			return *r, err
		}).(GetKMSKeyRingResultOutput)
}

// A collection of arguments for invoking getKMSKeyRing.
type GetKMSKeyRingOutputArgs struct {
	// The Google Cloud Platform location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location pulumi.StringInput `pulumi:"location"`
	// The KeyRing's name.
	// A KeyRing name must exist within the provided location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	Name pulumi.StringInput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetKMSKeyRingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKMSKeyRingArgs)(nil)).Elem()
}

// A collection of values returned by getKMSKeyRing.
type GetKMSKeyRingResultOutput struct{ *pulumi.OutputState }

func (GetKMSKeyRingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKMSKeyRingResult)(nil)).Elem()
}

func (o GetKMSKeyRingResultOutput) ToGetKMSKeyRingResultOutput() GetKMSKeyRingResultOutput {
	return o
}

func (o GetKMSKeyRingResultOutput) ToGetKMSKeyRingResultOutputWithContext(ctx context.Context) GetKMSKeyRingResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetKMSKeyRingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKMSKeyRingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKMSKeyRingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetKMSKeyRingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetKMSKeyRingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetKMSKeyRingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetKMSKeyRingResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKMSKeyRingResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKMSKeyRingResultOutput{})
}
