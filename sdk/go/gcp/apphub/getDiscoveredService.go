// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apphub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a discovered service from its uri.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apphub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apphub.GetDiscoveredService(ctx, &apphub.GetDiscoveredServiceArgs{
//				Location:   "my-location",
//				ServiceUri: "my-service-uri",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetDiscoveredService(ctx *pulumi.Context, args *GetDiscoveredServiceArgs, opts ...pulumi.InvokeOption) (*GetDiscoveredServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDiscoveredServiceResult
	err := ctx.Invoke("gcp:apphub/getDiscoveredService:getDiscoveredService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDiscoveredService.
type GetDiscoveredServiceArgs struct {
	// The location of the discovered service.
	Location string `pulumi:"location"`
	// The host project of the discovered service.
	Project *string `pulumi:"project"`
	// The uri of the service.
	ServiceUri string `pulumi:"serviceUri"`
}

// A collection of values returned by getDiscoveredService.
type GetDiscoveredServiceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The location that the underlying resource resides in.
	Location string `pulumi:"location"`
	// Resource name of a Service. Format: "projects/{host-project-id}/locations/{location}/applications/{application-id}/services/{service-id}".
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Properties of an underlying compute resource that can comprise a Service. Structure is documented below
	ServiceProperties []GetDiscoveredServiceServiceProperty `pulumi:"serviceProperties"`
	// Reference to an underlying networking resource that can comprise a Service. Structure is documented below
	ServiceReferences []GetDiscoveredServiceServiceReference `pulumi:"serviceReferences"`
	ServiceUri        string                                 `pulumi:"serviceUri"`
}

func GetDiscoveredServiceOutput(ctx *pulumi.Context, args GetDiscoveredServiceOutputArgs, opts ...pulumi.InvokeOption) GetDiscoveredServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDiscoveredServiceResult, error) {
			args := v.(GetDiscoveredServiceArgs)
			r, err := GetDiscoveredService(ctx, &args, opts...)
			var s GetDiscoveredServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDiscoveredServiceResultOutput)
}

// A collection of arguments for invoking getDiscoveredService.
type GetDiscoveredServiceOutputArgs struct {
	// The location of the discovered service.
	Location pulumi.StringInput `pulumi:"location"`
	// The host project of the discovered service.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The uri of the service.
	ServiceUri pulumi.StringInput `pulumi:"serviceUri"`
}

func (GetDiscoveredServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiscoveredServiceArgs)(nil)).Elem()
}

// A collection of values returned by getDiscoveredService.
type GetDiscoveredServiceResultOutput struct{ *pulumi.OutputState }

func (GetDiscoveredServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiscoveredServiceResult)(nil)).Elem()
}

func (o GetDiscoveredServiceResultOutput) ToGetDiscoveredServiceResultOutput() GetDiscoveredServiceResultOutput {
	return o
}

func (o GetDiscoveredServiceResultOutput) ToGetDiscoveredServiceResultOutputWithContext(ctx context.Context) GetDiscoveredServiceResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDiscoveredServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiscoveredServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location that the underlying resource resides in.
func (o GetDiscoveredServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiscoveredServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name of a Service. Format: "projects/{host-project-id}/locations/{location}/applications/{application-id}/services/{service-id}".
func (o GetDiscoveredServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiscoveredServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDiscoveredServiceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDiscoveredServiceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// Properties of an underlying compute resource that can comprise a Service. Structure is documented below
func (o GetDiscoveredServiceResultOutput) ServiceProperties() GetDiscoveredServiceServicePropertyArrayOutput {
	return o.ApplyT(func(v GetDiscoveredServiceResult) []GetDiscoveredServiceServiceProperty { return v.ServiceProperties }).(GetDiscoveredServiceServicePropertyArrayOutput)
}

// Reference to an underlying networking resource that can comprise a Service. Structure is documented below
func (o GetDiscoveredServiceResultOutput) ServiceReferences() GetDiscoveredServiceServiceReferenceArrayOutput {
	return o.ApplyT(func(v GetDiscoveredServiceResult) []GetDiscoveredServiceServiceReference { return v.ServiceReferences }).(GetDiscoveredServiceServiceReferenceArrayOutput)
}

func (o GetDiscoveredServiceResultOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiscoveredServiceResult) string { return v.ServiceUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDiscoveredServiceResultOutput{})
}
