// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctionsv2.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.cloudfunctionsv2.outputs.FunctionBuildConfigSourceRepoSource;
import com.pulumi.gcp.cloudfunctionsv2.outputs.FunctionBuildConfigSourceStorageSource;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FunctionBuildConfigSource {
    /**
     * @return If provided, get the source from this location in a Cloud Source Repository.
     * Structure is documented below.
     * 
     */
    private final @Nullable FunctionBuildConfigSourceRepoSource repoSource;
    /**
     * @return If provided, get the source from this location in Google Cloud Storage.
     * Structure is documented below.
     * 
     */
    private final @Nullable FunctionBuildConfigSourceStorageSource storageSource;

    @CustomType.Constructor
    private FunctionBuildConfigSource(
        @CustomType.Parameter("repoSource") @Nullable FunctionBuildConfigSourceRepoSource repoSource,
        @CustomType.Parameter("storageSource") @Nullable FunctionBuildConfigSourceStorageSource storageSource) {
        this.repoSource = repoSource;
        this.storageSource = storageSource;
    }

    /**
     * @return If provided, get the source from this location in a Cloud Source Repository.
     * Structure is documented below.
     * 
     */
    public Optional<FunctionBuildConfigSourceRepoSource> repoSource() {
        return Optional.ofNullable(this.repoSource);
    }
    /**
     * @return If provided, get the source from this location in Google Cloud Storage.
     * Structure is documented below.
     * 
     */
    public Optional<FunctionBuildConfigSourceStorageSource> storageSource() {
        return Optional.ofNullable(this.storageSource);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FunctionBuildConfigSource defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable FunctionBuildConfigSourceRepoSource repoSource;
        private @Nullable FunctionBuildConfigSourceStorageSource storageSource;

        public Builder() {
    	      // Empty
        }

        public Builder(FunctionBuildConfigSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.repoSource = defaults.repoSource;
    	      this.storageSource = defaults.storageSource;
        }

        public Builder repoSource(@Nullable FunctionBuildConfigSourceRepoSource repoSource) {
            this.repoSource = repoSource;
            return this;
        }
        public Builder storageSource(@Nullable FunctionBuildConfigSourceStorageSource storageSource) {
            this.storageSource = storageSource;
            return this;
        }        public FunctionBuildConfigSource build() {
            return new FunctionBuildConfigSource(repoSource, storageSource);
        }
    }
}
