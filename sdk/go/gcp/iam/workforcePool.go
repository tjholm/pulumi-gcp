// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a collection of external workforces. Provides namespaces for
// federated users that can be referenced in IAM policies.
//
// To get more information about WorkforcePool, see:
//
// * [API documentation](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools)
// * How-to Guides
//   - [Manage pools](https://cloud.google.com/iam/docs/manage-workforce-identity-pools-providers#manage_pools)
//
// > **Note:** Ask your Google Cloud account team to request access to workforce identity federation for
// your billing/quota project. The account team notifies you when the project is granted access.
//
// ## Example Usage
// ### Iam Workforce Pool Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewWorkforcePool(ctx, "example", &iam.WorkforcePoolArgs{
//				Location:        pulumi.String("global"),
//				Parent:          pulumi.String("organizations/123456789"),
//				WorkforcePoolId: pulumi.String("example-pool"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Iam Workforce Pool Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewWorkforcePool(ctx, "example", &iam.WorkforcePoolArgs{
//				Description:     pulumi.String("A sample workforce pool."),
//				Disabled:        pulumi.Bool(false),
//				DisplayName:     pulumi.String("Display name"),
//				Location:        pulumi.String("global"),
//				Parent:          pulumi.String("organizations/123456789"),
//				SessionDuration: pulumi.String("7200s"),
//				WorkforcePoolId: pulumi.String("example-pool"),
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
// # WorkforcePool can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:iam/workforcePool:WorkforcePool default locations/{{location}}/workforcePools/{{workforce_pool_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:iam/workforcePool:WorkforcePool default {{location}}/{{workforce_pool_id}}
//
// ```
type WorkforcePool struct {
	pulumi.CustomResourceState

	// A user-specified description of the pool. Cannot exceed 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
	// or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The location for the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. The resource name of the pool.
	// Format: `locations/{location}/workforcePools/{workforcePoolId}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The resource name of the parent. Format: `organizations/{org-id}`.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Duration that the Google Cloud access tokens, console sign-in sessions,
	// and `gcloud` sign-in sessions from this pool are valid.
	// Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
	// If `sessionDuration` is not configured, minted credentials have a default duration of one hour (3600s).
	// A duration in seconds with up to nine fractional digits, ending with '`s`'. Example: "`3.5s`".
	SessionDuration pulumi.StringPtrOutput `pulumi:"sessionDuration"`
	// Output only. The state of the pool.
	// * STATE_UNSPECIFIED: State unspecified.
	// * ACTIVE: The pool is active, and may be used in Google Cloud policies.
	// * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted
	//   after approximately 30 days. You can restore a soft-deleted pool using
	//   [workforcePools.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePool).
	//   You cannot reuse the ID of a soft-deleted pool until it is permanently deleted.
	//   While a pool is deleted, you cannot use it to exchange tokens, or use
	//   existing tokens to access resources. If the pool is undeleted, existing
	//   tokens grant access again.
	State pulumi.StringOutput `pulumi:"state"`
	// The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
	// digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
	// The prefix `gcp-` is reserved for use by Google, and may not be specified.
	WorkforcePoolId pulumi.StringOutput `pulumi:"workforcePoolId"`
}

// NewWorkforcePool registers a new resource with the given unique name, arguments, and options.
func NewWorkforcePool(ctx *pulumi.Context,
	name string, args *WorkforcePoolArgs, opts ...pulumi.ResourceOption) (*WorkforcePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.WorkforcePoolId == nil {
		return nil, errors.New("invalid value for required argument 'WorkforcePoolId'")
	}
	var resource WorkforcePool
	err := ctx.RegisterResource("gcp:iam/workforcePool:WorkforcePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkforcePool gets an existing WorkforcePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkforcePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkforcePoolState, opts ...pulumi.ResourceOption) (*WorkforcePool, error) {
	var resource WorkforcePool
	err := ctx.ReadResource("gcp:iam/workforcePool:WorkforcePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkforcePool resources.
type workforcePoolState struct {
	// A user-specified description of the pool. Cannot exceed 256 characters.
	Description *string `pulumi:"description"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
	// or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	Disabled *bool `pulumi:"disabled"`
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	DisplayName *string `pulumi:"displayName"`
	// The location for the resource.
	Location *string `pulumi:"location"`
	// Output only. The resource name of the pool.
	// Format: `locations/{location}/workforcePools/{workforcePoolId}`
	Name *string `pulumi:"name"`
	// Immutable. The resource name of the parent. Format: `organizations/{org-id}`.
	Parent *string `pulumi:"parent"`
	// Duration that the Google Cloud access tokens, console sign-in sessions,
	// and `gcloud` sign-in sessions from this pool are valid.
	// Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
	// If `sessionDuration` is not configured, minted credentials have a default duration of one hour (3600s).
	// A duration in seconds with up to nine fractional digits, ending with '`s`'. Example: "`3.5s`".
	SessionDuration *string `pulumi:"sessionDuration"`
	// Output only. The state of the pool.
	// * STATE_UNSPECIFIED: State unspecified.
	// * ACTIVE: The pool is active, and may be used in Google Cloud policies.
	// * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted
	//   after approximately 30 days. You can restore a soft-deleted pool using
	//   [workforcePools.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePool).
	//   You cannot reuse the ID of a soft-deleted pool until it is permanently deleted.
	//   While a pool is deleted, you cannot use it to exchange tokens, or use
	//   existing tokens to access resources. If the pool is undeleted, existing
	//   tokens grant access again.
	State *string `pulumi:"state"`
	// The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
	// digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
	// The prefix `gcp-` is reserved for use by Google, and may not be specified.
	WorkforcePoolId *string `pulumi:"workforcePoolId"`
}

type WorkforcePoolState struct {
	// A user-specified description of the pool. Cannot exceed 256 characters.
	Description pulumi.StringPtrInput
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
	// or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	Disabled pulumi.BoolPtrInput
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrInput
	// The location for the resource.
	Location pulumi.StringPtrInput
	// Output only. The resource name of the pool.
	// Format: `locations/{location}/workforcePools/{workforcePoolId}`
	Name pulumi.StringPtrInput
	// Immutable. The resource name of the parent. Format: `organizations/{org-id}`.
	Parent pulumi.StringPtrInput
	// Duration that the Google Cloud access tokens, console sign-in sessions,
	// and `gcloud` sign-in sessions from this pool are valid.
	// Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
	// If `sessionDuration` is not configured, minted credentials have a default duration of one hour (3600s).
	// A duration in seconds with up to nine fractional digits, ending with '`s`'. Example: "`3.5s`".
	SessionDuration pulumi.StringPtrInput
	// Output only. The state of the pool.
	// * STATE_UNSPECIFIED: State unspecified.
	// * ACTIVE: The pool is active, and may be used in Google Cloud policies.
	// * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted
	//   after approximately 30 days. You can restore a soft-deleted pool using
	//   [workforcePools.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePool).
	//   You cannot reuse the ID of a soft-deleted pool until it is permanently deleted.
	//   While a pool is deleted, you cannot use it to exchange tokens, or use
	//   existing tokens to access resources. If the pool is undeleted, existing
	//   tokens grant access again.
	State pulumi.StringPtrInput
	// The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
	// digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
	// The prefix `gcp-` is reserved for use by Google, and may not be specified.
	WorkforcePoolId pulumi.StringPtrInput
}

func (WorkforcePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*workforcePoolState)(nil)).Elem()
}

type workforcePoolArgs struct {
	// A user-specified description of the pool. Cannot exceed 256 characters.
	Description *string `pulumi:"description"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
	// or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	Disabled *bool `pulumi:"disabled"`
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	DisplayName *string `pulumi:"displayName"`
	// The location for the resource.
	Location string `pulumi:"location"`
	// Immutable. The resource name of the parent. Format: `organizations/{org-id}`.
	Parent string `pulumi:"parent"`
	// Duration that the Google Cloud access tokens, console sign-in sessions,
	// and `gcloud` sign-in sessions from this pool are valid.
	// Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
	// If `sessionDuration` is not configured, minted credentials have a default duration of one hour (3600s).
	// A duration in seconds with up to nine fractional digits, ending with '`s`'. Example: "`3.5s`".
	SessionDuration *string `pulumi:"sessionDuration"`
	// The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
	// digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
	// The prefix `gcp-` is reserved for use by Google, and may not be specified.
	WorkforcePoolId string `pulumi:"workforcePoolId"`
}

// The set of arguments for constructing a WorkforcePool resource.
type WorkforcePoolArgs struct {
	// A user-specified description of the pool. Cannot exceed 256 characters.
	Description pulumi.StringPtrInput
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
	// or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	Disabled pulumi.BoolPtrInput
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrInput
	// The location for the resource.
	Location pulumi.StringInput
	// Immutable. The resource name of the parent. Format: `organizations/{org-id}`.
	Parent pulumi.StringInput
	// Duration that the Google Cloud access tokens, console sign-in sessions,
	// and `gcloud` sign-in sessions from this pool are valid.
	// Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
	// If `sessionDuration` is not configured, minted credentials have a default duration of one hour (3600s).
	// A duration in seconds with up to nine fractional digits, ending with '`s`'. Example: "`3.5s`".
	SessionDuration pulumi.StringPtrInput
	// The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
	// digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
	// The prefix `gcp-` is reserved for use by Google, and may not be specified.
	WorkforcePoolId pulumi.StringInput
}

func (WorkforcePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workforcePoolArgs)(nil)).Elem()
}

type WorkforcePoolInput interface {
	pulumi.Input

	ToWorkforcePoolOutput() WorkforcePoolOutput
	ToWorkforcePoolOutputWithContext(ctx context.Context) WorkforcePoolOutput
}

func (*WorkforcePool) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkforcePool)(nil)).Elem()
}

func (i *WorkforcePool) ToWorkforcePoolOutput() WorkforcePoolOutput {
	return i.ToWorkforcePoolOutputWithContext(context.Background())
}

func (i *WorkforcePool) ToWorkforcePoolOutputWithContext(ctx context.Context) WorkforcePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkforcePoolOutput)
}

// WorkforcePoolArrayInput is an input type that accepts WorkforcePoolArray and WorkforcePoolArrayOutput values.
// You can construct a concrete instance of `WorkforcePoolArrayInput` via:
//
//	WorkforcePoolArray{ WorkforcePoolArgs{...} }
type WorkforcePoolArrayInput interface {
	pulumi.Input

	ToWorkforcePoolArrayOutput() WorkforcePoolArrayOutput
	ToWorkforcePoolArrayOutputWithContext(context.Context) WorkforcePoolArrayOutput
}

type WorkforcePoolArray []WorkforcePoolInput

func (WorkforcePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkforcePool)(nil)).Elem()
}

func (i WorkforcePoolArray) ToWorkforcePoolArrayOutput() WorkforcePoolArrayOutput {
	return i.ToWorkforcePoolArrayOutputWithContext(context.Background())
}

func (i WorkforcePoolArray) ToWorkforcePoolArrayOutputWithContext(ctx context.Context) WorkforcePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkforcePoolArrayOutput)
}

// WorkforcePoolMapInput is an input type that accepts WorkforcePoolMap and WorkforcePoolMapOutput values.
// You can construct a concrete instance of `WorkforcePoolMapInput` via:
//
//	WorkforcePoolMap{ "key": WorkforcePoolArgs{...} }
type WorkforcePoolMapInput interface {
	pulumi.Input

	ToWorkforcePoolMapOutput() WorkforcePoolMapOutput
	ToWorkforcePoolMapOutputWithContext(context.Context) WorkforcePoolMapOutput
}

type WorkforcePoolMap map[string]WorkforcePoolInput

func (WorkforcePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkforcePool)(nil)).Elem()
}

func (i WorkforcePoolMap) ToWorkforcePoolMapOutput() WorkforcePoolMapOutput {
	return i.ToWorkforcePoolMapOutputWithContext(context.Background())
}

func (i WorkforcePoolMap) ToWorkforcePoolMapOutputWithContext(ctx context.Context) WorkforcePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkforcePoolMapOutput)
}

type WorkforcePoolOutput struct{ *pulumi.OutputState }

func (WorkforcePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkforcePool)(nil)).Elem()
}

func (o WorkforcePoolOutput) ToWorkforcePoolOutput() WorkforcePoolOutput {
	return o
}

func (o WorkforcePoolOutput) ToWorkforcePoolOutputWithContext(ctx context.Context) WorkforcePoolOutput {
	return o
}

// A user-specified description of the pool. Cannot exceed 256 characters.
func (o WorkforcePoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
// or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
func (o WorkforcePoolOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
func (o WorkforcePoolOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The location for the resource.
func (o WorkforcePoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. The resource name of the pool.
// Format: `locations/{location}/workforcePools/{workforcePoolId}`
func (o WorkforcePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The resource name of the parent. Format: `organizations/{org-id}`.
func (o WorkforcePoolOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Duration that the Google Cloud access tokens, console sign-in sessions,
// and `gcloud` sign-in sessions from this pool are valid.
// Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
// If `sessionDuration` is not configured, minted credentials have a default duration of one hour (3600s).
// A duration in seconds with up to nine fractional digits, ending with '`s`'. Example: "`3.5s`".
func (o WorkforcePoolOutput) SessionDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringPtrOutput { return v.SessionDuration }).(pulumi.StringPtrOutput)
}

// Output only. The state of the pool.
//   - STATE_UNSPECIFIED: State unspecified.
//   - ACTIVE: The pool is active, and may be used in Google Cloud policies.
//   - DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted
//     after approximately 30 days. You can restore a soft-deleted pool using
//     [workforcePools.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePool).
//     You cannot reuse the ID of a soft-deleted pool until it is permanently deleted.
//     While a pool is deleted, you cannot use it to exchange tokens, or use
//     existing tokens to access resources. If the pool is undeleted, existing
//     tokens grant access again.
func (o WorkforcePoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
// digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
// The prefix `gcp-` is reserved for use by Google, and may not be specified.
func (o WorkforcePoolOutput) WorkforcePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkforcePool) pulumi.StringOutput { return v.WorkforcePoolId }).(pulumi.StringOutput)
}

type WorkforcePoolArrayOutput struct{ *pulumi.OutputState }

func (WorkforcePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkforcePool)(nil)).Elem()
}

func (o WorkforcePoolArrayOutput) ToWorkforcePoolArrayOutput() WorkforcePoolArrayOutput {
	return o
}

func (o WorkforcePoolArrayOutput) ToWorkforcePoolArrayOutputWithContext(ctx context.Context) WorkforcePoolArrayOutput {
	return o
}

func (o WorkforcePoolArrayOutput) Index(i pulumi.IntInput) WorkforcePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkforcePool {
		return vs[0].([]*WorkforcePool)[vs[1].(int)]
	}).(WorkforcePoolOutput)
}

type WorkforcePoolMapOutput struct{ *pulumi.OutputState }

func (WorkforcePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkforcePool)(nil)).Elem()
}

func (o WorkforcePoolMapOutput) ToWorkforcePoolMapOutput() WorkforcePoolMapOutput {
	return o
}

func (o WorkforcePoolMapOutput) ToWorkforcePoolMapOutputWithContext(ctx context.Context) WorkforcePoolMapOutput {
	return o
}

func (o WorkforcePoolMapOutput) MapIndex(k pulumi.StringInput) WorkforcePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkforcePool {
		return vs[0].(map[string]*WorkforcePool)[vs[1].(string)]
	}).(WorkforcePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkforcePoolInput)(nil)).Elem(), &WorkforcePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkforcePoolArrayInput)(nil)).Elem(), WorkforcePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkforcePoolMapInput)(nil)).Elem(), WorkforcePoolMap{})
	pulumi.RegisterOutputType(WorkforcePoolOutput{})
	pulumi.RegisterOutputType(WorkforcePoolArrayOutput{})
	pulumi.RegisterOutputType(WorkforcePoolMapOutput{})
}
