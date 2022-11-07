// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigqueryanalyticshub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Bigquery Analytics Hub DataExchange. Each of these resources serves a different use case:
//
// * `bigqueryanalyticshub.DataExchangeIamPolicy`: Authoritative. Sets the IAM policy for the dataexchange and replaces any existing policy already attached.
// * `bigqueryanalyticshub.DataExchangeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataexchange are preserved.
// * `bigqueryanalyticshub.DataExchangeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataexchange are preserved.
//
// > **Note:** `bigqueryanalyticshub.DataExchangeIamPolicy` **cannot** be used in conjunction with `bigqueryanalyticshub.DataExchangeIamBinding` and `bigqueryanalyticshub.DataExchangeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigqueryanalyticshub.DataExchangeIamBinding` resources **can be** used in conjunction with `bigqueryanalyticshub.DataExchangeIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigquery\_analytics\_hub\_data\_exchange\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigqueryanalyticshub"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					organizations.GetIAMPolicyBinding{
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
//			_, err = bigqueryanalyticshub.NewDataExchangeIamPolicy(ctx, "policy", &bigqueryanalyticshub.DataExchangeIamPolicyArgs{
//				Project:        pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Project),
//				Location:       pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Location),
//				DataExchangeId: pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Data_exchange_id),
//				PolicyData:     pulumi.String(admin.PolicyData),
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
// ## google\_bigquery\_analytics\_hub\_data\_exchange\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigqueryanalyticshub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigqueryanalyticshub.NewDataExchangeIamBinding(ctx, "binding", &bigqueryanalyticshub.DataExchangeIamBindingArgs{
//				Project:        pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Project),
//				Location:       pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Location),
//				DataExchangeId: pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Data_exchange_id),
//				Role:           pulumi.String("roles/viewer"),
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
// ## google\_bigquery\_analytics\_hub\_data\_exchange\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigqueryanalyticshub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigqueryanalyticshub.NewDataExchangeIamMember(ctx, "member", &bigqueryanalyticshub.DataExchangeIamMemberArgs{
//				Project:        pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Project),
//				Location:       pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Location),
//				DataExchangeId: pulumi.Any(google_bigquery_analytics_hub_data_exchange.Data_exchange.Data_exchange_id),
//				Role:           pulumi.String("roles/viewer"),
//				Member:         pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}} * {{project}}/{{location}}/{{data_exchange_id}} * {{location}}/{{data_exchange_id}} * {{data_exchange_id}} Any variables not passed in the import command will be taken from the provider configuration. Bigquery Analytics Hub dataexchange IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigqueryanalyticshub/dataExchangeIamBinding:DataExchangeIamBinding editor "projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigqueryanalyticshub/dataExchangeIamBinding:DataExchangeIamBinding editor "projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:bigqueryanalyticshub/dataExchangeIamBinding:DataExchangeIamBinding editor projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DataExchangeIamBinding struct {
	pulumi.CustomResourceState

	Condition DataExchangeIamBindingConditionPtrOutput `pulumi:"condition"`
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Used to find the parent resource to bind the IAM policy to
	DataExchangeId pulumi.StringOutput `pulumi:"dataExchangeId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the location this data exchange.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigqueryanalyticshub.DataExchangeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDataExchangeIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataExchangeIamBinding(ctx *pulumi.Context,
	name string, args *DataExchangeIamBindingArgs, opts ...pulumi.ResourceOption) (*DataExchangeIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataExchangeId == nil {
		return nil, errors.New("invalid value for required argument 'DataExchangeId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DataExchangeIamBinding
	err := ctx.RegisterResource("gcp:bigqueryanalyticshub/dataExchangeIamBinding:DataExchangeIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataExchangeIamBinding gets an existing DataExchangeIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataExchangeIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataExchangeIamBindingState, opts ...pulumi.ResourceOption) (*DataExchangeIamBinding, error) {
	var resource DataExchangeIamBinding
	err := ctx.ReadResource("gcp:bigqueryanalyticshub/dataExchangeIamBinding:DataExchangeIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataExchangeIamBinding resources.
type dataExchangeIamBindingState struct {
	Condition *DataExchangeIamBindingCondition `pulumi:"condition"`
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Used to find the parent resource to bind the IAM policy to
	DataExchangeId *string `pulumi:"dataExchangeId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the location this data exchange.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigqueryanalyticshub.DataExchangeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DataExchangeIamBindingState struct {
	Condition DataExchangeIamBindingConditionPtrInput
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Used to find the parent resource to bind the IAM policy to
	DataExchangeId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the location this data exchange.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigqueryanalyticshub.DataExchangeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DataExchangeIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeIamBindingState)(nil)).Elem()
}

type dataExchangeIamBindingArgs struct {
	Condition *DataExchangeIamBindingCondition `pulumi:"condition"`
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Used to find the parent resource to bind the IAM policy to
	DataExchangeId string `pulumi:"dataExchangeId"`
	// The name of the location this data exchange.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigqueryanalyticshub.DataExchangeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DataExchangeIamBinding resource.
type DataExchangeIamBindingArgs struct {
	Condition DataExchangeIamBindingConditionPtrInput
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Used to find the parent resource to bind the IAM policy to
	DataExchangeId pulumi.StringInput
	// The name of the location this data exchange.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigqueryanalyticshub.DataExchangeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DataExchangeIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeIamBindingArgs)(nil)).Elem()
}

type DataExchangeIamBindingInput interface {
	pulumi.Input

	ToDataExchangeIamBindingOutput() DataExchangeIamBindingOutput
	ToDataExchangeIamBindingOutputWithContext(ctx context.Context) DataExchangeIamBindingOutput
}

func (*DataExchangeIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchangeIamBinding)(nil)).Elem()
}

func (i *DataExchangeIamBinding) ToDataExchangeIamBindingOutput() DataExchangeIamBindingOutput {
	return i.ToDataExchangeIamBindingOutputWithContext(context.Background())
}

func (i *DataExchangeIamBinding) ToDataExchangeIamBindingOutputWithContext(ctx context.Context) DataExchangeIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeIamBindingOutput)
}

// DataExchangeIamBindingArrayInput is an input type that accepts DataExchangeIamBindingArray and DataExchangeIamBindingArrayOutput values.
// You can construct a concrete instance of `DataExchangeIamBindingArrayInput` via:
//
//	DataExchangeIamBindingArray{ DataExchangeIamBindingArgs{...} }
type DataExchangeIamBindingArrayInput interface {
	pulumi.Input

	ToDataExchangeIamBindingArrayOutput() DataExchangeIamBindingArrayOutput
	ToDataExchangeIamBindingArrayOutputWithContext(context.Context) DataExchangeIamBindingArrayOutput
}

type DataExchangeIamBindingArray []DataExchangeIamBindingInput

func (DataExchangeIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataExchangeIamBinding)(nil)).Elem()
}

func (i DataExchangeIamBindingArray) ToDataExchangeIamBindingArrayOutput() DataExchangeIamBindingArrayOutput {
	return i.ToDataExchangeIamBindingArrayOutputWithContext(context.Background())
}

func (i DataExchangeIamBindingArray) ToDataExchangeIamBindingArrayOutputWithContext(ctx context.Context) DataExchangeIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeIamBindingArrayOutput)
}

// DataExchangeIamBindingMapInput is an input type that accepts DataExchangeIamBindingMap and DataExchangeIamBindingMapOutput values.
// You can construct a concrete instance of `DataExchangeIamBindingMapInput` via:
//
//	DataExchangeIamBindingMap{ "key": DataExchangeIamBindingArgs{...} }
type DataExchangeIamBindingMapInput interface {
	pulumi.Input

	ToDataExchangeIamBindingMapOutput() DataExchangeIamBindingMapOutput
	ToDataExchangeIamBindingMapOutputWithContext(context.Context) DataExchangeIamBindingMapOutput
}

type DataExchangeIamBindingMap map[string]DataExchangeIamBindingInput

func (DataExchangeIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataExchangeIamBinding)(nil)).Elem()
}

func (i DataExchangeIamBindingMap) ToDataExchangeIamBindingMapOutput() DataExchangeIamBindingMapOutput {
	return i.ToDataExchangeIamBindingMapOutputWithContext(context.Background())
}

func (i DataExchangeIamBindingMap) ToDataExchangeIamBindingMapOutputWithContext(ctx context.Context) DataExchangeIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeIamBindingMapOutput)
}

type DataExchangeIamBindingOutput struct{ *pulumi.OutputState }

func (DataExchangeIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchangeIamBinding)(nil)).Elem()
}

func (o DataExchangeIamBindingOutput) ToDataExchangeIamBindingOutput() DataExchangeIamBindingOutput {
	return o
}

func (o DataExchangeIamBindingOutput) ToDataExchangeIamBindingOutputWithContext(ctx context.Context) DataExchangeIamBindingOutput {
	return o
}

func (o DataExchangeIamBindingOutput) Condition() DataExchangeIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DataExchangeIamBinding) DataExchangeIamBindingConditionPtrOutput { return v.Condition }).(DataExchangeIamBindingConditionPtrOutput)
}

// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Used to find the parent resource to bind the IAM policy to
func (o DataExchangeIamBindingOutput) DataExchangeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamBinding) pulumi.StringOutput { return v.DataExchangeId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o DataExchangeIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the location this data exchange.
// Used to find the parent resource to bind the IAM policy to
func (o DataExchangeIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataExchangeIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataExchangeIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o DataExchangeIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `bigqueryanalyticshub.DataExchangeIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o DataExchangeIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type DataExchangeIamBindingArrayOutput struct{ *pulumi.OutputState }

func (DataExchangeIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataExchangeIamBinding)(nil)).Elem()
}

func (o DataExchangeIamBindingArrayOutput) ToDataExchangeIamBindingArrayOutput() DataExchangeIamBindingArrayOutput {
	return o
}

func (o DataExchangeIamBindingArrayOutput) ToDataExchangeIamBindingArrayOutputWithContext(ctx context.Context) DataExchangeIamBindingArrayOutput {
	return o
}

func (o DataExchangeIamBindingArrayOutput) Index(i pulumi.IntInput) DataExchangeIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataExchangeIamBinding {
		return vs[0].([]*DataExchangeIamBinding)[vs[1].(int)]
	}).(DataExchangeIamBindingOutput)
}

type DataExchangeIamBindingMapOutput struct{ *pulumi.OutputState }

func (DataExchangeIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataExchangeIamBinding)(nil)).Elem()
}

func (o DataExchangeIamBindingMapOutput) ToDataExchangeIamBindingMapOutput() DataExchangeIamBindingMapOutput {
	return o
}

func (o DataExchangeIamBindingMapOutput) ToDataExchangeIamBindingMapOutputWithContext(ctx context.Context) DataExchangeIamBindingMapOutput {
	return o
}

func (o DataExchangeIamBindingMapOutput) MapIndex(k pulumi.StringInput) DataExchangeIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataExchangeIamBinding {
		return vs[0].(map[string]*DataExchangeIamBinding)[vs[1].(string)]
	}).(DataExchangeIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeIamBindingInput)(nil)).Elem(), &DataExchangeIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeIamBindingArrayInput)(nil)).Elem(), DataExchangeIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeIamBindingMapInput)(nil)).Elem(), DataExchangeIamBindingMap{})
	pulumi.RegisterOutputType(DataExchangeIamBindingOutput{})
	pulumi.RegisterOutputType(DataExchangeIamBindingArrayOutput{})
	pulumi.RegisterOutputType(DataExchangeIamBindingMapOutput{})
}
