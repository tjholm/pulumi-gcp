// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Dataproc metastore Service. Each of these resources serves a different use case:
//
// * `dataproc.MetastoreServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `dataproc.MetastoreServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `dataproc.MetastoreServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataproc.MetastoreServiceIamPolicy`: Retrieves the IAM policy for the service
//
// > **Note:** `dataproc.MetastoreServiceIamPolicy` **cannot** be used in conjunction with `dataproc.MetastoreServiceIamBinding` and `dataproc.MetastoreServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataproc.MetastoreServiceIamBinding` resources **can be** used in conjunction with `dataproc.MetastoreServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_dataproc\_metastore\_service\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
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
//			_, err = dataproc.NewMetastoreServiceIamPolicy(ctx, "policy", &dataproc.MetastoreServiceIamPolicyArgs{
//				Project:    pulumi.Any(_default.Project),
//				Location:   pulumi.Any(_default.Location),
//				ServiceId:  pulumi.Any(_default.ServiceId),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_metastore\_service\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewMetastoreServiceIamBinding(ctx, "binding", &dataproc.MetastoreServiceIamBindingArgs{
//				Project:   pulumi.Any(_default.Project),
//				Location:  pulumi.Any(_default.Location),
//				ServiceId: pulumi.Any(_default.ServiceId),
//				Role:      pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_metastore\_service\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewMetastoreServiceIamMember(ctx, "member", &dataproc.MetastoreServiceIamMemberArgs{
//				Project:   pulumi.Any(_default.Project),
//				Location:  pulumi.Any(_default.Location),
//				ServiceId: pulumi.Any(_default.ServiceId),
//				Role:      pulumi.String("roles/viewer"),
//				Member:    pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_metastore\_service\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
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
//			_, err = dataproc.NewMetastoreServiceIamPolicy(ctx, "policy", &dataproc.MetastoreServiceIamPolicyArgs{
//				Project:    pulumi.Any(_default.Project),
//				Location:   pulumi.Any(_default.Location),
//				ServiceId:  pulumi.Any(_default.ServiceId),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_metastore\_service\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewMetastoreServiceIamBinding(ctx, "binding", &dataproc.MetastoreServiceIamBindingArgs{
//				Project:   pulumi.Any(_default.Project),
//				Location:  pulumi.Any(_default.Location),
//				ServiceId: pulumi.Any(_default.ServiceId),
//				Role:      pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_metastore\_service\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewMetastoreServiceIamMember(ctx, "member", &dataproc.MetastoreServiceIamMemberArgs{
//				Project:   pulumi.Any(_default.Project),
//				Location:  pulumi.Any(_default.Location),
//				ServiceId: pulumi.Any(_default.ServiceId),
//				Role:      pulumi.String("roles/viewer"),
//				Member:    pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/{{location}}/services/{{service_id}}
//
// * {{project}}/{{location}}/{{service_id}}
//
// * {{location}}/{{service_id}}
//
// * {{service_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Dataproc metastore service IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:dataproc/metastoreServiceIamMember:MetastoreServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:dataproc/metastoreServiceIamMember:MetastoreServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:dataproc/metastoreServiceIamMember:MetastoreServiceIamMember editor projects/{{project}}/locations/{{location}}/services/{{service_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MetastoreServiceIamMember struct {
	pulumi.CustomResourceState

	Condition MetastoreServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location where the metastore service should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      pulumi.StringOutput `pulumi:"role"`
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
}

// NewMetastoreServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewMetastoreServiceIamMember(ctx *pulumi.Context,
	name string, args *MetastoreServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*MetastoreServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MetastoreServiceIamMember
	err := ctx.RegisterResource("gcp:dataproc/metastoreServiceIamMember:MetastoreServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetastoreServiceIamMember gets an existing MetastoreServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetastoreServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetastoreServiceIamMemberState, opts ...pulumi.ResourceOption) (*MetastoreServiceIamMember, error) {
	var resource MetastoreServiceIamMember
	err := ctx.ReadResource("gcp:dataproc/metastoreServiceIamMember:MetastoreServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetastoreServiceIamMember resources.
type metastoreServiceIamMemberState struct {
	Condition *MetastoreServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location where the metastore service should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      *string `pulumi:"role"`
	ServiceId *string `pulumi:"serviceId"`
}

type MetastoreServiceIamMemberState struct {
	Condition MetastoreServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location where the metastore service should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      pulumi.StringPtrInput
	ServiceId pulumi.StringPtrInput
}

func (MetastoreServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreServiceIamMemberState)(nil)).Elem()
}

type metastoreServiceIamMemberArgs struct {
	Condition *MetastoreServiceIamMemberCondition `pulumi:"condition"`
	// The location where the metastore service should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      string `pulumi:"role"`
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a MetastoreServiceIamMember resource.
type MetastoreServiceIamMemberArgs struct {
	Condition MetastoreServiceIamMemberConditionPtrInput
	// The location where the metastore service should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      pulumi.StringInput
	ServiceId pulumi.StringInput
}

func (MetastoreServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreServiceIamMemberArgs)(nil)).Elem()
}

type MetastoreServiceIamMemberInput interface {
	pulumi.Input

	ToMetastoreServiceIamMemberOutput() MetastoreServiceIamMemberOutput
	ToMetastoreServiceIamMemberOutputWithContext(ctx context.Context) MetastoreServiceIamMemberOutput
}

func (*MetastoreServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreServiceIamMember)(nil)).Elem()
}

func (i *MetastoreServiceIamMember) ToMetastoreServiceIamMemberOutput() MetastoreServiceIamMemberOutput {
	return i.ToMetastoreServiceIamMemberOutputWithContext(context.Background())
}

func (i *MetastoreServiceIamMember) ToMetastoreServiceIamMemberOutputWithContext(ctx context.Context) MetastoreServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreServiceIamMemberOutput)
}

// MetastoreServiceIamMemberArrayInput is an input type that accepts MetastoreServiceIamMemberArray and MetastoreServiceIamMemberArrayOutput values.
// You can construct a concrete instance of `MetastoreServiceIamMemberArrayInput` via:
//
//	MetastoreServiceIamMemberArray{ MetastoreServiceIamMemberArgs{...} }
type MetastoreServiceIamMemberArrayInput interface {
	pulumi.Input

	ToMetastoreServiceIamMemberArrayOutput() MetastoreServiceIamMemberArrayOutput
	ToMetastoreServiceIamMemberArrayOutputWithContext(context.Context) MetastoreServiceIamMemberArrayOutput
}

type MetastoreServiceIamMemberArray []MetastoreServiceIamMemberInput

func (MetastoreServiceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreServiceIamMember)(nil)).Elem()
}

func (i MetastoreServiceIamMemberArray) ToMetastoreServiceIamMemberArrayOutput() MetastoreServiceIamMemberArrayOutput {
	return i.ToMetastoreServiceIamMemberArrayOutputWithContext(context.Background())
}

func (i MetastoreServiceIamMemberArray) ToMetastoreServiceIamMemberArrayOutputWithContext(ctx context.Context) MetastoreServiceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreServiceIamMemberArrayOutput)
}

// MetastoreServiceIamMemberMapInput is an input type that accepts MetastoreServiceIamMemberMap and MetastoreServiceIamMemberMapOutput values.
// You can construct a concrete instance of `MetastoreServiceIamMemberMapInput` via:
//
//	MetastoreServiceIamMemberMap{ "key": MetastoreServiceIamMemberArgs{...} }
type MetastoreServiceIamMemberMapInput interface {
	pulumi.Input

	ToMetastoreServiceIamMemberMapOutput() MetastoreServiceIamMemberMapOutput
	ToMetastoreServiceIamMemberMapOutputWithContext(context.Context) MetastoreServiceIamMemberMapOutput
}

type MetastoreServiceIamMemberMap map[string]MetastoreServiceIamMemberInput

func (MetastoreServiceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreServiceIamMember)(nil)).Elem()
}

func (i MetastoreServiceIamMemberMap) ToMetastoreServiceIamMemberMapOutput() MetastoreServiceIamMemberMapOutput {
	return i.ToMetastoreServiceIamMemberMapOutputWithContext(context.Background())
}

func (i MetastoreServiceIamMemberMap) ToMetastoreServiceIamMemberMapOutputWithContext(ctx context.Context) MetastoreServiceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreServiceIamMemberMapOutput)
}

type MetastoreServiceIamMemberOutput struct{ *pulumi.OutputState }

func (MetastoreServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreServiceIamMember)(nil)).Elem()
}

func (o MetastoreServiceIamMemberOutput) ToMetastoreServiceIamMemberOutput() MetastoreServiceIamMemberOutput {
	return o
}

func (o MetastoreServiceIamMemberOutput) ToMetastoreServiceIamMemberOutputWithContext(ctx context.Context) MetastoreServiceIamMemberOutput {
	return o
}

func (o MetastoreServiceIamMemberOutput) Condition() MetastoreServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *MetastoreServiceIamMember) MetastoreServiceIamMemberConditionPtrOutput { return v.Condition }).(MetastoreServiceIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o MetastoreServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location where the metastore service should reside.
// The default value is `global`.
// Used to find the parent resource to bind the IAM policy to
func (o MetastoreServiceIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in `role`.
// Each entry can have one of the following values:
// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o MetastoreServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o MetastoreServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o MetastoreServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o MetastoreServiceIamMemberOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamMember) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

type MetastoreServiceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (MetastoreServiceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreServiceIamMember)(nil)).Elem()
}

func (o MetastoreServiceIamMemberArrayOutput) ToMetastoreServiceIamMemberArrayOutput() MetastoreServiceIamMemberArrayOutput {
	return o
}

func (o MetastoreServiceIamMemberArrayOutput) ToMetastoreServiceIamMemberArrayOutputWithContext(ctx context.Context) MetastoreServiceIamMemberArrayOutput {
	return o
}

func (o MetastoreServiceIamMemberArrayOutput) Index(i pulumi.IntInput) MetastoreServiceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetastoreServiceIamMember {
		return vs[0].([]*MetastoreServiceIamMember)[vs[1].(int)]
	}).(MetastoreServiceIamMemberOutput)
}

type MetastoreServiceIamMemberMapOutput struct{ *pulumi.OutputState }

func (MetastoreServiceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreServiceIamMember)(nil)).Elem()
}

func (o MetastoreServiceIamMemberMapOutput) ToMetastoreServiceIamMemberMapOutput() MetastoreServiceIamMemberMapOutput {
	return o
}

func (o MetastoreServiceIamMemberMapOutput) ToMetastoreServiceIamMemberMapOutputWithContext(ctx context.Context) MetastoreServiceIamMemberMapOutput {
	return o
}

func (o MetastoreServiceIamMemberMapOutput) MapIndex(k pulumi.StringInput) MetastoreServiceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetastoreServiceIamMember {
		return vs[0].(map[string]*MetastoreServiceIamMember)[vs[1].(string)]
	}).(MetastoreServiceIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreServiceIamMemberInput)(nil)).Elem(), &MetastoreServiceIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreServiceIamMemberArrayInput)(nil)).Elem(), MetastoreServiceIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreServiceIamMemberMapInput)(nil)).Elem(), MetastoreServiceIamMemberMap{})
	pulumi.RegisterOutputType(MetastoreServiceIamMemberOutput{})
	pulumi.RegisterOutputType(MetastoreServiceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(MetastoreServiceIamMemberMapOutput{})
}
