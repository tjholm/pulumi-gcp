// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataplex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Dataplex Asset resource
//
// ## Example Usage
// ### Basic_asset
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			basicBucket, err := storage.NewBucket(ctx, "basicBucket", &storage.BucketArgs{
//				Location:                 pulumi.String("us-west1"),
//				UniformBucketLevelAccess: pulumi.Bool(true),
//				Project:                  pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			basicLake, err := dataplex.NewLake(ctx, "basicLake", &dataplex.LakeArgs{
//				Location: pulumi.String("us-west1"),
//				Project:  pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			basicZone, err := dataplex.NewZone(ctx, "basicZone", &dataplex.ZoneArgs{
//				Location: pulumi.String("us-west1"),
//				Lake:     basicLake.Name,
//				Type:     pulumi.String("RAW"),
//				DiscoverySpec: &dataplex.ZoneDiscoverySpecArgs{
//					Enabled: pulumi.Bool(false),
//				},
//				ResourceSpec: &dataplex.ZoneResourceSpecArgs{
//					LocationType: pulumi.String("SINGLE_REGION"),
//				},
//				Project: pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dataplex.NewAsset(ctx, "primary", &dataplex.AssetArgs{
//				Location:     pulumi.String("us-west1"),
//				Lake:         basicLake.Name,
//				DataplexZone: basicZone.Name,
//				DiscoverySpec: &dataplex.AssetDiscoverySpecArgs{
//					Enabled: pulumi.Bool(false),
//				},
//				ResourceSpec: &dataplex.AssetResourceSpecArgs{
//					Name: pulumi.String("projects/my-project-name/buckets/bucket"),
//					Type: pulumi.String("STORAGE_BUCKET"),
//				},
//				Labels: pulumi.StringMap{
//					"env":      pulumi.String("foo"),
//					"my-asset": pulumi.String("exists"),
//				},
//				Project: pulumi.String("my-project-name"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				basicBucket,
//			}))
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
// # Asset can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:dataplex/asset:Asset default projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataplex/asset:Asset default {{project}}/{{location}}/{{lake}}/{{dataplex_zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataplex/asset:Asset default {{location}}/{{lake}}/{{dataplex_zone}}/{{name}}
//
// ```
type Asset struct {
	pulumi.CustomResourceState

	// Output only. The time when the asset was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The zone for the resource
	DataplexZone pulumi.StringOutput `pulumi:"dataplexZone"`
	// Optional. Description of the asset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec AssetDiscoverySpecOutput `pulumi:"discoverySpec"`
	// Output only. Status of the discovery feature applied to data referenced by this asset.
	DiscoveryStatuses AssetDiscoveryStatusArrayOutput `pulumi:"discoveryStatuses"`
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.MapOutput `pulumi:"effectiveLabels"`
	// Optional. User defined labels for the asset.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The lake for the resource
	Lake pulumi.StringOutput `pulumi:"lake"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the asset.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapOutput `pulumi:"pulumiLabels"`
	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec AssetResourceSpecOutput `pulumi:"resourceSpec"`
	// Output only. Status of the resource referenced by this asset.
	ResourceStatuses AssetResourceStatusArrayOutput `pulumi:"resourceStatuses"`
	// Output only. Status of the security policy applied to resource referenced by this asset.
	SecurityStatuses AssetSecurityStatusArrayOutput `pulumi:"securityStatuses"`
	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State pulumi.StringOutput `pulumi:"state"`
	// Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time when the asset was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAsset registers a new resource with the given unique name, arguments, and options.
func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataplexZone == nil {
		return nil, errors.New("invalid value for required argument 'DataplexZone'")
	}
	if args.DiscoverySpec == nil {
		return nil, errors.New("invalid value for required argument 'DiscoverySpec'")
	}
	if args.Lake == nil {
		return nil, errors.New("invalid value for required argument 'Lake'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceSpec == nil {
		return nil, errors.New("invalid value for required argument 'ResourceSpec'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Asset
	err := ctx.RegisterResource("gcp:dataplex/asset:Asset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAsset gets an existing Asset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetState, opts ...pulumi.ResourceOption) (*Asset, error) {
	var resource Asset
	err := ctx.ReadResource("gcp:dataplex/asset:Asset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Asset resources.
type assetState struct {
	// Output only. The time when the asset was created.
	CreateTime *string `pulumi:"createTime"`
	// The zone for the resource
	DataplexZone *string `pulumi:"dataplexZone"`
	// Optional. Description of the asset.
	Description *string `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec *AssetDiscoverySpec `pulumi:"discoverySpec"`
	// Output only. Status of the discovery feature applied to data referenced by this asset.
	DiscoveryStatuses []AssetDiscoveryStatus `pulumi:"discoveryStatuses"`
	// Optional. User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]interface{} `pulumi:"effectiveLabels"`
	// Optional. User defined labels for the asset.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake for the resource
	Lake *string `pulumi:"lake"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// The name of the asset.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels map[string]interface{} `pulumi:"pulumiLabels"`
	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec *AssetResourceSpec `pulumi:"resourceSpec"`
	// Output only. Status of the resource referenced by this asset.
	ResourceStatuses []AssetResourceStatus `pulumi:"resourceStatuses"`
	// Output only. Status of the security policy applied to resource referenced by this asset.
	SecurityStatuses []AssetSecurityStatus `pulumi:"securityStatuses"`
	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `pulumi:"state"`
	// Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
	Uid *string `pulumi:"uid"`
	// Output only. The time when the asset was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type AssetState struct {
	// Output only. The time when the asset was created.
	CreateTime pulumi.StringPtrInput
	// The zone for the resource
	DataplexZone pulumi.StringPtrInput
	// Optional. Description of the asset.
	Description pulumi.StringPtrInput
	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec AssetDiscoverySpecPtrInput
	// Output only. Status of the discovery feature applied to data referenced by this asset.
	DiscoveryStatuses AssetDiscoveryStatusArrayInput
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.MapInput
	// Optional. User defined labels for the asset.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The lake for the resource
	Lake pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// The name of the asset.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapInput
	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec AssetResourceSpecPtrInput
	// Output only. Status of the resource referenced by this asset.
	ResourceStatuses AssetResourceStatusArrayInput
	// Output only. Status of the security policy applied to resource referenced by this asset.
	SecurityStatuses AssetSecurityStatusArrayInput
	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State pulumi.StringPtrInput
	// Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
	Uid pulumi.StringPtrInput
	// Output only. The time when the asset was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (AssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetState)(nil)).Elem()
}

type assetArgs struct {
	// The zone for the resource
	DataplexZone string `pulumi:"dataplexZone"`
	// Optional. Description of the asset.
	Description *string `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec AssetDiscoverySpec `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// Optional. User defined labels for the asset.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake for the resource
	Lake string `pulumi:"lake"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The name of the asset.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec AssetResourceSpec `pulumi:"resourceSpec"`
}

// The set of arguments for constructing a Asset resource.
type AssetArgs struct {
	// The zone for the resource
	DataplexZone pulumi.StringInput
	// Optional. Description of the asset.
	Description pulumi.StringPtrInput
	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec AssetDiscoverySpecInput
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrInput
	// Optional. User defined labels for the asset.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The lake for the resource
	Lake pulumi.StringInput
	// The location for the resource
	Location pulumi.StringInput
	// The name of the asset.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec AssetResourceSpecInput
}

func (AssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetArgs)(nil)).Elem()
}

type AssetInput interface {
	pulumi.Input

	ToAssetOutput() AssetOutput
	ToAssetOutputWithContext(ctx context.Context) AssetOutput
}

func (*Asset) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (i *Asset) ToAssetOutput() AssetOutput {
	return i.ToAssetOutputWithContext(context.Background())
}

func (i *Asset) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetOutput)
}

func (i *Asset) ToOutput(ctx context.Context) pulumix.Output[*Asset] {
	return pulumix.Output[*Asset]{
		OutputState: i.ToAssetOutputWithContext(ctx).OutputState,
	}
}

// AssetArrayInput is an input type that accepts AssetArray and AssetArrayOutput values.
// You can construct a concrete instance of `AssetArrayInput` via:
//
//	AssetArray{ AssetArgs{...} }
type AssetArrayInput interface {
	pulumi.Input

	ToAssetArrayOutput() AssetArrayOutput
	ToAssetArrayOutputWithContext(context.Context) AssetArrayOutput
}

type AssetArray []AssetInput

func (AssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Asset)(nil)).Elem()
}

func (i AssetArray) ToAssetArrayOutput() AssetArrayOutput {
	return i.ToAssetArrayOutputWithContext(context.Background())
}

func (i AssetArray) ToAssetArrayOutputWithContext(ctx context.Context) AssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetArrayOutput)
}

func (i AssetArray) ToOutput(ctx context.Context) pulumix.Output[[]*Asset] {
	return pulumix.Output[[]*Asset]{
		OutputState: i.ToAssetArrayOutputWithContext(ctx).OutputState,
	}
}

// AssetMapInput is an input type that accepts AssetMap and AssetMapOutput values.
// You can construct a concrete instance of `AssetMapInput` via:
//
//	AssetMap{ "key": AssetArgs{...} }
type AssetMapInput interface {
	pulumi.Input

	ToAssetMapOutput() AssetMapOutput
	ToAssetMapOutputWithContext(context.Context) AssetMapOutput
}

type AssetMap map[string]AssetInput

func (AssetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Asset)(nil)).Elem()
}

func (i AssetMap) ToAssetMapOutput() AssetMapOutput {
	return i.ToAssetMapOutputWithContext(context.Background())
}

func (i AssetMap) ToAssetMapOutputWithContext(ctx context.Context) AssetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetMapOutput)
}

func (i AssetMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Asset] {
	return pulumix.Output[map[string]*Asset]{
		OutputState: i.ToAssetMapOutputWithContext(ctx).OutputState,
	}
}

type AssetOutput struct{ *pulumi.OutputState }

func (AssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (o AssetOutput) ToAssetOutput() AssetOutput {
	return o
}

func (o AssetOutput) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return o
}

func (o AssetOutput) ToOutput(ctx context.Context) pulumix.Output[*Asset] {
	return pulumix.Output[*Asset]{
		OutputState: o.OutputState,
	}
}

// Output only. The time when the asset was created.
func (o AssetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The zone for the resource
func (o AssetOutput) DataplexZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.DataplexZone }).(pulumi.StringOutput)
}

// Optional. Description of the asset.
func (o AssetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
func (o AssetOutput) DiscoverySpec() AssetDiscoverySpecOutput {
	return o.ApplyT(func(v *Asset) AssetDiscoverySpecOutput { return v.DiscoverySpec }).(AssetDiscoverySpecOutput)
}

// Output only. Status of the discovery feature applied to data referenced by this asset.
func (o AssetOutput) DiscoveryStatuses() AssetDiscoveryStatusArrayOutput {
	return o.ApplyT(func(v *Asset) AssetDiscoveryStatusArrayOutput { return v.DiscoveryStatuses }).(AssetDiscoveryStatusArrayOutput)
}

// Optional. User friendly display name.
func (o AssetOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o AssetOutput) EffectiveLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Asset) pulumi.MapOutput { return v.EffectiveLabels }).(pulumi.MapOutput)
}

// Optional. User defined labels for the asset.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o AssetOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The lake for the resource
func (o AssetOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

// The location for the resource
func (o AssetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the asset.
func (o AssetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o AssetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o AssetOutput) PulumiLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Asset) pulumi.MapOutput { return v.PulumiLabels }).(pulumi.MapOutput)
}

// Required. Immutable. Specification of the resource that is referenced by this asset.
func (o AssetOutput) ResourceSpec() AssetResourceSpecOutput {
	return o.ApplyT(func(v *Asset) AssetResourceSpecOutput { return v.ResourceSpec }).(AssetResourceSpecOutput)
}

// Output only. Status of the resource referenced by this asset.
func (o AssetOutput) ResourceStatuses() AssetResourceStatusArrayOutput {
	return o.ApplyT(func(v *Asset) AssetResourceStatusArrayOutput { return v.ResourceStatuses }).(AssetResourceStatusArrayOutput)
}

// Output only. Status of the security policy applied to resource referenced by this asset.
func (o AssetOutput) SecurityStatuses() AssetSecurityStatusArrayOutput {
	return o.ApplyT(func(v *Asset) AssetSecurityStatusArrayOutput { return v.SecurityStatuses }).(AssetSecurityStatusArrayOutput)
}

// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
func (o AssetOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
func (o AssetOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The time when the asset was last updated.
func (o AssetOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type AssetArrayOutput struct{ *pulumi.OutputState }

func (AssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Asset)(nil)).Elem()
}

func (o AssetArrayOutput) ToAssetArrayOutput() AssetArrayOutput {
	return o
}

func (o AssetArrayOutput) ToAssetArrayOutputWithContext(ctx context.Context) AssetArrayOutput {
	return o
}

func (o AssetArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Asset] {
	return pulumix.Output[[]*Asset]{
		OutputState: o.OutputState,
	}
}

func (o AssetArrayOutput) Index(i pulumi.IntInput) AssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Asset {
		return vs[0].([]*Asset)[vs[1].(int)]
	}).(AssetOutput)
}

type AssetMapOutput struct{ *pulumi.OutputState }

func (AssetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Asset)(nil)).Elem()
}

func (o AssetMapOutput) ToAssetMapOutput() AssetMapOutput {
	return o
}

func (o AssetMapOutput) ToAssetMapOutputWithContext(ctx context.Context) AssetMapOutput {
	return o
}

func (o AssetMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Asset] {
	return pulumix.Output[map[string]*Asset]{
		OutputState: o.OutputState,
	}
}

func (o AssetMapOutput) MapIndex(k pulumi.StringInput) AssetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Asset {
		return vs[0].(map[string]*Asset)[vs[1].(string)]
	}).(AssetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssetInput)(nil)).Elem(), &Asset{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetArrayInput)(nil)).Elem(), AssetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetMapInput)(nil)).Elem(), AssetMap{})
	pulumi.RegisterOutputType(AssetOutput{})
	pulumi.RegisterOutputType(AssetArrayOutput{})
	pulumi.RegisterOutputType(AssetMapOutput{})
}
