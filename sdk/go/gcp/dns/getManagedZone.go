// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides access to a zone's attributes within Google Cloud DNS.
// For more information see
// [the official documentation](https://cloud.google.com/dns/zones/)
// and
// [API](https://cloud.google.com/dns/api/v1/managedZones).
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			envDnsZone, err := dns.LookupManagedZone(ctx, &dns.LookupManagedZoneArgs{
//				Name: "qa-zone",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewRecordSet(ctx, "dns", &dns.RecordSetArgs{
//				Name:        pulumi.Sprintf("my-address.%v", envDnsZone.DnsName),
//				Type:        pulumi.String("TXT"),
//				Ttl:         pulumi.Int(300),
//				ManagedZone: pulumi.String(envDnsZone.Name),
//				Rrdatas: pulumi.StringArray{
//					pulumi.String("test"),
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
func LookupManagedZone(ctx *pulumi.Context, args *LookupManagedZoneArgs, opts ...pulumi.InvokeOption) (*LookupManagedZoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedZoneResult
	err := ctx.Invoke("gcp:dns/getManagedZone:getManagedZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getManagedZone.
type LookupManagedZoneArgs struct {
	// A unique name for the resource.
	Name string `pulumi:"name"`
	// The ID of the project for the Google Cloud DNS zone.  If this is not provided the default project will be used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getManagedZone.
type LookupManagedZoneResult struct {
	// A textual description field.
	Description string `pulumi:"description"`
	// The fully qualified DNS name of this zone, e.g. `example.io.`.
	DnsName       string `pulumi:"dnsName"`
	Id            string `pulumi:"id"`
	ManagedZoneId string `pulumi:"managedZoneId"`
	Name          string `pulumi:"name"`
	// The list of nameservers that will be authoritative for this
	// domain. Use NS records to redirect from your DNS provider to these names,
	// thus making Google Cloud DNS authoritative for this zone.
	NameServers []string `pulumi:"nameServers"`
	Project     *string  `pulumi:"project"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	Visibility string `pulumi:"visibility"`
}

func LookupManagedZoneOutput(ctx *pulumi.Context, args LookupManagedZoneOutputArgs, opts ...pulumi.InvokeOption) LookupManagedZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedZoneResult, error) {
			args := v.(LookupManagedZoneArgs)
			r, err := LookupManagedZone(ctx, &args, opts...)
			var s LookupManagedZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedZoneResultOutput)
}

// A collection of arguments for invoking getManagedZone.
type LookupManagedZoneOutputArgs struct {
	// A unique name for the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the project for the Google Cloud DNS zone.  If this is not provided the default project will be used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupManagedZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedZoneArgs)(nil)).Elem()
}

// A collection of values returned by getManagedZone.
type LookupManagedZoneResultOutput struct{ *pulumi.OutputState }

func (LookupManagedZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedZoneResult)(nil)).Elem()
}

func (o LookupManagedZoneResultOutput) ToLookupManagedZoneResultOutput() LookupManagedZoneResultOutput {
	return o
}

func (o LookupManagedZoneResultOutput) ToLookupManagedZoneResultOutputWithContext(ctx context.Context) LookupManagedZoneResultOutput {
	return o
}

// A textual description field.
func (o LookupManagedZoneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) string { return v.Description }).(pulumi.StringOutput)
}

// The fully qualified DNS name of this zone, e.g. `example.io.`.
func (o LookupManagedZoneResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) string { return v.DnsName }).(pulumi.StringOutput)
}

func (o LookupManagedZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedZoneResultOutput) ManagedZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) string { return v.ManagedZoneId }).(pulumi.StringOutput)
}

func (o LookupManagedZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of nameservers that will be authoritative for this
// domain. Use NS records to redirect from your DNS provider to these names,
// thus making Google Cloud DNS authoritative for this zone.
func (o LookupManagedZoneResultOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

func (o LookupManagedZoneResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// The zone's visibility: public zones are exposed to the Internet,
// while private zones are visible only to Virtual Private Cloud resources.
func (o LookupManagedZoneResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedZoneResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedZoneResultOutput{})
}
