// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows configuring a single access level condition to be appended to an access level's conditions.
// This resource is intended to be used in cases where it is not possible to compile a full list
// of conditions to include in a `accesscontextmanager.AccessLevel` resource,
// to enable them to be added separately.
//
// > **Note:** If this resource is used alongside a `accesscontextmanager.AccessLevel` resource,
// the access level resource must have a `lifecycle` block with `ignoreChanges = [basic[0].conditions]` so
// they don't fight over which service accounts should be included.
//
// To get more information about AccessLevelCondition, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
// * How-to Guides
//   - [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
// ### Access Context Manager Access Level Condition Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := accesscontextmanager.NewAccessPolicy(ctx, "access-policy", &accesscontextmanager.AccessPolicyArgs{
//				Parent: pulumi.String("organizations/123456789"),
//				Title:  pulumi.String("my policy"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = accesscontextmanager.NewAccessLevel(ctx, "access-level-service-account", &accesscontextmanager.AccessLevelArgs{
//				Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("accessPolicies/%v", name), nil
//				}).(pulumi.StringOutput),
//				Title: pulumi.String("tf_test_chromeos_no_lock"),
//				Basic: &accesscontextmanager.AccessLevelBasicArgs{
//					Conditions: accesscontextmanager.AccessLevelBasicConditionArray{
//						&accesscontextmanager.AccessLevelBasicConditionArgs{
//							DevicePolicy: &accesscontextmanager.AccessLevelBasicConditionDevicePolicyArgs{
//								RequireScreenLock: pulumi.Bool(true),
//								OsConstraints: accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArray{
//									&accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArgs{
//										OsType: pulumi.String("DESKTOP_CHROME_OS"),
//									},
//								},
//							},
//							Regions: pulumi.StringArray{
//								pulumi.String("CH"),
//								pulumi.String("IT"),
//								pulumi.String("US"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceAccount.NewAccount(ctx, "created-later", &serviceAccount.AccountArgs{
//				AccountId: pulumi.String("tf-test"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = accesscontextmanager.NewAccessLevelCondition(ctx, "access-level-conditions", &accesscontextmanager.AccessLevelConditionArgs{
//				AccessLevel: access_level_service_account.Name,
//				IpSubnetworks: pulumi.StringArray{
//					pulumi.String("192.0.4.0/24"),
//				},
//				Members: pulumi.StringArray{
//					pulumi.String("user:test@google.com"),
//					pulumi.String("user:test2@google.com"),
//					created_later.Email.ApplyT(func(email string) (string, error) {
//						return fmt.Sprintf("serviceAccount:%v", email), nil
//					}).(pulumi.StringOutput),
//				},
//				Negate: pulumi.Bool(false),
//				DevicePolicy: &accesscontextmanager.AccessLevelConditionDevicePolicyArgs{
//					RequireScreenLock:    pulumi.Bool(false),
//					RequireAdminApproval: pulumi.Bool(false),
//					RequireCorpOwned:     pulumi.Bool(true),
//					OsConstraints: accesscontextmanager.AccessLevelConditionDevicePolicyOsConstraintArray{
//						&accesscontextmanager.AccessLevelConditionDevicePolicyOsConstraintArgs{
//							OsType: pulumi.String("DESKTOP_CHROME_OS"),
//						},
//					},
//				},
//				Regions: pulumi.StringArray{
//					pulumi.String("IT"),
//					pulumi.String("US"),
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
// This resource does not support import.
type AccessLevelCondition struct {
	pulumi.CustomResourceState

	// The name of the Access Level to add this condition to.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy AccessLevelConditionDevicePolicyPtrOutput `pulumi:"devicePolicy"`
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks pulumi.StringArrayOutput `pulumi:"ipSubnetworks"`
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate pulumi.BoolPtrOutput `pulumi:"negate"`
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels pulumi.StringArrayOutput `pulumi:"requiredAccessLevels"`
}

// NewAccessLevelCondition registers a new resource with the given unique name, arguments, and options.
func NewAccessLevelCondition(ctx *pulumi.Context,
	name string, args *AccessLevelConditionArgs, opts ...pulumi.ResourceOption) (*AccessLevelCondition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessLevel == nil {
		return nil, errors.New("invalid value for required argument 'AccessLevel'")
	}
	var resource AccessLevelCondition
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessLevelCondition:AccessLevelCondition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLevelCondition gets an existing AccessLevelCondition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLevelCondition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLevelConditionState, opts ...pulumi.ResourceOption) (*AccessLevelCondition, error) {
	var resource AccessLevelCondition
	err := ctx.ReadResource("gcp:accesscontextmanager/accessLevelCondition:AccessLevelCondition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLevelCondition resources.
type accessLevelConditionState struct {
	// The name of the Access Level to add this condition to.
	AccessLevel *string `pulumi:"accessLevel"`
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy *AccessLevelConditionDevicePolicy `pulumi:"devicePolicy"`
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks []string `pulumi:"ipSubnetworks"`
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members []string `pulumi:"members"`
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `pulumi:"negate"`
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []string `pulumi:"regions"`
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []string `pulumi:"requiredAccessLevels"`
}

type AccessLevelConditionState struct {
	// The name of the Access Level to add this condition to.
	AccessLevel pulumi.StringPtrInput
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy AccessLevelConditionDevicePolicyPtrInput
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks pulumi.StringArrayInput
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members pulumi.StringArrayInput
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate pulumi.BoolPtrInput
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions pulumi.StringArrayInput
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels pulumi.StringArrayInput
}

func (AccessLevelConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelConditionState)(nil)).Elem()
}

type accessLevelConditionArgs struct {
	// The name of the Access Level to add this condition to.
	AccessLevel string `pulumi:"accessLevel"`
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy *AccessLevelConditionDevicePolicy `pulumi:"devicePolicy"`
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks []string `pulumi:"ipSubnetworks"`
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members []string `pulumi:"members"`
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `pulumi:"negate"`
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []string `pulumi:"regions"`
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []string `pulumi:"requiredAccessLevels"`
}

// The set of arguments for constructing a AccessLevelCondition resource.
type AccessLevelConditionArgs struct {
	// The name of the Access Level to add this condition to.
	AccessLevel pulumi.StringInput
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy AccessLevelConditionDevicePolicyPtrInput
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks pulumi.StringArrayInput
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members pulumi.StringArrayInput
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate pulumi.BoolPtrInput
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions pulumi.StringArrayInput
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels pulumi.StringArrayInput
}

func (AccessLevelConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelConditionArgs)(nil)).Elem()
}

type AccessLevelConditionInput interface {
	pulumi.Input

	ToAccessLevelConditionOutput() AccessLevelConditionOutput
	ToAccessLevelConditionOutputWithContext(ctx context.Context) AccessLevelConditionOutput
}

func (*AccessLevelCondition) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLevelCondition)(nil)).Elem()
}

func (i *AccessLevelCondition) ToAccessLevelConditionOutput() AccessLevelConditionOutput {
	return i.ToAccessLevelConditionOutputWithContext(context.Background())
}

func (i *AccessLevelCondition) ToAccessLevelConditionOutputWithContext(ctx context.Context) AccessLevelConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelConditionOutput)
}

// AccessLevelConditionArrayInput is an input type that accepts AccessLevelConditionArray and AccessLevelConditionArrayOutput values.
// You can construct a concrete instance of `AccessLevelConditionArrayInput` via:
//
//	AccessLevelConditionArray{ AccessLevelConditionArgs{...} }
type AccessLevelConditionArrayInput interface {
	pulumi.Input

	ToAccessLevelConditionArrayOutput() AccessLevelConditionArrayOutput
	ToAccessLevelConditionArrayOutputWithContext(context.Context) AccessLevelConditionArrayOutput
}

type AccessLevelConditionArray []AccessLevelConditionInput

func (AccessLevelConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessLevelCondition)(nil)).Elem()
}

func (i AccessLevelConditionArray) ToAccessLevelConditionArrayOutput() AccessLevelConditionArrayOutput {
	return i.ToAccessLevelConditionArrayOutputWithContext(context.Background())
}

func (i AccessLevelConditionArray) ToAccessLevelConditionArrayOutputWithContext(ctx context.Context) AccessLevelConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelConditionArrayOutput)
}

// AccessLevelConditionMapInput is an input type that accepts AccessLevelConditionMap and AccessLevelConditionMapOutput values.
// You can construct a concrete instance of `AccessLevelConditionMapInput` via:
//
//	AccessLevelConditionMap{ "key": AccessLevelConditionArgs{...} }
type AccessLevelConditionMapInput interface {
	pulumi.Input

	ToAccessLevelConditionMapOutput() AccessLevelConditionMapOutput
	ToAccessLevelConditionMapOutputWithContext(context.Context) AccessLevelConditionMapOutput
}

type AccessLevelConditionMap map[string]AccessLevelConditionInput

func (AccessLevelConditionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessLevelCondition)(nil)).Elem()
}

func (i AccessLevelConditionMap) ToAccessLevelConditionMapOutput() AccessLevelConditionMapOutput {
	return i.ToAccessLevelConditionMapOutputWithContext(context.Background())
}

func (i AccessLevelConditionMap) ToAccessLevelConditionMapOutputWithContext(ctx context.Context) AccessLevelConditionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelConditionMapOutput)
}

type AccessLevelConditionOutput struct{ *pulumi.OutputState }

func (AccessLevelConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLevelCondition)(nil)).Elem()
}

func (o AccessLevelConditionOutput) ToAccessLevelConditionOutput() AccessLevelConditionOutput {
	return o
}

func (o AccessLevelConditionOutput) ToAccessLevelConditionOutputWithContext(ctx context.Context) AccessLevelConditionOutput {
	return o
}

// The name of the Access Level to add this condition to.
func (o AccessLevelConditionOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLevelCondition) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// Device specific restrictions, all restrictions must hold for
// the Condition to be true. If not specified, all devices are
// allowed.
// Structure is documented below.
func (o AccessLevelConditionOutput) DevicePolicy() AccessLevelConditionDevicePolicyPtrOutput {
	return o.ApplyT(func(v *AccessLevelCondition) AccessLevelConditionDevicePolicyPtrOutput { return v.DevicePolicy }).(AccessLevelConditionDevicePolicyPtrOutput)
}

// A list of CIDR block IP subnetwork specification. May be IPv4
// or IPv6.
// Note that for a CIDR IP address block, the specified IP address
// portion must be properly truncated (i.e. all the host bits must
// be zero) or the input is considered malformed. For example,
// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
// is not. The originating IP of a request must be in one of the
// listed subnets in order for this Condition to be true.
// If empty, all IP addresses are allowed.
func (o AccessLevelConditionOutput) IpSubnetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessLevelCondition) pulumi.StringArrayOutput { return v.IpSubnetworks }).(pulumi.StringArrayOutput)
}

// An allowed list of members (users, service accounts).
// Using groups is not supported yet.
// The signed-in user originating the request must be a part of one
// of the provided members. If not specified, a request may come
// from any user (logged in/not logged in, not present in any
// groups, etc.).
// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
func (o AccessLevelConditionOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessLevelCondition) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Whether to negate the Condition. If true, the Condition becomes
// a NAND over its non-empty fields, each field must be false for
// the Condition overall to be satisfied. Defaults to false.
func (o AccessLevelConditionOutput) Negate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessLevelCondition) pulumi.BoolPtrOutput { return v.Negate }).(pulumi.BoolPtrOutput)
}

// The request must originate from one of the provided
// countries/regions.
// Format: A valid ISO 3166-1 alpha-2 code.
func (o AccessLevelConditionOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessLevelCondition) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// A list of other access levels defined in the same Policy,
// referenced by resource name. Referencing an AccessLevel which
// does not exist is an error. All access levels listed must be
// granted for the Condition to be true.
// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
func (o AccessLevelConditionOutput) RequiredAccessLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessLevelCondition) pulumi.StringArrayOutput { return v.RequiredAccessLevels }).(pulumi.StringArrayOutput)
}

type AccessLevelConditionArrayOutput struct{ *pulumi.OutputState }

func (AccessLevelConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessLevelCondition)(nil)).Elem()
}

func (o AccessLevelConditionArrayOutput) ToAccessLevelConditionArrayOutput() AccessLevelConditionArrayOutput {
	return o
}

func (o AccessLevelConditionArrayOutput) ToAccessLevelConditionArrayOutputWithContext(ctx context.Context) AccessLevelConditionArrayOutput {
	return o
}

func (o AccessLevelConditionArrayOutput) Index(i pulumi.IntInput) AccessLevelConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessLevelCondition {
		return vs[0].([]*AccessLevelCondition)[vs[1].(int)]
	}).(AccessLevelConditionOutput)
}

type AccessLevelConditionMapOutput struct{ *pulumi.OutputState }

func (AccessLevelConditionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessLevelCondition)(nil)).Elem()
}

func (o AccessLevelConditionMapOutput) ToAccessLevelConditionMapOutput() AccessLevelConditionMapOutput {
	return o
}

func (o AccessLevelConditionMapOutput) ToAccessLevelConditionMapOutputWithContext(ctx context.Context) AccessLevelConditionMapOutput {
	return o
}

func (o AccessLevelConditionMapOutput) MapIndex(k pulumi.StringInput) AccessLevelConditionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessLevelCondition {
		return vs[0].(map[string]*AccessLevelCondition)[vs[1].(string)]
	}).(AccessLevelConditionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLevelConditionInput)(nil)).Elem(), &AccessLevelCondition{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLevelConditionArrayInput)(nil)).Elem(), AccessLevelConditionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLevelConditionMapInput)(nil)).Elem(), AccessLevelConditionMap{})
	pulumi.RegisterOutputType(AccessLevelConditionOutput{})
	pulumi.RegisterOutputType(AccessLevelConditionArrayOutput{})
	pulumi.RegisterOutputType(AccessLevelConditionMapOutput{})
}
