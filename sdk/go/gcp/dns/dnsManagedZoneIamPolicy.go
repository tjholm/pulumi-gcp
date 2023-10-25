// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for Cloud DNS ManagedZone. Each of these resources serves a different use case:
//
// * `dns.DnsManagedZoneIamPolicy`: Authoritative. Sets the IAM policy for the managedzone and replaces any existing policy already attached.
// * `dns.DnsManagedZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedzone are preserved.
// * `dns.DnsManagedZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedzone are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dns.DnsManagedZoneIamPolicy`: Retrieves the IAM policy for the managedzone
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
//	$ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor projects/{{project}}/managedZones/{{managed_zone}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DnsManagedZoneIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone pulumi.StringOutput `pulumi:"managedZone"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDnsManagedZoneIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDnsManagedZoneIamPolicy(ctx *pulumi.Context,
	name string, args *DnsManagedZoneIamPolicyArgs, opts ...pulumi.ResourceOption) (*DnsManagedZoneIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedZone == nil {
		return nil, errors.New("invalid value for required argument 'ManagedZone'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnsManagedZoneIamPolicy
	err := ctx.RegisterResource("gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsManagedZoneIamPolicy gets an existing DnsManagedZoneIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsManagedZoneIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsManagedZoneIamPolicyState, opts ...pulumi.ResourceOption) (*DnsManagedZoneIamPolicy, error) {
	var resource DnsManagedZoneIamPolicy
	err := ctx.ReadResource("gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsManagedZoneIamPolicy resources.
type dnsManagedZoneIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone *string `pulumi:"managedZone"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
}

type DnsManagedZoneIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
}

func (DnsManagedZoneIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsManagedZoneIamPolicyState)(nil)).Elem()
}

type dnsManagedZoneIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone string `pulumi:"managedZone"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DnsManagedZoneIamPolicy resource.
type DnsManagedZoneIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	ManagedZone pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
}

func (DnsManagedZoneIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsManagedZoneIamPolicyArgs)(nil)).Elem()
}

type DnsManagedZoneIamPolicyInput interface {
	pulumi.Input

	ToDnsManagedZoneIamPolicyOutput() DnsManagedZoneIamPolicyOutput
	ToDnsManagedZoneIamPolicyOutputWithContext(ctx context.Context) DnsManagedZoneIamPolicyOutput
}

func (*DnsManagedZoneIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsManagedZoneIamPolicy)(nil)).Elem()
}

func (i *DnsManagedZoneIamPolicy) ToDnsManagedZoneIamPolicyOutput() DnsManagedZoneIamPolicyOutput {
	return i.ToDnsManagedZoneIamPolicyOutputWithContext(context.Background())
}

func (i *DnsManagedZoneIamPolicy) ToDnsManagedZoneIamPolicyOutputWithContext(ctx context.Context) DnsManagedZoneIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsManagedZoneIamPolicyOutput)
}

func (i *DnsManagedZoneIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*DnsManagedZoneIamPolicy] {
	return pulumix.Output[*DnsManagedZoneIamPolicy]{
		OutputState: i.ToDnsManagedZoneIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// DnsManagedZoneIamPolicyArrayInput is an input type that accepts DnsManagedZoneIamPolicyArray and DnsManagedZoneIamPolicyArrayOutput values.
// You can construct a concrete instance of `DnsManagedZoneIamPolicyArrayInput` via:
//
//	DnsManagedZoneIamPolicyArray{ DnsManagedZoneIamPolicyArgs{...} }
type DnsManagedZoneIamPolicyArrayInput interface {
	pulumi.Input

	ToDnsManagedZoneIamPolicyArrayOutput() DnsManagedZoneIamPolicyArrayOutput
	ToDnsManagedZoneIamPolicyArrayOutputWithContext(context.Context) DnsManagedZoneIamPolicyArrayOutput
}

type DnsManagedZoneIamPolicyArray []DnsManagedZoneIamPolicyInput

func (DnsManagedZoneIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsManagedZoneIamPolicy)(nil)).Elem()
}

func (i DnsManagedZoneIamPolicyArray) ToDnsManagedZoneIamPolicyArrayOutput() DnsManagedZoneIamPolicyArrayOutput {
	return i.ToDnsManagedZoneIamPolicyArrayOutputWithContext(context.Background())
}

func (i DnsManagedZoneIamPolicyArray) ToDnsManagedZoneIamPolicyArrayOutputWithContext(ctx context.Context) DnsManagedZoneIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsManagedZoneIamPolicyArrayOutput)
}

func (i DnsManagedZoneIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*DnsManagedZoneIamPolicy] {
	return pulumix.Output[[]*DnsManagedZoneIamPolicy]{
		OutputState: i.ToDnsManagedZoneIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// DnsManagedZoneIamPolicyMapInput is an input type that accepts DnsManagedZoneIamPolicyMap and DnsManagedZoneIamPolicyMapOutput values.
// You can construct a concrete instance of `DnsManagedZoneIamPolicyMapInput` via:
//
//	DnsManagedZoneIamPolicyMap{ "key": DnsManagedZoneIamPolicyArgs{...} }
type DnsManagedZoneIamPolicyMapInput interface {
	pulumi.Input

	ToDnsManagedZoneIamPolicyMapOutput() DnsManagedZoneIamPolicyMapOutput
	ToDnsManagedZoneIamPolicyMapOutputWithContext(context.Context) DnsManagedZoneIamPolicyMapOutput
}

type DnsManagedZoneIamPolicyMap map[string]DnsManagedZoneIamPolicyInput

func (DnsManagedZoneIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsManagedZoneIamPolicy)(nil)).Elem()
}

func (i DnsManagedZoneIamPolicyMap) ToDnsManagedZoneIamPolicyMapOutput() DnsManagedZoneIamPolicyMapOutput {
	return i.ToDnsManagedZoneIamPolicyMapOutputWithContext(context.Background())
}

func (i DnsManagedZoneIamPolicyMap) ToDnsManagedZoneIamPolicyMapOutputWithContext(ctx context.Context) DnsManagedZoneIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsManagedZoneIamPolicyMapOutput)
}

func (i DnsManagedZoneIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DnsManagedZoneIamPolicy] {
	return pulumix.Output[map[string]*DnsManagedZoneIamPolicy]{
		OutputState: i.ToDnsManagedZoneIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type DnsManagedZoneIamPolicyOutput struct{ *pulumi.OutputState }

func (DnsManagedZoneIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsManagedZoneIamPolicy)(nil)).Elem()
}

func (o DnsManagedZoneIamPolicyOutput) ToDnsManagedZoneIamPolicyOutput() DnsManagedZoneIamPolicyOutput {
	return o
}

func (o DnsManagedZoneIamPolicyOutput) ToDnsManagedZoneIamPolicyOutputWithContext(ctx context.Context) DnsManagedZoneIamPolicyOutput {
	return o
}

func (o DnsManagedZoneIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*DnsManagedZoneIamPolicy] {
	return pulumix.Output[*DnsManagedZoneIamPolicy]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o DnsManagedZoneIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o DnsManagedZoneIamPolicyOutput) ManagedZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamPolicy) pulumi.StringOutput { return v.ManagedZone }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o DnsManagedZoneIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
//   - **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
//   - **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
//   - **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o DnsManagedZoneIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type DnsManagedZoneIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (DnsManagedZoneIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsManagedZoneIamPolicy)(nil)).Elem()
}

func (o DnsManagedZoneIamPolicyArrayOutput) ToDnsManagedZoneIamPolicyArrayOutput() DnsManagedZoneIamPolicyArrayOutput {
	return o
}

func (o DnsManagedZoneIamPolicyArrayOutput) ToDnsManagedZoneIamPolicyArrayOutputWithContext(ctx context.Context) DnsManagedZoneIamPolicyArrayOutput {
	return o
}

func (o DnsManagedZoneIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DnsManagedZoneIamPolicy] {
	return pulumix.Output[[]*DnsManagedZoneIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o DnsManagedZoneIamPolicyArrayOutput) Index(i pulumi.IntInput) DnsManagedZoneIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsManagedZoneIamPolicy {
		return vs[0].([]*DnsManagedZoneIamPolicy)[vs[1].(int)]
	}).(DnsManagedZoneIamPolicyOutput)
}

type DnsManagedZoneIamPolicyMapOutput struct{ *pulumi.OutputState }

func (DnsManagedZoneIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsManagedZoneIamPolicy)(nil)).Elem()
}

func (o DnsManagedZoneIamPolicyMapOutput) ToDnsManagedZoneIamPolicyMapOutput() DnsManagedZoneIamPolicyMapOutput {
	return o
}

func (o DnsManagedZoneIamPolicyMapOutput) ToDnsManagedZoneIamPolicyMapOutputWithContext(ctx context.Context) DnsManagedZoneIamPolicyMapOutput {
	return o
}

func (o DnsManagedZoneIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DnsManagedZoneIamPolicy] {
	return pulumix.Output[map[string]*DnsManagedZoneIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o DnsManagedZoneIamPolicyMapOutput) MapIndex(k pulumi.StringInput) DnsManagedZoneIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsManagedZoneIamPolicy {
		return vs[0].(map[string]*DnsManagedZoneIamPolicy)[vs[1].(string)]
	}).(DnsManagedZoneIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsManagedZoneIamPolicyInput)(nil)).Elem(), &DnsManagedZoneIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsManagedZoneIamPolicyArrayInput)(nil)).Elem(), DnsManagedZoneIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsManagedZoneIamPolicyMapInput)(nil)).Elem(), DnsManagedZoneIamPolicyMap{})
	pulumi.RegisterOutputType(DnsManagedZoneIamPolicyOutput{})
	pulumi.RegisterOutputType(DnsManagedZoneIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(DnsManagedZoneIamPolicyMapOutput{})
}
