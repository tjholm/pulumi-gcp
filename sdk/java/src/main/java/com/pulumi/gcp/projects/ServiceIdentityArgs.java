// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.projects;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceIdentityArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceIdentityArgs Empty = new ServiceIdentityArgs();

    /**
     * The email address of the Google managed service account.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return The email address of the Google managed service account.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The service to generate identity for.
     * 
     */
    @Import(name="service", required=true)
    private Output<String> service;

    /**
     * @return The service to generate identity for.
     * 
     */
    public Output<String> service() {
        return this.service;
    }

    private ServiceIdentityArgs() {}

    private ServiceIdentityArgs(ServiceIdentityArgs $) {
        this.email = $.email;
        this.project = $.project;
        this.service = $.service;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceIdentityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceIdentityArgs $;

        public Builder() {
            $ = new ServiceIdentityArgs();
        }

        public Builder(ServiceIdentityArgs defaults) {
            $ = new ServiceIdentityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param email The email address of the Google managed service account.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The email address of the Google managed service account.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param service The service to generate identity for.
         * 
         * @return builder
         * 
         */
        public Builder service(Output<String> service) {
            $.service = service;
            return this;
        }

        /**
         * @param service The service to generate identity for.
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            return service(Output.of(service));
        }

        public ServiceIdentityArgs build() {
            $.service = Objects.requireNonNull($.service, "expected parameter 'service' to be non-null");
            return $;
        }
    }

}
