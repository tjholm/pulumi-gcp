// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterPscConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterPscConfigArgs Empty = new ClusterPscConfigArgs();

    /**
     * Create an instance that allows connections from Private Service Connect endpoints to the instance.
     * 
     */
    @Import(name="pscEnabled")
    private @Nullable Output<Boolean> pscEnabled;

    /**
     * @return Create an instance that allows connections from Private Service Connect endpoints to the instance.
     * 
     */
    public Optional<Output<Boolean>> pscEnabled() {
        return Optional.ofNullable(this.pscEnabled);
    }

    private ClusterPscConfigArgs() {}

    private ClusterPscConfigArgs(ClusterPscConfigArgs $) {
        this.pscEnabled = $.pscEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterPscConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterPscConfigArgs $;

        public Builder() {
            $ = new ClusterPscConfigArgs();
        }

        public Builder(ClusterPscConfigArgs defaults) {
            $ = new ClusterPscConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param pscEnabled Create an instance that allows connections from Private Service Connect endpoints to the instance.
         * 
         * @return builder
         * 
         */
        public Builder pscEnabled(@Nullable Output<Boolean> pscEnabled) {
            $.pscEnabled = pscEnabled;
            return this;
        }

        /**
         * @param pscEnabled Create an instance that allows connections from Private Service Connect endpoints to the instance.
         * 
         * @return builder
         * 
         */
        public Builder pscEnabled(Boolean pscEnabled) {
            return pscEnabled(Output.of(pscEnabled));
        }

        public ClusterPscConfigArgs build() {
            return $;
        }
    }

}
