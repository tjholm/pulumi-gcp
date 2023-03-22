// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beyondcorp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A BeyondCorp AppConnector resource represents an application facing component deployed proximal to
// and with direct access to the application instances. It is used to establish connectivity between the
// remote enterprise environment and GCP. It initiates connections to the applications and can proxy the
// data from users over the connection.
//
// To get more information about AppConnector, see:
//
// * [API documentation](https://cloud.google.com/beyondcorp/docs/reference/rest#rest-resource:-v1.projects.locations.appconnectors)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/beyondcorp-enterprise/docs/enable-app-connector)
//
// ## Example Usage
// ### Beyondcorp App Connector Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/beyondcorp"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			serviceAccount, err := serviceAccount.NewAccount(ctx, "serviceAccount", &serviceAccount.AccountArgs{
//				AccountId:   pulumi.String("my-account"),
//				DisplayName: pulumi.String("Test Service Account"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = beyondcorp.NewAppConnector(ctx, "appConnector", &beyondcorp.AppConnectorArgs{
//				PrincipalInfo: &beyondcorp.AppConnectorPrincipalInfoArgs{
//					ServiceAccount: &beyondcorp.AppConnectorPrincipalInfoServiceAccountArgs{
//						Email: serviceAccount.Email,
//					},
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
// ### Beyondcorp App Connector Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/beyondcorp"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			serviceAccount, err := serviceAccount.NewAccount(ctx, "serviceAccount", &serviceAccount.AccountArgs{
//				AccountId:   pulumi.String("my-account"),
//				DisplayName: pulumi.String("Test Service Account"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = beyondcorp.NewAppConnector(ctx, "appConnector", &beyondcorp.AppConnectorArgs{
//				Region:      pulumi.String("us-central1"),
//				DisplayName: pulumi.String("some display name"),
//				PrincipalInfo: &beyondcorp.AppConnectorPrincipalInfoArgs{
//					ServiceAccount: &beyondcorp.AppConnectorPrincipalInfoServiceAccountArgs{
//						Email: serviceAccount.Email,
//					},
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//					"bar": pulumi.String("baz"),
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
// # AppConnector can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:beyondcorp/appConnector:AppConnector default projects/{{project}}/locations/{{region}}/appConnectors/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:beyondcorp/appConnector:AppConnector default {{project}}/{{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:beyondcorp/appConnector:AppConnector default {{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:beyondcorp/appConnector:AppConnector default {{name}}
//
// ```
type AppConnector struct {
	pulumi.CustomResourceState

	// An arbitrary user-provided name for the AppConnector.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// ID of the AppConnector.
	Name pulumi.StringOutput `pulumi:"name"`
	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	PrincipalInfo AppConnectorPrincipalInfoOutput `pulumi:"principalInfo"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the AppConnector.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Represents the different states of a AppConnector.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewAppConnector registers a new resource with the given unique name, arguments, and options.
func NewAppConnector(ctx *pulumi.Context,
	name string, args *AppConnectorArgs, opts ...pulumi.ResourceOption) (*AppConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrincipalInfo == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalInfo'")
	}
	var resource AppConnector
	err := ctx.RegisterResource("gcp:beyondcorp/appConnector:AppConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppConnector gets an existing AppConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppConnectorState, opts ...pulumi.ResourceOption) (*AppConnector, error) {
	var resource AppConnector
	err := ctx.ReadResource("gcp:beyondcorp/appConnector:AppConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppConnector resources.
type appConnectorState struct {
	// An arbitrary user-provided name for the AppConnector.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// ID of the AppConnector.
	Name *string `pulumi:"name"`
	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	PrincipalInfo *AppConnectorPrincipalInfo `pulumi:"principalInfo"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the AppConnector.
	Region *string `pulumi:"region"`
	// Represents the different states of a AppConnector.
	State *string `pulumi:"state"`
}

type AppConnectorState struct {
	// An arbitrary user-provided name for the AppConnector.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapInput
	// ID of the AppConnector.
	Name pulumi.StringPtrInput
	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	PrincipalInfo AppConnectorPrincipalInfoPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the AppConnector.
	Region pulumi.StringPtrInput
	// Represents the different states of a AppConnector.
	State pulumi.StringPtrInput
}

func (AppConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectorState)(nil)).Elem()
}

type appConnectorArgs struct {
	// An arbitrary user-provided name for the AppConnector.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// ID of the AppConnector.
	Name *string `pulumi:"name"`
	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	PrincipalInfo AppConnectorPrincipalInfo `pulumi:"principalInfo"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the AppConnector.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a AppConnector resource.
type AppConnectorArgs struct {
	// An arbitrary user-provided name for the AppConnector.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapInput
	// ID of the AppConnector.
	Name pulumi.StringPtrInput
	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	PrincipalInfo AppConnectorPrincipalInfoInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the AppConnector.
	Region pulumi.StringPtrInput
}

func (AppConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectorArgs)(nil)).Elem()
}

type AppConnectorInput interface {
	pulumi.Input

	ToAppConnectorOutput() AppConnectorOutput
	ToAppConnectorOutputWithContext(ctx context.Context) AppConnectorOutput
}

func (*AppConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnector)(nil)).Elem()
}

func (i *AppConnector) ToAppConnectorOutput() AppConnectorOutput {
	return i.ToAppConnectorOutputWithContext(context.Background())
}

func (i *AppConnector) ToAppConnectorOutputWithContext(ctx context.Context) AppConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConnectorOutput)
}

// AppConnectorArrayInput is an input type that accepts AppConnectorArray and AppConnectorArrayOutput values.
// You can construct a concrete instance of `AppConnectorArrayInput` via:
//
//	AppConnectorArray{ AppConnectorArgs{...} }
type AppConnectorArrayInput interface {
	pulumi.Input

	ToAppConnectorArrayOutput() AppConnectorArrayOutput
	ToAppConnectorArrayOutputWithContext(context.Context) AppConnectorArrayOutput
}

type AppConnectorArray []AppConnectorInput

func (AppConnectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppConnector)(nil)).Elem()
}

func (i AppConnectorArray) ToAppConnectorArrayOutput() AppConnectorArrayOutput {
	return i.ToAppConnectorArrayOutputWithContext(context.Background())
}

func (i AppConnectorArray) ToAppConnectorArrayOutputWithContext(ctx context.Context) AppConnectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConnectorArrayOutput)
}

// AppConnectorMapInput is an input type that accepts AppConnectorMap and AppConnectorMapOutput values.
// You can construct a concrete instance of `AppConnectorMapInput` via:
//
//	AppConnectorMap{ "key": AppConnectorArgs{...} }
type AppConnectorMapInput interface {
	pulumi.Input

	ToAppConnectorMapOutput() AppConnectorMapOutput
	ToAppConnectorMapOutputWithContext(context.Context) AppConnectorMapOutput
}

type AppConnectorMap map[string]AppConnectorInput

func (AppConnectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppConnector)(nil)).Elem()
}

func (i AppConnectorMap) ToAppConnectorMapOutput() AppConnectorMapOutput {
	return i.ToAppConnectorMapOutputWithContext(context.Background())
}

func (i AppConnectorMap) ToAppConnectorMapOutputWithContext(ctx context.Context) AppConnectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConnectorMapOutput)
}

type AppConnectorOutput struct{ *pulumi.OutputState }

func (AppConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnector)(nil)).Elem()
}

func (o AppConnectorOutput) ToAppConnectorOutput() AppConnectorOutput {
	return o
}

func (o AppConnectorOutput) ToAppConnectorOutputWithContext(ctx context.Context) AppConnectorOutput {
	return o
}

// An arbitrary user-provided name for the AppConnector.
func (o AppConnectorOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppConnector) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource labels to represent user provided metadata.
func (o AppConnectorOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppConnector) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// ID of the AppConnector.
func (o AppConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Principal information about the Identity of the AppConnector.
// Structure is documented below.
func (o AppConnectorOutput) PrincipalInfo() AppConnectorPrincipalInfoOutput {
	return o.ApplyT(func(v *AppConnector) AppConnectorPrincipalInfoOutput { return v.PrincipalInfo }).(AppConnectorPrincipalInfoOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o AppConnectorOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnector) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region of the AppConnector.
func (o AppConnectorOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppConnector) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// Represents the different states of a AppConnector.
func (o AppConnectorOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnector) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type AppConnectorArrayOutput struct{ *pulumi.OutputState }

func (AppConnectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppConnector)(nil)).Elem()
}

func (o AppConnectorArrayOutput) ToAppConnectorArrayOutput() AppConnectorArrayOutput {
	return o
}

func (o AppConnectorArrayOutput) ToAppConnectorArrayOutputWithContext(ctx context.Context) AppConnectorArrayOutput {
	return o
}

func (o AppConnectorArrayOutput) Index(i pulumi.IntInput) AppConnectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppConnector {
		return vs[0].([]*AppConnector)[vs[1].(int)]
	}).(AppConnectorOutput)
}

type AppConnectorMapOutput struct{ *pulumi.OutputState }

func (AppConnectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppConnector)(nil)).Elem()
}

func (o AppConnectorMapOutput) ToAppConnectorMapOutput() AppConnectorMapOutput {
	return o
}

func (o AppConnectorMapOutput) ToAppConnectorMapOutputWithContext(ctx context.Context) AppConnectorMapOutput {
	return o
}

func (o AppConnectorMapOutput) MapIndex(k pulumi.StringInput) AppConnectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppConnector {
		return vs[0].(map[string]*AppConnector)[vs[1].(string)]
	}).(AppConnectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppConnectorInput)(nil)).Elem(), &AppConnector{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppConnectorArrayInput)(nil)).Elem(), AppConnectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppConnectorMapInput)(nil)).Elem(), AppConnectorMap{})
	pulumi.RegisterOutputType(AppConnectorOutput{})
	pulumi.RegisterOutputType(AppConnectorArrayOutput{})
	pulumi.RegisterOutputType(AppConnectorMapOutput{})
}
