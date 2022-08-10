// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpoints

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Endpoints ServiceConsumers. Each of these resources serves a different use case:
//
// * `endpoints.ConsumersIamPolicy`: Authoritative. Sets the IAM policy for the serviceconsumers and replaces any existing policy already attached.
// * `endpoints.ConsumersIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the serviceconsumers are preserved.
// * `endpoints.ConsumersIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the serviceconsumers are preserved.
//
// > **Note:** `endpoints.ConsumersIamPolicy` **cannot** be used in conjunction with `endpoints.ConsumersIamBinding` and `endpoints.ConsumersIamMember` or they will fight over what your policy should be.
//
// > **Note:** `endpoints.ConsumersIamBinding` resources **can be** used in conjunction with `endpoints.ConsumersIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* services/{{service_name}}/consumers/{{consumer_project}} * {{service_name}}/{{consumer_project}} * {{consumer_project}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Endpoints serviceconsumers IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor "services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor "services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor services/{{service_name}}/consumers/{{consumer_project}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ConsumersIamPolicy struct {
	pulumi.CustomResourceState

	ConsumerProject pulumi.StringOutput `pulumi:"consumerProject"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData  pulumi.StringOutput `pulumi:"policyData"`
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewConsumersIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewConsumersIamPolicy(ctx *pulumi.Context,
	name string, args *ConsumersIamPolicyArgs, opts ...pulumi.ResourceOption) (*ConsumersIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerProject == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerProject'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource ConsumersIamPolicy
	err := ctx.RegisterResource("gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumersIamPolicy gets an existing ConsumersIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumersIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumersIamPolicyState, opts ...pulumi.ResourceOption) (*ConsumersIamPolicy, error) {
	var resource ConsumersIamPolicy
	err := ctx.ReadResource("gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumersIamPolicy resources.
type consumersIamPolicyState struct {
	ConsumerProject *string `pulumi:"consumerProject"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData  *string `pulumi:"policyData"`
	ServiceName *string `pulumi:"serviceName"`
}

type ConsumersIamPolicyState struct {
	ConsumerProject pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData  pulumi.StringPtrInput
	ServiceName pulumi.StringPtrInput
}

func (ConsumersIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumersIamPolicyState)(nil)).Elem()
}

type consumersIamPolicyArgs struct {
	ConsumerProject string `pulumi:"consumerProject"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData  string `pulumi:"policyData"`
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ConsumersIamPolicy resource.
type ConsumersIamPolicyArgs struct {
	ConsumerProject pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData  pulumi.StringInput
	ServiceName pulumi.StringInput
}

func (ConsumersIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumersIamPolicyArgs)(nil)).Elem()
}

type ConsumersIamPolicyInput interface {
	pulumi.Input

	ToConsumersIamPolicyOutput() ConsumersIamPolicyOutput
	ToConsumersIamPolicyOutputWithContext(ctx context.Context) ConsumersIamPolicyOutput
}

func (*ConsumersIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumersIamPolicy)(nil)).Elem()
}

func (i *ConsumersIamPolicy) ToConsumersIamPolicyOutput() ConsumersIamPolicyOutput {
	return i.ToConsumersIamPolicyOutputWithContext(context.Background())
}

func (i *ConsumersIamPolicy) ToConsumersIamPolicyOutputWithContext(ctx context.Context) ConsumersIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumersIamPolicyOutput)
}

// ConsumersIamPolicyArrayInput is an input type that accepts ConsumersIamPolicyArray and ConsumersIamPolicyArrayOutput values.
// You can construct a concrete instance of `ConsumersIamPolicyArrayInput` via:
//
//	ConsumersIamPolicyArray{ ConsumersIamPolicyArgs{...} }
type ConsumersIamPolicyArrayInput interface {
	pulumi.Input

	ToConsumersIamPolicyArrayOutput() ConsumersIamPolicyArrayOutput
	ToConsumersIamPolicyArrayOutputWithContext(context.Context) ConsumersIamPolicyArrayOutput
}

type ConsumersIamPolicyArray []ConsumersIamPolicyInput

func (ConsumersIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumersIamPolicy)(nil)).Elem()
}

func (i ConsumersIamPolicyArray) ToConsumersIamPolicyArrayOutput() ConsumersIamPolicyArrayOutput {
	return i.ToConsumersIamPolicyArrayOutputWithContext(context.Background())
}

func (i ConsumersIamPolicyArray) ToConsumersIamPolicyArrayOutputWithContext(ctx context.Context) ConsumersIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumersIamPolicyArrayOutput)
}

// ConsumersIamPolicyMapInput is an input type that accepts ConsumersIamPolicyMap and ConsumersIamPolicyMapOutput values.
// You can construct a concrete instance of `ConsumersIamPolicyMapInput` via:
//
//	ConsumersIamPolicyMap{ "key": ConsumersIamPolicyArgs{...} }
type ConsumersIamPolicyMapInput interface {
	pulumi.Input

	ToConsumersIamPolicyMapOutput() ConsumersIamPolicyMapOutput
	ToConsumersIamPolicyMapOutputWithContext(context.Context) ConsumersIamPolicyMapOutput
}

type ConsumersIamPolicyMap map[string]ConsumersIamPolicyInput

func (ConsumersIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumersIamPolicy)(nil)).Elem()
}

func (i ConsumersIamPolicyMap) ToConsumersIamPolicyMapOutput() ConsumersIamPolicyMapOutput {
	return i.ToConsumersIamPolicyMapOutputWithContext(context.Background())
}

func (i ConsumersIamPolicyMap) ToConsumersIamPolicyMapOutputWithContext(ctx context.Context) ConsumersIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumersIamPolicyMapOutput)
}

type ConsumersIamPolicyOutput struct{ *pulumi.OutputState }

func (ConsumersIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumersIamPolicy)(nil)).Elem()
}

func (o ConsumersIamPolicyOutput) ToConsumersIamPolicyOutput() ConsumersIamPolicyOutput {
	return o
}

func (o ConsumersIamPolicyOutput) ToConsumersIamPolicyOutputWithContext(ctx context.Context) ConsumersIamPolicyOutput {
	return o
}

func (o ConsumersIamPolicyOutput) ConsumerProject() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumersIamPolicy) pulumi.StringOutput { return v.ConsumerProject }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o ConsumersIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumersIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o ConsumersIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumersIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o ConsumersIamPolicyOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumersIamPolicy) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type ConsumersIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (ConsumersIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumersIamPolicy)(nil)).Elem()
}

func (o ConsumersIamPolicyArrayOutput) ToConsumersIamPolicyArrayOutput() ConsumersIamPolicyArrayOutput {
	return o
}

func (o ConsumersIamPolicyArrayOutput) ToConsumersIamPolicyArrayOutputWithContext(ctx context.Context) ConsumersIamPolicyArrayOutput {
	return o
}

func (o ConsumersIamPolicyArrayOutput) Index(i pulumi.IntInput) ConsumersIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsumersIamPolicy {
		return vs[0].([]*ConsumersIamPolicy)[vs[1].(int)]
	}).(ConsumersIamPolicyOutput)
}

type ConsumersIamPolicyMapOutput struct{ *pulumi.OutputState }

func (ConsumersIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumersIamPolicy)(nil)).Elem()
}

func (o ConsumersIamPolicyMapOutput) ToConsumersIamPolicyMapOutput() ConsumersIamPolicyMapOutput {
	return o
}

func (o ConsumersIamPolicyMapOutput) ToConsumersIamPolicyMapOutputWithContext(ctx context.Context) ConsumersIamPolicyMapOutput {
	return o
}

func (o ConsumersIamPolicyMapOutput) MapIndex(k pulumi.StringInput) ConsumersIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsumersIamPolicy {
		return vs[0].(map[string]*ConsumersIamPolicy)[vs[1].(string)]
	}).(ConsumersIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumersIamPolicyInput)(nil)).Elem(), &ConsumersIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumersIamPolicyArrayInput)(nil)).Elem(), ConsumersIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumersIamPolicyMapInput)(nil)).Elem(), ConsumersIamPolicyMap{})
	pulumi.RegisterOutputType(ConsumersIamPolicyOutput{})
	pulumi.RegisterOutputType(ConsumersIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(ConsumersIamPolicyMapOutput{})
}
