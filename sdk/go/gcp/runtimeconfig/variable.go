// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// Example creating a RuntimeConfig variable.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/runtimeconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := runtimeconfig.NewConfig(ctx, "my-runtime-config", &runtimeconfig.ConfigArgs{
//				Description: pulumi.String("Runtime configuration values for my service"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = runtimeconfig.NewVariable(ctx, "environment", &runtimeconfig.VariableArgs{
//				Parent: my_runtime_config.Name,
//				Text:   pulumi.String("example.com"),
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
// You can also encode binary content using the `value` argument instead. The
// value must be base64 encoded.
//
// Example of using the `value` argument.
//
// ```go
// package main
//
// import (
//
//	"encoding/base64"
//	"os"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/runtimeconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func filebase64OrPanic(path string) pulumi.StringPtrInput {
//		if fileData, err := os.ReadFile(path); err == nil {
//			return pulumi.String(base64.StdEncoding.EncodeToString(fileData[:]))
//		} else {
//			panic(err.Error())
//		}
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := runtimeconfig.NewConfig(ctx, "my-runtime-config", &runtimeconfig.ConfigArgs{
//				Description: pulumi.String("Runtime configuration values for my service"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = runtimeconfig.NewVariable(ctx, "my-secret", &runtimeconfig.VariableArgs{
//				Parent: my_runtime_config.Name,
//				Value:  filebase64OrPanic("my-encrypted-secret.dat"),
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
// Runtime Config Variables can be imported using the `name` or full variable name, e.g.
//
// ```sh
//
//	$ pulumi import gcp:runtimeconfig/variable:Variable myvariable myconfig/myvariable
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:runtimeconfig/variable:Variable myvariable projects/my-gcp-project/configs/myconfig/variables/myvariable
//
// ```
//
//	When importing using only the name, the provider project must be set.
type Variable struct {
	pulumi.CustomResourceState

	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text pulumi.StringPtrOutput `pulumi:"text"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds, representing when the variable was last updated.
	// Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringOutput    `pulumi:"updateTime"`
	Value      pulumi.StringPtrOutput `pulumi:"value"`
}

// NewVariable registers a new resource with the given unique name, arguments, and options.
func NewVariable(ctx *pulumi.Context,
	name string, args *VariableArgs, opts ...pulumi.ResourceOption) (*Variable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.Text != nil {
		args.Text = pulumi.ToSecret(args.Text).(pulumi.StringPtrInput)
	}
	if args.Value != nil {
		args.Value = pulumi.ToSecret(args.Value).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"text",
		"value",
	})
	opts = append(opts, secrets)
	var resource Variable
	err := ctx.RegisterResource("gcp:runtimeconfig/variable:Variable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVariable gets an existing Variable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableState, opts ...pulumi.ResourceOption) (*Variable, error) {
	var resource Variable
	err := ctx.ReadResource("gcp:runtimeconfig/variable:Variable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Variable resources.
type variableState struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name *string `pulumi:"name"`
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent *string `pulumi:"parent"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text *string `pulumi:"text"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds, representing when the variable was last updated.
	// Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime *string `pulumi:"updateTime"`
	Value      *string `pulumi:"value"`
}

type VariableState struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name pulumi.StringPtrInput
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text pulumi.StringPtrInput
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds, representing when the variable was last updated.
	// Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringPtrInput
	Value      pulumi.StringPtrInput
}

func (VariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableState)(nil)).Elem()
}

type variableArgs struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name *string `pulumi:"name"`
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent string `pulumi:"parent"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text  *string `pulumi:"text"`
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a Variable resource.
type VariableArgs struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name pulumi.StringPtrInput
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text  pulumi.StringPtrInput
	Value pulumi.StringPtrInput
}

func (VariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableArgs)(nil)).Elem()
}

type VariableInput interface {
	pulumi.Input

	ToVariableOutput() VariableOutput
	ToVariableOutputWithContext(ctx context.Context) VariableOutput
}

func (*Variable) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (i *Variable) ToVariableOutput() VariableOutput {
	return i.ToVariableOutputWithContext(context.Background())
}

func (i *Variable) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableOutput)
}

// VariableArrayInput is an input type that accepts VariableArray and VariableArrayOutput values.
// You can construct a concrete instance of `VariableArrayInput` via:
//
//	VariableArray{ VariableArgs{...} }
type VariableArrayInput interface {
	pulumi.Input

	ToVariableArrayOutput() VariableArrayOutput
	ToVariableArrayOutputWithContext(context.Context) VariableArrayOutput
}

type VariableArray []VariableInput

func (VariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Variable)(nil)).Elem()
}

func (i VariableArray) ToVariableArrayOutput() VariableArrayOutput {
	return i.ToVariableArrayOutputWithContext(context.Background())
}

func (i VariableArray) ToVariableArrayOutputWithContext(ctx context.Context) VariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableArrayOutput)
}

// VariableMapInput is an input type that accepts VariableMap and VariableMapOutput values.
// You can construct a concrete instance of `VariableMapInput` via:
//
//	VariableMap{ "key": VariableArgs{...} }
type VariableMapInput interface {
	pulumi.Input

	ToVariableMapOutput() VariableMapOutput
	ToVariableMapOutputWithContext(context.Context) VariableMapOutput
}

type VariableMap map[string]VariableInput

func (VariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Variable)(nil)).Elem()
}

func (i VariableMap) ToVariableMapOutput() VariableMapOutput {
	return i.ToVariableMapOutputWithContext(context.Background())
}

func (i VariableMap) ToVariableMapOutputWithContext(ctx context.Context) VariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableMapOutput)
}

type VariableOutput struct{ *pulumi.OutputState }

func (VariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (o VariableOutput) ToVariableOutput() VariableOutput {
	return o
}

func (o VariableOutput) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return o
}

// The name of the variable to manage. Note that variable
// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
func (o VariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the RuntimeConfig resource containing this
// variable.
func (o VariableOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o VariableOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// or `value` - (Required) The content to associate with the variable.
// Exactly one of `text` or `variable` must be specified. If `text` is specified,
// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
// is specified, it must be base64 encoded and less than 4096 bytes in length.
func (o VariableOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.Text }).(pulumi.StringPtrOutput)
}

// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
// accurate to nanoseconds, representing when the variable was last updated.
// Example: "2016-10-09T12:33:37.578138407Z".
func (o VariableOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func (o VariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

type VariableArrayOutput struct{ *pulumi.OutputState }

func (VariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Variable)(nil)).Elem()
}

func (o VariableArrayOutput) ToVariableArrayOutput() VariableArrayOutput {
	return o
}

func (o VariableArrayOutput) ToVariableArrayOutputWithContext(ctx context.Context) VariableArrayOutput {
	return o
}

func (o VariableArrayOutput) Index(i pulumi.IntInput) VariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Variable {
		return vs[0].([]*Variable)[vs[1].(int)]
	}).(VariableOutput)
}

type VariableMapOutput struct{ *pulumi.OutputState }

func (VariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Variable)(nil)).Elem()
}

func (o VariableMapOutput) ToVariableMapOutput() VariableMapOutput {
	return o
}

func (o VariableMapOutput) ToVariableMapOutputWithContext(ctx context.Context) VariableMapOutput {
	return o
}

func (o VariableMapOutput) MapIndex(k pulumi.StringInput) VariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Variable {
		return vs[0].(map[string]*Variable)[vs[1].(string)]
	}).(VariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VariableInput)(nil)).Elem(), &Variable{})
	pulumi.RegisterInputType(reflect.TypeOf((*VariableArrayInput)(nil)).Elem(), VariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VariableMapInput)(nil)).Elem(), VariableMap{})
	pulumi.RegisterOutputType(VariableOutput{})
	pulumi.RegisterOutputType(VariableArrayOutput{})
	pulumi.RegisterOutputType(VariableMapOutput{})
}
