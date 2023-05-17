// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.sql.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTiersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTiersArgs Empty = new GetTiersArgs();

    /**
     * The Project ID for which to list tiers. If `project` is not provided, the project defined within the default provider configuration is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The Project ID for which to list tiers. If `project` is not provided, the project defined within the default provider configuration is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private GetTiersArgs() {}

    private GetTiersArgs(GetTiersArgs $) {
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTiersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTiersArgs $;

        public Builder() {
            $ = new GetTiersArgs();
        }

        public Builder(GetTiersArgs defaults) {
            $ = new GetTiersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The Project ID for which to list tiers. If `project` is not provided, the project defined within the default provider configuration is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The Project ID for which to list tiers. If `project` is not provided, the project defined within the default provider configuration is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetTiersArgs build() {
            return $;
        }
    }

}
