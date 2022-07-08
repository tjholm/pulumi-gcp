// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.clouddeploy.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DeliveryPipelineSerialPipelineStage {
    /**
     * @return Skaffold profiles to use when rendering the manifest for this stage&#39;s `Target`.
     * 
     */
    private final @Nullable List<String> profiles;
    /**
     * @return The target_id to which this stage points. This field refers exclusively to the last segment of a target name. For example, this field would just be `my-target` (rather than `projects/project/locations/location/targets/my-target`). The location of the `Target` is inferred to be the same as the location of the `DeliveryPipeline` that contains this `Stage`.
     * 
     */
    private final @Nullable String targetId;

    @CustomType.Constructor
    private DeliveryPipelineSerialPipelineStage(
        @CustomType.Parameter("profiles") @Nullable List<String> profiles,
        @CustomType.Parameter("targetId") @Nullable String targetId) {
        this.profiles = profiles;
        this.targetId = targetId;
    }

    /**
     * @return Skaffold profiles to use when rendering the manifest for this stage&#39;s `Target`.
     * 
     */
    public List<String> profiles() {
        return this.profiles == null ? List.of() : this.profiles;
    }
    /**
     * @return The target_id to which this stage points. This field refers exclusively to the last segment of a target name. For example, this field would just be `my-target` (rather than `projects/project/locations/location/targets/my-target`). The location of the `Target` is inferred to be the same as the location of the `DeliveryPipeline` that contains this `Stage`.
     * 
     */
    public Optional<String> targetId() {
        return Optional.ofNullable(this.targetId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DeliveryPipelineSerialPipelineStage defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> profiles;
        private @Nullable String targetId;

        public Builder() {
    	      // Empty
        }

        public Builder(DeliveryPipelineSerialPipelineStage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.profiles = defaults.profiles;
    	      this.targetId = defaults.targetId;
        }

        public Builder profiles(@Nullable List<String> profiles) {
            this.profiles = profiles;
            return this;
        }
        public Builder profiles(String... profiles) {
            return profiles(List.of(profiles));
        }
        public Builder targetId(@Nullable String targetId) {
            this.targetId = targetId;
            return this;
        }        public DeliveryPipelineSerialPipelineStage build() {
            return new DeliveryPipelineSerialPipelineStage(profiles, targetId);
        }
    }
}
