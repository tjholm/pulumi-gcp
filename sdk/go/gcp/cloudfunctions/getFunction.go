// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Google Cloud Function. For more information see
// the [official documentation](https://cloud.google.com/functions/docs/)
// and [API](https://cloud.google.com/functions/docs/apis).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfunctions.LookupFunction(ctx, &cloudfunctions.LookupFunctionArgs{
//				Name: "function",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	var rv LookupFunctionResult
	err := ctx.Invoke("gcp:cloudfunctions/getFunction:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunction.
type LookupFunctionArgs struct {
	// The name of a Cloud Function.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the resource belongs. If it
	// is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getFunction.
type LookupFunctionResult struct {
	// Available memory (in MB) to the function.
	AvailableMemoryMb         int                    `pulumi:"availableMemoryMb"`
	BuildEnvironmentVariables map[string]interface{} `pulumi:"buildEnvironmentVariables"`
	// Description of the function.
	Description      string `pulumi:"description"`
	DockerRegistry   string `pulumi:"dockerRegistry"`
	DockerRepository string `pulumi:"dockerRepository"`
	// Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
	EntryPoint           string                 `pulumi:"entryPoint"`
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below.
	EventTriggers             []GetFunctionEventTrigger `pulumi:"eventTriggers"`
	HttpsTriggerSecurityLevel string                    `pulumi:"httpsTriggerSecurityLevel"`
	// If function is triggered by HTTP, trigger URL is set here.
	HttpsTriggerUrl string `pulumi:"httpsTriggerUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Controls what traffic can reach the function.
	IngressSettings string `pulumi:"ingressSettings"`
	KmsKeyName      string `pulumi:"kmsKeyName"`
	// A map of labels applied to this function.
	Labels map[string]interface{} `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances int `pulumi:"maxInstances"`
	MinInstances int `pulumi:"minInstances"`
	// The name of the Cloud Function.
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The runtime in which the function is running.
	Runtime                    string                                 `pulumi:"runtime"`
	SecretEnvironmentVariables []GetFunctionSecretEnvironmentVariable `pulumi:"secretEnvironmentVariables"`
	SecretVolumes              []GetFunctionSecretVolume              `pulumi:"secretVolumes"`
	// The service account email to be assumed by the cloud function.
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket string `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject string `pulumi:"sourceArchiveObject"`
	// The URL of the Cloud Source Repository that the function is deployed from. Structure is documented below.
	SourceRepositories []GetFunctionSourceRepository `pulumi:"sourceRepositories"`
	// Function execution timeout (in seconds).
	Timeout int `pulumi:"timeout"`
	// If function is triggered by HTTP, this boolean is set.
	TriggerHttp bool `pulumi:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to.
	VpcConnector string `pulumi:"vpcConnector"`
	// The egress settings for the connector, controlling what traffic is diverted through it.
	VpcConnectorEgressSettings string `pulumi:"vpcConnectorEgressSettings"`
}

func LookupFunctionOutput(ctx *pulumi.Context, args LookupFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionResult, error) {
			args := v.(LookupFunctionArgs)
			r, err := LookupFunction(ctx, &args, opts...)
			var s LookupFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFunctionResultOutput)
}

// A collection of arguments for invoking getFunction.
type LookupFunctionOutputArgs struct {
	// The name of a Cloud Function.
	Name pulumi.StringInput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The region in which the resource belongs. If it
	// is not provided, the provider region is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionArgs)(nil)).Elem()
}

// A collection of values returned by getFunction.
type LookupFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionResult)(nil)).Elem()
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutput() LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutputWithContext(ctx context.Context) LookupFunctionResultOutput {
	return o
}

// Available memory (in MB) to the function.
func (o LookupFunctionResultOutput) AvailableMemoryMb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.AvailableMemoryMb }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) BuildEnvironmentVariables() pulumi.MapOutput {
	return o.ApplyT(func(v LookupFunctionResult) map[string]interface{} { return v.BuildEnvironmentVariables }).(pulumi.MapOutput)
}

// Description of the function.
func (o LookupFunctionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) DockerRegistry() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.DockerRegistry }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) DockerRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.DockerRepository }).(pulumi.StringOutput)
}

// Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
func (o LookupFunctionResultOutput) EntryPoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.EntryPoint }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) EnvironmentVariables() pulumi.MapOutput {
	return o.ApplyT(func(v LookupFunctionResult) map[string]interface{} { return v.EnvironmentVariables }).(pulumi.MapOutput)
}

// A source that fires events in response to a condition in another service. Structure is documented below.
func (o LookupFunctionResultOutput) EventTriggers() GetFunctionEventTriggerArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []GetFunctionEventTrigger { return v.EventTriggers }).(GetFunctionEventTriggerArrayOutput)
}

func (o LookupFunctionResultOutput) HttpsTriggerSecurityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.HttpsTriggerSecurityLevel }).(pulumi.StringOutput)
}

// If function is triggered by HTTP, trigger URL is set here.
func (o LookupFunctionResultOutput) HttpsTriggerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.HttpsTriggerUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Controls what traffic can reach the function.
func (o LookupFunctionResultOutput) IngressSettings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.IngressSettings }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

// A map of labels applied to this function.
func (o LookupFunctionResultOutput) Labels() pulumi.MapOutput {
	return o.ApplyT(func(v LookupFunctionResult) map[string]interface{} { return v.Labels }).(pulumi.MapOutput)
}

// The limit on the maximum number of function instances that may coexist at a given time.
func (o LookupFunctionResultOutput) MaxInstances() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.MaxInstances }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) MinInstances() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.MinInstances }).(pulumi.IntOutput)
}

// The name of the Cloud Function.
func (o LookupFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The runtime in which the function is running.
func (o LookupFunctionResultOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Runtime }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) SecretEnvironmentVariables() GetFunctionSecretEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []GetFunctionSecretEnvironmentVariable {
		return v.SecretEnvironmentVariables
	}).(GetFunctionSecretEnvironmentVariableArrayOutput)
}

func (o LookupFunctionResultOutput) SecretVolumes() GetFunctionSecretVolumeArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []GetFunctionSecretVolume { return v.SecretVolumes }).(GetFunctionSecretVolumeArrayOutput)
}

// The service account email to be assumed by the cloud function.
func (o LookupFunctionResultOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

// The GCS bucket containing the zip archive which contains the function.
func (o LookupFunctionResultOutput) SourceArchiveBucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.SourceArchiveBucket }).(pulumi.StringOutput)
}

// The source archive object (file) in archive bucket.
func (o LookupFunctionResultOutput) SourceArchiveObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.SourceArchiveObject }).(pulumi.StringOutput)
}

// The URL of the Cloud Source Repository that the function is deployed from. Structure is documented below.
func (o LookupFunctionResultOutput) SourceRepositories() GetFunctionSourceRepositoryArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []GetFunctionSourceRepository { return v.SourceRepositories }).(GetFunctionSourceRepositoryArrayOutput)
}

// Function execution timeout (in seconds).
func (o LookupFunctionResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.Timeout }).(pulumi.IntOutput)
}

// If function is triggered by HTTP, this boolean is set.
func (o LookupFunctionResultOutput) TriggerHttp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFunctionResult) bool { return v.TriggerHttp }).(pulumi.BoolOutput)
}

// The VPC Network Connector that this cloud function can connect to.
func (o LookupFunctionResultOutput) VpcConnector() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.VpcConnector }).(pulumi.StringOutput)
}

// The egress settings for the connector, controlling what traffic is diverted through it.
func (o LookupFunctionResultOutput) VpcConnectorEgressSettings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.VpcConnectorEgressSettings }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionResultOutput{})
}
