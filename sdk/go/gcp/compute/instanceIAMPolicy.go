// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for Compute Engine Instance. Each of these resources serves a different use case:
//
// * `compute.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `compute.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `compute.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `compute.InstanceIAMPolicy`: Retrieves the IAM policy for the instance
//
// > **Note:** `compute.InstanceIAMPolicy` **cannot** be used in conjunction with `compute.InstanceIAMBinding` and `compute.InstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.InstanceIAMBinding` resources **can be** used in conjunction with `compute.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_compute\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
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
//						Role: "roles/compute.osLogin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstanceIAMPolicy(ctx, "policy", &compute.InstanceIAMPolicyArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				PolicyData:   *pulumi.String(admin.PolicyData),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
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
//						Role: "roles/compute.osLogin",
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
//			_, err = compute.NewInstanceIAMPolicy(ctx, "policy", &compute.InstanceIAMPolicyArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				PolicyData:   *pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## google\_compute\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewInstanceIAMBinding(ctx, "binding", &compute.InstanceIAMBindingArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewInstanceIAMBinding(ctx, "binding", &compute.InstanceIAMBindingArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &compute.InstanceIAMBindingConditionArgs{
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
// ## google\_compute\_instance\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewInstanceIAMMember(ctx, "member", &compute.InstanceIAMMemberArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
//				Member:       pulumi.String("user:jane@example.com"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewInstanceIAMMember(ctx, "member", &compute.InstanceIAMMemberArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
//				Member:       pulumi.String("user:jane@example.com"),
//				Condition: &compute.InstanceIAMMemberConditionArgs{
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/zones/{{zone}}/instances/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine instance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceIAMPolicy:InstanceIAMPolicy editor "projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceIAMPolicy:InstanceIAMPolicy editor "projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceIAMPolicy:InstanceIAMPolicy editor projects/{{project}}/zones/{{zone}}/instances/{{instance}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
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
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMPolicy(ctx *pulumi.Context,
	name string, args *InstanceIAMPolicyArgs, opts ...pulumi.ResourceOption) (*InstanceIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceIAMPolicy
	err := ctx.RegisterResource("gcp:compute/instanceIAMPolicy:InstanceIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIAMPolicy gets an existing InstanceIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIAMPolicyState, opts ...pulumi.ResourceOption) (*InstanceIAMPolicy, error) {
	var resource InstanceIAMPolicy
	err := ctx.ReadResource("gcp:compute/instanceIAMPolicy:InstanceIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIAMPolicy resources.
type instanceIAMPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName *string `pulumi:"instanceName"`
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
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

type InstanceIAMPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringPtrInput
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
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (InstanceIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMPolicyState)(nil)).Elem()
}

type instanceIAMPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	InstanceName string `pulumi:"instanceName"`
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
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceIAMPolicy resource.
type InstanceIAMPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringInput
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
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (InstanceIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMPolicyArgs)(nil)).Elem()
}

type InstanceIAMPolicyInput interface {
	pulumi.Input

	ToInstanceIAMPolicyOutput() InstanceIAMPolicyOutput
	ToInstanceIAMPolicyOutputWithContext(ctx context.Context) InstanceIAMPolicyOutput
}

func (*InstanceIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMPolicy)(nil)).Elem()
}

func (i *InstanceIAMPolicy) ToInstanceIAMPolicyOutput() InstanceIAMPolicyOutput {
	return i.ToInstanceIAMPolicyOutputWithContext(context.Background())
}

func (i *InstanceIAMPolicy) ToInstanceIAMPolicyOutputWithContext(ctx context.Context) InstanceIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMPolicyOutput)
}

func (i *InstanceIAMPolicy) ToOutput(ctx context.Context) pulumix.Output[*InstanceIAMPolicy] {
	return pulumix.Output[*InstanceIAMPolicy]{
		OutputState: i.ToInstanceIAMPolicyOutputWithContext(ctx).OutputState,
	}
}

// InstanceIAMPolicyArrayInput is an input type that accepts InstanceIAMPolicyArray and InstanceIAMPolicyArrayOutput values.
// You can construct a concrete instance of `InstanceIAMPolicyArrayInput` via:
//
//	InstanceIAMPolicyArray{ InstanceIAMPolicyArgs{...} }
type InstanceIAMPolicyArrayInput interface {
	pulumi.Input

	ToInstanceIAMPolicyArrayOutput() InstanceIAMPolicyArrayOutput
	ToInstanceIAMPolicyArrayOutputWithContext(context.Context) InstanceIAMPolicyArrayOutput
}

type InstanceIAMPolicyArray []InstanceIAMPolicyInput

func (InstanceIAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIAMPolicy)(nil)).Elem()
}

func (i InstanceIAMPolicyArray) ToInstanceIAMPolicyArrayOutput() InstanceIAMPolicyArrayOutput {
	return i.ToInstanceIAMPolicyArrayOutputWithContext(context.Background())
}

func (i InstanceIAMPolicyArray) ToInstanceIAMPolicyArrayOutputWithContext(ctx context.Context) InstanceIAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMPolicyArrayOutput)
}

func (i InstanceIAMPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*InstanceIAMPolicy] {
	return pulumix.Output[[]*InstanceIAMPolicy]{
		OutputState: i.ToInstanceIAMPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// InstanceIAMPolicyMapInput is an input type that accepts InstanceIAMPolicyMap and InstanceIAMPolicyMapOutput values.
// You can construct a concrete instance of `InstanceIAMPolicyMapInput` via:
//
//	InstanceIAMPolicyMap{ "key": InstanceIAMPolicyArgs{...} }
type InstanceIAMPolicyMapInput interface {
	pulumi.Input

	ToInstanceIAMPolicyMapOutput() InstanceIAMPolicyMapOutput
	ToInstanceIAMPolicyMapOutputWithContext(context.Context) InstanceIAMPolicyMapOutput
}

type InstanceIAMPolicyMap map[string]InstanceIAMPolicyInput

func (InstanceIAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIAMPolicy)(nil)).Elem()
}

func (i InstanceIAMPolicyMap) ToInstanceIAMPolicyMapOutput() InstanceIAMPolicyMapOutput {
	return i.ToInstanceIAMPolicyMapOutputWithContext(context.Background())
}

func (i InstanceIAMPolicyMap) ToInstanceIAMPolicyMapOutputWithContext(ctx context.Context) InstanceIAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMPolicyMapOutput)
}

func (i InstanceIAMPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*InstanceIAMPolicy] {
	return pulumix.Output[map[string]*InstanceIAMPolicy]{
		OutputState: i.ToInstanceIAMPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type InstanceIAMPolicyOutput struct{ *pulumi.OutputState }

func (InstanceIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMPolicy)(nil)).Elem()
}

func (o InstanceIAMPolicyOutput) ToInstanceIAMPolicyOutput() InstanceIAMPolicyOutput {
	return o
}

func (o InstanceIAMPolicyOutput) ToInstanceIAMPolicyOutputWithContext(ctx context.Context) InstanceIAMPolicyOutput {
	return o
}

func (o InstanceIAMPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*InstanceIAMPolicy] {
	return pulumix.Output[*InstanceIAMPolicy]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o InstanceIAMPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o InstanceIAMPolicyOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMPolicy) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o InstanceIAMPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
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
func (o InstanceIAMPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
// zone is specified, it is taken from the provider configuration.
func (o InstanceIAMPolicyOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMPolicy) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceIAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (InstanceIAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIAMPolicy)(nil)).Elem()
}

func (o InstanceIAMPolicyArrayOutput) ToInstanceIAMPolicyArrayOutput() InstanceIAMPolicyArrayOutput {
	return o
}

func (o InstanceIAMPolicyArrayOutput) ToInstanceIAMPolicyArrayOutputWithContext(ctx context.Context) InstanceIAMPolicyArrayOutput {
	return o
}

func (o InstanceIAMPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*InstanceIAMPolicy] {
	return pulumix.Output[[]*InstanceIAMPolicy]{
		OutputState: o.OutputState,
	}
}

func (o InstanceIAMPolicyArrayOutput) Index(i pulumi.IntInput) InstanceIAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceIAMPolicy {
		return vs[0].([]*InstanceIAMPolicy)[vs[1].(int)]
	}).(InstanceIAMPolicyOutput)
}

type InstanceIAMPolicyMapOutput struct{ *pulumi.OutputState }

func (InstanceIAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIAMPolicy)(nil)).Elem()
}

func (o InstanceIAMPolicyMapOutput) ToInstanceIAMPolicyMapOutput() InstanceIAMPolicyMapOutput {
	return o
}

func (o InstanceIAMPolicyMapOutput) ToInstanceIAMPolicyMapOutputWithContext(ctx context.Context) InstanceIAMPolicyMapOutput {
	return o
}

func (o InstanceIAMPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*InstanceIAMPolicy] {
	return pulumix.Output[map[string]*InstanceIAMPolicy]{
		OutputState: o.OutputState,
	}
}

func (o InstanceIAMPolicyMapOutput) MapIndex(k pulumi.StringInput) InstanceIAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceIAMPolicy {
		return vs[0].(map[string]*InstanceIAMPolicy)[vs[1].(string)]
	}).(InstanceIAMPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIAMPolicyInput)(nil)).Elem(), &InstanceIAMPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIAMPolicyArrayInput)(nil)).Elem(), InstanceIAMPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIAMPolicyMapInput)(nil)).Elem(), InstanceIAMPolicyMap{})
	pulumi.RegisterOutputType(InstanceIAMPolicyOutput{})
	pulumi.RegisterOutputType(InstanceIAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(InstanceIAMPolicyMapOutput{})
}
