// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class FeatureSpecClusterupgradePostConditionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final FeatureSpecClusterupgradePostConditionsArgs Empty = new FeatureSpecClusterupgradePostConditionsArgs();

    /**
     * Amount of time to &#34;soak&#34; after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.
     * 
     */
    @Import(name="soaking", required=true)
    private Output<String> soaking;

    /**
     * @return Amount of time to &#34;soak&#34; after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.
     * 
     */
    public Output<String> soaking() {
        return this.soaking;
    }

    private FeatureSpecClusterupgradePostConditionsArgs() {}

    private FeatureSpecClusterupgradePostConditionsArgs(FeatureSpecClusterupgradePostConditionsArgs $) {
        this.soaking = $.soaking;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureSpecClusterupgradePostConditionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureSpecClusterupgradePostConditionsArgs $;

        public Builder() {
            $ = new FeatureSpecClusterupgradePostConditionsArgs();
        }

        public Builder(FeatureSpecClusterupgradePostConditionsArgs defaults) {
            $ = new FeatureSpecClusterupgradePostConditionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param soaking Amount of time to &#34;soak&#34; after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.
         * 
         * @return builder
         * 
         */
        public Builder soaking(Output<String> soaking) {
            $.soaking = soaking;
            return this;
        }

        /**
         * @param soaking Amount of time to &#34;soak&#34; after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.
         * 
         * @return builder
         * 
         */
        public Builder soaking(String soaking) {
            return soaking(Output.of(soaking));
        }

        public FeatureSpecClusterupgradePostConditionsArgs build() {
            if ($.soaking == null) {
                throw new MissingRequiredPropertyException("FeatureSpecClusterupgradePostConditionsArgs", "soaking");
            }
            return $;
        }
    }

}
