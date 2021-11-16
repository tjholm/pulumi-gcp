// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// App profile is a configuration object describing how Cloud Bigtable should treat traffic from a particular end user application.
//
// To get more information about AppProfile, see:
//
// * [API documentation](https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.appProfiles)
//
// ## Example Usage
// ### Bigtable App Profile Multicluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance, err := bigtable.NewInstance(ctx, "instance", &bigtable.InstanceArgs{
// 			Clusters: bigtable.InstanceClusterArray{
// 				&bigtable.InstanceClusterArgs{
// 					ClusterId:   pulumi.String("bt-instance"),
// 					Zone:        pulumi.String("us-central1-b"),
// 					NumNodes:    pulumi.Int(3),
// 					StorageType: pulumi.String("HDD"),
// 				},
// 			},
// 			DeletionProtection: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewAppProfile(ctx, "ap", &bigquery.AppProfileArgs{
// 			Instance:                  instance.Name,
// 			AppProfileId:              pulumi.String("bt-profile"),
// 			MultiClusterRoutingUseAny: pulumi.Bool(true),
// 			IgnoreWarnings:            pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Bigtable App Profile Singlecluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance, err := bigtable.NewInstance(ctx, "instance", &bigtable.InstanceArgs{
// 			Clusters: bigtable.InstanceClusterArray{
// 				&bigtable.InstanceClusterArgs{
// 					ClusterId:   pulumi.String("bt-instance"),
// 					Zone:        pulumi.String("us-central1-b"),
// 					NumNodes:    pulumi.Int(3),
// 					StorageType: pulumi.String("HDD"),
// 				},
// 			},
// 			DeletionProtection: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewAppProfile(ctx, "ap", &bigquery.AppProfileArgs{
// 			Instance:     instance.Name,
// 			AppProfileId: pulumi.String("bt-profile"),
// 			SingleClusterRouting: &bigquery.AppProfileSingleClusterRoutingArgs{
// 				ClusterId:                pulumi.String("bt-instance"),
// 				AllowTransactionalWrites: pulumi.Bool(true),
// 			},
// 			IgnoreWarnings: pulumi.Bool(true),
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
// AppProfile can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:bigquery/appProfile:AppProfile default projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/appProfile:AppProfile default {{project}}/{{instance}}/{{app_profile_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/appProfile:AppProfile default {{instance}}/{{app_profile_id}}
// ```
type AppProfile struct {
	pulumi.CustomResourceState

	// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	AppProfileId pulumi.StringOutput `pulumi:"appProfileId"`
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrOutput `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrOutput `pulumi:"instance"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrOutput `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting AppProfileSingleClusterRoutingPtrOutput `pulumi:"singleClusterRouting"`
}

// NewAppProfile registers a new resource with the given unique name, arguments, and options.
func NewAppProfile(ctx *pulumi.Context,
	name string, args *AppProfileArgs, opts ...pulumi.ResourceOption) (*AppProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppProfileId == nil {
		return nil, errors.New("invalid value for required argument 'AppProfileId'")
	}
	var resource AppProfile
	err := ctx.RegisterResource("gcp:bigquery/appProfile:AppProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppProfile gets an existing AppProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppProfileState, opts ...pulumi.ResourceOption) (*AppProfile, error) {
	var resource AppProfile
	err := ctx.ReadResource("gcp:bigquery/appProfile:AppProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppProfile resources.
type appProfileState struct {
	// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	AppProfileId *string `pulumi:"appProfileId"`
	// Long form description of the use case for this app profile.
	Description *string `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance *string `pulumi:"instance"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny *bool `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting *AppProfileSingleClusterRouting `pulumi:"singleClusterRouting"`
}

type AppProfileState struct {
	// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	AppProfileId pulumi.StringPtrInput
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrInput
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrInput
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrInput
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrInput
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting AppProfileSingleClusterRoutingPtrInput
}

func (AppProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*appProfileState)(nil)).Elem()
}

type appProfileArgs struct {
	// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	AppProfileId string `pulumi:"appProfileId"`
	// Long form description of the use case for this app profile.
	Description *string `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance *string `pulumi:"instance"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny *bool `pulumi:"multiClusterRoutingUseAny"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting *AppProfileSingleClusterRouting `pulumi:"singleClusterRouting"`
}

// The set of arguments for constructing a AppProfile resource.
type AppProfileArgs struct {
	// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	AppProfileId pulumi.StringInput
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrInput
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrInput
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrInput
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting AppProfileSingleClusterRoutingPtrInput
}

func (AppProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appProfileArgs)(nil)).Elem()
}

type AppProfileInput interface {
	pulumi.Input

	ToAppProfileOutput() AppProfileOutput
	ToAppProfileOutputWithContext(ctx context.Context) AppProfileOutput
}

func (*AppProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*AppProfile)(nil))
}

func (i *AppProfile) ToAppProfileOutput() AppProfileOutput {
	return i.ToAppProfileOutputWithContext(context.Background())
}

func (i *AppProfile) ToAppProfileOutputWithContext(ctx context.Context) AppProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProfileOutput)
}

func (i *AppProfile) ToAppProfilePtrOutput() AppProfilePtrOutput {
	return i.ToAppProfilePtrOutputWithContext(context.Background())
}

func (i *AppProfile) ToAppProfilePtrOutputWithContext(ctx context.Context) AppProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProfilePtrOutput)
}

type AppProfilePtrInput interface {
	pulumi.Input

	ToAppProfilePtrOutput() AppProfilePtrOutput
	ToAppProfilePtrOutputWithContext(ctx context.Context) AppProfilePtrOutput
}

type appProfilePtrType AppProfileArgs

func (*appProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppProfile)(nil))
}

func (i *appProfilePtrType) ToAppProfilePtrOutput() AppProfilePtrOutput {
	return i.ToAppProfilePtrOutputWithContext(context.Background())
}

func (i *appProfilePtrType) ToAppProfilePtrOutputWithContext(ctx context.Context) AppProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProfilePtrOutput)
}

// AppProfileArrayInput is an input type that accepts AppProfileArray and AppProfileArrayOutput values.
// You can construct a concrete instance of `AppProfileArrayInput` via:
//
//          AppProfileArray{ AppProfileArgs{...} }
type AppProfileArrayInput interface {
	pulumi.Input

	ToAppProfileArrayOutput() AppProfileArrayOutput
	ToAppProfileArrayOutputWithContext(context.Context) AppProfileArrayOutput
}

type AppProfileArray []AppProfileInput

func (AppProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppProfile)(nil)).Elem()
}

func (i AppProfileArray) ToAppProfileArrayOutput() AppProfileArrayOutput {
	return i.ToAppProfileArrayOutputWithContext(context.Background())
}

func (i AppProfileArray) ToAppProfileArrayOutputWithContext(ctx context.Context) AppProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProfileArrayOutput)
}

// AppProfileMapInput is an input type that accepts AppProfileMap and AppProfileMapOutput values.
// You can construct a concrete instance of `AppProfileMapInput` via:
//
//          AppProfileMap{ "key": AppProfileArgs{...} }
type AppProfileMapInput interface {
	pulumi.Input

	ToAppProfileMapOutput() AppProfileMapOutput
	ToAppProfileMapOutputWithContext(context.Context) AppProfileMapOutput
}

type AppProfileMap map[string]AppProfileInput

func (AppProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppProfile)(nil)).Elem()
}

func (i AppProfileMap) ToAppProfileMapOutput() AppProfileMapOutput {
	return i.ToAppProfileMapOutputWithContext(context.Background())
}

func (i AppProfileMap) ToAppProfileMapOutputWithContext(ctx context.Context) AppProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProfileMapOutput)
}

type AppProfileOutput struct{ *pulumi.OutputState }

func (AppProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppProfile)(nil))
}

func (o AppProfileOutput) ToAppProfileOutput() AppProfileOutput {
	return o
}

func (o AppProfileOutput) ToAppProfileOutputWithContext(ctx context.Context) AppProfileOutput {
	return o
}

func (o AppProfileOutput) ToAppProfilePtrOutput() AppProfilePtrOutput {
	return o.ToAppProfilePtrOutputWithContext(context.Background())
}

func (o AppProfileOutput) ToAppProfilePtrOutputWithContext(ctx context.Context) AppProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppProfile) *AppProfile {
		return &v
	}).(AppProfilePtrOutput)
}

type AppProfilePtrOutput struct{ *pulumi.OutputState }

func (AppProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppProfile)(nil))
}

func (o AppProfilePtrOutput) ToAppProfilePtrOutput() AppProfilePtrOutput {
	return o
}

func (o AppProfilePtrOutput) ToAppProfilePtrOutputWithContext(ctx context.Context) AppProfilePtrOutput {
	return o
}

func (o AppProfilePtrOutput) Elem() AppProfileOutput {
	return o.ApplyT(func(v *AppProfile) AppProfile {
		if v != nil {
			return *v
		}
		var ret AppProfile
		return ret
	}).(AppProfileOutput)
}

type AppProfileArrayOutput struct{ *pulumi.OutputState }

func (AppProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppProfile)(nil))
}

func (o AppProfileArrayOutput) ToAppProfileArrayOutput() AppProfileArrayOutput {
	return o
}

func (o AppProfileArrayOutput) ToAppProfileArrayOutputWithContext(ctx context.Context) AppProfileArrayOutput {
	return o
}

func (o AppProfileArrayOutput) Index(i pulumi.IntInput) AppProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppProfile {
		return vs[0].([]AppProfile)[vs[1].(int)]
	}).(AppProfileOutput)
}

type AppProfileMapOutput struct{ *pulumi.OutputState }

func (AppProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppProfile)(nil))
}

func (o AppProfileMapOutput) ToAppProfileMapOutput() AppProfileMapOutput {
	return o
}

func (o AppProfileMapOutput) ToAppProfileMapOutputWithContext(ctx context.Context) AppProfileMapOutput {
	return o
}

func (o AppProfileMapOutput) MapIndex(k pulumi.StringInput) AppProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppProfile {
		return vs[0].(map[string]AppProfile)[vs[1].(string)]
	}).(AppProfileOutput)
}

func init() {
	pulumi.RegisterOutputType(AppProfileOutput{})
	pulumi.RegisterOutputType(AppProfilePtrOutput{})
	pulumi.RegisterOutputType(AppProfileArrayOutput{})
	pulumi.RegisterOutputType(AppProfileMapOutput{})
}
