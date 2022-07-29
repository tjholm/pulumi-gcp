// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Dataproc Metastore Federation Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultMetastoreService, err := dataproc.NewMetastoreService(ctx, "defaultMetastoreService", &dataproc.MetastoreServiceArgs{
// 			ServiceId: pulumi.String("fed"),
// 			Location:  pulumi.String("us-central1"),
// 			Tier:      pulumi.String("DEVELOPER"),
// 			HiveMetastoreConfig: &dataproc.MetastoreServiceHiveMetastoreConfigArgs{
// 				Version:          pulumi.String("3.1.2"),
// 				EndpointProtocol: pulumi.String("GRPC"),
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataproc.NewMetastoreFederation(ctx, "defaultMetastoreFederation", &dataproc.MetastoreFederationArgs{
// 			Location:     pulumi.String("us-central1"),
// 			FederationId: pulumi.String("fed"),
// 			Version:      pulumi.String("3.1.2"),
// 			BackendMetastores: dataproc.MetastoreFederationBackendMetastoreArray{
// 				&dataproc.MetastoreFederationBackendMetastoreArgs{
// 					Rank:          pulumi.String("1"),
// 					Name:          defaultMetastoreService.ID(),
// 					MetastoreType: pulumi.String("DATAPROC_METASTORE"),
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
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
// Federation can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dataproc/metastoreFederation:MetastoreFederation default projects/{{project}}/locations/{{location}}/federations/{{federation_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/metastoreFederation:MetastoreFederation default {{project}}/{{location}}/{{federation_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/metastoreFederation:MetastoreFederation default {{location}}/{{federation_id}}
// ```
type MetastoreFederation struct {
	pulumi.CustomResourceState

	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
	// Structure is documented below.
	BackendMetastores MetastoreFederationBackendMetastoreArrayOutput `pulumi:"backendMetastores"`
	// The URI of the endpoint used to access the metastore federation.
	EndpointUri pulumi.StringOutput `pulumi:"endpointUri"`
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 63 characters.
	FederationId pulumi.StringOutput `pulumi:"federationId"`
	// User-defined labels for the metastore federation.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where the metastore federation should reside.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The relative resource name of the metastore that is being federated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The current state of the metastore federation.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of the metastore federation, if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The globally unique resource identifier of the metastore federation.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewMetastoreFederation registers a new resource with the given unique name, arguments, and options.
func NewMetastoreFederation(ctx *pulumi.Context,
	name string, args *MetastoreFederationArgs, opts ...pulumi.ResourceOption) (*MetastoreFederation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendMetastores == nil {
		return nil, errors.New("invalid value for required argument 'BackendMetastores'")
	}
	if args.FederationId == nil {
		return nil, errors.New("invalid value for required argument 'FederationId'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource MetastoreFederation
	err := ctx.RegisterResource("gcp:dataproc/metastoreFederation:MetastoreFederation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetastoreFederation gets an existing MetastoreFederation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetastoreFederation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetastoreFederationState, opts ...pulumi.ResourceOption) (*MetastoreFederation, error) {
	var resource MetastoreFederation
	err := ctx.ReadResource("gcp:dataproc/metastoreFederation:MetastoreFederation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetastoreFederation resources.
type metastoreFederationState struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
	// Structure is documented below.
	BackendMetastores []MetastoreFederationBackendMetastore `pulumi:"backendMetastores"`
	// The URI of the endpoint used to access the metastore federation.
	EndpointUri *string `pulumi:"endpointUri"`
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 63 characters.
	FederationId *string `pulumi:"federationId"`
	// User-defined labels for the metastore federation.
	Labels map[string]string `pulumi:"labels"`
	// The location where the metastore federation should reside.
	Location *string `pulumi:"location"`
	// The relative resource name of the metastore that is being federated.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The current state of the metastore federation.
	State *string `pulumi:"state"`
	// Additional information about the current state of the metastore federation, if available.
	StateMessage *string `pulumi:"stateMessage"`
	// The globally unique resource identifier of the metastore federation.
	Uid *string `pulumi:"uid"`
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
	Version *string `pulumi:"version"`
}

type MetastoreFederationState struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
	// Structure is documented below.
	BackendMetastores MetastoreFederationBackendMetastoreArrayInput
	// The URI of the endpoint used to access the metastore federation.
	EndpointUri pulumi.StringPtrInput
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 63 characters.
	FederationId pulumi.StringPtrInput
	// User-defined labels for the metastore federation.
	Labels pulumi.StringMapInput
	// The location where the metastore federation should reside.
	Location pulumi.StringPtrInput
	// The relative resource name of the metastore that is being federated.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The current state of the metastore federation.
	State pulumi.StringPtrInput
	// Additional information about the current state of the metastore federation, if available.
	StateMessage pulumi.StringPtrInput
	// The globally unique resource identifier of the metastore federation.
	Uid pulumi.StringPtrInput
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
	Version pulumi.StringPtrInput
}

func (MetastoreFederationState) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationState)(nil)).Elem()
}

type metastoreFederationArgs struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
	// Structure is documented below.
	BackendMetastores []MetastoreFederationBackendMetastore `pulumi:"backendMetastores"`
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 63 characters.
	FederationId string `pulumi:"federationId"`
	// User-defined labels for the metastore federation.
	Labels map[string]string `pulumi:"labels"`
	// The location where the metastore federation should reside.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a MetastoreFederation resource.
type MetastoreFederationArgs struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
	// Structure is documented below.
	BackendMetastores MetastoreFederationBackendMetastoreArrayInput
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 63 characters.
	FederationId pulumi.StringInput
	// User-defined labels for the metastore federation.
	Labels pulumi.StringMapInput
	// The location where the metastore federation should reside.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
	Version pulumi.StringInput
}

func (MetastoreFederationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationArgs)(nil)).Elem()
}

type MetastoreFederationInput interface {
	pulumi.Input

	ToMetastoreFederationOutput() MetastoreFederationOutput
	ToMetastoreFederationOutputWithContext(ctx context.Context) MetastoreFederationOutput
}

func (*MetastoreFederation) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederation)(nil)).Elem()
}

func (i *MetastoreFederation) ToMetastoreFederationOutput() MetastoreFederationOutput {
	return i.ToMetastoreFederationOutputWithContext(context.Background())
}

func (i *MetastoreFederation) ToMetastoreFederationOutputWithContext(ctx context.Context) MetastoreFederationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationOutput)
}

// MetastoreFederationArrayInput is an input type that accepts MetastoreFederationArray and MetastoreFederationArrayOutput values.
// You can construct a concrete instance of `MetastoreFederationArrayInput` via:
//
//          MetastoreFederationArray{ MetastoreFederationArgs{...} }
type MetastoreFederationArrayInput interface {
	pulumi.Input

	ToMetastoreFederationArrayOutput() MetastoreFederationArrayOutput
	ToMetastoreFederationArrayOutputWithContext(context.Context) MetastoreFederationArrayOutput
}

type MetastoreFederationArray []MetastoreFederationInput

func (MetastoreFederationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederation)(nil)).Elem()
}

func (i MetastoreFederationArray) ToMetastoreFederationArrayOutput() MetastoreFederationArrayOutput {
	return i.ToMetastoreFederationArrayOutputWithContext(context.Background())
}

func (i MetastoreFederationArray) ToMetastoreFederationArrayOutputWithContext(ctx context.Context) MetastoreFederationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationArrayOutput)
}

// MetastoreFederationMapInput is an input type that accepts MetastoreFederationMap and MetastoreFederationMapOutput values.
// You can construct a concrete instance of `MetastoreFederationMapInput` via:
//
//          MetastoreFederationMap{ "key": MetastoreFederationArgs{...} }
type MetastoreFederationMapInput interface {
	pulumi.Input

	ToMetastoreFederationMapOutput() MetastoreFederationMapOutput
	ToMetastoreFederationMapOutputWithContext(context.Context) MetastoreFederationMapOutput
}

type MetastoreFederationMap map[string]MetastoreFederationInput

func (MetastoreFederationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederation)(nil)).Elem()
}

func (i MetastoreFederationMap) ToMetastoreFederationMapOutput() MetastoreFederationMapOutput {
	return i.ToMetastoreFederationMapOutputWithContext(context.Background())
}

func (i MetastoreFederationMap) ToMetastoreFederationMapOutputWithContext(ctx context.Context) MetastoreFederationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationMapOutput)
}

type MetastoreFederationOutput struct{ *pulumi.OutputState }

func (MetastoreFederationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederation)(nil)).Elem()
}

func (o MetastoreFederationOutput) ToMetastoreFederationOutput() MetastoreFederationOutput {
	return o
}

func (o MetastoreFederationOutput) ToMetastoreFederationOutputWithContext(ctx context.Context) MetastoreFederationOutput {
	return o
}

// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
// Structure is documented below.
func (o MetastoreFederationOutput) BackendMetastores() MetastoreFederationBackendMetastoreArrayOutput {
	return o.ApplyT(func(v *MetastoreFederation) MetastoreFederationBackendMetastoreArrayOutput {
		return v.BackendMetastores
	}).(MetastoreFederationBackendMetastoreArrayOutput)
}

// The URI of the endpoint used to access the metastore federation.
func (o MetastoreFederationOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.EndpointUri }).(pulumi.StringOutput)
}

// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
// 3 and 63 characters.
func (o MetastoreFederationOutput) FederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.FederationId }).(pulumi.StringOutput)
}

// User-defined labels for the metastore federation.
func (o MetastoreFederationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where the metastore federation should reside.
func (o MetastoreFederationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The relative resource name of the metastore that is being federated.
func (o MetastoreFederationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o MetastoreFederationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The current state of the metastore federation.
func (o MetastoreFederationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current state of the metastore federation, if available.
func (o MetastoreFederationOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.StateMessage }).(pulumi.StringOutput)
}

// The globally unique resource identifier of the metastore federation.
func (o MetastoreFederationOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
func (o MetastoreFederationOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederation) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type MetastoreFederationArrayOutput struct{ *pulumi.OutputState }

func (MetastoreFederationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederation)(nil)).Elem()
}

func (o MetastoreFederationArrayOutput) ToMetastoreFederationArrayOutput() MetastoreFederationArrayOutput {
	return o
}

func (o MetastoreFederationArrayOutput) ToMetastoreFederationArrayOutputWithContext(ctx context.Context) MetastoreFederationArrayOutput {
	return o
}

func (o MetastoreFederationArrayOutput) Index(i pulumi.IntInput) MetastoreFederationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetastoreFederation {
		return vs[0].([]*MetastoreFederation)[vs[1].(int)]
	}).(MetastoreFederationOutput)
}

type MetastoreFederationMapOutput struct{ *pulumi.OutputState }

func (MetastoreFederationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederation)(nil)).Elem()
}

func (o MetastoreFederationMapOutput) ToMetastoreFederationMapOutput() MetastoreFederationMapOutput {
	return o
}

func (o MetastoreFederationMapOutput) ToMetastoreFederationMapOutputWithContext(ctx context.Context) MetastoreFederationMapOutput {
	return o
}

func (o MetastoreFederationMapOutput) MapIndex(k pulumi.StringInput) MetastoreFederationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetastoreFederation {
		return vs[0].(map[string]*MetastoreFederation)[vs[1].(string)]
	}).(MetastoreFederationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationInput)(nil)).Elem(), &MetastoreFederation{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationArrayInput)(nil)).Elem(), MetastoreFederationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationMapInput)(nil)).Elem(), MetastoreFederationMap{})
	pulumi.RegisterOutputType(MetastoreFederationOutput{})
	pulumi.RegisterOutputType(MetastoreFederationArrayOutput{})
	pulumi.RegisterOutputType(MetastoreFederationMapOutput{})
}
