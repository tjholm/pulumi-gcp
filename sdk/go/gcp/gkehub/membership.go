// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Membership contains information about a member cluster.
//
// To get more information about Membership, see:
//
// * [API documentation](https://cloud.google.com/anthos/multicluster-management/reference/rest/v1/projects.locations.memberships)
// * How-to Guides
//   - [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
//
// ## Example Usage
// ### Gkehub Membership Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
//				DeletionProtection: pulumi.Bool(true),
//				InitialNodeCount:   pulumi.Int(1),
//				Location:           pulumi.String("us-central1-a"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gkehub.NewMembership(ctx, "membership", &gkehub.MembershipArgs{
//				Endpoint: &gkehub.MembershipEndpointArgs{
//					GkeCluster: &gkehub.MembershipEndpointGkeClusterArgs{
//						ResourceLink: primary.ID().ApplyT(func(id string) (string, error) {
//							return fmt.Sprintf("//container.googleapis.com/%v", id), nil
//						}).(pulumi.StringOutput),
//					},
//				},
//				Labels: pulumi.StringMap{
//					"env": pulumi.String("test"),
//				},
//				MembershipId: pulumi.String("basic"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Gkehub Membership Issuer
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1-a"),
//				InitialNodeCount: pulumi.Int(1),
//				WorkloadIdentityConfig: &container.ClusterWorkloadIdentityConfigArgs{
//					WorkloadPool: pulumi.String("my-project-name.svc.id.goog"),
//				},
//				DeletionProtection: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gkehub.NewMembership(ctx, "membership", &gkehub.MembershipArgs{
//				MembershipId: pulumi.String("basic"),
//				Endpoint: &gkehub.MembershipEndpointArgs{
//					GkeCluster: &gkehub.MembershipEndpointGkeClusterArgs{
//						ResourceLink: primary.ID(),
//					},
//				},
//				Authority: &gkehub.MembershipAuthorityArgs{
//					Issuer: primary.ID().ApplyT(func(id string) (string, error) {
//						return fmt.Sprintf("https://container.googleapis.com/v1/%v", id), nil
//					}).(pulumi.StringOutput),
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
// ## Import
//
// # Membership can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:gkehub/membership:Membership default projects/{{project}}/locations/global/memberships/{{membership_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkehub/membership:Membership default {{project}}/{{membership_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkehub/membership:Membership default {{membership_id}}
//
// ```
type Membership struct {
	pulumi.CustomResourceState

	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	Authority MembershipAuthorityPtrOutput `pulumi:"authority"`
	// (Optional, Deprecated)
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// > **Warning:** `description` is deprecated and will be removed in a future major release.
	//
	// Deprecated: `description` is deprecated and will be removed in a future major release.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint MembershipEndpointPtrOutput `pulumi:"endpoint"`
	// Labels to apply to this membership.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The client-provided identifier of the membership.
	//
	// ***
	MembershipId pulumi.StringOutput `pulumi:"membershipId"`
	// The unique identifier of the membership.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
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
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// (Optional, Deprecated)
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// > **Warning:** `description` is deprecated and will be removed in a future major release.
	//
	// Deprecated: `description` is deprecated and will be removed in a future major release.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint *MembershipEndpoint `pulumi:"endpoint"`
	// Labels to apply to this membership.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The client-provided identifier of the membership.
	//
	// ***
	MembershipId *string `pulumi:"membershipId"`
	// The unique identifier of the membership.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
}

type MembershipState struct {
	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	Authority MembershipAuthorityPtrInput
	// (Optional, Deprecated)
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// > **Warning:** `description` is deprecated and will be removed in a future major release.
	//
	// Deprecated: `description` is deprecated and will be removed in a future major release.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint MembershipEndpointPtrInput
	// Labels to apply to this membership.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The client-provided identifier of the membership.
	//
	// ***
	MembershipId pulumi.StringPtrInput
	// The unique identifier of the membership.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
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
	// (Optional, Deprecated)
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// > **Warning:** `description` is deprecated and will be removed in a future major release.
	//
	// Deprecated: `description` is deprecated and will be removed in a future major release.
	Description *string `pulumi:"description"`
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint *MembershipEndpoint `pulumi:"endpoint"`
	// Labels to apply to this membership.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The client-provided identifier of the membership.
	//
	// ***
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
	// (Optional, Deprecated)
	// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
	//
	// > **Warning:** `description` is deprecated and will be removed in a future major release.
	//
	// Deprecated: `description` is deprecated and will be removed in a future major release.
	Description pulumi.StringPtrInput
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	Endpoint MembershipEndpointPtrInput
	// Labels to apply to this membership.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The client-provided identifier of the membership.
	//
	// ***
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
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (i *Membership) ToMembershipOutput() MembershipOutput {
	return i.ToMembershipOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipOutput)
}

func (i *Membership) ToOutput(ctx context.Context) pulumix.Output[*Membership] {
	return pulumix.Output[*Membership]{
		OutputState: i.ToMembershipOutputWithContext(ctx).OutputState,
	}
}

// MembershipArrayInput is an input type that accepts MembershipArray and MembershipArrayOutput values.
// You can construct a concrete instance of `MembershipArrayInput` via:
//
//	MembershipArray{ MembershipArgs{...} }
type MembershipArrayInput interface {
	pulumi.Input

	ToMembershipArrayOutput() MembershipArrayOutput
	ToMembershipArrayOutputWithContext(context.Context) MembershipArrayOutput
}

type MembershipArray []MembershipInput

func (MembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Membership)(nil)).Elem()
}

func (i MembershipArray) ToMembershipArrayOutput() MembershipArrayOutput {
	return i.ToMembershipArrayOutputWithContext(context.Background())
}

func (i MembershipArray) ToMembershipArrayOutputWithContext(ctx context.Context) MembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipArrayOutput)
}

func (i MembershipArray) ToOutput(ctx context.Context) pulumix.Output[[]*Membership] {
	return pulumix.Output[[]*Membership]{
		OutputState: i.ToMembershipArrayOutputWithContext(ctx).OutputState,
	}
}

// MembershipMapInput is an input type that accepts MembershipMap and MembershipMapOutput values.
// You can construct a concrete instance of `MembershipMapInput` via:
//
//	MembershipMap{ "key": MembershipArgs{...} }
type MembershipMapInput interface {
	pulumi.Input

	ToMembershipMapOutput() MembershipMapOutput
	ToMembershipMapOutputWithContext(context.Context) MembershipMapOutput
}

type MembershipMap map[string]MembershipInput

func (MembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Membership)(nil)).Elem()
}

func (i MembershipMap) ToMembershipMapOutput() MembershipMapOutput {
	return i.ToMembershipMapOutputWithContext(context.Background())
}

func (i MembershipMap) ToMembershipMapOutputWithContext(ctx context.Context) MembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipMapOutput)
}

func (i MembershipMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Membership] {
	return pulumix.Output[map[string]*Membership]{
		OutputState: i.ToMembershipMapOutputWithContext(ctx).OutputState,
	}
}

type MembershipOutput struct{ *pulumi.OutputState }

func (MembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (o MembershipOutput) ToMembershipOutput() MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return o
}

func (o MembershipOutput) ToOutput(ctx context.Context) pulumix.Output[*Membership] {
	return pulumix.Output[*Membership]{
		OutputState: o.OutputState,
	}
}

// Authority encodes how Google will recognize identities from this Membership.
// See the workload identity documentation for more details:
// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
// Structure is documented below.
func (o MembershipOutput) Authority() MembershipAuthorityPtrOutput {
	return o.ApplyT(func(v *Membership) MembershipAuthorityPtrOutput { return v.Authority }).(MembershipAuthorityPtrOutput)
}

// (Optional, Deprecated)
// The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
//
// > **Warning:** `description` is deprecated and will be removed in a future major release.
//
// Deprecated: `description` is deprecated and will be removed in a future major release.
func (o MembershipOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o MembershipOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
// Structure is documented below.
func (o MembershipOutput) Endpoint() MembershipEndpointPtrOutput {
	return o.ApplyT(func(v *Membership) MembershipEndpointPtrOutput { return v.Endpoint }).(MembershipEndpointPtrOutput)
}

// Labels to apply to this membership.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o MembershipOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The client-provided identifier of the membership.
//
// ***
func (o MembershipOutput) MembershipId() pulumi.StringOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringOutput { return v.MembershipId }).(pulumi.StringOutput)
}

// The unique identifier of the membership.
func (o MembershipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o MembershipOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o MembershipOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

type MembershipArrayOutput struct{ *pulumi.OutputState }

func (MembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Membership)(nil)).Elem()
}

func (o MembershipArrayOutput) ToMembershipArrayOutput() MembershipArrayOutput {
	return o
}

func (o MembershipArrayOutput) ToMembershipArrayOutputWithContext(ctx context.Context) MembershipArrayOutput {
	return o
}

func (o MembershipArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Membership] {
	return pulumix.Output[[]*Membership]{
		OutputState: o.OutputState,
	}
}

func (o MembershipArrayOutput) Index(i pulumi.IntInput) MembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Membership {
		return vs[0].([]*Membership)[vs[1].(int)]
	}).(MembershipOutput)
}

type MembershipMapOutput struct{ *pulumi.OutputState }

func (MembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Membership)(nil)).Elem()
}

func (o MembershipMapOutput) ToMembershipMapOutput() MembershipMapOutput {
	return o
}

func (o MembershipMapOutput) ToMembershipMapOutputWithContext(ctx context.Context) MembershipMapOutput {
	return o
}

func (o MembershipMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Membership] {
	return pulumix.Output[map[string]*Membership]{
		OutputState: o.OutputState,
	}
}

func (o MembershipMapOutput) MapIndex(k pulumi.StringInput) MembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Membership {
		return vs[0].(map[string]*Membership)[vs[1].(string)]
	}).(MembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipInput)(nil)).Elem(), &Membership{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipArrayInput)(nil)).Elem(), MembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipMapInput)(nil)).Elem(), MembershipMap{})
	pulumi.RegisterOutputType(MembershipOutput{})
	pulumi.RegisterOutputType(MembershipArrayOutput{})
	pulumi.RegisterOutputType(MembershipMapOutput{})
}
