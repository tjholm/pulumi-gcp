// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tpu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get TensorFlow versions available for a project. For more information see the [official documentation](https://cloud.google.com/tpu/docs/) and [API](https://cloud.google.com/tpu/docs/reference/rest/v1/projects.locations.tensorflowVersions).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tpu"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = tpu.GetTensorflowVersions(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Configure Basic TPU Node With Available Version
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tpu"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			available, err := tpu.GetTensorflowVersions(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = tpu.NewNode(ctx, "tpu", &tpu.NodeArgs{
//				Zone:              pulumi.String("us-central1-b"),
//				AcceleratorType:   pulumi.String("v3-8"),
//				TensorflowVersion: pulumi.String(available.Versions[0]),
//				CidrBlock:         pulumi.String("10.2.0.0/29"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTensorflowVersions(ctx *pulumi.Context, args *GetTensorflowVersionsArgs, opts ...pulumi.InvokeOption) (*GetTensorflowVersionsResult, error) {
	var rv GetTensorflowVersionsResult
	err := ctx.Invoke("gcp:tpu/getTensorflowVersions:getTensorflowVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTensorflowVersions.
type GetTensorflowVersionsArgs struct {
	// The project to list versions for. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone to list versions for. If it
	// is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getTensorflowVersions.
type GetTensorflowVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Project string `pulumi:"project"`
	// The list of TensorFlow versions available for the given project and zone.
	Versions []string `pulumi:"versions"`
	Zone     string   `pulumi:"zone"`
}

func GetTensorflowVersionsOutput(ctx *pulumi.Context, args GetTensorflowVersionsOutputArgs, opts ...pulumi.InvokeOption) GetTensorflowVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTensorflowVersionsResult, error) {
			args := v.(GetTensorflowVersionsArgs)
			r, err := GetTensorflowVersions(ctx, &args, opts...)
			var s GetTensorflowVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTensorflowVersionsResultOutput)
}

// A collection of arguments for invoking getTensorflowVersions.
type GetTensorflowVersionsOutputArgs struct {
	// The project to list versions for. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The zone to list versions for. If it
	// is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetTensorflowVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTensorflowVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getTensorflowVersions.
type GetTensorflowVersionsResultOutput struct{ *pulumi.OutputState }

func (GetTensorflowVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTensorflowVersionsResult)(nil)).Elem()
}

func (o GetTensorflowVersionsResultOutput) ToGetTensorflowVersionsResultOutput() GetTensorflowVersionsResultOutput {
	return o
}

func (o GetTensorflowVersionsResultOutput) ToGetTensorflowVersionsResultOutputWithContext(ctx context.Context) GetTensorflowVersionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTensorflowVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTensorflowVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTensorflowVersionsResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetTensorflowVersionsResult) string { return v.Project }).(pulumi.StringOutput)
}

// The list of TensorFlow versions available for the given project and zone.
func (o GetTensorflowVersionsResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTensorflowVersionsResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func (o GetTensorflowVersionsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetTensorflowVersionsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTensorflowVersionsResultOutput{})
}
