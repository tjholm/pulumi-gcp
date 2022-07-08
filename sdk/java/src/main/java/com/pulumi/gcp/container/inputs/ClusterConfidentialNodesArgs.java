// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;


public final class ClusterConfidentialNodesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterConfidentialNodesArgs Empty = new ClusterConfidentialNodesArgs();

    /**
     * Enable the PodSecurityPolicy controller for this cluster.
     * If enabled, pods must be valid under a PodSecurityPolicy to be created.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Enable the PodSecurityPolicy controller for this cluster.
     * If enabled, pods must be valid under a PodSecurityPolicy to be created.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    private ClusterConfidentialNodesArgs() {}

    private ClusterConfidentialNodesArgs(ClusterConfidentialNodesArgs $) {
        this.enabled = $.enabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterConfidentialNodesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterConfidentialNodesArgs $;

        public Builder() {
            $ = new ClusterConfidentialNodesArgs();
        }

        public Builder(ClusterConfidentialNodesArgs defaults) {
            $ = new ClusterConfidentialNodesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Enable the PodSecurityPolicy controller for this cluster.
         * If enabled, pods must be valid under a PodSecurityPolicy to be created.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Enable the PodSecurityPolicy controller for this cluster.
         * If enabled, pods must be valid under a PodSecurityPolicy to be created.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public ClusterConfidentialNodesArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            return $;
        }
    }

}
