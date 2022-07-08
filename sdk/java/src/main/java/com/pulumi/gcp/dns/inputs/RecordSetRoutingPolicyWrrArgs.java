// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class RecordSetRoutingPolicyWrrArgs extends com.pulumi.resources.ResourceArgs {

    public static final RecordSetRoutingPolicyWrrArgs Empty = new RecordSetRoutingPolicyWrrArgs();

    /**
     * Same as `rrdatas` above.
     * 
     */
    @Import(name="rrdatas", required=true)
    private Output<List<String>> rrdatas;

    /**
     * @return Same as `rrdatas` above.
     * 
     */
    public Output<List<String>> rrdatas() {
        return this.rrdatas;
    }

    /**
     * The ratio of traffic routed to the target.
     * 
     */
    @Import(name="weight", required=true)
    private Output<Double> weight;

    /**
     * @return The ratio of traffic routed to the target.
     * 
     */
    public Output<Double> weight() {
        return this.weight;
    }

    private RecordSetRoutingPolicyWrrArgs() {}

    private RecordSetRoutingPolicyWrrArgs(RecordSetRoutingPolicyWrrArgs $) {
        this.rrdatas = $.rrdatas;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RecordSetRoutingPolicyWrrArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RecordSetRoutingPolicyWrrArgs $;

        public Builder() {
            $ = new RecordSetRoutingPolicyWrrArgs();
        }

        public Builder(RecordSetRoutingPolicyWrrArgs defaults) {
            $ = new RecordSetRoutingPolicyWrrArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param rrdatas Same as `rrdatas` above.
         * 
         * @return builder
         * 
         */
        public Builder rrdatas(Output<List<String>> rrdatas) {
            $.rrdatas = rrdatas;
            return this;
        }

        /**
         * @param rrdatas Same as `rrdatas` above.
         * 
         * @return builder
         * 
         */
        public Builder rrdatas(List<String> rrdatas) {
            return rrdatas(Output.of(rrdatas));
        }

        /**
         * @param rrdatas Same as `rrdatas` above.
         * 
         * @return builder
         * 
         */
        public Builder rrdatas(String... rrdatas) {
            return rrdatas(List.of(rrdatas));
        }

        /**
         * @param weight The ratio of traffic routed to the target.
         * 
         * @return builder
         * 
         */
        public Builder weight(Output<Double> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight The ratio of traffic routed to the target.
         * 
         * @return builder
         * 
         */
        public Builder weight(Double weight) {
            return weight(Output.of(weight));
        }

        public RecordSetRoutingPolicyWrrArgs build() {
            $.rrdatas = Objects.requireNonNull($.rrdatas, "expected parameter 'rrdatas' to be non-null");
            $.weight = Objects.requireNonNull($.weight, "expected parameter 'weight' to be non-null");
            return $;
        }
    }

}
