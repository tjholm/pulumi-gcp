// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing bucket in Google Cloud Storage service (GCS).
// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#buckets)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = storage.LookupBucket(ctx, &storage.LookupBucketArgs{
//				Name: "my-bucket",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBucket(ctx *pulumi.Context, args *LookupBucketArgs, opts ...pulumi.InvokeOption) (*LookupBucketResult, error) {
	var rv LookupBucketResult
	err := ctx.Invoke("gcp:storage/getBucket:getBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucket.
type LookupBucketArgs struct {
	// The name of the bucket.
	Name string `pulumi:"name"`
}

// A collection of values returned by getBucket.
type LookupBucketResult struct {
	Cors                   []GetBucketCor                   `pulumi:"cors"`
	CustomPlacementConfigs []GetBucketCustomPlacementConfig `pulumi:"customPlacementConfigs"`
	DefaultEventBasedHold  bool                             `pulumi:"defaultEventBasedHold"`
	Encryptions            []GetBucketEncryption            `pulumi:"encryptions"`
	ForceDestroy           bool                             `pulumi:"forceDestroy"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string                     `pulumi:"id"`
	Labels                   map[string]string          `pulumi:"labels"`
	LifecycleRules           []GetBucketLifecycleRule   `pulumi:"lifecycleRules"`
	Location                 string                     `pulumi:"location"`
	Loggings                 []GetBucketLogging         `pulumi:"loggings"`
	Name                     string                     `pulumi:"name"`
	Project                  string                     `pulumi:"project"`
	PublicAccessPrevention   string                     `pulumi:"publicAccessPrevention"`
	RequesterPays            bool                       `pulumi:"requesterPays"`
	RetentionPolicies        []GetBucketRetentionPolicy `pulumi:"retentionPolicies"`
	SelfLink                 string                     `pulumi:"selfLink"`
	StorageClass             string                     `pulumi:"storageClass"`
	UniformBucketLevelAccess bool                       `pulumi:"uniformBucketLevelAccess"`
	Url                      string                     `pulumi:"url"`
	Versionings              []GetBucketVersioning      `pulumi:"versionings"`
	Websites                 []GetBucketWebsite         `pulumi:"websites"`
}

func LookupBucketOutput(ctx *pulumi.Context, args LookupBucketOutputArgs, opts ...pulumi.InvokeOption) LookupBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketResult, error) {
			args := v.(LookupBucketArgs)
			r, err := LookupBucket(ctx, &args, opts...)
			var s LookupBucketResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketResultOutput)
}

// A collection of arguments for invoking getBucket.
type LookupBucketOutputArgs struct {
	// The name of the bucket.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketArgs)(nil)).Elem()
}

// A collection of values returned by getBucket.
type LookupBucketResultOutput struct{ *pulumi.OutputState }

func (LookupBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketResult)(nil)).Elem()
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutput() LookupBucketResultOutput {
	return o
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutputWithContext(ctx context.Context) LookupBucketResultOutput {
	return o
}

func (o LookupBucketResultOutput) Cors() GetBucketCorArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketCor { return v.Cors }).(GetBucketCorArrayOutput)
}

func (o LookupBucketResultOutput) CustomPlacementConfigs() GetBucketCustomPlacementConfigArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketCustomPlacementConfig { return v.CustomPlacementConfigs }).(GetBucketCustomPlacementConfigArrayOutput)
}

func (o LookupBucketResultOutput) DefaultEventBasedHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketResult) bool { return v.DefaultEventBasedHold }).(pulumi.BoolOutput)
}

func (o LookupBucketResultOutput) Encryptions() GetBucketEncryptionArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketEncryption { return v.Encryptions }).(GetBucketEncryptionArrayOutput)
}

func (o LookupBucketResultOutput) ForceDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketResult) bool { return v.ForceDestroy }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBucketResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBucketResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupBucketResultOutput) LifecycleRules() GetBucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketLifecycleRule { return v.LifecycleRules }).(GetBucketLifecycleRuleArrayOutput)
}

func (o LookupBucketResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) Loggings() GetBucketLoggingArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketLogging { return v.Loggings }).(GetBucketLoggingArrayOutput)
}

func (o LookupBucketResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) PublicAccessPrevention() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.PublicAccessPrevention }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) RequesterPays() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketResult) bool { return v.RequesterPays }).(pulumi.BoolOutput)
}

func (o LookupBucketResultOutput) RetentionPolicies() GetBucketRetentionPolicyArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketRetentionPolicy { return v.RetentionPolicies }).(GetBucketRetentionPolicyArrayOutput)
}

func (o LookupBucketResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) UniformBucketLevelAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketResult) bool { return v.UniformBucketLevelAccess }).(pulumi.BoolOutput)
}

func (o LookupBucketResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketResult) string { return v.Url }).(pulumi.StringOutput)
}

func (o LookupBucketResultOutput) Versionings() GetBucketVersioningArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketVersioning { return v.Versionings }).(GetBucketVersioningArrayOutput)
}

func (o LookupBucketResultOutput) Websites() GetBucketWebsiteArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []GetBucketWebsite { return v.Websites }).(GetBucketWebsiteArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketResultOutput{})
}
