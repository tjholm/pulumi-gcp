// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity {
    /**
     * @return Number of characters after the finding to consider.
     * 
     */
    private @Nullable Integer windowAfter;
    /**
     * @return Number of characters before the finding to consider.
     * 
     */
    private @Nullable Integer windowBefore;

    private PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity() {}
    /**
     * @return Number of characters after the finding to consider.
     * 
     */
    public Optional<Integer> windowAfter() {
        return Optional.ofNullable(this.windowAfter);
    }
    /**
     * @return Number of characters before the finding to consider.
     * 
     */
    public Optional<Integer> windowBefore() {
        return Optional.ofNullable(this.windowBefore);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer windowAfter;
        private @Nullable Integer windowBefore;
        public Builder() {}
        public Builder(PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.windowAfter = defaults.windowAfter;
    	      this.windowBefore = defaults.windowBefore;
        }

        @CustomType.Setter
        public Builder windowAfter(@Nullable Integer windowAfter) {
            this.windowAfter = windowAfter;
            return this;
        }
        @CustomType.Setter
        public Builder windowBefore(@Nullable Integer windowBefore) {
            this.windowBefore = windowBefore;
            return this;
        }
        public PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity build() {
            final var o = new PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity();
            o.windowAfter = windowAfter;
            o.windowBefore = windowBefore;
            return o;
        }
    }
}
