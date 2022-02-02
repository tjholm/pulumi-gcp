// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Authoritatively manages the access control list (ACL) for an object in a Google
// Cloud Storage (GCS) bucket. Removing a `storage.ObjectACL` sets the
// acl to the `private` [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl).
//
// For more information see
// [the official documentation](https://cloud.google.com/storage/docs/access-control/lists)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls).
//
// > Want fine-grained control over object ACLs? Use `storage.ObjectAccessControl` to control individual
// role entity pairs.
//
// ## Example Usage
//
// Create an object ACL with one owner and one reader.
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
// 		_, err := storage.NewBucket(ctx, "image-store", &storage.BucketArgs{
// 			Location: pulumi.String("EU"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		image, err := storage.NewBucketObject(ctx, "image", &storage.BucketObjectArgs{
// 			Bucket: image_store.Name,
// 			Source: pulumi.NewFileAsset("image1.jpg"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewObjectACL(ctx, "image-store-acl", &storage.ObjectACLArgs{
// 			Bucket: image_store.Name,
// 			Object: image.OutputName,
// 			RoleEntities: pulumi.StringArray{
// 				pulumi.String("OWNER:user-my.email@gmail.com"),
// 				pulumi.String("READER:group-mygroup"),
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
// This resource does not support import.
type ObjectACL struct {
	pulumi.CustomResourceState

	// The name of the bucket the object is stored in.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The name of the object to apply the acl to.
	Object pulumi.StringOutput `pulumi:"object"`
	// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl pulumi.StringPtrOutput `pulumi:"predefinedAcl"`
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Must be set if `predefinedAcl` is not.
	RoleEntities pulumi.StringArrayOutput `pulumi:"roleEntities"`
}

// NewObjectACL registers a new resource with the given unique name, arguments, and options.
func NewObjectACL(ctx *pulumi.Context,
	name string, args *ObjectACLArgs, opts ...pulumi.ResourceOption) (*ObjectACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Object == nil {
		return nil, errors.New("invalid value for required argument 'Object'")
	}
	var resource ObjectACL
	err := ctx.RegisterResource("gcp:storage/objectACL:ObjectACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectACL gets an existing ObjectACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectACLState, opts ...pulumi.ResourceOption) (*ObjectACL, error) {
	var resource ObjectACL
	err := ctx.ReadResource("gcp:storage/objectACL:ObjectACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectACL resources.
type objectACLState struct {
	// The name of the bucket the object is stored in.
	Bucket *string `pulumi:"bucket"`
	// The name of the object to apply the acl to.
	Object *string `pulumi:"object"`
	// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl *string `pulumi:"predefinedAcl"`
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Must be set if `predefinedAcl` is not.
	RoleEntities []string `pulumi:"roleEntities"`
}

type ObjectACLState struct {
	// The name of the bucket the object is stored in.
	Bucket pulumi.StringPtrInput
	// The name of the object to apply the acl to.
	Object pulumi.StringPtrInput
	// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl pulumi.StringPtrInput
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Must be set if `predefinedAcl` is not.
	RoleEntities pulumi.StringArrayInput
}

func (ObjectACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectACLState)(nil)).Elem()
}

type objectACLArgs struct {
	// The name of the bucket the object is stored in.
	Bucket string `pulumi:"bucket"`
	// The name of the object to apply the acl to.
	Object string `pulumi:"object"`
	// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl *string `pulumi:"predefinedAcl"`
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Must be set if `predefinedAcl` is not.
	RoleEntities []string `pulumi:"roleEntities"`
}

// The set of arguments for constructing a ObjectACL resource.
type ObjectACLArgs struct {
	// The name of the bucket the object is stored in.
	Bucket pulumi.StringInput
	// The name of the object to apply the acl to.
	Object pulumi.StringInput
	// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl pulumi.StringPtrInput
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Must be set if `predefinedAcl` is not.
	RoleEntities pulumi.StringArrayInput
}

func (ObjectACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectACLArgs)(nil)).Elem()
}

type ObjectACLInput interface {
	pulumi.Input

	ToObjectACLOutput() ObjectACLOutput
	ToObjectACLOutputWithContext(ctx context.Context) ObjectACLOutput
}

func (*ObjectACL) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectACL)(nil)).Elem()
}

func (i *ObjectACL) ToObjectACLOutput() ObjectACLOutput {
	return i.ToObjectACLOutputWithContext(context.Background())
}

func (i *ObjectACL) ToObjectACLOutputWithContext(ctx context.Context) ObjectACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectACLOutput)
}

// ObjectACLArrayInput is an input type that accepts ObjectACLArray and ObjectACLArrayOutput values.
// You can construct a concrete instance of `ObjectACLArrayInput` via:
//
//          ObjectACLArray{ ObjectACLArgs{...} }
type ObjectACLArrayInput interface {
	pulumi.Input

	ToObjectACLArrayOutput() ObjectACLArrayOutput
	ToObjectACLArrayOutputWithContext(context.Context) ObjectACLArrayOutput
}

type ObjectACLArray []ObjectACLInput

func (ObjectACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectACL)(nil)).Elem()
}

func (i ObjectACLArray) ToObjectACLArrayOutput() ObjectACLArrayOutput {
	return i.ToObjectACLArrayOutputWithContext(context.Background())
}

func (i ObjectACLArray) ToObjectACLArrayOutputWithContext(ctx context.Context) ObjectACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectACLArrayOutput)
}

// ObjectACLMapInput is an input type that accepts ObjectACLMap and ObjectACLMapOutput values.
// You can construct a concrete instance of `ObjectACLMapInput` via:
//
//          ObjectACLMap{ "key": ObjectACLArgs{...} }
type ObjectACLMapInput interface {
	pulumi.Input

	ToObjectACLMapOutput() ObjectACLMapOutput
	ToObjectACLMapOutputWithContext(context.Context) ObjectACLMapOutput
}

type ObjectACLMap map[string]ObjectACLInput

func (ObjectACLMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectACL)(nil)).Elem()
}

func (i ObjectACLMap) ToObjectACLMapOutput() ObjectACLMapOutput {
	return i.ToObjectACLMapOutputWithContext(context.Background())
}

func (i ObjectACLMap) ToObjectACLMapOutputWithContext(ctx context.Context) ObjectACLMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectACLMapOutput)
}

type ObjectACLOutput struct{ *pulumi.OutputState }

func (ObjectACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectACL)(nil)).Elem()
}

func (o ObjectACLOutput) ToObjectACLOutput() ObjectACLOutput {
	return o
}

func (o ObjectACLOutput) ToObjectACLOutputWithContext(ctx context.Context) ObjectACLOutput {
	return o
}

type ObjectACLArrayOutput struct{ *pulumi.OutputState }

func (ObjectACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectACL)(nil)).Elem()
}

func (o ObjectACLArrayOutput) ToObjectACLArrayOutput() ObjectACLArrayOutput {
	return o
}

func (o ObjectACLArrayOutput) ToObjectACLArrayOutputWithContext(ctx context.Context) ObjectACLArrayOutput {
	return o
}

func (o ObjectACLArrayOutput) Index(i pulumi.IntInput) ObjectACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectACL {
		return vs[0].([]*ObjectACL)[vs[1].(int)]
	}).(ObjectACLOutput)
}

type ObjectACLMapOutput struct{ *pulumi.OutputState }

func (ObjectACLMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectACL)(nil)).Elem()
}

func (o ObjectACLMapOutput) ToObjectACLMapOutput() ObjectACLMapOutput {
	return o
}

func (o ObjectACLMapOutput) ToObjectACLMapOutputWithContext(ctx context.Context) ObjectACLMapOutput {
	return o
}

func (o ObjectACLMapOutput) MapIndex(k pulumi.StringInput) ObjectACLOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectACL {
		return vs[0].(map[string]*ObjectACL)[vs[1].(string)]
	}).(ObjectACLOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectACLInput)(nil)).Elem(), &ObjectACL{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectACLArrayInput)(nil)).Elem(), ObjectACLArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectACLMapInput)(nil)).Elem(), ObjectACLMap{})
	pulumi.RegisterOutputType(ObjectACLOutput{})
	pulumi.RegisterOutputType(ObjectACLArrayOutput{})
	pulumi.RegisterOutputType(ObjectACLMapOutput{})
}
