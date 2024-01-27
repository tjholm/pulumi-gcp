// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ObjectAccessControls resources represent the Access Control Lists
// (ACLs) for objects within Google Cloud Storage. ACLs let you specify
// who has access to your data and to what extent.
//
// There are two roles that can be assigned to an entity:
//
// READERs can get an object, though the acl property will not be revealed.
// OWNERs are READERs, and they can get the acl property, update an object,
// and call all objectAccessControls methods on the object. The owner of an
// object is always an OWNER.
// For more information, see Access Control, with the caveat that this API
// uses READER and OWNER instead of READ and FULL_CONTROL.
//
// To get more information about ObjectAccessControl, see:
//
// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)
//
// ## Example Usage
// ### Storage Object Access Control Public Object
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
//				Location: pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			object, err := storage.NewBucketObject(ctx, "object", &storage.BucketObjectArgs{
//				Bucket: bucket.Name,
//				Source: pulumi.NewFileAsset("../static/img/header-logo.png"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewObjectAccessControl(ctx, "publicRule", &storage.ObjectAccessControlArgs{
//				Object: object.OutputName,
//				Bucket: bucket.Name,
//				Role:   pulumi.String("READER"),
//				Entity: pulumi.String("allUsers"),
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
// ObjectAccessControl can be imported using any of these accepted formats* `{{bucket}}/{{object}}/{{entity}}` When using the `pulumi import` command, ObjectAccessControl can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:storage/objectAccessControl:ObjectAccessControl default {{bucket}}/{{object}}/{{entity}}
//
// ```
type ObjectAccessControl struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The ID for the entity
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation pulumi.IntOutput `pulumi:"generation"`
	// The name of the object to apply the access control to.
	Object pulumi.StringOutput `pulumi:"object"`
	// The project team associated with the entity
	// Structure is documented below.
	ProjectTeams ObjectAccessControlProjectTeamArrayOutput `pulumi:"projectTeams"`
	// The access permission for the entity.
	// Possible values are: `OWNER`, `READER`.
	//
	// ***
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewObjectAccessControl registers a new resource with the given unique name, arguments, and options.
func NewObjectAccessControl(ctx *pulumi.Context,
	name string, args *ObjectAccessControlArgs, opts ...pulumi.ResourceOption) (*ObjectAccessControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Entity == nil {
		return nil, errors.New("invalid value for required argument 'Entity'")
	}
	if args.Object == nil {
		return nil, errors.New("invalid value for required argument 'Object'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectAccessControl
	err := ctx.RegisterResource("gcp:storage/objectAccessControl:ObjectAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectAccessControl gets an existing ObjectAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectAccessControlState, opts ...pulumi.ResourceOption) (*ObjectAccessControl, error) {
	var resource ObjectAccessControl
	err := ctx.ReadResource("gcp:storage/objectAccessControl:ObjectAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectAccessControl resources.
type objectAccessControlState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity *string `pulumi:"entity"`
	// The ID for the entity
	EntityId *string `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation *int `pulumi:"generation"`
	// The name of the object to apply the access control to.
	Object *string `pulumi:"object"`
	// The project team associated with the entity
	// Structure is documented below.
	ProjectTeams []ObjectAccessControlProjectTeam `pulumi:"projectTeams"`
	// The access permission for the entity.
	// Possible values are: `OWNER`, `READER`.
	//
	// ***
	Role *string `pulumi:"role"`
}

type ObjectAccessControlState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The domain associated with the entity.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity pulumi.StringPtrInput
	// The ID for the entity
	EntityId pulumi.StringPtrInput
	// The content generation of the object, if applied to an object.
	Generation pulumi.IntPtrInput
	// The name of the object to apply the access control to.
	Object pulumi.StringPtrInput
	// The project team associated with the entity
	// Structure is documented below.
	ProjectTeams ObjectAccessControlProjectTeamArrayInput
	// The access permission for the entity.
	// Possible values are: `OWNER`, `READER`.
	//
	// ***
	Role pulumi.StringPtrInput
}

func (ObjectAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAccessControlState)(nil)).Elem()
}

type objectAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity string `pulumi:"entity"`
	// The name of the object to apply the access control to.
	Object string `pulumi:"object"`
	// The access permission for the entity.
	// Possible values are: `OWNER`, `READER`.
	//
	// ***
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ObjectAccessControl resource.
type ObjectAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity pulumi.StringInput
	// The name of the object to apply the access control to.
	Object pulumi.StringInput
	// The access permission for the entity.
	// Possible values are: `OWNER`, `READER`.
	//
	// ***
	Role pulumi.StringInput
}

func (ObjectAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAccessControlArgs)(nil)).Elem()
}

type ObjectAccessControlInput interface {
	pulumi.Input

	ToObjectAccessControlOutput() ObjectAccessControlOutput
	ToObjectAccessControlOutputWithContext(ctx context.Context) ObjectAccessControlOutput
}

func (*ObjectAccessControl) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAccessControl)(nil)).Elem()
}

func (i *ObjectAccessControl) ToObjectAccessControlOutput() ObjectAccessControlOutput {
	return i.ToObjectAccessControlOutputWithContext(context.Background())
}

func (i *ObjectAccessControl) ToObjectAccessControlOutputWithContext(ctx context.Context) ObjectAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAccessControlOutput)
}

// ObjectAccessControlArrayInput is an input type that accepts ObjectAccessControlArray and ObjectAccessControlArrayOutput values.
// You can construct a concrete instance of `ObjectAccessControlArrayInput` via:
//
//	ObjectAccessControlArray{ ObjectAccessControlArgs{...} }
type ObjectAccessControlArrayInput interface {
	pulumi.Input

	ToObjectAccessControlArrayOutput() ObjectAccessControlArrayOutput
	ToObjectAccessControlArrayOutputWithContext(context.Context) ObjectAccessControlArrayOutput
}

type ObjectAccessControlArray []ObjectAccessControlInput

func (ObjectAccessControlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectAccessControl)(nil)).Elem()
}

func (i ObjectAccessControlArray) ToObjectAccessControlArrayOutput() ObjectAccessControlArrayOutput {
	return i.ToObjectAccessControlArrayOutputWithContext(context.Background())
}

func (i ObjectAccessControlArray) ToObjectAccessControlArrayOutputWithContext(ctx context.Context) ObjectAccessControlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAccessControlArrayOutput)
}

// ObjectAccessControlMapInput is an input type that accepts ObjectAccessControlMap and ObjectAccessControlMapOutput values.
// You can construct a concrete instance of `ObjectAccessControlMapInput` via:
//
//	ObjectAccessControlMap{ "key": ObjectAccessControlArgs{...} }
type ObjectAccessControlMapInput interface {
	pulumi.Input

	ToObjectAccessControlMapOutput() ObjectAccessControlMapOutput
	ToObjectAccessControlMapOutputWithContext(context.Context) ObjectAccessControlMapOutput
}

type ObjectAccessControlMap map[string]ObjectAccessControlInput

func (ObjectAccessControlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectAccessControl)(nil)).Elem()
}

func (i ObjectAccessControlMap) ToObjectAccessControlMapOutput() ObjectAccessControlMapOutput {
	return i.ToObjectAccessControlMapOutputWithContext(context.Background())
}

func (i ObjectAccessControlMap) ToObjectAccessControlMapOutputWithContext(ctx context.Context) ObjectAccessControlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAccessControlMapOutput)
}

type ObjectAccessControlOutput struct{ *pulumi.OutputState }

func (ObjectAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAccessControl)(nil)).Elem()
}

func (o ObjectAccessControlOutput) ToObjectAccessControlOutput() ObjectAccessControlOutput {
	return o
}

func (o ObjectAccessControlOutput) ToObjectAccessControlOutputWithContext(ctx context.Context) ObjectAccessControlOutput {
	return o
}

// The name of the bucket.
func (o ObjectAccessControlOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The domain associated with the entity.
func (o ObjectAccessControlOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The email address associated with the entity.
func (o ObjectAccessControlOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The entity holding the permission, in one of the following forms:
// * user-{{userId}}
// * user-{{email}} (such as "user-liz@example.com")
// * group-{{groupId}}
// * group-{{email}} (such as "group-example@googlegroups.com")
// * domain-{{domain}} (such as "domain-example.com")
// * project-team-{{projectId}}
// * allUsers
// * allAuthenticatedUsers
func (o ObjectAccessControlOutput) Entity() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Entity }).(pulumi.StringOutput)
}

// The ID for the entity
func (o ObjectAccessControlOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// The content generation of the object, if applied to an object.
func (o ObjectAccessControlOutput) Generation() pulumi.IntOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.IntOutput { return v.Generation }).(pulumi.IntOutput)
}

// The name of the object to apply the access control to.
func (o ObjectAccessControlOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Object }).(pulumi.StringOutput)
}

// The project team associated with the entity
// Structure is documented below.
func (o ObjectAccessControlOutput) ProjectTeams() ObjectAccessControlProjectTeamArrayOutput {
	return o.ApplyT(func(v *ObjectAccessControl) ObjectAccessControlProjectTeamArrayOutput { return v.ProjectTeams }).(ObjectAccessControlProjectTeamArrayOutput)
}

// The access permission for the entity.
// Possible values are: `OWNER`, `READER`.
//
// ***
func (o ObjectAccessControlOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ObjectAccessControlArrayOutput struct{ *pulumi.OutputState }

func (ObjectAccessControlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectAccessControl)(nil)).Elem()
}

func (o ObjectAccessControlArrayOutput) ToObjectAccessControlArrayOutput() ObjectAccessControlArrayOutput {
	return o
}

func (o ObjectAccessControlArrayOutput) ToObjectAccessControlArrayOutputWithContext(ctx context.Context) ObjectAccessControlArrayOutput {
	return o
}

func (o ObjectAccessControlArrayOutput) Index(i pulumi.IntInput) ObjectAccessControlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectAccessControl {
		return vs[0].([]*ObjectAccessControl)[vs[1].(int)]
	}).(ObjectAccessControlOutput)
}

type ObjectAccessControlMapOutput struct{ *pulumi.OutputState }

func (ObjectAccessControlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectAccessControl)(nil)).Elem()
}

func (o ObjectAccessControlMapOutput) ToObjectAccessControlMapOutput() ObjectAccessControlMapOutput {
	return o
}

func (o ObjectAccessControlMapOutput) ToObjectAccessControlMapOutputWithContext(ctx context.Context) ObjectAccessControlMapOutput {
	return o
}

func (o ObjectAccessControlMapOutput) MapIndex(k pulumi.StringInput) ObjectAccessControlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectAccessControl {
		return vs[0].(map[string]*ObjectAccessControl)[vs[1].(string)]
	}).(ObjectAccessControlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAccessControlInput)(nil)).Elem(), &ObjectAccessControl{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAccessControlArrayInput)(nil)).Elem(), ObjectAccessControlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAccessControlMapInput)(nil)).Elem(), ObjectAccessControlMap{})
	pulumi.RegisterOutputType(ObjectAccessControlOutput{})
	pulumi.RegisterOutputType(ObjectAccessControlArrayOutput{})
	pulumi.RegisterOutputType(ObjectAccessControlMapOutput{})
}
