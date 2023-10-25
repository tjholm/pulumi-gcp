// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Monitoring Service is the root resource under which operational aspects of a
// generic service are accessible. A service is some discrete, autonomous, and
// network-accessible unit, designed to solve an individual concern
//
// A monitoring Istio Canonical Service is automatically created by GCP to monitor
// Istio Canonical Services.
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
// * How-to Guides
//   - [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
//   - [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
//
// ## Example Usage
// ### Monitoring Istio Canonical Service
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/monitoring"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := monitoring.GetIstioCanonicalService(ctx, &monitoring.GetIstioCanonicalServiceArgs{
//				CanonicalService:          "prometheus",
//				CanonicalServiceNamespace: "istio-system",
//				MeshUid:                   "proj-573164786102",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIstioCanonicalService(ctx *pulumi.Context, args *GetIstioCanonicalServiceArgs, opts ...pulumi.InvokeOption) (*GetIstioCanonicalServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIstioCanonicalServiceResult
	err := ctx.Invoke("gcp:monitoring/getIstioCanonicalService:getIstioCanonicalService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIstioCanonicalService.
type GetIstioCanonicalServiceArgs struct {
	// The name of the canonical service underlying this service.
	// Corresponds to the destinationCanonicalServiceName metric label in label in Istio metrics.
	//
	// ***
	//
	// Other optional fields include:
	CanonicalService string `pulumi:"canonicalService"`
	// The namespace of the canonical service underlying this service.
	// Corresponds to the destinationCanonicalServiceNamespace metric label in Istio metrics.
	CanonicalServiceNamespace string `pulumi:"canonicalServiceNamespace"`
	// Identifier for the mesh in which this Istio service is defined.
	// Corresponds to the meshUid metric label in Istio metrics.
	MeshUid string `pulumi:"meshUid"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getIstioCanonicalService.
type GetIstioCanonicalServiceResult struct {
	CanonicalService          string `pulumi:"canonicalService"`
	CanonicalServiceNamespace string `pulumi:"canonicalServiceNamespace"`
	// Name used for UI elements listing this (Monitoring) Service.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	MeshUid string `pulumi:"meshUid"`
	// The full REST resource name for this channel. The syntax is:
	// `projects/[PROJECT_ID]/services/[SERVICE_ID]`.
	Name      string  `pulumi:"name"`
	Project   *string `pulumi:"project"`
	ServiceId string  `pulumi:"serviceId"`
	// Configuration for how to query telemetry on the Service. Structure is documented below.
	Telemetries []GetIstioCanonicalServiceTelemetry `pulumi:"telemetries"`
	UserLabels  map[string]string                   `pulumi:"userLabels"`
}

func GetIstioCanonicalServiceOutput(ctx *pulumi.Context, args GetIstioCanonicalServiceOutputArgs, opts ...pulumi.InvokeOption) GetIstioCanonicalServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIstioCanonicalServiceResult, error) {
			args := v.(GetIstioCanonicalServiceArgs)
			r, err := GetIstioCanonicalService(ctx, &args, opts...)
			var s GetIstioCanonicalServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIstioCanonicalServiceResultOutput)
}

// A collection of arguments for invoking getIstioCanonicalService.
type GetIstioCanonicalServiceOutputArgs struct {
	// The name of the canonical service underlying this service.
	// Corresponds to the destinationCanonicalServiceName metric label in label in Istio metrics.
	//
	// ***
	//
	// Other optional fields include:
	CanonicalService pulumi.StringInput `pulumi:"canonicalService"`
	// The namespace of the canonical service underlying this service.
	// Corresponds to the destinationCanonicalServiceNamespace metric label in Istio metrics.
	CanonicalServiceNamespace pulumi.StringInput `pulumi:"canonicalServiceNamespace"`
	// Identifier for the mesh in which this Istio service is defined.
	// Corresponds to the meshUid metric label in Istio metrics.
	MeshUid pulumi.StringInput `pulumi:"meshUid"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetIstioCanonicalServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIstioCanonicalServiceArgs)(nil)).Elem()
}

// A collection of values returned by getIstioCanonicalService.
type GetIstioCanonicalServiceResultOutput struct{ *pulumi.OutputState }

func (GetIstioCanonicalServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIstioCanonicalServiceResult)(nil)).Elem()
}

func (o GetIstioCanonicalServiceResultOutput) ToGetIstioCanonicalServiceResultOutput() GetIstioCanonicalServiceResultOutput {
	return o
}

func (o GetIstioCanonicalServiceResultOutput) ToGetIstioCanonicalServiceResultOutputWithContext(ctx context.Context) GetIstioCanonicalServiceResultOutput {
	return o
}

func (o GetIstioCanonicalServiceResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetIstioCanonicalServiceResult] {
	return pulumix.Output[GetIstioCanonicalServiceResult]{
		OutputState: o.OutputState,
	}
}

func (o GetIstioCanonicalServiceResultOutput) CanonicalService() pulumi.StringOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) string { return v.CanonicalService }).(pulumi.StringOutput)
}

func (o GetIstioCanonicalServiceResultOutput) CanonicalServiceNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) string { return v.CanonicalServiceNamespace }).(pulumi.StringOutput)
}

// Name used for UI elements listing this (Monitoring) Service.
func (o GetIstioCanonicalServiceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIstioCanonicalServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIstioCanonicalServiceResultOutput) MeshUid() pulumi.StringOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) string { return v.MeshUid }).(pulumi.StringOutput)
}

// The full REST resource name for this channel. The syntax is:
// `projects/[PROJECT_ID]/services/[SERVICE_ID]`.
func (o GetIstioCanonicalServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetIstioCanonicalServiceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o GetIstioCanonicalServiceResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

// Configuration for how to query telemetry on the Service. Structure is documented below.
func (o GetIstioCanonicalServiceResultOutput) Telemetries() GetIstioCanonicalServiceTelemetryArrayOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) []GetIstioCanonicalServiceTelemetry { return v.Telemetries }).(GetIstioCanonicalServiceTelemetryArrayOutput)
}

func (o GetIstioCanonicalServiceResultOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetIstioCanonicalServiceResult) map[string]string { return v.UserLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIstioCanonicalServiceResultOutput{})
}
