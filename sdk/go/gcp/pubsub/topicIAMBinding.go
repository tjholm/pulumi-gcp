// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Pub/Sub Topic. Each of these resources serves a different use case:
//
// * `pubsub.TopicIAMPolicy`: Authoritative. Sets the IAM policy for the topic and replaces any existing policy already attached.
// * `pubsub.TopicIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the topic are preserved.
// * `pubsub.TopicIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the topic are preserved.
//
// > **Note:** `pubsub.TopicIAMPolicy` **cannot** be used in conjunction with `pubsub.TopicIAMBinding` and `pubsub.TopicIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `pubsub.TopicIAMBinding` resources **can be** used in conjunction with `pubsub.TopicIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_pubsub\_topic\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewTopicIAMPolicy(ctx, "policy", &pubsub.TopicIAMPolicyArgs{
// 			Project:    pulumi.Any(google_pubsub_topic.Example.Project),
// 			Topic:      pulumi.Any(google_pubsub_topic.Example.Name),
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
// ## google\_pubsub\_topic\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewTopicIAMBinding(ctx, "binding", &pubsub.TopicIAMBindingArgs{
// 			Project: pulumi.Any(google_pubsub_topic.Example.Project),
// 			Topic:   pulumi.Any(google_pubsub_topic.Example.Name),
// 			Role:    pulumi.String("roles/viewer"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_topic\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewTopicIAMMember(ctx, "member", &pubsub.TopicIAMMemberArgs{
// 			Project: pulumi.Any(google_pubsub_topic.Example.Project),
// 			Topic:   pulumi.Any(google_pubsub_topic.Example.Name),
// 			Role:    pulumi.String("roles/viewer"),
// 			Member:  pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/topics/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Pub/Sub topic IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:pubsub/topicIAMBinding:TopicIAMBinding editor "projects/{{project}}/topics/{{topic}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:pubsub/topicIAMBinding:TopicIAMBinding editor "projects/{{project}}/topics/{{topic}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:pubsub/topicIAMBinding:TopicIAMBinding editor projects/{{project}}/topics/{{topic}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TopicIAMBinding struct {
	pulumi.CustomResourceState

	Condition TopicIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewTopicIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewTopicIAMBinding(ctx *pulumi.Context,
	name string, args *TopicIAMBindingArgs, opts ...pulumi.ResourceOption) (*TopicIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	var resource TopicIAMBinding
	err := ctx.RegisterResource("gcp:pubsub/topicIAMBinding:TopicIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicIAMBinding gets an existing TopicIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicIAMBindingState, opts ...pulumi.ResourceOption) (*TopicIAMBinding, error) {
	var resource TopicIAMBinding
	err := ctx.ReadResource("gcp:pubsub/topicIAMBinding:TopicIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicIAMBinding resources.
type topicIAMBindingState struct {
	Condition *TopicIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Topic *string `pulumi:"topic"`
}

type TopicIAMBindingState struct {
	Condition TopicIAMBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Topic pulumi.StringPtrInput
}

func (TopicIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicIAMBindingState)(nil)).Elem()
}

type topicIAMBindingArgs struct {
	Condition *TopicIAMBindingCondition `pulumi:"condition"`
	Members   []string                  `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Topic string `pulumi:"topic"`
}

// The set of arguments for constructing a TopicIAMBinding resource.
type TopicIAMBindingArgs struct {
	Condition TopicIAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Topic pulumi.StringInput
}

func (TopicIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicIAMBindingArgs)(nil)).Elem()
}

type TopicIAMBindingInput interface {
	pulumi.Input

	ToTopicIAMBindingOutput() TopicIAMBindingOutput
	ToTopicIAMBindingOutputWithContext(ctx context.Context) TopicIAMBindingOutput
}

func (*TopicIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicIAMBinding)(nil)).Elem()
}

func (i *TopicIAMBinding) ToTopicIAMBindingOutput() TopicIAMBindingOutput {
	return i.ToTopicIAMBindingOutputWithContext(context.Background())
}

func (i *TopicIAMBinding) ToTopicIAMBindingOutputWithContext(ctx context.Context) TopicIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMBindingOutput)
}

// TopicIAMBindingArrayInput is an input type that accepts TopicIAMBindingArray and TopicIAMBindingArrayOutput values.
// You can construct a concrete instance of `TopicIAMBindingArrayInput` via:
//
//          TopicIAMBindingArray{ TopicIAMBindingArgs{...} }
type TopicIAMBindingArrayInput interface {
	pulumi.Input

	ToTopicIAMBindingArrayOutput() TopicIAMBindingArrayOutput
	ToTopicIAMBindingArrayOutputWithContext(context.Context) TopicIAMBindingArrayOutput
}

type TopicIAMBindingArray []TopicIAMBindingInput

func (TopicIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TopicIAMBinding)(nil)).Elem()
}

func (i TopicIAMBindingArray) ToTopicIAMBindingArrayOutput() TopicIAMBindingArrayOutput {
	return i.ToTopicIAMBindingArrayOutputWithContext(context.Background())
}

func (i TopicIAMBindingArray) ToTopicIAMBindingArrayOutputWithContext(ctx context.Context) TopicIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMBindingArrayOutput)
}

// TopicIAMBindingMapInput is an input type that accepts TopicIAMBindingMap and TopicIAMBindingMapOutput values.
// You can construct a concrete instance of `TopicIAMBindingMapInput` via:
//
//          TopicIAMBindingMap{ "key": TopicIAMBindingArgs{...} }
type TopicIAMBindingMapInput interface {
	pulumi.Input

	ToTopicIAMBindingMapOutput() TopicIAMBindingMapOutput
	ToTopicIAMBindingMapOutputWithContext(context.Context) TopicIAMBindingMapOutput
}

type TopicIAMBindingMap map[string]TopicIAMBindingInput

func (TopicIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TopicIAMBinding)(nil)).Elem()
}

func (i TopicIAMBindingMap) ToTopicIAMBindingMapOutput() TopicIAMBindingMapOutput {
	return i.ToTopicIAMBindingMapOutputWithContext(context.Background())
}

func (i TopicIAMBindingMap) ToTopicIAMBindingMapOutputWithContext(ctx context.Context) TopicIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMBindingMapOutput)
}

type TopicIAMBindingOutput struct{ *pulumi.OutputState }

func (TopicIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicIAMBinding)(nil)).Elem()
}

func (o TopicIAMBindingOutput) ToTopicIAMBindingOutput() TopicIAMBindingOutput {
	return o
}

func (o TopicIAMBindingOutput) ToTopicIAMBindingOutputWithContext(ctx context.Context) TopicIAMBindingOutput {
	return o
}

type TopicIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (TopicIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TopicIAMBinding)(nil)).Elem()
}

func (o TopicIAMBindingArrayOutput) ToTopicIAMBindingArrayOutput() TopicIAMBindingArrayOutput {
	return o
}

func (o TopicIAMBindingArrayOutput) ToTopicIAMBindingArrayOutputWithContext(ctx context.Context) TopicIAMBindingArrayOutput {
	return o
}

func (o TopicIAMBindingArrayOutput) Index(i pulumi.IntInput) TopicIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TopicIAMBinding {
		return vs[0].([]*TopicIAMBinding)[vs[1].(int)]
	}).(TopicIAMBindingOutput)
}

type TopicIAMBindingMapOutput struct{ *pulumi.OutputState }

func (TopicIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TopicIAMBinding)(nil)).Elem()
}

func (o TopicIAMBindingMapOutput) ToTopicIAMBindingMapOutput() TopicIAMBindingMapOutput {
	return o
}

func (o TopicIAMBindingMapOutput) ToTopicIAMBindingMapOutputWithContext(ctx context.Context) TopicIAMBindingMapOutput {
	return o
}

func (o TopicIAMBindingMapOutput) MapIndex(k pulumi.StringInput) TopicIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TopicIAMBinding {
		return vs[0].(map[string]*TopicIAMBinding)[vs[1].(string)]
	}).(TopicIAMBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicIAMBindingInput)(nil)).Elem(), &TopicIAMBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicIAMBindingArrayInput)(nil)).Elem(), TopicIAMBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicIAMBindingMapInput)(nil)).Elem(), TopicIAMBindingMap{})
	pulumi.RegisterOutputType(TopicIAMBindingOutput{})
	pulumi.RegisterOutputType(TopicIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(TopicIAMBindingMapOutput{})
}
