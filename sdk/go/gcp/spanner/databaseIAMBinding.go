// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
//
// * `spanner.DatabaseIAMPolicy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
//
// > **Warning:** It's entirely possibly to lock yourself out of your database using `spanner.DatabaseIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
//
// * `spanner.DatabaseIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
// * `spanner.DatabaseIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
//
// > **Note:** `spanner.DatabaseIAMPolicy` **cannot** be used in conjunction with `spanner.DatabaseIAMBinding` and `spanner.DatabaseIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `spanner.DatabaseIAMBinding` resources **can be** used in conjunction with `spanner.DatabaseIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_spanner\_database\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
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
//			_, err = spanner.NewDatabaseIAMPolicy(ctx, "database", &spanner.DatabaseIAMPolicyArgs{
//				Instance:   pulumi.String("your-instance-name"),
//				Database:   pulumi.String("your-database-name"),
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
// ## google\_spanner\_database\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := spanner.NewDatabaseIAMBinding(ctx, "database", &spanner.DatabaseIAMBindingArgs{
//				Database: pulumi.String("your-database-name"),
//				Instance: pulumi.String("your-instance-name"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role: pulumi.String("roles/compute.networkUser"),
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
// ## google\_spanner\_database\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := spanner.NewDatabaseIAMMember(ctx, "database", &spanner.DatabaseIAMMemberArgs{
//				Database: pulumi.String("your-database-name"),
//				Instance: pulumi.String("your-instance-name"),
//				Member:   pulumi.String("user:jane@example.com"),
//				Role:     pulumi.String("roles/compute.networkUser"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{instance}}/{{database}} * {{instance}}/{{database}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:spanner/databaseIAMBinding:DatabaseIAMBinding database "project-name/instance-name/database-name roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:spanner/databaseIAMBinding:DatabaseIAMBinding database "project-name/instance-name/database-name roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:spanner/databaseIAMBinding:DatabaseIAMBinding database project-name/instance-name/database-name
//
// ```
//
//	-> **Custom Roles:** If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatabaseIAMBinding struct {
	pulumi.CustomResourceState

	Condition DatabaseIAMBindingConditionPtrOutput `pulumi:"condition"`
	// The name of the Spanner database.
	Database pulumi.StringOutput `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatabaseIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewDatabaseIAMBinding(ctx *pulumi.Context,
	name string, args *DatabaseIAMBindingArgs, opts ...pulumi.ResourceOption) (*DatabaseIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseIAMBinding
	err := ctx.RegisterResource("gcp:spanner/databaseIAMBinding:DatabaseIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseIAMBinding gets an existing DatabaseIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseIAMBindingState, opts ...pulumi.ResourceOption) (*DatabaseIAMBinding, error) {
	var resource DatabaseIAMBinding
	err := ctx.ReadResource("gcp:spanner/databaseIAMBinding:DatabaseIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseIAMBinding resources.
type databaseIAMBindingState struct {
	Condition *DatabaseIAMBindingCondition `pulumi:"condition"`
	// The name of the Spanner database.
	Database *string `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatabaseIAMBindingState struct {
	Condition DatabaseIAMBindingConditionPtrInput
	// The name of the Spanner database.
	Database pulumi.StringPtrInput
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the Spanner instance the database belongs to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatabaseIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMBindingState)(nil)).Elem()
}

type databaseIAMBindingArgs struct {
	Condition *DatabaseIAMBindingCondition `pulumi:"condition"`
	// The name of the Spanner database.
	Database string `pulumi:"database"`
	// The name of the Spanner instance the database belongs to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatabaseIAMBinding resource.
type DatabaseIAMBindingArgs struct {
	Condition DatabaseIAMBindingConditionPtrInput
	// The name of the Spanner database.
	Database pulumi.StringInput
	// The name of the Spanner instance the database belongs to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatabaseIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMBindingArgs)(nil)).Elem()
}

type DatabaseIAMBindingInput interface {
	pulumi.Input

	ToDatabaseIAMBindingOutput() DatabaseIAMBindingOutput
	ToDatabaseIAMBindingOutputWithContext(ctx context.Context) DatabaseIAMBindingOutput
}

func (*DatabaseIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIAMBinding)(nil)).Elem()
}

func (i *DatabaseIAMBinding) ToDatabaseIAMBindingOutput() DatabaseIAMBindingOutput {
	return i.ToDatabaseIAMBindingOutputWithContext(context.Background())
}

func (i *DatabaseIAMBinding) ToDatabaseIAMBindingOutputWithContext(ctx context.Context) DatabaseIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMBindingOutput)
}

func (i *DatabaseIAMBinding) ToOutput(ctx context.Context) pulumix.Output[*DatabaseIAMBinding] {
	return pulumix.Output[*DatabaseIAMBinding]{
		OutputState: i.ToDatabaseIAMBindingOutputWithContext(ctx).OutputState,
	}
}

// DatabaseIAMBindingArrayInput is an input type that accepts DatabaseIAMBindingArray and DatabaseIAMBindingArrayOutput values.
// You can construct a concrete instance of `DatabaseIAMBindingArrayInput` via:
//
//	DatabaseIAMBindingArray{ DatabaseIAMBindingArgs{...} }
type DatabaseIAMBindingArrayInput interface {
	pulumi.Input

	ToDatabaseIAMBindingArrayOutput() DatabaseIAMBindingArrayOutput
	ToDatabaseIAMBindingArrayOutputWithContext(context.Context) DatabaseIAMBindingArrayOutput
}

type DatabaseIAMBindingArray []DatabaseIAMBindingInput

func (DatabaseIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseIAMBinding)(nil)).Elem()
}

func (i DatabaseIAMBindingArray) ToDatabaseIAMBindingArrayOutput() DatabaseIAMBindingArrayOutput {
	return i.ToDatabaseIAMBindingArrayOutputWithContext(context.Background())
}

func (i DatabaseIAMBindingArray) ToDatabaseIAMBindingArrayOutputWithContext(ctx context.Context) DatabaseIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMBindingArrayOutput)
}

func (i DatabaseIAMBindingArray) ToOutput(ctx context.Context) pulumix.Output[[]*DatabaseIAMBinding] {
	return pulumix.Output[[]*DatabaseIAMBinding]{
		OutputState: i.ToDatabaseIAMBindingArrayOutputWithContext(ctx).OutputState,
	}
}

// DatabaseIAMBindingMapInput is an input type that accepts DatabaseIAMBindingMap and DatabaseIAMBindingMapOutput values.
// You can construct a concrete instance of `DatabaseIAMBindingMapInput` via:
//
//	DatabaseIAMBindingMap{ "key": DatabaseIAMBindingArgs{...} }
type DatabaseIAMBindingMapInput interface {
	pulumi.Input

	ToDatabaseIAMBindingMapOutput() DatabaseIAMBindingMapOutput
	ToDatabaseIAMBindingMapOutputWithContext(context.Context) DatabaseIAMBindingMapOutput
}

type DatabaseIAMBindingMap map[string]DatabaseIAMBindingInput

func (DatabaseIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseIAMBinding)(nil)).Elem()
}

func (i DatabaseIAMBindingMap) ToDatabaseIAMBindingMapOutput() DatabaseIAMBindingMapOutput {
	return i.ToDatabaseIAMBindingMapOutputWithContext(context.Background())
}

func (i DatabaseIAMBindingMap) ToDatabaseIAMBindingMapOutputWithContext(ctx context.Context) DatabaseIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMBindingMapOutput)
}

func (i DatabaseIAMBindingMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DatabaseIAMBinding] {
	return pulumix.Output[map[string]*DatabaseIAMBinding]{
		OutputState: i.ToDatabaseIAMBindingMapOutputWithContext(ctx).OutputState,
	}
}

type DatabaseIAMBindingOutput struct{ *pulumi.OutputState }

func (DatabaseIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIAMBinding)(nil)).Elem()
}

func (o DatabaseIAMBindingOutput) ToDatabaseIAMBindingOutput() DatabaseIAMBindingOutput {
	return o
}

func (o DatabaseIAMBindingOutput) ToDatabaseIAMBindingOutputWithContext(ctx context.Context) DatabaseIAMBindingOutput {
	return o
}

func (o DatabaseIAMBindingOutput) ToOutput(ctx context.Context) pulumix.Output[*DatabaseIAMBinding] {
	return pulumix.Output[*DatabaseIAMBinding]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseIAMBindingOutput) Condition() DatabaseIAMBindingConditionPtrOutput {
	return o.ApplyT(func(v *DatabaseIAMBinding) DatabaseIAMBindingConditionPtrOutput { return v.Condition }).(DatabaseIAMBindingConditionPtrOutput)
}

// The name of the Spanner database.
func (o DatabaseIAMBindingOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseIAMBinding) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// (Computed) The etag of the database's IAM policy.
func (o DatabaseIAMBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseIAMBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the Spanner instance the database belongs to.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o DatabaseIAMBindingOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseIAMBinding) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

func (o DatabaseIAMBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseIAMBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o DatabaseIAMBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseIAMBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o DatabaseIAMBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseIAMBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type DatabaseIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (DatabaseIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseIAMBinding)(nil)).Elem()
}

func (o DatabaseIAMBindingArrayOutput) ToDatabaseIAMBindingArrayOutput() DatabaseIAMBindingArrayOutput {
	return o
}

func (o DatabaseIAMBindingArrayOutput) ToDatabaseIAMBindingArrayOutputWithContext(ctx context.Context) DatabaseIAMBindingArrayOutput {
	return o
}

func (o DatabaseIAMBindingArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DatabaseIAMBinding] {
	return pulumix.Output[[]*DatabaseIAMBinding]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseIAMBindingArrayOutput) Index(i pulumi.IntInput) DatabaseIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseIAMBinding {
		return vs[0].([]*DatabaseIAMBinding)[vs[1].(int)]
	}).(DatabaseIAMBindingOutput)
}

type DatabaseIAMBindingMapOutput struct{ *pulumi.OutputState }

func (DatabaseIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseIAMBinding)(nil)).Elem()
}

func (o DatabaseIAMBindingMapOutput) ToDatabaseIAMBindingMapOutput() DatabaseIAMBindingMapOutput {
	return o
}

func (o DatabaseIAMBindingMapOutput) ToDatabaseIAMBindingMapOutputWithContext(ctx context.Context) DatabaseIAMBindingMapOutput {
	return o
}

func (o DatabaseIAMBindingMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DatabaseIAMBinding] {
	return pulumix.Output[map[string]*DatabaseIAMBinding]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseIAMBindingMapOutput) MapIndex(k pulumi.StringInput) DatabaseIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseIAMBinding {
		return vs[0].(map[string]*DatabaseIAMBinding)[vs[1].(string)]
	}).(DatabaseIAMBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseIAMBindingInput)(nil)).Elem(), &DatabaseIAMBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseIAMBindingArrayInput)(nil)).Elem(), DatabaseIAMBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseIAMBindingMapInput)(nil)).Elem(), DatabaseIAMBindingMap{})
	pulumi.RegisterOutputType(DatabaseIAMBindingOutput{})
	pulumi.RegisterOutputType(DatabaseIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(DatabaseIAMBindingMapOutput{})
}
