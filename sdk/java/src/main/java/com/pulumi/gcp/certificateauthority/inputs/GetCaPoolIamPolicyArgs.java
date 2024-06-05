// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCaPoolIamPolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCaPoolIamPolicyArgs Empty = new GetCaPoolIamPolicyArgs();

    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="caPool", required=true)
    private Output<String> caPool;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> caPool() {
        return this.caPool;
    }

    /**
     * Location of the CaPool. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return Location of the CaPool. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private GetCaPoolIamPolicyArgs() {}

    private GetCaPoolIamPolicyArgs(GetCaPoolIamPolicyArgs $) {
        this.caPool = $.caPool;
        this.location = $.location;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCaPoolIamPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCaPoolIamPolicyArgs $;

        public Builder() {
            $ = new GetCaPoolIamPolicyArgs();
        }

        public Builder(GetCaPoolIamPolicyArgs defaults) {
            $ = new GetCaPoolIamPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param caPool Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder caPool(Output<String> caPool) {
            $.caPool = caPool;
            return this;
        }

        /**
         * @param caPool Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder caPool(String caPool) {
            return caPool(Output.of(caPool));
        }

        /**
         * @param location Location of the CaPool. A full list of valid locations can be found by
         * running `gcloud privateca locations list`.
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Location of the CaPool. A full list of valid locations can be found by
         * running `gcloud privateca locations list`.
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
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
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetCaPoolIamPolicyArgs build() {
            if ($.caPool == null) {
                throw new MissingRequiredPropertyException("GetCaPoolIamPolicyArgs", "caPool");
            }
            return $;
        }
    }

}
