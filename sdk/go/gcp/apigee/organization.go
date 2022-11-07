// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An `Organization` is the top-level container in Apigee.
//
// To get more information about Organization, see:
//
// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations)
// * How-to Guides
//   - [Creating an API organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org)
//
// ## Example Usage
//
// ## Import
//
// # Organization can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:apigee/organization:Organization default organizations/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigee/organization:Organization default {{name}}
//
// ```
type Organization struct {
	pulumi.CustomResourceState

	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion pulumi.StringPtrOutput `pulumi:"analyticsRegion"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork pulumi.StringPtrOutput `pulumi:"authorizedNetwork"`
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType pulumi.StringOutput `pulumi:"billingType"`
	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
	// is CLOUD. A base64-encoded string.
	CaCertificate pulumi.StringOutput `pulumi:"caCertificate"`
	// Description of the Apigee organization.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the Apigee organization.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Name of the property.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project ID associated with the Apigee organization.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	Properties OrganizationPropertiesOutput `pulumi:"properties"`
	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is `DELETION_RETENTION_UNSPECIFIED`.
	// Possible values are `DELETION_RETENTION_UNSPECIFIED` and `MINIMUM`.
	Retention pulumi.StringPtrOutput `pulumi:"retention"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName pulumi.StringPtrOutput `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType pulumi.StringPtrOutput `pulumi:"runtimeType"`
	// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
	// purposes only) or paid (full subscription has been purchased).
	SubscriptionType pulumi.StringOutput `pulumi:"subscriptionType"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Organization
	err := ctx.RegisterResource("gcp:apigee/organization:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("gcp:apigee/organization:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion *string `pulumi:"analyticsRegion"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType *string `pulumi:"billingType"`
	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
	// is CLOUD. A base64-encoded string.
	CaCertificate *string `pulumi:"caCertificate"`
	// Description of the Apigee organization.
	Description *string `pulumi:"description"`
	// The display name of the Apigee organization.
	DisplayName *string `pulumi:"displayName"`
	// Name of the property.
	Name *string `pulumi:"name"`
	// The project ID associated with the Apigee organization.
	ProjectId *string `pulumi:"projectId"`
	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	Properties *OrganizationProperties `pulumi:"properties"`
	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is `DELETION_RETENTION_UNSPECIFIED`.
	// Possible values are `DELETION_RETENTION_UNSPECIFIED` and `MINIMUM`.
	Retention *string `pulumi:"retention"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName *string `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType *string `pulumi:"runtimeType"`
	// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
	// purposes only) or paid (full subscription has been purchased).
	SubscriptionType *string `pulumi:"subscriptionType"`
}

type OrganizationState struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion pulumi.StringPtrInput
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork pulumi.StringPtrInput
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType pulumi.StringPtrInput
	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
	// is CLOUD. A base64-encoded string.
	CaCertificate pulumi.StringPtrInput
	// Description of the Apigee organization.
	Description pulumi.StringPtrInput
	// The display name of the Apigee organization.
	DisplayName pulumi.StringPtrInput
	// Name of the property.
	Name pulumi.StringPtrInput
	// The project ID associated with the Apigee organization.
	ProjectId pulumi.StringPtrInput
	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	Properties OrganizationPropertiesPtrInput
	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is `DELETION_RETENTION_UNSPECIFIED`.
	// Possible values are `DELETION_RETENTION_UNSPECIFIED` and `MINIMUM`.
	Retention pulumi.StringPtrInput
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName pulumi.StringPtrInput
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType pulumi.StringPtrInput
	// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
	// purposes only) or paid (full subscription has been purchased).
	SubscriptionType pulumi.StringPtrInput
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion *string `pulumi:"analyticsRegion"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType *string `pulumi:"billingType"`
	// Description of the Apigee organization.
	Description *string `pulumi:"description"`
	// The display name of the Apigee organization.
	DisplayName *string `pulumi:"displayName"`
	// The project ID associated with the Apigee organization.
	ProjectId string `pulumi:"projectId"`
	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	Properties *OrganizationProperties `pulumi:"properties"`
	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is `DELETION_RETENTION_UNSPECIFIED`.
	// Possible values are `DELETION_RETENTION_UNSPECIFIED` and `MINIMUM`.
	Retention *string `pulumi:"retention"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName *string `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType *string `pulumi:"runtimeType"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion pulumi.StringPtrInput
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork pulumi.StringPtrInput
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType pulumi.StringPtrInput
	// Description of the Apigee organization.
	Description pulumi.StringPtrInput
	// The display name of the Apigee organization.
	DisplayName pulumi.StringPtrInput
	// The project ID associated with the Apigee organization.
	ProjectId pulumi.StringInput
	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	Properties OrganizationPropertiesPtrInput
	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is `DELETION_RETENTION_UNSPECIFIED`.
	// Possible values are `DELETION_RETENTION_UNSPECIFIED` and `MINIMUM`.
	Retention pulumi.StringPtrInput
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName pulumi.StringPtrInput
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType pulumi.StringPtrInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

func (*Organization) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (i *Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i *Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

// OrganizationArrayInput is an input type that accepts OrganizationArray and OrganizationArrayOutput values.
// You can construct a concrete instance of `OrganizationArrayInput` via:
//
//	OrganizationArray{ OrganizationArgs{...} }
type OrganizationArrayInput interface {
	pulumi.Input

	ToOrganizationArrayOutput() OrganizationArrayOutput
	ToOrganizationArrayOutputWithContext(context.Context) OrganizationArrayOutput
}

type OrganizationArray []OrganizationInput

func (OrganizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Organization)(nil)).Elem()
}

func (i OrganizationArray) ToOrganizationArrayOutput() OrganizationArrayOutput {
	return i.ToOrganizationArrayOutputWithContext(context.Background())
}

func (i OrganizationArray) ToOrganizationArrayOutputWithContext(ctx context.Context) OrganizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationArrayOutput)
}

// OrganizationMapInput is an input type that accepts OrganizationMap and OrganizationMapOutput values.
// You can construct a concrete instance of `OrganizationMapInput` via:
//
//	OrganizationMap{ "key": OrganizationArgs{...} }
type OrganizationMapInput interface {
	pulumi.Input

	ToOrganizationMapOutput() OrganizationMapOutput
	ToOrganizationMapOutputWithContext(context.Context) OrganizationMapOutput
}

type OrganizationMap map[string]OrganizationInput

func (OrganizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Organization)(nil)).Elem()
}

func (i OrganizationMap) ToOrganizationMapOutput() OrganizationMapOutput {
	return i.ToOrganizationMapOutputWithContext(context.Background())
}

func (i OrganizationMap) ToOrganizationMapOutputWithContext(ctx context.Context) OrganizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationMapOutput)
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
func (o OrganizationOutput) AnalyticsRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.AnalyticsRegion }).(pulumi.StringPtrOutput)
}

// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
func (o OrganizationOutput) AuthorizedNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.AuthorizedNetwork }).(pulumi.StringPtrOutput)
}

// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
func (o OrganizationOutput) BillingType() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.BillingType }).(pulumi.StringOutput)
}

// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
// is CLOUD. A base64-encoded string.
func (o OrganizationOutput) CaCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.CaCertificate }).(pulumi.StringOutput)
}

// Description of the Apigee organization.
func (o OrganizationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the Apigee organization.
func (o OrganizationOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Name of the property.
func (o OrganizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project ID associated with the Apigee organization.
func (o OrganizationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Properties defined in the Apigee organization profile.
// Structure is documented below.
func (o OrganizationOutput) Properties() OrganizationPropertiesOutput {
	return o.ApplyT(func(v *Organization) OrganizationPropertiesOutput { return v.Properties }).(OrganizationPropertiesOutput)
}

// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
// operation completes. During this period, the Organization may be restored to its last known state.
// After this period, the Organization will no longer be able to be restored.
// Default value is `DELETION_RETENTION_UNSPECIFIED`.
// Possible values are `DELETION_RETENTION_UNSPECIFIED` and `MINIMUM`.
func (o OrganizationOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.Retention }).(pulumi.StringPtrOutput)
}

// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
// Update is not allowed after the organization is created.
// If not specified, a Google-Managed encryption key will be used.
// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
func (o OrganizationOutput) RuntimeDatabaseEncryptionKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.RuntimeDatabaseEncryptionKeyName }).(pulumi.StringPtrOutput)
}

// Runtime type of the Apigee organization based on the Apigee subscription purchased.
// Default value is `CLOUD`.
// Possible values are `CLOUD` and `HYBRID`.
func (o OrganizationOutput) RuntimeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.RuntimeType }).(pulumi.StringPtrOutput)
}

// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
// purposes only) or paid (full subscription has been purchased).
func (o OrganizationOutput) SubscriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.SubscriptionType }).(pulumi.StringOutput)
}

type OrganizationArrayOutput struct{ *pulumi.OutputState }

func (OrganizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Organization)(nil)).Elem()
}

func (o OrganizationArrayOutput) ToOrganizationArrayOutput() OrganizationArrayOutput {
	return o
}

func (o OrganizationArrayOutput) ToOrganizationArrayOutputWithContext(ctx context.Context) OrganizationArrayOutput {
	return o
}

func (o OrganizationArrayOutput) Index(i pulumi.IntInput) OrganizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Organization {
		return vs[0].([]*Organization)[vs[1].(int)]
	}).(OrganizationOutput)
}

type OrganizationMapOutput struct{ *pulumi.OutputState }

func (OrganizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Organization)(nil)).Elem()
}

func (o OrganizationMapOutput) ToOrganizationMapOutput() OrganizationMapOutput {
	return o
}

func (o OrganizationMapOutput) ToOrganizationMapOutputWithContext(ctx context.Context) OrganizationMapOutput {
	return o
}

func (o OrganizationMapOutput) MapIndex(k pulumi.StringInput) OrganizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Organization {
		return vs[0].(map[string]*Organization)[vs[1].(string)]
	}).(OrganizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationInput)(nil)).Elem(), &Organization{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationArrayInput)(nil)).Elem(), OrganizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMapInput)(nil)).Elem(), OrganizationMap{})
	pulumi.RegisterOutputType(OrganizationOutput{})
	pulumi.RegisterOutputType(OrganizationArrayOutput{})
	pulumi.RegisterOutputType(OrganizationMapOutput{})
}
