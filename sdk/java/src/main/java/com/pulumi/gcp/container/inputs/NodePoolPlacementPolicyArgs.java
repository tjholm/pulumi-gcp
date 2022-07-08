// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class NodePoolPlacementPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodePoolPlacementPolicyArgs Empty = new NodePoolPlacementPolicyArgs();

    /**
     * The type of the policy. Supports a single value: COMPACT.
     * Specifying COMPACT placement policy type places node pool&#39;s nodes in a closer
     * physical proximity in order to reduce network latency between nodes.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of the policy. Supports a single value: COMPACT.
     * Specifying COMPACT placement policy type places node pool&#39;s nodes in a closer
     * physical proximity in order to reduce network latency between nodes.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private NodePoolPlacementPolicyArgs() {}

    private NodePoolPlacementPolicyArgs(NodePoolPlacementPolicyArgs $) {
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodePoolPlacementPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodePoolPlacementPolicyArgs $;

        public Builder() {
            $ = new NodePoolPlacementPolicyArgs();
        }

        public Builder(NodePoolPlacementPolicyArgs defaults) {
            $ = new NodePoolPlacementPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type The type of the policy. Supports a single value: COMPACT.
         * Specifying COMPACT placement policy type places node pool&#39;s nodes in a closer
         * physical proximity in order to reduce network latency between nodes.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the policy. Supports a single value: COMPACT.
         * Specifying COMPACT placement policy type places node pool&#39;s nodes in a closer
         * physical proximity in order to reduce network latency between nodes.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public NodePoolPlacementPolicyArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
