// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * {{featurestore}}/entityTypes/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Vertex AI featurestoreentitytype IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamMember:AiFeatureStoreEntityTypeIamMember editor "{{featurestore}}/entityTypes/{{featurestore_entitytype}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamMember:AiFeatureStoreEntityTypeIamMember editor "{{featurestore}}/entityTypes/{{featurestore_entitytype}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamMember:AiFeatureStoreEntityTypeIamMember editor {{featurestore}}/entityTypes/{{featurestore_entitytype}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AiFeatureStoreEntityTypeIamMember struct {
	pulumi.CustomResourceState

	Condition AiFeatureStoreEntityTypeIamMemberConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Entitytype pulumi.StringOutput `pulumi:"entitytype"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringOutput `pulumi:"featurestore"`
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
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreEntityTypeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAiFeatureStoreEntityTypeIamMember registers a new resource with the given unique name, arguments, and options.
func NewAiFeatureStoreEntityTypeIamMember(ctx *pulumi.Context,
	name string, args *AiFeatureStoreEntityTypeIamMemberArgs, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityTypeIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Entitytype == nil {
		return nil, errors.New("invalid value for required argument 'Entitytype'")
	}
	if args.Featurestore == nil {
		return nil, errors.New("invalid value for required argument 'Featurestore'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiFeatureStoreEntityTypeIamMember
	err := ctx.RegisterResource("gcp:vertex/aiFeatureStoreEntityTypeIamMember:AiFeatureStoreEntityTypeIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiFeatureStoreEntityTypeIamMember gets an existing AiFeatureStoreEntityTypeIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiFeatureStoreEntityTypeIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiFeatureStoreEntityTypeIamMemberState, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityTypeIamMember, error) {
	var resource AiFeatureStoreEntityTypeIamMember
	err := ctx.ReadResource("gcp:vertex/aiFeatureStoreEntityTypeIamMember:AiFeatureStoreEntityTypeIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiFeatureStoreEntityTypeIamMember resources.
type aiFeatureStoreEntityTypeIamMemberState struct {
	Condition *AiFeatureStoreEntityTypeIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Entitytype *string `pulumi:"entitytype"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore *string `pulumi:"featurestore"`
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
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreEntityTypeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type AiFeatureStoreEntityTypeIamMemberState struct {
	Condition AiFeatureStoreEntityTypeIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Entitytype pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringPtrInput
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
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreEntityTypeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (AiFeatureStoreEntityTypeIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeIamMemberState)(nil)).Elem()
}

type aiFeatureStoreEntityTypeIamMemberArgs struct {
	Condition *AiFeatureStoreEntityTypeIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Entitytype string `pulumi:"entitytype"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore string `pulumi:"featurestore"`
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
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreEntityTypeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AiFeatureStoreEntityTypeIamMember resource.
type AiFeatureStoreEntityTypeIamMemberArgs struct {
	Condition AiFeatureStoreEntityTypeIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Entitytype pulumi.StringInput
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringInput
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
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreEntityTypeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (AiFeatureStoreEntityTypeIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeIamMemberArgs)(nil)).Elem()
}

type AiFeatureStoreEntityTypeIamMemberInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeIamMemberOutput() AiFeatureStoreEntityTypeIamMemberOutput
	ToAiFeatureStoreEntityTypeIamMemberOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamMemberOutput
}

func (*AiFeatureStoreEntityTypeIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityTypeIamMember)(nil)).Elem()
}

func (i *AiFeatureStoreEntityTypeIamMember) ToAiFeatureStoreEntityTypeIamMemberOutput() AiFeatureStoreEntityTypeIamMemberOutput {
	return i.ToAiFeatureStoreEntityTypeIamMemberOutputWithContext(context.Background())
}

func (i *AiFeatureStoreEntityTypeIamMember) ToAiFeatureStoreEntityTypeIamMemberOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeIamMemberOutput)
}

// AiFeatureStoreEntityTypeIamMemberArrayInput is an input type that accepts AiFeatureStoreEntityTypeIamMemberArray and AiFeatureStoreEntityTypeIamMemberArrayOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeIamMemberArrayInput` via:
//
//	AiFeatureStoreEntityTypeIamMemberArray{ AiFeatureStoreEntityTypeIamMemberArgs{...} }
type AiFeatureStoreEntityTypeIamMemberArrayInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeIamMemberArrayOutput() AiFeatureStoreEntityTypeIamMemberArrayOutput
	ToAiFeatureStoreEntityTypeIamMemberArrayOutputWithContext(context.Context) AiFeatureStoreEntityTypeIamMemberArrayOutput
}

type AiFeatureStoreEntityTypeIamMemberArray []AiFeatureStoreEntityTypeIamMemberInput

func (AiFeatureStoreEntityTypeIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreEntityTypeIamMember)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeIamMemberArray) ToAiFeatureStoreEntityTypeIamMemberArrayOutput() AiFeatureStoreEntityTypeIamMemberArrayOutput {
	return i.ToAiFeatureStoreEntityTypeIamMemberArrayOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeIamMemberArray) ToAiFeatureStoreEntityTypeIamMemberArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeIamMemberArrayOutput)
}

// AiFeatureStoreEntityTypeIamMemberMapInput is an input type that accepts AiFeatureStoreEntityTypeIamMemberMap and AiFeatureStoreEntityTypeIamMemberMapOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeIamMemberMapInput` via:
//
//	AiFeatureStoreEntityTypeIamMemberMap{ "key": AiFeatureStoreEntityTypeIamMemberArgs{...} }
type AiFeatureStoreEntityTypeIamMemberMapInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeIamMemberMapOutput() AiFeatureStoreEntityTypeIamMemberMapOutput
	ToAiFeatureStoreEntityTypeIamMemberMapOutputWithContext(context.Context) AiFeatureStoreEntityTypeIamMemberMapOutput
}

type AiFeatureStoreEntityTypeIamMemberMap map[string]AiFeatureStoreEntityTypeIamMemberInput

func (AiFeatureStoreEntityTypeIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreEntityTypeIamMember)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeIamMemberMap) ToAiFeatureStoreEntityTypeIamMemberMapOutput() AiFeatureStoreEntityTypeIamMemberMapOutput {
	return i.ToAiFeatureStoreEntityTypeIamMemberMapOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeIamMemberMap) ToAiFeatureStoreEntityTypeIamMemberMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeIamMemberMapOutput)
}

type AiFeatureStoreEntityTypeIamMemberOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityTypeIamMember)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeIamMemberOutput) ToAiFeatureStoreEntityTypeIamMemberOutput() AiFeatureStoreEntityTypeIamMemberOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamMemberOutput) ToAiFeatureStoreEntityTypeIamMemberOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamMemberOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamMemberOutput) Condition() AiFeatureStoreEntityTypeIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamMember) AiFeatureStoreEntityTypeIamMemberConditionPtrOutput {
		return v.Condition
	}).(AiFeatureStoreEntityTypeIamMemberConditionPtrOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o AiFeatureStoreEntityTypeIamMemberOutput) Entitytype() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamMember) pulumi.StringOutput { return v.Entitytype }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o AiFeatureStoreEntityTypeIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
func (o AiFeatureStoreEntityTypeIamMemberOutput) Featurestore() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamMember) pulumi.StringOutput { return v.Featurestore }).(pulumi.StringOutput)
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
func (o AiFeatureStoreEntityTypeIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `vertex.AiFeatureStoreEntityTypeIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o AiFeatureStoreEntityTypeIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type AiFeatureStoreEntityTypeIamMemberArrayOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreEntityTypeIamMember)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeIamMemberArrayOutput) ToAiFeatureStoreEntityTypeIamMemberArrayOutput() AiFeatureStoreEntityTypeIamMemberArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamMemberArrayOutput) ToAiFeatureStoreEntityTypeIamMemberArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamMemberArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamMemberArrayOutput) Index(i pulumi.IntInput) AiFeatureStoreEntityTypeIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiFeatureStoreEntityTypeIamMember {
		return vs[0].([]*AiFeatureStoreEntityTypeIamMember)[vs[1].(int)]
	}).(AiFeatureStoreEntityTypeIamMemberOutput)
}

type AiFeatureStoreEntityTypeIamMemberMapOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreEntityTypeIamMember)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeIamMemberMapOutput) ToAiFeatureStoreEntityTypeIamMemberMapOutput() AiFeatureStoreEntityTypeIamMemberMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamMemberMapOutput) ToAiFeatureStoreEntityTypeIamMemberMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamMemberMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamMemberMapOutput) MapIndex(k pulumi.StringInput) AiFeatureStoreEntityTypeIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiFeatureStoreEntityTypeIamMember {
		return vs[0].(map[string]*AiFeatureStoreEntityTypeIamMember)[vs[1].(string)]
	}).(AiFeatureStoreEntityTypeIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeIamMemberInput)(nil)).Elem(), &AiFeatureStoreEntityTypeIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeIamMemberArrayInput)(nil)).Elem(), AiFeatureStoreEntityTypeIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeIamMemberMapInput)(nil)).Elem(), AiFeatureStoreEntityTypeIamMemberMap{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeIamMemberOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeIamMemberArrayOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeIamMemberMapOutput{})
}
