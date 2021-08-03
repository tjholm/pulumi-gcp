// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Partially mask a string by replacing a given number of characters with a fixed character.
        /// Masking can start from the beginning or end of the string.
        /// Structure is documented below.
        /// </summary>
        [Input("characterMaskConfig")]
        public Input<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfigGetArgs>? CharacterMaskConfig { get; set; }

        /// <summary>
        /// Pseudonymization method that generates deterministic encryption for the given input. Outputs a base64 encoded representation of the encrypted output. Uses AES-SIV based on the RFC [https://tools.ietf.org/html/rfc5297](https://tools.ietf.org/html/rfc5297).
        /// Structure is documented below.
        /// </summary>
        [Input("cryptoDeterministicConfig")]
        public Input<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfigGetArgs>? CryptoDeterministicConfig { get; set; }

        /// <summary>
        /// Replaces an identifier with a surrogate using Format Preserving Encryption (FPE) with the FFX mode of operation; however when used in the `content.reidentify` API method, it serves the opposite function by reversing the surrogate back into the original identifier. The identifier must be encoded as ASCII. For a given crypto key and context, the same identifier will be replaced with the same surrogate. Identifiers must be at least two characters long. In the case that the identifier is the empty string, it will be skipped. See [https://cloud.google.com/dlp/docs/pseudonymization](https://cloud.google.com/dlp/docs/pseudonymization) to learn more.
        /// Note: We recommend using CryptoDeterministicConfig for all use cases which do not require preserving the input alphabet space and size, plus warrant referential integrity.
        /// Structure is documented below.
        /// </summary>
        [Input("cryptoReplaceFfxFpeConfig")]
        public Input<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoReplaceFfxFpeConfigGetArgs>? CryptoReplaceFfxFpeConfig { get; set; }

        /// <summary>
        /// Replace each input value with a given value.
        /// Structure is documented below.
        /// </summary>
        [Input("replaceConfig")]
        public Input<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigGetArgs>? ReplaceConfig { get; set; }

        /// <summary>
        /// Replace each matching finding with the name of the info type.
        /// </summary>
        [Input("replaceWithInfoTypeConfig")]
        public Input<bool>? ReplaceWithInfoTypeConfig { get; set; }

        public PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationGetArgs()
        {
        }
    }
}
