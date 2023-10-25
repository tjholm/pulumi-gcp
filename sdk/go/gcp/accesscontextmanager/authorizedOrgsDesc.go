// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An authorized organizations description describes a list of organizations
// (1) that have been authorized to use certain asset (for example, device) data
// owned by different organizations at the enforcement points, or (2) with certain
// asset (for example, device) have been authorized to access the resources in
// another organization at the enforcement points.
//
// To get more information about AuthorizedOrgsDesc, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.authorizedOrgsDescs)
// * How-to Guides
//   - [gcloud docs](https://cloud.google.com/beyondcorp-enterprise/docs/cross-org-authorization)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
// ### Access Context Manager Authorized Orgs Desc Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := accesscontextmanager.NewAccessPolicy(ctx, "test-access", &accesscontextmanager.AccessPolicyArgs{
//				Parent: pulumi.String("organizations/"),
//				Title:  pulumi.String("my policy"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = accesscontextmanager.NewAuthorizedOrgsDesc(ctx, "authorized-orgs-desc", &accesscontextmanager.AuthorizedOrgsDescArgs{
//				AssetType:              pulumi.String("ASSET_TYPE_CREDENTIAL_STRENGTH"),
//				AuthorizationDirection: pulumi.String("AUTHORIZATION_DIRECTION_TO"),
//				AuthorizationType:      pulumi.String("AUTHORIZATION_TYPE_TRUST"),
//				Orgs: pulumi.StringArray{
//					pulumi.String("organizations/12345"),
//					pulumi.String("organizations/98765"),
//				},
//				Parent: test_access.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("accessPolicies/%v", name), nil
//				}).(pulumi.StringOutput),
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
// AuthorizedOrgsDesc can be imported using any of these accepted formats:
//
// ```sh
//
//	$ pulumi import gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc default {{name}}
//
// ```
type AuthorizedOrgsDesc struct {
	pulumi.CustomResourceState

	// The type of entities that need to use the authorization relationship during
	// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
	// "ASSET_TYPE_CREDENTIAL_STRENGTH".
	// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	AssetType pulumi.StringPtrOutput `pulumi:"assetType"`
	// The direction of the authorization relationship between this organization
	// and the organizations listed in the "orgs" field. The valid values for this
	// field include the following:
	// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
	// in the organizations listed in the `orgs` field.
	// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
	// field to evaluate the traffic in this organization.
	// For the authorization relationship to take effect, all of the organizations
	// must authorize and specify the appropriate relationship direction. For
	// example, if organization A authorized organization B and C to evaluate its
	// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
	// direction, organizations B and C must specify
	// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
	// "AuthorizedOrgsDesc" resource.
	// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
	AuthorizationDirection pulumi.StringPtrOutput `pulumi:"authorizationDirection"`
	// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
	// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
	AuthorizationType pulumi.StringPtrOutput `pulumi:"authorizationType"`
	// Time the AuthorizedOrgsDesc was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Resource name for the `AuthorizedOrgsDesc`. Format:
	// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
	// The `authorizedOrgsDesc` component must begin with a letter, followed by
	// alphanumeric characters or `_`.
	// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of organization ids in this AuthorizedOrgsDesc.
	// Format: `organizations/<org_number>`
	// Example: `organizations/123456`
	Orgs pulumi.StringArrayOutput `pulumi:"orgs"`
	// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Time the AuthorizedOrgsDesc was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAuthorizedOrgsDesc registers a new resource with the given unique name, arguments, and options.
func NewAuthorizedOrgsDesc(ctx *pulumi.Context,
	name string, args *AuthorizedOrgsDescArgs, opts ...pulumi.ResourceOption) (*AuthorizedOrgsDesc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthorizedOrgsDesc
	err := ctx.RegisterResource("gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizedOrgsDesc gets an existing AuthorizedOrgsDesc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizedOrgsDesc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizedOrgsDescState, opts ...pulumi.ResourceOption) (*AuthorizedOrgsDesc, error) {
	var resource AuthorizedOrgsDesc
	err := ctx.ReadResource("gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizedOrgsDesc resources.
type authorizedOrgsDescState struct {
	// The type of entities that need to use the authorization relationship during
	// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
	// "ASSET_TYPE_CREDENTIAL_STRENGTH".
	// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	AssetType *string `pulumi:"assetType"`
	// The direction of the authorization relationship between this organization
	// and the organizations listed in the "orgs" field. The valid values for this
	// field include the following:
	// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
	// in the organizations listed in the `orgs` field.
	// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
	// field to evaluate the traffic in this organization.
	// For the authorization relationship to take effect, all of the organizations
	// must authorize and specify the appropriate relationship direction. For
	// example, if organization A authorized organization B and C to evaluate its
	// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
	// direction, organizations B and C must specify
	// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
	// "AuthorizedOrgsDesc" resource.
	// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
	AuthorizationDirection *string `pulumi:"authorizationDirection"`
	// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
	// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
	AuthorizationType *string `pulumi:"authorizationType"`
	// Time the AuthorizedOrgsDesc was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Resource name for the `AuthorizedOrgsDesc`. Format:
	// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
	// The `authorizedOrgsDesc` component must begin with a letter, followed by
	// alphanumeric characters or `_`.
	// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
	//
	// ***
	Name *string `pulumi:"name"`
	// The list of organization ids in this AuthorizedOrgsDesc.
	// Format: `organizations/<org_number>`
	// Example: `organizations/123456`
	Orgs []string `pulumi:"orgs"`
	// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
	Parent *string `pulumi:"parent"`
	// Time the AuthorizedOrgsDesc was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type AuthorizedOrgsDescState struct {
	// The type of entities that need to use the authorization relationship during
	// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
	// "ASSET_TYPE_CREDENTIAL_STRENGTH".
	// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	AssetType pulumi.StringPtrInput
	// The direction of the authorization relationship between this organization
	// and the organizations listed in the "orgs" field. The valid values for this
	// field include the following:
	// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
	// in the organizations listed in the `orgs` field.
	// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
	// field to evaluate the traffic in this organization.
	// For the authorization relationship to take effect, all of the organizations
	// must authorize and specify the appropriate relationship direction. For
	// example, if organization A authorized organization B and C to evaluate its
	// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
	// direction, organizations B and C must specify
	// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
	// "AuthorizedOrgsDesc" resource.
	// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
	AuthorizationDirection pulumi.StringPtrInput
	// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
	// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
	AuthorizationType pulumi.StringPtrInput
	// Time the AuthorizedOrgsDesc was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Resource name for the `AuthorizedOrgsDesc`. Format:
	// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
	// The `authorizedOrgsDesc` component must begin with a letter, followed by
	// alphanumeric characters or `_`.
	// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
	//
	// ***
	Name pulumi.StringPtrInput
	// The list of organization ids in this AuthorizedOrgsDesc.
	// Format: `organizations/<org_number>`
	// Example: `organizations/123456`
	Orgs pulumi.StringArrayInput
	// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
	Parent pulumi.StringPtrInput
	// Time the AuthorizedOrgsDesc was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (AuthorizedOrgsDescState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizedOrgsDescState)(nil)).Elem()
}

type authorizedOrgsDescArgs struct {
	// The type of entities that need to use the authorization relationship during
	// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
	// "ASSET_TYPE_CREDENTIAL_STRENGTH".
	// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	AssetType *string `pulumi:"assetType"`
	// The direction of the authorization relationship between this organization
	// and the organizations listed in the "orgs" field. The valid values for this
	// field include the following:
	// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
	// in the organizations listed in the `orgs` field.
	// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
	// field to evaluate the traffic in this organization.
	// For the authorization relationship to take effect, all of the organizations
	// must authorize and specify the appropriate relationship direction. For
	// example, if organization A authorized organization B and C to evaluate its
	// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
	// direction, organizations B and C must specify
	// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
	// "AuthorizedOrgsDesc" resource.
	// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
	AuthorizationDirection *string `pulumi:"authorizationDirection"`
	// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
	// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
	AuthorizationType *string `pulumi:"authorizationType"`
	// Resource name for the `AuthorizedOrgsDesc`. Format:
	// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
	// The `authorizedOrgsDesc` component must begin with a letter, followed by
	// alphanumeric characters or `_`.
	// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
	//
	// ***
	Name *string `pulumi:"name"`
	// The list of organization ids in this AuthorizedOrgsDesc.
	// Format: `organizations/<org_number>`
	// Example: `organizations/123456`
	Orgs []string `pulumi:"orgs"`
	// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
	Parent string `pulumi:"parent"`
}

// The set of arguments for constructing a AuthorizedOrgsDesc resource.
type AuthorizedOrgsDescArgs struct {
	// The type of entities that need to use the authorization relationship during
	// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
	// "ASSET_TYPE_CREDENTIAL_STRENGTH".
	// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	AssetType pulumi.StringPtrInput
	// The direction of the authorization relationship between this organization
	// and the organizations listed in the "orgs" field. The valid values for this
	// field include the following:
	// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
	// in the organizations listed in the `orgs` field.
	// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
	// field to evaluate the traffic in this organization.
	// For the authorization relationship to take effect, all of the organizations
	// must authorize and specify the appropriate relationship direction. For
	// example, if organization A authorized organization B and C to evaluate its
	// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
	// direction, organizations B and C must specify
	// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
	// "AuthorizedOrgsDesc" resource.
	// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
	AuthorizationDirection pulumi.StringPtrInput
	// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
	// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
	AuthorizationType pulumi.StringPtrInput
	// Resource name for the `AuthorizedOrgsDesc`. Format:
	// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
	// The `authorizedOrgsDesc` component must begin with a letter, followed by
	// alphanumeric characters or `_`.
	// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
	//
	// ***
	Name pulumi.StringPtrInput
	// The list of organization ids in this AuthorizedOrgsDesc.
	// Format: `organizations/<org_number>`
	// Example: `organizations/123456`
	Orgs pulumi.StringArrayInput
	// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
	Parent pulumi.StringInput
}

func (AuthorizedOrgsDescArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizedOrgsDescArgs)(nil)).Elem()
}

type AuthorizedOrgsDescInput interface {
	pulumi.Input

	ToAuthorizedOrgsDescOutput() AuthorizedOrgsDescOutput
	ToAuthorizedOrgsDescOutputWithContext(ctx context.Context) AuthorizedOrgsDescOutput
}

func (*AuthorizedOrgsDesc) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizedOrgsDesc)(nil)).Elem()
}

func (i *AuthorizedOrgsDesc) ToAuthorizedOrgsDescOutput() AuthorizedOrgsDescOutput {
	return i.ToAuthorizedOrgsDescOutputWithContext(context.Background())
}

func (i *AuthorizedOrgsDesc) ToAuthorizedOrgsDescOutputWithContext(ctx context.Context) AuthorizedOrgsDescOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizedOrgsDescOutput)
}

func (i *AuthorizedOrgsDesc) ToOutput(ctx context.Context) pulumix.Output[*AuthorizedOrgsDesc] {
	return pulumix.Output[*AuthorizedOrgsDesc]{
		OutputState: i.ToAuthorizedOrgsDescOutputWithContext(ctx).OutputState,
	}
}

// AuthorizedOrgsDescArrayInput is an input type that accepts AuthorizedOrgsDescArray and AuthorizedOrgsDescArrayOutput values.
// You can construct a concrete instance of `AuthorizedOrgsDescArrayInput` via:
//
//	AuthorizedOrgsDescArray{ AuthorizedOrgsDescArgs{...} }
type AuthorizedOrgsDescArrayInput interface {
	pulumi.Input

	ToAuthorizedOrgsDescArrayOutput() AuthorizedOrgsDescArrayOutput
	ToAuthorizedOrgsDescArrayOutputWithContext(context.Context) AuthorizedOrgsDescArrayOutput
}

type AuthorizedOrgsDescArray []AuthorizedOrgsDescInput

func (AuthorizedOrgsDescArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthorizedOrgsDesc)(nil)).Elem()
}

func (i AuthorizedOrgsDescArray) ToAuthorizedOrgsDescArrayOutput() AuthorizedOrgsDescArrayOutput {
	return i.ToAuthorizedOrgsDescArrayOutputWithContext(context.Background())
}

func (i AuthorizedOrgsDescArray) ToAuthorizedOrgsDescArrayOutputWithContext(ctx context.Context) AuthorizedOrgsDescArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizedOrgsDescArrayOutput)
}

func (i AuthorizedOrgsDescArray) ToOutput(ctx context.Context) pulumix.Output[[]*AuthorizedOrgsDesc] {
	return pulumix.Output[[]*AuthorizedOrgsDesc]{
		OutputState: i.ToAuthorizedOrgsDescArrayOutputWithContext(ctx).OutputState,
	}
}

// AuthorizedOrgsDescMapInput is an input type that accepts AuthorizedOrgsDescMap and AuthorizedOrgsDescMapOutput values.
// You can construct a concrete instance of `AuthorizedOrgsDescMapInput` via:
//
//	AuthorizedOrgsDescMap{ "key": AuthorizedOrgsDescArgs{...} }
type AuthorizedOrgsDescMapInput interface {
	pulumi.Input

	ToAuthorizedOrgsDescMapOutput() AuthorizedOrgsDescMapOutput
	ToAuthorizedOrgsDescMapOutputWithContext(context.Context) AuthorizedOrgsDescMapOutput
}

type AuthorizedOrgsDescMap map[string]AuthorizedOrgsDescInput

func (AuthorizedOrgsDescMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthorizedOrgsDesc)(nil)).Elem()
}

func (i AuthorizedOrgsDescMap) ToAuthorizedOrgsDescMapOutput() AuthorizedOrgsDescMapOutput {
	return i.ToAuthorizedOrgsDescMapOutputWithContext(context.Background())
}

func (i AuthorizedOrgsDescMap) ToAuthorizedOrgsDescMapOutputWithContext(ctx context.Context) AuthorizedOrgsDescMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizedOrgsDescMapOutput)
}

func (i AuthorizedOrgsDescMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AuthorizedOrgsDesc] {
	return pulumix.Output[map[string]*AuthorizedOrgsDesc]{
		OutputState: i.ToAuthorizedOrgsDescMapOutputWithContext(ctx).OutputState,
	}
}

type AuthorizedOrgsDescOutput struct{ *pulumi.OutputState }

func (AuthorizedOrgsDescOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizedOrgsDesc)(nil)).Elem()
}

func (o AuthorizedOrgsDescOutput) ToAuthorizedOrgsDescOutput() AuthorizedOrgsDescOutput {
	return o
}

func (o AuthorizedOrgsDescOutput) ToAuthorizedOrgsDescOutputWithContext(ctx context.Context) AuthorizedOrgsDescOutput {
	return o
}

func (o AuthorizedOrgsDescOutput) ToOutput(ctx context.Context) pulumix.Output[*AuthorizedOrgsDesc] {
	return pulumix.Output[*AuthorizedOrgsDesc]{
		OutputState: o.OutputState,
	}
}

// The type of entities that need to use the authorization relationship during
// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
// "ASSET_TYPE_CREDENTIAL_STRENGTH".
// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
func (o AuthorizedOrgsDescOutput) AssetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringPtrOutput { return v.AssetType }).(pulumi.StringPtrOutput)
}

// The direction of the authorization relationship between this organization
// and the organizations listed in the "orgs" field. The valid values for this
// field include the following:
// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
// in the organizations listed in the `orgs` field.
// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
// field to evaluate the traffic in this organization.
// For the authorization relationship to take effect, all of the organizations
// must authorize and specify the appropriate relationship direction. For
// example, if organization A authorized organization B and C to evaluate its
// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
// direction, organizations B and C must specify
// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
// "AuthorizedOrgsDesc" resource.
// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
func (o AuthorizedOrgsDescOutput) AuthorizationDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringPtrOutput { return v.AuthorizationDirection }).(pulumi.StringPtrOutput)
}

// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
func (o AuthorizedOrgsDescOutput) AuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringPtrOutput { return v.AuthorizationType }).(pulumi.StringPtrOutput)
}

// Time the AuthorizedOrgsDesc was created in UTC.
func (o AuthorizedOrgsDescOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Resource name for the `AuthorizedOrgsDesc`. Format:
// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
// The `authorizedOrgsDesc` component must begin with a letter, followed by
// alphanumeric characters or `_`.
// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
//
// ***
func (o AuthorizedOrgsDescOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of organization ids in this AuthorizedOrgsDesc.
// Format: `organizations/<org_number>`
// Example: `organizations/123456`
func (o AuthorizedOrgsDescOutput) Orgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringArrayOutput { return v.Orgs }).(pulumi.StringArrayOutput)
}

// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
func (o AuthorizedOrgsDescOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Time the AuthorizedOrgsDesc was updated in UTC.
func (o AuthorizedOrgsDescOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizedOrgsDesc) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type AuthorizedOrgsDescArrayOutput struct{ *pulumi.OutputState }

func (AuthorizedOrgsDescArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthorizedOrgsDesc)(nil)).Elem()
}

func (o AuthorizedOrgsDescArrayOutput) ToAuthorizedOrgsDescArrayOutput() AuthorizedOrgsDescArrayOutput {
	return o
}

func (o AuthorizedOrgsDescArrayOutput) ToAuthorizedOrgsDescArrayOutputWithContext(ctx context.Context) AuthorizedOrgsDescArrayOutput {
	return o
}

func (o AuthorizedOrgsDescArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AuthorizedOrgsDesc] {
	return pulumix.Output[[]*AuthorizedOrgsDesc]{
		OutputState: o.OutputState,
	}
}

func (o AuthorizedOrgsDescArrayOutput) Index(i pulumi.IntInput) AuthorizedOrgsDescOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthorizedOrgsDesc {
		return vs[0].([]*AuthorizedOrgsDesc)[vs[1].(int)]
	}).(AuthorizedOrgsDescOutput)
}

type AuthorizedOrgsDescMapOutput struct{ *pulumi.OutputState }

func (AuthorizedOrgsDescMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthorizedOrgsDesc)(nil)).Elem()
}

func (o AuthorizedOrgsDescMapOutput) ToAuthorizedOrgsDescMapOutput() AuthorizedOrgsDescMapOutput {
	return o
}

func (o AuthorizedOrgsDescMapOutput) ToAuthorizedOrgsDescMapOutputWithContext(ctx context.Context) AuthorizedOrgsDescMapOutput {
	return o
}

func (o AuthorizedOrgsDescMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AuthorizedOrgsDesc] {
	return pulumix.Output[map[string]*AuthorizedOrgsDesc]{
		OutputState: o.OutputState,
	}
}

func (o AuthorizedOrgsDescMapOutput) MapIndex(k pulumi.StringInput) AuthorizedOrgsDescOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthorizedOrgsDesc {
		return vs[0].(map[string]*AuthorizedOrgsDesc)[vs[1].(string)]
	}).(AuthorizedOrgsDescOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizedOrgsDescInput)(nil)).Elem(), &AuthorizedOrgsDesc{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizedOrgsDescArrayInput)(nil)).Elem(), AuthorizedOrgsDescArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizedOrgsDescMapInput)(nil)).Elem(), AuthorizedOrgsDescMap{})
	pulumi.RegisterOutputType(AuthorizedOrgsDescOutput{})
	pulumi.RegisterOutputType(AuthorizedOrgsDescArrayOutput{})
	pulumi.RegisterOutputType(AuthorizedOrgsDescMapOutput{})
}
