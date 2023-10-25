// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A connectivity test are a static analysis of your resource configurations
 * that enables you to evaluate connectivity to and from Google Cloud
 * resources in your Virtual Private Cloud (VPC) network.
 *
 * To get more information about ConnectivityTest, see:
 *
 * * [API documentation](https://cloud.google.com/network-intelligence-center/docs/connectivity-tests/reference/networkmanagement/rest/v1/projects.locations.global.connectivityTests)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/network-intelligence-center/docs)
 *
 * ## Example Usage
 * ### Network Management Connectivity Test Instances
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const vpc = new gcp.compute.Network("vpc", {});
 * const debian9 = gcp.compute.getImage({
 *     family: "debian-11",
 *     project: "debian-cloud",
 * });
 * const source = new gcp.compute.Instance("source", {
 *     machineType: "e2-medium",
 *     bootDisk: {
 *         initializeParams: {
 *             image: debian9.then(debian9 => debian9.id),
 *         },
 *     },
 *     networkInterfaces: [{
 *         network: vpc.id,
 *         accessConfigs: [{}],
 *     }],
 * });
 * const destination = new gcp.compute.Instance("destination", {
 *     machineType: "e2-medium",
 *     bootDisk: {
 *         initializeParams: {
 *             image: debian9.then(debian9 => debian9.id),
 *         },
 *     },
 *     networkInterfaces: [{
 *         network: vpc.id,
 *         accessConfigs: [{}],
 *     }],
 * });
 * const instance_test = new gcp.networkmanagement.ConnectivityTest("instance-test", {
 *     source: {
 *         instance: source.id,
 *     },
 *     destination: {
 *         instance: destination.id,
 *     },
 *     protocol: "TCP",
 * });
 * ```
 * ### Network Management Connectivity Test Addresses
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const vpc = new gcp.compute.Network("vpc", {});
 * const subnet = new gcp.compute.Subnetwork("subnet", {
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 *     network: vpc.id,
 * });
 * const source_addr = new gcp.compute.Address("source-addr", {
 *     subnetwork: subnet.id,
 *     addressType: "INTERNAL",
 *     address: "10.0.42.42",
 *     region: "us-central1",
 * });
 * const dest_addr = new gcp.compute.Address("dest-addr", {
 *     subnetwork: subnet.id,
 *     addressType: "INTERNAL",
 *     address: "10.0.43.43",
 *     region: "us-central1",
 * });
 * const address_test = new gcp.networkmanagement.ConnectivityTest("address-test", {
 *     source: {
 *         ipAddress: source_addr.address,
 *         projectId: source_addr.project,
 *         network: vpc.id,
 *         networkType: "GCP_NETWORK",
 *     },
 *     destination: {
 *         ipAddress: dest_addr.address,
 *         projectId: dest_addr.project,
 *         network: vpc.id,
 *     },
 *     protocol: "UDP",
 * });
 * ```
 *
 * ## Import
 *
 * ConnectivityTest can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networkmanagement/connectivityTest:ConnectivityTest default projects/{{project}}/locations/global/connectivityTests/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkmanagement/connectivityTest:ConnectivityTest default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkmanagement/connectivityTest:ConnectivityTest default {{name}}
 * ```
 */
export class ConnectivityTest extends pulumi.CustomResource {
    /**
     * Get an existing ConnectivityTest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectivityTestState, opts?: pulumi.CustomResourceOptions): ConnectivityTest {
        return new ConnectivityTest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networkmanagement/connectivityTest:ConnectivityTest';

    /**
     * Returns true if the given object is an instance of ConnectivityTest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectivityTest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectivityTest.__pulumiType;
    }

    /**
     * The user-supplied description of the Connectivity Test.
     * Maximum of 512 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Required. Destination specification of the Connectivity Test.
     * You can use a combination of destination IP address, Compute
     * Engine VM instance, or VPC network to uniquely identify the
     * destination location.
     * Even if the destination IP address is not unique, the source IP
     * location is unique. Usually, the analysis can infer the destination
     * endpoint from route information.
     * If the destination you specify is a VM instance and the instance has
     * multiple network interfaces, then you must also specify either a
     * destination IP address or VPC network to identify the destination
     * interface.
     * A reachability analysis proceeds even if the destination location
     * is ambiguous. However, the result can include endpoints that you
     * don't intend to test.
     * Structure is documented below.
     */
    public readonly destination!: pulumi.Output<outputs.networkmanagement.ConnectivityTestDestination>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Unique name for the connectivity test.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * IP Protocol of the test. When not provided, "TCP" is assumed.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * Other projects that may be relevant for reachability analysis.
     * This is applicable to scenarios where a test can cross project
     * boundaries.
     */
    public readonly relatedProjects!: pulumi.Output<string[] | undefined>;
    /**
     * Required. Source specification of the Connectivity Test.
     * You can use a combination of source IP address, virtual machine
     * (VM) instance, or Compute Engine network to uniquely identify the
     * source location.
     * Examples: If the source IP address is an internal IP address within
     * a Google Cloud Virtual Private Cloud (VPC) network, then you must
     * also specify the VPC network. Otherwise, specify the VM instance,
     * which already contains its internal IP address and VPC network
     * information.
     * If the source of the test is within an on-premises network, then
     * you must provide the destination VPC network.
     * If the source endpoint is a Compute Engine VM instance with multiple
     * network interfaces, the instance itself is not sufficient to
     * identify the endpoint. So, you must also specify the source IP
     * address or VPC network.
     * A reachability analysis proceeds even if the source location is
     * ambiguous. However, the test result may include endpoints that
     * you don't intend to test.
     * Structure is documented below.
     */
    public readonly source!: pulumi.Output<outputs.networkmanagement.ConnectivityTestSource>;

    /**
     * Create a ConnectivityTest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectivityTestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectivityTestArgs | ConnectivityTestState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectivityTestState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["relatedProjects"] = state ? state.relatedProjects : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
        } else {
            const args = argsOrState as ConnectivityTestArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["relatedProjects"] = args ? args.relatedProjects : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConnectivityTest.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectivityTest resources.
 */
export interface ConnectivityTestState {
    /**
     * The user-supplied description of the Connectivity Test.
     * Maximum of 512 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. Destination specification of the Connectivity Test.
     * You can use a combination of destination IP address, Compute
     * Engine VM instance, or VPC network to uniquely identify the
     * destination location.
     * Even if the destination IP address is not unique, the source IP
     * location is unique. Usually, the analysis can infer the destination
     * endpoint from route information.
     * If the destination you specify is a VM instance and the instance has
     * multiple network interfaces, then you must also specify either a
     * destination IP address or VPC network to identify the destination
     * interface.
     * A reachability analysis proceeds even if the destination location
     * is ambiguous. However, the result can include endpoints that you
     * don't intend to test.
     * Structure is documented below.
     */
    destination?: pulumi.Input<inputs.networkmanagement.ConnectivityTestDestination>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Unique name for the connectivity test.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * IP Protocol of the test. When not provided, "TCP" is assumed.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Other projects that may be relevant for reachability analysis.
     * This is applicable to scenarios where a test can cross project
     * boundaries.
     */
    relatedProjects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Source specification of the Connectivity Test.
     * You can use a combination of source IP address, virtual machine
     * (VM) instance, or Compute Engine network to uniquely identify the
     * source location.
     * Examples: If the source IP address is an internal IP address within
     * a Google Cloud Virtual Private Cloud (VPC) network, then you must
     * also specify the VPC network. Otherwise, specify the VM instance,
     * which already contains its internal IP address and VPC network
     * information.
     * If the source of the test is within an on-premises network, then
     * you must provide the destination VPC network.
     * If the source endpoint is a Compute Engine VM instance with multiple
     * network interfaces, the instance itself is not sufficient to
     * identify the endpoint. So, you must also specify the source IP
     * address or VPC network.
     * A reachability analysis proceeds even if the source location is
     * ambiguous. However, the test result may include endpoints that
     * you don't intend to test.
     * Structure is documented below.
     */
    source?: pulumi.Input<inputs.networkmanagement.ConnectivityTestSource>;
}

/**
 * The set of arguments for constructing a ConnectivityTest resource.
 */
export interface ConnectivityTestArgs {
    /**
     * The user-supplied description of the Connectivity Test.
     * Maximum of 512 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. Destination specification of the Connectivity Test.
     * You can use a combination of destination IP address, Compute
     * Engine VM instance, or VPC network to uniquely identify the
     * destination location.
     * Even if the destination IP address is not unique, the source IP
     * location is unique. Usually, the analysis can infer the destination
     * endpoint from route information.
     * If the destination you specify is a VM instance and the instance has
     * multiple network interfaces, then you must also specify either a
     * destination IP address or VPC network to identify the destination
     * interface.
     * A reachability analysis proceeds even if the destination location
     * is ambiguous. However, the result can include endpoints that you
     * don't intend to test.
     * Structure is documented below.
     */
    destination: pulumi.Input<inputs.networkmanagement.ConnectivityTestDestination>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Unique name for the connectivity test.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * IP Protocol of the test. When not provided, "TCP" is assumed.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Other projects that may be relevant for reachability analysis.
     * This is applicable to scenarios where a test can cross project
     * boundaries.
     */
    relatedProjects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Source specification of the Connectivity Test.
     * You can use a combination of source IP address, virtual machine
     * (VM) instance, or Compute Engine network to uniquely identify the
     * source location.
     * Examples: If the source IP address is an internal IP address within
     * a Google Cloud Virtual Private Cloud (VPC) network, then you must
     * also specify the VPC network. Otherwise, specify the VM instance,
     * which already contains its internal IP address and VPC network
     * information.
     * If the source of the test is within an on-premises network, then
     * you must provide the destination VPC network.
     * If the source endpoint is a Compute Engine VM instance with multiple
     * network interfaces, the instance itself is not sufficient to
     * identify the endpoint. So, you must also specify the source IP
     * address or VPC network.
     * A reachability analysis proceeds even if the source location is
     * ambiguous. However, the test result may include endpoints that
     * you don't intend to test.
     * Structure is documented below.
     */
    source: pulumi.Input<inputs.networkmanagement.ConnectivityTestSource>;
}
