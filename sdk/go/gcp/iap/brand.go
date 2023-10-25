// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### Iap Brand
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.NewProject(ctx, "project", &organizations.ProjectArgs{
//				ProjectId: pulumi.String("my-project"),
//				OrgId:     pulumi.String("123456789"),
//			})
//			if err != nil {
//				return err
//			}
//			projectService, err := projects.NewService(ctx, "projectService", &projects.ServiceArgs{
//				Project: project.ProjectId,
//				Service: pulumi.String("iap.googleapis.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iap.NewBrand(ctx, "projectBrand", &iap.BrandArgs{
//				SupportEmail:     pulumi.String("support@example.com"),
//				ApplicationTitle: pulumi.String("Cloud IAP protected Application"),
//				Project:          projectService.Project,
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
// # Brand can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:iap/brand:Brand default projects/{{project_id}}/brands/{{brand_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:iap/brand:Brand default projects/{{project_number}}/brands/{{brand_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:iap/brand:Brand default {{project_number}}/{{brand_id}}
//
// ```
type Brand struct {
	pulumi.CustomResourceState

	// Application name displayed on OAuth consent screen.
	//
	// ***
	ApplicationTitle pulumi.StringOutput `pulumi:"applicationTitle"`
	// Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
	// NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
	// NOTE: The brand identification corresponds to the project number as only one
	// brand can be created per project.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the brand is only intended for usage inside the GSuite organization only.
	OrgInternalOnly pulumi.BoolOutput `pulumi:"orgInternalOnly"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail pulumi.StringOutput `pulumi:"supportEmail"`
}

// NewBrand registers a new resource with the given unique name, arguments, and options.
func NewBrand(ctx *pulumi.Context,
	name string, args *BrandArgs, opts ...pulumi.ResourceOption) (*Brand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationTitle == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationTitle'")
	}
	if args.SupportEmail == nil {
		return nil, errors.New("invalid value for required argument 'SupportEmail'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Brand
	err := ctx.RegisterResource("gcp:iap/brand:Brand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBrand gets an existing Brand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBrand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BrandState, opts ...pulumi.ResourceOption) (*Brand, error) {
	var resource Brand
	err := ctx.ReadResource("gcp:iap/brand:Brand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Brand resources.
type brandState struct {
	// Application name displayed on OAuth consent screen.
	//
	// ***
	ApplicationTitle *string `pulumi:"applicationTitle"`
	// Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
	// NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
	// NOTE: The brand identification corresponds to the project number as only one
	// brand can be created per project.
	Name *string `pulumi:"name"`
	// Whether the brand is only intended for usage inside the GSuite organization only.
	OrgInternalOnly *bool `pulumi:"orgInternalOnly"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail *string `pulumi:"supportEmail"`
}

type BrandState struct {
	// Application name displayed on OAuth consent screen.
	//
	// ***
	ApplicationTitle pulumi.StringPtrInput
	// Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
	// NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
	// NOTE: The brand identification corresponds to the project number as only one
	// brand can be created per project.
	Name pulumi.StringPtrInput
	// Whether the brand is only intended for usage inside the GSuite organization only.
	OrgInternalOnly pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail pulumi.StringPtrInput
}

func (BrandState) ElementType() reflect.Type {
	return reflect.TypeOf((*brandState)(nil)).Elem()
}

type brandArgs struct {
	// Application name displayed on OAuth consent screen.
	//
	// ***
	ApplicationTitle string `pulumi:"applicationTitle"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail string `pulumi:"supportEmail"`
}

// The set of arguments for constructing a Brand resource.
type BrandArgs struct {
	// Application name displayed on OAuth consent screen.
	//
	// ***
	ApplicationTitle pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail pulumi.StringInput
}

func (BrandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*brandArgs)(nil)).Elem()
}

type BrandInput interface {
	pulumi.Input

	ToBrandOutput() BrandOutput
	ToBrandOutputWithContext(ctx context.Context) BrandOutput
}

func (*Brand) ElementType() reflect.Type {
	return reflect.TypeOf((**Brand)(nil)).Elem()
}

func (i *Brand) ToBrandOutput() BrandOutput {
	return i.ToBrandOutputWithContext(context.Background())
}

func (i *Brand) ToBrandOutputWithContext(ctx context.Context) BrandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrandOutput)
}

func (i *Brand) ToOutput(ctx context.Context) pulumix.Output[*Brand] {
	return pulumix.Output[*Brand]{
		OutputState: i.ToBrandOutputWithContext(ctx).OutputState,
	}
}

// BrandArrayInput is an input type that accepts BrandArray and BrandArrayOutput values.
// You can construct a concrete instance of `BrandArrayInput` via:
//
//	BrandArray{ BrandArgs{...} }
type BrandArrayInput interface {
	pulumi.Input

	ToBrandArrayOutput() BrandArrayOutput
	ToBrandArrayOutputWithContext(context.Context) BrandArrayOutput
}

type BrandArray []BrandInput

func (BrandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Brand)(nil)).Elem()
}

func (i BrandArray) ToBrandArrayOutput() BrandArrayOutput {
	return i.ToBrandArrayOutputWithContext(context.Background())
}

func (i BrandArray) ToBrandArrayOutputWithContext(ctx context.Context) BrandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrandArrayOutput)
}

func (i BrandArray) ToOutput(ctx context.Context) pulumix.Output[[]*Brand] {
	return pulumix.Output[[]*Brand]{
		OutputState: i.ToBrandArrayOutputWithContext(ctx).OutputState,
	}
}

// BrandMapInput is an input type that accepts BrandMap and BrandMapOutput values.
// You can construct a concrete instance of `BrandMapInput` via:
//
//	BrandMap{ "key": BrandArgs{...} }
type BrandMapInput interface {
	pulumi.Input

	ToBrandMapOutput() BrandMapOutput
	ToBrandMapOutputWithContext(context.Context) BrandMapOutput
}

type BrandMap map[string]BrandInput

func (BrandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Brand)(nil)).Elem()
}

func (i BrandMap) ToBrandMapOutput() BrandMapOutput {
	return i.ToBrandMapOutputWithContext(context.Background())
}

func (i BrandMap) ToBrandMapOutputWithContext(ctx context.Context) BrandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrandMapOutput)
}

func (i BrandMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Brand] {
	return pulumix.Output[map[string]*Brand]{
		OutputState: i.ToBrandMapOutputWithContext(ctx).OutputState,
	}
}

type BrandOutput struct{ *pulumi.OutputState }

func (BrandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Brand)(nil)).Elem()
}

func (o BrandOutput) ToBrandOutput() BrandOutput {
	return o
}

func (o BrandOutput) ToBrandOutputWithContext(ctx context.Context) BrandOutput {
	return o
}

func (o BrandOutput) ToOutput(ctx context.Context) pulumix.Output[*Brand] {
	return pulumix.Output[*Brand]{
		OutputState: o.OutputState,
	}
}

// Application name displayed on OAuth consent screen.
//
// ***
func (o BrandOutput) ApplicationTitle() pulumi.StringOutput {
	return o.ApplyT(func(v *Brand) pulumi.StringOutput { return v.ApplicationTitle }).(pulumi.StringOutput)
}

// Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
// NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
// NOTE: The brand identification corresponds to the project number as only one
// brand can be created per project.
func (o BrandOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Brand) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the brand is only intended for usage inside the GSuite organization only.
func (o BrandOutput) OrgInternalOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *Brand) pulumi.BoolOutput { return v.OrgInternalOnly }).(pulumi.BoolOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o BrandOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Brand) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Support email displayed on the OAuth consent screen. Can be either a
// user or group email. When a user email is specified, the caller must
// be the user with the associated email address. When a group email is
// specified, the caller can be either a user or a service account which
// is an owner of the specified group in Cloud Identity.
func (o BrandOutput) SupportEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *Brand) pulumi.StringOutput { return v.SupportEmail }).(pulumi.StringOutput)
}

type BrandArrayOutput struct{ *pulumi.OutputState }

func (BrandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Brand)(nil)).Elem()
}

func (o BrandArrayOutput) ToBrandArrayOutput() BrandArrayOutput {
	return o
}

func (o BrandArrayOutput) ToBrandArrayOutputWithContext(ctx context.Context) BrandArrayOutput {
	return o
}

func (o BrandArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Brand] {
	return pulumix.Output[[]*Brand]{
		OutputState: o.OutputState,
	}
}

func (o BrandArrayOutput) Index(i pulumi.IntInput) BrandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Brand {
		return vs[0].([]*Brand)[vs[1].(int)]
	}).(BrandOutput)
}

type BrandMapOutput struct{ *pulumi.OutputState }

func (BrandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Brand)(nil)).Elem()
}

func (o BrandMapOutput) ToBrandMapOutput() BrandMapOutput {
	return o
}

func (o BrandMapOutput) ToBrandMapOutputWithContext(ctx context.Context) BrandMapOutput {
	return o
}

func (o BrandMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Brand] {
	return pulumix.Output[map[string]*Brand]{
		OutputState: o.OutputState,
	}
}

func (o BrandMapOutput) MapIndex(k pulumi.StringInput) BrandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Brand {
		return vs[0].(map[string]*Brand)[vs[1].(string)]
	}).(BrandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BrandInput)(nil)).Elem(), &Brand{})
	pulumi.RegisterInputType(reflect.TypeOf((*BrandArrayInput)(nil)).Elem(), BrandArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BrandMapInput)(nil)).Elem(), BrandMap{})
	pulumi.RegisterOutputType(BrandOutput{})
	pulumi.RegisterOutputType(BrandArrayOutput{})
	pulumi.RegisterOutputType(BrandMapOutput{})
}
