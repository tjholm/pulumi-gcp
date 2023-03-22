// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy Tunnel. Each of these resources serves a different use case:
//
// * `iap.TunnelIamPolicy`: Authoritative. Sets the IAM policy for the tunnel and replaces any existing policy already attached.
// * `iap.TunnelIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnel are preserved.
// * `iap.TunnelIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnel are preserved.
//
// > **Note:** `iap.TunnelIamPolicy` **cannot** be used in conjunction with `iap.TunnelIamBinding` and `iap.TunnelIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.TunnelIamBinding` resources **can be** used in conjunction with `iap.TunnelIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_iap\_tunnel\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
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
//						Role: "roles/iap.tunnelResourceAccessor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iap.NewTunnelIamPolicy(ctx, "policy", &iap.TunnelIamPolicyArgs{
//				Project:    pulumi.Any(google_project_service.Project_service.Project),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
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
//						Role: "roles/iap.tunnelResourceAccessor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iap.NewTunnelIamPolicy(ctx, "policy", &iap.TunnelIamPolicyArgs{
//				Project:    pulumi.Any(google_project_service.Project_service.Project),
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
// ## google\_iap\_tunnel\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewTunnelIamBinding(ctx, "binding", &iap.TunnelIamBindingArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
//				Role:    pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewTunnelIamBinding(ctx, "binding", &iap.TunnelIamBindingArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
//				Role:    pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &iap.TunnelIamBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_iap\_tunnel\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewTunnelIamMember(ctx, "member", &iap.TunnelIamMemberArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
//				Role:    pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Member:  pulumi.String("user:jane@example.com"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iap.NewTunnelIamMember(ctx, "member", &iap.TunnelIamMemberArgs{
//				Project: pulumi.Any(google_project_service.Project_service.Project),
//				Role:    pulumi.String("roles/iap.tunnelResourceAccessor"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Condition: &iap.TunnelIamMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_tunnel * {{project}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy tunnel IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/tunnelIamPolicy:TunnelIamPolicy editor "projects/{{project}}/iap_tunnel roles/iap.tunnelResourceAccessor user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/tunnelIamPolicy:TunnelIamPolicy editor "projects/{{project}}/iap_tunnel roles/iap.tunnelResourceAccessor"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:iap/tunnelIamPolicy:TunnelIamPolicy editor projects/{{project}}/iap_tunnel
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TunnelIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewTunnelIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTunnelIamPolicy(ctx *pulumi.Context,
	name string, args *TunnelIamPolicyArgs, opts ...pulumi.ResourceOption) (*TunnelIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource TunnelIamPolicy
	err := ctx.RegisterResource("gcp:iap/tunnelIamPolicy:TunnelIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTunnelIamPolicy gets an existing TunnelIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTunnelIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TunnelIamPolicyState, opts ...pulumi.ResourceOption) (*TunnelIamPolicy, error) {
	var resource TunnelIamPolicy
	err := ctx.ReadResource("gcp:iap/tunnelIamPolicy:TunnelIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TunnelIamPolicy resources.
type tunnelIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type TunnelIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TunnelIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelIamPolicyState)(nil)).Elem()
}

type tunnelIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a TunnelIamPolicy resource.
type TunnelIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TunnelIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelIamPolicyArgs)(nil)).Elem()
}

type TunnelIamPolicyInput interface {
	pulumi.Input

	ToTunnelIamPolicyOutput() TunnelIamPolicyOutput
	ToTunnelIamPolicyOutputWithContext(ctx context.Context) TunnelIamPolicyOutput
}

func (*TunnelIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelIamPolicy)(nil)).Elem()
}

func (i *TunnelIamPolicy) ToTunnelIamPolicyOutput() TunnelIamPolicyOutput {
	return i.ToTunnelIamPolicyOutputWithContext(context.Background())
}

func (i *TunnelIamPolicy) ToTunnelIamPolicyOutputWithContext(ctx context.Context) TunnelIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelIamPolicyOutput)
}

// TunnelIamPolicyArrayInput is an input type that accepts TunnelIamPolicyArray and TunnelIamPolicyArrayOutput values.
// You can construct a concrete instance of `TunnelIamPolicyArrayInput` via:
//
//	TunnelIamPolicyArray{ TunnelIamPolicyArgs{...} }
type TunnelIamPolicyArrayInput interface {
	pulumi.Input

	ToTunnelIamPolicyArrayOutput() TunnelIamPolicyArrayOutput
	ToTunnelIamPolicyArrayOutputWithContext(context.Context) TunnelIamPolicyArrayOutput
}

type TunnelIamPolicyArray []TunnelIamPolicyInput

func (TunnelIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TunnelIamPolicy)(nil)).Elem()
}

func (i TunnelIamPolicyArray) ToTunnelIamPolicyArrayOutput() TunnelIamPolicyArrayOutput {
	return i.ToTunnelIamPolicyArrayOutputWithContext(context.Background())
}

func (i TunnelIamPolicyArray) ToTunnelIamPolicyArrayOutputWithContext(ctx context.Context) TunnelIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelIamPolicyArrayOutput)
}

// TunnelIamPolicyMapInput is an input type that accepts TunnelIamPolicyMap and TunnelIamPolicyMapOutput values.
// You can construct a concrete instance of `TunnelIamPolicyMapInput` via:
//
//	TunnelIamPolicyMap{ "key": TunnelIamPolicyArgs{...} }
type TunnelIamPolicyMapInput interface {
	pulumi.Input

	ToTunnelIamPolicyMapOutput() TunnelIamPolicyMapOutput
	ToTunnelIamPolicyMapOutputWithContext(context.Context) TunnelIamPolicyMapOutput
}

type TunnelIamPolicyMap map[string]TunnelIamPolicyInput

func (TunnelIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TunnelIamPolicy)(nil)).Elem()
}

func (i TunnelIamPolicyMap) ToTunnelIamPolicyMapOutput() TunnelIamPolicyMapOutput {
	return i.ToTunnelIamPolicyMapOutputWithContext(context.Background())
}

func (i TunnelIamPolicyMap) ToTunnelIamPolicyMapOutputWithContext(ctx context.Context) TunnelIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelIamPolicyMapOutput)
}

type TunnelIamPolicyOutput struct{ *pulumi.OutputState }

func (TunnelIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelIamPolicy)(nil)).Elem()
}

func (o TunnelIamPolicyOutput) ToTunnelIamPolicyOutput() TunnelIamPolicyOutput {
	return o
}

func (o TunnelIamPolicyOutput) ToTunnelIamPolicyOutputWithContext(ctx context.Context) TunnelIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o TunnelIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o TunnelIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o TunnelIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type TunnelIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TunnelIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TunnelIamPolicy)(nil)).Elem()
}

func (o TunnelIamPolicyArrayOutput) ToTunnelIamPolicyArrayOutput() TunnelIamPolicyArrayOutput {
	return o
}

func (o TunnelIamPolicyArrayOutput) ToTunnelIamPolicyArrayOutputWithContext(ctx context.Context) TunnelIamPolicyArrayOutput {
	return o
}

func (o TunnelIamPolicyArrayOutput) Index(i pulumi.IntInput) TunnelIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TunnelIamPolicy {
		return vs[0].([]*TunnelIamPolicy)[vs[1].(int)]
	}).(TunnelIamPolicyOutput)
}

type TunnelIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TunnelIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TunnelIamPolicy)(nil)).Elem()
}

func (o TunnelIamPolicyMapOutput) ToTunnelIamPolicyMapOutput() TunnelIamPolicyMapOutput {
	return o
}

func (o TunnelIamPolicyMapOutput) ToTunnelIamPolicyMapOutputWithContext(ctx context.Context) TunnelIamPolicyMapOutput {
	return o
}

func (o TunnelIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TunnelIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TunnelIamPolicy {
		return vs[0].(map[string]*TunnelIamPolicy)[vs[1].(string)]
	}).(TunnelIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelIamPolicyInput)(nil)).Elem(), &TunnelIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelIamPolicyArrayInput)(nil)).Elem(), TunnelIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelIamPolicyMapInput)(nil)).Elem(), TunnelIamPolicyMap{})
	pulumi.RegisterOutputType(TunnelIamPolicyOutput{})
	pulumi.RegisterOutputType(TunnelIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TunnelIamPolicyMapOutput{})
}
