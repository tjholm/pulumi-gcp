// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.vertex.outputs.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AiFeatureStoreEntityTypeMonitoringConfig {
    /**
     * @return Configuration of how features in Featurestore are monitored.
     * Structure is documented below.
     * 
     */
    private final @Nullable AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis snapshotAnalysis;

    @CustomType.Constructor
    private AiFeatureStoreEntityTypeMonitoringConfig(@CustomType.Parameter("snapshotAnalysis") @Nullable AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis snapshotAnalysis) {
        this.snapshotAnalysis = snapshotAnalysis;
    }

    /**
     * @return Configuration of how features in Featurestore are monitored.
     * Structure is documented below.
     * 
     */
    public Optional<AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis> snapshotAnalysis() {
        return Optional.ofNullable(this.snapshotAnalysis);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AiFeatureStoreEntityTypeMonitoringConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis snapshotAnalysis;

        public Builder() {
    	      // Empty
        }

        public Builder(AiFeatureStoreEntityTypeMonitoringConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.snapshotAnalysis = defaults.snapshotAnalysis;
        }

        public Builder snapshotAnalysis(@Nullable AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis snapshotAnalysis) {
            this.snapshotAnalysis = snapshotAnalysis;
            return this;
        }        public AiFeatureStoreEntityTypeMonitoringConfig build() {
            return new AiFeatureStoreEntityTypeMonitoringConfig(snapshotAnalysis);
        }
    }
}
