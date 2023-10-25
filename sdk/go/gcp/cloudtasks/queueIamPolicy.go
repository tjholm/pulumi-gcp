// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtasks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for Cloud Tasks Queue. Each of these resources serves a different use case:
//
// * `cloudtasks.QueueIamPolicy`: Authoritative. Sets the IAM policy for the queue and replaces any existing policy already attached.
// * `cloudtasks.QueueIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the queue are preserved.
// * `cloudtasks.QueueIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the queue are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `cloudtasks.QueueIamPolicy`: Retrieves the IAM policy for the queue
//
// > **Note:** `cloudtasks.QueueIamPolicy` **cannot** be used in conjunction with `cloudtasks.QueueIamBinding` and `cloudtasks.QueueIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudtasks.QueueIamBinding` resources **can be** used in conjunction with `cloudtasks.QueueIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_cloud\_tasks\_queue\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudtasks"
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
//			_, err = cloudtasks.NewQueueIamPolicy(ctx, "policy", &cloudtasks.QueueIamPolicyArgs{
//				Project:    pulumi.Any(google_cloud_tasks_queue.Default.Project),
//				Location:   pulumi.Any(google_cloud_tasks_queue.Default.Location),
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
// ## google\_cloud\_tasks\_queue\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudtasks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudtasks.NewQueueIamBinding(ctx, "binding", &cloudtasks.QueueIamBindingArgs{
//				Project:  pulumi.Any(google_cloud_tasks_queue.Default.Project),
//				Location: pulumi.Any(google_cloud_tasks_queue.Default.Location),
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
// ## google\_cloud\_tasks\_queue\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudtasks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudtasks.NewQueueIamMember(ctx, "member", &cloudtasks.QueueIamMemberArgs{
//				Project:  pulumi.Any(google_cloud_tasks_queue.Default.Project),
//				Location: pulumi.Any(google_cloud_tasks_queue.Default.Location),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/queues/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Tasks queue IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudtasks/queueIamPolicy:QueueIamPolicy editor "projects/{{project}}/locations/{{location}}/queues/{{queue}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudtasks/queueIamPolicy:QueueIamPolicy editor "projects/{{project}}/locations/{{location}}/queues/{{queue}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudtasks/queueIamPolicy:QueueIamPolicy editor projects/{{project}}/locations/{{location}}/queues/{{queue}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type QueueIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of the queue Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
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
}

// NewQueueIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewQueueIamPolicy(ctx *pulumi.Context,
	name string, args *QueueIamPolicyArgs, opts ...pulumi.ResourceOption) (*QueueIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueueIamPolicy
	err := ctx.RegisterResource("gcp:cloudtasks/queueIamPolicy:QueueIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueueIamPolicy gets an existing QueueIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueueIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueIamPolicyState, opts ...pulumi.ResourceOption) (*QueueIamPolicy, error) {
	var resource QueueIamPolicy
	err := ctx.ReadResource("gcp:cloudtasks/queueIamPolicy:QueueIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueIamPolicy resources.
type queueIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location of the queue Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
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
}

type QueueIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location of the queue Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
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
}

func (QueueIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueIamPolicyState)(nil)).Elem()
}

type queueIamPolicyArgs struct {
	// The location of the queue Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
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
}

// The set of arguments for constructing a QueueIamPolicy resource.
type QueueIamPolicyArgs struct {
	// The location of the queue Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
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
}

func (QueueIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueIamPolicyArgs)(nil)).Elem()
}

type QueueIamPolicyInput interface {
	pulumi.Input

	ToQueueIamPolicyOutput() QueueIamPolicyOutput
	ToQueueIamPolicyOutputWithContext(ctx context.Context) QueueIamPolicyOutput
}

func (*QueueIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueIamPolicy)(nil)).Elem()
}

func (i *QueueIamPolicy) ToQueueIamPolicyOutput() QueueIamPolicyOutput {
	return i.ToQueueIamPolicyOutputWithContext(context.Background())
}

func (i *QueueIamPolicy) ToQueueIamPolicyOutputWithContext(ctx context.Context) QueueIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueIamPolicyOutput)
}

func (i *QueueIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*QueueIamPolicy] {
	return pulumix.Output[*QueueIamPolicy]{
		OutputState: i.ToQueueIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// QueueIamPolicyArrayInput is an input type that accepts QueueIamPolicyArray and QueueIamPolicyArrayOutput values.
// You can construct a concrete instance of `QueueIamPolicyArrayInput` via:
//
//	QueueIamPolicyArray{ QueueIamPolicyArgs{...} }
type QueueIamPolicyArrayInput interface {
	pulumi.Input

	ToQueueIamPolicyArrayOutput() QueueIamPolicyArrayOutput
	ToQueueIamPolicyArrayOutputWithContext(context.Context) QueueIamPolicyArrayOutput
}

type QueueIamPolicyArray []QueueIamPolicyInput

func (QueueIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueueIamPolicy)(nil)).Elem()
}

func (i QueueIamPolicyArray) ToQueueIamPolicyArrayOutput() QueueIamPolicyArrayOutput {
	return i.ToQueueIamPolicyArrayOutputWithContext(context.Background())
}

func (i QueueIamPolicyArray) ToQueueIamPolicyArrayOutputWithContext(ctx context.Context) QueueIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueIamPolicyArrayOutput)
}

func (i QueueIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*QueueIamPolicy] {
	return pulumix.Output[[]*QueueIamPolicy]{
		OutputState: i.ToQueueIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// QueueIamPolicyMapInput is an input type that accepts QueueIamPolicyMap and QueueIamPolicyMapOutput values.
// You can construct a concrete instance of `QueueIamPolicyMapInput` via:
//
//	QueueIamPolicyMap{ "key": QueueIamPolicyArgs{...} }
type QueueIamPolicyMapInput interface {
	pulumi.Input

	ToQueueIamPolicyMapOutput() QueueIamPolicyMapOutput
	ToQueueIamPolicyMapOutputWithContext(context.Context) QueueIamPolicyMapOutput
}

type QueueIamPolicyMap map[string]QueueIamPolicyInput

func (QueueIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueueIamPolicy)(nil)).Elem()
}

func (i QueueIamPolicyMap) ToQueueIamPolicyMapOutput() QueueIamPolicyMapOutput {
	return i.ToQueueIamPolicyMapOutputWithContext(context.Background())
}

func (i QueueIamPolicyMap) ToQueueIamPolicyMapOutputWithContext(ctx context.Context) QueueIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueIamPolicyMapOutput)
}

func (i QueueIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueueIamPolicy] {
	return pulumix.Output[map[string]*QueueIamPolicy]{
		OutputState: i.ToQueueIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type QueueIamPolicyOutput struct{ *pulumi.OutputState }

func (QueueIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueIamPolicy)(nil)).Elem()
}

func (o QueueIamPolicyOutput) ToQueueIamPolicyOutput() QueueIamPolicyOutput {
	return o
}

func (o QueueIamPolicyOutput) ToQueueIamPolicyOutputWithContext(ctx context.Context) QueueIamPolicyOutput {
	return o
}

func (o QueueIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*QueueIamPolicy] {
	return pulumix.Output[*QueueIamPolicy]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o QueueIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location of the queue Used to find the parent resource to bind the IAM policy to
func (o QueueIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o QueueIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o QueueIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
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
func (o QueueIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type QueueIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (QueueIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueueIamPolicy)(nil)).Elem()
}

func (o QueueIamPolicyArrayOutput) ToQueueIamPolicyArrayOutput() QueueIamPolicyArrayOutput {
	return o
}

func (o QueueIamPolicyArrayOutput) ToQueueIamPolicyArrayOutputWithContext(ctx context.Context) QueueIamPolicyArrayOutput {
	return o
}

func (o QueueIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*QueueIamPolicy] {
	return pulumix.Output[[]*QueueIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o QueueIamPolicyArrayOutput) Index(i pulumi.IntInput) QueueIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QueueIamPolicy {
		return vs[0].([]*QueueIamPolicy)[vs[1].(int)]
	}).(QueueIamPolicyOutput)
}

type QueueIamPolicyMapOutput struct{ *pulumi.OutputState }

func (QueueIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueueIamPolicy)(nil)).Elem()
}

func (o QueueIamPolicyMapOutput) ToQueueIamPolicyMapOutput() QueueIamPolicyMapOutput {
	return o
}

func (o QueueIamPolicyMapOutput) ToQueueIamPolicyMapOutputWithContext(ctx context.Context) QueueIamPolicyMapOutput {
	return o
}

func (o QueueIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueueIamPolicy] {
	return pulumix.Output[map[string]*QueueIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o QueueIamPolicyMapOutput) MapIndex(k pulumi.StringInput) QueueIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QueueIamPolicy {
		return vs[0].(map[string]*QueueIamPolicy)[vs[1].(string)]
	}).(QueueIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueIamPolicyInput)(nil)).Elem(), &QueueIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueIamPolicyArrayInput)(nil)).Elem(), QueueIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueIamPolicyMapInput)(nil)).Elem(), QueueIamPolicyMap{})
	pulumi.RegisterOutputType(QueueIamPolicyOutput{})
	pulumi.RegisterOutputType(QueueIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(QueueIamPolicyMapOutput{})
}
