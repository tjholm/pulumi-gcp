// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs extends com.pulumi.resources.ResourceArgs {

    public static final VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs Empty = new VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs();

    /**
     * Hostname of the machine. VM&#39;s name will be used if this field is empty.
     * 
     */
    @Import(name="hostname", required=true)
    private Output<String> hostname;

    /**
     * @return Hostname of the machine. VM&#39;s name will be used if this field is empty.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }

    /**
     * IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    private VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs() {}

    private VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs(VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs $) {
        this.hostname = $.hostname;
        this.ip = $.ip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs $;

        public Builder() {
            $ = new VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs();
        }

        public Builder(VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs defaults) {
            $ = new VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hostname Hostname of the machine. VM&#39;s name will be used if this field is empty.
         * 
         * @return builder
         * 
         */
        public Builder hostname(Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname Hostname of the machine. VM&#39;s name will be used if this field is empty.
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param ip IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        public VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs build() {
            $.hostname = Objects.requireNonNull($.hostname, "expected parameter 'hostname' to be non-null");
            $.ip = Objects.requireNonNull($.ip, "expected parameter 'ip' to be non-null");
            return $;
        }
    }

}
