// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class JobSparksqlConfigLoggingConfig {
    private final Map<String,String> driverLogLevels;

    @CustomType.Constructor
    private JobSparksqlConfigLoggingConfig(@CustomType.Parameter("driverLogLevels") Map<String,String> driverLogLevels) {
        this.driverLogLevels = driverLogLevels;
    }

    public Map<String,String> driverLogLevels() {
        return this.driverLogLevels;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(JobSparksqlConfigLoggingConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Map<String,String> driverLogLevels;

        public Builder() {
    	      // Empty
        }

        public Builder(JobSparksqlConfigLoggingConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.driverLogLevels = defaults.driverLogLevels;
        }

        public Builder driverLogLevels(Map<String,String> driverLogLevels) {
            this.driverLogLevels = Objects.requireNonNull(driverLogLevels);
            return this;
        }        public JobSparksqlConfigLoggingConfig build() {
            return new JobSparksqlConfigLoggingConfig(driverLogLevels);
        }
    }
}
