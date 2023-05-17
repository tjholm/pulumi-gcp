// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VMwareClusterControlPlaneNodeVsphereConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final VMwareClusterControlPlaneNodeVsphereConfigArgs Empty = new VMwareClusterControlPlaneNodeVsphereConfigArgs();

    /**
     * (Output)
     * The Vsphere datastore used by the Control Plane Node.
     * 
     */
    @Import(name="datastore")
    private @Nullable Output<String> datastore;

    /**
     * @return (Output)
     * The Vsphere datastore used by the Control Plane Node.
     * 
     */
    public Optional<Output<String>> datastore() {
        return Optional.ofNullable(this.datastore);
    }

    private VMwareClusterControlPlaneNodeVsphereConfigArgs() {}

    private VMwareClusterControlPlaneNodeVsphereConfigArgs(VMwareClusterControlPlaneNodeVsphereConfigArgs $) {
        this.datastore = $.datastore;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VMwareClusterControlPlaneNodeVsphereConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VMwareClusterControlPlaneNodeVsphereConfigArgs $;

        public Builder() {
            $ = new VMwareClusterControlPlaneNodeVsphereConfigArgs();
        }

        public Builder(VMwareClusterControlPlaneNodeVsphereConfigArgs defaults) {
            $ = new VMwareClusterControlPlaneNodeVsphereConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datastore (Output)
         * The Vsphere datastore used by the Control Plane Node.
         * 
         * @return builder
         * 
         */
        public Builder datastore(@Nullable Output<String> datastore) {
            $.datastore = datastore;
            return this;
        }

        /**
         * @param datastore (Output)
         * The Vsphere datastore used by the Control Plane Node.
         * 
         * @return builder
         * 
         */
        public Builder datastore(String datastore) {
            return datastore(Output.of(datastore));
        }

        public VMwareClusterControlPlaneNodeVsphereConfigArgs build() {
            return $;
        }
    }

}
