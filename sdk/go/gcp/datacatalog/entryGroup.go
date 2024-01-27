// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An EntryGroup resource represents a logical grouping of zero or more Data Catalog Entry resources.
//
// To get more information about EntryGroup, see:
//
// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/data-catalog/docs)
//
// ## Example Usage
// ### Data Catalog Entry Group Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewEntryGroup(ctx, "basicEntryGroup", &datacatalog.EntryGroupArgs{
//				EntryGroupId: pulumi.String("my_group"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Data Catalog Entry Group Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewEntryGroup(ctx, "basicEntryGroup", &datacatalog.EntryGroupArgs{
//				Description:  pulumi.String("example entry group"),
//				DisplayName:  pulumi.String("entry group"),
//				EntryGroupId: pulumi.String("my_group"),
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
// EntryGroup can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, EntryGroup can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/entryGroup:EntryGroup default {{name}}
//
// ```
type EntryGroup struct {
	pulumi.CustomResourceState

	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	//
	// ***
	EntryGroupId pulumi.StringOutput `pulumi:"entryGroupId"`
	// The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// EntryGroup location region.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewEntryGroup registers a new resource with the given unique name, arguments, and options.
func NewEntryGroup(ctx *pulumi.Context,
	name string, args *EntryGroupArgs, opts ...pulumi.ResourceOption) (*EntryGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EntryGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EntryGroup
	err := ctx.RegisterResource("gcp:datacatalog/entryGroup:EntryGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroup gets an existing EntryGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupState, opts ...pulumi.ResourceOption) (*EntryGroup, error) {
	var resource EntryGroup
	err := ctx.ReadResource("gcp:datacatalog/entryGroup:EntryGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroup resources.
type entryGroupState struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description *string `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName *string `pulumi:"displayName"`
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	//
	// ***
	EntryGroupId *string `pulumi:"entryGroupId"`
	// The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// EntryGroup location region.
	Region *string `pulumi:"region"`
}

type EntryGroupState struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description pulumi.StringPtrInput
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName pulumi.StringPtrInput
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	//
	// ***
	EntryGroupId pulumi.StringPtrInput
	// The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// EntryGroup location region.
	Region pulumi.StringPtrInput
}

func (EntryGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupState)(nil)).Elem()
}

type entryGroupArgs struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description *string `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName *string `pulumi:"displayName"`
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	//
	// ***
	EntryGroupId string `pulumi:"entryGroupId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// EntryGroup location region.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a EntryGroup resource.
type EntryGroupArgs struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description pulumi.StringPtrInput
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName pulumi.StringPtrInput
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	//
	// ***
	EntryGroupId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// EntryGroup location region.
	Region pulumi.StringPtrInput
}

func (EntryGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupArgs)(nil)).Elem()
}

type EntryGroupInput interface {
	pulumi.Input

	ToEntryGroupOutput() EntryGroupOutput
	ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput
}

func (*EntryGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroup)(nil)).Elem()
}

func (i *EntryGroup) ToEntryGroupOutput() EntryGroupOutput {
	return i.ToEntryGroupOutputWithContext(context.Background())
}

func (i *EntryGroup) ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupOutput)
}

// EntryGroupArrayInput is an input type that accepts EntryGroupArray and EntryGroupArrayOutput values.
// You can construct a concrete instance of `EntryGroupArrayInput` via:
//
//	EntryGroupArray{ EntryGroupArgs{...} }
type EntryGroupArrayInput interface {
	pulumi.Input

	ToEntryGroupArrayOutput() EntryGroupArrayOutput
	ToEntryGroupArrayOutputWithContext(context.Context) EntryGroupArrayOutput
}

type EntryGroupArray []EntryGroupInput

func (EntryGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntryGroup)(nil)).Elem()
}

func (i EntryGroupArray) ToEntryGroupArrayOutput() EntryGroupArrayOutput {
	return i.ToEntryGroupArrayOutputWithContext(context.Background())
}

func (i EntryGroupArray) ToEntryGroupArrayOutputWithContext(ctx context.Context) EntryGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupArrayOutput)
}

// EntryGroupMapInput is an input type that accepts EntryGroupMap and EntryGroupMapOutput values.
// You can construct a concrete instance of `EntryGroupMapInput` via:
//
//	EntryGroupMap{ "key": EntryGroupArgs{...} }
type EntryGroupMapInput interface {
	pulumi.Input

	ToEntryGroupMapOutput() EntryGroupMapOutput
	ToEntryGroupMapOutputWithContext(context.Context) EntryGroupMapOutput
}

type EntryGroupMap map[string]EntryGroupInput

func (EntryGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntryGroup)(nil)).Elem()
}

func (i EntryGroupMap) ToEntryGroupMapOutput() EntryGroupMapOutput {
	return i.ToEntryGroupMapOutputWithContext(context.Background())
}

func (i EntryGroupMap) ToEntryGroupMapOutputWithContext(ctx context.Context) EntryGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupMapOutput)
}

type EntryGroupOutput struct{ *pulumi.OutputState }

func (EntryGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroup)(nil)).Elem()
}

func (o EntryGroupOutput) ToEntryGroupOutput() EntryGroupOutput {
	return o
}

func (o EntryGroupOutput) ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput {
	return o
}

// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
func (o EntryGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntryGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A short name to identify the entry group, for example, "analytics data - jan 2011".
func (o EntryGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntryGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The id of the entry group to create. The id must begin with a letter or underscore,
// contain only English letters, numbers and underscores, and be at most 64 characters.
//
// ***
func (o EntryGroupOutput) EntryGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroup) pulumi.StringOutput { return v.EntryGroupId }).(pulumi.StringOutput)
}

// The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
func (o EntryGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o EntryGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// EntryGroup location region.
func (o EntryGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type EntryGroupArrayOutput struct{ *pulumi.OutputState }

func (EntryGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntryGroup)(nil)).Elem()
}

func (o EntryGroupArrayOutput) ToEntryGroupArrayOutput() EntryGroupArrayOutput {
	return o
}

func (o EntryGroupArrayOutput) ToEntryGroupArrayOutputWithContext(ctx context.Context) EntryGroupArrayOutput {
	return o
}

func (o EntryGroupArrayOutput) Index(i pulumi.IntInput) EntryGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EntryGroup {
		return vs[0].([]*EntryGroup)[vs[1].(int)]
	}).(EntryGroupOutput)
}

type EntryGroupMapOutput struct{ *pulumi.OutputState }

func (EntryGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntryGroup)(nil)).Elem()
}

func (o EntryGroupMapOutput) ToEntryGroupMapOutput() EntryGroupMapOutput {
	return o
}

func (o EntryGroupMapOutput) ToEntryGroupMapOutputWithContext(ctx context.Context) EntryGroupMapOutput {
	return o
}

func (o EntryGroupMapOutput) MapIndex(k pulumi.StringInput) EntryGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EntryGroup {
		return vs[0].(map[string]*EntryGroup)[vs[1].(string)]
	}).(EntryGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupInput)(nil)).Elem(), &EntryGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupArrayInput)(nil)).Elem(), EntryGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupMapInput)(nil)).Elem(), EntryGroupMap{})
	pulumi.RegisterOutputType(EntryGroupOutput{})
	pulumi.RegisterOutputType(EntryGroupArrayOutput{})
	pulumi.RegisterOutputType(EntryGroupMapOutput{})
}
