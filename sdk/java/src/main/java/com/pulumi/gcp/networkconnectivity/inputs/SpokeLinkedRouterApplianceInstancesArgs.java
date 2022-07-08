// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkconnectivity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networkconnectivity.inputs.SpokeLinkedRouterApplianceInstancesInstanceArgs;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;


public final class SpokeLinkedRouterApplianceInstancesArgs extends com.pulumi.resources.ResourceArgs {

    public static final SpokeLinkedRouterApplianceInstancesArgs Empty = new SpokeLinkedRouterApplianceInstancesArgs();

    /**
     * The list of router appliance instances
     * 
     */
    @Import(name="instances", required=true)
    private Output<List<SpokeLinkedRouterApplianceInstancesInstanceArgs>> instances;

    /**
     * @return The list of router appliance instances
     * 
     */
    public Output<List<SpokeLinkedRouterApplianceInstancesInstanceArgs>> instances() {
        return this.instances;
    }

    /**
     * A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
     * 
     */
    @Import(name="siteToSiteDataTransfer", required=true)
    private Output<Boolean> siteToSiteDataTransfer;

    /**
     * @return A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
     * 
     */
    public Output<Boolean> siteToSiteDataTransfer() {
        return this.siteToSiteDataTransfer;
    }

    private SpokeLinkedRouterApplianceInstancesArgs() {}

    private SpokeLinkedRouterApplianceInstancesArgs(SpokeLinkedRouterApplianceInstancesArgs $) {
        this.instances = $.instances;
        this.siteToSiteDataTransfer = $.siteToSiteDataTransfer;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SpokeLinkedRouterApplianceInstancesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SpokeLinkedRouterApplianceInstancesArgs $;

        public Builder() {
            $ = new SpokeLinkedRouterApplianceInstancesArgs();
        }

        public Builder(SpokeLinkedRouterApplianceInstancesArgs defaults) {
            $ = new SpokeLinkedRouterApplianceInstancesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instances The list of router appliance instances
         * 
         * @return builder
         * 
         */
        public Builder instances(Output<List<SpokeLinkedRouterApplianceInstancesInstanceArgs>> instances) {
            $.instances = instances;
            return this;
        }

        /**
         * @param instances The list of router appliance instances
         * 
         * @return builder
         * 
         */
        public Builder instances(List<SpokeLinkedRouterApplianceInstancesInstanceArgs> instances) {
            return instances(Output.of(instances));
        }

        /**
         * @param instances The list of router appliance instances
         * 
         * @return builder
         * 
         */
        public Builder instances(SpokeLinkedRouterApplianceInstancesInstanceArgs... instances) {
            return instances(List.of(instances));
        }

        /**
         * @param siteToSiteDataTransfer A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
         * 
         * @return builder
         * 
         */
        public Builder siteToSiteDataTransfer(Output<Boolean> siteToSiteDataTransfer) {
            $.siteToSiteDataTransfer = siteToSiteDataTransfer;
            return this;
        }

        /**
         * @param siteToSiteDataTransfer A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
         * 
         * @return builder
         * 
         */
        public Builder siteToSiteDataTransfer(Boolean siteToSiteDataTransfer) {
            return siteToSiteDataTransfer(Output.of(siteToSiteDataTransfer));
        }

        public SpokeLinkedRouterApplianceInstancesArgs build() {
            $.instances = Objects.requireNonNull($.instances, "expected parameter 'instances' to be non-null");
            $.siteToSiteDataTransfer = Objects.requireNonNull($.siteToSiteDataTransfer, "expected parameter 'siteToSiteDataTransfer' to be non-null");
            return $;
        }
    }

}
