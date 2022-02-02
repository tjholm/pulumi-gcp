// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replace all existing Access Levels in an Access Policy with the Access Levels provided. This is done atomically.
// This is a bulk edit of all Access Levels and may override existing Access Levels created by `accesscontextmanager.AccessLevel`,
// thus causing a permadiff if used alongside `accesscontextmanager.AccessLevel` on the same parent.
//
// To get more information about AccessLevels, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
// * How-to Guides
//     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
//
// ## Example Usage
// ### Access Context Manager Access Levels Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := accesscontextmanager.NewAccessPolicy(ctx, "access-policy", &accesscontextmanager.AccessPolicyArgs{
// 			Parent: pulumi.String("organizations/123456789"),
// 			Title:  pulumi.String("my policy"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewAccessLevels(ctx, "access-levels", &accesscontextmanager.AccessLevelsArgs{
// 			AccessLevels: accesscontextmanager.AccessLevelsAccessLevelArray{
// 				&accesscontextmanager.AccessLevelsAccessLevelArgs{
// 					Basic: &accesscontextmanager.AccessLevelsAccessLevelBasicArgs{
// 						Conditions: accesscontextmanager.AccessLevelsAccessLevelBasicConditionArray{
// 							&accesscontextmanager.AccessLevelsAccessLevelBasicConditionArgs{
// 								DevicePolicy: &accesscontextmanager.AccessLevelsAccessLevelBasicConditionDevicePolicyArgs{
// 									OsConstraints: accesscontextmanager.AccessLevelsAccessLevelBasicConditionDevicePolicyOsConstraintArray{
// 										&accesscontextmanager.AccessLevelsAccessLevelBasicConditionDevicePolicyOsConstraintArgs{
// 											OsType: pulumi.String("DESKTOP_CHROME_OS"),
// 										},
// 									},
// 									RequireScreenLock: pulumi.Bool(true),
// 								},
// 								Regions: pulumi.StringArray{
// 									pulumi.String("CH"),
// 									pulumi.String("IT"),
// 									pulumi.String("US"),
// 								},
// 							},
// 						},
// 					},
// 					Name: access_policy.Name.ApplyT(func(name string) (string, error) {
// 						return fmt.Sprintf("%v%v%v", "accessPolicies/", name, "/accessLevels/chromeos_no_lock"), nil
// 					}).(pulumi.StringOutput),
// 					Title: pulumi.String("chromeos_no_lock"),
// 				},
// 				&accesscontextmanager.AccessLevelsAccessLevelArgs{
// 					Basic: &accesscontextmanager.AccessLevelsAccessLevelBasicArgs{
// 						Conditions: accesscontextmanager.AccessLevelsAccessLevelBasicConditionArray{
// 							&accesscontextmanager.AccessLevelsAccessLevelBasicConditionArgs{
// 								DevicePolicy: &accesscontextmanager.AccessLevelsAccessLevelBasicConditionDevicePolicyArgs{
// 									OsConstraints: accesscontextmanager.AccessLevelsAccessLevelBasicConditionDevicePolicyOsConstraintArray{
// 										&accesscontextmanager.AccessLevelsAccessLevelBasicConditionDevicePolicyOsConstraintArgs{
// 											OsType: pulumi.String("DESKTOP_MAC"),
// 										},
// 									},
// 									RequireScreenLock: pulumi.Bool(true),
// 								},
// 								Regions: pulumi.StringArray{
// 									pulumi.String("CH"),
// 									pulumi.String("IT"),
// 									pulumi.String("US"),
// 								},
// 							},
// 						},
// 					},
// 					Name: access_policy.Name.ApplyT(func(name string) (string, error) {
// 						return fmt.Sprintf("%v%v%v", "accessPolicies/", name, "/accessLevels/mac_no_lock"), nil
// 					}).(pulumi.StringOutput),
// 					Title: pulumi.String("mac_no_lock"),
// 				},
// 			},
// 			Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "accessPolicies/", name), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// AccessLevels can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/accessLevels:AccessLevels default {{parent}}/accessLevels
// ```
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/accessLevels:AccessLevels default {{parent}}
// ```
type AccessLevels struct {
	pulumi.CustomResourceState

	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels AccessLevelsAccessLevelArrayOutput `pulumi:"accessLevels"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
}

// NewAccessLevels registers a new resource with the given unique name, arguments, and options.
func NewAccessLevels(ctx *pulumi.Context,
	name string, args *AccessLevelsArgs, opts ...pulumi.ResourceOption) (*AccessLevels, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource AccessLevels
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessLevels:AccessLevels", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLevels gets an existing AccessLevels resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLevels(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLevelsState, opts ...pulumi.ResourceOption) (*AccessLevels, error) {
	var resource AccessLevels
	err := ctx.ReadResource("gcp:accesscontextmanager/accessLevels:AccessLevels", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLevels resources.
type accessLevelsState struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels []AccessLevelsAccessLevel `pulumi:"accessLevels"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `pulumi:"parent"`
}

type AccessLevelsState struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels AccessLevelsAccessLevelArrayInput
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringPtrInput
}

func (AccessLevelsState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelsState)(nil)).Elem()
}

type accessLevelsArgs struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels []AccessLevelsAccessLevel `pulumi:"accessLevels"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent string `pulumi:"parent"`
}

// The set of arguments for constructing a AccessLevels resource.
type AccessLevelsArgs struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels AccessLevelsAccessLevelArrayInput
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringInput
}

func (AccessLevelsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelsArgs)(nil)).Elem()
}

type AccessLevelsInput interface {
	pulumi.Input

	ToAccessLevelsOutput() AccessLevelsOutput
	ToAccessLevelsOutputWithContext(ctx context.Context) AccessLevelsOutput
}

func (*AccessLevels) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLevels)(nil)).Elem()
}

func (i *AccessLevels) ToAccessLevelsOutput() AccessLevelsOutput {
	return i.ToAccessLevelsOutputWithContext(context.Background())
}

func (i *AccessLevels) ToAccessLevelsOutputWithContext(ctx context.Context) AccessLevelsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelsOutput)
}

// AccessLevelsArrayInput is an input type that accepts AccessLevelsArray and AccessLevelsArrayOutput values.
// You can construct a concrete instance of `AccessLevelsArrayInput` via:
//
//          AccessLevelsArray{ AccessLevelsArgs{...} }
type AccessLevelsArrayInput interface {
	pulumi.Input

	ToAccessLevelsArrayOutput() AccessLevelsArrayOutput
	ToAccessLevelsArrayOutputWithContext(context.Context) AccessLevelsArrayOutput
}

type AccessLevelsArray []AccessLevelsInput

func (AccessLevelsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessLevels)(nil)).Elem()
}

func (i AccessLevelsArray) ToAccessLevelsArrayOutput() AccessLevelsArrayOutput {
	return i.ToAccessLevelsArrayOutputWithContext(context.Background())
}

func (i AccessLevelsArray) ToAccessLevelsArrayOutputWithContext(ctx context.Context) AccessLevelsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelsArrayOutput)
}

// AccessLevelsMapInput is an input type that accepts AccessLevelsMap and AccessLevelsMapOutput values.
// You can construct a concrete instance of `AccessLevelsMapInput` via:
//
//          AccessLevelsMap{ "key": AccessLevelsArgs{...} }
type AccessLevelsMapInput interface {
	pulumi.Input

	ToAccessLevelsMapOutput() AccessLevelsMapOutput
	ToAccessLevelsMapOutputWithContext(context.Context) AccessLevelsMapOutput
}

type AccessLevelsMap map[string]AccessLevelsInput

func (AccessLevelsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessLevels)(nil)).Elem()
}

func (i AccessLevelsMap) ToAccessLevelsMapOutput() AccessLevelsMapOutput {
	return i.ToAccessLevelsMapOutputWithContext(context.Background())
}

func (i AccessLevelsMap) ToAccessLevelsMapOutputWithContext(ctx context.Context) AccessLevelsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelsMapOutput)
}

type AccessLevelsOutput struct{ *pulumi.OutputState }

func (AccessLevelsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLevels)(nil)).Elem()
}

func (o AccessLevelsOutput) ToAccessLevelsOutput() AccessLevelsOutput {
	return o
}

func (o AccessLevelsOutput) ToAccessLevelsOutputWithContext(ctx context.Context) AccessLevelsOutput {
	return o
}

type AccessLevelsArrayOutput struct{ *pulumi.OutputState }

func (AccessLevelsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessLevels)(nil)).Elem()
}

func (o AccessLevelsArrayOutput) ToAccessLevelsArrayOutput() AccessLevelsArrayOutput {
	return o
}

func (o AccessLevelsArrayOutput) ToAccessLevelsArrayOutputWithContext(ctx context.Context) AccessLevelsArrayOutput {
	return o
}

func (o AccessLevelsArrayOutput) Index(i pulumi.IntInput) AccessLevelsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessLevels {
		return vs[0].([]*AccessLevels)[vs[1].(int)]
	}).(AccessLevelsOutput)
}

type AccessLevelsMapOutput struct{ *pulumi.OutputState }

func (AccessLevelsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessLevels)(nil)).Elem()
}

func (o AccessLevelsMapOutput) ToAccessLevelsMapOutput() AccessLevelsMapOutput {
	return o
}

func (o AccessLevelsMapOutput) ToAccessLevelsMapOutputWithContext(ctx context.Context) AccessLevelsMapOutput {
	return o
}

func (o AccessLevelsMapOutput) MapIndex(k pulumi.StringInput) AccessLevelsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessLevels {
		return vs[0].(map[string]*AccessLevels)[vs[1].(string)]
	}).(AccessLevelsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLevelsInput)(nil)).Elem(), &AccessLevels{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLevelsArrayInput)(nil)).Elem(), AccessLevelsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLevelsMapInput)(nil)).Elem(), AccessLevelsMap{})
	pulumi.RegisterOutputType(AccessLevelsOutput{})
	pulumi.RegisterOutputType(AccessLevelsArrayOutput{})
	pulumi.RegisterOutputType(AccessLevelsMapOutput{})
}
