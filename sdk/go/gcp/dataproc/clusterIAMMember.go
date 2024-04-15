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
// <!--Start PulumiCodeChooser -->
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
//				PolicyData: pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_cluster\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//				Role:    pulumi.String("roles/editor"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_cluster\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//				Role:    pulumi.String("roles/editor"),
//				Member:  pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_cluster\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
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
//				PolicyData: pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_cluster\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
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
//				Role:    pulumi.String("roles/editor"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_dataproc\_cluster\_iam\_member
//
// <!--Start PulumiCodeChooser -->
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
//				Role:    pulumi.String("roles/editor"),
//				Member:  pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// ### Importing IAM policies
//
// IAM policy imports use the `cluster` identifier of the Dataproc Cluster resource only. For example:
//
// * `projects/{project}/regions/{region}/clusters/{cluster}`
//
// An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
//
// tf
//
// import {
//
//	id = projects/{project}/regions/{region}/clusters/{cluster}
//
//	to = google_dataproc_cluster_iam_policy.default
//
// }
//
// The `pulumi import` command can also be used:
//
// ```sh
// $ pulumi import gcp:dataproc/clusterIAMMember:ClusterIAMMember default projects/{project}/regions/{region}/clusters/{cluster}
// ```
type ClusterIAMMember struct {
	pulumi.CustomResourceState

	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	Cluster   pulumi.StringOutput                `pulumi:"cluster"`
	Condition ClusterIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringOutput `pulumi:"member"`
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

// NewClusterIAMMember registers a new resource with the given unique name, arguments, and options.
func NewClusterIAMMember(ctx *pulumi.Context,
	name string, args *ClusterIAMMemberArgs, opts ...pulumi.ResourceOption) (*ClusterIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterIAMMember
	err := ctx.RegisterResource("gcp:dataproc/clusterIAMMember:ClusterIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterIAMMember gets an existing ClusterIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterIAMMemberState, opts ...pulumi.ResourceOption) (*ClusterIAMMember, error) {
	var resource ClusterIAMMember
	err := ctx.ReadResource("gcp:dataproc/clusterIAMMember:ClusterIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterIAMMember resources.
type clusterIAMMemberState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	Cluster   *string                    `pulumi:"cluster"`
	Condition *ClusterIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag *string `pulumi:"etag"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member *string `pulumi:"member"`
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

type ClusterIAMMemberState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	Cluster   pulumi.StringPtrInput
	Condition ClusterIAMMemberConditionPtrInput
	// (Computed) The etag of the clusters's IAM policy.
	Etag pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringPtrInput
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

func (ClusterIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMMemberState)(nil)).Elem()
}

type clusterIAMMemberArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	Cluster   string                     `pulumi:"cluster"`
	Condition *ClusterIAMMemberCondition `pulumi:"condition"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member string `pulumi:"member"`
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

// The set of arguments for constructing a ClusterIAMMember resource.
type ClusterIAMMemberArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	//
	// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
	Cluster   pulumi.StringInput
	Condition ClusterIAMMemberConditionPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringInput
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

func (ClusterIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMMemberArgs)(nil)).Elem()
}

type ClusterIAMMemberInput interface {
	pulumi.Input

	ToClusterIAMMemberOutput() ClusterIAMMemberOutput
	ToClusterIAMMemberOutputWithContext(ctx context.Context) ClusterIAMMemberOutput
}

func (*ClusterIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIAMMember)(nil)).Elem()
}

func (i *ClusterIAMMember) ToClusterIAMMemberOutput() ClusterIAMMemberOutput {
	return i.ToClusterIAMMemberOutputWithContext(context.Background())
}

func (i *ClusterIAMMember) ToClusterIAMMemberOutputWithContext(ctx context.Context) ClusterIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIAMMemberOutput)
}

// ClusterIAMMemberArrayInput is an input type that accepts ClusterIAMMemberArray and ClusterIAMMemberArrayOutput values.
// You can construct a concrete instance of `ClusterIAMMemberArrayInput` via:
//
//	ClusterIAMMemberArray{ ClusterIAMMemberArgs{...} }
type ClusterIAMMemberArrayInput interface {
	pulumi.Input

	ToClusterIAMMemberArrayOutput() ClusterIAMMemberArrayOutput
	ToClusterIAMMemberArrayOutputWithContext(context.Context) ClusterIAMMemberArrayOutput
}

type ClusterIAMMemberArray []ClusterIAMMemberInput

func (ClusterIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterIAMMember)(nil)).Elem()
}

func (i ClusterIAMMemberArray) ToClusterIAMMemberArrayOutput() ClusterIAMMemberArrayOutput {
	return i.ToClusterIAMMemberArrayOutputWithContext(context.Background())
}

func (i ClusterIAMMemberArray) ToClusterIAMMemberArrayOutputWithContext(ctx context.Context) ClusterIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIAMMemberArrayOutput)
}

// ClusterIAMMemberMapInput is an input type that accepts ClusterIAMMemberMap and ClusterIAMMemberMapOutput values.
// You can construct a concrete instance of `ClusterIAMMemberMapInput` via:
//
//	ClusterIAMMemberMap{ "key": ClusterIAMMemberArgs{...} }
type ClusterIAMMemberMapInput interface {
	pulumi.Input

	ToClusterIAMMemberMapOutput() ClusterIAMMemberMapOutput
	ToClusterIAMMemberMapOutputWithContext(context.Context) ClusterIAMMemberMapOutput
}

type ClusterIAMMemberMap map[string]ClusterIAMMemberInput

func (ClusterIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterIAMMember)(nil)).Elem()
}

func (i ClusterIAMMemberMap) ToClusterIAMMemberMapOutput() ClusterIAMMemberMapOutput {
	return i.ToClusterIAMMemberMapOutputWithContext(context.Background())
}

func (i ClusterIAMMemberMap) ToClusterIAMMemberMapOutputWithContext(ctx context.Context) ClusterIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIAMMemberMapOutput)
}

type ClusterIAMMemberOutput struct{ *pulumi.OutputState }

func (ClusterIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIAMMember)(nil)).Elem()
}

func (o ClusterIAMMemberOutput) ToClusterIAMMemberOutput() ClusterIAMMemberOutput {
	return o
}

func (o ClusterIAMMemberOutput) ToClusterIAMMemberOutputWithContext(ctx context.Context) ClusterIAMMemberOutput {
	return o
}

// The name or relative resource id of the cluster to manage IAM policies for.
//
// For `dataproc.ClusterIAMMember` or `dataproc.ClusterIAMBinding`:
func (o ClusterIAMMemberOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMMember) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

func (o ClusterIAMMemberOutput) Condition() ClusterIAMMemberConditionPtrOutput {
	return o.ApplyT(func(v *ClusterIAMMember) ClusterIAMMemberConditionPtrOutput { return v.Condition }).(ClusterIAMMemberConditionPtrOutput)
}

// (Computed) The etag of the clusters's IAM policy.
func (o ClusterIAMMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in `role`.
// Each entry can have one of the following values:
// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o ClusterIAMMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The project in which the cluster belongs. If it
// is not provided, the provider will use a default.
func (o ClusterIAMMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region in which the cluster belongs. If it
// is not provided, the provider will use a default.
func (o ClusterIAMMemberOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMMember) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
//
// `dataproc.ClusterIAMPolicy` only:
func (o ClusterIAMMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ClusterIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (ClusterIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterIAMMember)(nil)).Elem()
}

func (o ClusterIAMMemberArrayOutput) ToClusterIAMMemberArrayOutput() ClusterIAMMemberArrayOutput {
	return o
}

func (o ClusterIAMMemberArrayOutput) ToClusterIAMMemberArrayOutputWithContext(ctx context.Context) ClusterIAMMemberArrayOutput {
	return o
}

func (o ClusterIAMMemberArrayOutput) Index(i pulumi.IntInput) ClusterIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterIAMMember {
		return vs[0].([]*ClusterIAMMember)[vs[1].(int)]
	}).(ClusterIAMMemberOutput)
}

type ClusterIAMMemberMapOutput struct{ *pulumi.OutputState }

func (ClusterIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterIAMMember)(nil)).Elem()
}

func (o ClusterIAMMemberMapOutput) ToClusterIAMMemberMapOutput() ClusterIAMMemberMapOutput {
	return o
}

func (o ClusterIAMMemberMapOutput) ToClusterIAMMemberMapOutputWithContext(ctx context.Context) ClusterIAMMemberMapOutput {
	return o
}

func (o ClusterIAMMemberMapOutput) MapIndex(k pulumi.StringInput) ClusterIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterIAMMember {
		return vs[0].(map[string]*ClusterIAMMember)[vs[1].(string)]
	}).(ClusterIAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIAMMemberInput)(nil)).Elem(), &ClusterIAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIAMMemberArrayInput)(nil)).Elem(), ClusterIAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIAMMemberMapInput)(nil)).Elem(), ClusterIAMMemberMap{})
	pulumi.RegisterOutputType(ClusterIAMMemberOutput{})
	pulumi.RegisterOutputType(ClusterIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(ClusterIAMMemberMapOutput{})
}
