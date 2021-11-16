// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The DefaultObjectAccessControls resources represent the Access Control
// Lists (ACLs) applied to a new object within a Google Cloud Storage bucket
// when no ACL was provided for that object. ACLs let you specify who has
// access to your bucket contents and to what extent.
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
// To get more information about DefaultObjectAccessControl, see:
//
// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)
//
// ## Example Usage
// ### Storage Default Object Access Control Public
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
// 		bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
// 			Location: pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewDefaultObjectAccessControl(ctx, "publicRule", &storage.DefaultObjectAccessControlArgs{
// 			Bucket: bucket.Name,
// 			Role:   pulumi.String("READER"),
// 			Entity: pulumi.String("allUsers"),
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
// DefaultObjectAccessControl can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:storage/defaultObjectAccessControl:DefaultObjectAccessControl default {{bucket}}/{{entity}}
// ```
type DefaultObjectAccessControl struct {
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
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrOutput `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeams DefaultObjectAccessControlProjectTeamArrayOutput `pulumi:"projectTeams"`
	// The access permission for the entity.
	// Possible values are `OWNER` and `READER`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDefaultObjectAccessControl registers a new resource with the given unique name, arguments, and options.
func NewDefaultObjectAccessControl(ctx *pulumi.Context,
	name string, args *DefaultObjectAccessControlArgs, opts ...pulumi.ResourceOption) (*DefaultObjectAccessControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Entity == nil {
		return nil, errors.New("invalid value for required argument 'Entity'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DefaultObjectAccessControl
	err := ctx.RegisterResource("gcp:storage/defaultObjectAccessControl:DefaultObjectAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultObjectAccessControl gets an existing DefaultObjectAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultObjectAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultObjectAccessControlState, opts ...pulumi.ResourceOption) (*DefaultObjectAccessControl, error) {
	var resource DefaultObjectAccessControl
	err := ctx.ReadResource("gcp:storage/defaultObjectAccessControl:DefaultObjectAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultObjectAccessControl resources.
type defaultObjectAccessControlState struct {
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
	// The name of the object, if applied to an object.
	Object *string `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeams []DefaultObjectAccessControlProjectTeam `pulumi:"projectTeams"`
	// The access permission for the entity.
	// Possible values are `OWNER` and `READER`.
	Role *string `pulumi:"role"`
}

type DefaultObjectAccessControlState struct {
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
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrInput
	// The project team associated with the entity
	ProjectTeams DefaultObjectAccessControlProjectTeamArrayInput
	// The access permission for the entity.
	// Possible values are `OWNER` and `READER`.
	Role pulumi.StringPtrInput
}

func (DefaultObjectAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectAccessControlState)(nil)).Elem()
}

type defaultObjectAccessControlArgs struct {
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
	// The name of the object, if applied to an object.
	Object *string `pulumi:"object"`
	// The access permission for the entity.
	// Possible values are `OWNER` and `READER`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DefaultObjectAccessControl resource.
type DefaultObjectAccessControlArgs struct {
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
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrInput
	// The access permission for the entity.
	// Possible values are `OWNER` and `READER`.
	Role pulumi.StringInput
}

func (DefaultObjectAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectAccessControlArgs)(nil)).Elem()
}

type DefaultObjectAccessControlInput interface {
	pulumi.Input

	ToDefaultObjectAccessControlOutput() DefaultObjectAccessControlOutput
	ToDefaultObjectAccessControlOutputWithContext(ctx context.Context) DefaultObjectAccessControlOutput
}

func (*DefaultObjectAccessControl) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultObjectAccessControl)(nil))
}

func (i *DefaultObjectAccessControl) ToDefaultObjectAccessControlOutput() DefaultObjectAccessControlOutput {
	return i.ToDefaultObjectAccessControlOutputWithContext(context.Background())
}

func (i *DefaultObjectAccessControl) ToDefaultObjectAccessControlOutputWithContext(ctx context.Context) DefaultObjectAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectAccessControlOutput)
}

func (i *DefaultObjectAccessControl) ToDefaultObjectAccessControlPtrOutput() DefaultObjectAccessControlPtrOutput {
	return i.ToDefaultObjectAccessControlPtrOutputWithContext(context.Background())
}

func (i *DefaultObjectAccessControl) ToDefaultObjectAccessControlPtrOutputWithContext(ctx context.Context) DefaultObjectAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectAccessControlPtrOutput)
}

type DefaultObjectAccessControlPtrInput interface {
	pulumi.Input

	ToDefaultObjectAccessControlPtrOutput() DefaultObjectAccessControlPtrOutput
	ToDefaultObjectAccessControlPtrOutputWithContext(ctx context.Context) DefaultObjectAccessControlPtrOutput
}

type defaultObjectAccessControlPtrType DefaultObjectAccessControlArgs

func (*defaultObjectAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultObjectAccessControl)(nil))
}

func (i *defaultObjectAccessControlPtrType) ToDefaultObjectAccessControlPtrOutput() DefaultObjectAccessControlPtrOutput {
	return i.ToDefaultObjectAccessControlPtrOutputWithContext(context.Background())
}

func (i *defaultObjectAccessControlPtrType) ToDefaultObjectAccessControlPtrOutputWithContext(ctx context.Context) DefaultObjectAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectAccessControlPtrOutput)
}

// DefaultObjectAccessControlArrayInput is an input type that accepts DefaultObjectAccessControlArray and DefaultObjectAccessControlArrayOutput values.
// You can construct a concrete instance of `DefaultObjectAccessControlArrayInput` via:
//
//          DefaultObjectAccessControlArray{ DefaultObjectAccessControlArgs{...} }
type DefaultObjectAccessControlArrayInput interface {
	pulumi.Input

	ToDefaultObjectAccessControlArrayOutput() DefaultObjectAccessControlArrayOutput
	ToDefaultObjectAccessControlArrayOutputWithContext(context.Context) DefaultObjectAccessControlArrayOutput
}

type DefaultObjectAccessControlArray []DefaultObjectAccessControlInput

func (DefaultObjectAccessControlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultObjectAccessControl)(nil)).Elem()
}

func (i DefaultObjectAccessControlArray) ToDefaultObjectAccessControlArrayOutput() DefaultObjectAccessControlArrayOutput {
	return i.ToDefaultObjectAccessControlArrayOutputWithContext(context.Background())
}

func (i DefaultObjectAccessControlArray) ToDefaultObjectAccessControlArrayOutputWithContext(ctx context.Context) DefaultObjectAccessControlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectAccessControlArrayOutput)
}

// DefaultObjectAccessControlMapInput is an input type that accepts DefaultObjectAccessControlMap and DefaultObjectAccessControlMapOutput values.
// You can construct a concrete instance of `DefaultObjectAccessControlMapInput` via:
//
//          DefaultObjectAccessControlMap{ "key": DefaultObjectAccessControlArgs{...} }
type DefaultObjectAccessControlMapInput interface {
	pulumi.Input

	ToDefaultObjectAccessControlMapOutput() DefaultObjectAccessControlMapOutput
	ToDefaultObjectAccessControlMapOutputWithContext(context.Context) DefaultObjectAccessControlMapOutput
}

type DefaultObjectAccessControlMap map[string]DefaultObjectAccessControlInput

func (DefaultObjectAccessControlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultObjectAccessControl)(nil)).Elem()
}

func (i DefaultObjectAccessControlMap) ToDefaultObjectAccessControlMapOutput() DefaultObjectAccessControlMapOutput {
	return i.ToDefaultObjectAccessControlMapOutputWithContext(context.Background())
}

func (i DefaultObjectAccessControlMap) ToDefaultObjectAccessControlMapOutputWithContext(ctx context.Context) DefaultObjectAccessControlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectAccessControlMapOutput)
}

type DefaultObjectAccessControlOutput struct{ *pulumi.OutputState }

func (DefaultObjectAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultObjectAccessControl)(nil))
}

func (o DefaultObjectAccessControlOutput) ToDefaultObjectAccessControlOutput() DefaultObjectAccessControlOutput {
	return o
}

func (o DefaultObjectAccessControlOutput) ToDefaultObjectAccessControlOutputWithContext(ctx context.Context) DefaultObjectAccessControlOutput {
	return o
}

func (o DefaultObjectAccessControlOutput) ToDefaultObjectAccessControlPtrOutput() DefaultObjectAccessControlPtrOutput {
	return o.ToDefaultObjectAccessControlPtrOutputWithContext(context.Background())
}

func (o DefaultObjectAccessControlOutput) ToDefaultObjectAccessControlPtrOutputWithContext(ctx context.Context) DefaultObjectAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultObjectAccessControl) *DefaultObjectAccessControl {
		return &v
	}).(DefaultObjectAccessControlPtrOutput)
}

type DefaultObjectAccessControlPtrOutput struct{ *pulumi.OutputState }

func (DefaultObjectAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultObjectAccessControl)(nil))
}

func (o DefaultObjectAccessControlPtrOutput) ToDefaultObjectAccessControlPtrOutput() DefaultObjectAccessControlPtrOutput {
	return o
}

func (o DefaultObjectAccessControlPtrOutput) ToDefaultObjectAccessControlPtrOutputWithContext(ctx context.Context) DefaultObjectAccessControlPtrOutput {
	return o
}

func (o DefaultObjectAccessControlPtrOutput) Elem() DefaultObjectAccessControlOutput {
	return o.ApplyT(func(v *DefaultObjectAccessControl) DefaultObjectAccessControl {
		if v != nil {
			return *v
		}
		var ret DefaultObjectAccessControl
		return ret
	}).(DefaultObjectAccessControlOutput)
}

type DefaultObjectAccessControlArrayOutput struct{ *pulumi.OutputState }

func (DefaultObjectAccessControlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefaultObjectAccessControl)(nil))
}

func (o DefaultObjectAccessControlArrayOutput) ToDefaultObjectAccessControlArrayOutput() DefaultObjectAccessControlArrayOutput {
	return o
}

func (o DefaultObjectAccessControlArrayOutput) ToDefaultObjectAccessControlArrayOutputWithContext(ctx context.Context) DefaultObjectAccessControlArrayOutput {
	return o
}

func (o DefaultObjectAccessControlArrayOutput) Index(i pulumi.IntInput) DefaultObjectAccessControlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DefaultObjectAccessControl {
		return vs[0].([]DefaultObjectAccessControl)[vs[1].(int)]
	}).(DefaultObjectAccessControlOutput)
}

type DefaultObjectAccessControlMapOutput struct{ *pulumi.OutputState }

func (DefaultObjectAccessControlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DefaultObjectAccessControl)(nil))
}

func (o DefaultObjectAccessControlMapOutput) ToDefaultObjectAccessControlMapOutput() DefaultObjectAccessControlMapOutput {
	return o
}

func (o DefaultObjectAccessControlMapOutput) ToDefaultObjectAccessControlMapOutputWithContext(ctx context.Context) DefaultObjectAccessControlMapOutput {
	return o
}

func (o DefaultObjectAccessControlMapOutput) MapIndex(k pulumi.StringInput) DefaultObjectAccessControlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DefaultObjectAccessControl {
		return vs[0].(map[string]DefaultObjectAccessControl)[vs[1].(string)]
	}).(DefaultObjectAccessControlOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultObjectAccessControlOutput{})
	pulumi.RegisterOutputType(DefaultObjectAccessControlPtrOutput{})
	pulumi.RegisterOutputType(DefaultObjectAccessControlArrayOutput{})
	pulumi.RegisterOutputType(DefaultObjectAccessControlMapOutput{})
}
