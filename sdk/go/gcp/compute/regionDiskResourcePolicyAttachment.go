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

// Adds existing resource policies to a disk. You can only add one policy
// which will be applied to this disk for scheduling snapshot creation.
//
// > **Note:** This resource does not support zonal disks (`compute.Disk`). For zonal disks, please refer to the `compute.DiskResourcePolicyAttachment` resource.
//
// ## Example Usage
// ### Region Disk Resource Policy Attachment Basic
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
//			disk, err := compute.NewDisk(ctx, "disk", &compute.DiskArgs{
//				Image: pulumi.String("debian-cloud/debian-11"),
//				Size:  pulumi.Int(50),
//				Type:  pulumi.String("pd-ssd"),
//				Zone:  pulumi.String("us-central1-a"),
//			})
//			if err != nil {
//				return err
//			}
//			snapdisk, err := compute.NewSnapshot(ctx, "snapdisk", &compute.SnapshotArgs{
//				SourceDisk: disk.Name,
//				Zone:       pulumi.String("us-central1-a"),
//			})
//			if err != nil {
//				return err
//			}
//			ssd, err := compute.NewRegionDisk(ctx, "ssd", &compute.RegionDiskArgs{
//				ReplicaZones: pulumi.StringArray{
//					pulumi.String("us-central1-a"),
//					pulumi.String("us-central1-f"),
//				},
//				Snapshot: snapdisk.ID(),
//				Size:     pulumi.Int(50),
//				Type:     pulumi.String("pd-ssd"),
//				Region:   pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewRegionDiskResourcePolicyAttachment(ctx, "attachment", &compute.RegionDiskResourcePolicyAttachmentArgs{
//				Disk:   ssd.Name,
//				Region: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewResourcePolicy(ctx, "policy", &compute.ResourcePolicyArgs{
//				Region: pulumi.String("us-central1"),
//				SnapshotSchedulePolicy: &compute.ResourcePolicySnapshotSchedulePolicyArgs{
//					Schedule: &compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs{
//						DailySchedule: &compute.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs{
//							DaysInCycle: pulumi.Int(1),
//							StartTime:   pulumi.String("04:00"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.LookupImage(ctx, &compute.LookupImageArgs{
//				Family:  pulumi.StringRef("debian-11"),
//				Project: pulumi.StringRef("debian-cloud"),
//			}, nil)
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
// # RegionDiskResourcePolicyAttachment can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default projects/{{project}}/regions/{{region}}/disks/{{disk}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{project}}/{{region}}/{{disk}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{region}}/{{disk}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{disk}}/{{name}}
//
// ```
type RegionDiskResourcePolicyAttachment struct {
	pulumi.CustomResourceState

	// The name of the regional disk in which the resource policies are attached to.
	//
	// ***
	Disk pulumi.StringOutput `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewRegionDiskResourcePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewRegionDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, args *RegionDiskResourcePolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*RegionDiskResourcePolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disk == nil {
		return nil, errors.New("invalid value for required argument 'Disk'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegionDiskResourcePolicyAttachment
	err := ctx.RegisterResource("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionDiskResourcePolicyAttachment gets an existing RegionDiskResourcePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionDiskResourcePolicyAttachmentState, opts ...pulumi.ResourceOption) (*RegionDiskResourcePolicyAttachment, error) {
	var resource RegionDiskResourcePolicyAttachment
	err := ctx.ReadResource("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionDiskResourcePolicyAttachment resources.
type regionDiskResourcePolicyAttachmentState struct {
	// The name of the regional disk in which the resource policies are attached to.
	//
	// ***
	Disk *string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
}

type RegionDiskResourcePolicyAttachmentState struct {
	// The name of the regional disk in which the resource policies are attached to.
	//
	// ***
	Disk pulumi.StringPtrInput
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
}

func (RegionDiskResourcePolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskResourcePolicyAttachmentState)(nil)).Elem()
}

type regionDiskResourcePolicyAttachmentArgs struct {
	// The name of the regional disk in which the resource policies are attached to.
	//
	// ***
	Disk string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RegionDiskResourcePolicyAttachment resource.
type RegionDiskResourcePolicyAttachmentArgs struct {
	// The name of the regional disk in which the resource policies are attached to.
	//
	// ***
	Disk pulumi.StringInput
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
}

func (RegionDiskResourcePolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskResourcePolicyAttachmentArgs)(nil)).Elem()
}

type RegionDiskResourcePolicyAttachmentInput interface {
	pulumi.Input

	ToRegionDiskResourcePolicyAttachmentOutput() RegionDiskResourcePolicyAttachmentOutput
	ToRegionDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) RegionDiskResourcePolicyAttachmentOutput
}

func (*RegionDiskResourcePolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (i *RegionDiskResourcePolicyAttachment) ToRegionDiskResourcePolicyAttachmentOutput() RegionDiskResourcePolicyAttachmentOutput {
	return i.ToRegionDiskResourcePolicyAttachmentOutputWithContext(context.Background())
}

func (i *RegionDiskResourcePolicyAttachment) ToRegionDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) RegionDiskResourcePolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskResourcePolicyAttachmentOutput)
}

func (i *RegionDiskResourcePolicyAttachment) ToOutput(ctx context.Context) pulumix.Output[*RegionDiskResourcePolicyAttachment] {
	return pulumix.Output[*RegionDiskResourcePolicyAttachment]{
		OutputState: i.ToRegionDiskResourcePolicyAttachmentOutputWithContext(ctx).OutputState,
	}
}

// RegionDiskResourcePolicyAttachmentArrayInput is an input type that accepts RegionDiskResourcePolicyAttachmentArray and RegionDiskResourcePolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `RegionDiskResourcePolicyAttachmentArrayInput` via:
//
//	RegionDiskResourcePolicyAttachmentArray{ RegionDiskResourcePolicyAttachmentArgs{...} }
type RegionDiskResourcePolicyAttachmentArrayInput interface {
	pulumi.Input

	ToRegionDiskResourcePolicyAttachmentArrayOutput() RegionDiskResourcePolicyAttachmentArrayOutput
	ToRegionDiskResourcePolicyAttachmentArrayOutputWithContext(context.Context) RegionDiskResourcePolicyAttachmentArrayOutput
}

type RegionDiskResourcePolicyAttachmentArray []RegionDiskResourcePolicyAttachmentInput

func (RegionDiskResourcePolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (i RegionDiskResourcePolicyAttachmentArray) ToRegionDiskResourcePolicyAttachmentArrayOutput() RegionDiskResourcePolicyAttachmentArrayOutput {
	return i.ToRegionDiskResourcePolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i RegionDiskResourcePolicyAttachmentArray) ToRegionDiskResourcePolicyAttachmentArrayOutputWithContext(ctx context.Context) RegionDiskResourcePolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskResourcePolicyAttachmentArrayOutput)
}

func (i RegionDiskResourcePolicyAttachmentArray) ToOutput(ctx context.Context) pulumix.Output[[]*RegionDiskResourcePolicyAttachment] {
	return pulumix.Output[[]*RegionDiskResourcePolicyAttachment]{
		OutputState: i.ToRegionDiskResourcePolicyAttachmentArrayOutputWithContext(ctx).OutputState,
	}
}

// RegionDiskResourcePolicyAttachmentMapInput is an input type that accepts RegionDiskResourcePolicyAttachmentMap and RegionDiskResourcePolicyAttachmentMapOutput values.
// You can construct a concrete instance of `RegionDiskResourcePolicyAttachmentMapInput` via:
//
//	RegionDiskResourcePolicyAttachmentMap{ "key": RegionDiskResourcePolicyAttachmentArgs{...} }
type RegionDiskResourcePolicyAttachmentMapInput interface {
	pulumi.Input

	ToRegionDiskResourcePolicyAttachmentMapOutput() RegionDiskResourcePolicyAttachmentMapOutput
	ToRegionDiskResourcePolicyAttachmentMapOutputWithContext(context.Context) RegionDiskResourcePolicyAttachmentMapOutput
}

type RegionDiskResourcePolicyAttachmentMap map[string]RegionDiskResourcePolicyAttachmentInput

func (RegionDiskResourcePolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (i RegionDiskResourcePolicyAttachmentMap) ToRegionDiskResourcePolicyAttachmentMapOutput() RegionDiskResourcePolicyAttachmentMapOutput {
	return i.ToRegionDiskResourcePolicyAttachmentMapOutputWithContext(context.Background())
}

func (i RegionDiskResourcePolicyAttachmentMap) ToRegionDiskResourcePolicyAttachmentMapOutputWithContext(ctx context.Context) RegionDiskResourcePolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskResourcePolicyAttachmentMapOutput)
}

func (i RegionDiskResourcePolicyAttachmentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RegionDiskResourcePolicyAttachment] {
	return pulumix.Output[map[string]*RegionDiskResourcePolicyAttachment]{
		OutputState: i.ToRegionDiskResourcePolicyAttachmentMapOutputWithContext(ctx).OutputState,
	}
}

type RegionDiskResourcePolicyAttachmentOutput struct{ *pulumi.OutputState }

func (RegionDiskResourcePolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (o RegionDiskResourcePolicyAttachmentOutput) ToRegionDiskResourcePolicyAttachmentOutput() RegionDiskResourcePolicyAttachmentOutput {
	return o
}

func (o RegionDiskResourcePolicyAttachmentOutput) ToRegionDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) RegionDiskResourcePolicyAttachmentOutput {
	return o
}

func (o RegionDiskResourcePolicyAttachmentOutput) ToOutput(ctx context.Context) pulumix.Output[*RegionDiskResourcePolicyAttachment] {
	return pulumix.Output[*RegionDiskResourcePolicyAttachment]{
		OutputState: o.OutputState,
	}
}

// The name of the regional disk in which the resource policies are attached to.
//
// ***
func (o RegionDiskResourcePolicyAttachmentOutput) Disk() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Disk }).(pulumi.StringOutput)
}

// The resource policy to be attached to the disk for scheduling snapshot
// creation. Do not specify the self link.
func (o RegionDiskResourcePolicyAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RegionDiskResourcePolicyAttachmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A reference to the region where the disk resides.
func (o RegionDiskResourcePolicyAttachmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type RegionDiskResourcePolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (RegionDiskResourcePolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (o RegionDiskResourcePolicyAttachmentArrayOutput) ToRegionDiskResourcePolicyAttachmentArrayOutput() RegionDiskResourcePolicyAttachmentArrayOutput {
	return o
}

func (o RegionDiskResourcePolicyAttachmentArrayOutput) ToRegionDiskResourcePolicyAttachmentArrayOutputWithContext(ctx context.Context) RegionDiskResourcePolicyAttachmentArrayOutput {
	return o
}

func (o RegionDiskResourcePolicyAttachmentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RegionDiskResourcePolicyAttachment] {
	return pulumix.Output[[]*RegionDiskResourcePolicyAttachment]{
		OutputState: o.OutputState,
	}
}

func (o RegionDiskResourcePolicyAttachmentArrayOutput) Index(i pulumi.IntInput) RegionDiskResourcePolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionDiskResourcePolicyAttachment {
		return vs[0].([]*RegionDiskResourcePolicyAttachment)[vs[1].(int)]
	}).(RegionDiskResourcePolicyAttachmentOutput)
}

type RegionDiskResourcePolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (RegionDiskResourcePolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (o RegionDiskResourcePolicyAttachmentMapOutput) ToRegionDiskResourcePolicyAttachmentMapOutput() RegionDiskResourcePolicyAttachmentMapOutput {
	return o
}

func (o RegionDiskResourcePolicyAttachmentMapOutput) ToRegionDiskResourcePolicyAttachmentMapOutputWithContext(ctx context.Context) RegionDiskResourcePolicyAttachmentMapOutput {
	return o
}

func (o RegionDiskResourcePolicyAttachmentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RegionDiskResourcePolicyAttachment] {
	return pulumix.Output[map[string]*RegionDiskResourcePolicyAttachment]{
		OutputState: o.OutputState,
	}
}

func (o RegionDiskResourcePolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) RegionDiskResourcePolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionDiskResourcePolicyAttachment {
		return vs[0].(map[string]*RegionDiskResourcePolicyAttachment)[vs[1].(string)]
	}).(RegionDiskResourcePolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionDiskResourcePolicyAttachmentInput)(nil)).Elem(), &RegionDiskResourcePolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionDiskResourcePolicyAttachmentArrayInput)(nil)).Elem(), RegionDiskResourcePolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionDiskResourcePolicyAttachmentMapInput)(nil)).Elem(), RegionDiskResourcePolicyAttachmentMap{})
	pulumi.RegisterOutputType(RegionDiskResourcePolicyAttachmentOutput{})
	pulumi.RegisterOutputType(RegionDiskResourcePolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(RegionDiskResourcePolicyAttachmentMapOutput{})
}
