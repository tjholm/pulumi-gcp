// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a Persistent Disk Snapshot resource.
//
// Use snapshots to back up data from your persistent disks. Snapshots are
// different from public images and custom images, which are used primarily
// to create instances or configure instance templates. Snapshots are useful
// for periodic backup of the data on your persistent disks. You can create
// snapshots from persistent disks even while they are attached to running
// instances.
//
// Snapshots are incremental, so you can create regular snapshots on a
// persistent disk faster and at a much lower cost than if you regularly
// created a full image of the disk.
//
// To get more information about Snapshot, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
//
// > **Warning:** All arguments including `snapshot_encryption_key.raw_key` and `source_disk_encryption_key.raw_key` will be stored in the raw
// state as plain-text.
//
// ## Example Usage
// ### Snapshot Basic
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
//			debian, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
//				Family:  pulumi.StringRef("debian-11"),
//				Project: pulumi.StringRef("debian-cloud"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			persistent, err := compute.NewDisk(ctx, "persistent", &compute.DiskArgs{
//				Image: *pulumi.String(debian.SelfLink),
//				Size:  pulumi.Int(10),
//				Type:  pulumi.String("pd-ssd"),
//				Zone:  pulumi.String("us-central1-a"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewSnapshot(ctx, "snapshot", &compute.SnapshotArgs{
//				SourceDisk: persistent.ID(),
//				Zone:       pulumi.String("us-central1-a"),
//				Labels: pulumi.StringMap{
//					"my_label": pulumi.String("value"),
//				},
//				StorageLocations: pulumi.StringArray{
//					pulumi.String("us-central1"),
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
// ### Snapshot Chainname
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
//			debian, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
//				Family:  pulumi.StringRef("debian-11"),
//				Project: pulumi.StringRef("debian-cloud"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			persistent, err := compute.NewDisk(ctx, "persistent", &compute.DiskArgs{
//				Image: *pulumi.String(debian.SelfLink),
//				Size:  pulumi.Int(10),
//				Type:  pulumi.String("pd-ssd"),
//				Zone:  pulumi.String("us-central1-a"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewSnapshot(ctx, "snapshot", &compute.SnapshotArgs{
//				SourceDisk: persistent.ID(),
//				Zone:       pulumi.String("us-central1-a"),
//				ChainName:  pulumi.String("snapshot-chain"),
//				Labels: pulumi.StringMap{
//					"my_label": pulumi.String("value"),
//				},
//				StorageLocations: pulumi.StringArray{
//					pulumi.String("us-central1"),
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
// # Snapshot can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/snapshot:Snapshot default projects/{{project}}/global/snapshots/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/snapshot:Snapshot default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/snapshot:Snapshot default {{name}}
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	ChainName pulumi.StringPtrOutput `pulumi:"chainName"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Size of the snapshot, specified in GB.
	DiskSizeGb pulumi.IntOutput `pulumi:"diskSizeGb"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this Snapshot.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A list of public visible licenses that apply to this snapshot. This
	// can be because the original image had licenses attached (such as a
	// Windows image).  snapshotEncryptionKey nested object Encrypts the
	// snapshot using a customer-supplied encryption key.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	// Structure is documented below.
	SnapshotEncryptionKey SnapshotSnapshotEncryptionKeyPtrOutput `pulumi:"snapshotEncryptionKey"`
	// The unique identifier for the resource.
	SnapshotId pulumi.IntOutput `pulumi:"snapshotId"`
	// A reference to the disk used to create this snapshot.
	//
	// ***
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceDiskEncryptionKey SnapshotSourceDiskEncryptionKeyPtrOutput `pulumi:"sourceDiskEncryptionKey"`
	// A size of the storage used by the snapshot. As snapshots share
	// storage, this number is expected to change with snapshot
	// creation/deletion.
	StorageBytes pulumi.IntOutput `pulumi:"storageBytes"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceDisk == nil {
		return nil, errors.New("invalid value for required argument 'SourceDisk'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("gcp:compute/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("gcp:compute/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	ChainName *string `pulumi:"chainName"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Size of the snapshot, specified in GB.
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this Snapshot.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A list of public visible licenses that apply to this snapshot. This
	// can be because the original image had licenses attached (such as a
	// Windows image).  snapshotEncryptionKey nested object Encrypts the
	// snapshot using a customer-supplied encryption key.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	// Structure is documented below.
	SnapshotEncryptionKey *SnapshotSnapshotEncryptionKey `pulumi:"snapshotEncryptionKey"`
	// The unique identifier for the resource.
	SnapshotId *int `pulumi:"snapshotId"`
	// A reference to the disk used to create this snapshot.
	//
	// ***
	SourceDisk *string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceDiskEncryptionKey *SnapshotSourceDiskEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	// A size of the storage used by the snapshot. As snapshots share
	// storage, this number is expected to change with snapshot
	// creation/deletion.
	StorageBytes *int `pulumi:"storageBytes"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations []string `pulumi:"storageLocations"`
	// A reference to the zone where the disk is hosted.
	Zone *string `pulumi:"zone"`
}

type SnapshotState struct {
	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	ChainName pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Size of the snapshot, specified in GB.
	DiskSizeGb pulumi.IntPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this Snapshot.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A list of public visible licenses that apply to this snapshot. This
	// can be because the original image had licenses attached (such as a
	// Windows image).  snapshotEncryptionKey nested object Encrypts the
	// snapshot using a customer-supplied encryption key.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	// Structure is documented below.
	SnapshotEncryptionKey SnapshotSnapshotEncryptionKeyPtrInput
	// The unique identifier for the resource.
	SnapshotId pulumi.IntPtrInput
	// A reference to the disk used to create this snapshot.
	//
	// ***
	SourceDisk pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceDiskEncryptionKey SnapshotSourceDiskEncryptionKeyPtrInput
	// A size of the storage used by the snapshot. As snapshots share
	// storage, this number is expected to change with snapshot
	// creation/deletion.
	StorageBytes pulumi.IntPtrInput
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayInput
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	ChainName *string `pulumi:"chainName"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Labels to apply to this Snapshot.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	// Structure is documented below.
	SnapshotEncryptionKey *SnapshotSnapshotEncryptionKey `pulumi:"snapshotEncryptionKey"`
	// A reference to the disk used to create this snapshot.
	//
	// ***
	SourceDisk string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceDiskEncryptionKey *SnapshotSourceDiskEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations []string `pulumi:"storageLocations"`
	// A reference to the zone where the disk is hosted.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	ChainName pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Labels to apply to this Snapshot.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	// Structure is documented below.
	SnapshotEncryptionKey SnapshotSnapshotEncryptionKeyPtrInput
	// A reference to the disk used to create this snapshot.
	//
	// ***
	SourceDisk pulumi.StringInput
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceDiskEncryptionKey SnapshotSourceDiskEncryptionKeyPtrInput
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayInput
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

func (i *Snapshot) ToOutput(ctx context.Context) pulumix.Output[*Snapshot] {
	return pulumix.Output[*Snapshot]{
		OutputState: i.ToSnapshotOutputWithContext(ctx).OutputState,
	}
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

func (i SnapshotArray) ToOutput(ctx context.Context) pulumix.Output[[]*Snapshot] {
	return pulumix.Output[[]*Snapshot]{
		OutputState: i.ToSnapshotArrayOutputWithContext(ctx).OutputState,
	}
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

func (i SnapshotMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Snapshot] {
	return pulumix.Output[map[string]*Snapshot]{
		OutputState: i.ToSnapshotMapOutputWithContext(ctx).OutputState,
	}
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToOutput(ctx context.Context) pulumix.Output[*Snapshot] {
	return pulumix.Output[*Snapshot]{
		OutputState: o.OutputState,
	}
}

// Creates the new snapshot in the snapshot chain labeled with the
// specified name. The chain name must be 1-63 characters long and
// comply with RFC1035. This is an uncommon option only for advanced
// service owners who needs to create separate snapshot chains, for
// example, for chargeback tracking.  When you describe your snapshot
// resource, this field is visible only if it has a non-empty value.
func (o SnapshotOutput) ChainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.ChainName }).(pulumi.StringPtrOutput)
}

// Creation timestamp in RFC3339 text format.
func (o SnapshotOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o SnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Size of the snapshot, specified in GB.
func (o SnapshotOutput) DiskSizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.DiskSizeGb }).(pulumi.IntOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o SnapshotOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The fingerprint used for optimistic locking of this resource. Used
// internally during updates.
func (o SnapshotOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this Snapshot.
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o SnapshotOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// A list of public visible licenses that apply to this snapshot. This
// can be because the original image had licenses attached (such as a
// Windows image).  snapshotEncryptionKey nested object Encrypts the
// snapshot using a customer-supplied encryption key.
func (o SnapshotOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.Licenses }).(pulumi.StringArrayOutput)
}

// Name of the resource; provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and match
// the regular expression `a-z?` which means the
// first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o SnapshotOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o SnapshotOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The URI of the created resource.
func (o SnapshotOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Encrypts the snapshot using a customer-supplied encryption key.
// After you encrypt a snapshot using a customer-supplied key, you must
// provide the same key if you use the snapshot later. For example, you
// must provide the encryption key when you create a disk from the
// encrypted snapshot in a future request.
// Customer-supplied encryption keys do not protect access to metadata of
// the snapshot.
// If you do not provide an encryption key when creating the snapshot,
// then the snapshot will be encrypted using an automatically generated
// key and you do not need to provide a key to use the snapshot later.
// Structure is documented below.
func (o SnapshotOutput) SnapshotEncryptionKey() SnapshotSnapshotEncryptionKeyPtrOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotSnapshotEncryptionKeyPtrOutput { return v.SnapshotEncryptionKey }).(SnapshotSnapshotEncryptionKeyPtrOutput)
}

// The unique identifier for the resource.
func (o SnapshotOutput) SnapshotId() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.SnapshotId }).(pulumi.IntOutput)
}

// A reference to the disk used to create this snapshot.
//
// ***
func (o SnapshotOutput) SourceDisk() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceDisk }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source snapshot. Required
// if the source snapshot is protected by a customer-supplied encryption
// key.
// Structure is documented below.
func (o SnapshotOutput) SourceDiskEncryptionKey() SnapshotSourceDiskEncryptionKeyPtrOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotSourceDiskEncryptionKeyPtrOutput { return v.SourceDiskEncryptionKey }).(SnapshotSourceDiskEncryptionKeyPtrOutput)
}

// A size of the storage used by the snapshot. As snapshots share
// storage, this number is expected to change with snapshot
// creation/deletion.
func (o SnapshotOutput) StorageBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.StorageBytes }).(pulumi.IntOutput)
}

// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
func (o SnapshotOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

// A reference to the zone where the disk is hosted.
func (o SnapshotOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Snapshot] {
	return pulumix.Output[[]*Snapshot]{
		OutputState: o.OutputState,
	}
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Snapshot] {
	return pulumix.Output[map[string]*Snapshot]{
		OutputState: o.OutputState,
	}
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
