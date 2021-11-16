// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Google SQL SSL Cert on a Google SQL Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/mysql/admin-api/v1beta4/sslCerts).
//
// ## Example Usage
//
// Example creating a SQL Client Certificate.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomId(ctx, "dbNameSuffix", &random.RandomIdArgs{
// 			ByteLength: pulumi.Int(4),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		master, err := sql.NewDatabaseInstance(ctx, "master", &sql.DatabaseInstanceArgs{
// 			DatabaseVersion: pulumi.String("MYSQL_5_7"),
// 			Settings: &sql.DatabaseInstanceSettingsArgs{
// 				Tier: pulumi.String("db-f1-micro"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewSslCert(ctx, "clientCert", &sql.SslCertArgs{
// 			CommonName: pulumi.String("client-name"),
// 			Instance:   master.Name,
// 		})
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
// Since the contents of the certificate cannot be accessed after its creation, this resource cannot be imported.
type SslCert struct {
	pulumi.CustomResourceState

	// The actual certificate data for this client certificate.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// The serial number extracted from the certificate data.
	CertSerialNumber pulumi.StringOutput `pulumi:"certSerialNumber"`
	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The time when the certificate was created in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time when the certificate expires in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The private key associated with the client certificate.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The CA cert of the server this client cert was generated from.
	ServerCaCert pulumi.StringOutput `pulumi:"serverCaCert"`
	// The SHA1 Fingerprint of the certificate.
	Sha1Fingerprint pulumi.StringOutput `pulumi:"sha1Fingerprint"`
}

// NewSslCert registers a new resource with the given unique name, arguments, and options.
func NewSslCert(ctx *pulumi.Context,
	name string, args *SslCertArgs, opts ...pulumi.ResourceOption) (*SslCert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommonName == nil {
		return nil, errors.New("invalid value for required argument 'CommonName'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	var resource SslCert
	err := ctx.RegisterResource("gcp:sql/sslCert:SslCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslCert gets an existing SslCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslCertState, opts ...pulumi.ResourceOption) (*SslCert, error) {
	var resource SslCert
	err := ctx.ReadResource("gcp:sql/sslCert:SslCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslCert resources.
type sslCertState struct {
	// The actual certificate data for this client certificate.
	Cert *string `pulumi:"cert"`
	// The serial number extracted from the certificate data.
	CertSerialNumber *string `pulumi:"certSerialNumber"`
	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName *string `pulumi:"commonName"`
	// The time when the certificate was created in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	CreateTime *string `pulumi:"createTime"`
	// The time when the certificate expires in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	ExpirationTime *string `pulumi:"expirationTime"`
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance *string `pulumi:"instance"`
	// The private key associated with the client certificate.
	PrivateKey *string `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The CA cert of the server this client cert was generated from.
	ServerCaCert *string `pulumi:"serverCaCert"`
	// The SHA1 Fingerprint of the certificate.
	Sha1Fingerprint *string `pulumi:"sha1Fingerprint"`
}

type SslCertState struct {
	// The actual certificate data for this client certificate.
	Cert pulumi.StringPtrInput
	// The serial number extracted from the certificate data.
	CertSerialNumber pulumi.StringPtrInput
	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName pulumi.StringPtrInput
	// The time when the certificate was created in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	CreateTime pulumi.StringPtrInput
	// The time when the certificate expires in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	ExpirationTime pulumi.StringPtrInput
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance pulumi.StringPtrInput
	// The private key associated with the client certificate.
	PrivateKey pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The CA cert of the server this client cert was generated from.
	ServerCaCert pulumi.StringPtrInput
	// The SHA1 Fingerprint of the certificate.
	Sha1Fingerprint pulumi.StringPtrInput
}

func (SslCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslCertState)(nil)).Elem()
}

type sslCertArgs struct {
	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName string `pulumi:"commonName"`
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance string `pulumi:"instance"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a SslCert resource.
type SslCertArgs struct {
	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName pulumi.StringInput
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (SslCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslCertArgs)(nil)).Elem()
}

type SslCertInput interface {
	pulumi.Input

	ToSslCertOutput() SslCertOutput
	ToSslCertOutputWithContext(ctx context.Context) SslCertOutput
}

func (*SslCert) ElementType() reflect.Type {
	return reflect.TypeOf((*SslCert)(nil))
}

func (i *SslCert) ToSslCertOutput() SslCertOutput {
	return i.ToSslCertOutputWithContext(context.Background())
}

func (i *SslCert) ToSslCertOutputWithContext(ctx context.Context) SslCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCertOutput)
}

func (i *SslCert) ToSslCertPtrOutput() SslCertPtrOutput {
	return i.ToSslCertPtrOutputWithContext(context.Background())
}

func (i *SslCert) ToSslCertPtrOutputWithContext(ctx context.Context) SslCertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCertPtrOutput)
}

type SslCertPtrInput interface {
	pulumi.Input

	ToSslCertPtrOutput() SslCertPtrOutput
	ToSslCertPtrOutputWithContext(ctx context.Context) SslCertPtrOutput
}

type sslCertPtrType SslCertArgs

func (*sslCertPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SslCert)(nil))
}

func (i *sslCertPtrType) ToSslCertPtrOutput() SslCertPtrOutput {
	return i.ToSslCertPtrOutputWithContext(context.Background())
}

func (i *sslCertPtrType) ToSslCertPtrOutputWithContext(ctx context.Context) SslCertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCertPtrOutput)
}

// SslCertArrayInput is an input type that accepts SslCertArray and SslCertArrayOutput values.
// You can construct a concrete instance of `SslCertArrayInput` via:
//
//          SslCertArray{ SslCertArgs{...} }
type SslCertArrayInput interface {
	pulumi.Input

	ToSslCertArrayOutput() SslCertArrayOutput
	ToSslCertArrayOutputWithContext(context.Context) SslCertArrayOutput
}

type SslCertArray []SslCertInput

func (SslCertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslCert)(nil)).Elem()
}

func (i SslCertArray) ToSslCertArrayOutput() SslCertArrayOutput {
	return i.ToSslCertArrayOutputWithContext(context.Background())
}

func (i SslCertArray) ToSslCertArrayOutputWithContext(ctx context.Context) SslCertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCertArrayOutput)
}

// SslCertMapInput is an input type that accepts SslCertMap and SslCertMapOutput values.
// You can construct a concrete instance of `SslCertMapInput` via:
//
//          SslCertMap{ "key": SslCertArgs{...} }
type SslCertMapInput interface {
	pulumi.Input

	ToSslCertMapOutput() SslCertMapOutput
	ToSslCertMapOutputWithContext(context.Context) SslCertMapOutput
}

type SslCertMap map[string]SslCertInput

func (SslCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslCert)(nil)).Elem()
}

func (i SslCertMap) ToSslCertMapOutput() SslCertMapOutput {
	return i.ToSslCertMapOutputWithContext(context.Background())
}

func (i SslCertMap) ToSslCertMapOutputWithContext(ctx context.Context) SslCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCertMapOutput)
}

type SslCertOutput struct{ *pulumi.OutputState }

func (SslCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslCert)(nil))
}

func (o SslCertOutput) ToSslCertOutput() SslCertOutput {
	return o
}

func (o SslCertOutput) ToSslCertOutputWithContext(ctx context.Context) SslCertOutput {
	return o
}

func (o SslCertOutput) ToSslCertPtrOutput() SslCertPtrOutput {
	return o.ToSslCertPtrOutputWithContext(context.Background())
}

func (o SslCertOutput) ToSslCertPtrOutputWithContext(ctx context.Context) SslCertPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslCert) *SslCert {
		return &v
	}).(SslCertPtrOutput)
}

type SslCertPtrOutput struct{ *pulumi.OutputState }

func (SslCertPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslCert)(nil))
}

func (o SslCertPtrOutput) ToSslCertPtrOutput() SslCertPtrOutput {
	return o
}

func (o SslCertPtrOutput) ToSslCertPtrOutputWithContext(ctx context.Context) SslCertPtrOutput {
	return o
}

func (o SslCertPtrOutput) Elem() SslCertOutput {
	return o.ApplyT(func(v *SslCert) SslCert {
		if v != nil {
			return *v
		}
		var ret SslCert
		return ret
	}).(SslCertOutput)
}

type SslCertArrayOutput struct{ *pulumi.OutputState }

func (SslCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SslCert)(nil))
}

func (o SslCertArrayOutput) ToSslCertArrayOutput() SslCertArrayOutput {
	return o
}

func (o SslCertArrayOutput) ToSslCertArrayOutputWithContext(ctx context.Context) SslCertArrayOutput {
	return o
}

func (o SslCertArrayOutput) Index(i pulumi.IntInput) SslCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SslCert {
		return vs[0].([]SslCert)[vs[1].(int)]
	}).(SslCertOutput)
}

type SslCertMapOutput struct{ *pulumi.OutputState }

func (SslCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SslCert)(nil))
}

func (o SslCertMapOutput) ToSslCertMapOutput() SslCertMapOutput {
	return o
}

func (o SslCertMapOutput) ToSslCertMapOutputWithContext(ctx context.Context) SslCertMapOutput {
	return o
}

func (o SslCertMapOutput) MapIndex(k pulumi.StringInput) SslCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SslCert {
		return vs[0].(map[string]SslCert)[vs[1].(string)]
	}).(SslCertOutput)
}

func init() {
	pulumi.RegisterOutputType(SslCertOutput{})
	pulumi.RegisterOutputType(SslCertPtrOutput{})
	pulumi.RegisterOutputType(SslCertArrayOutput{})
	pulumi.RegisterOutputType(SslCertMapOutput{})
}
