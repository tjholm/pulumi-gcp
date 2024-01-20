// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.gkeonprem.inputs.VMwareNodePoolConfigVsphereConfigTagArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VMwareNodePoolConfigVsphereConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final VMwareNodePoolConfigVsphereConfigArgs Empty = new VMwareNodePoolConfigVsphereConfigArgs();

    /**
     * The name of the vCenter datastore. Inherited from the user cluster.
     * 
     */
    @Import(name="datastore")
    private @Nullable Output<String> datastore;

    /**
     * @return The name of the vCenter datastore. Inherited from the user cluster.
     * 
     */
    public Optional<Output<String>> datastore() {
        return Optional.ofNullable(this.datastore);
    }

    /**
     * Vsphere host groups to apply to all VMs in the node pool
     * 
     */
    @Import(name="hostGroups")
    private @Nullable Output<List<String>> hostGroups;

    /**
     * @return Vsphere host groups to apply to all VMs in the node pool
     * 
     */
    public Optional<Output<List<String>>> hostGroups() {
        return Optional.ofNullable(this.hostGroups);
    }

    /**
     * Tags to apply to VMs.
     * Structure is documented below.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<VMwareNodePoolConfigVsphereConfigTagArgs>> tags;

    /**
     * @return Tags to apply to VMs.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<VMwareNodePoolConfigVsphereConfigTagArgs>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private VMwareNodePoolConfigVsphereConfigArgs() {}

    private VMwareNodePoolConfigVsphereConfigArgs(VMwareNodePoolConfigVsphereConfigArgs $) {
        this.datastore = $.datastore;
        this.hostGroups = $.hostGroups;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VMwareNodePoolConfigVsphereConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VMwareNodePoolConfigVsphereConfigArgs $;

        public Builder() {
            $ = new VMwareNodePoolConfigVsphereConfigArgs();
        }

        public Builder(VMwareNodePoolConfigVsphereConfigArgs defaults) {
            $ = new VMwareNodePoolConfigVsphereConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datastore The name of the vCenter datastore. Inherited from the user cluster.
         * 
         * @return builder
         * 
         */
        public Builder datastore(@Nullable Output<String> datastore) {
            $.datastore = datastore;
            return this;
        }

        /**
         * @param datastore The name of the vCenter datastore. Inherited from the user cluster.
         * 
         * @return builder
         * 
         */
        public Builder datastore(String datastore) {
            return datastore(Output.of(datastore));
        }

        /**
         * @param hostGroups Vsphere host groups to apply to all VMs in the node pool
         * 
         * @return builder
         * 
         */
        public Builder hostGroups(@Nullable Output<List<String>> hostGroups) {
            $.hostGroups = hostGroups;
            return this;
        }

        /**
         * @param hostGroups Vsphere host groups to apply to all VMs in the node pool
         * 
         * @return builder
         * 
         */
        public Builder hostGroups(List<String> hostGroups) {
            return hostGroups(Output.of(hostGroups));
        }

        /**
         * @param hostGroups Vsphere host groups to apply to all VMs in the node pool
         * 
         * @return builder
         * 
         */
        public Builder hostGroups(String... hostGroups) {
            return hostGroups(List.of(hostGroups));
        }

        /**
         * @param tags Tags to apply to VMs.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<VMwareNodePoolConfigVsphereConfigTagArgs>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags to apply to VMs.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<VMwareNodePoolConfigVsphereConfigTagArgs> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags to apply to VMs.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder tags(VMwareNodePoolConfigVsphereConfigTagArgs... tags) {
            return tags(List.of(tags));
        }

        public VMwareNodePoolConfigVsphereConfigArgs build() {
            return $;
        }
    }

}
