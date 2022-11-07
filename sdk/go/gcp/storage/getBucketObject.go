// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
//
// ## Example Usage
//
// Example picture stored within a folder.
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
//			_, err = storage.LookupBucketObject(ctx, &storage.LookupBucketObjectArgs{
//				Bucket: pulumi.StringRef("image-store"),
//				Name:   pulumi.StringRef("folder/butterfly01.jpg"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBucketObject(ctx *pulumi.Context, args *LookupBucketObjectArgs, opts ...pulumi.InvokeOption) (*LookupBucketObjectResult, error) {
	var rv LookupBucketObjectResult
	err := ctx.Invoke("gcp:storage/getBucketObject:getBucketObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketObject.
type LookupBucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket *string `pulumi:"bucket"`
	// The name of the object.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getBucketObject.
type LookupBucketObjectResult struct {
	Bucket *string `pulumi:"bucket"`
	// (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl string `pulumi:"cacheControl"`
	Content      string `pulumi:"content"`
	// (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition string `pulumi:"contentDisposition"`
	// (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding string `pulumi:"contentEncoding"`
	// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage string `pulumi:"contentLanguage"`
	// (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType string `pulumi:"contentType"`
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c              string                              `pulumi:"crc32c"`
	CustomerEncryptions []GetBucketObjectCustomerEncryption `pulumi:"customerEncryptions"`
	DetectMd5hash       string                              `pulumi:"detectMd5hash"`
	// (Computed) Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold bool `pulumi:"eventBasedHold"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	KmsKeyName string `pulumi:"kmsKeyName"`
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash string `pulumi:"md5hash"`
	// (Computed) A url reference to download this object.
	MediaLink  string            `pulumi:"mediaLink"`
	Metadata   map[string]string `pulumi:"metadata"`
	Name       *string           `pulumi:"name"`
	OutputName string            `pulumi:"outputName"`
	// (Computed) A url reference to this object.
	SelfLink string `pulumi:"selfLink"`
	Source   string `pulumi:"source"`
	// (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass string `pulumi:"storageClass"`
	// (Computed) Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold bool `pulumi:"temporaryHold"`
}

func LookupBucketObjectOutput(ctx *pulumi.Context, args LookupBucketObjectOutputArgs, opts ...pulumi.InvokeOption) LookupBucketObjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketObjectResult, error) {
			args := v.(LookupBucketObjectArgs)
			r, err := LookupBucketObject(ctx, &args, opts...)
			var s LookupBucketObjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketObjectResultOutput)
}

// A collection of arguments for invoking getBucketObject.
type LookupBucketObjectOutputArgs struct {
	// The name of the containing bucket.
	Bucket pulumi.StringPtrInput `pulumi:"bucket"`
	// The name of the object.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupBucketObjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketObjectArgs)(nil)).Elem()
}

// A collection of values returned by getBucketObject.
type LookupBucketObjectResultOutput struct{ *pulumi.OutputState }

func (LookupBucketObjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketObjectResult)(nil)).Elem()
}

func (o LookupBucketObjectResultOutput) ToLookupBucketObjectResultOutput() LookupBucketObjectResultOutput {
	return o
}

func (o LookupBucketObjectResultOutput) ToLookupBucketObjectResultOutputWithContext(ctx context.Context) LookupBucketObjectResultOutput {
	return o
}

func (o LookupBucketObjectResultOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) *string { return v.Bucket }).(pulumi.StringPtrOutput)
}

// (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
func (o LookupBucketObjectResultOutput) CacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.CacheControl }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Content }).(pulumi.StringOutput)
}

// (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
func (o LookupBucketObjectResultOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentDisposition }).(pulumi.StringOutput)
}

// (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
func (o LookupBucketObjectResultOutput) ContentEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentEncoding }).(pulumi.StringOutput)
}

// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
func (o LookupBucketObjectResultOutput) ContentLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentLanguage }).(pulumi.StringOutput)
}

// (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
func (o LookupBucketObjectResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentType }).(pulumi.StringOutput)
}

// (Computed) Base 64 CRC32 hash of the uploaded data.
func (o LookupBucketObjectResultOutput) Crc32c() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Crc32c }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) CustomerEncryptions() GetBucketObjectCustomerEncryptionArrayOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) []GetBucketObjectCustomerEncryption { return v.CustomerEncryptions }).(GetBucketObjectCustomerEncryptionArrayOutput)
}

func (o LookupBucketObjectResultOutput) DetectMd5hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.DetectMd5hash }).(pulumi.StringOutput)
}

// (Computed) Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
func (o LookupBucketObjectResultOutput) EventBasedHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) bool { return v.EventBasedHold }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBucketObjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

// (Computed) Base 64 MD5 hash of the uploaded data.
func (o LookupBucketObjectResultOutput) Md5hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Md5hash }).(pulumi.StringOutput)
}

// (Computed) A url reference to download this object.
func (o LookupBucketObjectResultOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.MediaLink }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupBucketObjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupBucketObjectResultOutput) OutputName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.OutputName }).(pulumi.StringOutput)
}

// (Computed) A url reference to this object.
func (o LookupBucketObjectResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Source }).(pulumi.StringOutput)
}

// (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
func (o LookupBucketObjectResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

// (Computed) Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
func (o LookupBucketObjectResultOutput) TemporaryHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) bool { return v.TemporaryHold }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketObjectResultOutput{})
}
