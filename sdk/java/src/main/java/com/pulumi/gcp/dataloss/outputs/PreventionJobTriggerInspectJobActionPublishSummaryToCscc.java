// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.util.Objects;

@CustomType
public final class PreventionJobTriggerInspectJobActionPublishSummaryToCscc {
    private PreventionJobTriggerInspectJobActionPublishSummaryToCscc() {}

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PreventionJobTriggerInspectJobActionPublishSummaryToCscc defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        public Builder() {}
        public Builder(PreventionJobTriggerInspectJobActionPublishSummaryToCscc defaults) {
    	      Objects.requireNonNull(defaults);
        }

        public PreventionJobTriggerInspectJobActionPublishSummaryToCscc build() {
            final var o = new PreventionJobTriggerInspectJobActionPublishSummaryToCscc();
            return o;
        }
    }
}
