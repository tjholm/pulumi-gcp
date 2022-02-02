// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new bucket in Google cloud storage service (GCS).
// Once a bucket has been created, its location can't be changed.
//
// For more information see
// [the official documentation](https://cloud.google.com/storage/docs/overview)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
//
// **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
// determined which will require enabling the compute api.
//
// ## Example Usage
// ### Creating A Private Bucket In Standard Storage, In The EU Region. Bucket Configured As Static Website And CORS Configurations
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.NewBucket(ctx, "static-site", &storage.BucketArgs{
// 			Cors: storage.BucketCorArray{
// 				&storage.BucketCorArgs{
// 					MaxAgeSeconds: pulumi.Int(3600),
// 					Methods: pulumi.StringArray{
// 						pulumi.String("GET"),
// 						pulumi.String("HEAD"),
// 						pulumi.String("PUT"),
// 						pulumi.String("POST"),
// 						pulumi.String("DELETE"),
// 					},
// 					Origins: pulumi.StringArray{
// 						pulumi.String("http://image-store.com"),
// 					},
// 					ResponseHeaders: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 				},
// 			},
// 			ForceDestroy:             pulumi.Bool(true),
// 			Location:                 pulumi.String("EU"),
// 			UniformBucketLevelAccess: pulumi.Bool(true),
// 			Website: &storage.BucketWebsiteArgs{
// 				MainPageSuffix: pulumi.String("index.html"),
// 				NotFoundPage:   pulumi.String("404.html"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Life Cycle Settings For Storage Bucket Objects
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.NewBucket(ctx, "auto-expire", &storage.BucketArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 			LifecycleRules: storage.BucketLifecycleRuleArray{
// 				&storage.BucketLifecycleRuleArgs{
// 					Action: &storage.BucketLifecycleRuleActionArgs{
// 						Type: pulumi.String("Delete"),
// 					},
// 					Condition: &storage.BucketLifecycleRuleConditionArgs{
// 						Age: pulumi.Int(3),
// 					},
// 				},
// 			},
// 			Location: pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Storage buckets can be imported using the `name` or
//
// `project/name`. If the project is not passed to the import command it will be inferred from the provider block or environment variables. If it cannot be inferred it will be queried from the Compute API (this will fail if the API is not enabled). e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucket:Bucket image-store image-store-bucket
// ```
//
// ```sh
//  $ pulumi import gcp:storage/bucket:Bucket image-store tf-test-project/image-store-bucket
// ```
//
//  `false` in state. If you've set it to `true` in config, run `terraform apply` to update the value set in state. If you delete this resource before updating the value, objects in the bucket will not be destroyed.
type Bucket struct {
	pulumi.CustomResourceState

	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  BucketCorArrayOutput `pulumi:"cors"`
	DefaultEventBasedHold pulumi.BoolPtrOutput `pulumi:"defaultEventBasedHold"`
	// The bucket's encryption configuration. Structure is documented below.
	Encryption BucketEncryptionPtrOutput `pulumi:"encryption"`
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// A map of key/value label pairs to assign to the bucket.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location pulumi.StringOutput `pulumi:"location"`
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
	Logging BucketLoggingPtrOutput `pulumi:"logging"`
	// The name of the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Prevents public access to a bucket.
	PublicAccessPrevention pulumi.StringOutput `pulumi:"publicAccessPrevention"`
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays pulumi.BoolPtrOutput `pulumi:"requesterPays"`
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy BucketRetentionPolicyPtrOutput `pulumi:"retentionPolicy"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess pulumi.BoolOutput `pulumi:"uniformBucketLevelAccess"`
	// The base URL of the bucket, in the format `gs://<bucket-name>`.
	Url pulumi.StringOutput `pulumi:"url"`
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
	Versioning BucketVersioningPtrOutput `pulumi:"versioning"`
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website BucketWebsitePtrOutput `pulumi:"website"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Bucket
	err := ctx.RegisterResource("gcp:storage/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("gcp:storage/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  []BucketCor `pulumi:"cors"`
	DefaultEventBasedHold *bool       `pulumi:"defaultEventBasedHold"`
	// The bucket's encryption configuration. Structure is documented below.
	Encryption *BucketEncryption `pulumi:"encryption"`
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A map of key/value label pairs to assign to the bucket.
	Labels map[string]string `pulumi:"labels"`
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location *string `pulumi:"location"`
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
	Logging *BucketLogging `pulumi:"logging"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Prevents public access to a bucket.
	PublicAccessPrevention *string `pulumi:"publicAccessPrevention"`
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays *bool `pulumi:"requesterPays"`
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy *BucketRetentionPolicy `pulumi:"retentionPolicy"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass *string `pulumi:"storageClass"`
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess *bool `pulumi:"uniformBucketLevelAccess"`
	// The base URL of the bucket, in the format `gs://<bucket-name>`.
	Url *string `pulumi:"url"`
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
	Versioning *BucketVersioning `pulumi:"versioning"`
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website *BucketWebsite `pulumi:"website"`
}

type BucketState struct {
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  BucketCorArrayInput
	DefaultEventBasedHold pulumi.BoolPtrInput
	// The bucket's encryption configuration. Structure is documented below.
	Encryption BucketEncryptionPtrInput
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy pulumi.BoolPtrInput
	// A map of key/value label pairs to assign to the bucket.
	Labels pulumi.StringMapInput
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules BucketLifecycleRuleArrayInput
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location pulumi.StringPtrInput
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
	Logging BucketLoggingPtrInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Prevents public access to a bucket.
	PublicAccessPrevention pulumi.StringPtrInput
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays pulumi.BoolPtrInput
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy BucketRetentionPolicyPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass pulumi.StringPtrInput
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess pulumi.BoolPtrInput
	// The base URL of the bucket, in the format `gs://<bucket-name>`.
	Url pulumi.StringPtrInput
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
	Versioning BucketVersioningPtrInput
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website BucketWebsitePtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  []BucketCor `pulumi:"cors"`
	DefaultEventBasedHold *bool       `pulumi:"defaultEventBasedHold"`
	// The bucket's encryption configuration. Structure is documented below.
	Encryption *BucketEncryption `pulumi:"encryption"`
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A map of key/value label pairs to assign to the bucket.
	Labels map[string]string `pulumi:"labels"`
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location string `pulumi:"location"`
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
	Logging *BucketLogging `pulumi:"logging"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Prevents public access to a bucket.
	PublicAccessPrevention *string `pulumi:"publicAccessPrevention"`
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays *bool `pulumi:"requesterPays"`
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy *BucketRetentionPolicy `pulumi:"retentionPolicy"`
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass *string `pulumi:"storageClass"`
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess *bool `pulumi:"uniformBucketLevelAccess"`
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
	Versioning *BucketVersioning `pulumi:"versioning"`
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website *BucketWebsite `pulumi:"website"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  BucketCorArrayInput
	DefaultEventBasedHold pulumi.BoolPtrInput
	// The bucket's encryption configuration. Structure is documented below.
	Encryption BucketEncryptionPtrInput
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy pulumi.BoolPtrInput
	// A map of key/value label pairs to assign to the bucket.
	Labels pulumi.StringMapInput
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules BucketLifecycleRuleArrayInput
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location pulumi.StringInput
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
	Logging BucketLoggingPtrInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Prevents public access to a bucket.
	PublicAccessPrevention pulumi.StringPtrInput
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays pulumi.BoolPtrInput
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy BucketRetentionPolicyPtrInput
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass pulumi.StringPtrInput
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess pulumi.BoolPtrInput
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
	Versioning BucketVersioningPtrInput
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website BucketWebsitePtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//          BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//          BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].([]*Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].(map[string]*Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketArrayInput)(nil)).Elem(), BucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMapInput)(nil)).Elem(), BucketMap{})
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}
