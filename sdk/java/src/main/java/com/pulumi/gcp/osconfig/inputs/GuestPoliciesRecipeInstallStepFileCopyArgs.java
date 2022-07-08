// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.osconfig.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GuestPoliciesRecipeInstallStepFileCopyArgs extends com.pulumi.resources.ResourceArgs {

    public static final GuestPoliciesRecipeInstallStepFileCopyArgs Empty = new GuestPoliciesRecipeInstallStepFileCopyArgs();

    /**
     * The id of the relevant artifact in the recipe.
     * 
     */
    @Import(name="artifactId", required=true)
    private Output<String> artifactId;

    /**
     * @return The id of the relevant artifact in the recipe.
     * 
     */
    public Output<String> artifactId() {
        return this.artifactId;
    }

    /**
     * Directory to extract archive to. Defaults to / on Linux or C:\ on Windows.
     * 
     */
    @Import(name="destination", required=true)
    private Output<String> destination;

    /**
     * @return Directory to extract archive to. Defaults to / on Linux or C:\ on Windows.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }

    /**
     * Whether to allow this step to overwrite existing files.If this is false and the file already exists the file
     * is not overwritten and the step is considered a success. Defaults to false.
     * 
     */
    @Import(name="overwrite")
    private @Nullable Output<Boolean> overwrite;

    /**
     * @return Whether to allow this step to overwrite existing files.If this is false and the file already exists the file
     * is not overwritten and the step is considered a success. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> overwrite() {
        return Optional.ofNullable(this.overwrite);
    }

    /**
     * Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users
     * for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit
     * number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one
     * bit corresponds to the execute permission. Default behavior is 755.
     * Below are some examples of permissions and their associated values:
     * read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<String> permissions;

    /**
     * @return Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users
     * for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit
     * number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one
     * bit corresponds to the execute permission. Default behavior is 755.
     * Below are some examples of permissions and their associated values:
     * read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4
     * 
     */
    public Optional<Output<String>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    private GuestPoliciesRecipeInstallStepFileCopyArgs() {}

    private GuestPoliciesRecipeInstallStepFileCopyArgs(GuestPoliciesRecipeInstallStepFileCopyArgs $) {
        this.artifactId = $.artifactId;
        this.destination = $.destination;
        this.overwrite = $.overwrite;
        this.permissions = $.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GuestPoliciesRecipeInstallStepFileCopyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GuestPoliciesRecipeInstallStepFileCopyArgs $;

        public Builder() {
            $ = new GuestPoliciesRecipeInstallStepFileCopyArgs();
        }

        public Builder(GuestPoliciesRecipeInstallStepFileCopyArgs defaults) {
            $ = new GuestPoliciesRecipeInstallStepFileCopyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param artifactId The id of the relevant artifact in the recipe.
         * 
         * @return builder
         * 
         */
        public Builder artifactId(Output<String> artifactId) {
            $.artifactId = artifactId;
            return this;
        }

        /**
         * @param artifactId The id of the relevant artifact in the recipe.
         * 
         * @return builder
         * 
         */
        public Builder artifactId(String artifactId) {
            return artifactId(Output.of(artifactId));
        }

        /**
         * @param destination Directory to extract archive to. Defaults to / on Linux or C:\ on Windows.
         * 
         * @return builder
         * 
         */
        public Builder destination(Output<String> destination) {
            $.destination = destination;
            return this;
        }

        /**
         * @param destination Directory to extract archive to. Defaults to / on Linux or C:\ on Windows.
         * 
         * @return builder
         * 
         */
        public Builder destination(String destination) {
            return destination(Output.of(destination));
        }

        /**
         * @param overwrite Whether to allow this step to overwrite existing files.If this is false and the file already exists the file
         * is not overwritten and the step is considered a success. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder overwrite(@Nullable Output<Boolean> overwrite) {
            $.overwrite = overwrite;
            return this;
        }

        /**
         * @param overwrite Whether to allow this step to overwrite existing files.If this is false and the file already exists the file
         * is not overwritten and the step is considered a success. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder overwrite(Boolean overwrite) {
            return overwrite(Output.of(overwrite));
        }

        /**
         * @param permissions Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users
         * for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit
         * number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one
         * bit corresponds to the execute permission. Default behavior is 755.
         * Below are some examples of permissions and their associated values:
         * read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<String> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users
         * for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit
         * number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one
         * bit corresponds to the execute permission. Default behavior is 755.
         * Below are some examples of permissions and their associated values:
         * read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4
         * 
         * @return builder
         * 
         */
        public Builder permissions(String permissions) {
            return permissions(Output.of(permissions));
        }

        public GuestPoliciesRecipeInstallStepFileCopyArgs build() {
            $.artifactId = Objects.requireNonNull($.artifactId, "expected parameter 'artifactId' to be non-null");
            $.destination = Objects.requireNonNull($.destination, "expected parameter 'destination' to be non-null");
            return $;
        }
    }

}
