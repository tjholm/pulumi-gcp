// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Four different resources help you manage your IAM policy for a organization. Each of these resources serves a different use case:
//
// * `organizations.IAMPolicy`: Authoritative. Sets the IAM policy for the organization and replaces any existing policy already attached.
// * `organizations.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the organization are preserved.
// * `organizations.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the organization are preserved.
// * `organizations.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
//
// > **Note:** `organizations.IAMPolicy` **cannot** be used in conjunction with `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `organizations.IAMBinding` resources **can be** used in conjunction with `organizations.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_organization\_iam\_policy
//
// !> **Warning:** New organizations have several default policies which will,
//
//	without extreme caution, be **overwritten** by use of this resource.
//	The safest alternative is to use multiple `organizations.IAMBinding`
//	resources. This resource makes it easy to remove your own access to
//	an organization, which will require a call to Google Support to have
//	fixed, and can take multiple days to resolve.
//
//	In general, this resource should only be used with organizations
//	fully managed by this provider.I f you do use this resource,
//	the best way to be sure that you are not making dangerous changes is to start
//	by **importing** your existing policy, and examining the diff very closely.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/editor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = organizations.NewIAMPolicy(ctx, "organization", &organizations.IAMPolicyArgs{
//				OrgId:      pulumi.String("your-organization-id"),
//				PolicyData: *pulumi.String(admin.PolicyData),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Condition: {
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//							Title:       "expires_after_2019_12_31",
//						},
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Role: "roles/editor",
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = organizations.NewIAMPolicy(ctx, "organization", &organizations.IAMPolicyArgs{
//				OrgId:      pulumi.String("your-organization-id"),
//				PolicyData: *pulumi.String(admin.PolicyData),
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
// ## google\_organization\_iam\_binding
//
// > **Note:** If `role` is set to `roles/owner` and you don't specify a user or service account you have access to in `members`, you can lock yourself out of your organization.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewIAMBinding(ctx, "organization", &organizations.IAMBindingArgs{
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				OrgId: pulumi.String("your-organization-id"),
//				Role:  pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewIAMBinding(ctx, "organization", &organizations.IAMBindingArgs{
//				Condition: &organizations.IAMBindingConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				OrgId: pulumi.String("your-organization-id"),
//				Role:  pulumi.String("roles/editor"),
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
// ## google\_organization\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewIAMMember(ctx, "organization", &organizations.IAMMemberArgs{
//				Member: pulumi.String("user:jane@example.com"),
//				OrgId:  pulumi.String("your-organization-id"),
//				Role:   pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewIAMMember(ctx, "organization", &organizations.IAMMemberArgs{
//				Condition: &organizations.IAMMemberConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Member: pulumi.String("user:jane@example.com"),
//				OrgId:  pulumi.String("your-organization-id"),
//				Role:   pulumi.String("roles/editor"),
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
// ## google\_organization\_iam\_audit\_config
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewIamAuditConfig(ctx, "organization", &organizations.IamAuditConfigArgs{
//				AuditLogConfigs: organizations.IamAuditConfigAuditLogConfigArray{
//					&organizations.IamAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("ADMIN_READ"),
//					},
//					&organizations.IamAuditConfigAuditLogConfigArgs{
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@hashicorp.com"),
//						},
//						LogType: pulumi.String("DATA_READ"),
//					},
//				},
//				OrgId:   pulumi.String("your-organization-id"),
//				Service: pulumi.String("allServices"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `org_id`, role, and member e.g.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMMember:IAMMember my_organization "your-orgid roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `org_id` and role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMMember:IAMMember my_organization "your-org-id roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `org_id`.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMMember:IAMMember my_organization your-org-id
//
// ```
//
//	IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMMember:IAMMember my_organization "your-organization-id foo.googleapis.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMMember:IAMMember to include the title of condition, e.g. `google_organization_iam_binding.my_organization "your-org-id roles/{{role_id}} condition-title"`
//
// ```
type IAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource IAMMember
	err := ctx.RegisterResource("gcp:organizations/iAMMember:IAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMMemberState, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	var resource IAMMember
	err := ctx.ReadResource("gcp:organizations/iAMMember:IAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMMember resources.
type iammemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId *string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role *string `pulumi:"role"`
}

type IAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	// (Computed) The etag of the organization's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role pulumi.StringPtrInput
}

func (IAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberState)(nil)).Elem()
}

type iammemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	Member    string              `pulumi:"member"`
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId pulumi.StringInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role pulumi.StringInput
}

func (IAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberArgs)(nil)).Elem()
}

type IAMMemberInput interface {
	pulumi.Input

	ToIAMMemberOutput() IAMMemberOutput
	ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput
}

func (*IAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (i *IAMMember) ToIAMMemberOutput() IAMMemberOutput {
	return i.ToIAMMemberOutputWithContext(context.Background())
}

func (i *IAMMember) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberOutput)
}

// IAMMemberArrayInput is an input type that accepts IAMMemberArray and IAMMemberArrayOutput values.
// You can construct a concrete instance of `IAMMemberArrayInput` via:
//
//	IAMMemberArray{ IAMMemberArgs{...} }
type IAMMemberArrayInput interface {
	pulumi.Input

	ToIAMMemberArrayOutput() IAMMemberArrayOutput
	ToIAMMemberArrayOutputWithContext(context.Context) IAMMemberArrayOutput
}

type IAMMemberArray []IAMMemberInput

func (IAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (i IAMMemberArray) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return i.ToIAMMemberArrayOutputWithContext(context.Background())
}

func (i IAMMemberArray) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberArrayOutput)
}

// IAMMemberMapInput is an input type that accepts IAMMemberMap and IAMMemberMapOutput values.
// You can construct a concrete instance of `IAMMemberMapInput` via:
//
//	IAMMemberMap{ "key": IAMMemberArgs{...} }
type IAMMemberMapInput interface {
	pulumi.Input

	ToIAMMemberMapOutput() IAMMemberMapOutput
	ToIAMMemberMapOutputWithContext(context.Context) IAMMemberMapOutput
}

type IAMMemberMap map[string]IAMMemberInput

func (IAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (i IAMMemberMap) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return i.ToIAMMemberMapOutputWithContext(context.Background())
}

func (i IAMMemberMap) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberMapOutput)
}

type IAMMemberOutput struct{ *pulumi.OutputState }

func (IAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (o IAMMemberOutput) ToIAMMemberOutput() IAMMemberOutput {
	return o
}

func (o IAMMemberOutput) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o IAMMemberOutput) Condition() IAMMemberConditionPtrOutput {
	return o.ApplyT(func(v *IAMMember) IAMMemberConditionPtrOutput { return v.Condition }).(IAMMemberConditionPtrOutput)
}

// (Computed) The etag of the organization's IAM policy.
func (o IAMMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IAMMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
// will not be inferred from the provider.
func (o IAMMemberOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
// `organizations/{{org_id}}/roles/{{role_id}}`.
func (o IAMMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type IAMMemberArrayOutput struct{ *pulumi.OutputState }

func (IAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) Index(i pulumi.IntInput) IAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].([]*IAMMember)[vs[1].(int)]
	}).(IAMMemberOutput)
}

type IAMMemberMapOutput struct{ *pulumi.OutputState }

func (IAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) MapIndex(k pulumi.StringInput) IAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].(map[string]*IAMMember)[vs[1].(string)]
	}).(IAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberInput)(nil)).Elem(), &IAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberArrayInput)(nil)).Elem(), IAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberMapInput)(nil)).Elem(), IAMMemberMap{})
	pulumi.RegisterOutputType(IAMMemberOutput{})
	pulumi.RegisterOutputType(IAMMemberArrayOutput{})
	pulumi.RegisterOutputType(IAMMemberMapOutput{})
}
