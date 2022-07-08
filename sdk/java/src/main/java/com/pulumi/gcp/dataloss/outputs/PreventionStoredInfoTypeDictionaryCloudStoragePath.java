// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class PreventionStoredInfoTypeDictionaryCloudStoragePath {
    /**
     * @return A url representing a file or path (no wildcards) in Cloud Storage. Example: `gs://[BUCKET_NAME]/dictionary.txt`
     * 
     */
    private final String path;

    @CustomType.Constructor
    private PreventionStoredInfoTypeDictionaryCloudStoragePath(@CustomType.Parameter("path") String path) {
        this.path = path;
    }

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

    public static Builder builder(PreventionStoredInfoTypeDictionaryCloudStoragePath defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String path;

        public Builder() {
    	      // Empty
        }

        public Builder(PreventionStoredInfoTypeDictionaryCloudStoragePath defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.path = defaults.path;
        }

        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }        public PreventionStoredInfoTypeDictionaryCloudStoragePath build() {
            return new PreventionStoredInfoTypeDictionaryCloudStoragePath(path);
        }
    }
}
