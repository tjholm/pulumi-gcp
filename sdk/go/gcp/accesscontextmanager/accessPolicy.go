// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AccessPolicy is a container for AccessLevels (which define the necessary
// attributes to use GCP services) and ServicePerimeters (which define
// regions of services able to freely pass data within a perimeter). An
// access policy is globally visible within an organization, and the
// restrictions it specifies apply to all projects within an organization.
//
// To get more information about AccessPolicy, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies)
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
// ### Access Context Manager Access Policy Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := accesscontextmanager.NewAccessPolicy(ctx, "access-policy", &accesscontextmanager.AccessPolicyArgs{
//				Parent: pulumi.String("organizations/123456789"),
//				Title:  pulumi.String("Org Access Policy"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Access Context Manager Access Policy Scoped
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.NewProject(ctx, "project", &organizations.ProjectArgs{
//				OrgId:     pulumi.String("123456789"),
//				ProjectId: pulumi.String("acm-test-proj-123"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = accesscontextmanager.NewAccessPolicy(ctx, "access-policy", &accesscontextmanager.AccessPolicyArgs{
//				Parent: pulumi.String("organizations/123456789"),
//				Scopes: project.Number.ApplyT(func(number string) (string, error) {
//					return fmt.Sprintf("projects/%v", number), nil
//				}).(pulumi.StringOutput),
//				Title: pulumi.String("Scoped Access Policy"),
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
// # AccessPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:accesscontextmanager/accessPolicy:AccessPolicy default {{name}}
//
// ```
type AccessPolicy struct {
	pulumi.CustomResourceState

	// Time the AccessPolicy was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Resource name of the AccessPolicy. Format: {policy_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	Scopes pulumi.StringPtrOutput `pulumi:"scopes"`
	// Human readable title. Does not affect behavior.
	Title pulumi.StringOutput `pulumi:"title"`
	// Time the AccessPolicy was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource AccessPolicy
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessPolicy:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicy gets an existing AccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("gcp:accesscontextmanager/accessPolicy:AccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicy resources.
type accessPolicyState struct {
	// Time the AccessPolicy was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Resource name of the AccessPolicy. Format: {policy_id}
	Name *string `pulumi:"name"`
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent *string `pulumi:"parent"`
	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	Scopes *string `pulumi:"scopes"`
	// Human readable title. Does not affect behavior.
	Title *string `pulumi:"title"`
	// Time the AccessPolicy was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type AccessPolicyState struct {
	// Time the AccessPolicy was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Resource name of the AccessPolicy. Format: {policy_id}
	Name pulumi.StringPtrInput
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent pulumi.StringPtrInput
	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	Scopes pulumi.StringPtrInput
	// Human readable title. Does not affect behavior.
	Title pulumi.StringPtrInput
	// Time the AccessPolicy was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (AccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyState)(nil)).Elem()
}

type accessPolicyArgs struct {
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent string `pulumi:"parent"`
	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	Scopes *string `pulumi:"scopes"`
	// Human readable title. Does not affect behavior.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a AccessPolicy resource.
type AccessPolicyArgs struct {
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent pulumi.StringInput
	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	Scopes pulumi.StringPtrInput
	// Human readable title. Does not affect behavior.
	Title pulumi.StringInput
}

func (AccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyArgs)(nil)).Elem()
}

type AccessPolicyInput interface {
	pulumi.Input

	ToAccessPolicyOutput() AccessPolicyOutput
	ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput
}

func (*AccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil)).Elem()
}

func (i *AccessPolicy) ToAccessPolicyOutput() AccessPolicyOutput {
	return i.ToAccessPolicyOutputWithContext(context.Background())
}

func (i *AccessPolicy) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyOutput)
}

// AccessPolicyArrayInput is an input type that accepts AccessPolicyArray and AccessPolicyArrayOutput values.
// You can construct a concrete instance of `AccessPolicyArrayInput` via:
//
//	AccessPolicyArray{ AccessPolicyArgs{...} }
type AccessPolicyArrayInput interface {
	pulumi.Input

	ToAccessPolicyArrayOutput() AccessPolicyArrayOutput
	ToAccessPolicyArrayOutputWithContext(context.Context) AccessPolicyArrayOutput
}

type AccessPolicyArray []AccessPolicyInput

func (AccessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPolicy)(nil)).Elem()
}

func (i AccessPolicyArray) ToAccessPolicyArrayOutput() AccessPolicyArrayOutput {
	return i.ToAccessPolicyArrayOutputWithContext(context.Background())
}

func (i AccessPolicyArray) ToAccessPolicyArrayOutputWithContext(ctx context.Context) AccessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyArrayOutput)
}

// AccessPolicyMapInput is an input type that accepts AccessPolicyMap and AccessPolicyMapOutput values.
// You can construct a concrete instance of `AccessPolicyMapInput` via:
//
//	AccessPolicyMap{ "key": AccessPolicyArgs{...} }
type AccessPolicyMapInput interface {
	pulumi.Input

	ToAccessPolicyMapOutput() AccessPolicyMapOutput
	ToAccessPolicyMapOutputWithContext(context.Context) AccessPolicyMapOutput
}

type AccessPolicyMap map[string]AccessPolicyInput

func (AccessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPolicy)(nil)).Elem()
}

func (i AccessPolicyMap) ToAccessPolicyMapOutput() AccessPolicyMapOutput {
	return i.ToAccessPolicyMapOutputWithContext(context.Background())
}

func (i AccessPolicyMap) ToAccessPolicyMapOutputWithContext(ctx context.Context) AccessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyMapOutput)
}

type AccessPolicyOutput struct{ *pulumi.OutputState }

func (AccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil)).Elem()
}

func (o AccessPolicyOutput) ToAccessPolicyOutput() AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return o
}

// Time the AccessPolicy was created in UTC.
func (o AccessPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Resource name of the AccessPolicy. Format: {policy_id}
func (o AccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
// Format: organizations/{organization_id}
func (o AccessPolicyOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Folder or project on which this policy is applicable.
// Format: folders/{{folder_id}} or projects/{{project_id}}
func (o AccessPolicyOutput) Scopes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringPtrOutput { return v.Scopes }).(pulumi.StringPtrOutput)
}

// Human readable title. Does not affect behavior.
func (o AccessPolicyOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// Time the AccessPolicy was updated in UTC.
func (o AccessPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type AccessPolicyArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPolicy)(nil)).Elem()
}

func (o AccessPolicyArrayOutput) ToAccessPolicyArrayOutput() AccessPolicyArrayOutput {
	return o
}

func (o AccessPolicyArrayOutput) ToAccessPolicyArrayOutputWithContext(ctx context.Context) AccessPolicyArrayOutput {
	return o
}

func (o AccessPolicyArrayOutput) Index(i pulumi.IntInput) AccessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessPolicy {
		return vs[0].([]*AccessPolicy)[vs[1].(int)]
	}).(AccessPolicyOutput)
}

type AccessPolicyMapOutput struct{ *pulumi.OutputState }

func (AccessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPolicy)(nil)).Elem()
}

func (o AccessPolicyMapOutput) ToAccessPolicyMapOutput() AccessPolicyMapOutput {
	return o
}

func (o AccessPolicyMapOutput) ToAccessPolicyMapOutputWithContext(ctx context.Context) AccessPolicyMapOutput {
	return o
}

func (o AccessPolicyMapOutput) MapIndex(k pulumi.StringInput) AccessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessPolicy {
		return vs[0].(map[string]*AccessPolicy)[vs[1].(string)]
	}).(AccessPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyInput)(nil)).Elem(), &AccessPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyArrayInput)(nil)).Elem(), AccessPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyMapInput)(nil)).Elem(), AccessPolicyMap{})
	pulumi.RegisterOutputType(AccessPolicyOutput{})
	pulumi.RegisterOutputType(AccessPolicyArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyMapOutput{})
}
