// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A RegionSslCertificate resource, used for HTTPS load balancing. This resource
// provides a mechanism to upload an SSL key and certificate to
// the load balancer to serve secure connections from the user.
//
// To get more information about RegionSslCertificate, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionSslCertificates)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
//
// > **Warning:** All arguments including `certificate` and `privateKey` will be stored in the raw
// state as plain-text.
//
// ## Example Usage
// ### Region Ssl Certificate Basic
//
// ```go
// package main
//
// import (
//
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewRegionSslCertificate(ctx, "default", &compute.RegionSslCertificateArgs{
//				Region:      pulumi.String("us-central1"),
//				NamePrefix:  pulumi.String("my-certificate-"),
//				Description: pulumi.String("a description"),
//				PrivateKey:  readFileOrPanic("path/to/private.key"),
//				Certificate: readFileOrPanic("path/to/certificate.crt"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Region Ssl Certificate Random Provider
//
// ```go
// package main
//
// import (
//
//	"crypto/sha256"
//	"fmt"
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func filebase64sha256OrPanic(path string) pulumi.StringPtrInput {
//		if fileData, err := ioutil.ReadFile(path); err == nil {
//			hashedData := sha256.Sum256([]byte(fileData))
//			return pulumi.String(base64.StdEncoding.EncodeToString(hashedData[:]))
//		} else {
//			panic(err.Error())
//		}
//	}
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewRegionSslCertificate(ctx, "default", &compute.RegionSslCertificateArgs{
//				Region:      pulumi.String("us-central1"),
//				PrivateKey:  readFileOrPanic("path/to/private.key"),
//				Certificate: readFileOrPanic("path/to/certificate.crt"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = random.NewRandomId(ctx, "certificate", &random.RandomIdArgs{
//				ByteLength: pulumi.Int(4),
//				Prefix:     pulumi.String("my-certificate-"),
//				Keepers: pulumi.AnyMap{
//					"private_key": filebase64sha256OrPanic("path/to/private.key"),
//					"certificate": filebase64sha256OrPanic("path/to/certificate.crt"),
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
// # RegionSslCertificate can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default {{project}}/{{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default {{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionSslCertificate:RegionSslCertificate default {{name}}
//
// ```
type RegionSslCertificate struct {
	pulumi.CustomResourceState

	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The unique identifier for the resource.
	CertificateId pulumi.IntOutput `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewRegionSslCertificate registers a new resource with the given unique name, arguments, and options.
func NewRegionSslCertificate(ctx *pulumi.Context,
	name string, args *RegionSslCertificateArgs, opts ...pulumi.ResourceOption) (*RegionSslCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	var resource RegionSslCertificate
	err := ctx.RegisterResource("gcp:compute/regionSslCertificate:RegionSslCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionSslCertificate gets an existing RegionSslCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionSslCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionSslCertificateState, opts ...pulumi.ResourceOption) (*RegionSslCertificate, error) {
	var resource RegionSslCertificate
	err := ctx.ReadResource("gcp:compute/regionSslCertificate:RegionSslCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionSslCertificate resources.
type regionSslCertificateState struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate *string `pulumi:"certificate"`
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey *string `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type RegionSslCertificateState struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate pulumi.StringPtrInput
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (RegionSslCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionSslCertificateState)(nil)).Elem()
}

type regionSslCertificateArgs struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate string `pulumi:"certificate"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey string `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RegionSslCertificate resource.
type RegionSslCertificateArgs struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate pulumi.StringInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
}

func (RegionSslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionSslCertificateArgs)(nil)).Elem()
}

type RegionSslCertificateInput interface {
	pulumi.Input

	ToRegionSslCertificateOutput() RegionSslCertificateOutput
	ToRegionSslCertificateOutputWithContext(ctx context.Context) RegionSslCertificateOutput
}

func (*RegionSslCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionSslCertificate)(nil)).Elem()
}

func (i *RegionSslCertificate) ToRegionSslCertificateOutput() RegionSslCertificateOutput {
	return i.ToRegionSslCertificateOutputWithContext(context.Background())
}

func (i *RegionSslCertificate) ToRegionSslCertificateOutputWithContext(ctx context.Context) RegionSslCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionSslCertificateOutput)
}

// RegionSslCertificateArrayInput is an input type that accepts RegionSslCertificateArray and RegionSslCertificateArrayOutput values.
// You can construct a concrete instance of `RegionSslCertificateArrayInput` via:
//
//	RegionSslCertificateArray{ RegionSslCertificateArgs{...} }
type RegionSslCertificateArrayInput interface {
	pulumi.Input

	ToRegionSslCertificateArrayOutput() RegionSslCertificateArrayOutput
	ToRegionSslCertificateArrayOutputWithContext(context.Context) RegionSslCertificateArrayOutput
}

type RegionSslCertificateArray []RegionSslCertificateInput

func (RegionSslCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionSslCertificate)(nil)).Elem()
}

func (i RegionSslCertificateArray) ToRegionSslCertificateArrayOutput() RegionSslCertificateArrayOutput {
	return i.ToRegionSslCertificateArrayOutputWithContext(context.Background())
}

func (i RegionSslCertificateArray) ToRegionSslCertificateArrayOutputWithContext(ctx context.Context) RegionSslCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionSslCertificateArrayOutput)
}

// RegionSslCertificateMapInput is an input type that accepts RegionSslCertificateMap and RegionSslCertificateMapOutput values.
// You can construct a concrete instance of `RegionSslCertificateMapInput` via:
//
//	RegionSslCertificateMap{ "key": RegionSslCertificateArgs{...} }
type RegionSslCertificateMapInput interface {
	pulumi.Input

	ToRegionSslCertificateMapOutput() RegionSslCertificateMapOutput
	ToRegionSslCertificateMapOutputWithContext(context.Context) RegionSslCertificateMapOutput
}

type RegionSslCertificateMap map[string]RegionSslCertificateInput

func (RegionSslCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionSslCertificate)(nil)).Elem()
}

func (i RegionSslCertificateMap) ToRegionSslCertificateMapOutput() RegionSslCertificateMapOutput {
	return i.ToRegionSslCertificateMapOutputWithContext(context.Background())
}

func (i RegionSslCertificateMap) ToRegionSslCertificateMapOutputWithContext(ctx context.Context) RegionSslCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionSslCertificateMapOutput)
}

type RegionSslCertificateOutput struct{ *pulumi.OutputState }

func (RegionSslCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionSslCertificate)(nil)).Elem()
}

func (o RegionSslCertificateOutput) ToRegionSslCertificateOutput() RegionSslCertificateOutput {
	return o
}

func (o RegionSslCertificateOutput) ToRegionSslCertificateOutputWithContext(ctx context.Context) RegionSslCertificateOutput {
	return o
}

// The certificate in PEM format.
// The certificate chain must be no greater than 5 certs long.
// The chain must include at least one intermediate cert.
// **Note**: This property is sensitive and will not be displayed in the plan.
func (o RegionSslCertificateOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// The unique identifier for the resource.
func (o RegionSslCertificateOutput) CertificateId() pulumi.IntOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.IntOutput { return v.CertificateId }).(pulumi.IntOutput)
}

// Creation timestamp in RFC3339 text format.
func (o RegionSslCertificateOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o RegionSslCertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and match
// the regular expression `a-z?` which means the
// first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
func (o RegionSslCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the
// specified prefix. Conflicts with `name`.
func (o RegionSslCertificateOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// The write-only private key in PEM format.
// **Note**: This property is sensitive and will not be displayed in the plan.
func (o RegionSslCertificateOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RegionSslCertificateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The Region in which the created regional ssl certificate should reside.
// If it is not provided, the provider region is used.
func (o RegionSslCertificateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o RegionSslCertificateOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionSslCertificate) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

type RegionSslCertificateArrayOutput struct{ *pulumi.OutputState }

func (RegionSslCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionSslCertificate)(nil)).Elem()
}

func (o RegionSslCertificateArrayOutput) ToRegionSslCertificateArrayOutput() RegionSslCertificateArrayOutput {
	return o
}

func (o RegionSslCertificateArrayOutput) ToRegionSslCertificateArrayOutputWithContext(ctx context.Context) RegionSslCertificateArrayOutput {
	return o
}

func (o RegionSslCertificateArrayOutput) Index(i pulumi.IntInput) RegionSslCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionSslCertificate {
		return vs[0].([]*RegionSslCertificate)[vs[1].(int)]
	}).(RegionSslCertificateOutput)
}

type RegionSslCertificateMapOutput struct{ *pulumi.OutputState }

func (RegionSslCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionSslCertificate)(nil)).Elem()
}

func (o RegionSslCertificateMapOutput) ToRegionSslCertificateMapOutput() RegionSslCertificateMapOutput {
	return o
}

func (o RegionSslCertificateMapOutput) ToRegionSslCertificateMapOutputWithContext(ctx context.Context) RegionSslCertificateMapOutput {
	return o
}

func (o RegionSslCertificateMapOutput) MapIndex(k pulumi.StringInput) RegionSslCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionSslCertificate {
		return vs[0].(map[string]*RegionSslCertificate)[vs[1].(string)]
	}).(RegionSslCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionSslCertificateInput)(nil)).Elem(), &RegionSslCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionSslCertificateArrayInput)(nil)).Elem(), RegionSslCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionSslCertificateMapInput)(nil)).Elem(), RegionSslCertificateMap{})
	pulumi.RegisterOutputType(RegionSslCertificateOutput{})
	pulumi.RegisterOutputType(RegionSslCertificateArrayOutput{})
	pulumi.RegisterOutputType(RegionSslCertificateMapOutput{})
}
