// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a Machine Image resource. Machine images store all the configuration,
// metadata, permissions, and data from one or more disks required to create a
// Virtual machine (VM) instance.
//
// To get more information about MachineImage, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/machineImages)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/machine-images)
//
// ## Example Usage
// ### Machine Image Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vm, err := compute.NewInstance(ctx, "vm", &compute.InstanceArgs{
//				MachineType: pulumi.String("e2-medium"),
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
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewMachineImage(ctx, "image", &compute.MachineImageArgs{
//				SourceInstance: vm.SelfLink,
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Compute Machine Image Kms
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vm, err := compute.NewInstance(ctx, "vm", &compute.InstanceArgs{
//				MachineType: pulumi.String("e2-medium"),
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
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			keyRing, err := kms.NewKeyRing(ctx, "keyRing", &kms.KeyRingArgs{
//				Location: pulumi.String("us"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			cryptoKey, err := kms.NewCryptoKey(ctx, "cryptoKey", &kms.CryptoKeyArgs{
//				KeyRing: keyRing.ID(),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewMachineImage(ctx, "image", &compute.MachineImageArgs{
//				SourceInstance: vm.SelfLink,
//				MachineImageEncryptionKey: &compute.MachineImageMachineImageEncryptionKeyArgs{
//					KmsKeyName: cryptoKey.ID(),
//				},
//			}, pulumi.Provider(google_beta))
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
// # MachineImage can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/machineImage:MachineImage default projects/{{project}}/global/machineImages/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/machineImage:MachineImage default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/machineImage:MachineImage default {{name}}
//
// ```
type MachineImage struct {
	pulumi.CustomResourceState

	// A text description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrOutput `pulumi:"guestFlush"`
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey MachineImageMachineImageEncryptionKeyPtrOutput `pulumi:"machineImageEncryptionKey"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	//
	// ***
	SourceInstance pulumi.StringOutput `pulumi:"sourceInstance"`
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
}

// NewMachineImage registers a new resource with the given unique name, arguments, and options.
func NewMachineImage(ctx *pulumi.Context,
	name string, args *MachineImageArgs, opts ...pulumi.ResourceOption) (*MachineImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceInstance == nil {
		return nil, errors.New("invalid value for required argument 'SourceInstance'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MachineImage
	err := ctx.RegisterResource("gcp:compute/machineImage:MachineImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineImage gets an existing MachineImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineImageState, opts ...pulumi.ResourceOption) (*MachineImage, error) {
	var resource MachineImage
	err := ctx.ReadResource("gcp:compute/machineImage:MachineImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineImage resources.
type machineImageState struct {
	// A text description of the resource.
	Description *string `pulumi:"description"`
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush *bool `pulumi:"guestFlush"`
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey *MachineImageMachineImageEncryptionKey `pulumi:"machineImageEncryptionKey"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	//
	// ***
	SourceInstance *string `pulumi:"sourceInstance"`
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations []string `pulumi:"storageLocations"`
}

type MachineImageState struct {
	// A text description of the resource.
	Description pulumi.StringPtrInput
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrInput
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey MachineImageMachineImageEncryptionKeyPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	//
	// ***
	SourceInstance pulumi.StringPtrInput
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations pulumi.StringArrayInput
}

func (MachineImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageState)(nil)).Elem()
}

type machineImageArgs struct {
	// A text description of the resource.
	Description *string `pulumi:"description"`
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush *bool `pulumi:"guestFlush"`
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey *MachineImageMachineImageEncryptionKey `pulumi:"machineImageEncryptionKey"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	//
	// ***
	SourceInstance string `pulumi:"sourceInstance"`
}

// The set of arguments for constructing a MachineImage resource.
type MachineImageArgs struct {
	// A text description of the resource.
	Description pulumi.StringPtrInput
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrInput
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey MachineImageMachineImageEncryptionKeyPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	//
	// ***
	SourceInstance pulumi.StringInput
}

func (MachineImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageArgs)(nil)).Elem()
}

type MachineImageInput interface {
	pulumi.Input

	ToMachineImageOutput() MachineImageOutput
	ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput
}

func (*MachineImage) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImage)(nil)).Elem()
}

func (i *MachineImage) ToMachineImageOutput() MachineImageOutput {
	return i.ToMachineImageOutputWithContext(context.Background())
}

func (i *MachineImage) ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageOutput)
}

func (i *MachineImage) ToOutput(ctx context.Context) pulumix.Output[*MachineImage] {
	return pulumix.Output[*MachineImage]{
		OutputState: i.ToMachineImageOutputWithContext(ctx).OutputState,
	}
}

// MachineImageArrayInput is an input type that accepts MachineImageArray and MachineImageArrayOutput values.
// You can construct a concrete instance of `MachineImageArrayInput` via:
//
//	MachineImageArray{ MachineImageArgs{...} }
type MachineImageArrayInput interface {
	pulumi.Input

	ToMachineImageArrayOutput() MachineImageArrayOutput
	ToMachineImageArrayOutputWithContext(context.Context) MachineImageArrayOutput
}

type MachineImageArray []MachineImageInput

func (MachineImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineImage)(nil)).Elem()
}

func (i MachineImageArray) ToMachineImageArrayOutput() MachineImageArrayOutput {
	return i.ToMachineImageArrayOutputWithContext(context.Background())
}

func (i MachineImageArray) ToMachineImageArrayOutputWithContext(ctx context.Context) MachineImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageArrayOutput)
}

func (i MachineImageArray) ToOutput(ctx context.Context) pulumix.Output[[]*MachineImage] {
	return pulumix.Output[[]*MachineImage]{
		OutputState: i.ToMachineImageArrayOutputWithContext(ctx).OutputState,
	}
}

// MachineImageMapInput is an input type that accepts MachineImageMap and MachineImageMapOutput values.
// You can construct a concrete instance of `MachineImageMapInput` via:
//
//	MachineImageMap{ "key": MachineImageArgs{...} }
type MachineImageMapInput interface {
	pulumi.Input

	ToMachineImageMapOutput() MachineImageMapOutput
	ToMachineImageMapOutputWithContext(context.Context) MachineImageMapOutput
}

type MachineImageMap map[string]MachineImageInput

func (MachineImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineImage)(nil)).Elem()
}

func (i MachineImageMap) ToMachineImageMapOutput() MachineImageMapOutput {
	return i.ToMachineImageMapOutputWithContext(context.Background())
}

func (i MachineImageMap) ToMachineImageMapOutputWithContext(ctx context.Context) MachineImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageMapOutput)
}

func (i MachineImageMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MachineImage] {
	return pulumix.Output[map[string]*MachineImage]{
		OutputState: i.ToMachineImageMapOutputWithContext(ctx).OutputState,
	}
}

type MachineImageOutput struct{ *pulumi.OutputState }

func (MachineImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImage)(nil)).Elem()
}

func (o MachineImageOutput) ToMachineImageOutput() MachineImageOutput {
	return o
}

func (o MachineImageOutput) ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput {
	return o
}

func (o MachineImageOutput) ToOutput(ctx context.Context) pulumix.Output[*MachineImage] {
	return pulumix.Output[*MachineImage]{
		OutputState: o.OutputState,
	}
}

// A text description of the resource.
func (o MachineImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
func (o MachineImageOutput) GuestFlush() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineImage) pulumi.BoolPtrOutput { return v.GuestFlush }).(pulumi.BoolPtrOutput)
}

// Encrypts the machine image using a customer-supplied encryption key.
// After you encrypt a machine image with a customer-supplied key, you must
// provide the same key if you use the machine image later (e.g. to create a
// instance from the image)
// Structure is documented below.
func (o MachineImageOutput) MachineImageEncryptionKey() MachineImageMachineImageEncryptionKeyPtrOutput {
	return o.ApplyT(func(v *MachineImage) MachineImageMachineImageEncryptionKeyPtrOutput {
		return v.MachineImageEncryptionKey
	}).(MachineImageMachineImageEncryptionKeyPtrOutput)
}

// Name of the resource.
func (o MachineImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o MachineImageOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImage) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o MachineImageOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImage) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
//
// ***
func (o MachineImageOutput) SourceInstance() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineImage) pulumi.StringOutput { return v.SourceInstance }).(pulumi.StringOutput)
}

// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
func (o MachineImageOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MachineImage) pulumi.StringArrayOutput { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

type MachineImageArrayOutput struct{ *pulumi.OutputState }

func (MachineImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineImage)(nil)).Elem()
}

func (o MachineImageArrayOutput) ToMachineImageArrayOutput() MachineImageArrayOutput {
	return o
}

func (o MachineImageArrayOutput) ToMachineImageArrayOutputWithContext(ctx context.Context) MachineImageArrayOutput {
	return o
}

func (o MachineImageArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MachineImage] {
	return pulumix.Output[[]*MachineImage]{
		OutputState: o.OutputState,
	}
}

func (o MachineImageArrayOutput) Index(i pulumi.IntInput) MachineImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineImage {
		return vs[0].([]*MachineImage)[vs[1].(int)]
	}).(MachineImageOutput)
}

type MachineImageMapOutput struct{ *pulumi.OutputState }

func (MachineImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineImage)(nil)).Elem()
}

func (o MachineImageMapOutput) ToMachineImageMapOutput() MachineImageMapOutput {
	return o
}

func (o MachineImageMapOutput) ToMachineImageMapOutputWithContext(ctx context.Context) MachineImageMapOutput {
	return o
}

func (o MachineImageMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MachineImage] {
	return pulumix.Output[map[string]*MachineImage]{
		OutputState: o.OutputState,
	}
}

func (o MachineImageMapOutput) MapIndex(k pulumi.StringInput) MachineImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineImage {
		return vs[0].(map[string]*MachineImage)[vs[1].(string)]
	}).(MachineImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineImageInput)(nil)).Elem(), &MachineImage{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineImageArrayInput)(nil)).Elem(), MachineImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineImageMapInput)(nil)).Elem(), MachineImageMap{})
	pulumi.RegisterOutputType(MachineImageOutput{})
	pulumi.RegisterOutputType(MachineImageArrayOutput{})
	pulumi.RegisterOutputType(MachineImageMapOutput{})
}
