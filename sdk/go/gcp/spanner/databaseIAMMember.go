// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = spanner.NewDatabaseIAMPolicy(ctx, "database", &spanner.DatabaseIAMPolicyArgs{
// 			Instance:   pulumi.String("your-instance-name"),
// 			Database:   pulumi.String("your-database-name"),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_spanner\_database\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.NewDatabaseIAMBinding(ctx, "database", &spanner.DatabaseIAMBindingArgs{
// 			Database: pulumi.String("your-database-name"),
// 			Instance: pulumi.String("your-instance-name"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/compute.networkUser"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_spanner\_database\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.NewDatabaseIAMMember(ctx, "database", &spanner.DatabaseIAMMemberArgs{
// 			Database: pulumi.String("your-database-name"),
// 			Instance: pulumi.String("your-instance-name"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Role:     pulumi.String("roles/compute.networkUser"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{instance}}/{{database}} * {{instance}}/{{database}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/databaseIAMMember:DatabaseIAMMember database "project-name/instance-name/database-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/databaseIAMMember:DatabaseIAMMember database "project-name/instance-name/database-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/databaseIAMMember:DatabaseIAMMember database project-name/instance-name/database-name
// ```
//
//  -> **Custom Roles:** If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatabaseIAMMember struct {
	pulumi.CustomResourceState

	Condition DatabaseIAMMemberConditionPtrOutput `pulumi:"condition"`
	// The name of the Spanner database.
	Database pulumi.StringOutput `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringOutput `pulumi:"instance"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatabaseIAMMember registers a new resource with the given unique name, arguments, and options.
func NewDatabaseIAMMember(ctx *pulumi.Context,
	name string, args *DatabaseIAMMemberArgs, opts ...pulumi.ResourceOption) (*DatabaseIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DatabaseIAMMember
	err := ctx.RegisterResource("gcp:spanner/databaseIAMMember:DatabaseIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseIAMMember gets an existing DatabaseIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseIAMMemberState, opts ...pulumi.ResourceOption) (*DatabaseIAMMember, error) {
	var resource DatabaseIAMMember
	err := ctx.ReadResource("gcp:spanner/databaseIAMMember:DatabaseIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseIAMMember resources.
type databaseIAMMemberState struct {
	Condition *DatabaseIAMMemberCondition `pulumi:"condition"`
	// The name of the Spanner database.
	Database *string `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance *string `pulumi:"instance"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatabaseIAMMemberState struct {
	Condition DatabaseIAMMemberConditionPtrInput
	// The name of the Spanner database.
	Database pulumi.StringPtrInput
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatabaseIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMMemberState)(nil)).Elem()
}

type databaseIAMMemberArgs struct {
	Condition *DatabaseIAMMemberCondition `pulumi:"condition"`
	// The name of the Spanner database.
	Database string `pulumi:"database"`
	// The name of the Spanner instance the database belongs to.
	Instance string `pulumi:"instance"`
	Member   string `pulumi:"member"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatabaseIAMMember resource.
type DatabaseIAMMemberArgs struct {
	Condition DatabaseIAMMemberConditionPtrInput
	// The name of the Spanner database.
	Database pulumi.StringInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatabaseIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMMemberArgs)(nil)).Elem()
}

type DatabaseIAMMemberInput interface {
	pulumi.Input

	ToDatabaseIAMMemberOutput() DatabaseIAMMemberOutput
	ToDatabaseIAMMemberOutputWithContext(ctx context.Context) DatabaseIAMMemberOutput
}

func (*DatabaseIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIAMMember)(nil)).Elem()
}

func (i *DatabaseIAMMember) ToDatabaseIAMMemberOutput() DatabaseIAMMemberOutput {
	return i.ToDatabaseIAMMemberOutputWithContext(context.Background())
}

func (i *DatabaseIAMMember) ToDatabaseIAMMemberOutputWithContext(ctx context.Context) DatabaseIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMMemberOutput)
}

// DatabaseIAMMemberArrayInput is an input type that accepts DatabaseIAMMemberArray and DatabaseIAMMemberArrayOutput values.
// You can construct a concrete instance of `DatabaseIAMMemberArrayInput` via:
//
//          DatabaseIAMMemberArray{ DatabaseIAMMemberArgs{...} }
type DatabaseIAMMemberArrayInput interface {
	pulumi.Input

	ToDatabaseIAMMemberArrayOutput() DatabaseIAMMemberArrayOutput
	ToDatabaseIAMMemberArrayOutputWithContext(context.Context) DatabaseIAMMemberArrayOutput
}

type DatabaseIAMMemberArray []DatabaseIAMMemberInput

func (DatabaseIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseIAMMember)(nil)).Elem()
}

func (i DatabaseIAMMemberArray) ToDatabaseIAMMemberArrayOutput() DatabaseIAMMemberArrayOutput {
	return i.ToDatabaseIAMMemberArrayOutputWithContext(context.Background())
}

func (i DatabaseIAMMemberArray) ToDatabaseIAMMemberArrayOutputWithContext(ctx context.Context) DatabaseIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMMemberArrayOutput)
}

// DatabaseIAMMemberMapInput is an input type that accepts DatabaseIAMMemberMap and DatabaseIAMMemberMapOutput values.
// You can construct a concrete instance of `DatabaseIAMMemberMapInput` via:
//
//          DatabaseIAMMemberMap{ "key": DatabaseIAMMemberArgs{...} }
type DatabaseIAMMemberMapInput interface {
	pulumi.Input

	ToDatabaseIAMMemberMapOutput() DatabaseIAMMemberMapOutput
	ToDatabaseIAMMemberMapOutputWithContext(context.Context) DatabaseIAMMemberMapOutput
}

type DatabaseIAMMemberMap map[string]DatabaseIAMMemberInput

func (DatabaseIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseIAMMember)(nil)).Elem()
}

func (i DatabaseIAMMemberMap) ToDatabaseIAMMemberMapOutput() DatabaseIAMMemberMapOutput {
	return i.ToDatabaseIAMMemberMapOutputWithContext(context.Background())
}

func (i DatabaseIAMMemberMap) ToDatabaseIAMMemberMapOutputWithContext(ctx context.Context) DatabaseIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMMemberMapOutput)
}

type DatabaseIAMMemberOutput struct{ *pulumi.OutputState }

func (DatabaseIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIAMMember)(nil)).Elem()
}

func (o DatabaseIAMMemberOutput) ToDatabaseIAMMemberOutput() DatabaseIAMMemberOutput {
	return o
}

func (o DatabaseIAMMemberOutput) ToDatabaseIAMMemberOutputWithContext(ctx context.Context) DatabaseIAMMemberOutput {
	return o
}

type DatabaseIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (DatabaseIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseIAMMember)(nil)).Elem()
}

func (o DatabaseIAMMemberArrayOutput) ToDatabaseIAMMemberArrayOutput() DatabaseIAMMemberArrayOutput {
	return o
}

func (o DatabaseIAMMemberArrayOutput) ToDatabaseIAMMemberArrayOutputWithContext(ctx context.Context) DatabaseIAMMemberArrayOutput {
	return o
}

func (o DatabaseIAMMemberArrayOutput) Index(i pulumi.IntInput) DatabaseIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseIAMMember {
		return vs[0].([]*DatabaseIAMMember)[vs[1].(int)]
	}).(DatabaseIAMMemberOutput)
}

type DatabaseIAMMemberMapOutput struct{ *pulumi.OutputState }

func (DatabaseIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseIAMMember)(nil)).Elem()
}

func (o DatabaseIAMMemberMapOutput) ToDatabaseIAMMemberMapOutput() DatabaseIAMMemberMapOutput {
	return o
}

func (o DatabaseIAMMemberMapOutput) ToDatabaseIAMMemberMapOutputWithContext(ctx context.Context) DatabaseIAMMemberMapOutput {
	return o
}

func (o DatabaseIAMMemberMapOutput) MapIndex(k pulumi.StringInput) DatabaseIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseIAMMember {
		return vs[0].(map[string]*DatabaseIAMMember)[vs[1].(string)]
	}).(DatabaseIAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseIAMMemberInput)(nil)).Elem(), &DatabaseIAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseIAMMemberArrayInput)(nil)).Elem(), DatabaseIAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseIAMMemberMapInput)(nil)).Elem(), DatabaseIAMMemberMap{})
	pulumi.RegisterOutputType(DatabaseIAMMemberOutput{})
	pulumi.RegisterOutputType(DatabaseIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(DatabaseIAMMemberMapOutput{})
}
