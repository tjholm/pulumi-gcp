// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
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
// Cluster IAM resources can be imported using the project, region, cluster name, role and/or member.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/clusterIAMBinding:ClusterIAMBinding editor "projects/{project}/regions/{region}/clusters/{cluster}"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataproc/clusterIAMBinding:ClusterIAMBinding editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataproc/clusterIAMBinding:ClusterIAMBinding editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ClusterIAMBinding struct {
	pulumi.CustomResourceState

	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   pulumi.StringOutput                 `pulumi:"cluster"`
	Condition ClusterIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	Project pulumi.StringOutput      `pulumi:"project"`
	Region  pulumi.StringOutput      `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
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
	Cluster   *string                     `pulumi:"cluster"`
	Condition *ClusterIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	Project *string  `pulumi:"project"`
	Region  *string  `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ClusterIAMBindingState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   pulumi.StringPtrInput
	Condition ClusterIAMBindingConditionPtrInput
	// (Computed) The etag of the clusters's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ClusterIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMBindingState)(nil)).Elem()
}

type clusterIAMBindingArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   string                      `pulumi:"cluster"`
	Condition *ClusterIAMBindingCondition `pulumi:"condition"`
	Members   []string                    `pulumi:"members"`
	Project   *string                     `pulumi:"project"`
	Region    *string                     `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ClusterIAMBinding resource.
type ClusterIAMBindingArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   pulumi.StringInput
	Condition ClusterIAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	Project   pulumi.StringPtrInput
	Region    pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
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

func (o ClusterIAMBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ClusterIAMBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterIAMBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
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
