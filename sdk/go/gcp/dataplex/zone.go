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

// The Dataplex Zone resource
//
// ## Example Usage
// ### Basic_zone
// A basic example of a dataplex zone
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			basic, err := dataplex.NewLake(ctx, "basic", &dataplex.LakeArgs{
//				Location:    pulumi.String("us-west1"),
//				Description: pulumi.String("Lake for DCL"),
//				DisplayName: pulumi.String("Lake for DCL"),
//				Project:     pulumi.String("my-project-name"),
//				Labels: pulumi.StringMap{
//					"my-lake": pulumi.String("exists"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dataplex.NewZone(ctx, "primary", &dataplex.ZoneArgs{
//				DiscoverySpec: &dataplex.ZoneDiscoverySpecArgs{
//					Enabled: pulumi.Bool(false),
//				},
//				Lake:     basic.Name,
//				Location: pulumi.String("us-west1"),
//				ResourceSpec: &dataplex.ZoneResourceSpecArgs{
//					LocationType: pulumi.String("MULTI_REGION"),
//				},
//				Type:        pulumi.String("RAW"),
//				Description: pulumi.String("Zone for DCL"),
//				DisplayName: pulumi.String("Zone for DCL"),
//				Project:     pulumi.String("my-project-name"),
//				Labels:      nil,
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
// # Zone can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:dataplex/zone:Zone default projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataplex/zone:Zone default {{project}}/{{location}}/{{lake}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataplex/zone:Zone default {{location}}/{{lake}}/{{name}}
//
// ```
type Zone struct {
	pulumi.CustomResourceState

	// Output only. Aggregated status of the underlying assets of the zone.
	AssetStatuses ZoneAssetStatusArrayOutput `pulumi:"assetStatuses"`
	// Output only. The time when the zone was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Description of the zone.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec ZoneDiscoverySpecOutput `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.MapOutput `pulumi:"effectiveLabels"`
	// Optional. User defined labels for the zone.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The lake for the resource
	Lake pulumi.StringOutput `pulumi:"lake"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapOutput `pulumi:"pulumiLabels"`
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec ZoneResourceSpecOutput `pulumi:"resourceSpec"`
	// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State pulumi.StringOutput `pulumi:"state"`
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type pulumi.StringOutput `pulumi:"type"`
	// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time when the zone was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Zone
	err := ctx.RegisterResource("gcp:dataplex/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("gcp:dataplex/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// Output only. Aggregated status of the underlying assets of the zone.
	AssetStatuses []ZoneAssetStatus `pulumi:"assetStatuses"`
	// Output only. The time when the zone was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. Description of the zone.
	Description *string `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec *ZoneDiscoverySpec `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]interface{} `pulumi:"effectiveLabels"`
	// Optional. User defined labels for the zone.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake for the resource
	Lake *string `pulumi:"lake"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// The name of the zone.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels map[string]interface{} `pulumi:"pulumiLabels"`
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec *ZoneResourceSpec `pulumi:"resourceSpec"`
	// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `pulumi:"state"`
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type *string `pulumi:"type"`
	// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
	Uid *string `pulumi:"uid"`
	// Output only. The time when the zone was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type ZoneState struct {
	// Output only. Aggregated status of the underlying assets of the zone.
	AssetStatuses ZoneAssetStatusArrayInput
	// Output only. The time when the zone was created.
	CreateTime pulumi.StringPtrInput
	// Optional. Description of the zone.
	Description pulumi.StringPtrInput
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec ZoneDiscoverySpecPtrInput
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.MapInput
	// Optional. User defined labels for the zone.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The lake for the resource
	Lake pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// The name of the zone.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapInput
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec ZoneResourceSpecPtrInput
	// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State pulumi.StringPtrInput
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type pulumi.StringPtrInput
	// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
	Uid pulumi.StringPtrInput
	// Output only. The time when the zone was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// Optional. Description of the zone.
	Description *string `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec ZoneDiscoverySpec `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// Optional. User defined labels for the zone.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake for the resource
	Lake string `pulumi:"lake"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The name of the zone.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec ZoneResourceSpec `pulumi:"resourceSpec"`
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// Optional. Description of the zone.
	Description pulumi.StringPtrInput
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec ZoneDiscoverySpecInput
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrInput
	// Optional. User defined labels for the zone.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The lake for the resource
	Lake pulumi.StringInput
	// The location for the resource
	Location pulumi.StringInput
	// The name of the zone.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec ZoneResourceSpecInput
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type pulumi.StringInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

func (i *Zone) ToOutput(ctx context.Context) pulumix.Output[*Zone] {
	return pulumix.Output[*Zone]{
		OutputState: i.ToZoneOutputWithContext(ctx).OutputState,
	}
}

// ZoneArrayInput is an input type that accepts ZoneArray and ZoneArrayOutput values.
// You can construct a concrete instance of `ZoneArrayInput` via:
//
//	ZoneArray{ ZoneArgs{...} }
type ZoneArrayInput interface {
	pulumi.Input

	ToZoneArrayOutput() ZoneArrayOutput
	ToZoneArrayOutputWithContext(context.Context) ZoneArrayOutput
}

type ZoneArray []ZoneInput

func (ZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (i ZoneArray) ToZoneArrayOutput() ZoneArrayOutput {
	return i.ToZoneArrayOutputWithContext(context.Background())
}

func (i ZoneArray) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneArrayOutput)
}

func (i ZoneArray) ToOutput(ctx context.Context) pulumix.Output[[]*Zone] {
	return pulumix.Output[[]*Zone]{
		OutputState: i.ToZoneArrayOutputWithContext(ctx).OutputState,
	}
}

// ZoneMapInput is an input type that accepts ZoneMap and ZoneMapOutput values.
// You can construct a concrete instance of `ZoneMapInput` via:
//
//	ZoneMap{ "key": ZoneArgs{...} }
type ZoneMapInput interface {
	pulumi.Input

	ToZoneMapOutput() ZoneMapOutput
	ToZoneMapOutputWithContext(context.Context) ZoneMapOutput
}

type ZoneMap map[string]ZoneInput

func (ZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (i ZoneMap) ToZoneMapOutput() ZoneMapOutput {
	return i.ToZoneMapOutputWithContext(context.Background())
}

func (i ZoneMap) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMapOutput)
}

func (i ZoneMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Zone] {
	return pulumix.Output[map[string]*Zone]{
		OutputState: i.ToZoneMapOutputWithContext(ctx).OutputState,
	}
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

func (o ZoneOutput) ToOutput(ctx context.Context) pulumix.Output[*Zone] {
	return pulumix.Output[*Zone]{
		OutputState: o.OutputState,
	}
}

// Output only. Aggregated status of the underlying assets of the zone.
func (o ZoneOutput) AssetStatuses() ZoneAssetStatusArrayOutput {
	return o.ApplyT(func(v *Zone) ZoneAssetStatusArrayOutput { return v.AssetStatuses }).(ZoneAssetStatusArrayOutput)
}

// Output only. The time when the zone was created.
func (o ZoneOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the zone.
func (o ZoneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Required. Specification of the discovery feature applied to data in this zone.
func (o ZoneOutput) DiscoverySpec() ZoneDiscoverySpecOutput {
	return o.ApplyT(func(v *Zone) ZoneDiscoverySpecOutput { return v.DiscoverySpec }).(ZoneDiscoverySpecOutput)
}

// Optional. User friendly display name.
func (o ZoneOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o ZoneOutput) EffectiveLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Zone) pulumi.MapOutput { return v.EffectiveLabels }).(pulumi.MapOutput)
}

// Optional. User defined labels for the zone.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o ZoneOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The lake for the resource
func (o ZoneOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

// The location for the resource
func (o ZoneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the zone.
func (o ZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o ZoneOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ZoneOutput) PulumiLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Zone) pulumi.MapOutput { return v.PulumiLabels }).(pulumi.MapOutput)
}

// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
func (o ZoneOutput) ResourceSpec() ZoneResourceSpecOutput {
	return o.ApplyT(func(v *Zone) ZoneResourceSpecOutput { return v.ResourceSpec }).(ZoneResourceSpecOutput)
}

// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
func (o ZoneOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
func (o ZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
func (o ZoneOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The time when the zone was last updated.
func (o ZoneOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type ZoneArrayOutput struct{ *pulumi.OutputState }

func (ZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (o ZoneArrayOutput) ToZoneArrayOutput() ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Zone] {
	return pulumix.Output[[]*Zone]{
		OutputState: o.OutputState,
	}
}

func (o ZoneArrayOutput) Index(i pulumi.IntInput) ZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].([]*Zone)[vs[1].(int)]
	}).(ZoneOutput)
}

type ZoneMapOutput struct{ *pulumi.OutputState }

func (ZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (o ZoneMapOutput) ToZoneMapOutput() ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Zone] {
	return pulumix.Output[map[string]*Zone]{
		OutputState: o.OutputState,
	}
}

func (o ZoneMapOutput) MapIndex(k pulumi.StringInput) ZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].(map[string]*Zone)[vs[1].(string)]
	}).(ZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneArrayInput)(nil)).Elem(), ZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMapInput)(nil)).Elem(), ZoneMap{})
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZoneArrayOutput{})
	pulumi.RegisterOutputType(ZoneMapOutput{})
}
