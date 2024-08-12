// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Pub/Sub Schema. Each of these resources serves a different use case:
//
// * `pubsub.SchemaIamPolicy`: Authoritative. Sets the IAM policy for the schema and replaces any existing policy already attached.
// * `pubsub.SchemaIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the schema are preserved.
// * `pubsub.SchemaIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the schema are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `pubsub.SchemaIamPolicy`: Retrieves the IAM policy for the schema
//
// > **Note:** `pubsub.SchemaIamPolicy` **cannot** be used in conjunction with `pubsub.SchemaIamBinding` and `pubsub.SchemaIamMember` or they will fight over what your policy should be.
//
// > **Note:** `pubsub.SchemaIamBinding` resources **can be** used in conjunction with `pubsub.SchemaIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## pubsub.SchemaIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
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
//			_, err = pubsub.NewSchemaIamPolicy(ctx, "policy", &pubsub.SchemaIamPolicyArgs{
//				Project:    pulumi.Any(example.Project),
//				Schema:     pulumi.Any(example.Name),
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## pubsub.SchemaIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pubsub.NewSchemaIamBinding(ctx, "binding", &pubsub.SchemaIamBindingArgs{
//				Project: pulumi.Any(example.Project),
//				Schema:  pulumi.Any(example.Name),
//				Role:    pulumi.String("roles/viewer"),
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
// ## pubsub.SchemaIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pubsub.NewSchemaIamMember(ctx, "member", &pubsub.SchemaIamMemberArgs{
//				Project: pulumi.Any(example.Project),
//				Schema:  pulumi.Any(example.Name),
//				Role:    pulumi.String("roles/viewer"),
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
// ## This resource supports User Project Overrides.
//
// -
//
// # IAM policy for Cloud Pub/Sub Schema
// Three different resources help you manage your IAM policy for Cloud Pub/Sub Schema. Each of these resources serves a different use case:
//
// * `pubsub.SchemaIamPolicy`: Authoritative. Sets the IAM policy for the schema and replaces any existing policy already attached.
// * `pubsub.SchemaIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the schema are preserved.
// * `pubsub.SchemaIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the schema are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `pubsub.SchemaIamPolicy`: Retrieves the IAM policy for the schema
//
// > **Note:** `pubsub.SchemaIamPolicy` **cannot** be used in conjunction with `pubsub.SchemaIamBinding` and `pubsub.SchemaIamMember` or they will fight over what your policy should be.
//
// > **Note:** `pubsub.SchemaIamBinding` resources **can be** used in conjunction with `pubsub.SchemaIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## pubsub.SchemaIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
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
//			_, err = pubsub.NewSchemaIamPolicy(ctx, "policy", &pubsub.SchemaIamPolicyArgs{
//				Project:    pulumi.Any(example.Project),
//				Schema:     pulumi.Any(example.Name),
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## pubsub.SchemaIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pubsub.NewSchemaIamBinding(ctx, "binding", &pubsub.SchemaIamBindingArgs{
//				Project: pulumi.Any(example.Project),
//				Schema:  pulumi.Any(example.Name),
//				Role:    pulumi.String("roles/viewer"),
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
// ## pubsub.SchemaIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pubsub.NewSchemaIamMember(ctx, "member", &pubsub.SchemaIamMemberArgs{
//				Project: pulumi.Any(example.Project),
//				Schema:  pulumi.Any(example.Name),
//				Role:    pulumi.String("roles/viewer"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/schemas/{{name}}
//
// * {{project}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Cloud Pub/Sub schema IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:pubsub/schemaIamPolicy:SchemaIamPolicy editor "projects/{{project}}/schemas/{{schema}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:pubsub/schemaIamPolicy:SchemaIamPolicy editor "projects/{{project}}/schemas/{{schema}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:pubsub/schemaIamPolicy:SchemaIamPolicy editor projects/{{project}}/schemas/{{schema}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type SchemaIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Schema pulumi.StringOutput `pulumi:"schema"`
}

// NewSchemaIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewSchemaIamPolicy(ctx *pulumi.Context,
	name string, args *SchemaIamPolicyArgs, opts ...pulumi.ResourceOption) (*SchemaIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SchemaIamPolicy
	err := ctx.RegisterResource("gcp:pubsub/schemaIamPolicy:SchemaIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchemaIamPolicy gets an existing SchemaIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchemaIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaIamPolicyState, opts ...pulumi.ResourceOption) (*SchemaIamPolicy, error) {
	var resource SchemaIamPolicy
	err := ctx.ReadResource("gcp:pubsub/schemaIamPolicy:SchemaIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SchemaIamPolicy resources.
type schemaIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Schema *string `pulumi:"schema"`
}

type SchemaIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Schema pulumi.StringPtrInput
}

func (SchemaIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaIamPolicyState)(nil)).Elem()
}

type schemaIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Schema string `pulumi:"schema"`
}

// The set of arguments for constructing a SchemaIamPolicy resource.
type SchemaIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Schema pulumi.StringInput
}

func (SchemaIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaIamPolicyArgs)(nil)).Elem()
}

type SchemaIamPolicyInput interface {
	pulumi.Input

	ToSchemaIamPolicyOutput() SchemaIamPolicyOutput
	ToSchemaIamPolicyOutputWithContext(ctx context.Context) SchemaIamPolicyOutput
}

func (*SchemaIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaIamPolicy)(nil)).Elem()
}

func (i *SchemaIamPolicy) ToSchemaIamPolicyOutput() SchemaIamPolicyOutput {
	return i.ToSchemaIamPolicyOutputWithContext(context.Background())
}

func (i *SchemaIamPolicy) ToSchemaIamPolicyOutputWithContext(ctx context.Context) SchemaIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaIamPolicyOutput)
}

// SchemaIamPolicyArrayInput is an input type that accepts SchemaIamPolicyArray and SchemaIamPolicyArrayOutput values.
// You can construct a concrete instance of `SchemaIamPolicyArrayInput` via:
//
//	SchemaIamPolicyArray{ SchemaIamPolicyArgs{...} }
type SchemaIamPolicyArrayInput interface {
	pulumi.Input

	ToSchemaIamPolicyArrayOutput() SchemaIamPolicyArrayOutput
	ToSchemaIamPolicyArrayOutputWithContext(context.Context) SchemaIamPolicyArrayOutput
}

type SchemaIamPolicyArray []SchemaIamPolicyInput

func (SchemaIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SchemaIamPolicy)(nil)).Elem()
}

func (i SchemaIamPolicyArray) ToSchemaIamPolicyArrayOutput() SchemaIamPolicyArrayOutput {
	return i.ToSchemaIamPolicyArrayOutputWithContext(context.Background())
}

func (i SchemaIamPolicyArray) ToSchemaIamPolicyArrayOutputWithContext(ctx context.Context) SchemaIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaIamPolicyArrayOutput)
}

// SchemaIamPolicyMapInput is an input type that accepts SchemaIamPolicyMap and SchemaIamPolicyMapOutput values.
// You can construct a concrete instance of `SchemaIamPolicyMapInput` via:
//
//	SchemaIamPolicyMap{ "key": SchemaIamPolicyArgs{...} }
type SchemaIamPolicyMapInput interface {
	pulumi.Input

	ToSchemaIamPolicyMapOutput() SchemaIamPolicyMapOutput
	ToSchemaIamPolicyMapOutputWithContext(context.Context) SchemaIamPolicyMapOutput
}

type SchemaIamPolicyMap map[string]SchemaIamPolicyInput

func (SchemaIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SchemaIamPolicy)(nil)).Elem()
}

func (i SchemaIamPolicyMap) ToSchemaIamPolicyMapOutput() SchemaIamPolicyMapOutput {
	return i.ToSchemaIamPolicyMapOutputWithContext(context.Background())
}

func (i SchemaIamPolicyMap) ToSchemaIamPolicyMapOutputWithContext(ctx context.Context) SchemaIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaIamPolicyMapOutput)
}

type SchemaIamPolicyOutput struct{ *pulumi.OutputState }

func (SchemaIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaIamPolicy)(nil)).Elem()
}

func (o SchemaIamPolicyOutput) ToSchemaIamPolicyOutput() SchemaIamPolicyOutput {
	return o
}

func (o SchemaIamPolicyOutput) ToSchemaIamPolicyOutputWithContext(ctx context.Context) SchemaIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o SchemaIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o SchemaIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o SchemaIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o SchemaIamPolicyOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaIamPolicy) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

type SchemaIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (SchemaIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SchemaIamPolicy)(nil)).Elem()
}

func (o SchemaIamPolicyArrayOutput) ToSchemaIamPolicyArrayOutput() SchemaIamPolicyArrayOutput {
	return o
}

func (o SchemaIamPolicyArrayOutput) ToSchemaIamPolicyArrayOutputWithContext(ctx context.Context) SchemaIamPolicyArrayOutput {
	return o
}

func (o SchemaIamPolicyArrayOutput) Index(i pulumi.IntInput) SchemaIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SchemaIamPolicy {
		return vs[0].([]*SchemaIamPolicy)[vs[1].(int)]
	}).(SchemaIamPolicyOutput)
}

type SchemaIamPolicyMapOutput struct{ *pulumi.OutputState }

func (SchemaIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SchemaIamPolicy)(nil)).Elem()
}

func (o SchemaIamPolicyMapOutput) ToSchemaIamPolicyMapOutput() SchemaIamPolicyMapOutput {
	return o
}

func (o SchemaIamPolicyMapOutput) ToSchemaIamPolicyMapOutputWithContext(ctx context.Context) SchemaIamPolicyMapOutput {
	return o
}

func (o SchemaIamPolicyMapOutput) MapIndex(k pulumi.StringInput) SchemaIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SchemaIamPolicy {
		return vs[0].(map[string]*SchemaIamPolicy)[vs[1].(string)]
	}).(SchemaIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaIamPolicyInput)(nil)).Elem(), &SchemaIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaIamPolicyArrayInput)(nil)).Elem(), SchemaIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaIamPolicyMapInput)(nil)).Elem(), SchemaIamPolicyMap{})
	pulumi.RegisterOutputType(SchemaIamPolicyOutput{})
	pulumi.RegisterOutputType(SchemaIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(SchemaIamPolicyMapOutput{})
}
