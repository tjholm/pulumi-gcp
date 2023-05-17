// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class PreventionJobTriggerInspectJobInspectConfigCustomInfoTypeDictionaryCloudStoragePath {
    /**
     * @return A url representing a file or path (no wildcards) in Cloud Storage. Example: `gs://[BUCKET_NAME]/dictionary.txt`
     * 
     */
    private String path;

    private PreventionJobTriggerInspectJobInspectConfigCustomInfoTypeDictionaryCloudStoragePath() {}
    /**
     * @return A url representing a file or path (no wildcards) in Cloud Storage. Example: `gs://[BUCKET_NAME]/dictionary.txt`
     * 
     */
    public String path() {
        return this.path;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PreventionJobTriggerInspectJobInspectConfigCustomInfoTypeDictionaryCloudStoragePath defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String path;
        public Builder() {}
        public Builder(PreventionJobTriggerInspectJobInspectConfigCustomInfoTypeDictionaryCloudStoragePath defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.path = defaults.path;
        }

        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public PreventionJobTriggerInspectJobInspectConfigCustomInfoTypeDictionaryCloudStoragePath build() {
            final var o = new PreventionJobTriggerInspectJobInspectConfigCustomInfoTypeDictionaryCloudStoragePath();
            o.path = path;
            return o;
        }
    }
}
