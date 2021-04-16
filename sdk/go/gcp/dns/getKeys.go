// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the DNSKEY and DS records of DNSSEC-signed managed zones. For more information see the
// [official documentation](https://cloud.google.com/dns/docs/dnskeys/)
// and [API](https://cloud.google.com/dns/docs/reference/v1/dnsKeys).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/dns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := dns.NewManagedZone(ctx, "foo", &dns.ManagedZoneArgs{
// 			DnsName: pulumi.String("foo.bar."),
// 			DnssecConfig: &dns.ManagedZoneDnssecConfigArgs{
// 				State:        pulumi.String("on"),
// 				NonExistence: pulumi.String("nsec3"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("fooDnsDsRecord", fooDnsKeys.ApplyT(func(fooDnsKeys dns.GetKeysResult) (string, error) {
// 			return fooDnsKeys.KeySigningKeys[0].DsRecord, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetKeys(ctx *pulumi.Context, args *GetKeysArgs, opts ...pulumi.InvokeOption) (*GetKeysResult, error) {
	var rv GetKeysResult
	err := ctx.Invoke("gcp:dns/getKeys:getKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeys.
type GetKeysArgs struct {
	// The name or id of the Cloud DNS managed zone.
	ManagedZone string `pulumi:"managedZone"`
	// The ID of the project in which the resource belongs. If `project` is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getKeys.
type GetKeysResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Key-signing key (KSK) records. Structure is documented below. Additionally, the DS record is provided:
	KeySigningKeys []GetKeysKeySigningKey `pulumi:"keySigningKeys"`
	ManagedZone    string                 `pulumi:"managedZone"`
	Project        string                 `pulumi:"project"`
	// A list of Zone-signing key (ZSK) records. Structure is documented below.
	ZoneSigningKeys []GetKeysZoneSigningKey `pulumi:"zoneSigningKeys"`
}
