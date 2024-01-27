// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Persistent disks can be attached to a compute instance using the `attachedDisk`
// section within the compute instance configuration.
// However there may be situations where managing the attached disks via the compute
// instance config isn't preferable or possible, such as attaching dynamic
// numbers of disks using the `count` variable.
//
// To get more information about attaching disks, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
// * How-to Guides
//   - [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
//
// **Note:** When using `compute.AttachedDisk` you **must** use `lifecycle.ignore_changes = ["attachedDisk"]` on the `compute.Instance` resource that has the disks attached. Otherwise the two resources will fight for control of the attached disk block.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultInstance, err := compute.NewInstance(ctx, "defaultInstance", &compute.InstanceArgs{
//				MachineType: pulumi.String("e2-medium"),
//				Zone:        pulumi.String("us-west1-a"),
//				BootDisk: &compute.InstanceBootDiskArgs{
//					InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
//						Image: pulumi.String("debian-cloud/debian-11"),
//					},
//				},
//				NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
//					&compute.InstanceNetworkInterfaceArgs{
//						Network: pulumi.String("default"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewAttachedDisk(ctx, "defaultAttachedDisk", &compute.AttachedDiskArgs{
//				Disk:     pulumi.Any(google_compute_disk.Default.Id),
//				Instance: defaultInstance.ID(),
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
// Attached Disk can be imported the following ways* `projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}` * `{{project}}/{{zone}}/{{instance.name}}/{{disk.name}}` When using the `pulumi import` command, Attached Disk can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:compute/attachedDisk:AttachedDisk default projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/attachedDisk:AttachedDisk default {{project}}/{{zone}}/{{instance.name}}/{{disk.name}}
//
// ```
type AttachedDisk struct {
	pulumi.CustomResourceState

	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	//
	// If not specified, the server chooses a default device name to apply
	// to this disk, in the form persistent-disks-x, where x is a number
	// assigned by Google Compute Engine.
	DeviceName pulumi.StringOutput `pulumi:"deviceName"`
	// `name` or `selfLink` of the disk that will be attached.
	//
	// ***
	Disk pulumi.StringOutput `pulumi:"disk"`
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	//
	// Possible values:
	// "READ_ONLY"
	// "READ_WRITE"
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project pulumi.StringOutput `pulumi:"project"`
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewAttachedDisk registers a new resource with the given unique name, arguments, and options.
func NewAttachedDisk(ctx *pulumi.Context,
	name string, args *AttachedDiskArgs, opts ...pulumi.ResourceOption) (*AttachedDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disk == nil {
		return nil, errors.New("invalid value for required argument 'Disk'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AttachedDisk
	err := ctx.RegisterResource("gcp:compute/attachedDisk:AttachedDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDisk gets an existing AttachedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDiskState, opts ...pulumi.ResourceOption) (*AttachedDisk, error) {
	var resource AttachedDisk
	err := ctx.ReadResource("gcp:compute/attachedDisk:AttachedDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDisk resources.
type attachedDiskState struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	//
	// If not specified, the server chooses a default device name to apply
	// to this disk, in the form persistent-disks-x, where x is a number
	// assigned by Google Compute Engine.
	DeviceName *string `pulumi:"deviceName"`
	// `name` or `selfLink` of the disk that will be attached.
	//
	// ***
	Disk *string `pulumi:"disk"`
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance *string `pulumi:"instance"`
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	//
	// Possible values:
	// "READ_ONLY"
	// "READ_WRITE"
	Mode *string `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project *string `pulumi:"project"`
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone *string `pulumi:"zone"`
}

type AttachedDiskState struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	//
	// If not specified, the server chooses a default device name to apply
	// to this disk, in the form persistent-disks-x, where x is a number
	// assigned by Google Compute Engine.
	DeviceName pulumi.StringPtrInput
	// `name` or `selfLink` of the disk that will be attached.
	//
	// ***
	Disk pulumi.StringPtrInput
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance pulumi.StringPtrInput
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	//
	// Possible values:
	// "READ_ONLY"
	// "READ_WRITE"
	Mode pulumi.StringPtrInput
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project pulumi.StringPtrInput
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone pulumi.StringPtrInput
}

func (AttachedDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDiskState)(nil)).Elem()
}

type attachedDiskArgs struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	//
	// If not specified, the server chooses a default device name to apply
	// to this disk, in the form persistent-disks-x, where x is a number
	// assigned by Google Compute Engine.
	DeviceName *string `pulumi:"deviceName"`
	// `name` or `selfLink` of the disk that will be attached.
	//
	// ***
	Disk string `pulumi:"disk"`
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance string `pulumi:"instance"`
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	//
	// Possible values:
	// "READ_ONLY"
	// "READ_WRITE"
	Mode *string `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project *string `pulumi:"project"`
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a AttachedDisk resource.
type AttachedDiskArgs struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	//
	// If not specified, the server chooses a default device name to apply
	// to this disk, in the form persistent-disks-x, where x is a number
	// assigned by Google Compute Engine.
	DeviceName pulumi.StringPtrInput
	// `name` or `selfLink` of the disk that will be attached.
	//
	// ***
	Disk pulumi.StringInput
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance pulumi.StringInput
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	//
	// Possible values:
	// "READ_ONLY"
	// "READ_WRITE"
	Mode pulumi.StringPtrInput
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project pulumi.StringPtrInput
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone pulumi.StringPtrInput
}

func (AttachedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDiskArgs)(nil)).Elem()
}

type AttachedDiskInput interface {
	pulumi.Input

	ToAttachedDiskOutput() AttachedDiskOutput
	ToAttachedDiskOutputWithContext(ctx context.Context) AttachedDiskOutput
}

func (*AttachedDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDisk)(nil)).Elem()
}

func (i *AttachedDisk) ToAttachedDiskOutput() AttachedDiskOutput {
	return i.ToAttachedDiskOutputWithContext(context.Background())
}

func (i *AttachedDisk) ToAttachedDiskOutputWithContext(ctx context.Context) AttachedDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDiskOutput)
}

// AttachedDiskArrayInput is an input type that accepts AttachedDiskArray and AttachedDiskArrayOutput values.
// You can construct a concrete instance of `AttachedDiskArrayInput` via:
//
//	AttachedDiskArray{ AttachedDiskArgs{...} }
type AttachedDiskArrayInput interface {
	pulumi.Input

	ToAttachedDiskArrayOutput() AttachedDiskArrayOutput
	ToAttachedDiskArrayOutputWithContext(context.Context) AttachedDiskArrayOutput
}

type AttachedDiskArray []AttachedDiskInput

func (AttachedDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttachedDisk)(nil)).Elem()
}

func (i AttachedDiskArray) ToAttachedDiskArrayOutput() AttachedDiskArrayOutput {
	return i.ToAttachedDiskArrayOutputWithContext(context.Background())
}

func (i AttachedDiskArray) ToAttachedDiskArrayOutputWithContext(ctx context.Context) AttachedDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDiskArrayOutput)
}

// AttachedDiskMapInput is an input type that accepts AttachedDiskMap and AttachedDiskMapOutput values.
// You can construct a concrete instance of `AttachedDiskMapInput` via:
//
//	AttachedDiskMap{ "key": AttachedDiskArgs{...} }
type AttachedDiskMapInput interface {
	pulumi.Input

	ToAttachedDiskMapOutput() AttachedDiskMapOutput
	ToAttachedDiskMapOutputWithContext(context.Context) AttachedDiskMapOutput
}

type AttachedDiskMap map[string]AttachedDiskInput

func (AttachedDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttachedDisk)(nil)).Elem()
}

func (i AttachedDiskMap) ToAttachedDiskMapOutput() AttachedDiskMapOutput {
	return i.ToAttachedDiskMapOutputWithContext(context.Background())
}

func (i AttachedDiskMap) ToAttachedDiskMapOutputWithContext(ctx context.Context) AttachedDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDiskMapOutput)
}

type AttachedDiskOutput struct{ *pulumi.OutputState }

func (AttachedDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDisk)(nil)).Elem()
}

func (o AttachedDiskOutput) ToAttachedDiskOutput() AttachedDiskOutput {
	return o
}

func (o AttachedDiskOutput) ToAttachedDiskOutputWithContext(ctx context.Context) AttachedDiskOutput {
	return o
}

// Specifies a unique device name of your choice that is
// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
// system running within the instance. This name can be used to
// reference the device for mounting, resizing, and so on, from within
// the instance.
//
// If not specified, the server chooses a default device name to apply
// to this disk, in the form persistent-disks-x, where x is a number
// assigned by Google Compute Engine.
func (o AttachedDiskOutput) DeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDisk) pulumi.StringOutput { return v.DeviceName }).(pulumi.StringOutput)
}

// `name` or `selfLink` of the disk that will be attached.
//
// ***
func (o AttachedDiskOutput) Disk() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDisk) pulumi.StringOutput { return v.Disk }).(pulumi.StringOutput)
}

// `name` or `selfLink` of the compute instance that the disk will be attached to.
// If the `selfLink` is provided then `zone` and `project` are extracted from the
// self link. If only the name is used then `zone` and `project` must be defined
// as properties on the resource or provider.
func (o AttachedDiskOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDisk) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The mode in which to attach this disk, either READ_WRITE or
// READ_ONLY. If not specified, the default is to attach the disk in
// READ_WRITE mode.
//
// Possible values:
// "READ_ONLY"
// "READ_WRITE"
func (o AttachedDiskOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDisk) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The project that the referenced compute instance is a part of. If `instance` is referenced by its
// `selfLink` the project defined in the link will take precedence.
func (o AttachedDiskOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDisk) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The zone that the referenced compute instance is located within. If `instance` is referenced by its
// `selfLink` the zone defined in the link will take precedence.
func (o AttachedDiskOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDisk) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type AttachedDiskArrayOutput struct{ *pulumi.OutputState }

func (AttachedDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttachedDisk)(nil)).Elem()
}

func (o AttachedDiskArrayOutput) ToAttachedDiskArrayOutput() AttachedDiskArrayOutput {
	return o
}

func (o AttachedDiskArrayOutput) ToAttachedDiskArrayOutputWithContext(ctx context.Context) AttachedDiskArrayOutput {
	return o
}

func (o AttachedDiskArrayOutput) Index(i pulumi.IntInput) AttachedDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AttachedDisk {
		return vs[0].([]*AttachedDisk)[vs[1].(int)]
	}).(AttachedDiskOutput)
}

type AttachedDiskMapOutput struct{ *pulumi.OutputState }

func (AttachedDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttachedDisk)(nil)).Elem()
}

func (o AttachedDiskMapOutput) ToAttachedDiskMapOutput() AttachedDiskMapOutput {
	return o
}

func (o AttachedDiskMapOutput) ToAttachedDiskMapOutputWithContext(ctx context.Context) AttachedDiskMapOutput {
	return o
}

func (o AttachedDiskMapOutput) MapIndex(k pulumi.StringInput) AttachedDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AttachedDisk {
		return vs[0].(map[string]*AttachedDisk)[vs[1].(string)]
	}).(AttachedDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttachedDiskInput)(nil)).Elem(), &AttachedDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachedDiskArrayInput)(nil)).Elem(), AttachedDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachedDiskMapInput)(nil)).Elem(), AttachedDiskMap{})
	pulumi.RegisterOutputType(AttachedDiskOutput{})
	pulumi.RegisterOutputType(AttachedDiskArrayOutput{})
	pulumi.RegisterOutputType(AttachedDiskMapOutput{})
}
