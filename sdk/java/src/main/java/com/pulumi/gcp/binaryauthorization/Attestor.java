// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.binaryauthorization;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.binaryauthorization.AttestorArgs;
import com.pulumi.gcp.binaryauthorization.inputs.AttestorState;
import com.pulumi.gcp.binaryauthorization.outputs.AttestorAttestationAuthorityNote;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An attestor that attests to container image artifacts.
 * 
 * To get more information about Attestor, see:
 * 
 * * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/binary-authorization/)
 * 
 * ## Example Usage
 * 
 * ### Binary Authorization Attestor Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.containeranalysis.Note;
 * import com.pulumi.gcp.containeranalysis.NoteArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityHintArgs;
 * import com.pulumi.gcp.binaryauthorization.Attestor;
 * import com.pulumi.gcp.binaryauthorization.AttestorArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.AttestorAttestationAuthorityNoteArgs;
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
 *         var note = new Note("note", NoteArgs.builder()
 *             .name("test-attestor-note")
 *             .attestationAuthority(NoteAttestationAuthorityArgs.builder()
 *                 .hint(NoteAttestationAuthorityHintArgs.builder()
 *                     .humanReadableName("Attestor Note")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var attestor = new Attestor("attestor", AttestorArgs.builder()
 *             .name("test-attestor")
 *             .attestationAuthorityNote(AttestorAttestationAuthorityNoteArgs.builder()
 *                 .noteReference(note.name())
 *                 .publicKeys(AttestorAttestationAuthorityNotePublicKeyArgs.builder()
 *                     .asciiArmoredPgpPublicKey("""
 * mQENBFtP0doBCADF+joTiXWKVuP8kJt3fgpBSjT9h8ezMfKA4aXZctYLx5wslWQl
 * bB7Iu2ezkECNzoEeU7WxUe8a61pMCh9cisS9H5mB2K2uM4Jnf8tgFeXn3akJDVo0
 * oR1IC+Dp9mXbRSK3MAvKkOwWlG99sx3uEdvmeBRHBOO+grchLx24EThXFOyP9Fk6
 * V39j6xMjw4aggLD15B4V0v9JqBDdJiIYFzszZDL6pJwZrzcP0z8JO4rTZd+f64bD
 * Mpj52j/pQfA8lZHOaAgb1OrthLdMrBAjoDjArV4Ek7vSbrcgYWcI6BhsQrFoxKdX
 * 83TZKai55ZCfCLIskwUIzA1NLVwyzCS+fSN/ABEBAAG0KCJUZXN0IEF0dGVzdG9y
 * IiA8ZGFuYWhvZmZtYW5AZ29vZ2xlLmNvbT6JAU4EEwEIADgWIQRfWkqHt6hpTA1L
 * uY060eeM4dc66AUCW0/R2gIbLwULCQgHAgYVCgkICwIEFgIDAQIeAQIXgAAKCRA6
 * 0eeM4dc66HdpCAC4ot3b0OyxPb0Ip+WT2U0PbpTBPJklesuwpIrM4Lh0N+1nVRLC
 * 51WSmVbM8BiAFhLbN9LpdHhds1kUrHF7+wWAjdR8sqAj9otc6HGRM/3qfa2qgh+U
 * WTEk/3us/rYSi7T7TkMuutRMIa1IkR13uKiW56csEMnbOQpn9rDqwIr5R8nlZP5h
 * MAU9vdm1DIv567meMqTaVZgR3w7bck2P49AO8lO5ERFpVkErtu/98y+rUy9d789l
 * +OPuS1NGnxI1YKsNaWJF4uJVuvQuZ1twrhCbGNtVorO2U12+cEq+YtUxj7kmdOC1
 * qoIRW6y0+UlAc+MbqfL0ziHDOAmcqz1GnROg
 * =6Bvm
 *                     """)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Binary Authorization Attestor Kms
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.kms.inputs.CryptoKeyVersionTemplateArgs;
 * import com.pulumi.gcp.kms.KmsFunctions;
 * import com.pulumi.gcp.kms.inputs.GetKMSCryptoKeyVersionArgs;
 * import com.pulumi.gcp.containeranalysis.Note;
 * import com.pulumi.gcp.containeranalysis.NoteArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityHintArgs;
 * import com.pulumi.gcp.binaryauthorization.Attestor;
 * import com.pulumi.gcp.binaryauthorization.AttestorArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.AttestorAttestationAuthorityNoteArgs;
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
 *         var keyring = new KeyRing("keyring", KeyRingArgs.builder()
 *             .name("test-attestor-key-ring")
 *             .location("global")
 *             .build());
 * 
 *         var crypto_key = new CryptoKey("crypto-key", CryptoKeyArgs.builder()
 *             .name("test-attestor-key")
 *             .keyRing(keyring.id())
 *             .purpose("ASYMMETRIC_SIGN")
 *             .versionTemplate(CryptoKeyVersionTemplateArgs.builder()
 *                 .algorithm("RSA_SIGN_PKCS1_4096_SHA512")
 *                 .build())
 *             .build());
 * 
 *         final var version = KmsFunctions.getKMSCryptoKeyVersion(GetKMSCryptoKeyVersionArgs.builder()
 *             .cryptoKey(crypto_key.id())
 *             .build());
 * 
 *         var note = new Note("note", NoteArgs.builder()
 *             .name("test-attestor-note")
 *             .attestationAuthority(NoteAttestationAuthorityArgs.builder()
 *                 .hint(NoteAttestationAuthorityHintArgs.builder()
 *                     .humanReadableName("Attestor Note")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var attestor = new Attestor("attestor", AttestorArgs.builder()
 *             .name("test-attestor")
 *             .attestationAuthorityNote(AttestorAttestationAuthorityNoteArgs.builder()
 *                 .noteReference(note.name())
 *                 .publicKeys(AttestorAttestationAuthorityNotePublicKeyArgs.builder()
 *                     .id(version.applyValue(getKMSCryptoKeyVersionResult -> getKMSCryptoKeyVersionResult).applyValue(version -> version.applyValue(getKMSCryptoKeyVersionResult -> getKMSCryptoKeyVersionResult.id())))
 *                     .pkixPublicKey(AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs.builder()
 *                         .publicKeyPem(version.applyValue(getKMSCryptoKeyVersionResult -> getKMSCryptoKeyVersionResult).applyValue(version -> version.applyValue(getKMSCryptoKeyVersionResult -> getKMSCryptoKeyVersionResult.publicKeys()[0].pem())))
 *                         .signatureAlgorithm(version.applyValue(getKMSCryptoKeyVersionResult -> getKMSCryptoKeyVersionResult).applyValue(version -> version.applyValue(getKMSCryptoKeyVersionResult -> getKMSCryptoKeyVersionResult.publicKeys()[0].algorithm())))
 *                         .build())
 *                     .build())
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
 * Attestor can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/attestors/{{name}}`
 * 
 * * `{{project}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Attestor can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:binaryauthorization/attestor:Attestor default projects/{{project}}/attestors/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:binaryauthorization/attestor:Attestor default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:binaryauthorization/attestor:Attestor default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:binaryauthorization/attestor:Attestor")
public class Attestor extends com.pulumi.resources.CustomResource {
    /**
     * A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
     * Structure is documented below.
     * 
     */
    @Export(name="attestationAuthorityNote", refs={AttestorAttestationAuthorityNote.class}, tree="[0]")
    private Output<AttestorAttestationAuthorityNote> attestationAuthorityNote;

    /**
     * @return A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
     * Structure is documented below.
     * 
     */
    public Output<AttestorAttestationAuthorityNote> attestationAuthorityNote() {
        return this.attestationAuthorityNote;
    }
    /**
     * A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The resource name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Attestor(String name) {
        this(name, AttestorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Attestor(String name, AttestorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Attestor(String name, AttestorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:binaryauthorization/attestor:Attestor", name, args == null ? AttestorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Attestor(String name, Output<String> id, @Nullable AttestorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:binaryauthorization/attestor:Attestor", name, state, makeResourceOptions(options, id));
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
    public static Attestor get(String name, Output<String> id, @Nullable AttestorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Attestor(name, id, state, options);
    }
}
