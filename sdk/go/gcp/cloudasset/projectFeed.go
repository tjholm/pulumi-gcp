// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudasset

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Cloud Asset Inventory feed used to to listen to asset updates.
//
// To get more information about ProjectFeed, see:
//
// * [API documentation](https://cloud.google.com/asset-inventory/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/asset-inventory/docs)
//
// ## Example Usage
//
// ## Import
//
// ProjectFeed can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudasset/projectFeed:ProjectFeed default projects/{{project}}/feeds/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudasset/projectFeed:ProjectFeed default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudasset/projectFeed:ProjectFeed default {{name}}
// ```
type ProjectFeed struct {
	pulumi.CustomResourceState

	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames pulumi.StringArrayOutput `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes pulumi.StringArrayOutput `pulumi:"assetTypes"`
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing. If not specified, the resource's
	// project will be used.
	BillingProject pulumi.StringPtrOutput `pulumi:"billingProject"`
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition ProjectFeedConditionPtrOutput `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId pulumi.StringOutput `pulumi:"feedId"`
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig ProjectFeedFeedOutputConfigOutput `pulumi:"feedOutputConfig"`
	// The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectFeed registers a new resource with the given unique name, arguments, and options.
func NewProjectFeed(ctx *pulumi.Context,
	name string, args *ProjectFeedArgs, opts ...pulumi.ResourceOption) (*ProjectFeed, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeedId == nil {
		return nil, errors.New("invalid value for required argument 'FeedId'")
	}
	if args.FeedOutputConfig == nil {
		return nil, errors.New("invalid value for required argument 'FeedOutputConfig'")
	}
	var resource ProjectFeed
	err := ctx.RegisterResource("gcp:cloudasset/projectFeed:ProjectFeed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectFeed gets an existing ProjectFeed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectFeed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectFeedState, opts ...pulumi.ResourceOption) (*ProjectFeed, error) {
	var resource ProjectFeed
	err := ctx.ReadResource("gcp:cloudasset/projectFeed:ProjectFeed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectFeed resources.
type projectFeedState struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames []string `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes []string `pulumi:"assetTypes"`
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing. If not specified, the resource's
	// project will be used.
	BillingProject *string `pulumi:"billingProject"`
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition *ProjectFeedCondition `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType *string `pulumi:"contentType"`
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId *string `pulumi:"feedId"`
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig *ProjectFeedFeedOutputConfig `pulumi:"feedOutputConfig"`
	// The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ProjectFeedState struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames pulumi.StringArrayInput
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes pulumi.StringArrayInput
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing. If not specified, the resource's
	// project will be used.
	BillingProject pulumi.StringPtrInput
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition ProjectFeedConditionPtrInput
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType pulumi.StringPtrInput
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId pulumi.StringPtrInput
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig ProjectFeedFeedOutputConfigPtrInput
	// The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectFeedState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFeedState)(nil)).Elem()
}

type projectFeedArgs struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames []string `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes []string `pulumi:"assetTypes"`
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing. If not specified, the resource's
	// project will be used.
	BillingProject *string `pulumi:"billingProject"`
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition *ProjectFeedCondition `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType *string `pulumi:"contentType"`
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId string `pulumi:"feedId"`
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig ProjectFeedFeedOutputConfig `pulumi:"feedOutputConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectFeed resource.
type ProjectFeedArgs struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames pulumi.StringArrayInput
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes pulumi.StringArrayInput
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing. If not specified, the resource's
	// project will be used.
	BillingProject pulumi.StringPtrInput
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition ProjectFeedConditionPtrInput
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType pulumi.StringPtrInput
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId pulumi.StringInput
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig ProjectFeedFeedOutputConfigInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFeedArgs)(nil)).Elem()
}

type ProjectFeedInput interface {
	pulumi.Input

	ToProjectFeedOutput() ProjectFeedOutput
	ToProjectFeedOutputWithContext(ctx context.Context) ProjectFeedOutput
}

func (*ProjectFeed) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectFeed)(nil))
}

func (i *ProjectFeed) ToProjectFeedOutput() ProjectFeedOutput {
	return i.ToProjectFeedOutputWithContext(context.Background())
}

func (i *ProjectFeed) ToProjectFeedOutputWithContext(ctx context.Context) ProjectFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeedOutput)
}

func (i *ProjectFeed) ToProjectFeedPtrOutput() ProjectFeedPtrOutput {
	return i.ToProjectFeedPtrOutputWithContext(context.Background())
}

func (i *ProjectFeed) ToProjectFeedPtrOutputWithContext(ctx context.Context) ProjectFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeedPtrOutput)
}

type ProjectFeedPtrInput interface {
	pulumi.Input

	ToProjectFeedPtrOutput() ProjectFeedPtrOutput
	ToProjectFeedPtrOutputWithContext(ctx context.Context) ProjectFeedPtrOutput
}

type projectFeedPtrType ProjectFeedArgs

func (*projectFeedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFeed)(nil))
}

func (i *projectFeedPtrType) ToProjectFeedPtrOutput() ProjectFeedPtrOutput {
	return i.ToProjectFeedPtrOutputWithContext(context.Background())
}

func (i *projectFeedPtrType) ToProjectFeedPtrOutputWithContext(ctx context.Context) ProjectFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeedPtrOutput)
}

// ProjectFeedArrayInput is an input type that accepts ProjectFeedArray and ProjectFeedArrayOutput values.
// You can construct a concrete instance of `ProjectFeedArrayInput` via:
//
//          ProjectFeedArray{ ProjectFeedArgs{...} }
type ProjectFeedArrayInput interface {
	pulumi.Input

	ToProjectFeedArrayOutput() ProjectFeedArrayOutput
	ToProjectFeedArrayOutputWithContext(context.Context) ProjectFeedArrayOutput
}

type ProjectFeedArray []ProjectFeedInput

func (ProjectFeedArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ProjectFeed)(nil))
}

func (i ProjectFeedArray) ToProjectFeedArrayOutput() ProjectFeedArrayOutput {
	return i.ToProjectFeedArrayOutputWithContext(context.Background())
}

func (i ProjectFeedArray) ToProjectFeedArrayOutputWithContext(ctx context.Context) ProjectFeedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeedArrayOutput)
}

// ProjectFeedMapInput is an input type that accepts ProjectFeedMap and ProjectFeedMapOutput values.
// You can construct a concrete instance of `ProjectFeedMapInput` via:
//
//          ProjectFeedMap{ "key": ProjectFeedArgs{...} }
type ProjectFeedMapInput interface {
	pulumi.Input

	ToProjectFeedMapOutput() ProjectFeedMapOutput
	ToProjectFeedMapOutputWithContext(context.Context) ProjectFeedMapOutput
}

type ProjectFeedMap map[string]ProjectFeedInput

func (ProjectFeedMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ProjectFeed)(nil))
}

func (i ProjectFeedMap) ToProjectFeedMapOutput() ProjectFeedMapOutput {
	return i.ToProjectFeedMapOutputWithContext(context.Background())
}

func (i ProjectFeedMap) ToProjectFeedMapOutputWithContext(ctx context.Context) ProjectFeedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeedMapOutput)
}

type ProjectFeedOutput struct {
	*pulumi.OutputState
}

func (ProjectFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectFeed)(nil))
}

func (o ProjectFeedOutput) ToProjectFeedOutput() ProjectFeedOutput {
	return o
}

func (o ProjectFeedOutput) ToProjectFeedOutputWithContext(ctx context.Context) ProjectFeedOutput {
	return o
}

func (o ProjectFeedOutput) ToProjectFeedPtrOutput() ProjectFeedPtrOutput {
	return o.ToProjectFeedPtrOutputWithContext(context.Background())
}

func (o ProjectFeedOutput) ToProjectFeedPtrOutputWithContext(ctx context.Context) ProjectFeedPtrOutput {
	return o.ApplyT(func(v ProjectFeed) *ProjectFeed {
		return &v
	}).(ProjectFeedPtrOutput)
}

type ProjectFeedPtrOutput struct {
	*pulumi.OutputState
}

func (ProjectFeedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFeed)(nil))
}

func (o ProjectFeedPtrOutput) ToProjectFeedPtrOutput() ProjectFeedPtrOutput {
	return o
}

func (o ProjectFeedPtrOutput) ToProjectFeedPtrOutputWithContext(ctx context.Context) ProjectFeedPtrOutput {
	return o
}

type ProjectFeedArrayOutput struct{ *pulumi.OutputState }

func (ProjectFeedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectFeed)(nil))
}

func (o ProjectFeedArrayOutput) ToProjectFeedArrayOutput() ProjectFeedArrayOutput {
	return o
}

func (o ProjectFeedArrayOutput) ToProjectFeedArrayOutputWithContext(ctx context.Context) ProjectFeedArrayOutput {
	return o
}

func (o ProjectFeedArrayOutput) Index(i pulumi.IntInput) ProjectFeedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectFeed {
		return vs[0].([]ProjectFeed)[vs[1].(int)]
	}).(ProjectFeedOutput)
}

type ProjectFeedMapOutput struct{ *pulumi.OutputState }

func (ProjectFeedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProjectFeed)(nil))
}

func (o ProjectFeedMapOutput) ToProjectFeedMapOutput() ProjectFeedMapOutput {
	return o
}

func (o ProjectFeedMapOutput) ToProjectFeedMapOutputWithContext(ctx context.Context) ProjectFeedMapOutput {
	return o
}

func (o ProjectFeedMapOutput) MapIndex(k pulumi.StringInput) ProjectFeedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProjectFeed {
		return vs[0].(map[string]ProjectFeed)[vs[1].(string)]
	}).(ProjectFeedOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectFeedOutput{})
	pulumi.RegisterOutputType(ProjectFeedPtrOutput{})
	pulumi.RegisterOutputType(ProjectFeedArrayOutput{})
	pulumi.RegisterOutputType(ProjectFeedMapOutput{})
}
