// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing object content inside an existing bucket in Google Cloud Storage service (GCS).
// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
//
// > **Warning:** The object content will be saved in the state, and visiable to everyone who has access to the state file.
//
// ## Example Usage
//
// Example file object  stored within a folder.
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
//			key, err := storage.GetBucketObjectContent(ctx, &storage.GetBucketObjectContentArgs{
//				Name:   "encryptedkey",
//				Bucket: "keystore",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("encrypted", key.Content)
//			return nil
//		})
//	}
//
// ```
func GetBucketObjectContent(ctx *pulumi.Context, args *GetBucketObjectContentArgs, opts ...pulumi.InvokeOption) (*GetBucketObjectContentResult, error) {
	var rv GetBucketObjectContentResult
	err := ctx.Invoke("gcp:storage/getBucketObjectContent:getBucketObjectContent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketObjectContent.
type GetBucketObjectContentArgs struct {
	// The name of the containing bucket.
	Bucket string `pulumi:"bucket"`
	// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.
	Content *string `pulumi:"content"`
	// The name of the object.
	Name string `pulumi:"name"`
}

// A collection of values returned by getBucketObjectContent.
type GetBucketObjectContentResult struct {
	Bucket       string `pulumi:"bucket"`
	CacheControl string `pulumi:"cacheControl"`
	// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.
	Content             *string                                    `pulumi:"content"`
	ContentDisposition  string                                     `pulumi:"contentDisposition"`
	ContentEncoding     string                                     `pulumi:"contentEncoding"`
	ContentLanguage     string                                     `pulumi:"contentLanguage"`
	ContentType         string                                     `pulumi:"contentType"`
	Crc32c              string                                     `pulumi:"crc32c"`
	CustomerEncryptions []GetBucketObjectContentCustomerEncryption `pulumi:"customerEncryptions"`
	DetectMd5hash       string                                     `pulumi:"detectMd5hash"`
	EventBasedHold      bool                                       `pulumi:"eventBasedHold"`
	// The provider-assigned unique ID for this managed resource.
	Id            string            `pulumi:"id"`
	KmsKeyName    string            `pulumi:"kmsKeyName"`
	Md5hash       string            `pulumi:"md5hash"`
	MediaLink     string            `pulumi:"mediaLink"`
	Metadata      map[string]string `pulumi:"metadata"`
	Name          string            `pulumi:"name"`
	OutputName    string            `pulumi:"outputName"`
	SelfLink      string            `pulumi:"selfLink"`
	Source        string            `pulumi:"source"`
	StorageClass  string            `pulumi:"storageClass"`
	TemporaryHold bool              `pulumi:"temporaryHold"`
}

func GetBucketObjectContentOutput(ctx *pulumi.Context, args GetBucketObjectContentOutputArgs, opts ...pulumi.InvokeOption) GetBucketObjectContentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBucketObjectContentResult, error) {
			args := v.(GetBucketObjectContentArgs)
			r, err := GetBucketObjectContent(ctx, &args, opts...)
			var s GetBucketObjectContentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBucketObjectContentResultOutput)
}

// A collection of arguments for invoking getBucketObjectContent.
type GetBucketObjectContentOutputArgs struct {
	// The name of the containing bucket.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.
	Content pulumi.StringPtrInput `pulumi:"content"`
	// The name of the object.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetBucketObjectContentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketObjectContentArgs)(nil)).Elem()
}

// A collection of values returned by getBucketObjectContent.
type GetBucketObjectContentResultOutput struct{ *pulumi.OutputState }

func (GetBucketObjectContentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketObjectContentResult)(nil)).Elem()
}

func (o GetBucketObjectContentResultOutput) ToGetBucketObjectContentResultOutput() GetBucketObjectContentResultOutput {
	return o
}

func (o GetBucketObjectContentResultOutput) ToGetBucketObjectContentResultOutputWithContext(ctx context.Context) GetBucketObjectContentResultOutput {
	return o
}

func (o GetBucketObjectContentResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.Bucket }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) CacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.CacheControl }).(pulumi.StringOutput)
}

// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.
func (o GetBucketObjectContentResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o GetBucketObjectContentResultOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.ContentDisposition }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) ContentEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.ContentEncoding }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) ContentLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.ContentLanguage }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) Crc32c() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.Crc32c }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) CustomerEncryptions() GetBucketObjectContentCustomerEncryptionArrayOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) []GetBucketObjectContentCustomerEncryption {
		return v.CustomerEncryptions
	}).(GetBucketObjectContentCustomerEncryptionArrayOutput)
}

func (o GetBucketObjectContentResultOutput) DetectMd5hash() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.DetectMd5hash }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) EventBasedHold() pulumi.BoolOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) bool { return v.EventBasedHold }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBucketObjectContentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) Md5hash() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.Md5hash }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.MediaLink }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o GetBucketObjectContentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) OutputName() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.OutputName }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.Source }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

func (o GetBucketObjectContentResultOutput) TemporaryHold() pulumi.BoolOutput {
	return o.ApplyT(func(v GetBucketObjectContentResult) bool { return v.TemporaryHold }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBucketObjectContentResultOutput{})
}
