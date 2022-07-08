// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrun.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.cloudrun.outputs.DomainMappingStatusCondition;
import com.pulumi.gcp.cloudrun.outputs.DomainMappingStatusResourceRecord;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DomainMappingStatus {
    private final @Nullable List<DomainMappingStatusCondition> conditions;
    private final @Nullable String mappedRouteName;
    private final @Nullable Integer observedGeneration;
    private final @Nullable List<DomainMappingStatusResourceRecord> resourceRecords;

    @CustomType.Constructor
    private DomainMappingStatus(
        @CustomType.Parameter("conditions") @Nullable List<DomainMappingStatusCondition> conditions,
        @CustomType.Parameter("mappedRouteName") @Nullable String mappedRouteName,
        @CustomType.Parameter("observedGeneration") @Nullable Integer observedGeneration,
        @CustomType.Parameter("resourceRecords") @Nullable List<DomainMappingStatusResourceRecord> resourceRecords) {
        this.conditions = conditions;
        this.mappedRouteName = mappedRouteName;
        this.observedGeneration = observedGeneration;
        this.resourceRecords = resourceRecords;
    }

    public List<DomainMappingStatusCondition> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }
    public Optional<String> mappedRouteName() {
        return Optional.ofNullable(this.mappedRouteName);
    }
    public Optional<Integer> observedGeneration() {
        return Optional.ofNullable(this.observedGeneration);
    }
    public List<DomainMappingStatusResourceRecord> resourceRecords() {
        return this.resourceRecords == null ? List.of() : this.resourceRecords;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DomainMappingStatus defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<DomainMappingStatusCondition> conditions;
        private @Nullable String mappedRouteName;
        private @Nullable Integer observedGeneration;
        private @Nullable List<DomainMappingStatusResourceRecord> resourceRecords;

        public Builder() {
    	      // Empty
        }

        public Builder(DomainMappingStatus defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.conditions = defaults.conditions;
    	      this.mappedRouteName = defaults.mappedRouteName;
    	      this.observedGeneration = defaults.observedGeneration;
    	      this.resourceRecords = defaults.resourceRecords;
        }

        public Builder conditions(@Nullable List<DomainMappingStatusCondition> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(DomainMappingStatusCondition... conditions) {
            return conditions(List.of(conditions));
        }
        public Builder mappedRouteName(@Nullable String mappedRouteName) {
            this.mappedRouteName = mappedRouteName;
            return this;
        }
        public Builder observedGeneration(@Nullable Integer observedGeneration) {
            this.observedGeneration = observedGeneration;
            return this;
        }
        public Builder resourceRecords(@Nullable List<DomainMappingStatusResourceRecord> resourceRecords) {
            this.resourceRecords = resourceRecords;
            return this;
        }
        public Builder resourceRecords(DomainMappingStatusResourceRecord... resourceRecords) {
            return resourceRecords(List.of(resourceRecords));
        }        public DomainMappingStatus build() {
            return new DomainMappingStatus(conditions, mappedRouteName, observedGeneration, resourceRecords);
        }
    }
}
