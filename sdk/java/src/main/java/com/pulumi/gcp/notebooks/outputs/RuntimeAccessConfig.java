// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RuntimeAccessConfig {
    /**
     * @return The type of access mode this instance. For valid values, see
     * `https://cloud.google.com/vertex-ai/docs/workbench/reference/
     * rest/v1/projects.locations.runtimes#RuntimeAccessType`.
     * 
     */
    private final @Nullable String accessType;
    /**
     * @return -
     * The proxy endpoint that is used to access the runtime.
     * 
     */
    private final @Nullable String proxyUri;
    /**
     * @return The owner of this runtime after creation. Format: `alias@example.com`.
     * Currently supports one owner only.
     * 
     */
    private final @Nullable String runtimeOwner;

    @CustomType.Constructor
    private RuntimeAccessConfig(
        @CustomType.Parameter("accessType") @Nullable String accessType,
        @CustomType.Parameter("proxyUri") @Nullable String proxyUri,
        @CustomType.Parameter("runtimeOwner") @Nullable String runtimeOwner) {
        this.accessType = accessType;
        this.proxyUri = proxyUri;
        this.runtimeOwner = runtimeOwner;
    }

    /**
     * @return The type of access mode this instance. For valid values, see
     * `https://cloud.google.com/vertex-ai/docs/workbench/reference/
     * rest/v1/projects.locations.runtimes#RuntimeAccessType`.
     * 
     */
    public Optional<String> accessType() {
        return Optional.ofNullable(this.accessType);
    }
    /**
     * @return -
     * The proxy endpoint that is used to access the runtime.
     * 
     */
    public Optional<String> proxyUri() {
        return Optional.ofNullable(this.proxyUri);
    }
    /**
     * @return The owner of this runtime after creation. Format: `alias@example.com`.
     * Currently supports one owner only.
     * 
     */
    public Optional<String> runtimeOwner() {
        return Optional.ofNullable(this.runtimeOwner);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuntimeAccessConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String accessType;
        private @Nullable String proxyUri;
        private @Nullable String runtimeOwner;

        public Builder() {
    	      // Empty
        }

        public Builder(RuntimeAccessConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessType = defaults.accessType;
    	      this.proxyUri = defaults.proxyUri;
    	      this.runtimeOwner = defaults.runtimeOwner;
        }

        public Builder accessType(@Nullable String accessType) {
            this.accessType = accessType;
            return this;
        }
        public Builder proxyUri(@Nullable String proxyUri) {
            this.proxyUri = proxyUri;
            return this;
        }
        public Builder runtimeOwner(@Nullable String runtimeOwner) {
            this.runtimeOwner = runtimeOwner;
            return this;
        }        public RuntimeAccessConfig build() {
            return new RuntimeAccessConfig(accessType, proxyUri, runtimeOwner);
        }
    }
}
