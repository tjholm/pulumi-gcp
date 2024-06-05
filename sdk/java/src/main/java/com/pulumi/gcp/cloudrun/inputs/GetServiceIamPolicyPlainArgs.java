// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrun.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceIamPolicyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceIamPolicyPlainArgs Empty = new GetServiceIamPolicyPlainArgs();

    /**
     * The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="location")
    private @Nullable String location;

    /**
     * @return The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<String> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable String project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="service", required=true)
    private String service;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public String service() {
        return this.service;
    }

    private GetServiceIamPolicyPlainArgs() {}

    private GetServiceIamPolicyPlainArgs(GetServiceIamPolicyPlainArgs $) {
        this.location = $.location;
        this.project = $.project;
        this.service = $.service;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceIamPolicyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceIamPolicyPlainArgs $;

        public Builder() {
            $ = new GetServiceIamPolicyPlainArgs();
        }

        public Builder(GetServiceIamPolicyPlainArgs defaults) {
            $ = new GetServiceIamPolicyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param location The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable String location) {
            $.location = location;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable String project) {
            $.project = project;
            return this;
        }

        /**
         * @param service Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            $.service = service;
            return this;
        }

        public GetServiceIamPolicyPlainArgs build() {
            if ($.service == null) {
                throw new MissingRequiredPropertyException("GetServiceIamPolicyPlainArgs", "service");
            }
            return $;
        }
    }

}
