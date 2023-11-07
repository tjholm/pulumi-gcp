// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package filestore

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Google Cloud Filestore snapshot.
//
// To get more information about Snapshot, see:
//
// * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1/projects.locations.instances.snapshots)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/filestore/docs/snapshots)
//   - [Creating Snapshots](https://cloud.google.com/filestore/docs/create-snapshots)
//
// ## Example Usage
// ### Filestore Snapshot Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/filestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instance, err := filestore.NewInstance(ctx, "instance", &filestore.InstanceArgs{
//				Location: pulumi.String("us-east1"),
//				Tier:     pulumi.String("ENTERPRISE"),
//				FileShares: &filestore.InstanceFileSharesArgs{
//					CapacityGb: pulumi.Int(1024),
//					Name:       pulumi.String("share1"),
//				},
//				Networks: filestore.InstanceNetworkArray{
//					&filestore.InstanceNetworkArgs{
//						Network: pulumi.String("default"),
//						Modes: pulumi.StringArray{
//							pulumi.String("MODE_IPV4"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = filestore.NewSnapshot(ctx, "snapshot", &filestore.SnapshotArgs{
//				Instance: instance.Name,
//				Location: pulumi.String("us-east1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Filestore Snapshot Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/filestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instance, err := filestore.NewInstance(ctx, "instance", &filestore.InstanceArgs{
//				Location: pulumi.String("us-west1"),
//				Tier:     pulumi.String("ENTERPRISE"),
//				FileShares: &filestore.InstanceFileSharesArgs{
//					CapacityGb: pulumi.Int(1024),
//					Name:       pulumi.String("share1"),
//				},
//				Networks: filestore.InstanceNetworkArray{
//					&filestore.InstanceNetworkArgs{
//						Network: pulumi.String("default"),
//						Modes: pulumi.StringArray{
//							pulumi.String("MODE_IPV4"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = filestore.NewSnapshot(ctx, "snapshot", &filestore.SnapshotArgs{
//				Instance:    instance.Name,
//				Location:    pulumi.String("us-west1"),
//				Description: pulumi.String("Snapshot of test-instance-for-snapshot"),
//				Labels: pulumi.StringMap{
//					"my_label": pulumi.String("value"),
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
//	$ pulumi import gcp:filestore/snapshot:Snapshot default projects/{{project}}/locations/{{location}}/instances/{{instance}}/snapshots/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:filestore/snapshot:Snapshot default {{project}}/{{location}}/{{instance}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:filestore/snapshot:Snapshot default {{location}}/{{instance}}/{{name}}
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// The time when the snapshot was created in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The amount of bytes needed to allocate a full copy of the snapshot content.
	FilesystemUsedBytes pulumi.StringOutput `pulumi:"filesystemUsedBytes"`
	// The resource name of the filestore instance.
	//
	// ***
	Instance pulumi.StringOutput `pulumi:"instance"`
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the snapshot. The name must be unique within the specified instance.
	// The name must be 1-63 characters long, and comply with
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
	// The snapshot state.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("gcp:filestore/snapshot:Snapshot", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:filestore/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// The time when the snapshot was created in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The amount of bytes needed to allocate a full copy of the snapshot content.
	FilesystemUsedBytes *string `pulumi:"filesystemUsedBytes"`
	// The resource name of the filestore instance.
	//
	// ***
	Instance *string `pulumi:"instance"`
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location *string `pulumi:"location"`
	// The resource name of the snapshot. The name must be unique within the specified instance.
	// The name must be 1-63 characters long, and comply with
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
	// The snapshot state.
	State *string `pulumi:"state"`
}

type SnapshotState struct {
	// The time when the snapshot was created in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapInput
	// The amount of bytes needed to allocate a full copy of the snapshot content.
	FilesystemUsedBytes pulumi.StringPtrInput
	// The resource name of the filestore instance.
	//
	// ***
	Instance pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location pulumi.StringPtrInput
	// The resource name of the snapshot. The name must be unique within the specified instance.
	// The name must be 1-63 characters long, and comply with
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
	// The snapshot state.
	State pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description *string `pulumi:"description"`
	// The resource name of the filestore instance.
	//
	// ***
	Instance string `pulumi:"instance"`
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location string `pulumi:"location"`
	// The resource name of the snapshot. The name must be unique within the specified instance.
	// The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description pulumi.StringPtrInput
	// The resource name of the filestore instance.
	//
	// ***
	Instance pulumi.StringInput
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location pulumi.StringInput
	// The resource name of the snapshot. The name must be unique within the specified instance.
	// The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
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

// The time when the snapshot was created in RFC3339 text format.
func (o SnapshotOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
func (o SnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o SnapshotOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The amount of bytes needed to allocate a full copy of the snapshot content.
func (o SnapshotOutput) FilesystemUsedBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.FilesystemUsedBytes }).(pulumi.StringOutput)
}

// The resource name of the filestore instance.
//
// ***
func (o SnapshotOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// Resource labels to represent user-provided metadata.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o SnapshotOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
func (o SnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the snapshot. The name must be unique within the specified instance.
// The name must be 1-63 characters long, and comply with
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

// The snapshot state.
func (o SnapshotOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
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
