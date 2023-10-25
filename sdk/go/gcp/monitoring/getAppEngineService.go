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
// An App Engine monitoring service is automatically created by GCP to monitor
// App Engine services.
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
// * How-to Guides
//   - [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
//   - [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
//
// ## Example Usage
// ### Monitoring App Engine Service
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/appengine"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/monitoring"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
//				Location: pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			object, err := storage.NewBucketObject(ctx, "object", &storage.BucketObjectArgs{
//				Bucket: bucket.Name,
//				Source: pulumi.NewFileAsset("./test-fixtures/hello-world.zip"),
//			})
//			if err != nil {
//				return err
//			}
//			myapp, err := appengine.NewStandardAppVersion(ctx, "myapp", &appengine.StandardAppVersionArgs{
//				VersionId: pulumi.String("v1"),
//				Service:   pulumi.String("myapp"),
//				Runtime:   pulumi.String("nodejs10"),
//				Entrypoint: &appengine.StandardAppVersionEntrypointArgs{
//					Shell: pulumi.String("node ./app.js"),
//				},
//				Deployment: &appengine.StandardAppVersionDeploymentArgs{
//					Zip: &appengine.StandardAppVersionDeploymentZipArgs{
//						SourceUrl: pulumi.All(bucket.Name, object.Name).ApplyT(func(_args []interface{}) (string, error) {
//							bucketName := _args[0].(string)
//							objectName := _args[1].(string)
//							return fmt.Sprintf("https://storage.googleapis.com/%v/%v", bucketName, objectName), nil
//						}).(pulumi.StringOutput),
//					},
//				},
//				EnvVariables: pulumi.StringMap{
//					"port": pulumi.String("8080"),
//				},
//				DeleteServiceOnDestroy: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_ = monitoring.GetAppEngineServiceOutput(ctx, monitoring.GetAppEngineServiceOutputArgs{
//				ModuleId: myapp.Service,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetAppEngineService(ctx *pulumi.Context, args *GetAppEngineServiceArgs, opts ...pulumi.InvokeOption) (*GetAppEngineServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppEngineServiceResult
	err := ctx.Invoke("gcp:monitoring/getAppEngineService:getAppEngineService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppEngineService.
type GetAppEngineServiceArgs struct {
	// The ID of the App Engine module underlying this
	// service. Corresponds to the moduleId resource label in the [gaeApp](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
	//
	// ***
	//
	// Other optional fields include:
	ModuleId string `pulumi:"moduleId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getAppEngineService.
type GetAppEngineServiceResult struct {
	// Name used for UI elements listing this (Monitoring) Service.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	ModuleId string `pulumi:"moduleId"`
	// The full REST resource name for this channel. The syntax is:
	// `projects/[PROJECT_ID]/services/[SERVICE_ID]`.
	Name      string  `pulumi:"name"`
	Project   *string `pulumi:"project"`
	ServiceId string  `pulumi:"serviceId"`
	// Configuration for how to query telemetry on the Service. Structure is documented below.
	Telemetries []GetAppEngineServiceTelemetry `pulumi:"telemetries"`
	UserLabels  map[string]string              `pulumi:"userLabels"`
}

func GetAppEngineServiceOutput(ctx *pulumi.Context, args GetAppEngineServiceOutputArgs, opts ...pulumi.InvokeOption) GetAppEngineServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppEngineServiceResult, error) {
			args := v.(GetAppEngineServiceArgs)
			r, err := GetAppEngineService(ctx, &args, opts...)
			var s GetAppEngineServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppEngineServiceResultOutput)
}

// A collection of arguments for invoking getAppEngineService.
type GetAppEngineServiceOutputArgs struct {
	// The ID of the App Engine module underlying this
	// service. Corresponds to the moduleId resource label in the [gaeApp](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
	//
	// ***
	//
	// Other optional fields include:
	ModuleId pulumi.StringInput `pulumi:"moduleId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetAppEngineServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppEngineServiceArgs)(nil)).Elem()
}

// A collection of values returned by getAppEngineService.
type GetAppEngineServiceResultOutput struct{ *pulumi.OutputState }

func (GetAppEngineServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppEngineServiceResult)(nil)).Elem()
}

func (o GetAppEngineServiceResultOutput) ToGetAppEngineServiceResultOutput() GetAppEngineServiceResultOutput {
	return o
}

func (o GetAppEngineServiceResultOutput) ToGetAppEngineServiceResultOutputWithContext(ctx context.Context) GetAppEngineServiceResultOutput {
	return o
}

func (o GetAppEngineServiceResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetAppEngineServiceResult] {
	return pulumix.Output[GetAppEngineServiceResult]{
		OutputState: o.OutputState,
	}
}

// Name used for UI elements listing this (Monitoring) Service.
func (o GetAppEngineServiceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppEngineServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppEngineServiceResultOutput) ModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) string { return v.ModuleId }).(pulumi.StringOutput)
}

// The full REST resource name for this channel. The syntax is:
// `projects/[PROJECT_ID]/services/[SERVICE_ID]`.
func (o GetAppEngineServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetAppEngineServiceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o GetAppEngineServiceResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

// Configuration for how to query telemetry on the Service. Structure is documented below.
func (o GetAppEngineServiceResultOutput) Telemetries() GetAppEngineServiceTelemetryArrayOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) []GetAppEngineServiceTelemetry { return v.Telemetries }).(GetAppEngineServiceTelemetryArrayOutput)
}

func (o GetAppEngineServiceResultOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetAppEngineServiceResult) map[string]string { return v.UserLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppEngineServiceResultOutput{})
}
