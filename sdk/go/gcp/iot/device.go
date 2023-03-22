// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Google Cloud IoT Core device.
//
// To get more information about Device, see:
//
// * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/iot/docs/)
//
// ## Example Usage
// ### Cloudiot Device Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iot"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			registry, err := iot.NewRegistry(ctx, "registry", nil)
//			if err != nil {
//				return err
//			}
//			_, err = iot.NewDevice(ctx, "test-device", &iot.DeviceArgs{
//				Registry: registry.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Cloudiot Device Full
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iot"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			registry, err := iot.NewRegistry(ctx, "registry", nil)
//			if err != nil {
//				return err
//			}
//			_, err = iot.NewDevice(ctx, "test-device", &iot.DeviceArgs{
//				Registry: registry.ID(),
//				Credentials: iot.DeviceCredentialArray{
//					&iot.DeviceCredentialArgs{
//						PublicKey: &iot.DeviceCredentialPublicKeyArgs{
//							Format: pulumi.String("RSA_PEM"),
//							Key:    readFileOrPanic("test-fixtures/rsa_public.pem"),
//						},
//					},
//				},
//				Blocked:  pulumi.Bool(false),
//				LogLevel: pulumi.String("INFO"),
//				Metadata: pulumi.StringMap{
//					"test_key_1": pulumi.String("test_value_1"),
//				},
//				GatewayConfig: &iot.DeviceGatewayConfigArgs{
//					GatewayType: pulumi.String("NON_GATEWAY"),
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
// # Device can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:iot/device:Device default {{registry}}/devices/{{name}}
//
// ```
type Device struct {
	pulumi.CustomResourceState

	// If a device is blocked, connections or requests from this device will fail.
	Blocked pulumi.BoolPtrOutput `pulumi:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
	// Structure is documented below.
	Configs DeviceConfigArrayOutput `pulumi:"configs"`
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials DeviceCredentialArrayOutput `pulumi:"credentials"`
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig DeviceGatewayConfigPtrOutput `pulumi:"gatewayConfig"`
	// The last time a cloud-to-device config version acknowledgment was received from the device.
	LastConfigAckTime pulumi.StringOutput `pulumi:"lastConfigAckTime"`
	// The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime pulumi.StringOutput `pulumi:"lastConfigSendTime"`
	// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
	// Structure is documented below.
	LastErrorStatuses DeviceLastErrorStatusArrayOutput `pulumi:"lastErrorStatuses"`
	// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
	LastErrorTime pulumi.StringOutput `pulumi:"lastErrorTime"`
	// The last time a telemetry event was received.
	LastEventTime pulumi.StringOutput `pulumi:"lastEventTime"`
	// The last time an MQTT PINGREQ was received.
	LastHeartbeatTime pulumi.StringOutput `pulumi:"lastHeartbeatTime"`
	// The last time a state event was received.
	LastStateTime pulumi.StringOutput `pulumi:"lastStateTime"`
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel pulumi.StringPtrOutput `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// A unique name for the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// A server-defined unique numeric ID for the device.
	// This is a more compact way to identify devices, and it is globally unique.
	NumId pulumi.StringOutput `pulumi:"numId"`
	// The name of the device registry where this device should be created.
	Registry pulumi.StringOutput `pulumi:"registry"`
	// The state most recently received from the device.
	// Structure is documented below.
	States DeviceStateTypeArrayOutput `pulumi:"states"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Registry == nil {
		return nil, errors.New("invalid value for required argument 'Registry'")
	}
	var resource Device
	err := ctx.RegisterResource("gcp:iot/device:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("gcp:iot/device:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked *bool `pulumi:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
	// Structure is documented below.
	Configs []DeviceConfig `pulumi:"configs"`
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials []DeviceCredential `pulumi:"credentials"`
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig *DeviceGatewayConfig `pulumi:"gatewayConfig"`
	// The last time a cloud-to-device config version acknowledgment was received from the device.
	LastConfigAckTime *string `pulumi:"lastConfigAckTime"`
	// The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime *string `pulumi:"lastConfigSendTime"`
	// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
	// Structure is documented below.
	LastErrorStatuses []DeviceLastErrorStatus `pulumi:"lastErrorStatuses"`
	// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
	LastErrorTime *string `pulumi:"lastErrorTime"`
	// The last time a telemetry event was received.
	LastEventTime *string `pulumi:"lastEventTime"`
	// The last time an MQTT PINGREQ was received.
	LastHeartbeatTime *string `pulumi:"lastHeartbeatTime"`
	// The last time a state event was received.
	LastStateTime *string `pulumi:"lastStateTime"`
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel *string `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device.
	Metadata map[string]string `pulumi:"metadata"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// A server-defined unique numeric ID for the device.
	// This is a more compact way to identify devices, and it is globally unique.
	NumId *string `pulumi:"numId"`
	// The name of the device registry where this device should be created.
	Registry *string `pulumi:"registry"`
	// The state most recently received from the device.
	// Structure is documented below.
	States []DeviceStateType `pulumi:"states"`
}

type DeviceState struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked pulumi.BoolPtrInput
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
	// Structure is documented below.
	Configs DeviceConfigArrayInput
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials DeviceCredentialArrayInput
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig DeviceGatewayConfigPtrInput
	// The last time a cloud-to-device config version acknowledgment was received from the device.
	LastConfigAckTime pulumi.StringPtrInput
	// The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime pulumi.StringPtrInput
	// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
	// Structure is documented below.
	LastErrorStatuses DeviceLastErrorStatusArrayInput
	// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
	LastErrorTime pulumi.StringPtrInput
	// The last time a telemetry event was received.
	LastEventTime pulumi.StringPtrInput
	// The last time an MQTT PINGREQ was received.
	LastHeartbeatTime pulumi.StringPtrInput
	// The last time a state event was received.
	LastStateTime pulumi.StringPtrInput
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel pulumi.StringPtrInput
	// The metadata key-value pairs assigned to the device.
	Metadata pulumi.StringMapInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// A server-defined unique numeric ID for the device.
	// This is a more compact way to identify devices, and it is globally unique.
	NumId pulumi.StringPtrInput
	// The name of the device registry where this device should be created.
	Registry pulumi.StringPtrInput
	// The state most recently received from the device.
	// Structure is documented below.
	States DeviceStateTypeArrayInput
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked *bool `pulumi:"blocked"`
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials []DeviceCredential `pulumi:"credentials"`
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig *DeviceGatewayConfig `pulumi:"gatewayConfig"`
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel *string `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device.
	Metadata map[string]string `pulumi:"metadata"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// The name of the device registry where this device should be created.
	Registry string `pulumi:"registry"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked pulumi.BoolPtrInput
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials DeviceCredentialArrayInput
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig DeviceGatewayConfigPtrInput
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel pulumi.StringPtrInput
	// The metadata key-value pairs assigned to the device.
	Metadata pulumi.StringMapInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// The name of the device registry where this device should be created.
	Registry pulumi.StringInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

// DeviceArrayInput is an input type that accepts DeviceArray and DeviceArrayOutput values.
// You can construct a concrete instance of `DeviceArrayInput` via:
//
//	DeviceArray{ DeviceArgs{...} }
type DeviceArrayInput interface {
	pulumi.Input

	ToDeviceArrayOutput() DeviceArrayOutput
	ToDeviceArrayOutputWithContext(context.Context) DeviceArrayOutput
}

type DeviceArray []DeviceInput

func (DeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (i DeviceArray) ToDeviceArrayOutput() DeviceArrayOutput {
	return i.ToDeviceArrayOutputWithContext(context.Background())
}

func (i DeviceArray) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceArrayOutput)
}

// DeviceMapInput is an input type that accepts DeviceMap and DeviceMapOutput values.
// You can construct a concrete instance of `DeviceMapInput` via:
//
//	DeviceMap{ "key": DeviceArgs{...} }
type DeviceMapInput interface {
	pulumi.Input

	ToDeviceMapOutput() DeviceMapOutput
	ToDeviceMapOutputWithContext(context.Context) DeviceMapOutput
}

type DeviceMap map[string]DeviceInput

func (DeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (i DeviceMap) ToDeviceMapOutput() DeviceMapOutput {
	return i.ToDeviceMapOutputWithContext(context.Background())
}

func (i DeviceMap) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceMapOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

// If a device is blocked, connections or requests from this device will fail.
func (o DeviceOutput) Blocked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.BoolPtrOutput { return v.Blocked }).(pulumi.BoolPtrOutput)
}

// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
// Structure is documented below.
func (o DeviceOutput) Configs() DeviceConfigArrayOutput {
	return o.ApplyT(func(v *Device) DeviceConfigArrayOutput { return v.Configs }).(DeviceConfigArrayOutput)
}

// The credentials used to authenticate this device.
// Structure is documented below.
func (o DeviceOutput) Credentials() DeviceCredentialArrayOutput {
	return o.ApplyT(func(v *Device) DeviceCredentialArrayOutput { return v.Credentials }).(DeviceCredentialArrayOutput)
}

// Gateway-related configuration and state.
// Structure is documented below.
func (o DeviceOutput) GatewayConfig() DeviceGatewayConfigPtrOutput {
	return o.ApplyT(func(v *Device) DeviceGatewayConfigPtrOutput { return v.GatewayConfig }).(DeviceGatewayConfigPtrOutput)
}

// The last time a cloud-to-device config version acknowledgment was received from the device.
func (o DeviceOutput) LastConfigAckTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.LastConfigAckTime }).(pulumi.StringOutput)
}

// The last time a cloud-to-device config version was sent to the device.
func (o DeviceOutput) LastConfigSendTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.LastConfigSendTime }).(pulumi.StringOutput)
}

// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
// Structure is documented below.
func (o DeviceOutput) LastErrorStatuses() DeviceLastErrorStatusArrayOutput {
	return o.ApplyT(func(v *Device) DeviceLastErrorStatusArrayOutput { return v.LastErrorStatuses }).(DeviceLastErrorStatusArrayOutput)
}

// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
func (o DeviceOutput) LastErrorTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.LastErrorTime }).(pulumi.StringOutput)
}

// The last time a telemetry event was received.
func (o DeviceOutput) LastEventTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.LastEventTime }).(pulumi.StringOutput)
}

// The last time an MQTT PINGREQ was received.
func (o DeviceOutput) LastHeartbeatTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.LastHeartbeatTime }).(pulumi.StringOutput)
}

// The last time a state event was received.
func (o DeviceOutput) LastStateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.LastStateTime }).(pulumi.StringOutput)
}

// The logging verbosity for device activity.
// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
func (o DeviceOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.LogLevel }).(pulumi.StringPtrOutput)
}

// The metadata key-value pairs assigned to the device.
func (o DeviceOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Device) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// A unique name for the resource.
func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A server-defined unique numeric ID for the device.
// This is a more compact way to identify devices, and it is globally unique.
func (o DeviceOutput) NumId() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.NumId }).(pulumi.StringOutput)
}

// The name of the device registry where this device should be created.
func (o DeviceOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Registry }).(pulumi.StringOutput)
}

// The state most recently received from the device.
// Structure is documented below.
func (o DeviceOutput) States() DeviceStateTypeArrayOutput {
	return o.ApplyT(func(v *Device) DeviceStateTypeArrayOutput { return v.States }).(DeviceStateTypeArrayOutput)
}

type DeviceArrayOutput struct{ *pulumi.OutputState }

func (DeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (o DeviceArrayOutput) ToDeviceArrayOutput() DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) Index(i pulumi.IntInput) DeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Device {
		return vs[0].([]*Device)[vs[1].(int)]
	}).(DeviceOutput)
}

type DeviceMapOutput struct{ *pulumi.OutputState }

func (DeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (o DeviceMapOutput) ToDeviceMapOutput() DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) MapIndex(k pulumi.StringInput) DeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Device {
		return vs[0].(map[string]*Device)[vs[1].(string)]
	}).(DeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceInput)(nil)).Elem(), &Device{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceArrayInput)(nil)).Elem(), DeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceMapInput)(nil)).Elem(), DeviceMap{})
	pulumi.RegisterOutputType(DeviceOutput{})
	pulumi.RegisterOutputType(DeviceArrayOutput{})
	pulumi.RegisterOutputType(DeviceMapOutput{})
}
