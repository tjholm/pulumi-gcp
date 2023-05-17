// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PreventionJobTriggerInspectJobInspectConfigInfoType {
    /**
     * @return Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * 
     */
    private String name;
    /**
     * @return Version of the information type to use. By default, the version is set to stable.
     * 
     */
    private @Nullable String version;

    private PreventionJobTriggerInspectJobInspectConfigInfoType() {}
    /**
     * @return Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Version of the information type to use. By default, the version is set to stable.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PreventionJobTriggerInspectJobInspectConfigInfoType defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private @Nullable String version;
        public Builder() {}
        public Builder(PreventionJobTriggerInspectJobInspectConfigInfoType defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder version(@Nullable String version) {
            this.version = version;
            return this;
        }
        public PreventionJobTriggerInspectJobInspectConfigInfoType build() {
            final var o = new PreventionJobTriggerInspectJobInspectConfigInfoType();
            o.name = name;
            o.version = version;
            return o;
        }
    }
}
