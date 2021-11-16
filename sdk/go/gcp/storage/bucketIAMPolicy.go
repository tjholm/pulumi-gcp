// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Storage Bucket. Each of these resources serves a different use case:
//
// * `storage.BucketIAMPolicy`: Authoritative. Sets the IAM policy for the bucket and replaces any existing policy already attached.
// * `storage.BucketIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the bucket are preserved.
// * `storage.BucketIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the bucket are preserved.
//
// > **Note:** `storage.BucketIAMPolicy` **cannot** be used in conjunction with `storage.BucketIAMBinding` and `storage.BucketIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `storage.BucketIAMBinding` resources **can be** used in conjunction with `storage.BucketIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_storage\_bucket\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/storage.admin",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucketIAMPolicy(ctx, "policy", &storage.BucketIAMPolicyArgs{
// 			Bucket:     pulumi.Any(google_storage_bucket.Default.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/storage.admin",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucketIAMPolicy(ctx, "policy", &storage.BucketIAMPolicyArgs{
// 			Bucket:     pulumi.Any(google_storage_bucket.Default.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_storage\_bucket\_iam\_binding
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
// 		_, err := storage.NewBucketIAMBinding(ctx, "binding", &storage.BucketIAMBindingArgs{
// 			Bucket: pulumi.Any(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
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
// 		_, err := storage.NewBucketIAMBinding(ctx, "binding", &storage.BucketIAMBindingArgs{
// 			Bucket: pulumi.Any(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &storage.BucketIAMBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_storage\_bucket\_iam\_member
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
// 		_, err := storage.NewBucketIAMMember(ctx, "member", &storage.BucketIAMMemberArgs{
// 			Bucket: pulumi.Any(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Member: pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
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
// 		_, err := storage.NewBucketIAMMember(ctx, "member", &storage.BucketIAMMemberArgs{
// 			Bucket: pulumi.Any(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Member: pulumi.String("user:jane@example.com"),
// 			Condition: &storage.BucketIAMMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
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
// For all import syntaxes, the "resource in question" can take any of the following forms* b/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Storage bucket IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucketIAMPolicy:BucketIAMPolicy editor "b/{{bucket}} roles/storage.objectViewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucketIAMPolicy:BucketIAMPolicy editor "b/{{bucket}} roles/storage.objectViewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucketIAMPolicy:BucketIAMPolicy editor b/{{bucket}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BucketIAMPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewBucketIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketIAMPolicy(ctx *pulumi.Context,
	name string, args *BucketIAMPolicyArgs, opts ...pulumi.ResourceOption) (*BucketIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource BucketIAMPolicy
	err := ctx.RegisterResource("gcp:storage/bucketIAMPolicy:BucketIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIAMPolicy gets an existing BucketIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIAMPolicyState, opts ...pulumi.ResourceOption) (*BucketIAMPolicy, error) {
	var resource BucketIAMPolicy
	err := ctx.ReadResource("gcp:storage/bucketIAMPolicy:BucketIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIAMPolicy resources.
type bucketIAMPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket *string `pulumi:"bucket"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type BucketIAMPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (BucketIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMPolicyState)(nil)).Elem()
}

type bucketIAMPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket string `pulumi:"bucket"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a BucketIAMPolicy resource.
type BucketIAMPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (BucketIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMPolicyArgs)(nil)).Elem()
}

type BucketIAMPolicyInput interface {
	pulumi.Input

	ToBucketIAMPolicyOutput() BucketIAMPolicyOutput
	ToBucketIAMPolicyOutputWithContext(ctx context.Context) BucketIAMPolicyOutput
}

func (*BucketIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketIAMPolicy)(nil))
}

func (i *BucketIAMPolicy) ToBucketIAMPolicyOutput() BucketIAMPolicyOutput {
	return i.ToBucketIAMPolicyOutputWithContext(context.Background())
}

func (i *BucketIAMPolicy) ToBucketIAMPolicyOutputWithContext(ctx context.Context) BucketIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIAMPolicyOutput)
}

func (i *BucketIAMPolicy) ToBucketIAMPolicyPtrOutput() BucketIAMPolicyPtrOutput {
	return i.ToBucketIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *BucketIAMPolicy) ToBucketIAMPolicyPtrOutputWithContext(ctx context.Context) BucketIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIAMPolicyPtrOutput)
}

type BucketIAMPolicyPtrInput interface {
	pulumi.Input

	ToBucketIAMPolicyPtrOutput() BucketIAMPolicyPtrOutput
	ToBucketIAMPolicyPtrOutputWithContext(ctx context.Context) BucketIAMPolicyPtrOutput
}

type bucketIAMPolicyPtrType BucketIAMPolicyArgs

func (*bucketIAMPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIAMPolicy)(nil))
}

func (i *bucketIAMPolicyPtrType) ToBucketIAMPolicyPtrOutput() BucketIAMPolicyPtrOutput {
	return i.ToBucketIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *bucketIAMPolicyPtrType) ToBucketIAMPolicyPtrOutputWithContext(ctx context.Context) BucketIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIAMPolicyPtrOutput)
}

// BucketIAMPolicyArrayInput is an input type that accepts BucketIAMPolicyArray and BucketIAMPolicyArrayOutput values.
// You can construct a concrete instance of `BucketIAMPolicyArrayInput` via:
//
//          BucketIAMPolicyArray{ BucketIAMPolicyArgs{...} }
type BucketIAMPolicyArrayInput interface {
	pulumi.Input

	ToBucketIAMPolicyArrayOutput() BucketIAMPolicyArrayOutput
	ToBucketIAMPolicyArrayOutputWithContext(context.Context) BucketIAMPolicyArrayOutput
}

type BucketIAMPolicyArray []BucketIAMPolicyInput

func (BucketIAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketIAMPolicy)(nil)).Elem()
}

func (i BucketIAMPolicyArray) ToBucketIAMPolicyArrayOutput() BucketIAMPolicyArrayOutput {
	return i.ToBucketIAMPolicyArrayOutputWithContext(context.Background())
}

func (i BucketIAMPolicyArray) ToBucketIAMPolicyArrayOutputWithContext(ctx context.Context) BucketIAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIAMPolicyArrayOutput)
}

// BucketIAMPolicyMapInput is an input type that accepts BucketIAMPolicyMap and BucketIAMPolicyMapOutput values.
// You can construct a concrete instance of `BucketIAMPolicyMapInput` via:
//
//          BucketIAMPolicyMap{ "key": BucketIAMPolicyArgs{...} }
type BucketIAMPolicyMapInput interface {
	pulumi.Input

	ToBucketIAMPolicyMapOutput() BucketIAMPolicyMapOutput
	ToBucketIAMPolicyMapOutputWithContext(context.Context) BucketIAMPolicyMapOutput
}

type BucketIAMPolicyMap map[string]BucketIAMPolicyInput

func (BucketIAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketIAMPolicy)(nil)).Elem()
}

func (i BucketIAMPolicyMap) ToBucketIAMPolicyMapOutput() BucketIAMPolicyMapOutput {
	return i.ToBucketIAMPolicyMapOutputWithContext(context.Background())
}

func (i BucketIAMPolicyMap) ToBucketIAMPolicyMapOutputWithContext(ctx context.Context) BucketIAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIAMPolicyMapOutput)
}

type BucketIAMPolicyOutput struct{ *pulumi.OutputState }

func (BucketIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketIAMPolicy)(nil))
}

func (o BucketIAMPolicyOutput) ToBucketIAMPolicyOutput() BucketIAMPolicyOutput {
	return o
}

func (o BucketIAMPolicyOutput) ToBucketIAMPolicyOutputWithContext(ctx context.Context) BucketIAMPolicyOutput {
	return o
}

func (o BucketIAMPolicyOutput) ToBucketIAMPolicyPtrOutput() BucketIAMPolicyPtrOutput {
	return o.ToBucketIAMPolicyPtrOutputWithContext(context.Background())
}

func (o BucketIAMPolicyOutput) ToBucketIAMPolicyPtrOutputWithContext(ctx context.Context) BucketIAMPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BucketIAMPolicy) *BucketIAMPolicy {
		return &v
	}).(BucketIAMPolicyPtrOutput)
}

type BucketIAMPolicyPtrOutput struct{ *pulumi.OutputState }

func (BucketIAMPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIAMPolicy)(nil))
}

func (o BucketIAMPolicyPtrOutput) ToBucketIAMPolicyPtrOutput() BucketIAMPolicyPtrOutput {
	return o
}

func (o BucketIAMPolicyPtrOutput) ToBucketIAMPolicyPtrOutputWithContext(ctx context.Context) BucketIAMPolicyPtrOutput {
	return o
}

func (o BucketIAMPolicyPtrOutput) Elem() BucketIAMPolicyOutput {
	return o.ApplyT(func(v *BucketIAMPolicy) BucketIAMPolicy {
		if v != nil {
			return *v
		}
		var ret BucketIAMPolicy
		return ret
	}).(BucketIAMPolicyOutput)
}

type BucketIAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (BucketIAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketIAMPolicy)(nil))
}

func (o BucketIAMPolicyArrayOutput) ToBucketIAMPolicyArrayOutput() BucketIAMPolicyArrayOutput {
	return o
}

func (o BucketIAMPolicyArrayOutput) ToBucketIAMPolicyArrayOutputWithContext(ctx context.Context) BucketIAMPolicyArrayOutput {
	return o
}

func (o BucketIAMPolicyArrayOutput) Index(i pulumi.IntInput) BucketIAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BucketIAMPolicy {
		return vs[0].([]BucketIAMPolicy)[vs[1].(int)]
	}).(BucketIAMPolicyOutput)
}

type BucketIAMPolicyMapOutput struct{ *pulumi.OutputState }

func (BucketIAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BucketIAMPolicy)(nil))
}

func (o BucketIAMPolicyMapOutput) ToBucketIAMPolicyMapOutput() BucketIAMPolicyMapOutput {
	return o
}

func (o BucketIAMPolicyMapOutput) ToBucketIAMPolicyMapOutputWithContext(ctx context.Context) BucketIAMPolicyMapOutput {
	return o
}

func (o BucketIAMPolicyMapOutput) MapIndex(k pulumi.StringInput) BucketIAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BucketIAMPolicy {
		return vs[0].(map[string]BucketIAMPolicy)[vs[1].(string)]
	}).(BucketIAMPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketIAMPolicyOutput{})
	pulumi.RegisterOutputType(BucketIAMPolicyPtrOutput{})
	pulumi.RegisterOutputType(BucketIAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(BucketIAMPolicyMapOutput{})
}
