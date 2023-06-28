// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityArgs;
import com.pulumi.gcp.gkehub.inputs.FeatureSpecMulticlusteringressArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeatureSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final FeatureSpecArgs Empty = new FeatureSpecArgs();

    @Import(name="fleetobservability")
    private @Nullable Output<FeatureSpecFleetobservabilityArgs> fleetobservability;

    public Optional<Output<FeatureSpecFleetobservabilityArgs>> fleetobservability() {
        return Optional.ofNullable(this.fleetobservability);
    }

    /**
     * Multicluster Ingress-specific spec.
     * Structure is documented below.
     * 
     */
    @Import(name="multiclusteringress")
    private @Nullable Output<FeatureSpecMulticlusteringressArgs> multiclusteringress;

    /**
     * @return Multicluster Ingress-specific spec.
     * Structure is documented below.
     * 
     */
    public Optional<Output<FeatureSpecMulticlusteringressArgs>> multiclusteringress() {
        return Optional.ofNullable(this.multiclusteringress);
    }

    private FeatureSpecArgs() {}

    private FeatureSpecArgs(FeatureSpecArgs $) {
        this.fleetobservability = $.fleetobservability;
        this.multiclusteringress = $.multiclusteringress;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureSpecArgs $;

        public Builder() {
            $ = new FeatureSpecArgs();
        }

        public Builder(FeatureSpecArgs defaults) {
            $ = new FeatureSpecArgs(Objects.requireNonNull(defaults));
        }

        public Builder fleetobservability(@Nullable Output<FeatureSpecFleetobservabilityArgs> fleetobservability) {
            $.fleetobservability = fleetobservability;
            return this;
        }

        public Builder fleetobservability(FeatureSpecFleetobservabilityArgs fleetobservability) {
            return fleetobservability(Output.of(fleetobservability));
        }

        /**
         * @param multiclusteringress Multicluster Ingress-specific spec.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder multiclusteringress(@Nullable Output<FeatureSpecMulticlusteringressArgs> multiclusteringress) {
            $.multiclusteringress = multiclusteringress;
            return this;
        }

        /**
         * @param multiclusteringress Multicluster Ingress-specific spec.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder multiclusteringress(FeatureSpecMulticlusteringressArgs multiclusteringress) {
            return multiclusteringress(Output.of(multiclusteringress));
        }

        public FeatureSpecArgs build() {
            return $;
        }
    }

}
