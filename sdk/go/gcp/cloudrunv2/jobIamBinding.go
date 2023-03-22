// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrunv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Run (v2 API) Job. Each of these resources serves a different use case:
//
// * `cloudrunv2.JobIamPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
// * `cloudrunv2.JobIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
// * `cloudrunv2.JobIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
//
// > **Note:** `cloudrunv2.JobIamPolicy` **cannot** be used in conjunction with `cloudrunv2.JobIamBinding` and `cloudrunv2.JobIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudrunv2.JobIamBinding` resources **can be** used in conjunction with `cloudrunv2.JobIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_cloud\_run\_v2\_job\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
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
//			_, err = cloudrunv2.NewJobIamPolicy(ctx, "policy", &cloudrunv2.JobIamPolicyArgs{
//				Project:    pulumi.Any(google_cloud_run_v2_job.Default.Project),
//				Location:   pulumi.Any(google_cloud_run_v2_job.Default.Location),
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
// ## google\_cloud\_run\_v2\_job\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrunv2.NewJobIamBinding(ctx, "binding", &cloudrunv2.JobIamBindingArgs{
//				Project:  pulumi.Any(google_cloud_run_v2_job.Default.Project),
//				Location: pulumi.Any(google_cloud_run_v2_job.Default.Location),
//				Role:     pulumi.String("roles/viewer"),
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
// ## google\_cloud\_run\_v2\_job\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrunv2.NewJobIamMember(ctx, "member", &cloudrunv2.JobIamMemberArgs{
//				Project:  pulumi.Any(google_cloud_run_v2_job.Default.Project),
//				Location: pulumi.Any(google_cloud_run_v2_job.Default.Location),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/jobs/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Run (v2 API) job IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/jobIamBinding:JobIamBinding editor "projects/{{project}}/locations/{{location}}/jobs/{{job}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/jobIamBinding:JobIamBinding editor "projects/{{project}}/locations/{{location}}/jobs/{{job}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/jobIamBinding:JobIamBinding editor projects/{{project}}/locations/{{location}}/jobs/{{job}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type JobIamBinding struct {
	pulumi.CustomResourceState

	Condition JobIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of the cloud run job Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrunv2.JobIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewJobIamBinding registers a new resource with the given unique name, arguments, and options.
func NewJobIamBinding(ctx *pulumi.Context,
	name string, args *JobIamBindingArgs, opts ...pulumi.ResourceOption) (*JobIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource JobIamBinding
	err := ctx.RegisterResource("gcp:cloudrunv2/jobIamBinding:JobIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobIamBinding gets an existing JobIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobIamBindingState, opts ...pulumi.ResourceOption) (*JobIamBinding, error) {
	var resource JobIamBinding
	err := ctx.ReadResource("gcp:cloudrunv2/jobIamBinding:JobIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobIamBinding resources.
type jobIamBindingState struct {
	Condition *JobIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location of the cloud run job Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrunv2.JobIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type JobIamBindingState struct {
	Condition JobIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location of the cloud run job Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrunv2.JobIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (JobIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIamBindingState)(nil)).Elem()
}

type jobIamBindingArgs struct {
	Condition *JobIamBindingCondition `pulumi:"condition"`
	// The location of the cloud run job Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrunv2.JobIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a JobIamBinding resource.
type JobIamBindingArgs struct {
	Condition JobIamBindingConditionPtrInput
	// The location of the cloud run job Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrunv2.JobIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (JobIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIamBindingArgs)(nil)).Elem()
}

type JobIamBindingInput interface {
	pulumi.Input

	ToJobIamBindingOutput() JobIamBindingOutput
	ToJobIamBindingOutputWithContext(ctx context.Context) JobIamBindingOutput
}

func (*JobIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**JobIamBinding)(nil)).Elem()
}

func (i *JobIamBinding) ToJobIamBindingOutput() JobIamBindingOutput {
	return i.ToJobIamBindingOutputWithContext(context.Background())
}

func (i *JobIamBinding) ToJobIamBindingOutputWithContext(ctx context.Context) JobIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIamBindingOutput)
}

// JobIamBindingArrayInput is an input type that accepts JobIamBindingArray and JobIamBindingArrayOutput values.
// You can construct a concrete instance of `JobIamBindingArrayInput` via:
//
//	JobIamBindingArray{ JobIamBindingArgs{...} }
type JobIamBindingArrayInput interface {
	pulumi.Input

	ToJobIamBindingArrayOutput() JobIamBindingArrayOutput
	ToJobIamBindingArrayOutputWithContext(context.Context) JobIamBindingArrayOutput
}

type JobIamBindingArray []JobIamBindingInput

func (JobIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobIamBinding)(nil)).Elem()
}

func (i JobIamBindingArray) ToJobIamBindingArrayOutput() JobIamBindingArrayOutput {
	return i.ToJobIamBindingArrayOutputWithContext(context.Background())
}

func (i JobIamBindingArray) ToJobIamBindingArrayOutputWithContext(ctx context.Context) JobIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIamBindingArrayOutput)
}

// JobIamBindingMapInput is an input type that accepts JobIamBindingMap and JobIamBindingMapOutput values.
// You can construct a concrete instance of `JobIamBindingMapInput` via:
//
//	JobIamBindingMap{ "key": JobIamBindingArgs{...} }
type JobIamBindingMapInput interface {
	pulumi.Input

	ToJobIamBindingMapOutput() JobIamBindingMapOutput
	ToJobIamBindingMapOutputWithContext(context.Context) JobIamBindingMapOutput
}

type JobIamBindingMap map[string]JobIamBindingInput

func (JobIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobIamBinding)(nil)).Elem()
}

func (i JobIamBindingMap) ToJobIamBindingMapOutput() JobIamBindingMapOutput {
	return i.ToJobIamBindingMapOutputWithContext(context.Background())
}

func (i JobIamBindingMap) ToJobIamBindingMapOutputWithContext(ctx context.Context) JobIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIamBindingMapOutput)
}

type JobIamBindingOutput struct{ *pulumi.OutputState }

func (JobIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobIamBinding)(nil)).Elem()
}

func (o JobIamBindingOutput) ToJobIamBindingOutput() JobIamBindingOutput {
	return o
}

func (o JobIamBindingOutput) ToJobIamBindingOutputWithContext(ctx context.Context) JobIamBindingOutput {
	return o
}

func (o JobIamBindingOutput) Condition() JobIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *JobIamBinding) JobIamBindingConditionPtrOutput { return v.Condition }).(JobIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o JobIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *JobIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location of the cloud run job Used to find the parent resource to bind the IAM policy to
func (o JobIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *JobIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o JobIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JobIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o JobIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o JobIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *JobIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `cloudrunv2.JobIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o JobIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *JobIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type JobIamBindingArrayOutput struct{ *pulumi.OutputState }

func (JobIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobIamBinding)(nil)).Elem()
}

func (o JobIamBindingArrayOutput) ToJobIamBindingArrayOutput() JobIamBindingArrayOutput {
	return o
}

func (o JobIamBindingArrayOutput) ToJobIamBindingArrayOutputWithContext(ctx context.Context) JobIamBindingArrayOutput {
	return o
}

func (o JobIamBindingArrayOutput) Index(i pulumi.IntInput) JobIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *JobIamBinding {
		return vs[0].([]*JobIamBinding)[vs[1].(int)]
	}).(JobIamBindingOutput)
}

type JobIamBindingMapOutput struct{ *pulumi.OutputState }

func (JobIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobIamBinding)(nil)).Elem()
}

func (o JobIamBindingMapOutput) ToJobIamBindingMapOutput() JobIamBindingMapOutput {
	return o
}

func (o JobIamBindingMapOutput) ToJobIamBindingMapOutputWithContext(ctx context.Context) JobIamBindingMapOutput {
	return o
}

func (o JobIamBindingMapOutput) MapIndex(k pulumi.StringInput) JobIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *JobIamBinding {
		return vs[0].(map[string]*JobIamBinding)[vs[1].(string)]
	}).(JobIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobIamBindingInput)(nil)).Elem(), &JobIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobIamBindingArrayInput)(nil)).Elem(), JobIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobIamBindingMapInput)(nil)).Elem(), JobIamBindingMap{})
	pulumi.RegisterOutputType(JobIamBindingOutput{})
	pulumi.RegisterOutputType(JobIamBindingArrayOutput{})
	pulumi.RegisterOutputType(JobIamBindingMapOutput{})
}
