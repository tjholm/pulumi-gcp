// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Cloud AI Platform Notebook runtime.
//
// > **Note:** Due to limitations of the Notebooks Runtime API, many fields
// in this resource do not properly detect drift. These fields will also not
// appear in state once imported.
//
// To get more information about Runtime, see:
//
// * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
//
// ## Example Usage
// ### Notebook Runtime Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewRuntime(ctx, "runtime", &notebooks.RuntimeArgs{
// 			AccessConfig: &notebooks.RuntimeAccessConfigArgs{
// 				AccessType:   pulumi.String("SINGLE_USER"),
// 				RuntimeOwner: pulumi.String("admin@hashicorptest.com"),
// 			},
// 			Location: pulumi.String("us-central1"),
// 			VirtualMachine: &notebooks.RuntimeVirtualMachineArgs{
// 				VirtualMachineConfig: &notebooks.RuntimeVirtualMachineVirtualMachineConfigArgs{
// 					DataDisk: &notebooks.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs{
// 						InitializeParams: &notebooks.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs{
// 							DiskSizeGb: pulumi.Int(100),
// 							DiskType:   pulumi.String("PD_STANDARD"),
// 						},
// 					},
// 					MachineType: pulumi.String("n1-standard-4"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Notebook Runtime Basic Gpu
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewRuntime(ctx, "runtimeGpu", &notebooks.RuntimeArgs{
// 			AccessConfig: &notebooks.RuntimeAccessConfigArgs{
// 				AccessType:   pulumi.String("SINGLE_USER"),
// 				RuntimeOwner: pulumi.String("admin@hashicorptest.com"),
// 			},
// 			Location: pulumi.String("us-central1"),
// 			SoftwareConfig: &notebooks.RuntimeSoftwareConfigArgs{
// 				InstallGpuDriver: pulumi.Bool(true),
// 			},
// 			VirtualMachine: &notebooks.RuntimeVirtualMachineArgs{
// 				VirtualMachineConfig: &notebooks.RuntimeVirtualMachineVirtualMachineConfigArgs{
// 					AcceleratorConfig: &notebooks.RuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigArgs{
// 						CoreCount: pulumi.Int(1),
// 						Type:      pulumi.String("NVIDIA_TESLA_V100"),
// 					},
// 					DataDisk: &notebooks.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs{
// 						InitializeParams: &notebooks.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs{
// 							DiskSizeGb: pulumi.Int(100),
// 							DiskType:   pulumi.String("PD_STANDARD"),
// 						},
// 					},
// 					MachineType: pulumi.String("n1-standard-4"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Notebook Runtime Basic Container
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewRuntime(ctx, "runtimeContainer", &notebooks.RuntimeArgs{
// 			AccessConfig: &notebooks.RuntimeAccessConfigArgs{
// 				AccessType:   pulumi.String("SINGLE_USER"),
// 				RuntimeOwner: pulumi.String("admin@hashicorptest.com"),
// 			},
// 			Location: pulumi.String("us-central1"),
// 			VirtualMachine: &notebooks.RuntimeVirtualMachineArgs{
// 				VirtualMachineConfig: &notebooks.RuntimeVirtualMachineVirtualMachineConfigArgs{
// 					ContainerImages: notebooks.RuntimeVirtualMachineVirtualMachineConfigContainerImageArray{
// 						&notebooks.RuntimeVirtualMachineVirtualMachineConfigContainerImageArgs{
// 							Repository: pulumi.String("gcr.io/deeplearning-platform-release/base-cpu"),
// 							Tag:        pulumi.String("latest"),
// 						},
// 						&notebooks.RuntimeVirtualMachineVirtualMachineConfigContainerImageArgs{
// 							Repository: pulumi.String("gcr.io/deeplearning-platform-release/beam-notebooks"),
// 							Tag:        pulumi.String("latest"),
// 						},
// 					},
// 					DataDisk: &notebooks.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs{
// 						InitializeParams: &notebooks.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs{
// 							DiskSizeGb: pulumi.Int(100),
// 							DiskType:   pulumi.String("PD_STANDARD"),
// 						},
// 					},
// 					MachineType: pulumi.String("n1-standard-4"),
// 				},
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
// Runtime can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:notebooks/runtime:Runtime default projects/{{project}}/locations/{{location}}/runtimes/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:notebooks/runtime:Runtime default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:notebooks/runtime:Runtime default {{location}}/{{name}}
// ```
type Runtime struct {
	pulumi.CustomResourceState

	// The config settings for accessing runtime.
	// Structure is documented below.
	AccessConfig RuntimeAccessConfigPtrOutput `pulumi:"accessConfig"`
	// The health state of this runtime. For a list of possible output values, see
	// 'https://cloud.google.com/vertex-ai/docs/workbench/ reference/rest/v1/projects.locations.runtimes#healthstate'.
	HealthState pulumi.StringOutput `pulumi:"healthState"`
	// A reference to the zone where the machine resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// Contains Runtime daemon metrics such as Service status and JupyterLab status
	Metrics RuntimeMetricArrayOutput `pulumi:"metrics"`
	// The name specified for the Notebook instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The config settings for software inside the runtime.
	// Structure is documented below.
	SoftwareConfig RuntimeSoftwareConfigOutput `pulumi:"softwareConfig"`
	// The state of this runtime.
	State pulumi.StringOutput `pulumi:"state"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	// Structure is documented below.
	VirtualMachine RuntimeVirtualMachinePtrOutput `pulumi:"virtualMachine"`
}

// NewRuntime registers a new resource with the given unique name, arguments, and options.
func NewRuntime(ctx *pulumi.Context,
	name string, args *RuntimeArgs, opts ...pulumi.ResourceOption) (*Runtime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Runtime
	err := ctx.RegisterResource("gcp:notebooks/runtime:Runtime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntime gets an existing Runtime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeState, opts ...pulumi.ResourceOption) (*Runtime, error) {
	var resource Runtime
	err := ctx.ReadResource("gcp:notebooks/runtime:Runtime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Runtime resources.
type runtimeState struct {
	// The config settings for accessing runtime.
	// Structure is documented below.
	AccessConfig *RuntimeAccessConfig `pulumi:"accessConfig"`
	// The health state of this runtime. For a list of possible output values, see
	// 'https://cloud.google.com/vertex-ai/docs/workbench/ reference/rest/v1/projects.locations.runtimes#healthstate'.
	HealthState *string `pulumi:"healthState"`
	// A reference to the zone where the machine resides.
	Location *string `pulumi:"location"`
	// Contains Runtime daemon metrics such as Service status and JupyterLab status
	Metrics []RuntimeMetric `pulumi:"metrics"`
	// The name specified for the Notebook instance.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The config settings for software inside the runtime.
	// Structure is documented below.
	SoftwareConfig *RuntimeSoftwareConfig `pulumi:"softwareConfig"`
	// The state of this runtime.
	State *string `pulumi:"state"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	// Structure is documented below.
	VirtualMachine *RuntimeVirtualMachine `pulumi:"virtualMachine"`
}

type RuntimeState struct {
	// The config settings for accessing runtime.
	// Structure is documented below.
	AccessConfig RuntimeAccessConfigPtrInput
	// The health state of this runtime. For a list of possible output values, see
	// 'https://cloud.google.com/vertex-ai/docs/workbench/ reference/rest/v1/projects.locations.runtimes#healthstate'.
	HealthState pulumi.StringPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringPtrInput
	// Contains Runtime daemon metrics such as Service status and JupyterLab status
	Metrics RuntimeMetricArrayInput
	// The name specified for the Notebook instance.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The config settings for software inside the runtime.
	// Structure is documented below.
	SoftwareConfig RuntimeSoftwareConfigPtrInput
	// The state of this runtime.
	State pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the managed notebook instance.
	// Structure is documented below.
	VirtualMachine RuntimeVirtualMachinePtrInput
}

func (RuntimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeState)(nil)).Elem()
}

type runtimeArgs struct {
	// The config settings for accessing runtime.
	// Structure is documented below.
	AccessConfig *RuntimeAccessConfig `pulumi:"accessConfig"`
	// A reference to the zone where the machine resides.
	Location string `pulumi:"location"`
	// The name specified for the Notebook instance.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The config settings for software inside the runtime.
	// Structure is documented below.
	SoftwareConfig *RuntimeSoftwareConfig `pulumi:"softwareConfig"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	// Structure is documented below.
	VirtualMachine *RuntimeVirtualMachine `pulumi:"virtualMachine"`
}

// The set of arguments for constructing a Runtime resource.
type RuntimeArgs struct {
	// The config settings for accessing runtime.
	// Structure is documented below.
	AccessConfig RuntimeAccessConfigPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringInput
	// The name specified for the Notebook instance.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The config settings for software inside the runtime.
	// Structure is documented below.
	SoftwareConfig RuntimeSoftwareConfigPtrInput
	// Use a Compute Engine VM image to start the managed notebook instance.
	// Structure is documented below.
	VirtualMachine RuntimeVirtualMachinePtrInput
}

func (RuntimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeArgs)(nil)).Elem()
}

type RuntimeInput interface {
	pulumi.Input

	ToRuntimeOutput() RuntimeOutput
	ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput
}

func (*Runtime) ElementType() reflect.Type {
	return reflect.TypeOf((**Runtime)(nil)).Elem()
}

func (i *Runtime) ToRuntimeOutput() RuntimeOutput {
	return i.ToRuntimeOutputWithContext(context.Background())
}

func (i *Runtime) ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeOutput)
}

// RuntimeArrayInput is an input type that accepts RuntimeArray and RuntimeArrayOutput values.
// You can construct a concrete instance of `RuntimeArrayInput` via:
//
//          RuntimeArray{ RuntimeArgs{...} }
type RuntimeArrayInput interface {
	pulumi.Input

	ToRuntimeArrayOutput() RuntimeArrayOutput
	ToRuntimeArrayOutputWithContext(context.Context) RuntimeArrayOutput
}

type RuntimeArray []RuntimeInput

func (RuntimeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Runtime)(nil)).Elem()
}

func (i RuntimeArray) ToRuntimeArrayOutput() RuntimeArrayOutput {
	return i.ToRuntimeArrayOutputWithContext(context.Background())
}

func (i RuntimeArray) ToRuntimeArrayOutputWithContext(ctx context.Context) RuntimeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeArrayOutput)
}

// RuntimeMapInput is an input type that accepts RuntimeMap and RuntimeMapOutput values.
// You can construct a concrete instance of `RuntimeMapInput` via:
//
//          RuntimeMap{ "key": RuntimeArgs{...} }
type RuntimeMapInput interface {
	pulumi.Input

	ToRuntimeMapOutput() RuntimeMapOutput
	ToRuntimeMapOutputWithContext(context.Context) RuntimeMapOutput
}

type RuntimeMap map[string]RuntimeInput

func (RuntimeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Runtime)(nil)).Elem()
}

func (i RuntimeMap) ToRuntimeMapOutput() RuntimeMapOutput {
	return i.ToRuntimeMapOutputWithContext(context.Background())
}

func (i RuntimeMap) ToRuntimeMapOutputWithContext(ctx context.Context) RuntimeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeMapOutput)
}

type RuntimeOutput struct{ *pulumi.OutputState }

func (RuntimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Runtime)(nil)).Elem()
}

func (o RuntimeOutput) ToRuntimeOutput() RuntimeOutput {
	return o
}

func (o RuntimeOutput) ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput {
	return o
}

type RuntimeArrayOutput struct{ *pulumi.OutputState }

func (RuntimeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Runtime)(nil)).Elem()
}

func (o RuntimeArrayOutput) ToRuntimeArrayOutput() RuntimeArrayOutput {
	return o
}

func (o RuntimeArrayOutput) ToRuntimeArrayOutputWithContext(ctx context.Context) RuntimeArrayOutput {
	return o
}

func (o RuntimeArrayOutput) Index(i pulumi.IntInput) RuntimeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Runtime {
		return vs[0].([]*Runtime)[vs[1].(int)]
	}).(RuntimeOutput)
}

type RuntimeMapOutput struct{ *pulumi.OutputState }

func (RuntimeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Runtime)(nil)).Elem()
}

func (o RuntimeMapOutput) ToRuntimeMapOutput() RuntimeMapOutput {
	return o
}

func (o RuntimeMapOutput) ToRuntimeMapOutputWithContext(ctx context.Context) RuntimeMapOutput {
	return o
}

func (o RuntimeMapOutput) MapIndex(k pulumi.StringInput) RuntimeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Runtime {
		return vs[0].(map[string]*Runtime)[vs[1].(string)]
	}).(RuntimeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeInput)(nil)).Elem(), &Runtime{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeArrayInput)(nil)).Elem(), RuntimeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeMapInput)(nil)).Elem(), RuntimeMap{})
	pulumi.RegisterOutputType(RuntimeOutput{})
	pulumi.RegisterOutputType(RuntimeArrayOutput{})
	pulumi.RegisterOutputType(RuntimeMapOutput{})
}
