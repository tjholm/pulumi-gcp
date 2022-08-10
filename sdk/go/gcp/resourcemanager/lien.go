// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Lien represents an encumbrance on the actions that can be performed on a resource.
//
// ## Example Usage
// ### Resource Manager Lien
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.NewProject(ctx, "project", &organizations.ProjectArgs{
//				ProjectId: pulumi.String("staging-project"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourcemanager.NewLien(ctx, "lien", &resourcemanager.LienArgs{
//				Origin: pulumi.String("machine-readable-explanation"),
//				Parent: project.Number.ApplyT(func(number string) (string, error) {
//					return fmt.Sprintf("projects/%v", number), nil
//				}).(pulumi.StringOutput),
//				Reason: pulumi.String("This project is an important environment"),
//				Restrictions: pulumi.StringArray{
//					pulumi.String("resourcemanager.projects.delete"),
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
// # Lien can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:resourcemanager/lien:Lien default {{parent}}/{{name}}
//
// ```
type Lien struct {
	pulumi.CustomResourceState

	// Time of creation
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A system-generated unique identifier for this Lien.
	Name pulumi.StringOutput `pulumi:"name"`
	// A stable, user-visible/meaningful string identifying the origin
	// of the Lien, intended to be inspected programmatically. Maximum length of
	// 200 characters.
	Origin pulumi.StringOutput `pulumi:"origin"`
	// A reference to the resource this Lien is attached to.
	// The server will validate the parent against those for which Liens are supported.
	// Since a variety of objects can have Liens against them, you must provide the type
	// prefix (e.g. "projects/my-project-name").
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Concise user-visible strings indicating why an action cannot be performed
	// on a resource. Maximum length of 200 characters.
	Reason pulumi.StringOutput `pulumi:"reason"`
	// The types of operations which should be blocked as a result of this Lien.
	// Each value should correspond to an IAM permission. The server will validate
	// the permissions against those for which Liens are supported.  An empty
	// list is meaningless and will be rejected.
	// e.g. ['resourcemanager.projects.delete']
	Restrictions pulumi.StringArrayOutput `pulumi:"restrictions"`
}

// NewLien registers a new resource with the given unique name, arguments, and options.
func NewLien(ctx *pulumi.Context,
	name string, args *LienArgs, opts ...pulumi.ResourceOption) (*Lien, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Origin == nil {
		return nil, errors.New("invalid value for required argument 'Origin'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.Reason == nil {
		return nil, errors.New("invalid value for required argument 'Reason'")
	}
	if args.Restrictions == nil {
		return nil, errors.New("invalid value for required argument 'Restrictions'")
	}
	var resource Lien
	err := ctx.RegisterResource("gcp:resourcemanager/lien:Lien", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLien gets an existing Lien resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLien(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LienState, opts ...pulumi.ResourceOption) (*Lien, error) {
	var resource Lien
	err := ctx.ReadResource("gcp:resourcemanager/lien:Lien", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lien resources.
type lienState struct {
	// Time of creation
	CreateTime *string `pulumi:"createTime"`
	// A system-generated unique identifier for this Lien.
	Name *string `pulumi:"name"`
	// A stable, user-visible/meaningful string identifying the origin
	// of the Lien, intended to be inspected programmatically. Maximum length of
	// 200 characters.
	Origin *string `pulumi:"origin"`
	// A reference to the resource this Lien is attached to.
	// The server will validate the parent against those for which Liens are supported.
	// Since a variety of objects can have Liens against them, you must provide the type
	// prefix (e.g. "projects/my-project-name").
	Parent *string `pulumi:"parent"`
	// Concise user-visible strings indicating why an action cannot be performed
	// on a resource. Maximum length of 200 characters.
	Reason *string `pulumi:"reason"`
	// The types of operations which should be blocked as a result of this Lien.
	// Each value should correspond to an IAM permission. The server will validate
	// the permissions against those for which Liens are supported.  An empty
	// list is meaningless and will be rejected.
	// e.g. ['resourcemanager.projects.delete']
	Restrictions []string `pulumi:"restrictions"`
}

type LienState struct {
	// Time of creation
	CreateTime pulumi.StringPtrInput
	// A system-generated unique identifier for this Lien.
	Name pulumi.StringPtrInput
	// A stable, user-visible/meaningful string identifying the origin
	// of the Lien, intended to be inspected programmatically. Maximum length of
	// 200 characters.
	Origin pulumi.StringPtrInput
	// A reference to the resource this Lien is attached to.
	// The server will validate the parent against those for which Liens are supported.
	// Since a variety of objects can have Liens against them, you must provide the type
	// prefix (e.g. "projects/my-project-name").
	Parent pulumi.StringPtrInput
	// Concise user-visible strings indicating why an action cannot be performed
	// on a resource. Maximum length of 200 characters.
	Reason pulumi.StringPtrInput
	// The types of operations which should be blocked as a result of this Lien.
	// Each value should correspond to an IAM permission. The server will validate
	// the permissions against those for which Liens are supported.  An empty
	// list is meaningless and will be rejected.
	// e.g. ['resourcemanager.projects.delete']
	Restrictions pulumi.StringArrayInput
}

func (LienState) ElementType() reflect.Type {
	return reflect.TypeOf((*lienState)(nil)).Elem()
}

type lienArgs struct {
	// A stable, user-visible/meaningful string identifying the origin
	// of the Lien, intended to be inspected programmatically. Maximum length of
	// 200 characters.
	Origin string `pulumi:"origin"`
	// A reference to the resource this Lien is attached to.
	// The server will validate the parent against those for which Liens are supported.
	// Since a variety of objects can have Liens against them, you must provide the type
	// prefix (e.g. "projects/my-project-name").
	Parent string `pulumi:"parent"`
	// Concise user-visible strings indicating why an action cannot be performed
	// on a resource. Maximum length of 200 characters.
	Reason string `pulumi:"reason"`
	// The types of operations which should be blocked as a result of this Lien.
	// Each value should correspond to an IAM permission. The server will validate
	// the permissions against those for which Liens are supported.  An empty
	// list is meaningless and will be rejected.
	// e.g. ['resourcemanager.projects.delete']
	Restrictions []string `pulumi:"restrictions"`
}

// The set of arguments for constructing a Lien resource.
type LienArgs struct {
	// A stable, user-visible/meaningful string identifying the origin
	// of the Lien, intended to be inspected programmatically. Maximum length of
	// 200 characters.
	Origin pulumi.StringInput
	// A reference to the resource this Lien is attached to.
	// The server will validate the parent against those for which Liens are supported.
	// Since a variety of objects can have Liens against them, you must provide the type
	// prefix (e.g. "projects/my-project-name").
	Parent pulumi.StringInput
	// Concise user-visible strings indicating why an action cannot be performed
	// on a resource. Maximum length of 200 characters.
	Reason pulumi.StringInput
	// The types of operations which should be blocked as a result of this Lien.
	// Each value should correspond to an IAM permission. The server will validate
	// the permissions against those for which Liens are supported.  An empty
	// list is meaningless and will be rejected.
	// e.g. ['resourcemanager.projects.delete']
	Restrictions pulumi.StringArrayInput
}

func (LienArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lienArgs)(nil)).Elem()
}

type LienInput interface {
	pulumi.Input

	ToLienOutput() LienOutput
	ToLienOutputWithContext(ctx context.Context) LienOutput
}

func (*Lien) ElementType() reflect.Type {
	return reflect.TypeOf((**Lien)(nil)).Elem()
}

func (i *Lien) ToLienOutput() LienOutput {
	return i.ToLienOutputWithContext(context.Background())
}

func (i *Lien) ToLienOutputWithContext(ctx context.Context) LienOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LienOutput)
}

// LienArrayInput is an input type that accepts LienArray and LienArrayOutput values.
// You can construct a concrete instance of `LienArrayInput` via:
//
//	LienArray{ LienArgs{...} }
type LienArrayInput interface {
	pulumi.Input

	ToLienArrayOutput() LienArrayOutput
	ToLienArrayOutputWithContext(context.Context) LienArrayOutput
}

type LienArray []LienInput

func (LienArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Lien)(nil)).Elem()
}

func (i LienArray) ToLienArrayOutput() LienArrayOutput {
	return i.ToLienArrayOutputWithContext(context.Background())
}

func (i LienArray) ToLienArrayOutputWithContext(ctx context.Context) LienArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LienArrayOutput)
}

// LienMapInput is an input type that accepts LienMap and LienMapOutput values.
// You can construct a concrete instance of `LienMapInput` via:
//
//	LienMap{ "key": LienArgs{...} }
type LienMapInput interface {
	pulumi.Input

	ToLienMapOutput() LienMapOutput
	ToLienMapOutputWithContext(context.Context) LienMapOutput
}

type LienMap map[string]LienInput

func (LienMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Lien)(nil)).Elem()
}

func (i LienMap) ToLienMapOutput() LienMapOutput {
	return i.ToLienMapOutputWithContext(context.Background())
}

func (i LienMap) ToLienMapOutputWithContext(ctx context.Context) LienMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LienMapOutput)
}

type LienOutput struct{ *pulumi.OutputState }

func (LienOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lien)(nil)).Elem()
}

func (o LienOutput) ToLienOutput() LienOutput {
	return o
}

func (o LienOutput) ToLienOutputWithContext(ctx context.Context) LienOutput {
	return o
}

// Time of creation
func (o LienOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A system-generated unique identifier for this Lien.
func (o LienOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A stable, user-visible/meaningful string identifying the origin
// of the Lien, intended to be inspected programmatically. Maximum length of
// 200 characters.
func (o LienOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// A reference to the resource this Lien is attached to.
// The server will validate the parent against those for which Liens are supported.
// Since a variety of objects can have Liens against them, you must provide the type
// prefix (e.g. "projects/my-project-name").
func (o LienOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Concise user-visible strings indicating why an action cannot be performed
// on a resource. Maximum length of 200 characters.
func (o LienOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Reason }).(pulumi.StringOutput)
}

// The types of operations which should be blocked as a result of this Lien.
// Each value should correspond to an IAM permission. The server will validate
// the permissions against those for which Liens are supported.  An empty
// list is meaningless and will be rejected.
// e.g. ['resourcemanager.projects.delete']
func (o LienOutput) Restrictions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringArrayOutput { return v.Restrictions }).(pulumi.StringArrayOutput)
}

type LienArrayOutput struct{ *pulumi.OutputState }

func (LienArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Lien)(nil)).Elem()
}

func (o LienArrayOutput) ToLienArrayOutput() LienArrayOutput {
	return o
}

func (o LienArrayOutput) ToLienArrayOutputWithContext(ctx context.Context) LienArrayOutput {
	return o
}

func (o LienArrayOutput) Index(i pulumi.IntInput) LienOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Lien {
		return vs[0].([]*Lien)[vs[1].(int)]
	}).(LienOutput)
}

type LienMapOutput struct{ *pulumi.OutputState }

func (LienMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Lien)(nil)).Elem()
}

func (o LienMapOutput) ToLienMapOutput() LienMapOutput {
	return o
}

func (o LienMapOutput) ToLienMapOutputWithContext(ctx context.Context) LienMapOutput {
	return o
}

func (o LienMapOutput) MapIndex(k pulumi.StringInput) LienOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Lien {
		return vs[0].(map[string]*Lien)[vs[1].(string)]
	}).(LienOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LienInput)(nil)).Elem(), &Lien{})
	pulumi.RegisterInputType(reflect.TypeOf((*LienArrayInput)(nil)).Elem(), LienArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LienMapInput)(nil)).Elem(), LienMap{})
	pulumi.RegisterOutputType(LienOutput{})
	pulumi.RegisterOutputType(LienArrayOutput{})
	pulumi.RegisterOutputType(LienMapOutput{})
}
