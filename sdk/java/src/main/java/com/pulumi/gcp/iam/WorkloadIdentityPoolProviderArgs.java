// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iam;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.iam.inputs.WorkloadIdentityPoolProviderAwsArgs;
import com.pulumi.gcp.iam.inputs.WorkloadIdentityPoolProviderOidcArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkloadIdentityPoolProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkloadIdentityPoolProviderArgs Empty = new WorkloadIdentityPoolProviderArgs();

    /**
     * [A Common Expression Language](https://opensource.google/projects/cel) expression, in
     * plain text, to restrict what otherwise valid authentication credentials issued by the
     * provider should not be accepted.
     * The expression must output a boolean representing whether to allow the federation.
     * The following keywords may be referenced in the expressions:
     * * `assertion`: JSON representing the authentication credential issued by the provider.
     * * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
     * * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
     *   The maximum length of the attribute condition expression is 4096 characters. If
     *   unspecified, all valid authentication credential are accepted.
     *   The following example shows how to only allow credentials with a mapped `google.groups`
     *   value of `admins`:
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    @Import(name="attributeCondition")
    private @Nullable Output<String> attributeCondition;

    /**
     * @return [A Common Expression Language](https://opensource.google/projects/cel) expression, in
     * plain text, to restrict what otherwise valid authentication credentials issued by the
     * provider should not be accepted.
     * The expression must output a boolean representing whether to allow the federation.
     * The following keywords may be referenced in the expressions:
     * * `assertion`: JSON representing the authentication credential issued by the provider.
     * * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
     * * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
     *   The maximum length of the attribute condition expression is 4096 characters. If
     *   unspecified, all valid authentication credential are accepted.
     *   The following example shows how to only allow credentials with a mapped `google.groups`
     *   value of `admins`:
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    public Optional<Output<String>> attributeCondition() {
        return Optional.ofNullable(this.attributeCondition);
    }

    /**
     * Maps attributes from authentication credentials issued by an external identity provider
     * to Google Cloud attributes, such as `subject` and `segment`.
     * Each key must be a string specifying the Google Cloud IAM attribute to map to.
     * The following keys are supported:
     * * `google.subject`: The principal IAM is authenticating. You can reference this value
     *   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
     *   Cannot exceed 127 characters.
     * * `google.groups`: Groups the external identity belongs to. You can grant groups
     *   access to resources using an IAM `principalSet` binding; access applies to all
     *   members of the group.
     *   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
     *   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
     *   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
     *   is 100 characters, and the key may only contain the characters [a-z0-9_].
     *   You can reference these attributes in IAM policies to define fine-grained access for a
     *   workload to Google Cloud resources. For example:
     * * `google.subject`:
     *   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
     * * `google.groups`:
     *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
     * * `attribute.{custom_attribute}`:
     *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
     *   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
     *   function that maps an identity provider credential to the normalized attribute specified
     *   by the corresponding map key.
     *   You can use the `assertion` keyword in the expression to access a JSON representation of
     *   the authentication credential issued by the provider.
     *   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
     *   the total size of all mapped attributes must not exceed 8KB.
     *   For AWS providers, the following rules apply:
     * - If no attribute mapping is defined, the following default mapping applies:
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * - If any custom attribute mappings are defined, they must include a mapping to the
     *   `google.subject` attribute.
     *   For OIDC providers, the following rules apply:
     * - Custom attribute mappings must be defined, and must include a mapping to the
     *   `google.subject` attribute. For example, the following maps the `sub` claim of the
     *   incoming credential to the `subject` attribute on a Google token.
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    @Import(name="attributeMapping")
    private @Nullable Output<Map<String,String>> attributeMapping;

    /**
     * @return Maps attributes from authentication credentials issued by an external identity provider
     * to Google Cloud attributes, such as `subject` and `segment`.
     * Each key must be a string specifying the Google Cloud IAM attribute to map to.
     * The following keys are supported:
     * * `google.subject`: The principal IAM is authenticating. You can reference this value
     *   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
     *   Cannot exceed 127 characters.
     * * `google.groups`: Groups the external identity belongs to. You can grant groups
     *   access to resources using an IAM `principalSet` binding; access applies to all
     *   members of the group.
     *   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
     *   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
     *   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
     *   is 100 characters, and the key may only contain the characters [a-z0-9_].
     *   You can reference these attributes in IAM policies to define fine-grained access for a
     *   workload to Google Cloud resources. For example:
     * * `google.subject`:
     *   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
     * * `google.groups`:
     *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
     * * `attribute.{custom_attribute}`:
     *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
     *   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
     *   function that maps an identity provider credential to the normalized attribute specified
     *   by the corresponding map key.
     *   You can use the `assertion` keyword in the expression to access a JSON representation of
     *   the authentication credential issued by the provider.
     *   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
     *   the total size of all mapped attributes must not exceed 8KB.
     *   For AWS providers, the following rules apply:
     * - If no attribute mapping is defined, the following default mapping applies:
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * - If any custom attribute mappings are defined, they must include a mapping to the
     *   `google.subject` attribute.
     *   For OIDC providers, the following rules apply:
     * - Custom attribute mappings must be defined, and must include a mapping to the
     *   `google.subject` attribute. For example, the following maps the `sub` claim of the
     *   incoming credential to the `subject` attribute on a Google token.
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    public Optional<Output<Map<String,String>>> attributeMapping() {
        return Optional.ofNullable(this.attributeMapping);
    }

    /**
     * An Amazon Web Services identity provider. Not compatible with the property oidc.
     * Structure is documented below.
     * 
     */
    @Import(name="aws")
    private @Nullable Output<WorkloadIdentityPoolProviderAwsArgs> aws;

    /**
     * @return An Amazon Web Services identity provider. Not compatible with the property oidc.
     * Structure is documented below.
     * 
     */
    public Optional<Output<WorkloadIdentityPoolProviderAwsArgs>> aws() {
        return Optional.ofNullable(this.aws);
    }

    /**
     * A description for the provider. Cannot exceed 256 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the provider. Cannot exceed 256 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
     * However, existing tokens still grant access.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
     * However, existing tokens still grant access.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * A display name for the provider. Cannot exceed 32 characters.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return A display name for the provider. Cannot exceed 32 characters.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
     * Structure is documented below.
     * 
     */
    @Import(name="oidc")
    private @Nullable Output<WorkloadIdentityPoolProviderOidcArgs> oidc;

    /**
     * @return An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
     * Structure is documented below.
     * 
     */
    public Optional<Output<WorkloadIdentityPoolProviderOidcArgs>> oidc() {
        return Optional.ofNullable(this.oidc);
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
     * The ID used for the pool, which is the final component of the pool resource name. This
     * value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
     * `gcp-` is reserved for use by Google, and may not be specified.
     * 
     */
    @Import(name="workloadIdentityPoolId", required=true)
    private Output<String> workloadIdentityPoolId;

    /**
     * @return The ID used for the pool, which is the final component of the pool resource name. This
     * value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
     * `gcp-` is reserved for use by Google, and may not be specified.
     * 
     */
    public Output<String> workloadIdentityPoolId() {
        return this.workloadIdentityPoolId;
    }

    /**
     * The ID for the provider, which becomes the final component of the resource name. This
     * value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
     * `gcp-` is reserved for use by Google, and may not be specified.
     * 
     */
    @Import(name="workloadIdentityPoolProviderId", required=true)
    private Output<String> workloadIdentityPoolProviderId;

    /**
     * @return The ID for the provider, which becomes the final component of the resource name. This
     * value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
     * `gcp-` is reserved for use by Google, and may not be specified.
     * 
     */
    public Output<String> workloadIdentityPoolProviderId() {
        return this.workloadIdentityPoolProviderId;
    }

    private WorkloadIdentityPoolProviderArgs() {}

    private WorkloadIdentityPoolProviderArgs(WorkloadIdentityPoolProviderArgs $) {
        this.attributeCondition = $.attributeCondition;
        this.attributeMapping = $.attributeMapping;
        this.aws = $.aws;
        this.description = $.description;
        this.disabled = $.disabled;
        this.displayName = $.displayName;
        this.oidc = $.oidc;
        this.project = $.project;
        this.workloadIdentityPoolId = $.workloadIdentityPoolId;
        this.workloadIdentityPoolProviderId = $.workloadIdentityPoolProviderId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkloadIdentityPoolProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkloadIdentityPoolProviderArgs $;

        public Builder() {
            $ = new WorkloadIdentityPoolProviderArgs();
        }

        public Builder(WorkloadIdentityPoolProviderArgs defaults) {
            $ = new WorkloadIdentityPoolProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributeCondition [A Common Expression Language](https://opensource.google/projects/cel) expression, in
         * plain text, to restrict what otherwise valid authentication credentials issued by the
         * provider should not be accepted.
         * The expression must output a boolean representing whether to allow the federation.
         * The following keywords may be referenced in the expressions:
         * * `assertion`: JSON representing the authentication credential issued by the provider.
         * * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
         * * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
         *   The maximum length of the attribute condition expression is 4096 characters. If
         *   unspecified, all valid authentication credential are accepted.
         *   The following example shows how to only allow credentials with a mapped `google.groups`
         *   value of `admins`:
         * ```java
         * package generated_program;
         * 
         * import java.util.*;
         * import java.io.*;
         * import java.nio.*;
         * import com.pulumi.*;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder attributeCondition(@Nullable Output<String> attributeCondition) {
            $.attributeCondition = attributeCondition;
            return this;
        }

        /**
         * @param attributeCondition [A Common Expression Language](https://opensource.google/projects/cel) expression, in
         * plain text, to restrict what otherwise valid authentication credentials issued by the
         * provider should not be accepted.
         * The expression must output a boolean representing whether to allow the federation.
         * The following keywords may be referenced in the expressions:
         * * `assertion`: JSON representing the authentication credential issued by the provider.
         * * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
         * * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
         *   The maximum length of the attribute condition expression is 4096 characters. If
         *   unspecified, all valid authentication credential are accepted.
         *   The following example shows how to only allow credentials with a mapped `google.groups`
         *   value of `admins`:
         * ```java
         * package generated_program;
         * 
         * import java.util.*;
         * import java.io.*;
         * import java.nio.*;
         * import com.pulumi.*;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder attributeCondition(String attributeCondition) {
            return attributeCondition(Output.of(attributeCondition));
        }

        /**
         * @param attributeMapping Maps attributes from authentication credentials issued by an external identity provider
         * to Google Cloud attributes, such as `subject` and `segment`.
         * Each key must be a string specifying the Google Cloud IAM attribute to map to.
         * The following keys are supported:
         * * `google.subject`: The principal IAM is authenticating. You can reference this value
         *   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
         *   Cannot exceed 127 characters.
         * * `google.groups`: Groups the external identity belongs to. You can grant groups
         *   access to resources using an IAM `principalSet` binding; access applies to all
         *   members of the group.
         *   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
         *   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
         *   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
         *   is 100 characters, and the key may only contain the characters [a-z0-9_].
         *   You can reference these attributes in IAM policies to define fine-grained access for a
         *   workload to Google Cloud resources. For example:
         * * `google.subject`:
         *   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
         * * `google.groups`:
         *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
         * * `attribute.{custom_attribute}`:
         *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
         *   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
         *   function that maps an identity provider credential to the normalized attribute specified
         *   by the corresponding map key.
         *   You can use the `assertion` keyword in the expression to access a JSON representation of
         *   the authentication credential issued by the provider.
         *   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
         *   the total size of all mapped attributes must not exceed 8KB.
         *   For AWS providers, the following rules apply:
         * - If no attribute mapping is defined, the following default mapping applies:
         * ```java
         * package generated_program;
         * 
         * import java.util.*;
         * import java.io.*;
         * import java.nio.*;
         * import com.pulumi.*;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * - If any custom attribute mappings are defined, they must include a mapping to the
         *   `google.subject` attribute.
         *   For OIDC providers, the following rules apply:
         * - Custom attribute mappings must be defined, and must include a mapping to the
         *   `google.subject` attribute. For example, the following maps the `sub` claim of the
         *   incoming credential to the `subject` attribute on a Google token.
         * ```java
         * package generated_program;
         * 
         * import java.util.*;
         * import java.io.*;
         * import java.nio.*;
         * import com.pulumi.*;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder attributeMapping(@Nullable Output<Map<String,String>> attributeMapping) {
            $.attributeMapping = attributeMapping;
            return this;
        }

        /**
         * @param attributeMapping Maps attributes from authentication credentials issued by an external identity provider
         * to Google Cloud attributes, such as `subject` and `segment`.
         * Each key must be a string specifying the Google Cloud IAM attribute to map to.
         * The following keys are supported:
         * * `google.subject`: The principal IAM is authenticating. You can reference this value
         *   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
         *   Cannot exceed 127 characters.
         * * `google.groups`: Groups the external identity belongs to. You can grant groups
         *   access to resources using an IAM `principalSet` binding; access applies to all
         *   members of the group.
         *   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
         *   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
         *   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
         *   is 100 characters, and the key may only contain the characters [a-z0-9_].
         *   You can reference these attributes in IAM policies to define fine-grained access for a
         *   workload to Google Cloud resources. For example:
         * * `google.subject`:
         *   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
         * * `google.groups`:
         *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
         * * `attribute.{custom_attribute}`:
         *   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
         *   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
         *   function that maps an identity provider credential to the normalized attribute specified
         *   by the corresponding map key.
         *   You can use the `assertion` keyword in the expression to access a JSON representation of
         *   the authentication credential issued by the provider.
         *   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
         *   the total size of all mapped attributes must not exceed 8KB.
         *   For AWS providers, the following rules apply:
         * - If no attribute mapping is defined, the following default mapping applies:
         * ```java
         * package generated_program;
         * 
         * import java.util.*;
         * import java.io.*;
         * import java.nio.*;
         * import com.pulumi.*;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * - If any custom attribute mappings are defined, they must include a mapping to the
         *   `google.subject` attribute.
         *   For OIDC providers, the following rules apply:
         * - Custom attribute mappings must be defined, and must include a mapping to the
         *   `google.subject` attribute. For example, the following maps the `sub` claim of the
         *   incoming credential to the `subject` attribute on a Google token.
         * ```java
         * package generated_program;
         * 
         * import java.util.*;
         * import java.io.*;
         * import java.nio.*;
         * import com.pulumi.*;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder attributeMapping(Map<String,String> attributeMapping) {
            return attributeMapping(Output.of(attributeMapping));
        }

        /**
         * @param aws An Amazon Web Services identity provider. Not compatible with the property oidc.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder aws(@Nullable Output<WorkloadIdentityPoolProviderAwsArgs> aws) {
            $.aws = aws;
            return this;
        }

        /**
         * @param aws An Amazon Web Services identity provider. Not compatible with the property oidc.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder aws(WorkloadIdentityPoolProviderAwsArgs aws) {
            return aws(Output.of(aws));
        }

        /**
         * @param description A description for the provider. Cannot exceed 256 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the provider. Cannot exceed 256 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disabled Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
         * However, existing tokens still grant access.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
         * However, existing tokens still grant access.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param displayName A display name for the provider. Cannot exceed 32 characters.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName A display name for the provider. Cannot exceed 32 characters.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param oidc An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder oidc(@Nullable Output<WorkloadIdentityPoolProviderOidcArgs> oidc) {
            $.oidc = oidc;
            return this;
        }

        /**
         * @param oidc An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder oidc(WorkloadIdentityPoolProviderOidcArgs oidc) {
            return oidc(Output.of(oidc));
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
         * @param workloadIdentityPoolId The ID used for the pool, which is the final component of the pool resource name. This
         * value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
         * `gcp-` is reserved for use by Google, and may not be specified.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityPoolId(Output<String> workloadIdentityPoolId) {
            $.workloadIdentityPoolId = workloadIdentityPoolId;
            return this;
        }

        /**
         * @param workloadIdentityPoolId The ID used for the pool, which is the final component of the pool resource name. This
         * value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
         * `gcp-` is reserved for use by Google, and may not be specified.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityPoolId(String workloadIdentityPoolId) {
            return workloadIdentityPoolId(Output.of(workloadIdentityPoolId));
        }

        /**
         * @param workloadIdentityPoolProviderId The ID for the provider, which becomes the final component of the resource name. This
         * value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
         * `gcp-` is reserved for use by Google, and may not be specified.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityPoolProviderId(Output<String> workloadIdentityPoolProviderId) {
            $.workloadIdentityPoolProviderId = workloadIdentityPoolProviderId;
            return this;
        }

        /**
         * @param workloadIdentityPoolProviderId The ID for the provider, which becomes the final component of the resource name. This
         * value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
         * `gcp-` is reserved for use by Google, and may not be specified.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityPoolProviderId(String workloadIdentityPoolProviderId) {
            return workloadIdentityPoolProviderId(Output.of(workloadIdentityPoolProviderId));
        }

        public WorkloadIdentityPoolProviderArgs build() {
            $.workloadIdentityPoolId = Objects.requireNonNull($.workloadIdentityPoolId, "expected parameter 'workloadIdentityPoolId' to be non-null");
            $.workloadIdentityPoolProviderId = Objects.requireNonNull($.workloadIdentityPoolProviderId, "expected parameter 'workloadIdentityPoolProviderId' to be non-null");
            return $;
        }
    }

}
