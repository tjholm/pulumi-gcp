// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S)
// load balancing.
//
// An HTTP(S) load balancer can direct traffic to specified URLs to a
// backend bucket rather than a backend service. It can send requests for
// static content to a Cloud Storage bucket and requests for dynamic content
// to a virtual machine instance.
//
// To get more information about BackendBucket, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendBuckets)
// * How-to Guides
//   - [Using a Cloud Storage bucket as a load balancer backend](https://cloud.google.com/compute/docs/load-balancing/http/backend-bucket)
//
// ## Example Usage
// ### Backend Bucket Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			imageBucket, err := storage.NewBucket(ctx, "imageBucket", &storage.BucketArgs{
//				Location: pulumi.String("EU"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewBackendBucket(ctx, "imageBackend", &compute.BackendBucketArgs{
//				Description: pulumi.String("Contains beautiful images"),
//				BucketName:  imageBucket.Name,
//				EnableCdn:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Backend Bucket Security Policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			imageBackendBucket, err := storage.NewBucket(ctx, "imageBackendBucket", &storage.BucketArgs{
//				Location: pulumi.String("EU"),
//			})
//			if err != nil {
//				return err
//			}
//			policy, err := compute.NewSecurityPolicy(ctx, "policy", &compute.SecurityPolicyArgs{
//				Description: pulumi.String("basic security policy"),
//				Type:        pulumi.String("CLOUD_ARMOR_EDGE"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewBackendBucket(ctx, "imageBackendBackendBucket", &compute.BackendBucketArgs{
//				Description:        pulumi.String("Contains beautiful images"),
//				BucketName:         imageBackendBucket.Name,
//				EnableCdn:          pulumi.Bool(true),
//				EdgeSecurityPolicy: policy.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Backend Bucket Query String Whitelist
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			imageBucket, err := storage.NewBucket(ctx, "imageBucket", &storage.BucketArgs{
//				Location: pulumi.String("EU"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewBackendBucket(ctx, "imageBackend", &compute.BackendBucketArgs{
//				Description: pulumi.String("Contains beautiful images"),
//				BucketName:  imageBucket.Name,
//				EnableCdn:   pulumi.Bool(true),
//				CdnPolicy: &compute.BackendBucketCdnPolicyArgs{
//					CacheKeyPolicy: &compute.BackendBucketCdnPolicyCacheKeyPolicyArgs{
//						QueryStringWhitelists: pulumi.StringArray{
//							pulumi.String("image-version"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Backend Bucket Include Http Headers
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			imageBucket, err := storage.NewBucket(ctx, "imageBucket", &storage.BucketArgs{
//				Location: pulumi.String("EU"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewBackendBucket(ctx, "imageBackend", &compute.BackendBucketArgs{
//				Description: pulumi.String("Contains beautiful images"),
//				BucketName:  imageBucket.Name,
//				EnableCdn:   pulumi.Bool(true),
//				CdnPolicy: &compute.BackendBucketCdnPolicyArgs{
//					CacheKeyPolicy: &compute.BackendBucketCdnPolicyCacheKeyPolicyArgs{
//						IncludeHttpHeaders: pulumi.StringArray{
//							pulumi.String("X-My-Header-Field"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # BackendBucket can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucket:BackendBucket default projects/{{project}}/global/backendBuckets/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucket:BackendBucket default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucket:BackendBucket default {{name}}
//
// ```
type BackendBucket struct {
	pulumi.CustomResourceState

	// Cloud Storage bucket name.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	CdnPolicy BackendBucketCdnPolicyOutput `pulumi:"cdnPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders pulumi.StringArrayOutput `pulumi:"customResponseHeaders"`
	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy pulumi.StringPtrOutput `pulumi:"edgeSecurityPolicy"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn pulumi.BoolPtrOutput `pulumi:"enableCdn"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewBackendBucket registers a new resource with the given unique name, arguments, and options.
func NewBackendBucket(ctx *pulumi.Context,
	name string, args *BackendBucketArgs, opts ...pulumi.ResourceOption) (*BackendBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	var resource BackendBucket
	err := ctx.RegisterResource("gcp:compute/backendBucket:BackendBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendBucket gets an existing BackendBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendBucketState, opts ...pulumi.ResourceOption) (*BackendBucket, error) {
	var resource BackendBucket
	err := ctx.ReadResource("gcp:compute/backendBucket:BackendBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendBucket resources.
type backendBucketState struct {
	// Cloud Storage bucket name.
	BucketName *string `pulumi:"bucketName"`
	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	CdnPolicy *BackendBucketCdnPolicy `pulumi:"cdnPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders []string `pulumi:"customResponseHeaders"`
	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description *string `pulumi:"description"`
	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy *string `pulumi:"edgeSecurityPolicy"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn *bool `pulumi:"enableCdn"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type BackendBucketState struct {
	// Cloud Storage bucket name.
	BucketName pulumi.StringPtrInput
	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	CdnPolicy BackendBucketCdnPolicyPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders pulumi.StringArrayInput
	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description pulumi.StringPtrInput
	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy pulumi.StringPtrInput
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (BackendBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketState)(nil)).Elem()
}

type backendBucketArgs struct {
	// Cloud Storage bucket name.
	BucketName string `pulumi:"bucketName"`
	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	CdnPolicy *BackendBucketCdnPolicy `pulumi:"cdnPolicy"`
	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders []string `pulumi:"customResponseHeaders"`
	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description *string `pulumi:"description"`
	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy *string `pulumi:"edgeSecurityPolicy"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn *bool `pulumi:"enableCdn"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackendBucket resource.
type BackendBucketArgs struct {
	// Cloud Storage bucket name.
	BucketName pulumi.StringInput
	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	CdnPolicy BackendBucketCdnPolicyPtrInput
	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders pulumi.StringArrayInput
	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description pulumi.StringPtrInput
	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy pulumi.StringPtrInput
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketArgs)(nil)).Elem()
}

type BackendBucketInput interface {
	pulumi.Input

	ToBackendBucketOutput() BackendBucketOutput
	ToBackendBucketOutputWithContext(ctx context.Context) BackendBucketOutput
}

func (*BackendBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucket)(nil)).Elem()
}

func (i *BackendBucket) ToBackendBucketOutput() BackendBucketOutput {
	return i.ToBackendBucketOutputWithContext(context.Background())
}

func (i *BackendBucket) ToBackendBucketOutputWithContext(ctx context.Context) BackendBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketOutput)
}

// BackendBucketArrayInput is an input type that accepts BackendBucketArray and BackendBucketArrayOutput values.
// You can construct a concrete instance of `BackendBucketArrayInput` via:
//
//	BackendBucketArray{ BackendBucketArgs{...} }
type BackendBucketArrayInput interface {
	pulumi.Input

	ToBackendBucketArrayOutput() BackendBucketArrayOutput
	ToBackendBucketArrayOutputWithContext(context.Context) BackendBucketArrayOutput
}

type BackendBucketArray []BackendBucketInput

func (BackendBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucket)(nil)).Elem()
}

func (i BackendBucketArray) ToBackendBucketArrayOutput() BackendBucketArrayOutput {
	return i.ToBackendBucketArrayOutputWithContext(context.Background())
}

func (i BackendBucketArray) ToBackendBucketArrayOutputWithContext(ctx context.Context) BackendBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketArrayOutput)
}

// BackendBucketMapInput is an input type that accepts BackendBucketMap and BackendBucketMapOutput values.
// You can construct a concrete instance of `BackendBucketMapInput` via:
//
//	BackendBucketMap{ "key": BackendBucketArgs{...} }
type BackendBucketMapInput interface {
	pulumi.Input

	ToBackendBucketMapOutput() BackendBucketMapOutput
	ToBackendBucketMapOutputWithContext(context.Context) BackendBucketMapOutput
}

type BackendBucketMap map[string]BackendBucketInput

func (BackendBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucket)(nil)).Elem()
}

func (i BackendBucketMap) ToBackendBucketMapOutput() BackendBucketMapOutput {
	return i.ToBackendBucketMapOutputWithContext(context.Background())
}

func (i BackendBucketMap) ToBackendBucketMapOutputWithContext(ctx context.Context) BackendBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketMapOutput)
}

type BackendBucketOutput struct{ *pulumi.OutputState }

func (BackendBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucket)(nil)).Elem()
}

func (o BackendBucketOutput) ToBackendBucketOutput() BackendBucketOutput {
	return o
}

func (o BackendBucketOutput) ToBackendBucketOutputWithContext(ctx context.Context) BackendBucketOutput {
	return o
}

// Cloud Storage bucket name.
func (o BackendBucketOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// Cloud CDN configuration for this Backend Bucket.
// Structure is documented below.
func (o BackendBucketOutput) CdnPolicy() BackendBucketCdnPolicyOutput {
	return o.ApplyT(func(v *BackendBucket) BackendBucketCdnPolicyOutput { return v.CdnPolicy }).(BackendBucketCdnPolicyOutput)
}

// Creation timestamp in RFC3339 text format.
func (o BackendBucketOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Headers that the HTTP/S load balancer should add to proxied responses.
func (o BackendBucketOutput) CustomResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringArrayOutput { return v.CustomResponseHeaders }).(pulumi.StringArrayOutput)
}

// An optional textual description of the resource; provided by the
// client when the resource is created.
func (o BackendBucketOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The security policy associated with this backend bucket.
func (o BackendBucketOutput) EdgeSecurityPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringPtrOutput { return v.EdgeSecurityPolicy }).(pulumi.StringPtrOutput)
}

// If true, enable Cloud CDN for this BackendBucket.
func (o BackendBucketOutput) EnableCdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.BoolPtrOutput { return v.EnableCdn }).(pulumi.BoolPtrOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035.  Specifically, the name must be 1-63 characters long and
// match the regular expression `a-z?` which means
// the first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the
// last character, which cannot be a dash.
func (o BackendBucketOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o BackendBucketOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o BackendBucketOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucket) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

type BackendBucketArrayOutput struct{ *pulumi.OutputState }

func (BackendBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucket)(nil)).Elem()
}

func (o BackendBucketArrayOutput) ToBackendBucketArrayOutput() BackendBucketArrayOutput {
	return o
}

func (o BackendBucketArrayOutput) ToBackendBucketArrayOutputWithContext(ctx context.Context) BackendBucketArrayOutput {
	return o
}

func (o BackendBucketArrayOutput) Index(i pulumi.IntInput) BackendBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendBucket {
		return vs[0].([]*BackendBucket)[vs[1].(int)]
	}).(BackendBucketOutput)
}

type BackendBucketMapOutput struct{ *pulumi.OutputState }

func (BackendBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucket)(nil)).Elem()
}

func (o BackendBucketMapOutput) ToBackendBucketMapOutput() BackendBucketMapOutput {
	return o
}

func (o BackendBucketMapOutput) ToBackendBucketMapOutputWithContext(ctx context.Context) BackendBucketMapOutput {
	return o
}

func (o BackendBucketMapOutput) MapIndex(k pulumi.StringInput) BackendBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendBucket {
		return vs[0].(map[string]*BackendBucket)[vs[1].(string)]
	}).(BackendBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketInput)(nil)).Elem(), &BackendBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketArrayInput)(nil)).Elem(), BackendBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketMapInput)(nil)).Elem(), BackendBucketMap{})
	pulumi.RegisterOutputType(BackendBucketOutput{})
	pulumi.RegisterOutputType(BackendBucketArrayOutput{})
	pulumi.RegisterOutputType(BackendBucketMapOutput{})
}
