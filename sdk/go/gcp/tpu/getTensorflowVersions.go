// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tpu

import (
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/tpu"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := tpu.GetTensorflowVersions(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Configure Basic TPU Node With Available Version
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/tpu"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		available, err := tpu.GetTensorflowVersions(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = tpu.NewNode(ctx, "tpu", &tpu.NodeArgs{
// 			Zone:              pulumi.String("us-central1-b"),
// 			AcceleratorType:   pulumi.String("v3-8"),
// 			TensorflowVersion: pulumi.String(available.Versions[0]),
// 			CidrBlock:         pulumi.String("10.2.0.0/29"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
