// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.diagflow.CxWebhookArgs;
import com.pulumi.gcp.diagflow.inputs.CxWebhookState;
import com.pulumi.gcp.diagflow.outputs.CxWebhookGenericWebService;
import com.pulumi.gcp.diagflow.outputs.CxWebhookServiceDirectory;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Webhooks host the developer&#39;s business logic. During a session, webhooks allow the developer to use the data extracted by Dialogflow&#39;s natural language processing to generate dynamic responses, validate collected data, or trigger actions on the backend.
 * 
 * To get more information about Webhook, see:
 * 
 * * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.webhooks)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
 * 
 * ## Example Usage
 * 
 * ### Dialogflowcx Webhook Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.diagflow.CxAgent;
 * import com.pulumi.gcp.diagflow.CxAgentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxAgentSpeechToTextSettingsArgs;
 * import com.pulumi.gcp.diagflow.CxWebhook;
 * import com.pulumi.gcp.diagflow.CxWebhookArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxWebhookGenericWebServiceArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var agent = new CxAgent("agent", CxAgentArgs.builder()
 *             .displayName("dialogflowcx-agent")
 *             .location("global")
 *             .defaultLanguageCode("en")
 *             .supportedLanguageCodes(            
 *                 "it",
 *                 "de",
 *                 "es")
 *             .timeZone("America/New_York")
 *             .description("Example description.")
 *             .avatarUri("https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png")
 *             .enableStackdriverLogging(true)
 *             .enableSpellCorrection(true)
 *             .speechToTextSettings(CxAgentSpeechToTextSettingsArgs.builder()
 *                 .enableSpeechAdaptation(true)
 *                 .build())
 *             .build());
 * 
 *         var basicWebhook = new CxWebhook("basicWebhook", CxWebhookArgs.builder()
 *             .parent(agent.id())
 *             .displayName("MyFlow")
 *             .genericWebService(CxWebhookGenericWebServiceArgs.builder()
 *                 .uri("https://example.com")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Webhook can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/webhooks/{{name}}`
 * 
 * * `{{parent}}/{{name}}`
 * 
 * When using the `pulumi import` command, Webhook can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:diagflow/cxWebhook:CxWebhook default {{parent}}/webhooks/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:diagflow/cxWebhook:CxWebhook default {{parent}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:diagflow/cxWebhook:CxWebhook")
public class CxWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Indicates whether the webhook is disabled.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return Indicates whether the webhook is disabled.
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    /**
     * The human-readable name of the webhook, unique within the agent.
     * 
     * ***
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The human-readable name of the webhook, unique within the agent.
     * 
     * ***
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Indicates if automatic spell correction is enabled in detect intent requests.
     * 
     */
    @Export(name="enableSpellCorrection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableSpellCorrection;

    /**
     * @return Indicates if automatic spell correction is enabled in detect intent requests.
     * 
     */
    public Output<Optional<Boolean>> enableSpellCorrection() {
        return Codegen.optional(this.enableSpellCorrection);
    }
    /**
     * Determines whether this agent should log conversation queries.
     * 
     */
    @Export(name="enableStackdriverLogging", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableStackdriverLogging;

    /**
     * @return Determines whether this agent should log conversation queries.
     * 
     */
    public Output<Optional<Boolean>> enableStackdriverLogging() {
        return Codegen.optional(this.enableStackdriverLogging);
    }
    /**
     * Configuration for a generic web service.
     * Structure is documented below.
     * 
     */
    @Export(name="genericWebService", refs={CxWebhookGenericWebService.class}, tree="[0]")
    private Output</* @Nullable */ CxWebhookGenericWebService> genericWebService;

    /**
     * @return Configuration for a generic web service.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CxWebhookGenericWebService>> genericWebService() {
        return Codegen.optional(this.genericWebService);
    }
    /**
     * The unique identifier of the webhook.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/webhooks/&lt;Webhook ID&gt;.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique identifier of the webhook.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/webhooks/&lt;Webhook ID&gt;.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The agent to create a webhook for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parent;

    /**
     * @return The agent to create a webhook for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    public Output<Optional<String>> parent() {
        return Codegen.optional(this.parent);
    }
    /**
     * Name of the SecuritySettings reference for the agent. Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/securitySettings/&lt;Security Settings ID&gt;.
     * 
     */
    @Export(name="securitySettings", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securitySettings;

    /**
     * @return Name of the SecuritySettings reference for the agent. Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/securitySettings/&lt;Security Settings ID&gt;.
     * 
     */
    public Output<Optional<String>> securitySettings() {
        return Codegen.optional(this.securitySettings);
    }
    /**
     * Configuration for a Service Directory service.
     * Structure is documented below.
     * 
     */
    @Export(name="serviceDirectory", refs={CxWebhookServiceDirectory.class}, tree="[0]")
    private Output</* @Nullable */ CxWebhookServiceDirectory> serviceDirectory;

    /**
     * @return Configuration for a Service Directory service.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CxWebhookServiceDirectory>> serviceDirectory() {
        return Codegen.optional(this.serviceDirectory);
    }
    /**
     * Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;.
     * 
     */
    @Export(name="startFlow", refs={String.class}, tree="[0]")
    private Output<String> startFlow;

    /**
     * @return Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;.
     * 
     */
    public Output<String> startFlow() {
        return this.startFlow;
    }
    /**
     * Webhook execution timeout.
     * 
     */
    @Export(name="timeout", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timeout;

    /**
     * @return Webhook execution timeout.
     * 
     */
    public Output<Optional<String>> timeout() {
        return Codegen.optional(this.timeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CxWebhook(java.lang.String name) {
        this(name, CxWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CxWebhook(java.lang.String name, CxWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CxWebhook(java.lang.String name, CxWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/cxWebhook:CxWebhook", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CxWebhook(java.lang.String name, Output<java.lang.String> id, @Nullable CxWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/cxWebhook:CxWebhook", name, state, makeResourceOptions(options, id), false);
    }

    private static CxWebhookArgs makeArgs(CxWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CxWebhookArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CxWebhook get(java.lang.String name, Output<java.lang.String> id, @Nullable CxWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CxWebhook(name, id, state, options);
    }
}
