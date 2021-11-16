// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables the Google Compute Engine
// [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
// feature for a project, assigning it as a Shared VPC host project.
//
// For more information, see,
// [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
// where the Shared VPC feature is referred to by its former name "XPN".
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		host, err := compute.NewSharedVPCHostProject(ctx, "host", &compute.SharedVPCHostProjectArgs{
// 			Project: pulumi.String("host-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSharedVPCServiceProject(ctx, "service1", &compute.SharedVPCServiceProjectArgs{
// 			HostProject:    host.Project,
// 			ServiceProject: pulumi.String("service-project-id-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSharedVPCServiceProject(ctx, "service2", &compute.SharedVPCServiceProjectArgs{
// 			HostProject:    host.Project,
// 			ServiceProject: pulumi.String("service-project-id-2"),
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
// Google Compute Engine Shared VPC host project feature can be imported using the `project`, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/sharedVPCHostProject:SharedVPCHostProject host host-project-id
// ```
type SharedVPCHostProject struct {
	pulumi.CustomResourceState

	// The ID of the project that will serve as a Shared VPC host project
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewSharedVPCHostProject registers a new resource with the given unique name, arguments, and options.
func NewSharedVPCHostProject(ctx *pulumi.Context,
	name string, args *SharedVPCHostProjectArgs, opts ...pulumi.ResourceOption) (*SharedVPCHostProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource SharedVPCHostProject
	err := ctx.RegisterResource("gcp:compute/sharedVPCHostProject:SharedVPCHostProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedVPCHostProject gets an existing SharedVPCHostProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedVPCHostProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedVPCHostProjectState, opts ...pulumi.ResourceOption) (*SharedVPCHostProject, error) {
	var resource SharedVPCHostProject
	err := ctx.ReadResource("gcp:compute/sharedVPCHostProject:SharedVPCHostProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedVPCHostProject resources.
type sharedVPCHostProjectState struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project *string `pulumi:"project"`
}

type SharedVPCHostProjectState struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project pulumi.StringPtrInput
}

func (SharedVPCHostProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedVPCHostProjectState)(nil)).Elem()
}

type sharedVPCHostProjectArgs struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a SharedVPCHostProject resource.
type SharedVPCHostProjectArgs struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project pulumi.StringInput
}

func (SharedVPCHostProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedVPCHostProjectArgs)(nil)).Elem()
}

type SharedVPCHostProjectInput interface {
	pulumi.Input

	ToSharedVPCHostProjectOutput() SharedVPCHostProjectOutput
	ToSharedVPCHostProjectOutputWithContext(ctx context.Context) SharedVPCHostProjectOutput
}

func (*SharedVPCHostProject) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedVPCHostProject)(nil))
}

func (i *SharedVPCHostProject) ToSharedVPCHostProjectOutput() SharedVPCHostProjectOutput {
	return i.ToSharedVPCHostProjectOutputWithContext(context.Background())
}

func (i *SharedVPCHostProject) ToSharedVPCHostProjectOutputWithContext(ctx context.Context) SharedVPCHostProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedVPCHostProjectOutput)
}

func (i *SharedVPCHostProject) ToSharedVPCHostProjectPtrOutput() SharedVPCHostProjectPtrOutput {
	return i.ToSharedVPCHostProjectPtrOutputWithContext(context.Background())
}

func (i *SharedVPCHostProject) ToSharedVPCHostProjectPtrOutputWithContext(ctx context.Context) SharedVPCHostProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedVPCHostProjectPtrOutput)
}

type SharedVPCHostProjectPtrInput interface {
	pulumi.Input

	ToSharedVPCHostProjectPtrOutput() SharedVPCHostProjectPtrOutput
	ToSharedVPCHostProjectPtrOutputWithContext(ctx context.Context) SharedVPCHostProjectPtrOutput
}

type sharedVPCHostProjectPtrType SharedVPCHostProjectArgs

func (*sharedVPCHostProjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedVPCHostProject)(nil))
}

func (i *sharedVPCHostProjectPtrType) ToSharedVPCHostProjectPtrOutput() SharedVPCHostProjectPtrOutput {
	return i.ToSharedVPCHostProjectPtrOutputWithContext(context.Background())
}

func (i *sharedVPCHostProjectPtrType) ToSharedVPCHostProjectPtrOutputWithContext(ctx context.Context) SharedVPCHostProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedVPCHostProjectPtrOutput)
}

// SharedVPCHostProjectArrayInput is an input type that accepts SharedVPCHostProjectArray and SharedVPCHostProjectArrayOutput values.
// You can construct a concrete instance of `SharedVPCHostProjectArrayInput` via:
//
//          SharedVPCHostProjectArray{ SharedVPCHostProjectArgs{...} }
type SharedVPCHostProjectArrayInput interface {
	pulumi.Input

	ToSharedVPCHostProjectArrayOutput() SharedVPCHostProjectArrayOutput
	ToSharedVPCHostProjectArrayOutputWithContext(context.Context) SharedVPCHostProjectArrayOutput
}

type SharedVPCHostProjectArray []SharedVPCHostProjectInput

func (SharedVPCHostProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SharedVPCHostProject)(nil)).Elem()
}

func (i SharedVPCHostProjectArray) ToSharedVPCHostProjectArrayOutput() SharedVPCHostProjectArrayOutput {
	return i.ToSharedVPCHostProjectArrayOutputWithContext(context.Background())
}

func (i SharedVPCHostProjectArray) ToSharedVPCHostProjectArrayOutputWithContext(ctx context.Context) SharedVPCHostProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedVPCHostProjectArrayOutput)
}

// SharedVPCHostProjectMapInput is an input type that accepts SharedVPCHostProjectMap and SharedVPCHostProjectMapOutput values.
// You can construct a concrete instance of `SharedVPCHostProjectMapInput` via:
//
//          SharedVPCHostProjectMap{ "key": SharedVPCHostProjectArgs{...} }
type SharedVPCHostProjectMapInput interface {
	pulumi.Input

	ToSharedVPCHostProjectMapOutput() SharedVPCHostProjectMapOutput
	ToSharedVPCHostProjectMapOutputWithContext(context.Context) SharedVPCHostProjectMapOutput
}

type SharedVPCHostProjectMap map[string]SharedVPCHostProjectInput

func (SharedVPCHostProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SharedVPCHostProject)(nil)).Elem()
}

func (i SharedVPCHostProjectMap) ToSharedVPCHostProjectMapOutput() SharedVPCHostProjectMapOutput {
	return i.ToSharedVPCHostProjectMapOutputWithContext(context.Background())
}

func (i SharedVPCHostProjectMap) ToSharedVPCHostProjectMapOutputWithContext(ctx context.Context) SharedVPCHostProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedVPCHostProjectMapOutput)
}

type SharedVPCHostProjectOutput struct{ *pulumi.OutputState }

func (SharedVPCHostProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedVPCHostProject)(nil))
}

func (o SharedVPCHostProjectOutput) ToSharedVPCHostProjectOutput() SharedVPCHostProjectOutput {
	return o
}

func (o SharedVPCHostProjectOutput) ToSharedVPCHostProjectOutputWithContext(ctx context.Context) SharedVPCHostProjectOutput {
	return o
}

func (o SharedVPCHostProjectOutput) ToSharedVPCHostProjectPtrOutput() SharedVPCHostProjectPtrOutput {
	return o.ToSharedVPCHostProjectPtrOutputWithContext(context.Background())
}

func (o SharedVPCHostProjectOutput) ToSharedVPCHostProjectPtrOutputWithContext(ctx context.Context) SharedVPCHostProjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedVPCHostProject) *SharedVPCHostProject {
		return &v
	}).(SharedVPCHostProjectPtrOutput)
}

type SharedVPCHostProjectPtrOutput struct{ *pulumi.OutputState }

func (SharedVPCHostProjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedVPCHostProject)(nil))
}

func (o SharedVPCHostProjectPtrOutput) ToSharedVPCHostProjectPtrOutput() SharedVPCHostProjectPtrOutput {
	return o
}

func (o SharedVPCHostProjectPtrOutput) ToSharedVPCHostProjectPtrOutputWithContext(ctx context.Context) SharedVPCHostProjectPtrOutput {
	return o
}

func (o SharedVPCHostProjectPtrOutput) Elem() SharedVPCHostProjectOutput {
	return o.ApplyT(func(v *SharedVPCHostProject) SharedVPCHostProject {
		if v != nil {
			return *v
		}
		var ret SharedVPCHostProject
		return ret
	}).(SharedVPCHostProjectOutput)
}

type SharedVPCHostProjectArrayOutput struct{ *pulumi.OutputState }

func (SharedVPCHostProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedVPCHostProject)(nil))
}

func (o SharedVPCHostProjectArrayOutput) ToSharedVPCHostProjectArrayOutput() SharedVPCHostProjectArrayOutput {
	return o
}

func (o SharedVPCHostProjectArrayOutput) ToSharedVPCHostProjectArrayOutputWithContext(ctx context.Context) SharedVPCHostProjectArrayOutput {
	return o
}

func (o SharedVPCHostProjectArrayOutput) Index(i pulumi.IntInput) SharedVPCHostProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedVPCHostProject {
		return vs[0].([]SharedVPCHostProject)[vs[1].(int)]
	}).(SharedVPCHostProjectOutput)
}

type SharedVPCHostProjectMapOutput struct{ *pulumi.OutputState }

func (SharedVPCHostProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SharedVPCHostProject)(nil))
}

func (o SharedVPCHostProjectMapOutput) ToSharedVPCHostProjectMapOutput() SharedVPCHostProjectMapOutput {
	return o
}

func (o SharedVPCHostProjectMapOutput) ToSharedVPCHostProjectMapOutputWithContext(ctx context.Context) SharedVPCHostProjectMapOutput {
	return o
}

func (o SharedVPCHostProjectMapOutput) MapIndex(k pulumi.StringInput) SharedVPCHostProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SharedVPCHostProject {
		return vs[0].(map[string]SharedVPCHostProject)[vs[1].(string)]
	}).(SharedVPCHostProjectOutput)
}

func init() {
	pulumi.RegisterOutputType(SharedVPCHostProjectOutput{})
	pulumi.RegisterOutputType(SharedVPCHostProjectPtrOutput{})
	pulumi.RegisterOutputType(SharedVPCHostProjectArrayOutput{})
	pulumi.RegisterOutputType(SharedVPCHostProjectMapOutput{})
}
