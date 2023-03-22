// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud DNS ManagedZone. Each of these resources serves a different use case:
//
// * `dns.DnsManagedZoneIamPolicy`: Authoritative. Sets the IAM policy for the managedzone and replaces any existing policy already attached.
// * `dns.DnsManagedZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedzone are preserved.
// * `dns.DnsManagedZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedzone are preserved.
//
// > **Note:** `dns.DnsManagedZoneIamPolicy` **cannot** be used in conjunction with `dns.DnsManagedZoneIamBinding` and `dns.DnsManagedZoneIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dns.DnsManagedZoneIamBinding` resources **can be** used in conjunction with `dns.DnsManagedZoneIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_dns\_managed\_zone\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
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
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewDnsManagedZoneIamPolicy(ctx, "policy", &dns.DnsManagedZoneIamPolicyArgs{
//				Project:     pulumi.Any(google_dns_managed_zone.Default.Project),
//				ManagedZone: pulumi.Any(google_dns_managed_zone.Default.Name),
//				PolicyData:  *pulumi.String(admin.PolicyData),
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
// ## google\_dns\_managed\_zone\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dns.NewDnsManagedZoneIamBinding(ctx, "binding", &dns.DnsManagedZoneIamBindingArgs{
//				Project:     pulumi.Any(google_dns_managed_zone.Default.Project),
//				ManagedZone: pulumi.Any(google_dns_managed_zone.Default.Name),
//				Role:        pulumi.String("roles/viewer"),
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
// ## google\_dns\_managed\_zone\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dns.NewDnsManagedZoneIamMember(ctx, "member", &dns.DnsManagedZoneIamMemberArgs{
//				Project:     pulumi.Any(google_dns_managed_zone.Default.Project),
//				ManagedZone: pulumi.Any(google_dns_managed_zone.Default.Name),
//				Role:        pulumi.String("roles/viewer"),
//				Member:      pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/managedZones/{{managed_zone}} * {{project}}/{{managed_zone}} * {{managed_zone}} Any variables not passed in the import command will be taken from the provider configuration. Cloud DNS managedzone IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dns/dnsManagedZoneIamMember:DnsManagedZoneIamMember editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dns/dnsManagedZoneIamMember:DnsManagedZoneIamMember editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dns/dnsManagedZoneIamMember:DnsManagedZoneIamMember editor projects/{{project}}/managedZones/{{managed_zone}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DnsManagedZoneIamMember struct {
	pulumi.CustomResourceState

	Condition DnsManagedZoneIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone pulumi.StringOutput `pulumi:"managedZone"`
	Member      pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDnsManagedZoneIamMember registers a new resource with the given unique name, arguments, and options.
func NewDnsManagedZoneIamMember(ctx *pulumi.Context,
	name string, args *DnsManagedZoneIamMemberArgs, opts ...pulumi.ResourceOption) (*DnsManagedZoneIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedZone == nil {
		return nil, errors.New("invalid value for required argument 'ManagedZone'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DnsManagedZoneIamMember
	err := ctx.RegisterResource("gcp:dns/dnsManagedZoneIamMember:DnsManagedZoneIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsManagedZoneIamMember gets an existing DnsManagedZoneIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsManagedZoneIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsManagedZoneIamMemberState, opts ...pulumi.ResourceOption) (*DnsManagedZoneIamMember, error) {
	var resource DnsManagedZoneIamMember
	err := ctx.ReadResource("gcp:dns/dnsManagedZoneIamMember:DnsManagedZoneIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsManagedZoneIamMember resources.
type dnsManagedZoneIamMemberState struct {
	Condition *DnsManagedZoneIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone *string `pulumi:"managedZone"`
	Member      *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DnsManagedZoneIamMemberState struct {
	Condition DnsManagedZoneIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone pulumi.StringPtrInput
	Member      pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DnsManagedZoneIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsManagedZoneIamMemberState)(nil)).Elem()
}

type dnsManagedZoneIamMemberArgs struct {
	Condition *DnsManagedZoneIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone string `pulumi:"managedZone"`
	Member      string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DnsManagedZoneIamMember resource.
type DnsManagedZoneIamMemberArgs struct {
	Condition DnsManagedZoneIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone pulumi.StringInput
	Member      pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DnsManagedZoneIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsManagedZoneIamMemberArgs)(nil)).Elem()
}

type DnsManagedZoneIamMemberInput interface {
	pulumi.Input

	ToDnsManagedZoneIamMemberOutput() DnsManagedZoneIamMemberOutput
	ToDnsManagedZoneIamMemberOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberOutput
}

func (*DnsManagedZoneIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsManagedZoneIamMember)(nil)).Elem()
}

func (i *DnsManagedZoneIamMember) ToDnsManagedZoneIamMemberOutput() DnsManagedZoneIamMemberOutput {
	return i.ToDnsManagedZoneIamMemberOutputWithContext(context.Background())
}

func (i *DnsManagedZoneIamMember) ToDnsManagedZoneIamMemberOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsManagedZoneIamMemberOutput)
}

// DnsManagedZoneIamMemberArrayInput is an input type that accepts DnsManagedZoneIamMemberArray and DnsManagedZoneIamMemberArrayOutput values.
// You can construct a concrete instance of `DnsManagedZoneIamMemberArrayInput` via:
//
//	DnsManagedZoneIamMemberArray{ DnsManagedZoneIamMemberArgs{...} }
type DnsManagedZoneIamMemberArrayInput interface {
	pulumi.Input

	ToDnsManagedZoneIamMemberArrayOutput() DnsManagedZoneIamMemberArrayOutput
	ToDnsManagedZoneIamMemberArrayOutputWithContext(context.Context) DnsManagedZoneIamMemberArrayOutput
}

type DnsManagedZoneIamMemberArray []DnsManagedZoneIamMemberInput

func (DnsManagedZoneIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsManagedZoneIamMember)(nil)).Elem()
}

func (i DnsManagedZoneIamMemberArray) ToDnsManagedZoneIamMemberArrayOutput() DnsManagedZoneIamMemberArrayOutput {
	return i.ToDnsManagedZoneIamMemberArrayOutputWithContext(context.Background())
}

func (i DnsManagedZoneIamMemberArray) ToDnsManagedZoneIamMemberArrayOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsManagedZoneIamMemberArrayOutput)
}

// DnsManagedZoneIamMemberMapInput is an input type that accepts DnsManagedZoneIamMemberMap and DnsManagedZoneIamMemberMapOutput values.
// You can construct a concrete instance of `DnsManagedZoneIamMemberMapInput` via:
//
//	DnsManagedZoneIamMemberMap{ "key": DnsManagedZoneIamMemberArgs{...} }
type DnsManagedZoneIamMemberMapInput interface {
	pulumi.Input

	ToDnsManagedZoneIamMemberMapOutput() DnsManagedZoneIamMemberMapOutput
	ToDnsManagedZoneIamMemberMapOutputWithContext(context.Context) DnsManagedZoneIamMemberMapOutput
}

type DnsManagedZoneIamMemberMap map[string]DnsManagedZoneIamMemberInput

func (DnsManagedZoneIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsManagedZoneIamMember)(nil)).Elem()
}

func (i DnsManagedZoneIamMemberMap) ToDnsManagedZoneIamMemberMapOutput() DnsManagedZoneIamMemberMapOutput {
	return i.ToDnsManagedZoneIamMemberMapOutputWithContext(context.Background())
}

func (i DnsManagedZoneIamMemberMap) ToDnsManagedZoneIamMemberMapOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsManagedZoneIamMemberMapOutput)
}

type DnsManagedZoneIamMemberOutput struct{ *pulumi.OutputState }

func (DnsManagedZoneIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsManagedZoneIamMember)(nil)).Elem()
}

func (o DnsManagedZoneIamMemberOutput) ToDnsManagedZoneIamMemberOutput() DnsManagedZoneIamMemberOutput {
	return o
}

func (o DnsManagedZoneIamMemberOutput) ToDnsManagedZoneIamMemberOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberOutput {
	return o
}

func (o DnsManagedZoneIamMemberOutput) Condition() DnsManagedZoneIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) DnsManagedZoneIamMemberConditionPtrOutput { return v.Condition }).(DnsManagedZoneIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o DnsManagedZoneIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o DnsManagedZoneIamMemberOutput) ManagedZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.ManagedZone }).(pulumi.StringOutput)
}

func (o DnsManagedZoneIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o DnsManagedZoneIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o DnsManagedZoneIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type DnsManagedZoneIamMemberArrayOutput struct{ *pulumi.OutputState }

func (DnsManagedZoneIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsManagedZoneIamMember)(nil)).Elem()
}

func (o DnsManagedZoneIamMemberArrayOutput) ToDnsManagedZoneIamMemberArrayOutput() DnsManagedZoneIamMemberArrayOutput {
	return o
}

func (o DnsManagedZoneIamMemberArrayOutput) ToDnsManagedZoneIamMemberArrayOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberArrayOutput {
	return o
}

func (o DnsManagedZoneIamMemberArrayOutput) Index(i pulumi.IntInput) DnsManagedZoneIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsManagedZoneIamMember {
		return vs[0].([]*DnsManagedZoneIamMember)[vs[1].(int)]
	}).(DnsManagedZoneIamMemberOutput)
}

type DnsManagedZoneIamMemberMapOutput struct{ *pulumi.OutputState }

func (DnsManagedZoneIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsManagedZoneIamMember)(nil)).Elem()
}

func (o DnsManagedZoneIamMemberMapOutput) ToDnsManagedZoneIamMemberMapOutput() DnsManagedZoneIamMemberMapOutput {
	return o
}

func (o DnsManagedZoneIamMemberMapOutput) ToDnsManagedZoneIamMemberMapOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberMapOutput {
	return o
}

func (o DnsManagedZoneIamMemberMapOutput) MapIndex(k pulumi.StringInput) DnsManagedZoneIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsManagedZoneIamMember {
		return vs[0].(map[string]*DnsManagedZoneIamMember)[vs[1].(string)]
	}).(DnsManagedZoneIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsManagedZoneIamMemberInput)(nil)).Elem(), &DnsManagedZoneIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsManagedZoneIamMemberArrayInput)(nil)).Elem(), DnsManagedZoneIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsManagedZoneIamMemberMapInput)(nil)).Elem(), DnsManagedZoneIamMemberMap{})
	pulumi.RegisterOutputType(DnsManagedZoneIamMemberOutput{})
	pulumi.RegisterOutputType(DnsManagedZoneIamMemberArrayOutput{})
	pulumi.RegisterOutputType(DnsManagedZoneIamMemberMapOutput{})
}
