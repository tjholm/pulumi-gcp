// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.oslogin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.oslogin.SshPublicKeyArgs;
import com.pulumi.gcp.oslogin.inputs.SshPublicKeyState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The SSH public key information associated with a Google account.
 * 
 * To get more information about SSHPublicKey, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/oslogin/rest/v1/users.sshPublicKeys)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/oslogin)
 * 
 * ## Example Usage
 * ### Os Login Ssh Key Basic
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
 *         final var me = Output.of(OrganizationsFunctions.getClientOpenIdUserInfo());
 * 
 *         var cache = new SshPublicKey(&#34;cache&#34;, SshPublicKeyArgs.builder()        
 *             .user(me.apply(getClientOpenIdUserInfoResult -&gt; getClientOpenIdUserInfoResult.email()))
 *             .key(Files.readString(&#34;path/to/id_rsa.pub&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SSHPublicKey can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default users/{{user}}/sshPublicKeys/{{fingerprint}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default {{user}}/{{fingerprint}}
 * ```
 * 
 */
@ResourceType(type="gcp:oslogin/sshPublicKey:SshPublicKey")
public class SshPublicKey extends com.pulumi.resources.CustomResource {
    /**
     * An expiration time in microseconds since epoch.
     * 
     */
    @Export(name="expirationTimeUsec", type=String.class, parameters={})
    private Output</* @Nullable */ String> expirationTimeUsec;

    /**
     * @return An expiration time in microseconds since epoch.
     * 
     */
    public Output<Optional<String>> expirationTimeUsec() {
        return Codegen.optional(this.expirationTimeUsec);
    }
    /**
     * The SHA-256 fingerprint of the SSH public key.
     * 
     */
    @Export(name="fingerprint", type=String.class, parameters={})
    private Output<String> fingerprint;

    /**
     * @return The SHA-256 fingerprint of the SSH public key.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * Public key text in SSH format, defined by RFC4253 section 6.6.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return Public key text in SSH format, defined by RFC4253 section 6.6.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The project ID of the Google Cloud Platform project.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output</* @Nullable */ String> project;

    /**
     * @return The project ID of the Google Cloud Platform project.
     * 
     */
    public Output<Optional<String>> project() {
        return Codegen.optional(this.project);
    }
    /**
     * The user email.
     * 
     */
    @Export(name="user", type=String.class, parameters={})
    private Output<String> user;

    /**
     * @return The user email.
     * 
     */
    public Output<String> user() {
        return this.user;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SshPublicKey(String name) {
        this(name, SshPublicKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SshPublicKey(String name, SshPublicKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SshPublicKey(String name, SshPublicKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:oslogin/sshPublicKey:SshPublicKey", name, args == null ? SshPublicKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SshPublicKey(String name, Output<String> id, @Nullable SshPublicKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:oslogin/sshPublicKey:SshPublicKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static SshPublicKey get(String name, Output<String> id, @Nullable SshPublicKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SshPublicKey(name, id, state, options);
    }
}
