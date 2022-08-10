// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Cloud AI Platform Notebook environment.
//
// To get more information about Environment, see:
//
// * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
//
// ## Example Usage
// ### Notebook Environment Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := notebooks.NewEnvironment(ctx, "environment", &notebooks.EnvironmentArgs{
//				ContainerImage: &notebooks.EnvironmentContainerImageArgs{
//					Repository: pulumi.String("gcr.io/deeplearning-platform-release/base-cpu"),
//				},
//				Location: pulumi.String("us-west1-a"),
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
// # Environment can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:notebooks/environment:Environment default projects/{{project}}/locations/{{location}}/environments/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:notebooks/environment:Environment default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:notebooks/environment:Environment default {{location}}/{{name}}
//
// ```
type Environment struct {
	pulumi.CustomResourceState

	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage EnvironmentContainerImagePtrOutput `pulumi:"containerImage"`
	// Instance creation time
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A brief description of this environment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrOutput `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringOutput `pulumi:"project"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage EnvironmentVmImagePtrOutput `pulumi:"vmImage"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Environment
	err := ctx.RegisterResource("gcp:notebooks/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("gcp:notebooks/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage *EnvironmentContainerImage `pulumi:"containerImage"`
	// Instance creation time
	CreateTime *string `pulumi:"createTime"`
	// A brief description of this environment.
	Description *string `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName *string `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location *string `pulumi:"location"`
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name *string `pulumi:"name"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript *string `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project *string `pulumi:"project"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage *EnvironmentVmImage `pulumi:"vmImage"`
}

type EnvironmentState struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage EnvironmentContainerImagePtrInput
	// Instance creation time
	CreateTime pulumi.StringPtrInput
	// A brief description of this environment.
	Description pulumi.StringPtrInput
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringPtrInput
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name pulumi.StringPtrInput
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrInput
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage EnvironmentVmImagePtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage *EnvironmentContainerImage `pulumi:"containerImage"`
	// A brief description of this environment.
	Description *string `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName *string `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location string `pulumi:"location"`
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name *string `pulumi:"name"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript *string `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project *string `pulumi:"project"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage *EnvironmentVmImage `pulumi:"vmImage"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage EnvironmentContainerImagePtrInput
	// A brief description of this environment.
	Description pulumi.StringPtrInput
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringInput
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name pulumi.StringPtrInput
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrInput
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage EnvironmentVmImagePtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

// EnvironmentArrayInput is an input type that accepts EnvironmentArray and EnvironmentArrayOutput values.
// You can construct a concrete instance of `EnvironmentArrayInput` via:
//
//	EnvironmentArray{ EnvironmentArgs{...} }
type EnvironmentArrayInput interface {
	pulumi.Input

	ToEnvironmentArrayOutput() EnvironmentArrayOutput
	ToEnvironmentArrayOutputWithContext(context.Context) EnvironmentArrayOutput
}

type EnvironmentArray []EnvironmentInput

func (EnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (i EnvironmentArray) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return i.ToEnvironmentArrayOutputWithContext(context.Background())
}

func (i EnvironmentArray) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentArrayOutput)
}

// EnvironmentMapInput is an input type that accepts EnvironmentMap and EnvironmentMapOutput values.
// You can construct a concrete instance of `EnvironmentMapInput` via:
//
//	EnvironmentMap{ "key": EnvironmentArgs{...} }
type EnvironmentMapInput interface {
	pulumi.Input

	ToEnvironmentMapOutput() EnvironmentMapOutput
	ToEnvironmentMapOutputWithContext(context.Context) EnvironmentMapOutput
}

type EnvironmentMap map[string]EnvironmentInput

func (EnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (i EnvironmentMap) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return i.ToEnvironmentMapOutputWithContext(context.Background())
}

func (i EnvironmentMap) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMapOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

// Use a container image to start the notebook instance.
// Structure is documented below.
func (o EnvironmentOutput) ContainerImage() EnvironmentContainerImagePtrOutput {
	return o.ApplyT(func(v *Environment) EnvironmentContainerImagePtrOutput { return v.ContainerImage }).(EnvironmentContainerImagePtrOutput)
}

// Instance creation time
func (o EnvironmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A brief description of this environment.
func (o EnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of this environment for the UI.
func (o EnvironmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// A reference to the zone where the machine resides.
func (o EnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name specified for the Environment instance.
// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
func (o EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path to a Bash script that automatically runs after a notebook instance fully boots up.
// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
func (o EnvironmentOutput) PostStartupScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.PostStartupScript }).(pulumi.StringPtrOutput)
}

// The name of the Google Cloud project that this VM image belongs to.
// Format: projects/{project_id}
func (o EnvironmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Use a Compute Engine VM image to start the notebook instance.
// Structure is documented below.
func (o EnvironmentOutput) VmImage() EnvironmentVmImagePtrOutput {
	return o.ApplyT(func(v *Environment) EnvironmentVmImagePtrOutput { return v.VmImage }).(EnvironmentVmImagePtrOutput)
}

type EnvironmentArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) Index(i pulumi.IntInput) EnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].([]*Environment)[vs[1].(int)]
	}).(EnvironmentOutput)
}

type EnvironmentMapOutput struct{ *pulumi.OutputState }

func (EnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) MapIndex(k pulumi.StringInput) EnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].(map[string]*Environment)[vs[1].(string)]
	}).(EnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentArrayInput)(nil)).Elem(), EnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentMapInput)(nil)).Elem(), EnvironmentMap{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentMapOutput{})
}
