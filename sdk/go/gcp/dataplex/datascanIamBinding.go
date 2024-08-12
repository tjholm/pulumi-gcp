// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataplex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Dataplex Datascan. Each of these resources serves a different use case:
//
// * `dataplex.DatascanIamPolicy`: Authoritative. Sets the IAM policy for the datascan and replaces any existing policy already attached.
// * `dataplex.DatascanIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the datascan are preserved.
// * `dataplex.DatascanIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the datascan are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.DatascanIamPolicy`: Retrieves the IAM policy for the datascan
//
// > **Note:** `dataplex.DatascanIamPolicy` **cannot** be used in conjunction with `dataplex.DatascanIamBinding` and `dataplex.DatascanIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.DatascanIamBinding` resources **can be** used in conjunction with `dataplex.DatascanIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.DatascanIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
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
//			_, err = dataplex.NewDatascanIamPolicy(ctx, "policy", &dataplex.DatascanIamPolicyArgs{
//				Project:    pulumi.Any(basicProfile.Project),
//				Location:   pulumi.Any(basicProfile.Location),
//				DataScanId: pulumi.Any(basicProfile.DataScanId),
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
// ## dataplex.DatascanIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewDatascanIamBinding(ctx, "binding", &dataplex.DatascanIamBindingArgs{
//				Project:    pulumi.Any(basicProfile.Project),
//				Location:   pulumi.Any(basicProfile.Location),
//				DataScanId: pulumi.Any(basicProfile.DataScanId),
//				Role:       pulumi.String("roles/viewer"),
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
// ## dataplex.DatascanIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewDatascanIamMember(ctx, "member", &dataplex.DatascanIamMemberArgs{
//				Project:    pulumi.Any(basicProfile.Project),
//				Location:   pulumi.Any(basicProfile.Location),
//				DataScanId: pulumi.Any(basicProfile.DataScanId),
//				Role:       pulumi.String("roles/viewer"),
//				Member:     pulumi.String("user:jane@example.com"),
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
// # IAM policy for Dataplex Datascan
// Three different resources help you manage your IAM policy for Dataplex Datascan. Each of these resources serves a different use case:
//
// * `dataplex.DatascanIamPolicy`: Authoritative. Sets the IAM policy for the datascan and replaces any existing policy already attached.
// * `dataplex.DatascanIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the datascan are preserved.
// * `dataplex.DatascanIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the datascan are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.DatascanIamPolicy`: Retrieves the IAM policy for the datascan
//
// > **Note:** `dataplex.DatascanIamPolicy` **cannot** be used in conjunction with `dataplex.DatascanIamBinding` and `dataplex.DatascanIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.DatascanIamBinding` resources **can be** used in conjunction with `dataplex.DatascanIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.DatascanIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
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
//			_, err = dataplex.NewDatascanIamPolicy(ctx, "policy", &dataplex.DatascanIamPolicyArgs{
//				Project:    pulumi.Any(basicProfile.Project),
//				Location:   pulumi.Any(basicProfile.Location),
//				DataScanId: pulumi.Any(basicProfile.DataScanId),
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
// ## dataplex.DatascanIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewDatascanIamBinding(ctx, "binding", &dataplex.DatascanIamBindingArgs{
//				Project:    pulumi.Any(basicProfile.Project),
//				Location:   pulumi.Any(basicProfile.Location),
//				DataScanId: pulumi.Any(basicProfile.DataScanId),
//				Role:       pulumi.String("roles/viewer"),
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
// ## dataplex.DatascanIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewDatascanIamMember(ctx, "member", &dataplex.DatascanIamMemberArgs{
//				Project:    pulumi.Any(basicProfile.Project),
//				Location:   pulumi.Any(basicProfile.Location),
//				DataScanId: pulumi.Any(basicProfile.DataScanId),
//				Role:       pulumi.String("roles/viewer"),
//				Member:     pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}
//
// * {{project}}/{{location}}/{{data_scan_id}}
//
// * {{location}}/{{data_scan_id}}
//
// * {{data_scan_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Dataplex datascan IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/datascanIamBinding:DatascanIamBinding editor "projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/datascanIamBinding:DatascanIamBinding editor "projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/datascanIamBinding:DatascanIamBinding editor projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatascanIamBinding struct {
	pulumi.CustomResourceState

	Condition  DatascanIamBindingConditionPtrOutput `pulumi:"condition"`
	DataScanId pulumi.StringOutput                  `pulumi:"dataScanId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location where the data scan should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataplex.DatascanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatascanIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDatascanIamBinding(ctx *pulumi.Context,
	name string, args *DatascanIamBindingArgs, opts ...pulumi.ResourceOption) (*DatascanIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataScanId == nil {
		return nil, errors.New("invalid value for required argument 'DataScanId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatascanIamBinding
	err := ctx.RegisterResource("gcp:dataplex/datascanIamBinding:DatascanIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatascanIamBinding gets an existing DatascanIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatascanIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatascanIamBindingState, opts ...pulumi.ResourceOption) (*DatascanIamBinding, error) {
	var resource DatascanIamBinding
	err := ctx.ReadResource("gcp:dataplex/datascanIamBinding:DatascanIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatascanIamBinding resources.
type datascanIamBindingState struct {
	Condition  *DatascanIamBindingCondition `pulumi:"condition"`
	DataScanId *string                      `pulumi:"dataScanId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location where the data scan should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataplex.DatascanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatascanIamBindingState struct {
	Condition  DatascanIamBindingConditionPtrInput
	DataScanId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location where the data scan should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataplex.DatascanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatascanIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*datascanIamBindingState)(nil)).Elem()
}

type datascanIamBindingArgs struct {
	Condition  *DatascanIamBindingCondition `pulumi:"condition"`
	DataScanId string                       `pulumi:"dataScanId"`
	// The location where the data scan should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataplex.DatascanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatascanIamBinding resource.
type DatascanIamBindingArgs struct {
	Condition  DatascanIamBindingConditionPtrInput
	DataScanId pulumi.StringInput
	// The location where the data scan should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataplex.DatascanIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatascanIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datascanIamBindingArgs)(nil)).Elem()
}

type DatascanIamBindingInput interface {
	pulumi.Input

	ToDatascanIamBindingOutput() DatascanIamBindingOutput
	ToDatascanIamBindingOutputWithContext(ctx context.Context) DatascanIamBindingOutput
}

func (*DatascanIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DatascanIamBinding)(nil)).Elem()
}

func (i *DatascanIamBinding) ToDatascanIamBindingOutput() DatascanIamBindingOutput {
	return i.ToDatascanIamBindingOutputWithContext(context.Background())
}

func (i *DatascanIamBinding) ToDatascanIamBindingOutputWithContext(ctx context.Context) DatascanIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatascanIamBindingOutput)
}

// DatascanIamBindingArrayInput is an input type that accepts DatascanIamBindingArray and DatascanIamBindingArrayOutput values.
// You can construct a concrete instance of `DatascanIamBindingArrayInput` via:
//
//	DatascanIamBindingArray{ DatascanIamBindingArgs{...} }
type DatascanIamBindingArrayInput interface {
	pulumi.Input

	ToDatascanIamBindingArrayOutput() DatascanIamBindingArrayOutput
	ToDatascanIamBindingArrayOutputWithContext(context.Context) DatascanIamBindingArrayOutput
}

type DatascanIamBindingArray []DatascanIamBindingInput

func (DatascanIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatascanIamBinding)(nil)).Elem()
}

func (i DatascanIamBindingArray) ToDatascanIamBindingArrayOutput() DatascanIamBindingArrayOutput {
	return i.ToDatascanIamBindingArrayOutputWithContext(context.Background())
}

func (i DatascanIamBindingArray) ToDatascanIamBindingArrayOutputWithContext(ctx context.Context) DatascanIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatascanIamBindingArrayOutput)
}

// DatascanIamBindingMapInput is an input type that accepts DatascanIamBindingMap and DatascanIamBindingMapOutput values.
// You can construct a concrete instance of `DatascanIamBindingMapInput` via:
//
//	DatascanIamBindingMap{ "key": DatascanIamBindingArgs{...} }
type DatascanIamBindingMapInput interface {
	pulumi.Input

	ToDatascanIamBindingMapOutput() DatascanIamBindingMapOutput
	ToDatascanIamBindingMapOutputWithContext(context.Context) DatascanIamBindingMapOutput
}

type DatascanIamBindingMap map[string]DatascanIamBindingInput

func (DatascanIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatascanIamBinding)(nil)).Elem()
}

func (i DatascanIamBindingMap) ToDatascanIamBindingMapOutput() DatascanIamBindingMapOutput {
	return i.ToDatascanIamBindingMapOutputWithContext(context.Background())
}

func (i DatascanIamBindingMap) ToDatascanIamBindingMapOutputWithContext(ctx context.Context) DatascanIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatascanIamBindingMapOutput)
}

type DatascanIamBindingOutput struct{ *pulumi.OutputState }

func (DatascanIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatascanIamBinding)(nil)).Elem()
}

func (o DatascanIamBindingOutput) ToDatascanIamBindingOutput() DatascanIamBindingOutput {
	return o
}

func (o DatascanIamBindingOutput) ToDatascanIamBindingOutputWithContext(ctx context.Context) DatascanIamBindingOutput {
	return o
}

func (o DatascanIamBindingOutput) Condition() DatascanIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DatascanIamBinding) DatascanIamBindingConditionPtrOutput { return v.Condition }).(DatascanIamBindingConditionPtrOutput)
}

func (o DatascanIamBindingOutput) DataScanId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatascanIamBinding) pulumi.StringOutput { return v.DataScanId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o DatascanIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DatascanIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location where the data scan should reside.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o DatascanIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DatascanIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o DatascanIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatascanIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o DatascanIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatascanIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataplex.DatascanIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o DatascanIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DatascanIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type DatascanIamBindingArrayOutput struct{ *pulumi.OutputState }

func (DatascanIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatascanIamBinding)(nil)).Elem()
}

func (o DatascanIamBindingArrayOutput) ToDatascanIamBindingArrayOutput() DatascanIamBindingArrayOutput {
	return o
}

func (o DatascanIamBindingArrayOutput) ToDatascanIamBindingArrayOutputWithContext(ctx context.Context) DatascanIamBindingArrayOutput {
	return o
}

func (o DatascanIamBindingArrayOutput) Index(i pulumi.IntInput) DatascanIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatascanIamBinding {
		return vs[0].([]*DatascanIamBinding)[vs[1].(int)]
	}).(DatascanIamBindingOutput)
}

type DatascanIamBindingMapOutput struct{ *pulumi.OutputState }

func (DatascanIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatascanIamBinding)(nil)).Elem()
}

func (o DatascanIamBindingMapOutput) ToDatascanIamBindingMapOutput() DatascanIamBindingMapOutput {
	return o
}

func (o DatascanIamBindingMapOutput) ToDatascanIamBindingMapOutputWithContext(ctx context.Context) DatascanIamBindingMapOutput {
	return o
}

func (o DatascanIamBindingMapOutput) MapIndex(k pulumi.StringInput) DatascanIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatascanIamBinding {
		return vs[0].(map[string]*DatascanIamBinding)[vs[1].(string)]
	}).(DatascanIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatascanIamBindingInput)(nil)).Elem(), &DatascanIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatascanIamBindingArrayInput)(nil)).Elem(), DatascanIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatascanIamBindingMapInput)(nil)).Elem(), DatascanIamBindingMap{})
	pulumi.RegisterOutputType(DatascanIamBindingOutput{})
	pulumi.RegisterOutputType(DatascanIamBindingArrayOutput{})
	pulumi.RegisterOutputType(DatascanIamBindingMapOutput{})
}
