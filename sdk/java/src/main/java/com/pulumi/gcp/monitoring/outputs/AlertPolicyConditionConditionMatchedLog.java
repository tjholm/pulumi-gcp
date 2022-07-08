// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class AlertPolicyConditionConditionMatchedLog {
    /**
     * @return A logs-based filter.
     * 
     */
    private final String filter;
    /**
     * @return A map from a label key to an extractor expression, which is used to
     * extract the value for this label key. Each entry in this map is
     * a specification for how data should be extracted from log entries that
     * match filter. Each combination of extracted values is treated as
     * a separate rule for the purposes of triggering notifications.
     * Label keys and corresponding values can be used in notifications
     * generated by this condition.
     * 
     */
    private final @Nullable Map<String,String> labelExtractors;

    @CustomType.Constructor
    private AlertPolicyConditionConditionMatchedLog(
        @CustomType.Parameter("filter") String filter,
        @CustomType.Parameter("labelExtractors") @Nullable Map<String,String> labelExtractors) {
        this.filter = filter;
        this.labelExtractors = labelExtractors;
    }

    /**
     * @return A logs-based filter.
     * 
     */
    public String filter() {
        return this.filter;
    }
    /**
     * @return A map from a label key to an extractor expression, which is used to
     * extract the value for this label key. Each entry in this map is
     * a specification for how data should be extracted from log entries that
     * match filter. Each combination of extracted values is treated as
     * a separate rule for the purposes of triggering notifications.
     * Label keys and corresponding values can be used in notifications
     * generated by this condition.
     * 
     */
    public Map<String,String> labelExtractors() {
        return this.labelExtractors == null ? Map.of() : this.labelExtractors;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AlertPolicyConditionConditionMatchedLog defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String filter;
        private @Nullable Map<String,String> labelExtractors;

        public Builder() {
    	      // Empty
        }

        public Builder(AlertPolicyConditionConditionMatchedLog defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filter = defaults.filter;
    	      this.labelExtractors = defaults.labelExtractors;
        }

        public Builder filter(String filter) {
            this.filter = Objects.requireNonNull(filter);
            return this;
        }
        public Builder labelExtractors(@Nullable Map<String,String> labelExtractors) {
            this.labelExtractors = labelExtractors;
            return this;
        }        public AlertPolicyConditionConditionMatchedLog build() {
            return new AlertPolicyConditionConditionMatchedLog(filter, labelExtractors);
        }
    }
}
