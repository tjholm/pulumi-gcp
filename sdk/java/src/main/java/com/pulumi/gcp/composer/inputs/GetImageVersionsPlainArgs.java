// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.composer.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetImageVersionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetImageVersionsPlainArgs Empty = new GetImageVersionsPlainArgs();

    /**
     * The ID of the project to list versions in.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable String project;

    /**
     * @return The ID of the project to list versions in.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The location to list versions in.
     * If it is not provider, the provider region is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The location to list versions in.
     * If it is not provider, the provider region is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetImageVersionsPlainArgs() {}

    private GetImageVersionsPlainArgs(GetImageVersionsPlainArgs $) {
        this.project = $.project;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetImageVersionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetImageVersionsPlainArgs $;

        public Builder() {
            $ = new GetImageVersionsPlainArgs();
        }

        public Builder(GetImageVersionsPlainArgs defaults) {
            $ = new GetImageVersionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The ID of the project to list versions in.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable String project) {
            $.project = project;
            return this;
        }

        /**
         * @param region The location to list versions in.
         * If it is not provider, the provider region is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetImageVersionsPlainArgs build() {
            return $;
        }
    }

}
