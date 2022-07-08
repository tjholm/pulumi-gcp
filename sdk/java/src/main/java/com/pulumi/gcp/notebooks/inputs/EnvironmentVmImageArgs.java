// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnvironmentVmImageArgs extends com.pulumi.resources.ResourceArgs {

    public static final EnvironmentVmImageArgs Empty = new EnvironmentVmImageArgs();

    /**
     * Use this VM image family to find the image; the newest image in this family will be used.
     * 
     */
    @Import(name="imageFamily")
    private @Nullable Output<String> imageFamily;

    /**
     * @return Use this VM image family to find the image; the newest image in this family will be used.
     * 
     */
    public Optional<Output<String>> imageFamily() {
        return Optional.ofNullable(this.imageFamily);
    }

    /**
     * Use VM image name to find the image.
     * 
     */
    @Import(name="imageName")
    private @Nullable Output<String> imageName;

    /**
     * @return Use VM image name to find the image.
     * 
     */
    public Optional<Output<String>> imageName() {
        return Optional.ofNullable(this.imageName);
    }

    /**
     * The name of the Google Cloud project that this VM image belongs to.
     * Format: projects/{project_id}
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The name of the Google Cloud project that this VM image belongs to.
     * Format: projects/{project_id}
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    private EnvironmentVmImageArgs() {}

    private EnvironmentVmImageArgs(EnvironmentVmImageArgs $) {
        this.imageFamily = $.imageFamily;
        this.imageName = $.imageName;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnvironmentVmImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnvironmentVmImageArgs $;

        public Builder() {
            $ = new EnvironmentVmImageArgs();
        }

        public Builder(EnvironmentVmImageArgs defaults) {
            $ = new EnvironmentVmImageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param imageFamily Use this VM image family to find the image; the newest image in this family will be used.
         * 
         * @return builder
         * 
         */
        public Builder imageFamily(@Nullable Output<String> imageFamily) {
            $.imageFamily = imageFamily;
            return this;
        }

        /**
         * @param imageFamily Use this VM image family to find the image; the newest image in this family will be used.
         * 
         * @return builder
         * 
         */
        public Builder imageFamily(String imageFamily) {
            return imageFamily(Output.of(imageFamily));
        }

        /**
         * @param imageName Use VM image name to find the image.
         * 
         * @return builder
         * 
         */
        public Builder imageName(@Nullable Output<String> imageName) {
            $.imageName = imageName;
            return this;
        }

        /**
         * @param imageName Use VM image name to find the image.
         * 
         * @return builder
         * 
         */
        public Builder imageName(String imageName) {
            return imageName(Output.of(imageName));
        }

        /**
         * @param project The name of the Google Cloud project that this VM image belongs to.
         * Format: projects/{project_id}
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name of the Google Cloud project that this VM image belongs to.
         * Format: projects/{project_id}
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public EnvironmentVmImageArgs build() {
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            return $;
        }
    }

}
