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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
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
//  $ pulumi import gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy database "project-name/instance-name/database-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy database "project-name/instance-name/database-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy database project-name/instance-name/database-name
// ```
//
//  -> **Custom Roles:** If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatabaseIAMPolicy struct {
	pulumi.CustomResourceState

	// The name of the Spanner database.
	Database pulumi.StringOutput `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDatabaseIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatabaseIAMPolicy(ctx *pulumi.Context,
	name string, args *DatabaseIAMPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource DatabaseIAMPolicy
	err := ctx.RegisterResource("gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseIAMPolicy gets an existing DatabaseIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseIAMPolicyState, opts ...pulumi.ResourceOption) (*DatabaseIAMPolicy, error) {
	var resource DatabaseIAMPolicy
	err := ctx.ReadResource("gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseIAMPolicy resources.
type databaseIAMPolicyState struct {
	// The name of the Spanner database.
	Database *string `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance *string `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type DatabaseIAMPolicyState struct {
	// The name of the Spanner database.
	Database pulumi.StringPtrInput
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatabaseIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMPolicyState)(nil)).Elem()
}

type databaseIAMPolicyArgs struct {
	// The name of the Spanner database.
	Database string `pulumi:"database"`
	// The name of the Spanner instance the database belongs to.
	Instance string `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DatabaseIAMPolicy resource.
type DatabaseIAMPolicyArgs struct {
	// The name of the Spanner database.
	Database pulumi.StringInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatabaseIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMPolicyArgs)(nil)).Elem()
}

type DatabaseIAMPolicyInput interface {
	pulumi.Input

	ToDatabaseIAMPolicyOutput() DatabaseIAMPolicyOutput
	ToDatabaseIAMPolicyOutputWithContext(ctx context.Context) DatabaseIAMPolicyOutput
}

func (*DatabaseIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIAMPolicy)(nil))
}

func (i *DatabaseIAMPolicy) ToDatabaseIAMPolicyOutput() DatabaseIAMPolicyOutput {
	return i.ToDatabaseIAMPolicyOutputWithContext(context.Background())
}

func (i *DatabaseIAMPolicy) ToDatabaseIAMPolicyOutputWithContext(ctx context.Context) DatabaseIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMPolicyOutput)
}

func (i *DatabaseIAMPolicy) ToDatabaseIAMPolicyPtrOutput() DatabaseIAMPolicyPtrOutput {
	return i.ToDatabaseIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *DatabaseIAMPolicy) ToDatabaseIAMPolicyPtrOutputWithContext(ctx context.Context) DatabaseIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMPolicyPtrOutput)
}

type DatabaseIAMPolicyPtrInput interface {
	pulumi.Input

	ToDatabaseIAMPolicyPtrOutput() DatabaseIAMPolicyPtrOutput
	ToDatabaseIAMPolicyPtrOutputWithContext(ctx context.Context) DatabaseIAMPolicyPtrOutput
}

type databaseIAMPolicyPtrType DatabaseIAMPolicyArgs

func (*databaseIAMPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIAMPolicy)(nil))
}

func (i *databaseIAMPolicyPtrType) ToDatabaseIAMPolicyPtrOutput() DatabaseIAMPolicyPtrOutput {
	return i.ToDatabaseIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *databaseIAMPolicyPtrType) ToDatabaseIAMPolicyPtrOutputWithContext(ctx context.Context) DatabaseIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMPolicyPtrOutput)
}

// DatabaseIAMPolicyArrayInput is an input type that accepts DatabaseIAMPolicyArray and DatabaseIAMPolicyArrayOutput values.
// You can construct a concrete instance of `DatabaseIAMPolicyArrayInput` via:
//
//          DatabaseIAMPolicyArray{ DatabaseIAMPolicyArgs{...} }
type DatabaseIAMPolicyArrayInput interface {
	pulumi.Input

	ToDatabaseIAMPolicyArrayOutput() DatabaseIAMPolicyArrayOutput
	ToDatabaseIAMPolicyArrayOutputWithContext(context.Context) DatabaseIAMPolicyArrayOutput
}

type DatabaseIAMPolicyArray []DatabaseIAMPolicyInput

func (DatabaseIAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DatabaseIAMPolicy)(nil))
}

func (i DatabaseIAMPolicyArray) ToDatabaseIAMPolicyArrayOutput() DatabaseIAMPolicyArrayOutput {
	return i.ToDatabaseIAMPolicyArrayOutputWithContext(context.Background())
}

func (i DatabaseIAMPolicyArray) ToDatabaseIAMPolicyArrayOutputWithContext(ctx context.Context) DatabaseIAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMPolicyArrayOutput)
}

// DatabaseIAMPolicyMapInput is an input type that accepts DatabaseIAMPolicyMap and DatabaseIAMPolicyMapOutput values.
// You can construct a concrete instance of `DatabaseIAMPolicyMapInput` via:
//
//          DatabaseIAMPolicyMap{ "key": DatabaseIAMPolicyArgs{...} }
type DatabaseIAMPolicyMapInput interface {
	pulumi.Input

	ToDatabaseIAMPolicyMapOutput() DatabaseIAMPolicyMapOutput
	ToDatabaseIAMPolicyMapOutputWithContext(context.Context) DatabaseIAMPolicyMapOutput
}

type DatabaseIAMPolicyMap map[string]DatabaseIAMPolicyInput

func (DatabaseIAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DatabaseIAMPolicy)(nil))
}

func (i DatabaseIAMPolicyMap) ToDatabaseIAMPolicyMapOutput() DatabaseIAMPolicyMapOutput {
	return i.ToDatabaseIAMPolicyMapOutputWithContext(context.Background())
}

func (i DatabaseIAMPolicyMap) ToDatabaseIAMPolicyMapOutputWithContext(ctx context.Context) DatabaseIAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMPolicyMapOutput)
}

type DatabaseIAMPolicyOutput struct {
	*pulumi.OutputState
}

func (DatabaseIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIAMPolicy)(nil))
}

func (o DatabaseIAMPolicyOutput) ToDatabaseIAMPolicyOutput() DatabaseIAMPolicyOutput {
	return o
}

func (o DatabaseIAMPolicyOutput) ToDatabaseIAMPolicyOutputWithContext(ctx context.Context) DatabaseIAMPolicyOutput {
	return o
}

func (o DatabaseIAMPolicyOutput) ToDatabaseIAMPolicyPtrOutput() DatabaseIAMPolicyPtrOutput {
	return o.ToDatabaseIAMPolicyPtrOutputWithContext(context.Background())
}

func (o DatabaseIAMPolicyOutput) ToDatabaseIAMPolicyPtrOutputWithContext(ctx context.Context) DatabaseIAMPolicyPtrOutput {
	return o.ApplyT(func(v DatabaseIAMPolicy) *DatabaseIAMPolicy {
		return &v
	}).(DatabaseIAMPolicyPtrOutput)
}

type DatabaseIAMPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (DatabaseIAMPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIAMPolicy)(nil))
}

func (o DatabaseIAMPolicyPtrOutput) ToDatabaseIAMPolicyPtrOutput() DatabaseIAMPolicyPtrOutput {
	return o
}

func (o DatabaseIAMPolicyPtrOutput) ToDatabaseIAMPolicyPtrOutputWithContext(ctx context.Context) DatabaseIAMPolicyPtrOutput {
	return o
}

type DatabaseIAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (DatabaseIAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseIAMPolicy)(nil))
}

func (o DatabaseIAMPolicyArrayOutput) ToDatabaseIAMPolicyArrayOutput() DatabaseIAMPolicyArrayOutput {
	return o
}

func (o DatabaseIAMPolicyArrayOutput) ToDatabaseIAMPolicyArrayOutputWithContext(ctx context.Context) DatabaseIAMPolicyArrayOutput {
	return o
}

func (o DatabaseIAMPolicyArrayOutput) Index(i pulumi.IntInput) DatabaseIAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseIAMPolicy {
		return vs[0].([]DatabaseIAMPolicy)[vs[1].(int)]
	}).(DatabaseIAMPolicyOutput)
}

type DatabaseIAMPolicyMapOutput struct{ *pulumi.OutputState }

func (DatabaseIAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabaseIAMPolicy)(nil))
}

func (o DatabaseIAMPolicyMapOutput) ToDatabaseIAMPolicyMapOutput() DatabaseIAMPolicyMapOutput {
	return o
}

func (o DatabaseIAMPolicyMapOutput) ToDatabaseIAMPolicyMapOutputWithContext(ctx context.Context) DatabaseIAMPolicyMapOutput {
	return o
}

func (o DatabaseIAMPolicyMapOutput) MapIndex(k pulumi.StringInput) DatabaseIAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabaseIAMPolicy {
		return vs[0].(map[string]DatabaseIAMPolicy)[vs[1].(string)]
	}).(DatabaseIAMPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseIAMPolicyOutput{})
	pulumi.RegisterOutputType(DatabaseIAMPolicyPtrOutput{})
	pulumi.RegisterOutputType(DatabaseIAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(DatabaseIAMPolicyMapOutput{})
}
