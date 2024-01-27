// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
//
// * `dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
// * `dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
// * `dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
//
// > **Note:** `dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `dataproc.ClusterIAMBinding` and `dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `dataproc.ClusterIAMPolicy` replaces the entire policy.
//
// > **Note:** `dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_dataproc\_cluster\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
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
//			_, err = dataproc.NewClusterIAMPolicy(ctx, "editor", &dataproc.ClusterIAMPolicyArgs{
//				Project:    pulumi.String("your-project"),
//				Region:     pulumi.String("your-region"),
//				Cluster:    pulumi.String("your-dataproc-cluster"),
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
// ## google\_dataproc\_cluster\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewClusterIAMBinding(ctx, "editor", &dataproc.ClusterIAMBindingArgs{
//				Cluster: pulumi.String("your-dataproc-cluster"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role: pulumi.String("roles/editor"),
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
// ## google\_dataproc\_cluster\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataproc.NewClusterIAMMember(ctx, "editor", &dataproc.ClusterIAMMemberArgs{
//				Cluster: pulumi.String("your-dataproc-cluster"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Role:    pulumi.String("roles/editor"),
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
// ### Importing IAM policies IAM policy imports use the `cluster` identifier of the Dataproc Cluster resource only. For example* `projects/{project}/regions/{region}/clusters/{cluster}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
//
//	id = projects/{project}/regions/{region}/clusters/{cluster}
//
//	to = google_dataproc_cluster_iam_policy.default } The `pulumi import` command can also be used
//
// ```sh
//
//	$ pulumi import gcp:dataproc/clusterIAMBinding:ClusterIAMBinding default projects/{project}/regions/{region}/clusters/{cluster}
//
// ```
type ClusterIAMBinding struct {
	pulumi.CustomResourceState

	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Cluster   pulumi.StringOutput                 `pulumi:"cluster"`
	Condition ClusterIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	//
	// `dataproc.ClusterIAMPolicy` only:
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewClusterIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewClusterIAMBinding(ctx *pulumi.Context,
	name string, args *ClusterIAMBindingArgs, opts ...pulumi.ResourceOption) (*ClusterIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterIAMBinding
	err := ctx.RegisterResource("gcp:dataproc/clusterIAMBinding:ClusterIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterIAMBinding gets an existing ClusterIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterIAMBindingState, opts ...pulumi.ResourceOption) (*ClusterIAMBinding, error) {
	var resource ClusterIAMBinding
	err := ctx.ReadResource("gcp:dataproc/clusterIAMBinding:ClusterIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterIAMBinding resources.
type clusterIAMBindingState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Cluster   *string                     `pulumi:"cluster"`
	Condition *ClusterIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	//
	// `dataproc.ClusterIAMPolicy` only:
	Role *string `pulumi:"role"`
}

type ClusterIAMBindingState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Cluster   pulumi.StringPtrInput
	Condition ClusterIAMBindingConditionPtrInput
	// (Computed) The etag of the clusters's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	//
	// `dataproc.ClusterIAMPolicy` only:
	Role pulumi.StringPtrInput
}

func (ClusterIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMBindingState)(nil)).Elem()
}

type clusterIAMBindingArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Cluster   string                      `pulumi:"cluster"`
	Condition *ClusterIAMBindingCondition `pulumi:"condition"`
	Members   []string                    `pulumi:"members"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	//
	// `dataproc.ClusterIAMPolicy` only:
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ClusterIAMBinding resource.
type ClusterIAMBindingArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Cluster   pulumi.StringInput
	Condition ClusterIAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	//
	// `dataproc.ClusterIAMPolicy` only:
	Role pulumi.StringInput
}

func (ClusterIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMBindingArgs)(nil)).Elem()
}

type ClusterIAMBindingInput interface {
	pulumi.Input

	ToClusterIAMBindingOutput() ClusterIAMBindingOutput
	ToClusterIAMBindingOutputWithContext(ctx context.Context) ClusterIAMBindingOutput
}

func (*ClusterIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIAMBinding)(nil)).Elem()
}

func (i *ClusterIAMBinding) ToClusterIAMBindingOutput() ClusterIAMBindingOutput {
	return i.ToClusterIAMBindingOutputWithContext(context.Background())
}

func (i *ClusterIAMBinding) ToClusterIAMBindingOutputWithContext(ctx context.Context) ClusterIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIAMBindingOutput)
}

// ClusterIAMBindingArrayInput is an input type that accepts ClusterIAMBindingArray and ClusterIAMBindingArrayOutput values.
// You can construct a concrete instance of `ClusterIAMBindingArrayInput` via:
//
//	ClusterIAMBindingArray{ ClusterIAMBindingArgs{...} }
type ClusterIAMBindingArrayInput interface {
	pulumi.Input

	ToClusterIAMBindingArrayOutput() ClusterIAMBindingArrayOutput
	ToClusterIAMBindingArrayOutputWithContext(context.Context) ClusterIAMBindingArrayOutput
}

type ClusterIAMBindingArray []ClusterIAMBindingInput

func (ClusterIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterIAMBinding)(nil)).Elem()
}

func (i ClusterIAMBindingArray) ToClusterIAMBindingArrayOutput() ClusterIAMBindingArrayOutput {
	return i.ToClusterIAMBindingArrayOutputWithContext(context.Background())
}

func (i ClusterIAMBindingArray) ToClusterIAMBindingArrayOutputWithContext(ctx context.Context) ClusterIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIAMBindingArrayOutput)
}

// ClusterIAMBindingMapInput is an input type that accepts ClusterIAMBindingMap and ClusterIAMBindingMapOutput values.
// You can construct a concrete instance of `ClusterIAMBindingMapInput` via:
//
//	ClusterIAMBindingMap{ "key": ClusterIAMBindingArgs{...} }
type ClusterIAMBindingMapInput interface {
	pulumi.Input

	ToClusterIAMBindingMapOutput() ClusterIAMBindingMapOutput
	ToClusterIAMBindingMapOutputWithContext(context.Context) ClusterIAMBindingMapOutput
}

type ClusterIAMBindingMap map[string]ClusterIAMBindingInput

func (ClusterIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterIAMBinding)(nil)).Elem()
}

func (i ClusterIAMBindingMap) ToClusterIAMBindingMapOutput() ClusterIAMBindingMapOutput {
	return i.ToClusterIAMBindingMapOutputWithContext(context.Background())
}

func (i ClusterIAMBindingMap) ToClusterIAMBindingMapOutputWithContext(ctx context.Context) ClusterIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIAMBindingMapOutput)
}

type ClusterIAMBindingOutput struct{ *pulumi.OutputState }

func (ClusterIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIAMBinding)(nil)).Elem()
}

func (o ClusterIAMBindingOutput) ToClusterIAMBindingOutput() ClusterIAMBindingOutput {
	return o
}

func (o ClusterIAMBindingOutput) ToClusterIAMBindingOutputWithContext(ctx context.Context) ClusterIAMBindingOutput {
	return o
}

// The name or relative resource id of the cluster to manage IAM policies for.
//
// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o ClusterIAMBindingOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

func (o ClusterIAMBindingOutput) Condition() ClusterIAMBindingConditionPtrOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) ClusterIAMBindingConditionPtrOutput { return v.Condition }).(ClusterIAMBindingConditionPtrOutput)
}

// (Computed) The etag of the clusters's IAM policy.
func (o ClusterIAMBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ClusterIAMBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The project in which the cluster belongs. If it
// is not provided, the provider will use a default.
func (o ClusterIAMBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region in which the cluster belongs. If it
// is not provided, the provider will use a default.
func (o ClusterIAMBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
//
// `dataproc.ClusterIAMPolicy` only:
func (o ClusterIAMBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ClusterIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (ClusterIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterIAMBinding)(nil)).Elem()
}

func (o ClusterIAMBindingArrayOutput) ToClusterIAMBindingArrayOutput() ClusterIAMBindingArrayOutput {
	return o
}

func (o ClusterIAMBindingArrayOutput) ToClusterIAMBindingArrayOutputWithContext(ctx context.Context) ClusterIAMBindingArrayOutput {
	return o
}

func (o ClusterIAMBindingArrayOutput) Index(i pulumi.IntInput) ClusterIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterIAMBinding {
		return vs[0].([]*ClusterIAMBinding)[vs[1].(int)]
	}).(ClusterIAMBindingOutput)
}

type ClusterIAMBindingMapOutput struct{ *pulumi.OutputState }

func (ClusterIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterIAMBinding)(nil)).Elem()
}

func (o ClusterIAMBindingMapOutput) ToClusterIAMBindingMapOutput() ClusterIAMBindingMapOutput {
	return o
}

func (o ClusterIAMBindingMapOutput) ToClusterIAMBindingMapOutputWithContext(ctx context.Context) ClusterIAMBindingMapOutput {
	return o
}

func (o ClusterIAMBindingMapOutput) MapIndex(k pulumi.StringInput) ClusterIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterIAMBinding {
		return vs[0].(map[string]*ClusterIAMBinding)[vs[1].(string)]
	}).(ClusterIAMBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIAMBindingInput)(nil)).Elem(), &ClusterIAMBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIAMBindingArrayInput)(nil)).Elem(), ClusterIAMBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIAMBindingMapInput)(nil)).Elem(), ClusterIAMBindingMap{})
	pulumi.RegisterOutputType(ClusterIAMBindingOutput{})
	pulumi.RegisterOutputType(ClusterIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(ClusterIAMBindingMapOutput{})
}
