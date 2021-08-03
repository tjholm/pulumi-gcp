// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Membership contains information about a member cluster.
//
// To get more information about Membership, see:
//
// * [API documentation](https://cloud.google.com/anthos/multicluster-management/reference/rest/v1/projects.locations.memberships)
// * How-to Guides
//     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
//
// ## Example Usage
// ### Gkehub Membership Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/gkehub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
// 			InitialNodeCount: pulumi.Int(1),
// 			Location:         pulumi.String("us-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gkehub.NewMembership(ctx, "membership", &gkehub.MembershipArgs{
// 			Endpoint: &gkehub.MembershipEndpointArgs{
// 				GkeCluster: &gkehub.MembershipEndpointGkeClusterArgs{
// 					ResourceLink: primary.ID().ApplyT(func(id string) (string, error) {
// 						return fmt.Sprintf("%v%v", "//container.googleapis.com/", id), nil
// 					}).(pulumi.StringOutput),
// 				},
// 			},
// 			MembershipId: pulumi.String("basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Gkehub Membership Issuer
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/gkehub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
// 			InitialNodeCount: pulumi.Int(1),
// 			Location:         pulumi.String("us-central1-a"),
// 			WorkloadIdentityConfig: &container.ClusterWorkloadIdentityConfigArgs{
// 				IdentityNamespace: pulumi.String("my-project-name.svc.id.goog"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gkehub.NewMembership(ctx, "membership", &gkehub.MembershipArgs{
// 			Authority: &gkehub.MembershipAuthorityArgs{
// 				Issuer: primary.ID().ApplyT(func(id string) (string, error) {
// 					return fmt.Sprintf("%v%v", "https://container.googleapis.com/v1/", id), nil
// 				}).(pulumi.StringOutput),
// 			},
// 			Endpoint: &gkehub.MembershipEndpointArgs{
// 				GkeCluster: &gkehub.MembershipEndpointGkeClusterArgs{
// 					ResourceLink: primary.ID().ApplyT(func(id string) (string, error) {
// 						return fmt.Sprintf("%v%v", "//container.googleapis.com/", id), nil
// 					}).(pulumi.StringOutput),
// 				},
// 			},
// 			MembershipId: pulumi.String("basic"),
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
// Membership can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:gkehub/membership:Membership default {{name}}
// ```
type Membership struct {
	pulumi.CustomResourceState

	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	Authority MembershipAuthorityPtrOutput `pulumi:"authority"`
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// Deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint MembershipEndpointPtrOutput `pulumi:"endpoint"`
	// Labels to apply to this membership.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The client-provided identifier of the membership.
	MembershipId pulumi.StringOutput `pulumi:"membershipId"`
	// The unique identifier of the membership.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewMembership registers a new resource with the given unique name, arguments, and options.
func NewMembership(ctx *pulumi.Context,
	name string, args *MembershipArgs, opts ...pulumi.ResourceOption) (*Membership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MembershipId == nil {
		return nil, errors.New("invalid value for required argument 'MembershipId'")
	}
	var resource Membership
	err := ctx.RegisterResource("gcp:gkehub/membership:Membership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembership gets an existing Membership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipState, opts ...pulumi.ResourceOption) (*Membership, error) {
	var resource Membership
	err := ctx.ReadResource("gcp:gkehub/membership:Membership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Membership resources.
type membershipState struct {
	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	Authority *MembershipAuthority `pulumi:"authority"`
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// Deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.
	Description *string `pulumi:"description"`
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint *MembershipEndpoint `pulumi:"endpoint"`
	// Labels to apply to this membership.
	Labels map[string]string `pulumi:"labels"`
	// The client-provided identifier of the membership.
	MembershipId *string `pulumi:"membershipId"`
	// The unique identifier of the membership.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type MembershipState struct {
	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	Authority MembershipAuthorityPtrInput
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// Deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.
	Description pulumi.StringPtrInput
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint MembershipEndpointPtrInput
	// Labels to apply to this membership.
	Labels pulumi.StringMapInput
	// The client-provided identifier of the membership.
	MembershipId pulumi.StringPtrInput
	// The unique identifier of the membership.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipState)(nil)).Elem()
}

type membershipArgs struct {
	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	Authority *MembershipAuthority `pulumi:"authority"`
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// Deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.
	Description *string `pulumi:"description"`
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint *MembershipEndpoint `pulumi:"endpoint"`
	// Labels to apply to this membership.
	Labels map[string]string `pulumi:"labels"`
	// The client-provided identifier of the membership.
	MembershipId string `pulumi:"membershipId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Membership resource.
type MembershipArgs struct {
	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	Authority MembershipAuthorityPtrInput
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// Deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.
	Description pulumi.StringPtrInput
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint MembershipEndpointPtrInput
	// Labels to apply to this membership.
	Labels pulumi.StringMapInput
	// The client-provided identifier of the membership.
	MembershipId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipArgs)(nil)).Elem()
}

type MembershipInput interface {
	pulumi.Input

	ToMembershipOutput() MembershipOutput
	ToMembershipOutputWithContext(ctx context.Context) MembershipOutput
}

func (*Membership) ElementType() reflect.Type {
	return reflect.TypeOf((*Membership)(nil))
}

func (i *Membership) ToMembershipOutput() MembershipOutput {
	return i.ToMembershipOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipOutput)
}

func (i *Membership) ToMembershipPtrOutput() MembershipPtrOutput {
	return i.ToMembershipPtrOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipPtrOutputWithContext(ctx context.Context) MembershipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipPtrOutput)
}

type MembershipPtrInput interface {
	pulumi.Input

	ToMembershipPtrOutput() MembershipPtrOutput
	ToMembershipPtrOutputWithContext(ctx context.Context) MembershipPtrOutput
}

type membershipPtrType MembershipArgs

func (*membershipPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil))
}

func (i *membershipPtrType) ToMembershipPtrOutput() MembershipPtrOutput {
	return i.ToMembershipPtrOutputWithContext(context.Background())
}

func (i *membershipPtrType) ToMembershipPtrOutputWithContext(ctx context.Context) MembershipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipPtrOutput)
}

// MembershipArrayInput is an input type that accepts MembershipArray and MembershipArrayOutput values.
// You can construct a concrete instance of `MembershipArrayInput` via:
//
//          MembershipArray{ MembershipArgs{...} }
type MembershipArrayInput interface {
	pulumi.Input

	ToMembershipArrayOutput() MembershipArrayOutput
	ToMembershipArrayOutputWithContext(context.Context) MembershipArrayOutput
}

type MembershipArray []MembershipInput

func (MembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Membership)(nil))
}

func (i MembershipArray) ToMembershipArrayOutput() MembershipArrayOutput {
	return i.ToMembershipArrayOutputWithContext(context.Background())
}

func (i MembershipArray) ToMembershipArrayOutputWithContext(ctx context.Context) MembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipArrayOutput)
}

// MembershipMapInput is an input type that accepts MembershipMap and MembershipMapOutput values.
// You can construct a concrete instance of `MembershipMapInput` via:
//
//          MembershipMap{ "key": MembershipArgs{...} }
type MembershipMapInput interface {
	pulumi.Input

	ToMembershipMapOutput() MembershipMapOutput
	ToMembershipMapOutputWithContext(context.Context) MembershipMapOutput
}

type MembershipMap map[string]MembershipInput

func (MembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Membership)(nil))
}

func (i MembershipMap) ToMembershipMapOutput() MembershipMapOutput {
	return i.ToMembershipMapOutputWithContext(context.Background())
}

func (i MembershipMap) ToMembershipMapOutputWithContext(ctx context.Context) MembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipMapOutput)
}

type MembershipOutput struct {
	*pulumi.OutputState
}

func (MembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Membership)(nil))
}

func (o MembershipOutput) ToMembershipOutput() MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipPtrOutput() MembershipPtrOutput {
	return o.ToMembershipPtrOutputWithContext(context.Background())
}

func (o MembershipOutput) ToMembershipPtrOutputWithContext(ctx context.Context) MembershipPtrOutput {
	return o.ApplyT(func(v Membership) *Membership {
		return &v
	}).(MembershipPtrOutput)
}

type MembershipPtrOutput struct {
	*pulumi.OutputState
}

func (MembershipPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil))
}

func (o MembershipPtrOutput) ToMembershipPtrOutput() MembershipPtrOutput {
	return o
}

func (o MembershipPtrOutput) ToMembershipPtrOutputWithContext(ctx context.Context) MembershipPtrOutput {
	return o
}

type MembershipArrayOutput struct{ *pulumi.OutputState }

func (MembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Membership)(nil))
}

func (o MembershipArrayOutput) ToMembershipArrayOutput() MembershipArrayOutput {
	return o
}

func (o MembershipArrayOutput) ToMembershipArrayOutputWithContext(ctx context.Context) MembershipArrayOutput {
	return o
}

func (o MembershipArrayOutput) Index(i pulumi.IntInput) MembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Membership {
		return vs[0].([]Membership)[vs[1].(int)]
	}).(MembershipOutput)
}

type MembershipMapOutput struct{ *pulumi.OutputState }

func (MembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Membership)(nil))
}

func (o MembershipMapOutput) ToMembershipMapOutput() MembershipMapOutput {
	return o
}

func (o MembershipMapOutput) ToMembershipMapOutputWithContext(ctx context.Context) MembershipMapOutput {
	return o
}

func (o MembershipMapOutput) MapIndex(k pulumi.StringInput) MembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Membership {
		return vs[0].(map[string]Membership)[vs[1].(string)]
	}).(MembershipOutput)
}

func init() {
	pulumi.RegisterOutputType(MembershipOutput{})
	pulumi.RegisterOutputType(MembershipPtrOutput{})
	pulumi.RegisterOutputType(MembershipArrayOutput{})
	pulumi.RegisterOutputType(MembershipMapOutput{})
}
